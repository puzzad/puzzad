<script lang="ts">
  import {getGameClient} from '$lib/db'
  import {pb} from '$lib/auth.ts'

  export let puzzle = 0
  export let gameID = ''

  let guess = ''
  let checking = false

  const submit = async () => {
    checking = true
    pb.collection("guesses").create({content: guess, puzzle: puzzle, game: gameID})
        .then(() => guess = '')
        .finally(() => checking = false)
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