<script lang="ts">
  import {title} from '$lib/title'
  import DatabasePage from '$components/DatabasePage.svelte'

  title.set('Puzzad: About & Privacy')
</script>

<h2>About &amp; Privacy</h2>
<div class="about">
  <DatabasePage page="privacy"></DatabasePage>
</div>