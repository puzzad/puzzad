<script lang="ts">
  import {supabase} from '$lib/db'
  import {goto} from '$app/navigation'
  import {toasts} from 'svelte-toasts'

  let email
  let password

  const login = async () => {
    let {error} = await supabase.auth.signInWithPassword({email, password})
    if (error) {
      toasts.add({
        title: 'Error',
        description: error.message,
        duration: 10000,
        type: 'error',
      })
    } else {
      await goto('/')
      toasts.add({
        title: 'Success',
        description: 'Login success.',
        duration: 10000,
        type: 'success',
      })
    }
  }
</script>

<style>
  h3 {
    margin-top: 0;
    font-size: larger;
  }

  section {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border: 1px solid #666666;
    border-radius: 20px;
    padding: 1em;
    text-align: center;
  }

  input, button {
    width: 100%;
    height: 3em;
    margin-bottom: 1em;
  }

  label {
    text-transform: lowercase;
    font-variant: small-caps;
    text-align: left;
    width: 100%;
  }

  form {
    display: contents;
  }
</style>

<section>
  <h3>Puzzad native?</h3>
  <p>
    Login using your Puzzad account:
  </p>
  <form on:submit|preventDefault={login}>
    <label for="email">E-mail</label>
    <input
        id="email"
        type="email"
        name="email"
        bind:value={email}
        required
    />
    <label for="password">Password</label>
    <input
        id="password"
        type="password"
        name="password"
        bind:value={password}
        required
    />
    <button type="submit">
      Login
    </button>
  </form>
  <p class="signup">
    Don't have a Puzzad account and don't want to use an external service? <a href="/signup">Sign up here!</a>
  </p>
</section>