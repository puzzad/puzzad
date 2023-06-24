<style lang="scss">
  @use '../style/colours';

  section {
    display: flex;
    justify-content: center;
  }

  @keyframes spinner {
    0% {
      transform: translate3d(0%, 0%, 0) rotate(0deg);
    }
    100% {
      transform: translate3d(0%, 0%, 0) rotate(360deg);
    }
  }

  div {
    animation: 1s linear infinite spinner;
    will-change: transform;
    animation-play-state: inherit;
    font-size: 50px;
    padding: 50px 0;
    user-select: none;
    filter:
        #{'grayscale(100%) contrast(1000%) invert()'}
        colours.$filter-black-to-brand;
  }
</style>
<section>
  <div>🧩</div>
</section>