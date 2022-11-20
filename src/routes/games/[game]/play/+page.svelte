<script lang="ts">
  import {goto} from '$app/navigation'
  import {getGameClient} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import Hints from '$components/Hints.svelte'
  import PuzzleContent from '$components/PuzzleContent.svelte'
  import PuzzleAnswer from '$components/PuzzleAnswer.svelte'
  import VictoryDialog from '$components/VictoryDialog.svelte'
  import GuessMonitor from '$components/GuessMonitor.svelte'

  let hints
  let solved = false

  export let data

  const load = async () =>
      getGameClient(data.game).
          then((client) => client.from('games').
              select('puzzle (id, title, content, next, storage_slug, adventure (name))').
              eq('code', data.game).
              throwOnError().
              single(),
          ).
          then(({data: {puzzle}}) => puzzle).
          then((data) => {
            title.set(`Puzzad: ${data.adventure.name}: ${data.title}`)
            return data
          }).
          catch(() => goto(`/games/${data.game}`))
  
  const reload = () => {
    solved = false
    gameData = load()
  }

  let gameData

  $: if (data.game) {
    reload()
  }

</script>

{#await gameData}
  <Spinner/>
{:then gameData}
  <h2>{gameData.adventure.name}: {gameData.title}</h2>

  <PuzzleContent gameCode={data.game} storageSlug={gameData.storage_slug} content={gameData.content}></PuzzleContent>
  <PuzzleAnswer gameCode={data.game} puzzle={gameData.id}></PuzzleAnswer>
  <Hints gameCode={data.game} puzzleId={gameData.id} bind:this={hints}></Hints>

  {#if solved}
    <VictoryDialog game={data.game} finished={gameData.next === null} on:next={reload}></VictoryDialog>
  {/if}

  <GuessMonitor game={data.game} on:hint={() => {hints && hints.refresh()}}
                on:solve={() => {solved = true}}></GuessMonitor>
{/await}