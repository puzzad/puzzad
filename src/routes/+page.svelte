<script>
  import {goto} from '$app/navigation'
  import {title} from '$lib/title.ts'
  import DatabasePage from '$lib/components/DatabasePage.svelte'

  let code = ''

  const handleGameCode = () => {
    goto(`/games/${code.toLowerCase().replaceAll(/[^a-z-]/g, '')}`)
  }
  title.set('Puzzad')
</script>
<style>
  main {
    grid-template-columns: 50% 40%;
    gap: 2em;
  }

  @media (min-width: 480px) {
    main {
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

  #hero {
    display: flex;
    flex-direction: column;
    align-items: end;
    justify-content: space-between;
    background-image: url('../lib/assets/hero.jpg');
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    position: relative;
  }

  #hero h3 {
    left: 30%;
    right: 0.5em;
    top: 0.5em;
    font-size: xx-large;
    color: #000000;
    font-weight: bold;
    text-align: right;
    text-shadow: 0 0 5px #ffffff;
  }

  #hero a {
    bottom: 1em;
    right: 1em;
    font-size: large;
    font-weight: bold;
    text-shadow: 0 0 5px #000000;
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
<main>
  <section id="hero">
    <h3>Ready to start an adventure?</h3>
    <a href="/adventures">Let's go &raquo;</a>
  </section>
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
</main>