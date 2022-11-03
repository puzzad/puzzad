import{S as x,i as q,s as L,l as n,p as v,r as t,a as z,w as s,a1 as A,B as H,a2 as P,n as I,d as k,Y as T,A as Y}from"./index.58a3881b.js";function $(p){let l,h,e,u,y,i,f,w,r,o,_,c,b,d,g,C;return{c(){l=n("h2"),l.textContent="Puzzad: Puzzle Adventures",h=v(),e=n("main"),u=n("section"),u.innerHTML=`<h3 class="svelte-1eu2ksu">Ready to start an adventure?</h3> 
        <a href="/#/Adventures" class="svelte-1eu2ksu">Let&#39;s go \xBB</a>`,y=v(),i=n("section"),f=n("h3"),f.textContent="Got a game code?",w=v(),r=n("form"),o=n("input"),_=v(),c=n("input"),b=v(),d=n("section"),d.innerHTML=`<h3 class="svelte-1eu2ksu">What is a puzzle adventure?</h3> 
        <p>A puzzle adventure is part escape room, part puzzle hunt.
            Each adventure is a story interleaved with a variety of puzzles.
            To progress from one part of the story to the next you have to
            solve the puzzle you&#39;re presented with \u2014 but don&#39;t worry,
            there are ample hints available if you get stuck!</p> 
        <h3 class="svelte-1eu2ksu">Is there a time limit?</h3> 
        <p>Nope! Puzzle adventures are completely self-driven. If you want
            to stop and carry on later that&#39;s not a problem. Want to treat
            it like an advent calendar and open one door a day? Go ahead!</p> 
        <p>For those who are more competitive we&#39;ll show your total solve
            time when you finish the adventure (along with number of hints
            used) so you can compare against your friends.</p> 
        <h3 class="svelte-1eu2ksu">Can I play with friends?</h3> 
        <p>Yes, and we actively encourage it! When you start an adventure
            you&#39;ll get a code that looks something like <code>revolving-magenta-ocelot</code>.
            Simply give that code to your friends, they put it in the box
            above and you can all play together.</p> 
        <p>If you&#39;re not physically present, we recommend jumping on a voice
            chat with your team (take your pick from Discord, Zoom, Teams,
            and so on). You can see your teammate&#39;s guesses in real time,
            but it helps if you can ask them <em>why</em> they keep guessing
            &quot;elephant&quot;.</p> 
        <h3 class="svelte-1eu2ksu">How difficult is it?</h3> 
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
            to rope some less puzzle-inclined friends into the hobby!</p>`,t(u,"id","hero"),t(u,"class","svelte-1eu2ksu"),t(f,"class","svelte-1eu2ksu"),t(o,"type","text"),t(o,"placeholder","revolving-magenta-ocelot"),t(o,"class","svelte-1eu2ksu"),t(c,"type","submit"),c.value="Play!",t(c,"class","svelte-1eu2ksu"),t(r,"class","svelte-1eu2ksu"),t(i,"id","play"),t(i,"class","svelte-1eu2ksu"),t(d,"id","intro"),t(d,"class","svelte-1eu2ksu"),t(e,"class","svelte-1eu2ksu")},m(a,m){z(a,l,m),z(a,h,m),z(a,e,m),s(e,u),s(e,y),s(e,i),s(i,f),s(i,w),s(i,r),s(r,o),A(o,p[0]),s(r,_),s(r,c),s(e,b),s(e,d),g||(C=[H(o,"input",p[2]),H(r,"submit",P(p[1]))],g=!0)},p(a,[m]){m&1&&o.value!==a[0]&&A(o,a[0])},i:I,o:I,d(a){a&&k(l),a&&k(h),a&&k(e),g=!1,T(C)}}}function G(p,l,h){let e="";const u=()=>{Y("/game/"+e.toLowerCase().replaceAll(/[^a-z-]/g,""))};function y(){e=this.value,h(0,e)}return[e,u,y]}class W extends x{constructor(l){super(),q(this,l,G,$,L,{})}}export{W as default};
