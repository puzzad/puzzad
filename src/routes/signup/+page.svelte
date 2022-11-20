<script lang="ts">
  import {title} from '$lib/title'
  import {session} from '$lib/auth'
  import SignUp from '$components/SignUp.svelte'
  import AlreadyLoggedIn from '$components/AlreadyLoggedIn.svelte'

  title.set('Puzzad: Signup')
</script>

{#if !$session?.user}
  <SignUp/>
{:else}
  <AlreadyLoggedIn></AlreadyLoggedIn>
{/if}