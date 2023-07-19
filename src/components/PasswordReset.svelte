<script lang="ts">
  import {client} from '$lib/api.ts'
  import {toasts} from 'svelte-toasts'

  export let code

  let password
  const reset = () => {
    client.collection("users").confirmPasswordReset(code, password, password)
    .then(_ => {
      password = ""
      toasts.add({
        title: 'Success',
        description: 'Password reset!',
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
    console.log(code)
  }
</script>
<style lang="scss">
  @use '../style/colours';
  div {
    display: grid;
    grid-template-columns: auto 40%;
    grid-gap: 2em;
    align-items: start;

    @media (max-width: 1000px) {
      grid-gap: 0;
      grid-template-columns: auto;
    }
  }
</style>

<div>
  <form class="basic" on:submit|preventDefault={reset}>
    <label for="password">New Password</label>
    <input
        id="password"
        type="password"
        name="password"
        bind:value={password}
        placeholder="MySuperSecurePassword123!"
        required
    />
    <button>
      Submit
    </button>
  </form>
  <section>
    <p>
      To protect themselves from evil doers, adventurers must have
      a password that:
    </p>
    <ul>
      <li>is at least 8 characters long</li>
      <li>contains an uppercase letter</li>
      <li>contains a lowercase letter</li>
      <li>contains a number</li>
      <li>contains a special character</li>
    </ul>
    <p>
      We highly recommend a capable assistant like <a href="https://bitwarden.com/">BitWarden</a>
      for generating and securely storing passwords.
    </p>
  </section>
</div>