import{S as q,i as A,s as H,e as i,a as v,b as e,d as z,f as o,a9 as x,l as I,p as L,n as P,g as w,r as T,h as G}from"./index.f7bc07cf.js";import{t as S}from"./title.a72c6b81.js";import{g as W}from"./index.4d5a1717.js";function Y(p){let r,c,t,s,f,n,d,b,u,l,_,h,k,y,g,C;return{c(){r=i("h2"),r.textContent="Puzzad: Puzzle Adventures",c=v(),t=i("main"),s=i("section"),s.innerHTML=`<h3 class="svelte-1o6upsy">Ready to start an adventure?</h3> 
    <a href="/adventures" class="svelte-1o6upsy">Let&#39;s go \xBB</a>`,f=v(),n=i("section"),d=i("h3"),d.textContent="Got a game code?",b=v(),u=i("form"),l=i("input"),_=v(),h=i("input"),k=v(),y=i("section"),y.innerHTML=`<h3 class="svelte-1o6upsy">What is a puzzle adventure?</h3> 
    <p>A puzzle adventure is part escape room, part puzzle hunt.
      Each adventure is a story interleaved with a variety of puzzles.
      To progress from one part of the story to the next you have to
      solve the puzzle you&#39;re presented with \u2014 but don&#39;t worry,
      there are ample hints available if you get stuck!</p> 
    <h3 class="svelte-1o6upsy">Is there a time limit?</h3> 
    <p>Nope! Puzzle adventures are completely self-driven. If you want
      to stop and carry on later that&#39;s not a problem. Want to treat
      it like an advent calendar and open one door a day? Go ahead!</p> 
    <p>For those who are more competitive we&#39;ll show your total solve
      time when you finish the adventure (along with number of hints
      used) so you can compare against your friends.</p> 
    <h3 class="svelte-1o6upsy">Can I play with friends?</h3> 
    <p>Yes, and we actively encourage it! When you start an adventure
      you&#39;ll get a code that looks something like <code>revolving-magenta-ocelot</code>.
      Simply give that code to your friends, they put it in the box
      above and you can all play together.</p> 
    <p>If you&#39;re not physically present, we recommend jumping on a voice
      chat with your team (take your pick from Discord, Zoom, Teams,
      and so on). You can see your teammate&#39;s guesses in real time,
      but it helps if you can ask them <em>why</em> they keep guessing
      &quot;elephant&quot;.</p> 
    <h3 class="svelte-1o6upsy">How difficult is it?</h3> 
    <p>Our starter adventure, Conspiracy, is designed to be around the
      difficulty of a medium escape room. Having some knowledge of
      typical escape rooms or puzzle hunts challenges will help, but
      there is an extensive hint system that can point you in the right
      direction if you&#39;re struggling. You can even skip to the end and
      just get the answer, if you&#39;re really fed up with one of the
      puzzles. It&#39;s your adventure!</p> 
    <p>If you&#39;re a dedicated puzzler who takes parts in puzzle
      hunts or thrives on the more difficult escape rooms, we recommend
      avoiding all hints and seeing how quickly you can complete the
      full adventure. Alternatively, you could use it as a gentle introduction
      to rope some less puzzle-inclined friends into the hobby!</p>`,e(s,"id","hero"),e(s,"class","svelte-1o6upsy"),e(d,"class","svelte-1o6upsy"),e(l,"type","text"),e(l,"placeholder","revolving-magenta-ocelot"),e(l,"class","svelte-1o6upsy"),e(h,"type","submit"),h.value="Play!",e(h,"class","svelte-1o6upsy"),e(u,"class","svelte-1o6upsy"),e(n,"id","play"),e(n,"class","svelte-1o6upsy"),e(y,"id","intro"),e(y,"class","svelte-1o6upsy"),e(t,"class","svelte-1o6upsy")},m(a,m){z(a,r,m),z(a,c,m),z(a,t,m),o(t,s),o(t,f),o(t,n),o(n,d),o(n,b),o(n,u),o(u,l),x(l,p[0]),o(u,_),o(u,h),o(t,k),o(t,y),g||(C=[I(l,"input",p[2]),I(u,"submit",L(p[1]))],g=!0)},p(a,[m]){m&1&&l.value!==a[0]&&x(l,a[0])},i:P,o:P,d(a){a&&w(r),a&&w(c),a&&w(t),g=!1,T(C)}}}function j(p,r,c){let t;G(p,W,d=>c(3,t=d));let s="";const f=()=>{t("/games/[code]",{code:s.toLowerCase().replaceAll(/[^a-z-]/g,"")})};S.set("Puzzad");function n(){s=this.value,c(0,s)}return[s,f,n]}class E extends q{constructor(r){super(),A(this,r,j,Y,H,{})}}export{E as default};
