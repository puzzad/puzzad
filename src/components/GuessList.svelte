<script lang="ts">
  import {toasts} from 'svelte-toasts'
  import {getGameClient, getRealTimeClient, supabase} from '$lib/db'
  import {onDestroy, onMount} from 'svelte'
  import {createEventDispatcher} from 'svelte'
  import Error from '$components/Error.svelte'
  import Spinner from '$components/Spinner.svelte'

  let realTimeClient
  const dispatch = createEventDispatcher()

  export let game
  export let puzzle

  onMount(async () => {
    realTimeClient = await getRealTimeClient(game)
    await realTimeClient.channel('public:guesses:game=eq.' + game).on('postgres_changes', {
      event: 'INSERT',
      schema: 'public',
      table: 'guesses',
      filter: 'game=eq.' + game,
    }, handleStreamedGuess).subscribe()
  })

  onDestroy(function() {
    if (realTimeClient) {
      realTimeClient.disconnect()
    }
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

  let previousGuesses = getGameClient(game).
      then((gc) => gc.from('guesses').
          select('id,created_at,content,correct').
          eq('game', game).
          eq('puzzle', puzzle).
          order('created_at', {ascending: true}).
          throwOnError()).
      then(({data}) => data)

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
        <p class="date">{guess.created_at}</p>
      </li>
    {/each}
  {:catch error}
    <Error error={error}/>
  {/await}
</ul>