<script lang="ts">
  import AdventureBanner from '$components/AdventureBanner.svelte'
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import Error from '$components/Error.svelte'

  let adventures = supabase.from('adventures').select('id,name,price,public').throwOnError().then(({data}) => data)
  title.set('Puzzad: Adventures')
</script>

<h2>Which adventure would you like to go on?</h2>
{#await adventures}
  <Spinner/>
{:then adventures}
  {#each adventures as adventure}
    <AdventureBanner
        adventureName="{adventure.name}"
        price="{adventure.price}"
        isPublic="{adventure.public}"
        code=""
        status=""
    />
  {:else}
    <p>No Adventures, sorry.</p>
  {/each}
{:catch error}
  <Error error={error}></Error>
{/await}