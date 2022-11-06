<script lang="ts">
    import {getGameClient} from "$lib/db";
    import Spinner from "$comps/Spinner.svelte";
    import Error from "$comps/Error.svelte";

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

<style global>
    section.tip {
        position: relative;
        border-top: 10px solid #e3bc5e;
        border-bottom: 2px solid #e3bc5e;
        background-color: rgba(227, 188, 94, .3);
        padding: 0 0.7em;
    }

    section.tip::before {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "ℹ️ information";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #e3bc5e;
        color: black;
    }

    section.story {
        position: relative;
        border-top: 10px solid #1c76c5;
        border-bottom: 2px solid #1c76c5;
        background-color: rgba(28, 118, 197, .3);
        padding: 0 0.7em;
    }

    section.story::before {
        float: right;
        border-bottom-right-radius: 5px;
        border-bottom-left-radius: 5px;
        padding: 5px 10px;
        margin: 0 20px;
        content: "📖 story";
        font-weight: bold;
        font-variant: small-caps;
        background-color: #1c76c5;
        color: white;
    }
</style>

{#await html}
    <Spinner/>
{:then html}
    {@html html}
{:catch error}
    <Error error={error}></Error>
{/await}