<script>
  import {getGameClient} from '$lib/db'
  import Spinner from '$comps/Spinner.svelte'
  import {title} from '$lib/title.ts'
  import {goto, params} from '@roxi/routify'
  import Hints from '$comps/Hints.svelte'
  import PuzzleContent from '$comps/PuzzleContent.svelte'
  import PuzzleAnswer from '$comps/PuzzleAnswer.svelte'
  import VictoryDialog from '$comps/VictoryDialog.svelte'
  import GuessMonitor from '$comps/GuessMonitor.svelte'

  let data = {}
  let solved = false
  let initial = true
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

<GuessMonitor onHint={() => { hints && hints.refresh() }} onSolve={() => {solved = true}}></GuessMonitor>