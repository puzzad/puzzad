<script>
    import AdventureBanner from '$lib/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import {onMount} from 'svelte'

    let games = []
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('games')
                .select(`
                    id,status, adventures ( name, description ), status, code
                `)
        if (!error) {
            games = obtained
        }
    })
</script>
<section>
    This is a list of your games.
    {#each games as game}
        <AdventureBanner
                price=''
                status='{game.status}'
                name='{game.adventures?.name ?? "Unknown"}'
                description='{game.adventures?.description ?? "Unknown"}'
                code='{game.code}'
        />
    {:else}
        <p>You haven't  purchased any games.</p>
    {/each}
</section>