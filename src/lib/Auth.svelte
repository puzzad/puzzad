<script lang='ts'>
    import {supabase} from "./db";

    interface InfoText {
        error: boolean;
        text: string | null;
    }

    let email: string = "";
    let password: string = "";
    let infoText: InfoText = {error: null, text: null};
    let disabled = false

    const handleLogin = async (type) => {
        disabled = true
        const { _, error, } =
                type === "LOGIN"
                        ? await supabase.auth.signInWithPassword({email, password})
                        : await supabase.auth.signUp({email, password});
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
</style>

<div>
    <label for='email'>Email:</label>
    <input
            id='email'
            class='bg-gray-100 border py-1 px-3'
            type='email'
            name='email'
            bind:value={email}
            required
    />
    <label for='password'>Password:</label>
    <input
            id='password'
            class='bg-gray-100 border py-1 px-3'
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
    <div class='col-start-2'>
        <div>
            <button
                    type='submit'
                    on:click={() => handleLogin("REGISTER")}
                    disabled={disabled}
                    value='Sign up'>
                Sign Up
            </button>
            <button
                    on:click={() => handleLogin("LOGIN")}
                    disabled={disabled}
                    type='button'>
                Sign In
            </button>
        </div>
    </div>
</div>
