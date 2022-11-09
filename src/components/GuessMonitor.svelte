<script>
  import {params} from '@roxi/routify'
  import {toasts} from 'svelte-toasts'
  import {getRealTimeClient} from '$lib/db'
  import {onDestroy, onMount} from 'svelte'

  export let onHint = () => {}
  export let onSolve = () => {}
  let realTimeClient

  onMount(async () => {
    realTimeClient = await getRealTimeClient($params.code)
    await realTimeClient.channel('public:guesses:game=eq.' + $params.code).on('postgres_changes', {
      event: 'INSERT',
      schema: 'public',
      table: 'guesses',
      filter: 'game=eq.' + $params.code,
    }, handleStreamedGuess).subscribe()
  })

  onDestroy(function() {
    if (realTimeClient) {
      realTimeClient.disconnect()
    }
  })

  const handleStreamedGuess = function(payload) {
    if (payload.record.puzzle.toString() === $params.puzzle) {
      if (payload.record.content === '*hint') {
        onHint()
      } else if (payload.record.correct) {
        onSolve()
      } else {
        toasts.add({
          title: 'Incorrect guess',
          description: payload.record.content,
          duration: 10000,
          type: 'error',
        })
      }
    }
  }
</script>