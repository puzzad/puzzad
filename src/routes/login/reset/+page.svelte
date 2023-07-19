<script lang="ts">
  import {client} from '$lib/api'
  import {toasts} from 'svelte-toasts'

  let email

  const login = async () => {
    client.collection('users').requestPasswordReset(email)
    .then(_ => {
      email = ""
      toasts.add({
        title: 'Success',
        description: 'Password reset email sent.',
        duration: 10000,
        type: 'success',
      })
    }).catch(err => {
      toasts.add({
        title: 'Error',
        description: err.message,
        duration: 10000,
        type: 'error',
      })
    })
  }
</script>
<section>
  <p>
    If you've forgotten your password, you can initiate a password reset from here.
  </p>
  <form class="basic" on:submit|preventDefault={login}>
    <label for="email">E-mail</label>
    <input
        id="email"
        type="email"
        name="email"
        bind:value={email}
        placeholder="puzzler@example.com"
        required
    />
    <button type="submit">
      Initiate Reset
    </button>
  </form>
</section>