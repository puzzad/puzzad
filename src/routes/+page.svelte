<script lang="ts">
  import {goto} from '$app/navigation'
  import {title} from '$lib/title.ts'
  import DatabasePage from '$components/DatabasePage.svelte'
  import StartAdventureCallout from '../components/StartAdventureCallout.svelte'

  let code = ''

  const handleGameCode = () => {
    goto(`/games/${code.toLowerCase().replace(/[^a-z-]/g, '')}`)
  }
  title.set('Puzzad')
</script>
<style>
  #container {
    grid-template-columns: 50% calc(50% - 2em);
    gap: 2em;
  }

  @media (min-width: 480px) {
    #container {
      display: grid;
    }
  }

  @media (max-width: 480px) {
    #play {
      margin-bottom: 2em;
    }

    #hero {
      margin-bottom: 2em;
    }
  }

  section {
    border-radius: 15px;
    padding: 1em;
    margin: 0;
  }

  h3 {
    margin: 0;
  }

  #play {
    background-color: #39254D;
    display: flex;
    flex-direction: column;
    align-items: stretch;
    gap: 0.5em;
  }

  #play h3, #play input {
    text-align: center;
  }

  form {
    display: contents;
  }

  #intro {
    grid-column-start: 1;
    grid-column-end: 3;
    background-color: #222222;
  }

  #intro :global(h3) {
    border-bottom: 2px solid #e3bc5e;
    padding: 0;
    margin: 0;
    font-weight: bold;
    font-size: large;
  }
</style>

<h2>Puzzad: Puzzle Adventures</h2>
<div id="container">
  <div id="hero">
    <StartAdventureCallout></StartAdventureCallout>
  </div>
  <section id="play">
    <h3>Got a game code?</h3>
    <form on:submit|preventDefault={handleGameCode}>
      <input type="text" placeholder="revolving-magenta-ocelot" bind:value={code}>
      <input type="submit" value="Play!">
    </form>
  </section>
  <section id="intro">
    <DatabasePage page="homepage"></DatabasePage>
  </section>
</div>