<script lang="ts">
  export let options = []

  let choice = options[Math.floor(Math.random() * options.length)] || ''
</script>

{choice}