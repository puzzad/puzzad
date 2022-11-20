<script lang="ts">
  import {toasts} from 'svelte-toasts'
  import {getRealTimeClient} from '$lib/db'
  import {onDestroy, onMount} from 'svelte'
  import {createEventDispatcher} from 'svelte'

  let realTimeClient
  const dispatch = createEventDispatcher()

  export let game

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
</script>