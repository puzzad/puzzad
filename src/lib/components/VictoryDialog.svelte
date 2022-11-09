<script>
    import {goto} from '$app/navigation'
    import RandomText from '$lib/components/RandomText.svelte'
    import {Fireworks} from '@fireworks-js/svelte'

    export let next
    export let data

    const goToNextPuzzle = async function () {
        goto(`/games/${data.game}/${next}`)
    }

    const goToGamePage = async function () {
        goto(`/games/${data.game}`)
    }

    const congratsMessages = [
        'Congratulations!',
        'Well done!',
        'You did it!',
        'Huzzah!',
        'Good job!',
        'Nice work!',
    ]
</script>

<style>
    dialog {
        position: fixed;
        padding: 3em;
        top: 30vh;
        left: 30vh;
        right: 30vh;
        background: linear-gradient(to top, #106310, #3C8E2B);
        border-radius: 10px;
        text-align: center;
        z-index: 1001;
        animation: show 200ms linear;
    }

    @keyframes show {
        from {
            transform: scale(0.001);
        }
        to {
            transform: scale(1);
        }
    }

    dialog h3 {
        margin: 0 0 1em 0;
    }

    dialog p {
        margin: 0 0 2em 0;
    }

    dialog button {
        width: 100%;
        text-transform: uppercase;
        font-weight: bold;
        font-size: x-large;
        font-variant: all-small-caps;
    }

    /*noinspection CssUnusedSymbol*/
    :global(.fireworks) {
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        position: fixed;
        z-index: 1000;
        background: rgba(0, 0, 0, 0.7);
    }
</style>

<dialog open="open">
    <h3>
        <RandomText options={congratsMessages}></RandomText>
    </h3>
    {#if next}
        <p>You have completed this step in the adventure!</p>
        <button on:click={() => goToNextPuzzle()}>Next puzzle &raquo;</button>
    {:else}
        <p>You have completed the adventure!</p>
        <button on:click={() => goToGamePage()}>Continue &raquo;</button>
    {/if}
</dialog>

<Fireworks autostart={true} class="fireworks"/>