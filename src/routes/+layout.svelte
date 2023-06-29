<script lang="ts">
  import '../../src/style/global.scss'
  import {ToastContainer} from 'svelte-toasts'
  import {title} from '$lib/title'
  import {logout, isAdmin, currentUser} from '$lib/api'
  import Nav from '$components/Nav.svelte'
  import Footer from '$components/Footer.svelte'
  import Toast from '$components/Toast.svelte'
</script>
<svelte:head>
  <title>{$title}</title>
</svelte:head>
<Nav loggedIn={$currentUser} on:logout={logout} />
<main>
  <slot/>
</main>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
  <Toast type={data.type} title={data.title} message={data.description}/>
</ToastContainer>

<Footer isAdmin={$isAdmin}/>