<script>
    export let params = {}
    import {supabase} from '$lib/db'
    import {onMount} from 'svelte'

    let data = {}
    let initial = true
    let displayError
    onMount(async function () {
        let {data: game, error} =
                    await supabase.from('games')
                                  .select('puzzle, adventures ( name, description)')
                                  .eq('id', params.id)
        initial = false
        if (game[0]?.adventures) {
            data.puzzle = game[0]?.puzzle
            data.adventure = game[0]?.adventures
        } else {
            error = "Unable to find game"
        }
        if (error) {
            displayError = error
        }
    })
</script>
{#if initial}
{:else if displayError || !data}
    <h1>Error finding game</h1>
    <p>{displayError}</p>
{:else}
    <h1>{data.adventure.name}</h1>
    <p>{data.adventure.description}</p>
    <p>Current puzzle ID: {data.puzzle}</p>
{/if}
