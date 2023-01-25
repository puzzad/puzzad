<script>
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import {formatDuration} from '$lib/time'

  title.set('Puzzad: Admin: Games')

  let games = new Promise(async (resolve, reject) => {
    const { data, error } = await supabase.rpc('listallgames').throwOnError()
    if (error) {
      reject(error)
    } else {
      resolve(data)
    }
  })
  .then(data => data)
  .catch(_ => [])
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
        <th>Puzzle</th>
        <th>Start Time</th>
        <th>End Time</th>
        <th>Duration</th>
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
            {#if game.endTime}
              {new Date(game.endTime).toLocaleString()}
            {:else} In Progress
            {/if}
          </td>
          <td>
            {#if game.endTime }
              {formatDuration(game.startTime, game.endTime)}
            {:else}
              {formatDuration(game.startTime, Date())}
            {/if}
          </td>
        </tr>
      {/each}
      </tbody>
    </table>
  {/await}
</section>