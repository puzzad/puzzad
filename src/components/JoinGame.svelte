<script>
  import tickets from '$assets/tickets.webp'
  import {goto} from '$app/navigation'
  import Spinner from '$components/Spinner.svelte'
  import {getGameClient} from '$lib/api'
  import {toasts} from 'svelte-toasts'

  let code = ''
  let loading = false

  const handleGameCode = () => {
    const normalisedCode = code.toLowerCase().replace(/[^a-z.]/g, '')

    loading = true

    getGameClient(normalisedCode)
    .then(() => goto(`/games/${normalisedCode}`))
    .catch(() => {
      loading = false
      code = ''
      toasts.add({
        title: 'Error',
        description: 'That game code is not valid',
        duration: 10000,
        type: 'error',
      })
    })
  }
</script>

<style lang="scss">
  @use "../style/colours";


  form {
    position: relative;
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
    color: colours.$text-on-brand;

    .overlay {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: rgba(0, 0, 0, 0.5);
      border-radius: 10px;
    }
  }

  h3 {
    border-bottom: 0;
  }
</style>

<form on:submit|preventDefault={handleGameCode} style="background-image: url('{tickets}')">
  {#if loading}
    <div class="overlay">
      <Spinner/>
    </div>
  {/if}
  <h3>Got a game code?</h3>
  <input type="text" placeholder="revolving.magenta.ocelot" bind:value={code}>
  <input type="submit" value="Play!">
</form>
