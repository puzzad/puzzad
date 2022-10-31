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
                                  .eq('code', params.code)
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
<style>
    section {
        display: flex;
        justify-content: center;
    }
    @keyframes spinner {
        0% {
            transform: translate3d(-50%, -50%, 0) rotate(0deg);
        }
        100% {
            transform: translate3d(-50%, -50%, 0) rotate(360deg);
        }
    }
    .spin::before {
        animation: 1.5s linear infinite spinner;
        animation-play-state: inherit;
        border: solid 5px var(--border);
        border-bottom-color: var(--links);
        border-radius: 50%;
        content: "";
        height: 40px;
        width: 40px;
        position: absolute;
        transform: translate3d(-50%, -50%, 0);
        will-change: transform;
    }
</style>
{#if initial}
    <section>
        <div class='spin'></div>
    </section>
{:else if displayError || !data}
    <h1>Error finding game</h1>
    <p>{displayError}</p>
{:else}
    <h1>{data.adventure.name}</h1>
    <p>{data.adventure.description}</p>
    <p>Current puzzle ID: {data.puzzle}</p>
{/if}
