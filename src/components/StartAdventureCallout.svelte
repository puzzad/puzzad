<style>
  a {
    display: flex;
    flex-direction: column;
    align-items: end;
    justify-content: space-between;
    position: relative;
    border-radius: 15px;
    margin: 0;
    padding: 1em;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  a::before {
    z-index: -1;
    content: '';
    background-image: url('../assets/hero.jpg');
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    transition: transform 500ms ease-in-out;
  }

  a:hover::before {
    transform: scale(1.1);
  }

  a:hover {
    text-decoration: none;
  }

  h3 {
    margin: 0;
    left: 30%;
    right: 0.5em;
    top: 0.5em;
    font-size: xx-large;
    color: #000000;
    font-weight: bold;
    text-align: right;
    text-shadow: 0 0 5px #ffffff;
  }

  p {
    bottom: 1em;
    right: 1em;
    font-size: large;
    font-weight: bold;
    text-shadow: 0 0 5px #000000;
    margin: 0;
  }

  a:hover p {
    animation: throb 1.5s infinite 500ms;
  }

  @keyframes throb {
    from {
      transform: scale(1);
    }
    50% {
      transform: scale(1.1);
    }
    to {
      transform: scale(1);
    }
  }
</style>

<a href="/adventures">
  <h3>Ready to start an adventure?</h3>
  <p>Let's go &raquo;</p>
</a>