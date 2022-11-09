<script>
    import {goto} from '$app/navigation'
    import {getGameClient} from '$lib/db'
    import Spinner from '$lib/components/Spinner.svelte'
    import {title} from '$lib/title.ts'
    import Hints from '$lib/components/Hints.svelte'
    import PuzzleContent from '$lib/components/PuzzleContent.svelte'
    import PuzzleAnswer from '$lib/components/PuzzleAnswer.svelte'
    import VictoryDialog from '$lib/components/VictoryDialog.svelte'
    import GuessMonitor from '$lib/components/GuessMonitor.svelte'

    let hints
    let solved = false

    export let data

    console.log(data)

    const load = async (game, puzzle) =>
            getGameClient(game).
                               then((client) => client.from('puzzles').
                                    select('title, content, next, storage_slug, adventure (name)').
                                                      eq('id', puzzle).
                                                      throwOnError().
                                                      single(),
            ).
                               then(({data}) => data).
                               then((data) => {
                title.set(`Puzzad: ${data.adventure.name}: ${data.title}`)
                return data
            }).
                               catch(() => goto(`/games/${game}`))

    $: if (data.game || data.puzzle) { solved = false }

</script>

{#await load(data.game, data.puzzle)}
    <Spinner/>
{:then data}
    <h2>{data.adventure.name}: {data.title}</h2>

    <PuzzleContent gameCode={data.code} storageSlug={data.storage_slug} content={data.content}></PuzzleContent>
    <PuzzleAnswer gameCode={data.code} puzzle={data.puzzle}></PuzzleAnswer>
    <Hints gameCode={data.code} puzzleId={data.puzzle} bind:this={hints}></Hints>

    {#if solved}
        <VictoryDialog next={data.next}></VictoryDialog>
    {/if}

    <GuessMonitor on:hint={() => {hints && hints.refresh()}} on:solve={() => {solved = true}}></GuessMonitor>
{/await}