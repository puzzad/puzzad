<script lang="ts">
  import RandomText from '$components/RandomText.svelte'
  import {getGameClient} from '$lib/db'
  import {createEventDispatcher} from 'svelte'

  const dispatch = createEventDispatcher()

  export let gameCode = ''
  export let puzzleId = 0
  let details: HTMLElement

  export const close = () => {
    details.removeAttribute('open')
  }

  export const refresh = async () => {
    const gameClient = await getGameClient(gameCode)
    const {data, error} = await gameClient.rpc('gethints', {puzzleid: puzzleId, gamecode: gameCode})
    if (error) {
      throw error
    }
    hints = data
  }

  const request = async function(id) {
    const gameClient = await getGameClient(gameCode)
    await gameClient.rpc('requesthint', {puzzleid: puzzleId, gamecode: gameCode, hintid: id})
  }

  let hints = []

  $: if (gameCode && puzzleId) {
    refresh()
  }

  let opened = false

  $: if (opened) {
    dispatch('open')
  }

  const hintMessages = [
    'Need a hint? Browse our extensive collection below!',
    'Not sure where to go? Try one of our finest hints!',
    'Get your hints here! Freshly plucked from the hint tree!',
    'Psst... Can I interest you in a hint?',
    'We\'ve got the finest hints in all the land. Don\'t believe me? Try one for free!',
    'Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?',
  ]
</script>

<style lang="scss">
  details {
    display: contents;
  }

  .container {
    flex-shrink: 1;
    max-height: 80vh;
    overflow-y: auto;
  }

  div p.locked {
    filter: blur(5px);
  }

  div p {
    transition: filter 1s;
  }

  .hint {
    display: grid;

    * {
      grid-column: 1;
      grid-row: 1;
      justify-self: center;
      align-self: center;
    }

    button {
      z-index: 10;
    }
  }
</style>

<details bind:open={opened} bind:this={details}>
  <summary>
    <RandomText options={hintMessages}></RandomText>
  </summary>
  <div class="container">
    {#each hints as hint (hint.id)}
      <h4>{hint.title}</h4>
      <div class="hint">
        <p class="{hint.locked ? 'locked' : ''}">
          {#if hint.locked}
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
            labore et dolore magna aliqua.
          {:else}
            {hint.message}
          {/if}
        </p>
        {#if hint.locked}
          <button on:click={() => request(hint.id)}>Reveal this hint</button>
        {/if}
      </div>
    {/each}
  </div>
</details>