<script lang='ts'>
    import {supabase} from "$lib/db";
    import {replace} from 'svelte-spa-router'

    interface InfoText {
        error: boolean;
        text: string | null;
    }

    let email: string = "";
    let password: string = "";
    let infoText: InfoText = {error: null, text: null};
    let disabled = false

    const handleLogin = async () => {
        disabled = true
        const {user, error,} = await supabase.auth.signInWithPassword({email, password})
        disabled = false
        if (error) {
            infoText = {error: true, text: error.message};
        } else if (!error) {
            await replace('#/')
        }
    };
</script>

<style>
    [role='alert'] {
        background: indianred;
    }
    @media (max-width: 480px) {
        form {
            display: grid;
            grid-template-columns: auto;
            grid-auto-flow: row;
        }
        form > label  {
            grid-row: auto;
        }
        form > input,
        form > button {
            grid-row: auto;
            padding: 1em;
            margin: .5em 0 .5em 0;
        }
    }

    @media (min-width: 480px) {
        form {
            display: grid;
            grid-template-columns: [labels] 1fr [controls] auto;
            grid-auto-flow: row;
        }
        form > label  {
            grid-column: labels;
            grid-row: auto;
            padding-right: 1em;
        }
        form > input,
        form > button {
            grid-column: controls;
            grid-row: auto;
            border: none;
            padding: 1em;
        }
        section {
            grid-column: controls;
            grid-row: auto;
            margin: 0;
            padding: 1em;
            border-radius: 6px;
            text-align: center;
            line-height: normal;
        }
    }
</style>

<form on:submit|preventDefault>
    <label for='email'>Email:</label>
    <input
            id='email'
            type='email'
            name='email'
            bind:value={email}
            required />
    <label for='password'>Password:</label>
    <input
            id='password'
            type='password'
            name='password'
            bind:value={password}
            required
    />
    <button
            type='submit'
            on:click={() => handleLogin()}
            disabled={disabled}>
        Sign In
    </button>
</form>
{#if !!infoText.text}
    <div role={infoText.error ? "alert" : null}>
        {infoText.text}
    </div>
{/if}
<div>
    Don't have an account? <a href='#/signup'>Sign up!</a>
</div>