<script lang="ts">
  import Spinner from '$components/Spinner.svelte'
  import SvelteHcaptcha from 'svelte-hcaptcha'
  import {toasts} from 'svelte-toasts'
  import {client} from '$lib/api.ts'

  let loading = true
  let useAccount = true
  let supabaseToken = ''
  let email = ''
  let captchaToken = ''
  let success = false
  let needConfirm = false


  if (client.authStore.isValid) {
    useAccount = true
      if (client.authStore.isAdmin || client.authStore.isAuthRecord) {
          email = client.authStore.model.email
          useAccount = true
          supabaseToken = client.authStore.token
      } else {
        useAccount = false
      }
      loading = false
  } else {
    useAccount = false
    loading = false
  }

  const handleSubmit = () => {
    loading = true
    client.send('/wom/subscribe', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: email,
        captcha: captchaToken,
      })
    }).then((r) => {
      needConfirm = r.NeedConfirm
      success = true
      loading = false
    }).catch((e) => {
      toasts.add({
        title: 'Error joining mailing list',
        description: e,
        duration: 10000,
        type: 'error',
      })
      loading = false
    })
  }

  const captchaSuccess = (token) => {
    captchaToken = token.detail.token
  }
  const captchaFail = () => {
    captchaToken = ''
  }
</script>

<style lang="scss">
  a {
    font-size: small;
    margin-bottom: 10px;
  }

  section {
    display: flex;
    flex-direction: column;
  }

  form {
    display: flex;
    flex-direction: column;
    max-width: 500px;
    margin-top: 10px;
  }

  label {
    display: none;
  }

  section :global(iframe) {
    margin-top: 1em;
  }
</style>

<section>
  {#if loading}
    <Spinner></Spinner>
  {:else if success}
    <h3>Subscription request received!</h3>
    {#if needConfirm}
      <p>Nearly there - we've sent you an e-mail with a confirmation link.</p>
    {:else}
      <p>All done! You're successfully signed up to our mailing list!</p>
    {/if}
  {:else}
    <h3>Receive updates from the town crier?</h3>
    <p>
      Eager to hear when new adventures are available? Join our mailing list!
      We'll only send you the occasional announcement, and you can unsubscribe at any time.
    </p>
    <form on:submit|preventDefault={handleSubmit}>
      <label for="email">E-mail:</label>
      <input type="email" id="email" placeholder="email@example.com" bind:value={email} readonly={useAccount}
             disabled={useAccount}>
      {#if useAccount}
        <a href="#" on:click|preventDefault={() => useAccount = false}>Want to use a different address?</a>
      {/if}
      {#if !useAccount}
        <SvelteHcaptcha sitekey={import.meta.env.VITE_HCAPTCHA_SITE_KEY}
                        on:success={captchaSuccess}
                        on:error={captchaFail}
                        on:expired={captchaFail}
        />
      {/if}
      <input type="submit" value="Subscribe" disabled={!(useAccount || captchaToken !== '')}>
    </form>
  {/if}
</section>