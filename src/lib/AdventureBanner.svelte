<script>
    export let name
    export let description
    export let code
    export let status
</script>
<style>
    article {
        margin: 2rem 0;
        padding: 0;
        border-radius: 1rem 1rem 0 0;
    }

    article div {
        padding: 0.5rem;
        margin: 0;
        border-radius: 1rem 1rem 0 0;
    }

    article div.UNPAID {
        background: indianred;
    }

    article div.PAID {
        background: darkseagreen;
    }

    article div.EXPIRED {
        background: darkslategray;
    }

    article h2 {
        margin: 1rem;
        padding: 0;
    }

    article p {
        margin: 1rem;
        padding: 0;
    }
</style>
<article>
    {#if status}
        <div class='{status}'>{status}</div>
    {/if}
    <h2>{name}</h2>
    <p>{description}</p>
    {#if code}
        <hr>
        <p>Code: {code}</p>
    {/if}
</article>