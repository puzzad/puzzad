<script lang="ts">
  export let logo
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
</style>
<nav>
  <h1><a href="/"><img width="178" height="40" alt="Puzzad" src="{logo}"></a></h1>
</nav>