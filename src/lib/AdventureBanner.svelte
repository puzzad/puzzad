<script lang="ts">
    import {supabase} from "$lib/db";
    import AdventureLogo from "$lib/AdventureLogo.svelte";

    export let adventureName
    export let code
    export let status
    export let price
    if (price > 0) {
        price = "£" + price
    } else {
        price = "Free"
    }

    let backgroundUrl = adventureName && supabase.storage.from('adventures')
        .getPublicUrl(adventureName + '/background.jpg')
        .data.publicUrl
</script>
<style>
    a {
        position: relative;
        display: block;
        height: 300px;
        margin: 50px 0;
        border-radius: 10px;
        border: 5px solid transparent;
        background-origin: border-box;
        background-clip: border-box;
        transition: border-color .2s;
    }

    a:hover {
        border-color: #ccc;
    }

    /*noinspection CssUnusedSymbol*/
    a.expired::before {
        content: "";
        position: absolute;
        top: -5px;
        left: -5px;
        width: calc(100% + 10px);
        height: calc(100% + 10px);
        border-radius: 10px;
        background-color: rgba(180, 180, 180, 0.7);
        animation: shortStampSettle 0.3s ease-in;
        z-index: 1;
    }

    @keyframes shortStampSettle {
        0% {
            opacity: 0;
        }
        100% {
            opacity: 1;
        }
    }

    /*noinspection CssUnusedSymbol*/
    a.expired::after {
        content: "";
        position: absolute;
        bottom: 10%;
        right: 10px;
        width: 40%;
        height: 50%;
        background-image: url('../assets/complete.png');
        background-size: 100%;
        background-repeat: no-repeat;
        background-position: bottom right;
        transform: rotate(-10deg);
        animation: longStampSettle 0.8s ease-in;
        z-index: 2;
    }

    @keyframes longStampSettle {
        0% {
            opacity: 0;
        }
        80% {
            opacity: 0;
        }
        100% {
            opacity: 1;
        }
    }

    a code {
        position: absolute;
        top: 15px;
        right: 15px;
        background-color: white;
        color: black;
        font-size: x-large;
        padding: 2px 1em;
        box-shadow: 0 0 5px 5px rgba(255, 255, 255, .9);
    }

    :global(a.adventurebanner img) {
        max-width: 50%;
        position: absolute;
        bottom: 10%;
        left: 5%;
        animation: logoSettle 1s ease-out;
    }

    @keyframes logoSettle {
        0% {
            margin-left: 20px;
        }
        100% {
            margin-left: 0;
        }
    }
</style>
<a class="adventurebanner {status.toLowerCase()}"
   style="background-image: url('{backgroundUrl}')"
   href="{code ? '/#/game/' + code : '/#/adventure/' + adventureName}">
    <AdventureLogo bind:name={adventureName}></AdventureLogo>
    {#if code}
        <code>{code}</code>
    {/if}
</a>