<script lang="ts">
  import {title} from '../../lib/title'
  import DatabasePage from '$components/DatabasePage.svelte'

  title.set('Puzzad: About & Privacy')
</script>

<style>
  main :global(section) {
    background-color: #222222;
    border-radius: 15px;
    padding: 1em;
    margin: 0 0 1em 0;
  }

  main :global(h3) {
    border-bottom: 2px solid #e3bc5e;
    padding: 0;
    margin: 0;
    font-weight: bold;
    font-size: large;
  }

  main :global(h4) {
    border-bottom: 1px solid #e3bc5e;
    padding: 0;
    margin: 0.5em 0;
    font-weight: bold;
    font-size: medium;
  }

  main :global(h5) {
    padding: 0;
    margin: 1em 0;
    font-weight: bold;
    font-size: medium;
  }
</style>

<h2>About &amp; Privacy</h2>
<main>
  <DatabasePage page="privacy"></DatabasePage>
</main>