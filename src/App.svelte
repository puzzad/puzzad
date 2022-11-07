<script lang="ts">
    import bahunya from './assets/bahunya.min.css'
    import puzzad from './assets/puzzad.css'
    import logo from './assets/logo.png'
    import Router from 'svelte-spa-router'
    import {wrap} from 'svelte-spa-router/wrap'
    import NavBar from '$comps/NavBar.svelte'
    import {logout} from "$lib/auth.ts";
    import {title} from '$lib/title.ts'
    import {FlatToast, ToastContainer} from "svelte-toasts";

    const routes = {
        '/adventures': wrap({ asyncComponent: () => import('$routes/Adventures.svelte')}),
        '/adventure/:name': wrap({ asyncComponent: () => import('$routes/Adventure.svelte')}),
        '/games': wrap({ asyncComponent: () => import('$routes/Games.svelte')}),
        '/game/:code': wrap({ asyncComponent: () => import('$routes/Game.svelte')}),
        '/game/:code/:puzzle': wrap({ asyncComponent: () => import('$routes/Puzzle.svelte')}),
        '/login': wrap({ props: { type: "login" }, asyncComponent: () => import('$comps/Auth.svelte')}),
        '/logout': wrap({ props: { type: "logout" }, asyncComponent: () => import('$comps/Auth.svelte')}),
        '/signup': wrap({ props: { type: "signup" }, asyncComponent: () => import('$comps/Auth.svelte')}),
        '/': wrap({ asyncComponent: () => import('$routes/Home.svelte')}),
        '*': wrap({ asyncComponent: () => import('$routes/NotFound.svelte')}),
    }
</script>
<svelte:head><title>{$title}</title></svelte:head>
<NavBar logo='{logo}' />
<Router {routes}/>
<ToastContainer placement="bottom-right" theme="dark" let:data={data}>
    <FlatToast {data}/>
</ToastContainer>