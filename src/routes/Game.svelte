<script>
    export let params = {}
    import {getGameClient} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onMount} from 'svelte'
    import {push} from 'svelte-spa-router'

    let data = {}
    let initial = true
    let displayError
    let gc

    onMount(async function () {
        gc = await getGameClient(params.code)

        let {data: game, error} =
            await gc.from('games')
                .select('status, puzzle, startTime, endTime, adventures ( name, promoLogo, description)')
                .eq('code', params.code)

        if (game[0]?.adventures) {
            data = game[0]
        } else {
            error = "Unable to find game"
        }
        if (error) {
            displayError = error
        }
        initial = false
    })

    const handleStartAdventure = async function() {
        let { data: puzzle } = await gc.rpc('startadventure')
        if (puzzle) {
            await push('/game/' + params.code + '/' + puzzle)
        }
    }

    const handleContinueAdventure = async function() {
        await push('/game/' + params.code + '/' + data.puzzle)
    }
</script>

<style>
    code {
        display: block;
        background-image: url('../assets/code-bg.png');
        background-color: transparent;
        color: black;
        text-align: center;
        background-size: contain;
        background-repeat: no-repeat;
        background-position: center;
        padding: 70px 0;
        font-size: xxx-large;
        font-family: monospace;
        margin: 1em;
    }

    button {
        width: 100%;
        margin: 1em 0;
        font-size: x-large;
        padding: 0.5em;
    }
</style>

{#if initial}
    <Spinner/>
{:else if displayError || !data}
    <h1>Error finding game</h1>
    <p>{displayError}</p>
{:else}
    <h1><img src="{data.adventures.promoLogo}" alt="{data.adventures.name}"></h1>
    {#if data.status === 'EXPIRED'}
        <p>Congratulations! You finished the adventure!</p>
        <p>You took {Math.round((new Date(data.endTime).getTime() - new Date(data.startTime).getTime())/10/60)/100} minutes!</p>
    {:else if data.status === 'PAID'}
        <p>
            You've not yet started your adventure! Remember, it's dangerous to go alone.
            Take this if you want to recruit others to help you:
        </p>
        <code>
            {params.code}
        </code>
        <p>
            Once you're ready, press below to go to the first puzzle in the adventure!
        </p>
        <button on:click={() => handleStartAdventure()}>
            Begin the adventure!
        </button>
    {:else if data.status === 'ACTIVE'}
        <p>This adventure is already in progress!</p>
        <button on:click={() => handleContinueAdventure()}>
            Go to the current puzzle
        </button>
    {/if}
{/if}
