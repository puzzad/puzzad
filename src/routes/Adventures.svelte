<script>
    import AdventureBanner from '$lib/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import {onMount} from 'svelte'

    let adventures = []
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('adventures')
                .select(`
                    id,name,description,price
                `)
        if (!error) {
            adventures = obtained
        }
    })
</script>

{#each adventures as adventure}
    <AdventureBanner
            name='{adventure.name}'
            description='{adventure.description}'
            price='{adventure.price}'
            code=''
            status=''
    />
{:else}
    <p>No Adventures, sorry.</p>
{/each}