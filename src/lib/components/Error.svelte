<script lang="ts">
    import * as Sentry from "@sentry/svelte";

    export let error;

    const reload = () => document.location.reload()

    $: if (error) {
        Sentry.captureException(error);
    }
</script>

<style>
  section {
    border: 10px solid #8b0000;
    padding: 20px;
    border-radius: 10px;
    color: white;
    background: repeating-linear-gradient(-45deg, #330000 0, #330000 20px, #550000 20px, #550000 40px);
  }

  h2 {
    color: white;
    margin: 0 0 0.5em 0;
  }
</style>

<section>
  <h2>We've taken an unexpected turn...</h2>
  <p>
    Unfortunately, we've encountered an error. We have filed a report with the
    adventurers' guild who hope to have it fixed shortly.
  </p>
  <p>
    In the mean time, please try <a href="#" on:click|preventDefault={reload}>reloading the page</a> and see
    if you can sneak past. Sorry for any inconvenience!
  </p>
</section>