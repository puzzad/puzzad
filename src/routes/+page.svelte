<script lang="ts">
  import {goto} from '$app/navigation'
  import {title} from '$lib/title.ts'
  import preview from '$assets/preview.webp'

  let code = ''

  const handleGameCode = () => {
    goto(`/games/${code.toLowerCase().replace(/[^a-z.]/g, '')}`)
  }
  title.set('Puzzad')
</script>

<style lang="scss">
  @use "../style/colours";

  .code {
    display: grid;
    grid-template-columns: 55% 45%;
    padding: 3rem;
    grid-column-gap: 3rem;

    form {
      max-width: 800px;
      background-color: colours.$border;
      border-radius: 10px;
      box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.5);
      padding: var(--large-space) var(--small-space);
      display: flex;
      flex-direction: column;
      gap: var(--small-space);
      align-items: stretch;
      text-align: center;
    }

    h3 {
      border-bottom: 0;
    }

    @media (max-width: 1200px) {
      display: flex;
      flex-direction: column-reverse;
      gap: var(--large-space);
    }
  }

  .hero {
    display: grid;
    grid-template-columns: 45% 55%;
    padding: 3rem;
    grid-column-gap: 3rem;

    @media (max-width: 1200px) {
      display: flex;
      flex-direction: column;
      gap: var(--large-space);
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
        font-size: x-large;
        padding: 1em;
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

  .blurb {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  h2 {
    font-size: xxx-large;
    text-transform: uppercase;
    margin-bottom: 0.7em;
    white-space: pre-line;
    line-height: 1em;

    @media (max-width: 1200px) {
      white-space: normal;
    }
  }

  p {
    font-size: x-large;
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
    <img src="{preview}"
         alt="Preview of several puzzles from Conspiracy: a fridge, a section of door mat, a ticket machine, and a locked book">
    <button on:click={() => goto('/adventures/Conspiracy')}>
      Play &lsquo;Conspiracy&rsquo; for free now &raquo;
    </button>
  </div>
</div>

<div class="code">
  <form on:submit|preventDefault={handleGameCode}>
    <h3>Got a game code?</h3>
    <input type="text" placeholder="revolving-magenta-ocelot" bind:value={code}>
    <input type="submit" value="Play!">
  </form>
  <div class="blurb">
    <h2>
      Help A Friend
      On Their Adventure
    </h2>
    <p>
      When you start a puzzle adventure, you get a unique game code like <code>revolving-magneta-ocelot</code>.
      Share this code with your friends to let them join your adventure! You'll progress as a team, and be able
      to see what one another are guessing in real time!
    </p>
  </div>
</div>