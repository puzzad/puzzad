<script lang="ts">
  import {title} from '$lib/title'
  import {isLoggedIn} from '$lib/auth'
  import SignUp from '$components/SignUp.svelte'
  import AlreadyLoggedIn from '$components/AlreadyLoggedIn.svelte'
  import Spinner from '$components/Spinner.svelte'

  title.set('Puzzad: Signup')
</script>

{#if $isLoggedIn === null}
  <Spinner></Spinner>
{:else if !$isLoggedIn}
  <SignUp/>
{:else}
  <AlreadyLoggedIn></AlreadyLoggedIn>
{/if}