<script>
    export let user
    export let logo
    let dropdownMenu = false
    window.matchMedia('(max-width: 480px)').addEventListener('change', ev => dropdownMenu = ev.matches)
    let showDropdownMenu = false
    const handleMobileIconClick = () => showDropdownMenu = !showDropdownMenu
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

    div:last-of-type {
        display: flex;
        flex-grow: 2;
        justify-content: end;
        padding-right: 1em;
        padding-top: 0.5em;
    }

    .dropdown {
        display: block;
    }

    dialog[open] {
        animation: show 200ms linear;
    }

    @keyframes show {
        from {
            transform: scale(0.001);
        }
        to {
            transform: scale(1);
        }
    }
</style>
<nav>
    <h1><a href='/#/'><img src='{logo}' alt='Puzzad'></a></h1>
    {#if dropdownMenu}
        <div on:click={handleMobileIconClick}>🍔</div>
        <dialog open='{showDropdownMenu}' on:click={() => showDropdownMenu=false}>
            <ul class='{showDropdownMenu ? "dropdown" : ""}'>
                <li><a href='/#/Adventures'>Browse adventures</a></li>
            </ul>
            <ul>
                {#if !user}
                    <li><a href='#/login'>Login</a></li>
                {:else}
                    <li><a href='/#/games'>My Games</a></li>
                    <li><a href='/#/logout'>Logout</a></li>
                {/if}
            </ul>
        </dialog>
    {/if}
    {#if !dropdownMenu}
        <ul class='{showDropdownMenu ? "dropdown" : ""}'>
            <li><a href='/#/Adventures'>Browse adventures</a></li>
        </ul>
        <ul>
            {#if !user}
                <li><a href='#/login'>Login</a></li>
            {:else}
                <li><a href='/#/games'>My Games</a></li>
                <li><a href='/#/logout'>Logout</a></li>
            {/if}
        </ul>
    {/if}
</nav>