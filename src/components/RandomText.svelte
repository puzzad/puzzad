<script lang="ts">
    export let options: string[] = []

    let choice = options[Math.floor(Math.random() * options.length)] || ''
</script>

{choice}