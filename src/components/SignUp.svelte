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

<style lang="scss">
  @use '../style/colours';

  .disclaimer {
    grid-column: 1 / 3;
    margin: 10px 0;
    color: colours.$text-secondary;
    font-size: small;
  }

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

  section {
    margin: 0;
    border-left: 1px solid colours.$border;
    padding: 0 var(--small-space);

    @media (max-width: 1000px) {
      border-left: none;
      padding: var(--small-space) 0;
      border-top: 1px solid colours.$border;
    }
  }
</style>

<h2>Sign up</h2>
<div>
  <form class="basic" on:submit|preventDefault={authAction}>
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
      Sign up
    </button>
    <p class="disclaimer">
      If you create an account, we will only use your e-mail address to send you information
      related to your account (e.g. password reset links and receipts). We won't use it for marketing
      purposes unless you explicitly sign up for our mailing list, and we'll never sell your data or
      give it to anyone who will abuse it. See the <a href="/privacy">privacy</a> page for further
      information.
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
