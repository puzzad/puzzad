<script>
    export let params = {}
    import {supabase} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onMount, onDestroy} from 'svelte'
    import {toasts, ToastContainer, FlatToast} from "svelte-toasts";

    let data = {}
    let guess = ''
    let solved = false
    let checkingGuess = false
    let initial = true
    let displayError = null
    let guessesChannel = null
    $: disableInput = checkingGuess || solved

    onMount(async function () {
        let {data: puzzle, error} =
            await supabase.from('puzzles')
                .select('title, content, next, adventure (name)')
                .eq('id', params.puzzle)
        initial = false
        if (puzzle.length > 0) {
            data = puzzle[0]

            guessesChannel = supabase
                .channel('public:guesses:game=eq.' + params.code)
                .on('postgres_changes', {
                    event: 'INSERT',
                    schema: 'public',
                    table: 'guesses',
                    filter: 'game=eq.' + params.code
                }, handleStreamedGuess)
                .subscribe()
        } else {
            error = "Unable to find puzzle"
        }
        if (error) {
            displayError = error
        }
    })

    onDestroy(async function () {
        if (guessesChannel != null) {
            await supabase.removeChannel(guessesChannel)
        }
    })

    const handleStreamedGuess = function (payload) {
        if (payload.new.correct) {
            solved = true;

            toasts.add({
                title: 'Correct guess!',
                description: payload.new.content,
                duration: 0,
                type: 'success',
            });
        } else {
            toasts.add({
                title: 'Incorrect guess',
                description: payload.new.content,
                duration: 10000,
                type: 'error',
            });
        }
    }

    const handleGuess = async function () {
        checkingGuess = true
        await supabase.from('guesses')
            .insert({content: guess, puzzle: params.puzzle, game: params.code})
        checkingGuess = false
        guess = ''
    }
</script>

<style>
    section.tip {
        position: relative;
        border-top: 10px solid #e3bc5e;
        border-bottom: 2px solid #e3bc5e;
        background-color: rgba(227, 188, 94, .3);
        padding: 0 0.7em;
    }

    section.tip::before {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "💡 tip";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #e3bc5e;
        color: black;
    }

    section.story {
        position: relative;
        border-top: 10px solid #1c76c5;
        border-bottom: 2px solid #1c76c5;
        background-color: rgba(28, 118, 197, .3);
        padding: 0 0.7em;
    }

    section.story::before {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "📖 story";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #1c76c5;
        color: white;
    }

    section.answer fieldset {
        display: flex;
    }

    section.answer fieldset input[type=text] {
        flex-grow: 1;
    }

    section.hints h4 {
        border-bottom: 2px solid #e3bc5e;
        font-variant: small-caps;
        text-transform: lowercase;
        padding: 0 5px;
        font-weight: bold;
    }
</style>

{#if initial}
    <Spinner/>
{:else if displayError || !data}
    <h1>Error finding puzzle</h1>
    <p>{displayError}</p>
{:else}
    <h1>{data.adventure.name}: {data.title}</h1>
    <p>{data.content}</p>
    <section class="answer">
        <form on:submit|preventDefault={() => handleGuess()}>
            <fieldset>
                <legend>Enter a guess</legend>
                <input type="text" bind:value={guess} disabled={disableInput}>
                <input type="submit" value="Submit" disabled={disableInput}>
            </fieldset>
        </form>
    </section>

    <ToastContainer placement="bottom-right" theme="dark" let:data={data}>
        <FlatToast {data}/>
    </ToastContainer>
{/if}
