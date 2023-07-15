<script lang="ts">
  import SvelteHcaptcha from 'svelte-hcaptcha'
  import Spinner from '$components/Spinner.svelte'
  import {toasts} from 'svelte-toasts'
  import {title} from '$lib/title'
  import {currentUser, client} from '$lib/api'

  title.set('Puzzad: Contact')

  let name, nameError, email, emailError, message, messageError, success, captchaToken = '', loggedIn
  let loading = false
  const handleSubmit = () => {
    fetch(import.meta.env.VITE_SUPABASE_URL + 'wom/contact', {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + client.authStore.token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        'token': captchaToken,
        'name': name,
        'email': email,
        'message': message,
      })
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
  if ($currentUser) {
    loggedIn = true
    email = $currentUser?.email ?? ""
  } else {
    loggedIn = false
  }
  const captchaSuccess = (token) => {
    captchaToken = token.detail.token
  }
  const captchaFail = () => {
    captchaToken = ''
  }
</script>
<style>
  section :global(iframe) {
    margin-top: 1em;
  }

  #captcha {
    grid-column: 2;
  }
</style>
<section>
  <h2>Contact us</h2>
  {#if loading}
    <Spinner/>
  {:else if success}
    <p>Your message has been sent</p>
  {:else}
    <form method="POST" class="basic" on:submit|preventDefault={handleSubmit}>
      <label for="name">Name
        {#if nameError}<span class="error">The name field is required</span>{/if}
      </label>
      <input id="name" name="name" type="text" bind:value={name}/>
      {#if !loggedIn || email===""}
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
        <div id="captcha">
          <SvelteHcaptcha sitekey={import.meta.env.VITE_HCAPTCHA_SITE_KEY}
                          theme="dark"
                          reCaptchaCompa="false"
                          on:success={captchaSuccess}
                          on:error={captchaFail}
                          on:expired={captchaFail}
          />
        </div>
      {/if}
      <button type="submit" disabled={!(loggedIn || captchaToken !== '')}>Submit</button>
    </form>
  {/if}
</section>