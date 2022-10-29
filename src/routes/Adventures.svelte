<script>
    import {supabase} from '$lib/db'
    import { onMount } from "svelte";
    let adventures = []
    onMount(async function () {
        let {data: obtained, error}  = await supabase
                .from('adventures')
                .select(`
                    id,name,description,price
                `)
        if (!error) {
            adventures = obtained
        }
    });
</script>

<div>
    <ul>
        {#each adventures as adventure}
            <li>
                {adventure.name} - {adventure.description}
            </li>
        {/each}
    </ul>
</div>