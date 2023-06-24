<script>
  import Header from '$components/Header.svelte'
  import {createEventDispatcher} from 'svelte'

  const dispatch = createEventDispatcher()

  export let loggedIn
</script>

<style lang="scss">
  @use '../style/colours';
  @use '../style/fonts';

  nav {
    border-bottom: 8px solid colours.$brand;
    display: flex;
    align-items: end;
    position: relative;
    margin-bottom: 10px;
  }

  ul {
    display: flex;
    margin: 0 15px 0 15px;
    padding: 0;
    flex-grow: 1;
    justify-content: space-between;
    list-style-type: none;

    @media (max-width: 1000px) {
      flex-direction: row;
      position: absolute;
      bottom: -35px;
      left: 0;
      right: 0;
      margin: 0;
      gap: 10px;
    }
  }

  a {
    font-family: fonts.$header;
    font-weight: 800;
    font-size: 20px;
    line-height: 20px;
    text-transform: uppercase;
    color: colours.$link;
    text-decoration: none;
  }
</style>

<nav>
  <Header/>

  <ul>
    <li><a href="/adventures">Start an adventure</a></li>
    {#if loggedIn}
      <li><a href="" on:click|preventDefault={() => dispatch('logout')}>Logout</a></li>
    {:else}
      <li><a href="/login">Login</a></li>
    {/if}
  </ul>
</nav>