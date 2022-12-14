<script>
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {DateTime} from 'luxon'

  let games = supabase.rpc('listallgames')
  .then((data, error) => {
    if (error) {
      throw Promise.reject('No valid token')
    } else {
      return data.data
    }
  })
  .catch(_ => [])

  const settings = {columnFilter: true}
</script>
<style>
  section {
    margin-left: calc(((90vw - 50rem )/ 2) * -1);
    width: 90vw;
  }
</style>
<section>
  <h2>Games</h2>
  {#await games}
    <Spinner/>
  {:then games}
    <table>
      <thead>
      <tr>
        <th>User</th>
        <th>Code</th>
        <th>Adventure</th>
        <th>Progress</th>
        <th>Start Time</th>
        <th>End Time</th>
        <th>Time since starting or Duration</th>
      </tr>
      </thead>
      <tbody>
      {#each games as game}
        <tr>
          <td>{game.useremail}</td>
          <td>{game.code}</td>
          <td>{game.adventurename}</td>
          <td>{game.puzzletitle ?? "Not started"}</td>
          <td>{new Date(game.startTime).toLocaleString()}</td>
          <td>
            {#if (game.endTime)}
              {new Date(game.endTime).toLocaleString()}
            {:else} In Progress
            {/if}
          </td>
          <td>
            {#if game.endTime }
              {DateTime.fromISO(game.endTime).
                  diff(DateTime.fromISO(game.startTime), ['days', 'hours', 'minutes']).
                  toHuman({listStyle: "short", unitDisplay: "short", maximumSignificantDigits: 2})}
            {:else}
              {DateTime.now().
                  diff(DateTime.fromISO(game.startTime), ['days', 'hours', 'minutes']).
                  toHuman({unitDisplay: "short", maximumSignificantDigits: 2})}
            {/if}
          </td>
        </tr>
      {/each}
      </tbody>
    </table>
  {/await}
</section>