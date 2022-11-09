<script>
    import {goto} from '$app/navigation'
    import {supabase} from '$lib/db.ts'
    import {toasts} from 'svelte-toasts'
    import {logout, session} from '$lib/auth.ts'
    import {title} from '$lib/title.ts'
    import FaGoogle from 'svelte-icons/fa/FaGoogle.svelte'
    import FaDiscord from 'svelte-icons/fa/FaDiscord.svelte'
    import FaTwitch from 'svelte-icons/fa/FaTwitch.svelte'

    export let type = 'login'

    let email
    let password
    let disabled
    let altText
    let buttonText
    export const authAction = async () => {
        let successText
        let error
        switch (type) {
            case 'signup':
                ({error} = await supabase.auth.signUp({email, password}))
                successText = 'An email has been sent to you for verification!'
                break
            case 'login':
                ({error} = await supabase.auth.signInWithPassword({email, password}))
                successText = 'Login success.'
                await goto('/')
                break
            case 'logout':
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
    const oauthAction = async (name) => {
        const {error} = await supabase.auth.signInWithOAuth({provider: name})
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
                description: 'Login success.',
                duration: 10000,
                type: 'success',
            })
        }
    }
    switch (type) {
        case 'signup':
            title.set('Puzzad: Signup')
            altText = 'Already have an account? <a href=\'/login\'>Login!</a>'
            buttonText = 'Sign up'
            break
        case 'login':
            title.set('Puzzad: Login')
            altText = 'Don\'t have an account? <a href=\'/signup\'>Sign up!</a>'
            buttonText = 'Sign In'
            break
        case 'logout':
            title.set('Puzzad: Logout')
            logout()
    }
</script>
<style>
    .icon {
        width: 32px;
        height: 32px;
        cursor: pointer;
    }

    .socialicons {
        display: flex;
        flex-direction: row;
        padding: 0.5em;
        justify-content: center;
        gap: 1em;
        margin: 0;
    }

    .socialwrapper {
        grid-column: controls;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        margin: 0;
        padding-bottom: 1em;
    }

    p {
        margin: 0;
        padding: 0;
    }
</style>
{#if !$session?.user}
    <form class="basic" on:submit|preventDefault={authAction}>
        <div class="socialwrapper">
            <p>Sign in with one of these third party providers</p>
            <div class="socialicons">
                <div class="icon" on:click={() => oauthAction("google")}>
                    <FaGoogle/>
                </div>
                <div class="icon" on:click={() => oauthAction("discord")}>
                    <FaDiscord/>
                </div>
                <div class="icon" on:click={() => oauthAction("twitch")}>
                    <FaTwitch/>
                </div>
            </div>
            <p>or</p>
            <p>{buttonText} with email</p>
        </div>
        <label for="email">Email:</label>
        <input
                id="email"
                type="email"
                name="email"
                bind:value={email}
                required
        />
        <label for="password">Password:</label>
        <input
                id="password"
                type="password"
                name="password"
                bind:value={password}
                required
        />
        <button
                type="submit"
                disabled={disabled}>
            {buttonText}
        </button>
        <section>
            {@html altText}
        </section>
    </form>
{:else}
    <p>You're already logged in, <a href="" on:click|preventDefault={logout}>Logout</a>?</p>
{/if}
