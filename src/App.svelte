<script lang="ts">
    import classes from './assets/bahunya.min.css'
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
        '/login': wrap({ asyncComponent: () => import('./routes/Login.svelte')}),
        '/logout': wrap({ asyncComponent: () => import('./routes/Logout.svelte')}),
        '/signup': wrap({ asyncComponent: () => import('./routes/Signup.svelte')}),
        '/': wrap({ asyncComponent: () => import('./routes/Home.svelte')}),
        '*': wrap({ asyncComponent: () => import('./routes/NotFound.svelte')}),
    }
</script>
<style>
    nav:first-of-type {
        padding: 0;
    }
    nav ul:last-of-type {
        display: flex;
        flex-grow: 2;
        justify-content: end;
    }
</style>
<nav>
    <ul>
        <li><a href="#/">Puzzad</a></li>
        <li><a href="#/Adventures">Adventures</a></li>
    </ul>
    <ul>
    {#if !user}
        <li><a href='#/login'>Login</a></li>
    {:else}
        <li><a href='#/games'>Games</a></li>
        <li><a href='#/logout'>Logout</a></li>
    {/if}
    </ul>
</nav>
<Router {routes}/>