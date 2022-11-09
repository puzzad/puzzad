<script lang="ts">
    import {getGameClient} from "$lib/db.ts"
    import {formatDuration} from "$lib/time.ts"
    import Error from '$lib/components/Error.svelte'

    export let code
    export let startTime
    let stats = getGameClient(code)
        .then((gc) => gc.rpc('getstats', {gamecode: code}).throwOnError())
        .then(({data: stats}) => stats)
        .then((rows) => {
            let lastTime = startTime
            rows.forEach((r) => {
                r.time = formatDuration(lastTime, r.solvetime)
                lastTime = r.solvetime
            })
            return rows
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
        <td>{puzzle.hints}</td>
      </tr>
    {/each}
    </tbody>
  </table>
{:catch error}
  <Error error={error}></Error>
{/await}