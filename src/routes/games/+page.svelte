<script>
    import AdventureBanner from '$lib/components/AdventureBanner.svelte'
    import {supabase} from '$lib/db'
    import Spinner from '$lib/components/Spinner.svelte'
    import RandomText from '$lib/components/RandomText.svelte'
    import {title} from '$lib/title.ts'
    import Error from '$lib/components/Error.svelte'

    title.set('Puzzad: My Games')
    let games = supabase.from('games').select('id, status, adventures (name, public), status, code').throwOnError().then(({data}) => data)

    const quotes = [
        '“Never say \'no\' to adventures. Always say \'yes,\' otherwise you\'ll lead a very dull life.”\n― Ian Fleming',
        '“If happiness is the goal – and it should be, then adventures should be a priority.”\n― Richard Branson',
        '“Let us step into the night and pursue that flighty temptress, adventure“\n― J K Rowling',
        '“Life is either a daring adventure or nothing at all.”\n― Helen Keller',
        '“Adventure is worthwhile in itself.”\n― Amelia Earhart',
        '“Never fear quarrels, but seek hazardous adventures.”\n― Alexandre Dumas',
        '“When you see someone putting on his Big Boots, you can be pretty sure that an Adventure is going to happen.”\n― A.A. Milne',
        '“We are plain quiet folk and have no use for adventures. Nasty disturbing uncomfortable things! Make you late for dinner!”\n― J.R.R. Tolkien',
        '“Make your choice, adventurous Stranger,\nStrike the bell and bide the danger,\nOr wonder, till it drives you mad,\nWhat would have followed if you had.”\n― C.S. Lewis',
    ]
</script>

<style>
    blockquote {
        white-space: pre-line;
        font-size: xx-large;
        font-style: italic;
        border: 0;
        background-color: #222222;
        border-radius: 15px;
        padding: 1em;
        margin: 0;
    }
</style>

<section>
    <h2>Your adventures</h2>
    {#await games}
        <Spinner/>
    {:then games}
        {#each games as game}
            <AdventureBanner
                    status="{game.status}"
                    adventureName='{game.adventures?.name ?? "Unknown"}'
                    code="{game.code}"
                    isPublic="{game.adventures?.public}"
            />
        {:else}
            <blockquote>
                <RandomText options={quotes}></RandomText>
            </blockquote>
            <p>
                You aren't part of any adventures! You can
                <a href="/adventures">browse the available adventures</a>, or if you
                want to join a friend you can <a href="/#/">enter the game code</a> to
                jump straight in.
            </p>
        {/each}
    {:catch error}
        <Error error={error}></Error>
    {/await}
</section>