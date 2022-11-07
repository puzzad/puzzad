<script lang='ts'>
    import NavBar from '$comps/NavBar.svelte'
    import logo from '../assets/logo.png'
    import bahunya from '../assets/bahunya.min.css'
    import puzzad from '../assets/puzzad.css'
    import {FlatToast, ToastContainer} from "svelte-toasts";
    import {title} from '$lib/title.ts'
</script>
<svelte:head><title>{$title}</title></svelte:head>
<NavBar logo='{logo}' />
<slot />
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
    <FlatToast {data}/>
</ToastContainer>