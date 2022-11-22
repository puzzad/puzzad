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
    grid-column: 1 / 3;
    margin: 2em 0;
    font-size: small;
    color: #999999;
  }

  @media (min-width: 480px) {
    div {
      display: grid;
      grid-template-columns: auto 40%;
      grid-gap: 2em;
    }

    .disclaimer {
      text-align: center;
      border-top: 1px dotted #999999;
      padding-top: 2em;
    }
  }

  form {
    align-content: center;
  }

  section {
    background-color: #222222;
    padding: 0.5em 1em;
    margin: 0;
    border-radius: 15px;
  }
</style>

<h2>Sign up</h2>
<div>
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
  <p class="disclaimer">
    If you create an account, we will only use your e-mail address to send you information
    related to your account (e.g. password reset links and receipts). We won't use it for marketing
    purposes unless you explicitly sign up for our mailing list, and we'll never sell your data or
    give it to anyone who will abuse it.
  </p>
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
