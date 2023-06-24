<script lang="ts">
  import {goto} from '$app/navigation'
  import {title} from '$lib/title.ts'
  import preview from '$assets/preview.webp'

  import DatabasePage from '$components/DatabasePage.svelte'
  import StartAdventureCallout from '$components/StartAdventureCallout.svelte'

  let code = ''

  const handleGameCode = () => {
    goto(`/games/${code.toLowerCase().replace(/[^a-z-]/g, '')}`)
  }
  title.set('Puzzad')
</script>

<style lang="scss">
  .hero {
    display: grid;
    grid-template-columns: 45% 55%;
    padding: 3rem;
    grid-column-gap: 3rem;

    @media (max-width: 1200px) {
      grid-template-columns: 100%;
    }

    .blurb {
      display: flex;
      flex-direction: column;
      justify-content: center;

      h2 {
        font-family: var(--header-font);
        font-size: var(--font-size-xxl);
        margin-bottom: 0.7em;
        white-space: pre-line;

        @media (max-width: 1200px) {
          white-space: normal;
        }
      }

      p {
        font-size: var(--font-size-xl);
      }
    }

    .preview {
      display: grid;
      align-items: center;
      justify-content: center;

      img {
        grid-row: 1;
        grid-column: 1;
      }

      button {
        cursor: pointer;
        background-color: var(--accent);
        color: white;
        border: 0;
        border-radius: 0.5em;
        font-size: var(--font-size-xl);
        padding: 0.6em;
        text-decoration: underline;
        max-width: 80%;
        justify-self: center;

        grid-row: 1;
        grid-column: 1;

        @media (max-width: 800px) {
          grid-row: 2;
          min-width: 100%;
          max-width: 100%;
        }
      }
    }
  }
</style>

<div class="hero">
  <div class="blurb">
    <h2>
      Start Your Own
      Puzzle Adventure
    </h2>
    <p>
      Crack codes, solve puzzles, and tackle problems to progress through the story.
      Work alone, or collaborate with friends. It's just like an escape room, but
      with no time limit and without having to move!
    </p>
  </div>
  <div class="preview">
    <img src="{preview}" alt="Preview of several puzzles from Conspiracy: a picture of a fridge, a section of door mat, a ticket machine, and a locked book">
    <button>
      Play &lsquo;Conspiracy&rsquo; for free now &raquo;
    </button>
  </div>
</div>

<!--<div id="container">-->
<!--  <div id="hero">-->
<!--    <StartAdventureCallout></StartAdventureCallout>-->
<!--  </div>-->
<!--  <section id="play">-->
<!--    <h3>Got a game code?</h3>-->
<!--    <form on:submit|preventDefault={handleGameCode}>-->
<!--      <input type="text" placeholder="revolving-magenta-ocelot" bind:value={code}>-->
<!--      <input type="submit" value="Play!">-->
<!--    </form>-->
<!--  </section>-->
<!--  <section id="intro" class="text">-->
<!--    <DatabasePage page="homepage"></DatabasePage>-->
<!--  </section>-->
<!--</div>-->