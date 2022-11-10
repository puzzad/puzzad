<script lang="ts">
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import Error from '$components/Error.svelte'

  export let page

  let data = supabase.from('content').
      select('content').
      eq('page', page).
      throwOnError().
      single().
      then(({data}) => data.content)
</script>

{#await data}
  <Spinner></Spinner>
{:then content}
  {@html content}
{:catch error}
  <Error error={error}></Error>
{/await}