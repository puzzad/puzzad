<script>
  import Logo from '$components/Logo.svelte'
</script>

<style lang="scss">
  @use '../style/colours';
  @use '../style/fonts';

  a {
    display: flex;
    align-items: end;
  }

  span {
    flex-shrink: 0;
    display: flex;
  }

  #text {
    font-family: fonts.$header;
    font-size: 80px;
    line-height: 70px;
    font-weight: 800;
    padding-left: 10px;
    color: colours.$text;

    @media (max-width: 800px) {
      font-size: 60px;
      line-height: 50px;
    }
  }
</style>

<h1>
  <a href="/">
    <span id="logo"><Logo/></span>
    <span id="text">PUZZAD</span>
  </a>
</h1>