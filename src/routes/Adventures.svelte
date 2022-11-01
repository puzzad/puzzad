<script>
    import AdventureBanner from '$lib/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onMount} from 'svelte'

    let adventures = []
    let initial = true
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('adventures')
                .select(`
                    id,name,promoBackground,promoLogo,price
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
    {#each adventures as adventure}
        <AdventureBanner
                adventureName='{adventure.name}'
                backgroundUrl='{adventure.promoBackground}'
                logoUrl='{adventure.promoLogo}'
                price='{adventure.price}'
                code=''
                status=''
        />
    {:else}
        <p>No Adventures, sorry.</p>
    {/each}
{/if}