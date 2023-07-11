<script lang="ts">
  import {goto} from '$app/navigation'
  import AdventureLogo from '$components/AdventureLogo.svelte'
  import {formatDuration} from '$lib/time'
  import Spinner from '$components/Spinner.svelte'
  import {title} from '$lib/title'
  import Error from '$components/Error.svelte'
  import Certificate from '$components/Certificate.svelte'
  import {client, getGameClient} from '$lib/api'
  import GameStats from '$components/GameStats.svelte'
  import SubscribeToMailingList from '$components/SubscribeToMailingList.svelte'

  export let data

  let game = getGameClient(data.game).
  then(client => client.collection("games").getFirstListItem("username='"+data.game+"'", {expand: "adventure"}))
      .then((game) => {
        title.set(`Puzzad: ${game.expand.adventure.name} - ${data.game}`)
        return game
      })

  const handleStartAdventure = async () =>{
    client.send("/wom/startgame",{
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ game: data.game })
    })
    .then(() => goto(`/games/${data.game}/play`))
  }

  const handleContinueAdventure = async function() {
    await goto(`/games/${data.game}/play`)
  }
</script>

<style>
  code {
    display: block;
    background-image: url('../../../assets/code-bg.png');
    background-color: transparent;
    color: black;
    text-align: center;
    background-size: contain;
    background-repeat: no-repeat;
    background-position: center;
    padding: 70px 0;
    font-size: xxx-large;
    font-family: monospace;
    margin: 1em;
  }

  button {
    width: 100%;
    margin: 1em 0;
    font-size: x-large;
    padding: 0.5em;
  }

</style>

{#await game}
  <Spinner/>
{:then gameData}
  <h2>
    <AdventureLogo name={gameData.expand.adventure.name}></AdventureLogo>
  </h2>
  {#if gameData.status === 'EXPIRED'}
    <p>Congratulations! You finished the adventure!</p>
    <Certificate adventureName={gameData.expand.adventure.name} teamName={data.game}
                 completionDate={gameData.end}></Certificate>
    <section class="stats">
      <h3>Adventure statistics</h3>
      <p>You took {formatDuration(gameData.start, gameData.end)}!</p>
      <GameStats code={data.game}></GameStats>
    </section>
    <SubscribeToMailingList></SubscribeToMailingList>
  {:else if gameData.status === 'PAID'}
    <p>
      You've not yet started your adventure! Remember, it's dangerous to go alone.
      Take this if you want to recruit others to help you:
    </p>
    <code>
      {data.game}
    </code>
    <p>
      Once you're ready, press below to go to the first puzzle in the adventure!
    </p>
    <button on:click={() => handleStartAdventure()}>
      Begin the adventure!
    </button>
  {:else if gameData.status === 'ACTIVE'}
    <p>This adventure is already in progress!</p>
    <button on:click={() => handleContinueAdventure()}>
      Go to the current puzzle
    </button>
  {/if}
{:catch error}
  {#if error.message === 'Invalid team code'}
    <h1>Game not found</h1>
    <p>
      The adventure you are seeking does not seem to exist.
      If you have been given a code, please double check your spelling.
      If you registered the adventure yourself, you can find it listed in <a href="/games">My Games</a>.
    </p>
  {:else}
    <Error error={error}></Error>
  {/if}
{/await}