<script>
  import {goto} from '$app/navigation'
  import {isAdmin} from '$lib/auth'
  import {browser} from '$app/environment'

  if (browser && !$isAdmin) {
    goto('/')
  }
</script>
{#if $isAdmin}
  <h1>Admin Page</h1>
  <slot />
{/if}