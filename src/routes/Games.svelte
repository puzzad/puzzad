<script>
    import AdventureBanner from '$lib/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onMount} from 'svelte'

    let games = []
    let initial = true
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('games')
                .select(`
                    id,status, adventures ( name, promoBackground, promoLogo ), status, code
                `)
        if (!error) {
            initial = false
            games = obtained
        }
    })
</script>

<section>
    This is a list of your games.
    {#if initial}
        <Spinner />
    {:else}
        {#each games as game}
            <a href='/#/game/{game.id}'>
                <AdventureBanner
                        price=''
                        status='{game.status}'
                        adventureName='{game.adventures?.name ?? "Unknown"}'
                        backgroundUrl='{game.adventures?.promoBackground}'
                        logoUrl='{game.adventures?.promoLogo}'
                        code='{game.code}'
                />
            </a>
        {:else}
            <p>You haven't purchased any games.</p>
        {/each}
    {/if}
</section>