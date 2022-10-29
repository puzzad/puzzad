<script lang='ts'>
    import {supabase} from "$lib/db";

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
        const { _, error, } = await supabase.auth.signUp({email, password});
        disabled = false
        if (error) {
            infoText = {error: true, text: error.message};
        } else if (!error) {
            infoText = {
                error: false,
                text: "An email has been sent to you for verification!",
            };
        }
    };
</script>

<style>
    [role='alert'] {
        background: indianred;
    }
    form {
        display: grid;
        grid-template-columns: [labels] auto [controls] 1fr;
        grid-auto-flow: row;
    }
    form > label  {
        grid-column: labels;
        grid-row: auto;
        padding-right: 1em;
    }
    form > input,
    form > textarea,
    form > button {
        grid-column: controls;
        grid-row: auto;
        border: none;
        padding: 1em;
    }
    div {
        display: flex;
        justify-content: center;
        gap: 1em;
    }
</style>

<form on:submit|preventDefault>
    <label for='email'>Email:</label>
    <input
            id='email'
            type='email'
            name='email'
            bind:value={email}
            required
    />
    <label for='password'>Password:</label>
    <input
            id='password'
            type='password'
            name='password'
            bind:value={password}
            required
    />
    {#if !!infoText.text}
        <section role={infoText.error ? "alert" : null}>
            {infoText.text}
        </section>
    {/if}
    <button
            type='submit'
            on:click={() => handleLogin()}
            disabled={disabled}
            value='Sign up'>
        Sign Up
    </button>
</form>
