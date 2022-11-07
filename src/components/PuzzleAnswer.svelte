<script lang="ts">
    import {getGameClient} from "$lib/db";

    export let puzzle = 0
    export let gameCode = ''

    let guess = ''
    let checking = false

    const submit = async () => {
        checking = true
        await getGameClient(gameCode)
            .then((gc) => gc
                .from('guesses')
                .insert({content: guess, puzzle: puzzle, game: gameCode})
                .throwOnError()
            )
            .then(() => guess = '')
            .finally(() => checking = false)
    }
</script>

<style>
    @media (max-width: 480px) {
        fieldset {
            flex-direction: column;
        }
    }

    fieldset {
        display: flex;
    }

    fieldset input[type=text] {
        flex-grow: 1;
    }
</style>

<section>
    <form on:submit|preventDefault={submit}>
        <fieldset>
            <legend>Enter a guess</legend>
            <input type="text" bind:value={guess} disabled={checking}>
            <input type="submit" value="Submit" disabled={checking}>
        </fieldset>
    </form>
</section>