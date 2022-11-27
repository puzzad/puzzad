<script lang="ts">
  import {supabase} from '$lib/db.ts'
  import SvelteHcaptcha from 'svelte-hcaptcha'
  import Spinner from '$components/Spinner.svelte'
  import {toasts} from 'svelte-toasts'

  let name, nameError, email, emailError, message, messageError, success, captchaToken = '', supabaseToken, loggedIn
  let loading = false
  const handleSubmit = () => {
    fetch(import.meta.env.VITE_SUPABASE_URL + '/mail/contact', {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + supabaseToken,
        'Content-Type': 'Application/json',
      },
      body: JSON.stringify({
        'token': captchaToken,
        'name': name,
        'email': email,
        'message': message,
      }),
    }).then(response => {
      if (!response.ok) {
        return response.json().then((j) => {
          throw j.error
        })
      }
      return response
    }).then(_ => {
      success = true
      loading = false
    }).catch(error => {
      toasts.add({
        title: 'Error submitting form',
        description: error,
        duration: 10000,
        type: 'error',
      })
    })
  }
  supabase.auth.getSession().then(response => {
    if (response.data.session && response.data.session.user.email) {
      loggedIn = true
      email = response.data.session.user.email
      supabaseToken = response.data.session.access_token
    } else {
      loggedIn = false
    }
    loading = false
  })
  const captchaSuccess = (token) => {
    captchaToken = token.detail.token
  }
  const captchaFail = () => {
    captchaToken = ''
  }
</script>
<style>
  form {
    display: grid;
  }

  input {
    margin-bottom: 1em;
  }

  button {
    margin-top: 1em;
    margin-right: 0;
  }

  section :global(iframe) {
    margin-top: 1em;
  }
</style>
<section>
  <h2>Contact us</h2>
  {#if loading}
    <Spinner/>
  {:else if success}
    <p>Your message has been sent</p>
  {:else}
    <form method="POST" on:submit|preventDefault={handleSubmit}>
      <label for="name">Name
        {#if nameError}<span class="error">The name field is required</span>{/if}
      </label>
      <input id="name" name="name" bind:value={name}/>
      {#if !loggedIn}
        <label for="email">Email
          {#if emailError}<span class="error">The email field is required</span>{/if}
        </label>
        <input type="email" id="email" name="email" bind:value={email}/>
      {/if}
      <label for="message">Query
        {#if messageError}<span class="error">The query field is required</span>{/if}
      </label>
      <textarea id="message" name="message" bind:value={message}></textarea>
      {#if !loggedIn}
        <SvelteHcaptcha sitekey={import.meta.env.VITE_HCAPTCHA_SITE_KEY}
                        theme="dark"
                        reCaptchaCompa="false"
                        on:success={captchaSuccess}
                        on:error={captchaFail}
                        on:expired={captchaFail}
        />
      {/if}
      <button type="submit" disabled={!(loggedIn || captchaToken !== '')}>Submit</button>
    </form>
  {/if}
</section>