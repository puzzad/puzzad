<script lang="ts">
  import {goto} from '$app/navigation'
  import {getGameClient} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import Hints from '$components/Hints.svelte'
  import PuzzleAnswer from '$components/PuzzleAnswer.svelte'
  import VictoryDialog from '$components/VictoryDialog.svelte'
  import {parsePuzzleContent} from '$lib/puzzle'
  import GuessList from '$components/GuessList.svelte'

  let previousGuesses
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
            parsePuzzleContent(data.game, gameData.storage_slug, gameData.content),
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

  function handleHintsOpened() {
    previousGuesses.removeAttribute('open')
  }

  let guessesOpened = false
  $: if (guessesOpened) {
    hints.close()
  }
</script>

<style lang="scss">
  @use "../src/style/colours";
  @use "../src/style/fonts";

  #container {
    display: grid;
    grid-template-columns: 70% 30%;
    gap: var(--small-space);

    @media (max-width: 800px) {
      grid-template-columns: 1fr;
    }
  }

  #sidebar {
    border-left: 1px solid colours.$border;
    padding: 0 var(--small-space);

    @media (min-width: 800px) {
      position: sticky;
      top: 0;
      max-height: 100vh;
      display: flex;
      flex-direction: column;
    }

    @media (max-width: 800px) {
      border-top: 1px solid colours.$border;
      border-left: 0;
      padding: var(--small-space) 0;
    }
  }

  #sidebar h3 {
    font-family: fonts.$header;
    font-weight: 800;
    text-transform: uppercase;
    border: 0;
  }

  .guesses {
    max-height: 80vh;
    overflow-y: scroll;
    margin-top: 1em;
    flex-shrink: 1;
  }

  .hints {
    display: contents;
  }

  details {
    display: contents;

    summary {
      margin-top: 1em;
    }
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
    </div>
    <div id="sidebar">
      {#if gameData.sections.tip}
        <h3>Information</h3>
        {@html gameData.sections.tip}
        <hr>
      {/if}
      <h3>Guess</h3>
      <PuzzleAnswer gameCode={data.game} puzzle={gameData.id}></PuzzleAnswer>
      <details bind:open={guessesOpened} bind:this={previousGuesses}>
        <summary>See previous guesses</summary>
        <div class="guesses">
          <GuessList game={data.game} puzzle={gameData.id} on:hint={() => {hints && hints.refresh()}}
                     on:solve={() => {solved = true}}/>
        </div>
      </details>
      <h3>Hints</h3>
      <div class="hints">
        <Hints gameCode={data.game} puzzleId={gameData.id} bind:this={hints} on:open={handleHintsOpened}></Hints>
      </div>
    </div>
  </div>

  {#if solved}
    <VictoryDialog game={data.game} finished={gameData.next === null} on:next={reload}></VictoryDialog>
  {/if}

{/await}