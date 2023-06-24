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
  import {parsePuzzleContent} from '$lib/puzzle'

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
          then((gameData) => {
            title.set(`Puzzad: ${gameData.adventure.name}: ${gameData.title}`)
            return gameData
          }).
          then((gameData) => Promise.all([
              gameData,
            parsePuzzleContent(data.game, gameData.storage_slug, gameData.content)
          ])).
          then(([gameData, parsedContent]) => {
            gameData.sections = parsedContent
            return gameData
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

<style lang="scss">
  @use "../src/style/colours";
  @use "../src/style/fonts";

  #container {
    display: grid;
    grid-template-columns: 70% 30%;
    gap: var(--small-space);
  }

  #sidebar {
    border-left: 1px solid colours.$border;
    padding: 0 var(--small-space);
  }

  #sidebar h3 {
    font-family: fonts.$header;
    font-weight: 800;
    text-transform: uppercase;
    border: 0;
  }
</style>

{#await gameData}
  <Spinner/>
{:then gameData}
  <h2>{gameData.adventure.name}: {gameData.title}</h2>

  <div id="container">
    <div id="main-content">
      {@html gameData.sections.story || ''}
      {@html gameData.sections.puzzle || ''}
      <PuzzleAnswer gameCode={data.game} puzzle={gameData.id}></PuzzleAnswer>
      <Hints gameCode={data.game} puzzleId={gameData.id} bind:this={hints}></Hints>
    </div>
    <div id="sidebar">
      {#if gameData.sections.tip}
        <h3>Information</h3>
        {@html gameData.sections.tip}
        <hr>
      {/if}
      <h3>Guesses</h3>
      <ul>
        <li>
          Not the answer (3min ago)
        </li>
        <li>
          Help :( (6min ago)
        </li>
      </ul>
    </div>
  </div>

  {#if solved}
    <VictoryDialog game={data.game} finished={gameData.next === null} on:next={reload}></VictoryDialog>
  {/if}

  <GuessMonitor game={data.game} on:hint={() => {hints && hints.refresh()}}
                on:solve={() => {solved = true}}></GuessMonitor>
{/await}