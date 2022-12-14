<script>
  import {supabase} from '$lib/db'
  import Spinner from '$components/Spinner.svelte'

  let users = new Promise(async (resolve, reject) => {
    const { data, error } = await supabase.rpc('listallusers').throwOnError()
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
  <h2>Users</h2>
  {#await users}
    <Spinner/>
  {:then users}
    <table>
      <thead>
      <tr>
        <th>Email</th>
        <th>Created</th>
        <th>Confirmed</th>
        <th>Last Sign In</th>
        <th>Adventures</th>
      </tr>
      </thead>
      <tbody>
      {#each users as user}
        <tr>
          <td>{user.email}</td>
          <td>{new Date(user.created).toLocaleString()}</td>
          <td>{new Date(user.confirmed).toLocaleString()}</td>
          <td>{new Date(user.lastSignin).toLocaleString()}</td>
          <td>
            <table>
            {#each user.adventures as adventure}
              <tr>
                <td>{adventure.adventure}</td>
                <td>{adventure.code}</td>
              </tr>
              {:else}
              <tr>
                <td colspan="2">No Adventures</td>
              </tr>
            {/each}
            </table>
          </td>
        </tr>
        {/each}
      </tbody>
    </table>
  {/await}
</section>