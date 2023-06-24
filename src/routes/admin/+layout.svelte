<script>
  import {goto} from '$app/navigation'
  import {isAdmin} from '$lib/auth'
  import Spinner from '$components/Spinner.svelte'
</script>

{#if $isAdmin === null}
  <Spinner></Spinner>
{:else if $isAdmin === true}
  <slot />
{:else}
  {goto('/')}
{/if}