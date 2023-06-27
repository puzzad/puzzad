<script lang="ts">
  import {loginNative} from '$lib/auth'
  import {toasts} from 'svelte-toasts'

  let email
  let password

  const login = async () => {
    loginNative(email, password)
        .then(response => {
          toasts.add({
        title: 'Success',
        description: 'Login success.',
        duration: 10000,
        type: 'success',
      })
      //return goto('/')
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

<style lang="scss">
  section {
    display: flex;
    flex-direction: column;
    gap: var(--small-space);
  }

</style>

<section>
  <p>
    Login using your Puzzad account:
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
    <label for="password">Password</label>
    <input
        id="password"
        type="password"
        name="password"
        bind:value={password}
        placeholder="MySuperSecurePassword123!"
        required
    />
    <button type="submit">
      Login
    </button>
  </form>
  <p>Don't have a Puzzad account and don't want to use an external service? <a href="/signup">Sign up here!</a></p>
</section>