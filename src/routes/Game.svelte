<script>
    export let params = {}
    import {getGameClient} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onMount} from 'svelte'
    import {push} from 'svelte-spa-router'

    let data = {}
    let initial = true
    let displayError
    onMount(async function () {
        const gameclient = await getGameClient(params.code)

        let {data: game, error} =
                    await gameclient.from('games')
                                  .select('puzzle, adventures ( name, description)')
                                  .eq('code', params.code)

        if (game[0]?.adventures) {
            data.puzzle = game[0]?.puzzle
            data.adventure = game[0]?.adventures
            if (data.puzzle) {
                await push('/game/' + params.code + '/' + data.puzzle)
            } else {
                initial = false
            }
        } else {
            error = "Unable to find game"
            initial = false
        }
        if (error) {
            displayError = error
        }
    })
</script>
{#if initial}
    <Spinner />
{:else if displayError || !data}
    <h1>Error finding game</h1>
    <p>{displayError}</p>
{:else}
    <h1>{data.adventure.name}</h1>
    <p>{data.adventure.description}</p>
    <p>Current puzzle ID: <a href="/#/game/{params.code}/{data.puzzle}">{data.puzzle}</a></p>
{/if}
