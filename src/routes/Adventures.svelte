<script>
    import AdventureBanner from '$comps/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import Spinner from '$comps/Spinner.svelte'
    import {onMount} from 'svelte'

    let adventures = []
    let initial = true
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('adventures')
                .select(`
                    id,name,price,public
                `)
        if (!error) {
            initial = false
            adventures = obtained
        }
    })
</script>

{#if initial}
    <Spinner />
{:else}
    <h2>Which adventure would you like to go on?</h2>
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
{/if}