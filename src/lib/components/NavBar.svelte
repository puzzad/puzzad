<script lang="ts">
  import {logout, session} from '$lib/auth'
  import {browser} from '$app/environment'

  export let logo
  let dropdownMenu = false
  if (browser) {
    dropdownMenu = window.matchMedia('(max-width: 480px)').matches
    window.matchMedia('(max-width: 480px)').addEventListener('change', ev => dropdownMenu = ev.matches)
  }
  let showDropdownMenu = false
  const handleMobileIconClick = () => showDropdownMenu = !showDropdownMenu
</script>
<style>
  nav {
    align-items: stretch;
    display: flex;
    padding: 0 0 2px 0;
  }

  nav ul {
    align-items: stretch;
    display: flex;
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
    border-bottom: 1px dashed var(--links);
    font-size: large;
    font-variant: all-small-caps;
    line-height: 1em;
    margin-bottom: 0;
  }

  nav li a:hover {
    filter: brightness(1.5);
    text-decoration: none;
  }

  h1, h1 a {
    line-height: 0;
    margin: 0;
    padding: 0;
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

  .hamburgercontainer:hover {
    cursor: pointer;
  }

  .hamburger {
    box-shadow: inset 0 0 0 32px, 0 -6px, 0 6px;
    height: 2px;
    margin: 16px 7px;
    width: 20px;
  }
</style>
<nav>
  <h1><a href="/"><img alt="Puzzad" src="{logo}"></a></h1>
  {#if dropdownMenu}
    <div class="hamburgercontainer" on:click={handleMobileIconClick}><span class="hamburger"></span></div>
    <dialog open="{showDropdownMenu}" on:click={() => showDropdownMenu=false}>
      <ul class='{showDropdownMenu ? "dropdown" : ""}'>
        <li><a href="/adventures">Browse adventures</a></li>
      </ul>
      <ul>
        {#if !$session?.user}
          <li><a href="/login">Login</a></li>
        {:else}
          <li><a href="/games">My Games</a></li>
          <li><a on:click={logout()}>Logout</a></li>
        {/if}
      </ul>
    </dialog>
  {/if}
  {#if !dropdownMenu}
    <ul class='{showDropdownMenu ? "dropdown" : ""}'>
      <li><a href="/adventures">Browse adventures</a></li>
    </ul>
    <ul>
      {#if !$session?.user}
        <li><a href="/login">Login</a></li>
      {:else}
        <li><a href="/games">My Games</a></li>
        <li><a href="" on:click|preventDefault={logout}>Logout</a></li>
      {/if}
    </ul>
  {/if}
</nav>