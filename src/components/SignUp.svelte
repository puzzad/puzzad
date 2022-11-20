<script lang="ts">
  import {supabase} from '$lib/db'
  import {toasts} from 'svelte-toasts'

  let email
  let password

  export const authAction = async () => {
    let {error} = await supabase.auth.signUp({email, password})
    if (error) {
      toasts.add({
        title: 'Error',
        description: error.message,
        duration: 10000,
        type: 'error',
      })
    } else {
      toasts.add({
        title: 'Success',
        description: 'An email has been sent to you for verification!',
        duration: 10000,
        type: 'success',
      })
    }
  }
</script>

<style>
  .disclaimer {
    font-size: small;
    color: #999999;
  }
</style>

<h2>Sign up</h2>
<form class="basic" on:submit|preventDefault={authAction}>
  <label for="email">E-mail:</label>
  <input
      id="email"
      type="email"
      name="email"
      bind:value={email}
      required
  />
  <label for="password">Password:</label>
  <input
      id="password"
      type="password"
      name="password"
      bind:value={password}
      required
  />
  <button type="submit">
    Sign up
  </button>
</form>
<p class="disclaimer">
  If you create an account, we will only use your e-mail address to send you information
  related to your account (e.g. password reset links and receipts). We won't use it for marketing
  purposes unless you explicitly sign up for our mailing list, and we'll never sell your data or
  give it to anyone who will abuse it.
</p>
