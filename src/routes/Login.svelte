<script lang='ts'>
    import {supabase} from "$lib/db";
    import {replace} from 'svelte-spa-router'
    import {toasts, ToastContainer, FlatToast} from "svelte-toasts";

    let email: string = "";
    let password: string = "";
    let disabled = false
    let data = {}

    const handleLogin = async () => {
        disabled = true
        const {error} = await supabase.auth.signInWithPassword({email, password})
        disabled = false
        if (error) {
            toasts.add({
                title: 'Login',
                description: error.message,
                duration: 10000,
                type: 'error',
            })
        } else {
            await replace('#/')
        }
    };
</script>

<form class='basic' on:submit|preventDefault={handleLogin}>
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
            disabled={disabled}>
        Sign In
    </button>
    <section>
        Don't have an account? <a href='#/signup'>Sign up!</a>
    </section>
</form>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
    <FlatToast {data}/>
</ToastContainer>