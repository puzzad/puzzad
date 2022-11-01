<script>
    import {supabase} from "$lib/db";
    import {replace} from 'svelte-spa-router'

    async function signOut() {
        const { error } = await supabase.auth.signOut()
        if (!error) {
            await replace("/");
        }
    }
    import { onMount } from "svelte";
    onMount(async function () {
        await signOut()
    });
</script>
<section>
    Logging out...
</section>