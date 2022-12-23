<style>
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
    /* sepia is 39 degrees, our accent colour is 42 degrees, so rotate by 3 */
    filter: grayscale(100%) sepia(100%) hue-rotate(3deg) contrast(40%) saturate(400%);
  }
</style>
<section>
  <div>🧩</div>
</section>