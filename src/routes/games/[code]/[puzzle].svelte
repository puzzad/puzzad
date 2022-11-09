<script>
  import {getGameClient, getRealTimeClient} from '$lib/db'
  import Spinner from '$comps/Spinner.svelte'
  import {title} from '$lib/title.ts'
  import {goto, params} from '@roxi/routify'
  import {onDestroy} from 'svelte'
  import {toasts} from 'svelte-toasts'
  import Hints from '$comps/Hints.svelte'
  import PuzzleContent from '$comps/PuzzleContent.svelte'
  import PuzzleAnswer from '$comps/PuzzleAnswer.svelte'
  import VictoryDialog from '$comps/VictoryDialog.svelte'

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

  $: if ($params.puzzle && $params.code) {
    load()
  }

</script>

{#if initial}
  <Spinner/>
{:else}
  <h2>{data.adventure.name}: {data.title}</h2>

  <PuzzleContent gameCode={$params.code} storageSlug={data.storage_slug} content={data.content}></PuzzleContent>
  <PuzzleAnswer gameCode={$params.code} puzzle={$params.puzzle}></PuzzleAnswer>
  <Hints gameCode={$params.code} puzzleId={$params.puzzle} bind:this={hints}></Hints>

  {#if solved}
    <VictoryDialog next={data.next}></VictoryDialog>
  {/if}
{/if}
