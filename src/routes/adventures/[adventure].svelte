<script>
  import {supabase} from '$lib/db.ts'
  import Spinner from '$comps/Spinner.svelte'
  import AdventureLogo from '$comps/AdventureLogo.svelte'
  import {title} from '$lib/title.ts'
  import Error from '$comps/Error.svelte'
  import {goto} from '@roxi/routify'

  export let adventure = ''

  let details = supabase.from('adventures').
      select(`id,name,description,price`).
      eq('name', adventure).
      then(({data, error}) => {
        if (error) {
          throw error
        }
        title.set('Puzzad: ' + data[0].name)
        return data[0]
      })

  let loggedIn = supabase.auth.getSession().then(({data: {session}}) => !!session?.user)

  const addAdventure = async (details) => {
    if (details.price === 0) {
      let {data, error} = await supabase.rpc('addfreeadventure', {adventureid: details.id})
      if (data) {
        $goto('/games/[code]', {code: data})
      }
    }
  }
</script>
{#await details}
  <Spinner/>
{:then details}
  <h2>
    <AdventureLogo name={details.name}></AdventureLogo>
  </h2>
  <section>
    {@html details.description}
  </section>
  {#await loggedIn then isLoggedIn}
    {#if isLoggedIn}
      <button on:click={() => addAdventure(details)}>Start this adventure for {details.price === 0 ? 'free' : '£' +
          details.price}</button>
    {:else}
      <p>You must <a href="/login">login</a> before you can start an adventure.</p>
    {/if}
  {/await}
{:catch error}
  <Error error={error}></Error>
{/await}