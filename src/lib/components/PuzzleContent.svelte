<script lang="ts">
    import {getGameClient} from "$lib/db.ts";
    import Spinner from "$lib/components/Spinner.svelte";
    import Error from "$lib/components/Error.svelte";

    export let gameCode = '';
    export let storageSlug = '';
    export let content = '';

    const embedRegex = /\$[^$]+?\$/g

    let html = Promise
        .all([
            getGameClient(gameCode),
            storageSlug,
            content,
        ])
        .then(([gameClient, storageSlug, content]) => {
            const fileNames = content.match(embedRegex) || []
            return Promise.all([
                content,
                ...fileNames
                    .map((f) => f.substring(1, f.length - 1))
                    .map((f) => gameClient
                        .storage
                        .from('puzzles')
                        .createSignedUrl(`${storageSlug}/${f}`, 60 * 60)
                        .then(({data: {signedUrl}, error}) => {
                            if (error) {
                                throw error
                            }
                            return signedUrl
                        })
                    ),
            ])
        })
        .then(([content, ...urls]) => content.replace(embedRegex, () => urls.shift()))
</script>

{#await html}
  <Spinner/>
{:then html}
  {@html html}
{:catch error}
  <Error error={error}></Error>
{/await}