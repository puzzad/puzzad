<script lang="ts">
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import SvelteHcaptcha from 'svelte-hcaptcha'
  import {toasts} from 'svelte-toasts'

  let loading = true
  let useAccount = true
  let supabaseToken = ''
  let email = ''
  let captchaToken = ''
  let success = false
  let needConfirm = false

  supabase.auth.getSession().then(response => {
    if (response.data.session && response.data.session.user.email) {
      useAccount = true
      email = response.data.session.user.email
      supabaseToken = response.data.session.access_token
    } else {
      useAccount = false
    }
    loading = false
  })

  const handleSubmit = () => {
    loading = true
    fetch(import.meta.env.VITE_SUPABASE_URL + '/mail/subscribe', {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + supabaseToken,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: email,
        captcha: captchaToken,
      }),
    }).then((response) => {
      if (!response.ok) {
        return response.json().then((j) => {
          throw j.error
        })
      }
      return response.json()
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
    console.log(token)
    captchaToken = token.detail.token
  }
</script>

<style>
  a {
    font-size: small;
  }

  section {
    border-radius: 15px;
    padding: 1em;
    margin: 0;
    background-color: #39254D;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  form {
    display: contents;
  }

  h3 {
    margin: 0;
  }

  p {
    margin-bottom: 0;
    text-align: center;
  }

  input, input[type=submit] {
    margin: 1em 0 0 0;
    width: 50%;
    padding: 1em;
  }

  label {
    display: none;
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
        <SvelteHcaptcha sitekey={import.meta.env.VITE_HCAPTCHA_SITE_KEY} theme="dark"
                        on:success={captchaSuccess}></SvelteHcaptcha>
      {/if}
      <input type="submit" value="Subscribe" disabled={!(useAccount || captchaToken !== '')}>
    </form>
  {/if}
</section>