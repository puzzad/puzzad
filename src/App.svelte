<script lang="ts">
    import classes from './assets/bahunya.min.css'
    import logo from './assets/logo.png'
    import {onMount} from "svelte"
    import {supabase} from "$lib/db"
    import type {User} from "@supabase/supabase-js"
    import Router from 'svelte-spa-router'
    import {wrap} from 'svelte-spa-router/wrap'

    let user: User;

    onMount(() => {
        supabase.auth.getSession().then(({data: {session}}) => {
            user = session?.user ?? null;
        });

        const {subscription: authListener} = supabase.auth.onAuthStateChange(
            (event, session) => {
                const currentUser = session?.user;
                user = currentUser ?? null;
            }
        );

        return () => {
            authListener?.unsubscribe();
        };
    });
    const routes = {
        '/adventures': wrap({ asyncComponent: () => import('./routes/Adventures.svelte')}),
        '/adventure/:name': wrap({ asyncComponent: () => import('./routes/Adventure.svelte')}),
        '/games': wrap({ asyncComponent: () => import('./routes/Games.svelte')}),
        '/game/:code': wrap({ asyncComponent: () => import('./routes/Game.svelte')}),
        '/game/:code/:puzzle': wrap({ asyncComponent: () => import('./routes/Puzzle.svelte')}),
        '/login': wrap({ asyncComponent: () => import('./routes/Login.svelte')}),
        '/logout': wrap({ asyncComponent: () => import('./routes/Logout.svelte')}),
        '/signup': wrap({ asyncComponent: () => import('./routes/Signup.svelte')}),
        '/': wrap({ asyncComponent: () => import('./routes/Home.svelte')}),
        '*': wrap({ asyncComponent: () => import('./routes/NotFound.svelte')}),
    }
</script>
<style>
    nav {
        padding: 0 0 2px 0;
        align-items: stretch;
    }

    nav ul {
        display: flex;
        align-items: stretch;
    }

    nav ul:first-of-type {
        margin-left: 1em;
    }

    nav ul:last-of-type {
        flex-grow: 2;
        justify-content: end;
    }

    nav li {
        display: flex;
        flex-direction: column;
        justify-content: end;
    }

    nav li a {
        margin-bottom: 0;
        border-bottom: 1px dashed var(--links);
        font-variant: all-small-caps;
        font-size: large;
        line-height: 1em;
    }

    nav li a:hover {
        text-decoration: none;
        filter: brightness(1.5);
    }

    h1, h1 a {
        margin: 0;
        padding: 0;
        line-height: 0;
    }

    h1 img {
        height: 1em;
    }
</style>
<nav>
    <h1><a href="/#/"><img src="{logo}" alt="Puzzad"></a></h1>
    <ul>
        <li><a href="/#/Adventures">Browse adventures</a></li>
    </ul>
    <ul>
    {#if !user}
        <li><a href='#/login'>Login</a></li>
    {:else}
        <li><a href='/#/games'>My Games</a></li>
        <li><a href='/#/logout'>Logout</a></li>
    {/if}
    </ul>
</nav>
<Router {routes}/>