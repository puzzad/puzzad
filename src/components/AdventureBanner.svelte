<script lang="ts">
    import {supabase} from "$lib/../lib/db";
    import AdventureLogo from "$comps/AdventureLogo.svelte";

    export let adventureName
    export let code
    export let status
    export let isPublic = true
    export let price = null

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
        overflow: clip;
        overflow-clip-margin: 5px;
    }

    a:hover {
        border-color: #ccc;
    }

    /*noinspection CssUnusedSymbol*/
    .expired::before {
        content: '';
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
    .expired::after {
        content: '';
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

    code {
        position: absolute;
        top: 15px;
        right: 15px;
        background-color: white;
        color: black;
        font-size: x-large;
        padding: 2px 1em;
        box-shadow: 0 0 5px 5px rgba(255, 255, 255, .9);
    }

    .logo {
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

    .admin {
        position: absolute;
        top: 40px;
        left: -45px;
        width: 200px;
        text-align: center;
        transform: rotate(-45deg);
        background: repeating-linear-gradient(-30deg, #ff9999 0, #ff9999 10px, white 10px, white 20px);
        font-weight: bold;
        font-size: large;
        box-shadow: 0 2px 2px black, 0 -2px 2px black;
        color: black;
    }

    .price {
        transform-origin: 50% 50%;
        transform: rotate(-45deg);
        position: absolute;
        right: -65px;
        bottom: -40px;
        background-color: #265118;
        color: white;
        width: 150px;
        height: 100px;
        text-align: center;
        font-size: x-large;
    }
</style>
<a class="{status.toLowerCase()}"
   style="background-image: url('{backgroundUrl}')"
   href="{code ? '/games/' + code : '/adventures/' + adventureName}">
    <div class="logo">
        <AdventureLogo bind:name={adventureName}></AdventureLogo>
    </div>
    {#if code}
        <code>{code}</code>
    {/if}
    {#if !isPublic}
        <div class="admin">ADMIN ONLY</div>
    {/if}
    {#if price !== null}
        <div class="price">{price === 0 ? 'Free!' : '£' + price}</div>
    {/if}
</a>