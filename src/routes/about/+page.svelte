<script lang="ts">
  import {title} from '../../lib/title'
  import DatabasePage from '$components/DatabasePage.svelte'
  export const csr = false;
  title.set('Puzzad: About & Privacy')
</script>

<style>
  div :global(section) {
    background-color: #222222;
    border-radius: 15px;
    padding: 1em;
    margin: 0 0 1em 0;
  }

  div :global(h3) {
    border-bottom: 2px solid #e3bc5e;
    padding: 0;
    margin: 0;
    font-weight: bold;
    font-size: large;
  }

  div :global(h4) {
    border-bottom: 1px solid #e3bc5e;
    padding: 0;
    margin: 0.5em 0;
    font-weight: bold;
    font-size: medium;
  }

  div :global(h5) {
    padding: 0;
    margin: 1em 0;
    font-weight: bold;
    font-size: medium;
  }
</style>

<h2>About &amp; Privacy</h2>
<div class="about">
  <DatabasePage page="privacy"></DatabasePage>
</div>