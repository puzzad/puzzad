<script lang="ts">
  import {supabase} from '$lib/db'
  import {title} from '$lib/title'
  import Spinner from '$components/Spinner.svelte'
  import Error from '$components/Error.svelte'

  let extras = supabase.from('extras').
      select('filename,title,description').
      order('sort', {ascending: true}).
      throwOnError().
      then(({data}) => data)
  title.set('Puzzad: Extras')
</script>

<h2>Extras</h2>
{#await extras}
  <Spinner/>
{:then extras}
  <p>Interested in some bonus puzzles, or maybe a helpful guide? Check out our extras!</p>
  {#each extras as e}
    <h3><a href="{supabase.storage.from('extras').getPublicUrl(e.filename).data.publicUrl}">{e.title}</a></h3>
    <p>{e.description}</p>
  {/each}
{:catch error}
  <Error error={error}></Error>
{/await}