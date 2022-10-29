<script>
    import {supabase} from '$lib/db'
    import {onMount} from 'svelte'

    let adventures = []
    onMount(async function () {
        let {data: obtained, error} = await supabase
                .from('adventures')
                .select(`
                    id,name,description,price
                `)
        if (!error) {
            adventures = obtained
        }
    })
</script>

{#each adventures as adventure}
    <article>
        <h2>{adventure.name}</h2>
        <p>{adventure.description}</p>
    </article>
{:else}
    <p>No Adventures, sorry.</p>
{/each}