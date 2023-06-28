<script>
  import {title} from '$lib/title'
  import ExampleQuestion from '$components/ExampleQuestion.svelte'
  import gaming from '$assets/examples/gaming.webp'

  title.set('Puzzad: What is a puzzle hunt?')
</script>

<style lang="scss">
  @use "../../../style/colours";

  .hero-quote {
    margin-bottom: var(--small-space);

    @media (min-width: 1200px) {
      float: right;
      margin: 0 0 var(--small-space) var(--small-space);
    }
  }

  .flavour {
    color: colours.$text-secondary;
  }

  .puzzle {
    margin: 15px 0 var(--small-space) 0;
  }

  .example {
    border: 2px solid colours.$border;
    padding: var(--small-space);
    margin: var(--small-space) 0;
    display: flex;
    justify-content: center;
  }

  details {
    border: 1px solid colours.$border;
    margin: var(--small-space) 0;
    padding: 1em;

    &[open] summary {
      margin-bottom: var(--small-space);
    }
  }
</style>

<h2>What is a puzzle hunt?</h2>

<blockquote class="hero-quote">
  "Imagine a word search.
  Now imagine you aren’t told what words to look for.
  Now imagine you aren’t told it’s a word search.
  Now imagine it isn’t a word search."
  <span class="attribution">&mdash; Mike Develin</span>
</blockquote>

<p>
  A puzzle hunt is a competition where people <strong>compete to solve puzzles</strong>. The puzzles can be almost
  anything, but generally <strong>a puzzle will provide an answer</strong> and that answer is usually a
  <strong>single word or a short phrase</strong>. The other defining feature is that there
  <strong>won't be instructions</strong> telling you what to do: figuring that out is part of the puzzle itself!
  Beyond those basic constraints, puzzles come in all sorts of shapes and sizes: one puzzle might have pictures of
  toy unicorns, another might have audio clips from Star Trek episodes, and a third might have a video of a chess match.
</p>

<p>
  When you first see a puzzle in a hunt you almost certainly won't know how to solve it, but you might have some ideas
  for things to investigate. After following some of those lines of investigation, you might notice a pattern starting
  to form. That pattern might lead you to look at the puzzle in a different way, and that might give you some more ideas
  to investigate. At some point you'll likely have an "Aha!" moment where you understand what is going on and what you
  need to do. Then all you need to do is get the answer!
</p>

<h3>Anatomy of a puzzle</h3>

<p>
  In most puzzle hunts, each puzzle will have a <strong>title</strong>, some <strong>flavour text</strong>, and then
  the actual content of the <strong>puzzle itself</strong>. Let's look at an example:
</p>

<aside class="example">
  <ExampleQuestion answer="shone" header={false}>
    <h4>An elementary riddle</h4>
    <p class="flavour">
      You enter the pyramid of the great Pharaoh Vee Lednem. You see a sarcophagus, and on it is an inscription. It
      reads:
    </p>
    <blockquote class="puzzle">
      I’m often made bitter, I’m found by volcanoes.
      The first of many, unique yet plentiful; I fill three quarters of known mass.
      I'm found underwater, and even in the skies; if you don't use me you'll die.
      One of six nobilities, I demonstrate myself in many signs.
    </blockquote>
  </ExampleQuestion>
</aside>

<p>
  Here the title is "An elementary riddle", and the flavour text is the paragraph that starts "You enter the
  pyramid...". The puzzle itself is the larger text in the box. The flavour text is there to relate the puzzle to
  an ongoing story or theme, and may also give some hints as to how the solve the puzzle. The title is usually a
  hint as to what the puzzle is about, but it's generally not too obvious about it.
</p>

<p>
  This was a real puzzle from a hunt in 2023, although we've changed the title and flavour text here. This is a fairly
  straight-forward puzzle &mdash; see if you can solve it! Never fear, though: there's a full walkthrough below if
  you get stuck.
</p>

<details>
  <summary>Walkthrough of "An elementary riddle"</summary>
  <h4>Initial thoughts</h4>
  <p>
    The first thing we notice here is that the title of the puzzle has the word "riddle", and the inscription seems
    to take the form of a typical riddle. Some words might stick out to you, though: why is it an "elementary" riddle,
    and why is the Pharaoh's name "Vee Lednem"? Maybe nothing springs to mind right away: that's fine, let's carry on
    looking at the puzzle.
  </p>
  <p>
    We can see the puzzle text itself is made up of four sentences. Based on the wording, each sentence is
    likely a separate riddle: it's unlikely the same thing will be "found by volcanoes" and also "found underwater" or
    "in the sky"; similarly one is "first of many" while another is "one of six". We can't be totally confident about
    that, but it seems like a reasonable assumption at this stage. If it doesn't work out, we can come back and revisit
    the idea.
  </p>
  <h4>The "Aha!" moment</h4>
  <p>
    So we have four riddles, and presumably need to get four answers. But where do we go from here? You might have
    noticed the word "nobilities" in the fourth riddle. Does that make you think of anything? What about when you think
    about the "elementary" in the title? If you're familiar with the periodic table of elements, you'll know there's
    a group of elements called the "noble gases", and there are six of them. (If you're not at all familiar with the
    periodic table of elements, you may have to use a search engine to figure this out. That's fine!)
  </p>
  <p>
    So we think the riddles might be about chemical elements. <strong>Aha!</strong> "Vee Lednem" is "Mendeleev", the
    creator of the periodic table, backwards! That confirms we're on the right track! (You can obviously solve the
    puzzle entirely without noticing that. Or maybe you noticed it immediately and didn't need any other hints?!)
  </p>
  <h4>Solving</h4>
  <p>
    Now we can solve the riddles, with either our intuition or a search engine. The answers are: Sulphur, Hydrogen,
    Oxygen, and Neon. But now what? That doesn't feel like a satisfying answer. We must need to do something else with
    these answers. What could it be?
  </p>
  <p>
    One thing that all elements on the periodic table have is a <strong>symbol</strong>. For example, the symbol for
    Hydrogen is <code>H</code>. If we look up all four, we get: <code>S H O Ne</code>. That's a word! If you're
    trying to solve a puzzle and a word pops out, that's a very good sign. If we put <code>shone</code> into the
    answer box, it tells us we're correct! Hurrah!
  </p>
</details>

<h3>More complex puzzles</h3>

<p>
  Now let's look at a more complex puzzle! This is also from a 2023 hunt, but came much later than the previous one.
  Generally the puzzles will get more complex as you progress through the hunt, so don't worry if this seems a bit
  daunting.
</p>

<aside class="example">
  <ExampleQuestion answer="magic" header={false}>
    <h4>Gaming break!</h4>
    <p class="flavour">After a long hard day puzzling, you decide to take a break and play some games.</p>
    <img class="puzzle" src="{gaming}"
         alt="Various screenshots of games contained in blocky numbers, each with a different coloured border">
  </ExampleQuestion>
</aside>

<details>
  <summary>Walkthrough of "Gaming break!"</summary>

  <h4>Initial thoughts and research</h4>

  <p>
    As with the previous puzzle, let's gather our initial thoughts: the title and the flavour text reference gaming,
    so we can safely assume that all the images are from video games. We can note that the numbers have
    different coloured borders, and also that the adjacent "1" and "0"s seem to be the same image and colour, so they're
    probably the number "10" rather than two separate numbers.
  </p>

  <p>
    One obvious thing we can start doing is identifying the games. A lot of the time you can use reverse image search
    tools, but the shapes of these pictures makes that quite difficult. Instead, we can try to identify the games by
    the text that we can see. The actual image file is a lot bigger than you see on your screen: if you open it
    in a new browser tab, or download it and view it in an image editor, you'll be able to zoom in on sections.
  </p>

  <p>
    Taking the number 3 as an example, you can see in the bottom right a little section of what looks like character
    information: it has a player name, a level, and then the word <code>Horde</code>. Searching the Internet for
    "Horde" you'll find multiple references from the game <strong>World of Warcraft</strong>. If you then search for
    some screenshots of the game you might recognise the art style or the UI elements, which confirms you have the
    right answer. Repeating this process for each set of images gives us:
  </p>

  <ul>
    <li>Counter-Strike: Global Offensive</li>
    <li>World of Warcraft</li>
    <li>Guild Wars 2</li>
    <li>Lord of the Rings Online</li>
    <li>Hogwarts Legacy</li>
    <li>Elder Scrolls Online</li>
  </ul>

  <h4>Further thoughts and more research</h4>

  <p>
    We have some text and we have some numbers, so we can try to apply one common puzzle technique: indexing.
    This means taking the <em>n</em>th letter of each bit of text, where <em>n</em> is the number associated with it.
    So "Counter-Strike: Global Offensive" has the number 10, so we take the 10th letter (usually ignoring any spaces or
    punctuation): <code>r</code>. If we do that with all the games we get <code>rrurc</code>. Huh. That doesn't look
    good.
  </p>

  <p>
    When you get stuck on a puzzle, it's always helpful to go over which parts of the puzzle you've used so far:
    we've used the screenshots to get the games, and we've tried to use the numbers but didn't get anywhere good...
    What we've not used is the border colours. Maybe we should try to use those somehow?
  </p>

  <p>
    If we do another web search for (for example) <code>Elder Scrolls Online +blue</code> we find a smattering of
    results that mention "blue items". The same is true for a few of the other game and colour combinations. There
    aren't great results for some of them, but it's an interesting lead. If we do some digging into these results we
    find that the games have a system where they colour items or equipment according to rarity, and &mdash;
    <strong>aha!</strong> &mdash; the border colours seem to match up with the rarity colours. Now we're armed with
    better search terms, we can find the corresponding "rarity" levels for each game and colour:
  </p>

  <ul>
    <li>Counter-Strike: Global Offensive: Restricted</li>
    <li>World of Warcraft: Epic</li>
    <li>Guild Wars 2: Masterwork</li>
    <li>Lord of the Rings Online: Incomparable</li>
    <li>Hogwarts Legacy: - Well-Appointed</li>
    <li>Elder Scrolls Online: Superior</li>
  </ul>

  <h4>They think it's all over... It is now!</h4>

  <p>
    It doesn't look like the first letters of these rarities spells anything good (<code>remiws</code>), but we still
    have the numbers to use. If we try indexing into the rarities we get <code>diablo</code>. That's a word! You
    might think you're finished now: that's <em>obviously</em> the answer, right? But if you put it in the answer box,
    you don't get a satisfying green "success" message. What's going on?
  </p>

  <p>
    This is actually a fairly common puzzle trope: you've done <em>something</em> to get an intermediate answer, and
    in order to get the final answer you need to do that same <em>something</em> once more. In this case, we've got
    "Diablo" and some more searching will confirm that is the name of the video game, and that it has item rarities
    like the other games. But which one do we care about?
  </p>

  <p>
    The solution is, once again, to go back and re-check what information we've used: we've used the numbers, the
    screenshots, and the borders around the screenshots. But there's also a border around the whole image! You might
    have ignored it as just a bit of decoration, but if you look at the colour you'll notice it corresponds to one of
    the item rarities in Diablo! <strong>Aha!</strong> That item rarity is called <code>magic</code>, which gives us the
    answer.
  </p>
</details>

<p>
  As you can see, the puzzles can be quite involved, taking multiple steps to get to the final answer. Generally, each
  stage will be in some way self-confirming, though, so you're able to tell you're on the right track. It's also
  important to remember that you'll almost certainly feel "stuck", and not sure how to proceed. On average, it
  took participants in the hunt 1 hour to solve this puzzle, and they were given some extra hints along the way.
  It can be painful when you feel stuck, but it makes that "<strong>aha!</strong>" moment even sweeter when you do
  eventually find your way.
</p>

<h3>Other features of a hunt</h3>

<h4>Structure</h4>

<p>
  We've taken a look at the puzzles, but how do they all go together to form a puzzle hunt? Generally, a hunt consists
  of a number of <strong>episodes</strong> or <strong>rounds</strong>, which each contain some puzzles. There's usually
  also a <strong>meta-puzzle</strong> in each of the groups, which references the other puzzles somehow: it might
  need you to insert the answers from the other puzzles in some way, or you might have to use mechanics from the earlier
  puzzles in a novel manner. Completing the meta usually counts as finishing the entire group, even if you didn't
  complete all the individual puzzles (but skipping puzzles will generally make it harder to solve the meta, so it's a
  tricky balancing act!)
</p>

<p>
  In most hunts, the number of puzzles you can access initially will be limited. In some hunts you'll need to solve
  one puzzle at a time to get to the next one; in others you might have a small selection which will gradually expand
  as you solve more puzzles. The way this works will usually be explained up front in the hunt's introduction &mdash;
  you don't have to figure that out for yourself! Where you have multiple puzzles accessible, your team can split up
  and work on different puzzles at the same time, or you can all work together on the same puzzle and switch if you
  get too stuck. Sometimes taking a break from a puzzle is the best way to progress on it!
</p>

<h4>Help and hinting</h4>

<p>
  Hunts will usually come with some kind of <strong>hinting system</strong>. This usually takes one of two forms:
  either you communicate directly with the hunt admins, and they explain exactly what you're stuck on, or there is an
  automated system that dispenses hints that may or may not help you immediately. Again, this will generally be
  explained before the hunt starts so everyone knows what to expect. In any competitive hunt, the admins will strive
  to ensure hints are given fairly to all teams and don't give away too much too soon, so you might have to be patient
  when waiting for help.
</p>
<p>
  Some hunts also have <strong>unlocks</strong> or <strong>eurekas</strong>, which are automatic
  responses to things you put in the answer box. These will generally confirm that you're on the right track, but try
  not to rely on them too much as they'll often be used more sparingly in harder puzzles.
</p>

<h4>Physicality</h4>

<p>
  Some hunts are <strong>in-person</strong> events: you have to physically be at a location to do some of the puzzles.
  This limits who can take part, but opens up a wide variety of puzzle types that wouldn't be possible otherwise.
  It also enables interaction between teams &mdash; it makes it a much more social experience when you can bump into
  other teams while following clues. In-person hunts usually still have a majority of puzzles that can be solved
  remotely, so it's possible to team up with someone who's on site and work together, with them acting as the "runner"
  to do the physical parts.
</p>

<p>
  Some hunts that are online-only may also included physical elements: check the instructions to see if you'll need
  access to a printer or other resources. Some hunts even have physical elements that they post out to teams in advance,
  so be sure to check the details well in advance of the hunt starting!
</p>
