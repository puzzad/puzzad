<script lang="ts">
  import {pb} from '$lib/auth'

  export let name = ''

  let logoURL = pb.collection('adventures').getFirstListItem('name=\''+name+'\'')
  .then(response => {
    return pb.files.getUrl(response, response.logo)
  })
</script>

{#await logoURL}
{:then logoURL}
<img src="{logoURL}" alt="{name}">
{/await}