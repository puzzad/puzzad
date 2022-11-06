<script>
    import {supabase} from "$lib/db";

    export let name = ''

    let logoUrl = name && supabase.storage.from('adventures')
        .getPublicUrl(name + '/logo.png')
        .data.publicUrl
</script>

<img src="{logoUrl}" alt="{name}">