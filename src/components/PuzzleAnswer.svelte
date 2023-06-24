<script lang="ts">
  import {getGameClient} from '$lib/db'

  export let puzzle = 0
  export let gameCode = ''

  let guess = ''
  let checking = false

  const submit = async () => {
    checking = true
    await getGameClient(gameCode).
        then((gc) => gc.from('guesses').insert({content: guess, puzzle: puzzle, game: gameCode}).throwOnError()).
        then(() => guess = '').
        finally(() => checking = false)
  }
</script>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
</style>

<form on:submit|preventDefault={submit}>
  <input id="guess" type="text" bind:value={guess} disabled={checking} placeholder="Enter a guess...">
  <input type="submit" value="Submit" disabled={checking}>
</form>
