<script lang="ts">
  import Spinner from '$components/Spinner.svelte'
  import {toasts} from 'svelte-toasts'
  import {goto} from '$app/navigation'

  export let data

  $: if (data.code) {
    fetch(import.meta.env.VITE_API_URL + 'wom/unsubscribe', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        token: data.code,
      }),
    }).then((response) => {
      if (!response.ok) {
        return response.json().then((j) => {
          throw j.error
        })
      }
      return response
    }).then(() => {
      toasts.add({
        title: 'Success',
        description: 'Unsubscribed from mailing list.',
        duration: 10000,
        type: 'success',
      })
      return goto('/')
    }).catch((e) => {
      toasts.add({
        title: 'Error unsubscribing',
        description: e,
        duration: 10000,
        type: 'error',
      })
      return goto('/')
    })
  }
</script>

<Spinner></Spinner>