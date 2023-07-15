<script lang="ts">
  import {goto} from '$app/navigation'
  import RandomText from '$components/RandomText.svelte'
  import {Fireworks} from '@fireworks-js/svelte'
  import {createEventDispatcher} from 'svelte'

  const dispatch = createEventDispatcher()

  export let finished
  export let game

  const goToNextPuzzle = async function() {
    dispatch('next')
  }

  const goToGamePage = async function() {
    goto(`/games/${game}`)
  }

  const congratsMessages = [
    'Congratulations!',
    'Well done!',
    'You did it!',
    'Huzzah!',
    'Good job!',
    'Nice work!',
  ]
</script>

<style lang="scss">
  @use "../style/colours";

  dialog {
    position: fixed;
    padding: var(--small-space) var(--large-space);
    top: 30vh;
    left: 30vh;
    right: 30vh;
    background: colours.$success;
    color: colours.$text-on-brand;
    border: 1px solid colours.$border;
    text-align: center;
    z-index: 1001;
    animation: show 200ms linear;
    display: flex;
    flex-direction: column;
    gap: var(--small-space);
  }

  h3 {
    border-bottom: 0;
  }

  @keyframes show {
    from {
      transform: scale(0.001);
    }
    to {
      transform: scale(1);
    }
  }

  dialog button {
    width: 100%;
  }

  /*noinspection CssUnusedSymbol*/
  :global(.fireworks) {
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    position: fixed;
    z-index: 1000;
    background: rgba(0, 0, 0, 0.7);
  }
</style>

<dialog open="open">
  <h3>
    <RandomText options={congratsMessages}></RandomText>
  </h3>
  {#if finished}
    <p>You have completed the adventure!</p>
    <button on:click={() => goToGamePage()}>Continue &raquo;</button>
  {:else}
    <p>You have completed this step in the adventure!</p>
    <button on:click={() => goToNextPuzzle()}>Next puzzle &raquo;</button>
  {/if}
</dialog>

<Fireworks autostart={true} class="fireworks"/>