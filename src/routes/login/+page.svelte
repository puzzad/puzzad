<script>
  import ThirdPartyLogin from '$components/ThirdPartyLogin.svelte'
  import NativeLogin from '$components/NativeLogin.svelte'
  import AlreadyLoggedIn from '$components/AlreadyLoggedIn.svelte'
  import {session} from '$lib/auth'
  import {title} from '$lib/title'

  title.set('Puzzad: Login')
</script>

<style>
  @media (min-width: 480px) {
    div {
      display: grid;
      grid-template-columns: 40% auto;
      grid-gap: 1em;
    }
  }
</style>

{#if !$session?.user}
  <h2>Login</h2>
  <div>
    <ThirdPartyLogin></ThirdPartyLogin>
    <NativeLogin></NativeLogin>
  </div>
{:else}
  <AlreadyLoggedIn></AlreadyLoggedIn>
{/if}