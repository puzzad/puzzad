<script lang="ts">
    import {onMount} from "svelte"
    import {supabase} from "$lib/db"
    import type {User} from "@supabase/supabase-js"
    import Home from "./routes/Home.svelte"
    import Router from 'svelte-spa-router'
    import Adventure from "./routes/Adventure.svelte"
    import NotFound from "./routes/NotFound.svelte"
    import Login from "./routes/Login.svelte";
    import Logout from "./routes/Logout.svelte";
    import Adventures from "./routes/Adventures.svelte";
    import Signup from "./routes/Signup.svelte";
    import Games from "./routes/Games.svelte";

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
        '/adventures': Adventures,
        '/adventure/:name': Adventure,
        '/games': Games,
        '/login': Login,
        '/logout': Logout,
        '/signup': Signup,
        '/': Home,
        '*': NotFound,
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