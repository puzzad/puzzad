<script>
  import AdventureLogo from '$comps/AdventureLogo.svelte'
  import {getGameClient} from '$lib/db.ts'
  import Spinner from '$comps/Spinner.svelte'
  import {title} from '$lib/title.ts'
  import {goto} from '@roxi/routify'
  import {onMount} from 'svelte'
  import {params} from '@roxi/routify'

  $: code = $params.code
  let data = {}
  let initial = true
  let stats
  let displayError
  let gc

  onMount(async function() {
    try {
      gc = await getGameClient(code)
    } catch (e) {
      displayError = e.message || e
      initial = false
      return
    }

    let {data: game, error} =
        await gc.from('games').
            select('status, puzzle, startTime, endTime, adventures ( name, description)').
            eq('code', code)

    if (game[0]?.adventures) {
      data = game[0]
      stats = gc.rpc('getstats', {gamecode: code}).then(({data: stats, error}) => {
        if (error) {
          return []
        } else {
          return stats
        }
      }).then((rows) => {
        let lastTime = data.startTime
        rows.forEach((r) => {
          r.time = delta(lastTime, r.solvetime)
          lastTime = r.solvetime
        })
        return rows
      })
      title.set('Puzzad: ' + game[0]?.adventures.name + ' - ' + code)
    } else {
      title.set('Puzzad: Game not found')
      error = 'Unable to find game'
    }
    if (error) {
      displayError = error
    }
    initial = false
  })

  const handleStartAdventure = async function() {
    let {data: puzzle} = await gc.rpc('startadventure')
    console.log(puzzle)
    if (puzzle) {
      $goto('/games/[code]/[puzzle]', {code: code, puzzle: puzzle})
    }
  }

  const handleContinueAdventure = async function() {
    $goto('/games/[code]/[puzzle]', {code: code, puzzle: data.puzzle})
  }

  const delta = function(start, end) {
    const diff = ((new Date(end)).getTime() - (new Date(start)).getTime()) / 1000
    const text = (value, label) => {
      const rounded = Math.floor(value)
      if (rounded === 1) {
        return `${rounded} ${label}`
      } else if (rounded > 1) {
        return `${rounded} ${label}s`
      } else {
        return ''
      }
    }

    return [
      text(diff % 60, 'second'),
      text((diff / 60) % 60, 'minute'),
      text((diff / 3600) % 24, 'hour'),
      text(diff / 86400, 'day'),
    ].filter((t) => t !== '').reverse().join(', ')
  }
</script>

<style>
  code {
    display: block;
    background-image: url('../../../assets/code-bg.png');
    background-color: transparent;
    color: black;
    text-align: center;
    background-size: contain;
    background-repeat: no-repeat;
    background-position: center;
    padding: 70px 0;
    font-size: xxx-large;
    font-family: monospace;
    margin: 1em;
  }

  button {
    width: 100%;
    margin: 1em 0;
    font-size: x-large;
    padding: 0.5em;
  }
</style>

{#if initial}
  <Spinner/>
{:else if displayError || !data}
  <h1>Error finding game</h1>
  <p>{displayError}</p>
{:else}
  <h1>
    <AdventureLogo bind:name={data.adventures.name}></AdventureLogo>
  </h1>
  {#if data.status === 'EXPIRED'}
    <p>Congratulations! You finished the adventure!</p>
    <p>You took {delta(data.startTime, data.endTime)}!</p>
    {#await stats then puzzles}
      <table class="stats">
        <thead>
        <tr>
          <th>Puzzle</th>
          <th>Time</th>
          <th>Hints</th>
        </tr>
        </thead>
        <tbody>
        {#each puzzles as puzzle}
          <tr>
            <td>{puzzle.title}</td>
            <td class="time">{puzzle.time}</td>
            <td>{puzzle.hints}</td>
          </tr>
        {/each}
        </tbody>
      </table>
    {/await}
  {:else if data.status === 'PAID'}
    <p>
      You've not yet started your adventure! Remember, it's dangerous to go alone.
      Take this if you want to recruit others to help you:
    </p>
    <code>
      {code}
    </code>
    <p>
      Once you're ready, press below to go to the first puzzle in the adventure!
    </p>
    <button on:click={() => handleStartAdventure()}>
      Begin the adventure!
    </button>
  {:else if data.status === 'ACTIVE'}
    <p>This adventure is already in progress!</p>
    <button on:click={() => handleContinueAdventure()}>
      Go to the current puzzle
    </button>
  {/if}
{/if}
