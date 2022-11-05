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

<form class='basic' on:submit|preventDefault>
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
    <button
            type='submit'
            on:click={() => handleLogin()}
            disabled={disabled}
            value='Sign up'>
        Sign Up
    </button>
    {#if !!infoText.text}
        <section role={infoText.error ? "alert" : null}>
            {infoText.text}
        </section>
    {/if}
</form>
