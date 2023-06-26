<script>
  export let type = 'success'
  export let title = ''
  export let message = ''
</script>

<style lang="scss">
  @use "../style/colours";

  div {
    border: 1px solid colours.$border;
    min-width: 250px;
    max-width: max(450px, 25vw);
    color: colours.$text-on-brand;
    box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.5);
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: var(--small-space);

    h5 {
      padding: 0;
      border-bottom: none;
    }

    &.success {
      background-color: colours.$success;
    }

    &.error {
      background-color: colours.$error;
    }
  }
</style>

<div class="toast {type}">
  {#if title}<h5>{title}</h5>{/if}
  <p>{message}</p>
</div>