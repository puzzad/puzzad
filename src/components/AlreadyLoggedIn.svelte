<script lang="ts">
  import RandomText from '$components/RandomText.svelte'

  const jokes = [
    'A party of adventurers walk into a tavern, fully armed and sit down at a table.\n' +
    'The bartender comes over and asks, "Hey, why do you guys have your weapons ready?"\n' +
    'The party leader replies, "Mimics."\n' +
    'The bartender laughs.\n' +
    'The party laughs.\n' +
    'The table laughs.',
    'I used to be an adventurer until I went to Antarctica...\n\n' +
    'That\'s when things went south.',
    'My son went on a camping adventure...\n\n' +
    'He told me it was in-tents.',
    'A dragon sees two knights approaching and sighs...\n\n' +
    '"Canned food, again!"',
    'Why are dragons amazing musicians?\n\n' +
    'They know their scales.',
  ]
</script>

<style>
  blockquote {
    white-space: pre-line;
    font-size: xx-large;
    font-style: italic;
    border: 0;
    background-color: #222222;
    border-radius: 15px;
    padding: 1em;
    margin: 0;
  }
</style>

<h2>You're logged in</h2>
<p>There's not a lot we can show you here if you're already logged in. How about a joke?</p>
<blockquote>
  <RandomText options={jokes}></RandomText>
</blockquote>