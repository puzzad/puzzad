<script lang="ts">

  import {toasts} from 'svelte-toasts'

  export let answer: string

  let guess = ''
  let correct = false

  function handleGuess() {
    if (guess.toLowerCase() === answer.toLowerCase()) {
      toasts.success('Correct!')
      correct = true
      guess += ' ✅'
    } else {
      toasts.error('That\'s not right. Try again!')
      guess = ''
    }
  }
</script>

<style lang="scss">
  @use "../style/colours";

  aside {
    display: flex;
    flex-direction: column;

    @media (min-width: 1000px) {
      flex-direction: row;
      align-items: center;

      h4 {
        writing-mode: vertical-rl;
        transform: rotate(180deg);
        border: 0;
      }
    }

    div {
      @media (min-width: 1000px) {
        text-align: center;
        border-left: 1px solid colours.$border;
        margin-left: 10px;
        padding-left: var(--small-space);
      }
    }

    form {
      display: flex;
      gap: 20px;

      input[type=text] {
        flex-grow: 1;
      }
    }
  }
</style>

<aside>
  <h4>Example</h4>
  <div>
    <slot/>
    <form on:submit|preventDefault={handleGuess}>
      <input type="text" placeholder="Put your answer here" bind:value={guess} disabled={correct}>
      <input type="submit" value="Guess" disabled={correct}>
    </form>
  </div>
</aside>