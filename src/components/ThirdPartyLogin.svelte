<script lang="ts">
  import FaGoogle from 'svelte-icons/fa/FaGoogle.svelte'
  import FaDiscord from 'svelte-icons/fa/FaDiscord.svelte'
  import FaTwitch from 'svelte-icons/fa/FaTwitch.svelte'
  import {toasts} from 'svelte-toasts'
  import {goto} from '$app/navigation'
  import {loginOauth} from '$lib/api'

  const handleOauthLogin = async (name) => {
    loginOauth(name)
    .then(response => {
      console.log(response)
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

<style lang="scss">
  @use "../style/colours";

  section {
    display: flex;
    flex-direction: column;
    gap: var(--small-space);
  }

  .icon {
    width: 32px;
    height: 32px;
    display: inline-block;
    vertical-align: middle;
    margin-right: 0.5em;
  }

  .socialicons {
    display: flex;
    flex-direction: column;
    gap: 10px;

    button {
      width: 100%;
      max-width: 400px;
      display: flex;
      justify-content: center;
      align-items: center;

      @media (max-width: 1000px) {
        max-width: 100%;
      }
    }
  }

  .disclaimer {
    color: colours.$text-secondary;
    font-size: small;
  }
</style>

<section>
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