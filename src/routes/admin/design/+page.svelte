<script>
  import {title} from '$lib/title'
  import Logo from '$components/Logo.svelte'
  import Header from '$components/Header.svelte'
  import Nav from '$components/Nav.svelte'
  import Spinner from '$components/Spinner.svelte'
  import Footer from '$components/Footer.svelte'

  title.set('Puzzad: Design System')
</script>

<style lang="scss">
  @use '../../../style/colours';
  @use '../../../style/fonts';

  h4 {
    font-family: monospace;
  }

  .colour {
    width: 80px;
    height: 80px;
    border: 1px solid black;
  }

  .colour-brand {
    background-color: colours.$brand;
  }

  .colour-background {
    background-color: colours.$background;
  }

  .colour-border {
    background-color: colours.$border;
  }

  .colour-text {
    background-color: colours.$text;
  }

  .colour-text-secondary {
    background-color: colours.$text-secondary;
  }

  .colour-text-inverted {
    background-color: colours.$text-inverted;
  }

  .colour-error {
    background-color: colours.$error;
  }

  .colour-success {
    background-color: colours.$success;
  }

  .colour-link {
    background-color: colours.$link;
  }

  .colour-link-active {
    background-color: colours.$link-active;
  }

  .font-specimen, .html-specimen {
    border-left: 15px solid colours.$border;
    padding: 20px;
    margin: 10px 0;
  }

  .font-header {
    font-family: fonts.$header;
    font-size: 70px;
    font-weight: 800;
  }

  .font-primary {
    font-family: fonts.$primary;
    font-size: 21px;
  }
</style>

<section>
  <h2>Design system</h2>


  <h3>Colours</h3>

  <table class="colours">
    <tr>
      <td>
        <div class="colour colour-brand"></div>
      </td>
      <td>
        <h4>colours.$brand</h4>
        <p>Used sparingly as an accent to tie everything together.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-background"></div>
      </td>
      <td>
        <h4>colours.$background</h4>
        <p>Used for the background of pages.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-border"></div>
      </td>
      <td>
        <h4>colours.$border</h4>
        <p>Used for borders and dividers.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-text"></div>
      </td>
      <td>
        <h4>colours.$text</h4>
        <p>Used for primary text.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-text-secondary"></div>
      </td>
      <td>
        <h4>colours.$text-secondary</h4>
        <p>Used for text that's less important than primary text.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-text-inverted"></div>
      </td>
      <td>
        <h4>colours.$text-inverted</h4>
        <p>Used for text that's on a dark background (such as colours.$brand).</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-error"></div>
      </td>
      <td>
        <h4>colours.$error</h4>
        <p>Used to highlight error events.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-success"></div>
      </td>
      <td>
        <h4>colours.$success</h4>
        <p>Used to highlight success events.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-link"></div>
      </td>
      <td>
        <h4>colours.$link</h4>
        <p>Used for links, menus, etc.</p>
      </td>
    </tr>

    <tr>
      <td>
        <div class="colour colour-link-active"></div>
      </td>
      <td>
        <h4>colours.$link-active</h4>
        <p>Used on hover or focus for links.</p>
      </td>
    </tr>
  </table>

  <h3>Typography</h3>

  <h4>fonts.$header</h4>

  <p class="font-specimen font-header">PUZZAD</p>

  <p>Kanit Extra Bold, for use in site header. font-weight 800, size 70px</p>

  <h4>fonts.$primary</h4>

  <p class="font-specimen font-primary">
    “Make your choice, adventurous Stranger,<br>
    Strike the bell and bide the danger,<br>
    Or <strong>wonder</strong>, till it <em>drives</em> you mad,<br>
    What would have <strong><em>followed</em></strong> if you had.”<br>
    ― C.S. Lewis
  </p>

  <p>Kanit, for general copy. Includes variants for bold/italic/bolditalic.</p>

  <h3>HTML/CSS elements</h3>

  <h4>var(--small-space)</h4>

  <section class="html-specimen">
    <div style="background-color: black; width: var(--small-space); height: var(--small-space);"></div>
  </section>

  <h4>var(--large-space)</h4>

  <section class="html-specimen">
    <div style="background-color: black; width: var(--large-space); height: var(--large-space);"></div>
  </section>

  <h4>&lt;hr&gt;</h4>


  <section class="html-specimen">
    <hr>
  </section>

  <h4>&lt;h2&gt;&lt;h3&gt;&lt;h4&gt;&lt;h5&gt;&lt;h6&gt;</h4>

  <p>(h1 is reserved for the site-wide title, and is not styled globally)</p>

  <section class="html-specimen">
    <h2>Header 2</h2>
    <h3>Header 3</h3>
    <h4>Header 4</h4>
    <h5>Header 5</h5>
    <h6>Header 6</h6>
  </section>

  <h4>&lt;ul&gt;</h4>

  <section class="html-specimen">
    <ul>
      <li>List item 1</li>
      <li>List item 2</li>
      <li>
        List item 3
        <ul>
          <li>Nested item i</li>
          <li>Nested item ii</li>
        </ul>
      </li>
    </ul>
  </section>

  <h4>&lt;a&gt;</h4>

  <section class="html-specimen">
    <p><a href="https://www.puzzad.com/">Some link</a></p>
  </section>

  <h4>&lt;input type="text"&gt;</h4>

  <section class="html-specimen">
    <input type="text" placeholder="Some placeholder">
  </section>

  <h4>&lt;input type="submit"&gt;</h4>

  <section class="html-specimen">
    <input type="submit" value="Submit me">
  </section>

  <h3>Components</h3>

  <h4>&lt;Logo/&gt;</h4>

  <Logo/>

  <h4>&lt;Header/&gt;</h4>

  <Header/>

  <h4>&lt;Nav loggedIn={"{true}"} /&gt;</h4>

  <Nav loggedIn={true}/>

  <h4>&lt;Nav loggedIn={"{false}"} /&gt;</h4>

  <Nav loggedIn={false}/>

  <h4>&lt;Spinner/&gt;</h4>

  <Spinner/>

  <h4>&lt;Footer isAdmin={"{false}"}}/&gt;</h4>

  <Footer isAdmin={false}/>

  <h4>&lt;Footer isAdmin={"{true}"}}/&gt;</h4>

  <Footer isAdmin={true}/>
</section>