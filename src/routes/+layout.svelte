<script lang="ts">
    import NavBar from '$lib/components/NavBar.svelte'
    import logo from '$lib/assets//logo.png'
    import bahunya from '$lib/assets/bahunya.min.css'
    import puzzad from '$lib/assets/puzzad.css'
    import {FlatToast, ToastContainer} from "svelte-toasts";
</script>
<NavBar logo="{logo}"/>
<slot/>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
    <FlatToast {data}/>
</ToastContainer>