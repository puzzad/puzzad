<script type='ts'>
    import {supabase} from "$lib/db.ts";
    import {replace} from 'svelte-spa-router'
    import {toasts, ToastContainer, FlatToast} from "svelte-toasts";
    import {AuthError} from "@supabase/gotrue-js/src/lib/errors.ts";
    import {logout, session} from "$lib/auth";
    import {title} from "$lib/title.ts";

    export let type: string = "login"

    let email: string
    let password: string
    let disabled: boolean
    let altText: string
    let buttonText: string
    export const authAction = async () => {
        let successText: string
        let error: AuthError
        switch (type) {
            case "signup":
                ({error} = await supabase.auth.signUp({email, password}))
                successText = "An email has been sent to you for verification!"
                break;
            case "login":
                ({error} = await supabase.auth.signInWithPassword({email, password}))
                successText = "Login success, redirecting."
                await replace('/')
                break;
            case "logout":
                break
            }
        if (error) {
            toasts.add({
                title: 'Error',
                description: error.message,
                duration: 10000,
                type: 'error',
            })
        } else {
            toasts.add({
                title: 'Success',
                description: successText,
                duration: 10000,
                type: 'success',
            })
        }
    }
    switch (type) {
        case "signup":
            title.set("Puzzad: Signup")
            altText = "Already have an account? <a href='#/Login'>Login!</a>"
            buttonText = "Sign up"
            break;
        case "login":
            title.set("Puzzad: Login")
            altText = "Don't have an account? <a href='#/signup'>Sign up!</a>"
            buttonText = "Sign In"
            break;
        case "logout":
            title.set("Puzzad: Logout");
            logout()
    }
</script>
{#if !$session?.user}
    <form class='basic' on:submit|preventDefault={authAction}>
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
                disabled={disabled}>
            {buttonText}
        </button>
        <section>
            {@html altText}
        </section>
    </form>
{:else}
    <p>You're already logged in, <a href='' on:click|preventDefault={logout}>Logout</a>?</p>
{/if}
<ToastContainer placement='bottom-right' theme='dark' let:data={data}>
    <FlatToast {data} />
</ToastContainer>