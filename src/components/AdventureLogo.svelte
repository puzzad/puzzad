<script lang="ts">
  import {client} from '$lib/api'

  export let name = ''

  let logoURL = client.collection('adventures').getFirstListItem('name=\''+name+'\'')
  .then(response => {
    return client.files.getUrl(response, response.logo)
  })
</script>

{#await logoURL}
{:then logoURL}
<img src="{logoURL}" alt="{name}">
{/await}