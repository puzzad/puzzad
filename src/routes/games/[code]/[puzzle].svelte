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

  let hints
  let solved = false

  const load = async (code, puzzle) =>
      getGameClient(code).
          then((client) => client.from('puzzles').
              select('title, content, next, storage_slug, adventure (name)').
              eq('id', puzzle).
              throwOnError().
              single(),
          ).
          then(({data}) => data).
          then((data) => {
            title.set(`Puzzad: ${data.adventure.name}: ${data.title}`)
            return data
          }).
          catch(() => $goto('/games/[code]', {code}))

  $: if ($params.code || $params.puzzle) { solved = false }

</script>

{#await load($params.code, $params.puzzle)}
  <Spinner/>
{:then data}
  <h2>{data.adventure.name}: {data.title}</h2>

  <PuzzleContent gameCode={$params.code} storageSlug={data.storage_slug} content={data.content}></PuzzleContent>
  <PuzzleAnswer gameCode={$params.code} puzzle={$params.puzzle}></PuzzleAnswer>
  <Hints gameCode={$params.code} puzzleId={$params.puzzle} bind:this={hints}></Hints>

  {#if solved}
    <VictoryDialog next={data.next}></VictoryDialog>
  {/if}

  <GuessMonitor on:hint={() => {hints && hints.refresh()}} on:solve={() => {solved = true}}></GuessMonitor>
{/await}