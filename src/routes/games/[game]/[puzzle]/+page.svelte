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

    const load = async (game, puzzle) =>
        getGameClient(game).then((client) => client.from('puzzles').select('title, content, next, storage_slug, adventure (name)').eq('id', puzzle).throwOnError().single(),
        ).then(({data}) => data).then((data) => {
            title.set(`Puzzad: ${data.adventure.name}: ${data.title}`)
            return data
        }).catch(() => goto(`/games/${game}`))

    $: if (data.game || data.puzzle) {
        solved = false
    }

</script>

{#await load(data.game, data.puzzle)}
    <Spinner/>
{:then gameData}
    <h2>{gameData.adventure.name}: {gameData.title}</h2>

    <PuzzleContent gameCode={data.game} storageSlug={gameData.storage_slug} content={gameData.content}></PuzzleContent>
    <PuzzleAnswer gameCode={data.game} puzzle={data.puzzle}></PuzzleAnswer>
    <Hints gameCode={data.game} puzzleId={data.puzzle} bind:this={hints}></Hints>

    {#if solved}
        <VictoryDialog data={data} next={gameData.next}></VictoryDialog>
    {/if}

    <GuessMonitor data={data} on:hint={() => {hints && hints.refresh()}}
                  on:solve={() => {solved = true}}></GuessMonitor>
{/await}