<script>
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
<style>
    article {
        margin: 2rem 0;
        padding: 0;
        border-radius: 1rem 1rem 0 0;
    }
    article div {
        padding: 0.5rem;
        margin: 0;
        border-radius: 1rem 1rem 0 0;
    }
    article div.UNPAID {
        background: indianred;
    }
    article div.PAID {
        background: darkseagreen;
    }
    article div.EXPIRED {
        background: darkslategray;
    }
    article h2 {
        margin: 1rem;
        padding: 0;
    }
    article p {
        margin: 1rem;
        padding: 0;
    }
</style>
<section>
    This is a list of your games.
    {#each games as game}
        <article>
            <div class='{game.status}'>{game.status}</div>
            <h2>{game.adventures?.name}</h2>
            <p>{game.adventures?.description}</p>
            <hr>
            <p>Code: {game.code}</p>
        </article>
    {:else}
        <p>You haven't  purchased any games.</p>
    {/each}
</section>