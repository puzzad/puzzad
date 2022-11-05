<script lang="ts">
    import classes from './assets/bahunya.min.css'
    import logo from './assets/logo.png'
    import {onMount} from "svelte"
    import {supabase} from "$lib/db"
    import type {User} from "@supabase/supabase-js"
    import Router from 'svelte-spa-router'
    import {wrap} from 'svelte-spa-router/wrap'
    import NavBar from '$lib/NavBar.svelte'

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
<style global>
    table {
        display: table;
    }
    select {
        background: var(--background);
        -webkit-appearance: menulist;
    }
    [role='alert'] {
        background: indianred;
    }
    @media (max-width: 480px) {
        form.basic {
            display: grid;
            grid-template-columns: auto;
            grid-auto-flow: row;
        }
        form.basic > label  {
            grid-row: auto;
        }
        form.basic > input,
        form.basic > button {
            grid-row: auto;
            padding: 1em;
            margin: .5em 0 .5em 0;
        }
    }

    @media (min-width: 480px) {
        form.basic {
            display: grid;
            grid-template-columns: [labels] auto [controls] 1fr;
            grid-auto-flow: row;
        }
        form.basic > label  {
            grid-column: labels;
            grid-row: auto;
            padding-right: 1em;
        }
        form.basic > input,
        form.basic > button {
            grid-column: controls;
            grid-row: auto;
            border: none;
            padding: 1em;
        }
        form.basic > section {
            grid-column: controls;
            grid-row: auto;
            margin: 0;
            padding: 1em;
            border-radius: 6px;
            text-align: center;
            line-height: normal;
        }
    }
</style>
<NavBar user='{user}' logo='{logo}' />
<Router {routes}/>