<script lang='ts'>
    import {supabase} from "$lib/db";
    import {toasts, ToastContainer, FlatToast} from "svelte-toasts";

    let email: string = "";
    let password: string = "";
    let disabled = false
    let data = {}

    const handleLogin = async () => {
        disabled = true
        const { _, error, } = await supabase.auth.signUp({email, password});
        disabled = false
        if (error) {
            toasts.add({
                title: 'Signup',
                description: error.message,
                duration: 10000,
                type: 'error',
            })
        } else {
            toasts.add({
                title: 'Signup',
                description: "An email has been sent to you for verification!",
                duration: 10000,
                type: 'success',
            })
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
            disabled={disabled}
            value='Sign up'>
        Sign Up
    </button>
</form>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
    <FlatToast {data}/>
</ToastContainer>