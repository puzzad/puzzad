<script lang="ts">
  import FaLock from 'svelte-icons/fa/FaLock.svelte'
  import FaUserFriends from 'svelte-icons/fa/FaUserFriends.svelte'
  import FaPuzzlePiece from 'svelte-icons/fa/FaPuzzlePiece.svelte'
  import FaPrint from 'svelte-icons/fa/FaPrint.svelte'
  import FaAssistiveListeningSystems from 'svelte-icons/fa/FaAssistiveListeningSystems.svelte'
  import FaLowVision from 'svelte-icons/fa/FaLowVision.svelte'
  import FaPaintBrush from 'svelte-icons/fa/FaPaintBrush.svelte'
  import FaVideo from 'svelte-icons/fa/FaVideo.svelte'
  import {goto} from '$app/navigation'
  import Spinner from '$components/Spinner.svelte'
  import AdventureLogo from '$components/AdventureLogo.svelte'
  import Error from '$components/Error.svelte'
  import {supabase} from '$lib/db'
  import {title} from '$lib/title'
  import {isLoggedIn} from '$lib/auth'

  export let data

  const features = [
    ['Difficulty', FaLock, (f) => f.difficulty],
    ['Team size', FaUserFriends, (f) => f.people],
    ['Number of puzzles', FaPuzzlePiece, (f) => f.puzzles],
    ['Equipment', FaPrint, (f) => f.equipment],
    ['Accessibility: audio', FaAssistiveListeningSystems, (f) => f.accessibility.hearing],
    ['Accessibility: visual', FaLowVision, (f) => f.accessibility.vision],
    ['Accessibility: colours', FaPaintBrush, (f) => f.accessibility.colours],
    ['Accessibility: motion', FaVideo, (f) => f.accessibility.motion],
  ]

  let details = supabase.from('adventures').
      select(`id,name,description,price,features`).
      eq('name', data.adventure).
      throwOnError().
      single().
      then(({data}) => {
        title.set('Puzzad: ' + data.name)
        data.displayFeatures = features.
            map(([name, comp, value]) => [name, comp, value(data.features)]).
            filter(([name, comp, value]) => value !== undefined)
        data.description = data.description.replace(
            /\$(.*?)\$/g,
            (r, g) => supabase.storage.
                from('adventures').
                getPublicUrl(data.name + '/' + g).
                data.publicUrl,
        )
        return data
      }).catch(_ => {
        return goto(`/adventures`)
      })

  const addAdventure = async (details) => {
    if (details.price === 0) {
      let {data, error} = await supabase.rpc('addfreeadventure', {adventureid: details.id})
      console.log(error)
      if (data) {
        await goto(`/games/${data}`)
      }
    }
  }
</script>

<style>
  @media (min-width: 480px) {
    div {
      display: grid;
      grid-template-columns: 60% auto;
      grid-gap: 2em;
    }
  }

  .sidebar p, .sidebar h4 {
    margin: 0.5em 0;
    padding: 0.2em;
  }

  .sidebar p {
    font-size: small;
  }

  .sidebar h4 {
    font-size: medium;
    border-image: linear-gradient(to right, #e3bc5e 0, #e3bc5e 50%, #222222 100%) 30 1;
    border-bottom-width: 1px;
    border-bottom-style: solid;
  }

  .sidebar h4 > :global(svg) {
    height: 0.8em;
    width: 0.8em;
    display: inline-block;
    vertical-align: baseline;
  }

  .description :global(p:first-of-type) {
    margin-top: 0;
  }

  button {
    width: 100%;
    padding: 1em;
    background-color: #3C8E2B;
  }
</style>

{#await details}
  <Spinner/>
{:then details}
  <h2>
    <AdventureLogo name={details.name}></AdventureLogo>
  </h2>
  <div>
    <section class="description">
      {@html details.description}
      {#if $isLoggedIn === null}
        <Spinner></Spinner>
      {:else if $isLoggedIn}
        <button on:click={() => addAdventure(details)}>
          Start this adventure for {details.price === 0 ? 'free' : '??' + details.price}</button>
      {:else}
        <p>You must <a href="/login">login</a> before you can start an adventure.</p>
      {/if}
    </section>
    <section class="sidebar">
      {#each details.displayFeatures as feature}
        <h4>
          <svelte:component this={feature[1]}></svelte:component>
          {feature[0]}
        </h4>
        <p>{feature[2]}</p>
      {/each}
    </section>
  </div>
{:catch error}
  <Error error={error}></Error>
{/await}