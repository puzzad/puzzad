<script>
    import {push} from 'svelte-spa-router'

    let code = '';

    const handleGameCode = () => {
        push('/game/' + code.toLowerCase().replaceAll(/[^a-z-]/g, ''))
    }
</script>
<style>
    main {
        display: grid;
        grid-template-columns: 60% 40%;
        gap: 2em;
    }

    section {
        border-radius: 15px;
        padding: 1em;
        margin: 0;
    }

    h3 {
        margin: 0;
    }

    #hero {
        background-image: url('../assets/hero.jpg');
        background-repeat: no-repeat;
        background-size: cover;
        background-position: center;
        position: relative;
    }

    #hero h3 {
        position: absolute;
        left: 30%;
        right: 0.5em;
        top: 0.5em;
        font-size: xx-large;
        color: #000000;
        font-weight: bold;
        text-align: right;
        text-shadow: 0 0 5px #ffffff;
    }

    #hero a {
        position: absolute;
        bottom: 1em;
        right: 1em;
        font-size: large;
        font-weight: bold;
        text-shadow: 0 0 5px #000000;
    }

    #play {
        background-color: #39254D;
        display: flex;
        flex-direction: column;
        align-items: stretch;
        gap: 0.5em;
    }

    #play h3, #play input {
        text-align: center;
    }

    form {
        display: contents;
    }

    #intro {
        grid-column-start: 1;
        grid-column-end: 3;
        background-color: #222222;
    }

    #intro h3 {
        border-bottom: 2px solid #e3bc5e;
        padding: 0;
        margin: 0;
        font-weight: bold;
        font-size: large;
    }
</style>

<h2>Puzzad: Puzzle Adventures</h2>
<main>
    <section id="hero">
        <h3>Ready to start an adventure?</h3>
        <a href="/#/Adventures">Let's go &raquo;</a>
    </section>
    <section id="play">
        <h3>Got a game code?</h3>
        <form on:submit|preventDefault={handleGameCode}>
            <input type="text" placeholder="revolving-magenta-ocelot" bind:value={code}>
            <input type="submit" value="Play!">
        </form>
    </section>
    <section id="intro">
        <h3>What is a puzzle adventure?</h3>
        <p>
            A puzzle adventure is part escape room, part puzzle hunt.
            Each adventure is a story interleaved with a variety of puzzles.
            To progress from one part of the story to the next you have to
            solve the puzzle you're presented with &mdash; but don't worry,
            there are ample hints available if you get stuck!
        </p>
        <h3>Is there a time limit?</h3>
        <p>
            Nope! Puzzle adventures are completely self-driven. If you want
            to stop and carry on later that's not a problem. Want to treat
            it like an advent calendar and open one door a day? Go ahead!
        </p>
        <p>
            For those who are more competitive we'll show your total solve
            time when you finish the adventure (along with number of hints
            used) so you can compare against your friends.
        </p>
        <h3>Can I play with friends?</h3>
        <p>
            Yes, and we actively encourage it! When you start an adventure
            you'll get a code that looks something like <code>revolving-magenta-ocelot</code>.
            Simply give that code to your friends, they put it in the box
            above and you can all play together.
        </p>
        <p>
            If you're not physically present, we recommend jumping on a voice
            chat with your team (take your pick from Discord, Zoom, Teams,
            and so on). You can see your teammate's guesses in real time,
            but it helps if you can ask them <em>why</em> they keep guessing
            "elephant".
        </p>
        <h3>How difficult is it?</h3>
        <p>
            Our starter adventure, Conspiracy, is designed to be around the
            difficulty of a medium escape room. Having some knowledge of
            typical escape rooms or puzzle hunts challenges will help, but
            there is an extensive hint system that can point you in the right
            direction if you're struggling. You can even skip to the end and
            just get the answer, if you're really fed up with one of the
            puzzles. It's your adventure!
        </p>
        <p>
            If you're a dedicated puzzler who takes parts in puzzle
            hunts or thrives on the more difficult escape rooms, we recommend
            avoiding all hints and seeing how quickly you can complete the
            full adventure. Alternatively, you could use it as a gentle introduction
            to rope some less puzzle-inclined friends into the hobby!
        </p>
    </section>
</main>