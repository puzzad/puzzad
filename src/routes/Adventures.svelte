<script>
    import AdventureBanner from '$lib/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import {onMount} from 'svelte'

    let adventures = []
    onMount(async function () {
        let {data: obtained, error} = await supabase
            .from('adventures')
            .select(`
                    id,name,promoBackground,promoLogo,price
                `)
        if (!error) {
            adventures = obtained
        }
    })
</script>

{#each adventures as adventure}
    <AdventureBanner
            adventureName='{adventure.name}'
            backgroundUrl="{adventure.promoBackground}"
            logoUrl="{adventure.promoLogo}"
            price='{adventure.price}'
            code=''
            status=''
    />
{:else}
    <p>No Adventures, sorry.</p>
{/each}