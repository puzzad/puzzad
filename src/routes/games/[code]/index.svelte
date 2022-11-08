<script>
  import AdventureLogo from '$comps/AdventureLogo.svelte'
  import {getGameClient} from '$lib/db'
  import {formatDuration} from '$lib/time'
  import Spinner from '$comps/Spinner.svelte'
  import {title} from '$lib/title.ts'
  import {goto} from '@roxi/routify'
  import {params} from '@roxi/routify'
  import Error from '$comps/Error.svelte'

  let stats

  let game = getGameClient($params.code).
      then((gc) =>
          gc.from('games').
              select('status, puzzle, startTime, endTime, adventures ( name, description)').
              eq('code', $params.code).
              throwOnError().
              single(),
      ).
      then(({data: game}) => game).
      then((game) => {
        title.set(`Puzzad: ${game.adventures.name} - ${$params.code}`)
        if (game.status === 'EXPIRED') {
          stats = fetchStats(game.startTime)
        }
        return game
      })

  const fetchStats = async (startTime) => getGameClient($params.code).
      then((gc) => gc.rpc('getstats', {gamecode: $params.code}).throwOnError()).
      then(({data: stats}) => stats).
      then((rows) => {
        let lastTime = startTime
        rows.forEach((r) => {
          r.time = formatDuration(lastTime, r.solvetime)
          lastTime = r.solvetime
        })
        return rows
      })

  const handleStartAdventure = async () =>
      getGameClient($params.code).
          then((gc) => gc.rpc('startadventure').throwOnError()).
          then(({data: puzzle}) => {
            if (puzzle) {
              return $goto('/games/[code]/[puzzle]', {code: $params.code, puzzle: puzzle})
            }
          })

  const handleContinueAdventure = async function(puzzle) {
    $goto('/games/[code]/[puzzle]', {code: $params.code, puzzle})
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

{#await game}
  <Spinner/>
{:then data}
  <h1>
    <AdventureLogo name={data.adventures.name}></AdventureLogo>
  </h1>
  {#if data.status === 'EXPIRED'}
    <p>Congratulations! You finished the adventure!</p>
    <p>You took {formatDuration(data.startTime, data.endTime)}!</p>
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
      {$params.code}
    </code>
    <p>
      Once you're ready, press below to go to the first puzzle in the adventure!
    </p>
    <button on:click={() => handleStartAdventure()}>
      Begin the adventure!
    </button>
  {:else if data.status === 'ACTIVE'}
    <p>This adventure is already in progress!</p>
    <button on:click={() => handleContinueAdventure(data.puzzle)}>
      Go to the current puzzle
    </button>
  {/if}
{:catch error}
  {#if error.message === 'Invalid team code'}
    <h1>Game not found</h1>
    <p>
      The adventure you are seeking does not seem to exist.
      If you have been given a code, please double check your spelling.
      If you registered the adventure yourself, you can find it listed in <a href="/games">My Games</a>.
    </p>
  {:else}
    <Error error={error}></Error>
  {/if}
{/await}
