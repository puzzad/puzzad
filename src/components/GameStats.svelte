<script lang="ts">
    import {formatDuration} from '$lib/time'
  import Error from '$components/Error.svelte'
    import {getGameClient} from '$lib/api.ts'

  export let code

  let stats = getGameClient(code).
      then(client => client.collection("solvetimes").getFullList()).
      then(solveTimes => {
        solveTimes.forEach(solveTime => {
          solveTime.time = formatDuration(solveTime.gameStart, solveTime.timeSolved)
        })
        return solveTimes
      })
</script>

{#await stats then puzzles}
  <table class="stats">
    <thead>
    <tr>
      <th>Puzzle</th>
      <th>Time</th>
      <th>Hints</th>
    </tr>
    </thead>
    <tbody>
    {#each puzzles as puzzle}
      <tr>
        <td>{puzzle.title}</td>
        <td class="time">{puzzle.time}</td>
        <td>{puzzle.usedhints}</td>
      </tr>
    {/each}
    </tbody>
  </table>
{:catch error}
  <Error error={error}></Error>
{/await}