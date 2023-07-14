<script lang="ts">
  import {toasts} from 'svelte-toasts'
  import {onMount} from 'svelte'
  import {createEventDispatcher} from 'svelte'
  import Error from '$components/Error.svelte'
  import Spinner from '$components/Spinner.svelte'

  const dispatch = createEventDispatcher()

  export let gameClient

  onMount(async () => {
    let unsub = await gameClient.collection('guesses').subscribe('*', handleStreamedGuess)
    return () => unsub()
  })

  const handleStreamedGuess = function(payload) {
    newGuesses = newGuesses.concat(payload.record)
    if (payload.record.content === '*hint') {
      dispatch('hint')
    } else if (payload.record.correct) {
      dispatch('solve')
    } else {
      toasts.add({
        title: 'Incorrect guess',
        description: payload.record.content,
        duration: 10000,
        type: 'error',
      })
    }
  }

  let previousGuesses = gameClient.collection('guesses').
          getList(1, 50, {sort: 'created'}).
          then(({items}) => items)

  let newGuesses = []
</script>

<style lang="scss">
  @use "../style/colours";

  ul, li {
    list-style-type: none;
    margin: 0;
    padding: 0;
  }

  .date {
    font-size: small;
    color: colours.$text-secondary;
  }

  .hint {
    color: colours.$brand;
  }

  .correct {
    color: colours.$success;
  }
</style>

<ul>
  {#await previousGuesses}
    <Spinner/>
  {:then previousGuesses}
    {#each previousGuesses.concat(newGuesses) as guess (guess.id)}
      <li class="{(guess.content === '*hint' && 'hint') || (guess.correct && 'correct') || 'wrong'}">
        <p class="guess">{(guess.content === '*hint' && '💡 A hint was unlocked') || guess.content}</p>
        <p class="date">{guess.created}</p>
      </li>
    {:else}
      <p>You haven't guessed anything yet!</p>
    {/each}
  {:catch error}
    <Error error={error}/>
  {/await}
</ul>