<script>
  import coffee from '$assets/examples/coffee.png'
  import trains from '$assets/examples/trains.png'
  import buildings from '$assets/examples/buildings.png'
  import charts from '$assets/examples/charts.png'
  import robot from '$assets/examples/robot.png'
  import ExampleQuestion from '$components/ExampleQuestion.svelte'
  import {title} from '$lib/title'

  title.set('Puzzad: Common codes and techniques used in puzzles')
</script>

<style lang="scss">
  @use "../../../style/colours";

  section {
    display: grid;
    grid-template-columns: 40% 1fr var(--small-space);
    grid-column-gap: var(--large-space);

    @media (max-width: 1000px) {
      display: flex;
      flex-direction: column;
      gap: var(--large-space);
    }

    & > h3 {
      grid-column: 1 / span 3;
    }
  }

  .braille {
    font-size: xx-large;
    line-height: 1em;
  }

  table {
    width: 100%;

    td {
      text-align: left;
    }
  }

  blockquote {
    white-space: pre-line;
    color: colours.$brand;
    font-size: xx-large;
  }
</style>

<h2>Common codes and techniques used in puzzles</h2>

<p>
  There are some codes that you will see come up again and again in Escape Rooms and Puzzle Hunts. They're all fairly
  simple to use, but can be presented in ways that can make them difficult to spot. In Escape Rooms you'll nearly
  always be given some kind of reference or hint towards the code, but in most Puzzle Hunts there is more emphasis
  on working it out yourself.
</p>

<p>
  There are plenty of tools and references online to help you with these codes. We've not linked to any here, but
  you should be able to find them easily using your favourite search engine.
</p>

<section>
  <h3>Letters to numbers</h3>

  <div>
    <p>
      Let's start with the basics! If we say that <code>A</code> equals <code>1</code>, <code>B</code> equals
      <code>2</code> and so on we can suddenly write text as numbers. So for example, <code>hello</code> becomes
      <code>8 5 12 12 15</code>. If you ever find yourself with numbers that are all between 1 and 26, you should
      try converting them to letters.
    </p>

    <p>
      This is often called an <strong>a1z26 cipher</strong>. It's very common as a final step in puzzle hunts, where
      the main part of the puzzle produces numbers, and you then have to convert them to letters. Sometimes deciding
      <em>which</em> numbers to convert is a puzzle in itself, though!
    </p>
  </div>

  <ExampleQuestion answer="decaf">
    <p>
      You asked Bob what kind of coffee he wanted, and he sent you a bunch of emojis instead. Can you figure out what
      he wants?
    </p>
    <figure>
      <img src="{coffee}"
           alt="A series of coffee cup emojis: a line of 4, then a line of 5, then 3, then 1 and finally 6.">
    </figure>
  </ExampleQuestion>
</section>

<section>
  <h3>Morse code</h3>

  <div>
    <p>
      Morse code is a way of converting letters into a series of short and long tones, commonly called "dits" and
      "dahs". In text form, short tones are represented by <code>.</code> and long tones by <code>-</code>. For example
      the letter <code>S</code> in morse is four dits: <code>....</code> while an <code>O</code> is three dahs:
      <code>---</code>, so to send out an SOS signal you combine these together: <code>.... --- ....</code>. Fun fact:
      the default notification on old Nokia phones when you received a text message was the morse code for "SMS":
      <code>.... -- ....</code>!
    </p>

    <p>
      While morse code has a distinctive sound, it's quite easy for puzzles to hide morse code visually. Dahs are three
      times the length of dits, so if you find yourself with a set of groups that are either 1 or 3 long, Morse code is
      a good guess. It also naturally appears when you have items that can be categorised as long or short.
    </p>
  </div>

  <ExampleQuestion answer="rain">
    <p>
      Next you ask Bob what his favourite kind of weather is. Instead of answering he makes you watch a series
      of trains go past. Can you figure out what weather he likes?
    </p>
    <figure>
      <img src="{trains}"
           alt="A series of trains: a short one, a long one, then another short; a short one and a long one; two short ones; a long one and a short one">
    </figure>
  </ExampleQuestion>
</section>

<section>
  <h3>Braille</h3>

  <div>
    <p>
      Braille is a system for writing text using raised bumps that can be read tactilely instead of by sight. These
      bumps are arranged into grids consisting of two columns of three rows (although very occasionally a fourth row
      is used) &mdash; seeing something in that arrangement is good hint to think about Braille! For example,
      <code>Y</code> in Braille is <span class="braille">⠽</span> &mdash; you can see it uses five of the six bumps.
    </p>

    <p>
      One interesting aspect of Braille is that it uses the same set of bumps for letters and digits: <code>A</code>
      doubles as <code>1</code>, <code>B</code> as <code>2</code> through to <code>J</code> as <code>0</code>. Sometimes
      the special formatting mark <span class="braille">⠼</span> is used to indicate a number follows, but other times
      you'll just need to figure that out yourself.
    </p>
  </div>

  <ExampleQuestion answer="bumpy">
    <p>
      Still persevering, you ask Bob how his train journey was. He answers by running into a block of flats and
      closing and opening windows. What could he mean?
    </p>
    <figure>
      <img src="{buildings}"
           alt="A series of buildings each with a 2x3 grid of windows, some of which are open">
    </figure>
  </ExampleQuestion>
</section>

<section>
  <h3>Semaphore</h3>

  <div>
    <p>
      Flag Semaphore is a system for communicating over long distances by holding two flags in certain positions.
      It was widely used for ship-to-ship communication prior to the widespread adoption of radios. Semaphore
      flags are normally red and yellow, but when found in puzzles they will rarely be so obvious &mdash; there
      probably won't be any flags at all! Instead, keep an eye out for anything that can be interpreted as pairs
      of angles.
    </p>

    <p>
      Like with Braille, the first nine letters can also be interpreted as numbers. There's a slight difference,
      though: Flag Semaphore uses <code>K</code> to represent <code>0</code>, and instead <code>J</code> is used
      to signal a return to letters. Similarly, the easiest way to decode Semaphore tends to be by hand.
    </p>
  </div>

  <ExampleQuestion answer="ship">
    <p>
      "If you don't like trains, how would you rather get around?" you gamely ask Bob. As usual, he doesn't answer
      but instead shows you a powerpoint presentation containing a series of pie charts.
    </p>
    <figure>
      <img src="{charts}"
           alt="A set of four pie charts, each with one solid area and one unfilled area">
    </figure>
  </ExampleQuestion>
</section>

<section>
  <h3>Pigpen</h3>

  <div>
    <p>
      Pigpen is a system for writing text that uses two grids split into 9 areas, and two diamonds split into 4.
      The letters are inserted alphabetically, so an <code>A</code> is in the top-left cell of the first grid, which
      is represented simply by the internal borders of the grid: <code>┛</code>; <code>E</code> sits in the middle so
      is surrounded on all sides: <code>□</code>; and <code>I</code> is in the bottom-right so is the opposite of
      <code>A</code>: <code>┏</code>. The next nine letters go into the second grid, and are distinguished with a dot
      in the center of the cell. The final eight go clockwise in the two diamonds: first without a dot, then with a dot.
    </p>
    <p>
      Pigpen is common in Escape Rooms, and if you remember the placement order you can decode it fairly easily before
      you find the key. It's also occasionally used in Puzzle Hunts as it can be hidden within an image in interesting
      ways.
    </p>
  </div>

  <ExampleQuestion answer="oink">
    <p>
      Your thoughts are interrupted by a strange sound. You ask Bob what it was, and he produces a note. It looks...
      like a robot? What does he mean?
    </p>
    <figure>
      <img src="{robot}"
           alt="A line drawing of a robot with an unusual set of square facial features">
    </figure>
  </ExampleQuestion>
</section>

<section>
  <h3>Only using certain letters</h3>

  <div>
    <p>
      Have a look at this sentence: <code>all new solvers want easy results</code>. There's an answer hidden in plain
      sight: the first letter of each word spells out <code>ANSWER</code>! Another variation on this theme is taking
      only letters that are capitalised, or bolded, or have some other indicator.
    </p>
    <p>
      A more advanced version of this that is used extremely often in Puzzle Hunts is called <strong>indexing</strong>.
      Indexing is where you have some text and a number, and you use the number to decide which character to take. For
      example if you had <code>3 queens, 1 groom</code> you could take the 3rd letter of "queens" and the 1st letter
      of "groom" to get <code>EG</code>. When indexing, you generally skip over spaces and other punctuation as though
      they're not there.
    </p>
    <p>
      If you're having to index into things, the easiest way to handle it is to put your data into a spreadsheet and
      use the <code>MID</code> function to select the letter for you. That way if you realise you need to change the
      text or the number (or both) the result will automatically update.
    </p>
  </div>
  <ExampleQuestion answer="excel">
    <p>
      You confront Bob about his puzzling responses. "How am I meant to decode these?!" you cry. Bob silently handles
      you a slip of paper in reply. It seems to be how he rates certain TV shows? What does he mean by that?
    </p>
    <table>
      <thead>
      <tr>
        <th>TV Show</th>
        <th>Bob's rating</th>
      </tr>
      </thead>
      <tbody>
      <tr>
        <td>The X-Files</td>
        <td>⭐⭐⭐⭐⭐⭐⭐⭐</td>
      </tr>
      <tr>
        <td>Xena: Warrior Princess</td>
        <td>⭐</td>
      </tr>
      <tr>
        <td>Spaced</td>
        <td>⭐⭐⭐⭐</td>
      </tr>
      <tr>
        <td>Elementary</td>
        <td>⭐⭐⭐</td>
      </tr>
      <tr>
        <td>The Blacklist</td>
        <td>⭐⭐⭐⭐⭐⭐⭐⭐⭐</td>
      </tr>
      </tbody>
    </table>
  </ExampleQuestion>
</section>

<section>
  <h3>Airport codes and other properties</h3>
  <div>
    <p>
      Lots of things in the real world have 'codes' associated with them. Airports have IATA codes: Heathrow
      airport in London is <code>LHR</code>, for example. Countries have ISO-3166 codes: the United Kingdom is
      <code>GB</code>. There are codes for currencies (<code>GBP</code> for the British Pound),
      UK train stations (<code>KGX</code> for Kings Cross), health problems (<code>W59.22XA</code> for "Struck by
      turtle, initial encounter"), and many more. If you look up information about anything during a puzzle keep
      an eye out for any standardised code that might be associated with it, and if your puzzle has some odd letters
      or numbers you don't recognise, try searching for them.
    </p>
    <p>
      It's not just codes, though! Think about an episode of a TV show: it will have a writer, a director, a title,
      a season number, and an episode number. Sometimes the thing you need to discover will be hinted (but maybe not
      in an obvious manner); in harder puzzles you might just need to search through a number of properties
      before something stands out.
    </p>
  </div>
  <ExampleQuestion answer="inland">
    <p>
      You're feeling a bit restless, and decide to take a trip. You ask Bob for a destination suggestion, and he &mdash;
      of course &mdash; hands you a slip of paper. On it is a list of places. Surely there's a hidden meaning, though?
    </p>
    <blockquote>
      Indiana
      Louisiana
      North Dakota
    </blockquote>
  </ExampleQuestion>
</section>

<section>
  <h3>Alternative number formats</h3>
  <div>
    <p>
      When we deal with numbers, we usually use base 10, e.g. <code>42</code>. We can represent the same number in
      other bases easily, though: <code>101010</code> in binary (base 2), <code>222</code> in base 4,
      <code>2A</code> in hex (base 16), and so on. The smaller bases are often used to hide information: if you
      have things that have two different states you can encode them to binary; if you have groups of three you might
      be looking at ternary (base 3), etc.
    </p>
    <p>
      Changing the base isn't the only way to obfuscate numbers, though. Roman numerals are an Escape Room favourite:
      <code>XLII</code>. Then there are tally marks: <code>𝍸𝍸𝍸𝍸𝍸𝍸𝍸𝍸𝍷𝍷</code> in the Western style, or
      <code>𝍬𝍢</code> in rod numerals, among many others. There are also numerous historical systems for writing
      numbers, such as Egyptian Hieroglyphs, Cistercian Numerals and so on.
    </p>
  </div>
  <ExampleQuestion answer="7">
    <p>
      As you and Bob travel, he offers to entertain you with a magic trick. "I'm thinking of a number between 1 and 10",
      he says, "What is it?". He hands you another slip of paper, saying "I'll give you a hint: it's not any of these!"
    </p>
    <blockquote>
      𝍧 I 𝍸 A
      6 00000010 IV
      IX ###
    </blockquote>
  </ExampleQuestion>
</section>

<section>
  <h3>ASCII codes</h3>
  <div>
    <p>
      ASCII is a standard for how computers represent text as numbers. It's a bit more complicated than the A1Z26
      encoding, but not by much. Standard ASCII encoding ranges from <code>0</code> to <code>127</code>. The bits
      we are mostly interested in are: <code>32</code> (space), <code>48-57</code> (digits 0-9),
      <code>65-90</code> (A-Z uppercase), and <code>97-122</code> (a-z in lowercase). If you see numbers in those
      ranges, think ASCII!
    </p>
    <p>
      One interesting property of ASCII is that if you express it in binary, the first bit will always be 0 (because
      ASCII doesn't use numbers above 127). That means if you have a long string of binary you can quickly check the
      1st, 9th, 17th, etc, bits and if they're all 0 you know it's potentially ASCII. Another interesting property
      is that lowercase and uppercase characters have the same value except for the 3rd bit (which is <code>1</code>
      for lowercase and <code>0</code> for uppercase).
    </p>
    <p>
      There are other computer text encodings such as Unicode, but ASCII is far more commonly used in Puzzle Hunts.
    </p>
  </div>
  <ExampleQuestion answer="hello">
    <p>
      You finally reach your destination, and see a welcoming committee. They introduce themselves, but you sense
      a hidden message...
    </p>
    <blockquote style="white-space: normal">
      "I'm Gerald, the 72nd Mayor of Puzzleville. This is my wife, Alice &mdash; she once stayed up for 101 hours to
      finish a Hunt! If you need anything, we live at 76 Puzzle Avenue. That's 76. You can't miss it: there are chimneys
      on the roof that look like the number 111!"
    </blockquote>
  </ExampleQuestion>
</section>