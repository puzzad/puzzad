<script lang="ts">
  import FaGoogle from 'svelte-icons/fa/FaGoogle.svelte'
  import FaDiscord from 'svelte-icons/fa/FaDiscord.svelte'
  import FaTwitch from 'svelte-icons/fa/FaTwitch.svelte'
  import {supabase} from '$lib/db'
  import {toasts} from 'svelte-toasts'
  import {goto} from '$app/navigation'

  const handleOauthLogin = async (name) => {
    supabase.auth.signInWithOAuth({provider: name})
    .then(response => {
      if (response.error) {
        return Promise.reject(response.error)
      }
      toasts.add({
        title: 'Success',
        description: 'Login success.',
        duration: 10000,
        type: 'success',
      })
      return goto('/')
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

<style>
  section {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-radius: 20px;
    padding: 1em;
    text-align: center;
    background-color: #222222;
  }

  .icon {
    width: 32px;
    height: 32px;
    display: inline-block;
    vertical-align: middle;
    margin-right: 0.5em;
  }

  button {
    display: block;
    background-color: var(--links);
    border: 0;
    width: 100%;
    color: black;
    border-radius: 10px;
    margin: 1em 0;
    padding: 0.5em;
  }

  button:hover {
    text-decoration: none;
    filter: brightness(1.2);
  }

  .disclaimer {
    font-size: small;
    color: #999999;
  }

  h3 {
    margin-top: 0;
    font-size: larger;
  }
</style>

<section>
  <h3>From another land?</h3>
  <p>
    Login with an external service:
  </p>
  <div class="socialicons">
    <button on:click={() => handleOauthLogin('google')}>
      <span class="icon">
        <FaGoogle/>
      </span>
      Google
    </button>
    <button on:click={() => handleOauthLogin('discord')}>
      <span class="icon">
        <FaDiscord/>
      </span>
      Discord
    </button>
    <button on:click={() => handleOauthLogin('twitch')}>
      <span class="icon">
        <FaTwitch/>
      </span>
      Twitch
    </button>
  </div>
  <p class="disclaimer">
    When signing in with a third-party provider, Puzzad will receive your profile details such as your
    name and e-mail address. We can't do anything else with your account like read your e-mail, check
    your mentions, or walk your dog.
  </p>
</section>