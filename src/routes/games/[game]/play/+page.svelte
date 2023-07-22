<script lang="ts">
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import {onMount} from 'svelte'
  import {getGameClient} from '$lib/api.ts'
  import {goto} from '$app/navigation'
  import PuzzleAnswer from '$components/PuzzleAnswer.svelte'
  import GuessList from '$components/GuessList.svelte'
  import Hints from "$components/Hints.svelte";
  import VictoryDialog from "$components/VictoryDialog.svelte";

  let hints
  let root
  let solved = false
  let details = []

  export let data

  onMount(() => {
    const observer = new MutationObserver(records => {
      records.flatMap(r => Array.from(r.addedNodes.values())).filter(n => n.tagName === 'DETAILS').forEach(e => {
        details.push(e)
        e.addEventListener('toggle', () => {
          if (e.open) {
            details.forEach(d => {
              if (d !== e) {
                d.removeAttribute('open')
              }
            })
          }
        })
      })
    })

    observer.observe(root, { subtree: true, childList: true })

    return () => observer.disconnect()
  })

  const load = async () => {
    let gameClient = await getGameClient(data.game)
    return await gameClient.collection("games").
        getList(1, 1, {expand: "adventure,puzzle"}).
        then(({items}) => items[0]).
        then((gameData) => {
          if (gameData.end !== "") {
            solved = true
          }
          title.set(`Puzzad: ${gameData.expand.adventure.name}: ${gameData.expand?.puzzle?.title ?? ""}`)
          gameData.client = gameClient
          return gameData
        })
  }

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

  details.info[open] summary {
    display: none;
  }
</style>

<div id="root" bind:this={root}>
  {#await gameData}
    <Spinner/>
  {:then gameData}
    {#if solved}
      <VictoryDialog game={data.game} finished={gameData.expand.puzzle.next === ""} on:next={reload}></VictoryDialog>
    {:else if gameData.puzzle !== ""}
      <h2>{gameData.expand.adventure.name}: {gameData.expand.puzzle.title}</h2>

      <div id="container">
        <div id="main-content">
          {@html gameData.expand.puzzle.story || ''}
          {@html gameData.expand.puzzle.puzzle || ''}
        </div>
        <div id="sidebar">
          {#if gameData.expand.puzzle.information}
            <h3>Information</h3>
            <details open class="info">
              <summary>Show information</summary>
              {@html gameData.expand.puzzle.information}
            </details>
            <hr>
          {/if}
          <h3>Guess</h3>
          <PuzzleAnswer gameClient={gameData.client} gameID={gameData.id} puzzle={gameData.expand.puzzle.id}></PuzzleAnswer>
          <details>
            <summary>See previous guesses</summary>
            <div class="guesses">
              <GuessList gameClient={gameData.client} on:hint={() => {hints && hints.refresh()}}
                         on:solve={() => {solved = true}}/>
            </div>
          </details>
          <h3>Hints</h3>
          <div class="hints">
            <Hints gameClient={gameData.client} gameCode={data.game} puzzleId={gameData.id} bind:this={hints}></Hints>
          </div>
        </div>
      </div>
    {/if}
  {/await}
</div>