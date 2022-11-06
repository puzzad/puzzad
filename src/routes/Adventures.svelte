<script>
    import AdventureBanner from '$comps/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import Spinner from '$comps/Spinner.svelte'

    let adventures = supabase
        .from('adventures')
        .select('id,name,price,public')
        .throwOnError()
        .then(({data}) => data)
</script>

<h2>Which adventure would you like to go on?</h2>
{#await adventures}
    <Spinner/>
{:then adventures}
    {#each adventures as adventure}
        <AdventureBanner
                adventureName='{adventure.name}'
                price='{adventure.price}'
                isPublic='{adventure.public}'
                code=''
                status=''
        />
    {:else}
        <p>No Adventures, sorry.</p>
    {/each}
{:catch error}
    <p>A problem occurred trying to retrieve adventures.</p>
{/await}