<script>
  import headerLogo from '$assets/logo.png'
</script>

<style lang="scss">
  @use '../style/colours';

  img {
    filter: colours.$filter-black-to-brand;

    @media (max-width: 800px) {
      height: 50px;
    }
  }
</style>

<img src="{headerLogo}" alt="Logo consisting of a man with climbing gear ascending the side of a puzzle piece">