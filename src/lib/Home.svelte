<script>
    import {supabase} from './db'
    import { onMount } from "svelte";
    export let user
    let adventures = []
    onMount(async function () {
        let {data: obtained, error}  = await supabase
                .from('adventures')
                .select('id,name')
        if (!error) {
            adventures = obtained
        }
    });

</script>

<div>
    <button
            on:click={async () => {
                const { error } = await supabase.auth.signOut();
                if (error) console.log("Error logging out:", error.message);
            }}>
        Logout
    </button>
</div>
<div>
    <ul>
    {#each adventures as adventure}
       <li>{adventure.name}</li>
    {/each}
    </ul>
</div>
