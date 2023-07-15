<style lang="scss">
  @use "../style/colours";
  @use "../style/fonts";

  footer {
    margin: var(--small-space) 0 0 0;
    background-color: colours.$border;
    border-top: 2px solid colours.$brand;
    padding: var(--small-space) var(--large-space);
    color: colours.$text-secondary;

    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: space-around;

    @media (max-width: 1200px) {
      flex-direction: column;
    }

    & > div {
      flex-basis: 40%;
    }
  }

  h3 {
    border: 0;
    font-family: fonts.$header;
    font-weight: 800;
    text-transform: uppercase;
  }

  ul {
    list-style-type: disc;
    padding-left: 40px;
  }
</style>

<footer>
  <div>
    <h3 class="faux-logo">PUZZAD</h3>
    <p>
      Puzzle adventures: self-paced, story-driven puzzle hunts.
    </p>
    <p>
      &copy; 2023 Puzzad.com. All rights reserved. <a href="/about">About &amp; Privacy</a>.
    </p>
  </div>
  <div>
    <h3>Learn more</h3>
    <ul>
      <li><a href="/content/what-is-a-puzzle-hunt">What is a puzzle hunt?</a></li>
      <li><a href="/content/common-codes">Common codes and techniques</a></li>
    </ul>
  </div>
</footer>