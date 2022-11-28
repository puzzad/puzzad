<script lang="ts">
  import {title} from '$lib/title.ts'
  import SubscribeToMailingList from '$components/SubscribeToMailingList.svelte'

  title.set('Puzzad')
</script>
<style>
  #container {
    display: flex;
    flex-direction: column;
    align-content: center;
  }

  p {
    margin: 0.5em 0 3em;
  }
</style>

<h2>Puzzad: Puzzle Adventures</h2>
<div id="container">
  <p>
    Please mind the dust &mdash; we're not quite ready to start any adventures yet.
    Check back soon, or if you'd like to be notified when we go live, sign up for our mailing list below!
  </p>
  <SubscribeToMailingList></SubscribeToMailingList>
</div>