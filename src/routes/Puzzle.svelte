<script>
    export let params = {}
    import {getGameClient} from '$lib/db'
    import Spinner from '$lib/Spinner.svelte'
    import {onDestroy} from 'svelte'
    import {toasts, ToastContainer, FlatToast} from "svelte-toasts";
    import {Confetti} from "svelte-confetti";
    import {replace} from 'svelte-spa-router'

    let data = {}
    let guess = ''
    let solved = false
    let checkingGuess = false
    let initial = true
    let guessesChannel = null
    let gameClient = null
    let hints = []

    $: if (params.code && params.puzzle) {
        load()
    }

    const load = () => {
        try {
            reset()
            getGameClient(params.code)
                .then((client) => {
                    gameClient = client
                    return client.from('puzzles')
                        .select('title, content, next, storage_slug, adventure (name)')
                        .eq('id', params.puzzle)
                })
                .then(checkQueryResults)
                .then(obtainUrlReplacements)
                .then(performUrlReplacements)
                .then((puzzle) => data = puzzle)
                .then(() => {
                    if (!guessesChannel) {
                        startMonitoringGuesses()
                    }
                    initial = false
                })
                .then(refreshHints)
                .catch(() => replace('/game/' + params.code))
        } catch (e) {
            replace('/game/' + params.code)
        }
    }

    const reset = () => {
        console.log("Resetting...")
        initial = true
        solved = false
        guess = ''
        checkingGuess = false
        hints = []
    }

    const checkQueryResults = ({data, error}) => {
        if (error) {
            throw error
        } else if (data.length === 0) {
            throw 'Puzzle not found'
        } else {
            return data[0]
        }
    }

    const obtainUrlReplacements = (puzzle) => {
        let res = [puzzle]
        const urls = puzzle.content.match(/\$[^$]+?\$/g)
        if (urls) {
            urls.forEach((slug) => {
                res.push(
                    slug,
                    gameClient
                        .storage
                        .from('puzzles')
                        .createSignedUrl(puzzle.storage_slug + '/' + slug.substring(1, slug.length - 1), 60 * 60),
                )
            })
        }
        return Promise.all(res)
    }

    const performUrlReplacements = (results) => {
        let [puzzle, ...urls] = results
        for (let i = 0; i < urls.length; i += 2) {
            const {data: {signedUrl}, error} = urls[i + 1]
            if (error) {
                throw error
            }

            puzzle.content = puzzle.content.replace(urls[i], signedUrl)
        }
        return puzzle
    }

    const refreshHints = async () => {
        const {data, error} = await gameClient.rpc('gethints', {puzzleid: params.puzzle, gamecode: params.code})
        if (error) {
            throw error
        }
        hints = data
    }

    onDestroy(async function () {
        if (guessesChannel) {
            await gameClient.removeChannel(guessesChannel)
        }
    })

    const startMonitoringGuesses = async function () {
        guessesChannel = await gameClient
            .channel('public:guesses:game=eq.' + params.code)
            .on('postgres_changes', {
                event: 'INSERT',
                schema: 'public',
                table: 'guesses',
                filter: 'game=eq.' + params.code
            }, handleStreamedGuess)
            .subscribe()
    }

    const handleStreamedGuess = function (payload) {
        if (payload.new.puzzle.toString() === params.puzzle) {
            if (payload.new.content === '*hint') {
                refreshHints()
            } else if (payload.new.correct) {
                solved = true
            } else {
                toasts.add({
                    title: 'Incorrect guess',
                    description: payload.new.content,
                    duration: 10000,
                    type: 'error',
                });
            }
        }
    }

    const handleGuess = async function () {
        checkingGuess = true
        await gameClient.from('guesses')
            .insert({content: guess, puzzle: params.puzzle, game: params.code})
        checkingGuess = false
        guess = ''
    }

    const requestHint = async function (id) {
        await gameClient.rpc('requesthint', {puzzleid: params.puzzle, gamecode: params.code, hintid: id})
    }

    const goToNextPuzzle = async function () {
        await replace('/game/' + params.code + '/' + data.next)
    }

    const goToGamePage = async function () {
        await replace('/game/' + params.code)
    }

    const hintMessages = [
        'Need a hint? Browse our extensive collection below!',
        'Not sure where to go? Try one of our finest hints!',
        'Get your hints here! Freshly plucked from the hint tree!',
        'Psst... Can I interest you in a hint?',
        'We\'ve got the finest hints in all the land. Don\'t believe me? Try one for free!',
        'Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?',
    ]
    const hintMessage = hintMessages[Math.floor(Math.random() * hintMessages.length)];

    const congratsMessages = [
        'Congratulations!',
        'Well done!',
        'You did it!',
        'Huzzah!',
        'Good job!',
        'Nice work!',
    ]
    const congratsMessage = congratsMessages[Math.floor(Math.random() * congratsMessages.length)];
</script>

<style>
    :global(section.tip) {
        position: relative;
        border-top: 10px solid #e3bc5e;
        border-bottom: 2px solid #e3bc5e;
        background-color: rgba(227, 188, 94, .3);
        padding: 0 0.7em;
    }

    :global(section.tip::before) {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "ℹ️ information";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #e3bc5e;
        color: black;
    }

    :global(section.story) {
        position: relative;
        border-top: 10px solid #1c76c5;
        border-bottom: 2px solid #1c76c5;
        background-color: rgba(28, 118, 197, .3);
        padding: 0 0.7em;
    }

    :global(section.story::before) {
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

    section.hints {
        position: relative;
        border-top: 10px solid #333333;
        border-bottom: 2px solid #333333;
        background-color: rgba(51, 51, 51, .3);
        padding: 0 0.7em;
    }

    section.hints::before {
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

    section.hints h4 {
        border-bottom: 2px solid #e3bc5e;
        font-variant: small-caps;
        text-transform: lowercase;
        padding: 0;
        margin: 0;
        font-weight: bold;
    }

    .locked {
        position: relative;
    }

    .locked p {
        filter: blur(5px);
    }

    .locked button {
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
    }

    dialog[open] {
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

    .confetti-bg {
        position: fixed;
        top: -50px;
        left: 0;
        bottom: 0;
        right: 0;
        display: flex;
        justify-content: center;
        overflow: hidden;
        background-color: rgba(0, 0, 0, 0.8);
        z-index: 1000;
    }
</style>

{#if initial}
    <Spinner/>
{:else}
    <h2>{data.adventure.name}: {data.title}</h2>
    {@html data.content}
    <section class="answer">
        <form on:submit|preventDefault={() => handleGuess()}>
            <fieldset>
                <legend>Enter a guess</legend>
                <input type="text" bind:value={guess} disabled={checkingGuess}>
                <input type="submit" value="Submit" disabled={checkingGuess}>
            </fieldset>
        </form>
    </section>

    <section class="hints">
        <p>{hintMessage}</p>
        {#each hints as hint}
            <h4>{hint.title}</h4>
            {#if hint.locked}
                <div class="locked">
                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                        labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                        laboris nisi ut aliquip ex ea commodo consequat.</p>
                    <button on:click={() => requestHint(hint.id)}>Reveal this hint</button>
                </div>
            {:else}
                <p>{hint.message}</p>
            {/if}
        {/each}
    </section>

    <dialog open={solved}>
        <h3>{congratsMessage}</h3>
        {#if data.next}
            <p>You have completed this step in the adventure!</p>
            <button on:click={() => goToNextPuzzle()}>Next puzzle &raquo;</button>
        {:else}
            <p>You have completed the adventure!</p>
            <button on:click={() => goToGamePage()}>Continue &raquo;</button>
        {/if}
    </dialog>

    {#if solved}
        <div class="confetti-bg">
            <Confetti x={[-5, 5]} y={[0, 0.1]} delay={[0, 2000]} infinite duration=5000 amount=400
                      fallDistance="110vh"/>
        </div>
    {/if}

    <ToastContainer placement="bottom-right" theme="dark" let:data={data}>
        <FlatToast {data}/>
    </ToastContainer>
{/if}
