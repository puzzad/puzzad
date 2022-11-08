<script>
  import RandomText from '$comps/RandomText.svelte'
  import {getGameClient, getRealTimeClient} from '$lib/db'
  import Spinner from '$comps/Spinner.svelte'
  import {title} from '$lib/title.ts'
  import {goto, params} from '@roxi/routify'
  import {onDestroy, onMount} from 'svelte'
  import {toasts} from 'svelte-toasts'
  import {Confetti} from 'svelte-confetti'
  import Hints from '$comps/Hints.svelte'
  import PuzzleContent from '$comps/PuzzleContent.svelte'
  import PuzzleAnswer from '$comps/PuzzleAnswer.svelte'

  let data = {}
  let solved = false
  let initial = true
  let realTimeClient = null
  let hints

  title.set('Puzzad: Loading...')

  const load = () => {
    reset()
    getGameClient($params.code).
        then((client) => client.from('puzzles').
            select('title, content, next, storage_slug, adventure (name)').
            eq('id', $params.puzzle).
            throwOnError(),
        ).
        then(checkQueryResults).
        then((puzzle) => data = puzzle).
        then(() => {
          title.set('Puzzad: ' + data.adventure.name + ': ' + data.title)
          startMonitoringGuesses()
          initial = false
        }).
        catch(() => $goto('/games/[code]', {code: $params.code}))
  }

  const reset = () => {
    console.log('Resetting...')
    initial = true
    solved = false
  }

  const checkQueryResults = ({data, error}) => {
    if (error) {
      throw error
    } else if (data.length === 0) {
      throw 'Puzzle not found'
    } else {
      return data[0]
    }
  }

  onDestroy(function() {
    if (realTimeClient) {
      realTimeClient.disconnect()
    }
  })

  const startMonitoringGuesses = async function() {
    if (realTimeClient) {
      return
    }

    realTimeClient = await getRealTimeClient($params.code)
    await realTimeClient.channel('public:guesses:game=eq.' + $params.code).on('postgres_changes', {
      event: 'INSERT',
      schema: 'public',
      table: 'guesses',
      filter: 'game=eq.' + $params.code,
    }, handleStreamedGuess).subscribe()
  }

  const handleStreamedGuess = function(payload) {
    if (payload.record.puzzle.toString() === $params.puzzle) {
      if (payload.record.content === '*hint') {
        hints.refresh()
      } else if (payload.record.correct) {
        solved = true
      } else {
        toasts.add({
          title: 'Incorrect guess',
          description: payload.record.content,
          duration: 10000,
          type: 'error',
        })
      }
    }
  }

  const goToNextPuzzle = async function() {
    $goto('/games/[code]/[puzzle]', {code: $params.code, puzzle: data.next})
  }

  const goToGamePage = async function() {
    $goto('/games/[code]', {code: $params.code})
  }

  $: if ($params.puzzle && $params.code) {
    load()
  }

  const congratsMessages = [
    'Congratulations!',
    'Well done!',
    'You did it!',
    'Huzzah!',
    'Good job!',
    'Nice work!',
  ]
</script>

<style>
  dialog {
    position: fixed;
    padding: 3em;
    top: 30vh;
    left: 30vh;
    right: 30vh;
    background: linear-gradient(to top, #106310, #3C8E2B);
    border-radius: 10px;
    text-align: center;
    z-index: 1001;
  }

  dialog[open] {
    animation: show 200ms linear;
  }

  @keyframes show {
    from {
      transform: scale(0.001);
    }
    to {
      transform: scale(1);
    }
  }

  dialog h3 {
    margin: 0 0 1em 0;
  }

  dialog p {
    margin: 0 0 2em 0;
  }

  dialog button {
    width: 100%;
    text-transform: uppercase;
    font-weight: bold;
    font-size: x-large;
    font-variant: all-small-caps;
  }

  .confetti-bg {
    position: fixed;
    top: -50px;
    left: 0;
    bottom: 0;
    right: 0;
    display: flex;
    justify-content: center;
    overflow: hidden;
    background-color: rgba(0, 0, 0, 0.8);
    z-index: 1000;
  }
</style>

{#if initial}
  <Spinner/>
{:else}
  <h2>{data.adventure.name}: {data.title}</h2>

  <PuzzleContent gameCode={$params.code} storageSlug={data.storage_slug} content={data.content}></PuzzleContent>
  <PuzzleAnswer gameCode={$params.code} puzzle={$params.puzzle}></PuzzleAnswer>
  <Hints gameCode={$params.code} puzzleId={$params.puzzle} bind:this={hints}></Hints>

  <dialog open={solved}>
    <h3>
      <RandomText options={congratsMessages}></RandomText>
    </h3>
    {#if data.next}
      <p>You have completed this step in the adventure!</p>
      <button on:click={() => goToNextPuzzle()}>Next puzzle &raquo;</button>
    {:else}
      <p>You have completed the adventure!</p>
      <button on:click={() => goToGamePage()}>Continue &raquo;</button>
    {/if}
  </dialog>

  {#if solved}
    <div class="confetti-bg">
      <Confetti x={[-5, 5]} y={[0, 0.1]} delay={[0, 2000]} infinite duration="5000" amount="400"
                fallDistance="110vh"/>
    </div>
  {/if}
{/if}
