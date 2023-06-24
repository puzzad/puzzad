<script>
  export let type = 'success';
  export let title = '';
  export let message = '';
</script>

<style lang="scss">
  @use "../style/colours";

  div {
    border: 1px solid colours.$border;
    width: 350px;
    color: colours.$text-inverted;
    box-shadow: 5px 5px 5px rgba(0, 0, 0, 0.5);

    h5 {
      margin: var(--small-space) var(--small-space) 0 var(--small-space);
      padding: 0;
      border-bottom: none;
    }

    p {
      padding: 10px var(--small-space) var(--small-space) var(--small-space);
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
  <h5>{title}</h5>
  <p>{message}</p>
</div>