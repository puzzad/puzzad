<script lang="ts">
    import {onMount} from "svelte"
    import {supabase} from "$lib/db"
    import type {User} from "@supabase/supabase-js"
    import Home from "$lib/Home.svelte"
    import Router from 'svelte-spa-router'
    import Adventure from "$lib/Adventure.svelte"
    import NotFound from "./routes/NotFound.svelte"
    import Login from "./routes/Login.svelte";
    import Logout from "./routes/Logout.svelte";
    import Adventures from "./routes/Adventures.svelte";
    import Signup from "./routes/Signup.svelte";

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
    nav a:last-of-type {
        display: flex;
        flex-grow: 2;
        justify-content: end;
    }
</style>
<nav>
    <a href="#/">Puzzad</a>
    <a href="#/Adventures">Adventures</a>
    {#if !user}
        <a href='#/login'>Login</a>
    {:else}
        <a href='#/logout'>Logout</a>
    {/if}
</nav>
<Router {routes}/>