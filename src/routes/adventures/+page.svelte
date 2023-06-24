<script lang="ts">
  import AdventureBanner from '$components/AdventureBanner.svelte'
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import Error from '$components/Error.svelte'

  let adventures = supabase.from('adventures').
      select('id,name,price,public').
      throwOnError().
      then(({data}) => data)

  let games = supabase.from('games').
      select('id, status, adventures (name, public), status, code').
      throwOnError().
      then(({data}) => data)

  title.set('Puzzad: Adventures')
</script>

{#await games}
  <Spinner/>
{:then games}
  {@const finishedGames = games.filter(g => g.status === 'EXPIRED')}
  {@const activeGames = games.filter(g => g.status !== 'EXPIRED')}
  {#if games.length > 0}
    <h2>Your adventures</h2>
    {#each activeGames as game}
      <AdventureBanner
          status="{game.status}"
          adventureName='{game.adventures?.name ?? "Unknown"}'
          code="{game.code}"
          isPublic="{game.adventures?.public}"
      />
    {/each}
    <details>
      <summary>{finishedGames.length} finished adventure{finishedGames.length === 1 ? '' : 's'}</summary>
      {#each finishedGames as game}
        <AdventureBanner
            status="{game.status}"
            adventureName='{game.adventures?.name ?? "Unknown"}'
            code="{game.code}"
            isPublic="{game.adventures?.public}"
        />
      {/each}
    </details>
  {/if}
{:catch error}
  <Error error={error}></Error>
{/await}

<h2>Want to start a new adventure?</h2>
{#await adventures}
  <Spinner/>
{:then adventures}
  {#each adventures as adventure}
    <AdventureBanner
        adventureName="{adventure.name}"
        price="{adventure.price}"
        isPublic="{adventure.public}"
        code=""
        status=""
    />
  {:else}
    <p>No Adventures, sorry.</p>
  {/each}
{:catch error}
  <Error error={error}></Error>
{/await}