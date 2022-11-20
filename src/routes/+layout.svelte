<script lang="ts">
  import NavBar from '$components/NavBar.svelte'
  import logo from '$assets/logo.png'
  import bahunya from '$assets/bahunya.min.css'
  import puzzad from '$assets/puzzad.css'
  import {FlatToast, ToastContainer} from 'svelte-toasts'
  import {title} from '$lib/title'
  export const ssr = false;
</script>
<svelte:head>
  <title>{$title}</title>
</svelte:head>
<NavBar logo="{logo}"/>
<main>
<slot/>
</main>
  <ToastContainer placement="bottom-right" theme="dark" let:data={data}>
  <FlatToast {data}/>
</ToastContainer>
<footer>
  <p>&copy; 2022 Puzzad.com. All rights reserved. <a href="/about">About &amp; Privacy</a></p>
</footer>