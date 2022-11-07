<script lang="ts">
    import RandomText from "$comps/RandomText.svelte"
    import {getGameClient} from "$lib/db";

    export let gameCode = ''
    export let puzzleId = 0

    export const refresh = async () => {
        const gameClient = await getGameClient(gameCode)
        const {data, error} = await gameClient.rpc('gethints', {puzzleid: puzzleId, gamecode: gameCode})
        if (error) {
            throw error
        }
        hints = data
    }

    const request = async function (id) {
        const gameClient = await getGameClient(gameCode)
        await gameClient.rpc('requesthint', {puzzleid: puzzleId, gamecode: gameCode, hintid: id})
    }

    let hints = []

    $: if (gameCode && puzzleId) {
        refresh()
    }

    const hintMessages = [
        'Need a hint? Browse our extensive collection below!',
        'Not sure where to go? Try one of our finest hints!',
        'Get your hints here! Freshly plucked from the hint tree!',
        'Psst... Can I interest you in a hint?',
        'We\'ve got the finest hints in all the land. Don\'t believe me? Try one for free!',
        'Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?',
    ]
</script>

<style>
    section {
        position: relative;
        border-top: 10px solid #333333;
        border-bottom: 2px solid #333333;
        background-color: rgba(51, 51, 51, .3);
        padding: 0 0.7em;
    }

    section::before {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "💡 hints";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #333333;
        color: white;
    }

    section h4 {
        border-bottom: 2px solid #e3bc5e;
        font-variant: small-caps;
        text-transform: lowercase;
        padding: 0;
        margin: 0;
        font-weight: bold;
    }

    div {
        position: relative;
    }

    div p.locked {
        filter: blur(5px);
    }

    div p {
        transition: filter 1s;
    }

    button {
        display: block;
        position: absolute;
        top: 25%;
        bottom: 25%;
        left: 25%;
        right: 25%;
        width: 50%;
        height: 50%;
        text-align: center;
    }
</style>

<section>
    <p>
        <RandomText options={hintMessages}></RandomText>
    </p>
    {#each hints as hint (hint.id)}
        <h4>{hint.title}</h4>
        <div>
            <p class="{hint.locked ? 'locked' : ''}">
                {#if hint.locked}
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                    labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                    laboris nisi ut aliquip ex ea commodo consequat.
                {:else}
                    {hint.message}
                {/if}
            </p>
            {#if hint.locked}
                <button on:click={() => request(hint.id)}>Reveal this hint</button>
            {/if}
        </div>
    {/each}
</section>