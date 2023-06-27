<script>
  import ThirdPartyLogin from '$components/ThirdPartyLogin.svelte'
  import NativeLogin from '$components/NativeLogin.svelte'
  import AlreadyLoggedIn from '$components/AlreadyLoggedIn.svelte'
  import Spinner from '$components/Spinner.svelte'
  import {currentUser} from '$lib/auth'
  import {title} from '$lib/title'

  title.set('Puzzad: Login')
</script>

<style lang="scss">
  @use "../../style/colours";

  @media (min-width: 1000px) {
    #root {
      display: grid;
      grid-template-columns: 30% auto;
      grid-gap: 1em;

      #left {
        border-right: 1px solid colours.$border;
        padding-right: var(--small-space);
      }

      #right {
        padding-left: var(--small-space);
      }
    }
  }

  @media (max-width: 1000px) {
    #root {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;

      #left {
        border-bottom: 1px solid colours.$border;
        padding-bottom: var(--small-space);
        margin-bottom: var(--small-space);
      }
    }
  }
</style>
{#if !$currentUser}
  <h2>Login</h2>
  <div id="root">
    <div id="left"><ThirdPartyLogin></ThirdPartyLogin></div>
    <div id="right"><NativeLogin></NativeLogin></div>
  </div>
{:else}
  <AlreadyLoggedIn></AlreadyLoggedIn>
{/if}