import{S as I,i as x,s as A,e as n,a as v,w as t,c as z,y as a,R as k,C,D as P,n as H,f as w,T,j as G,G as L}from"./index.85c9ce2f.js";function $(c){let l,p,e,i,m,r,y,q,u,o,j,h,_,d,g,b;return{c(){l=n("h2"),l.textContent="Puzzad: Puzzle Adventures",p=v(),e=n("main"),i=n("section"),i.innerHTML=`<h3 class="svelte-1fqfvjc">Ready to start an adventure?</h3> 
        <a href="/#/Adventures" class="svelte-1fqfvjc">Let&#39;s go \xBB</a>`,m=v(),r=n("section"),y=n("h3"),y.textContent="Got a game code?",q=v(),u=n("form"),o=n("input"),j=v(),h=n("input"),_=v(),d=n("section"),d.innerHTML=`<h3 class="svelte-1fqfvjc">What is a puzzle adventure?</h3> 
        <p>A puzzle adventure is part escape room, part puzzle hunt.
            Each adventure is a story interleaved with a variety of puzzles.
            To progress from one part of the story to the next you have to
            solve the puzzle you&#39;re presented with \u2014 but don&#39;t worry,
            there are ample hints available if you get stuck!</p> 
        <h3 class="svelte-1fqfvjc">Is there a time limit?</h3> 
        <p>Nope! Puzzle adventures are completely self-driven. If you want
            to stop and carry on later that&#39;s not a problem. Want to treat
            it like an advent calendar and open one door a day? Go ahead!</p> 
        <p>For those who are more competitive we&#39;ll show your total solve
            time when you finish the adventure (along with number of hints
            used) so you can compare against your friends.</p> 
        <h3 class="svelte-1fqfvjc">Can I play with friends?</h3> 
        <p>Yes, and we actively encourage it! When you start an adventure
            you&#39;ll get a code that looks something like <code>revolving-magenta-ocelot</code>.
            Simply give that code to your friends, they put it in the box
            above and you can all play together.</p> 
        <p>If you&#39;re not physically present, we recommend jumping on a voice
            chat with your team (take your pick from Discord, Zoom, Teams,
            and so on). You can see your teammate&#39;s guesses in real time,
            but it helps if you can ask them <em>why</em> they keep guessing
            &quot;elephant&quot;.</p> 
        <h3 class="svelte-1fqfvjc">How difficult is it?</h3> 
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
            to rope some less puzzle-inclined friends into the hobby!</p>`,t(i,"id","hero"),t(i,"class","svelte-1fqfvjc"),t(y,"class","svelte-1fqfvjc"),t(o,"type","text"),t(o,"placeholder","revolving-magenta-ocelot"),t(o,"class","svelte-1fqfvjc"),t(h,"type","submit"),h.value="Play!",t(h,"class","svelte-1fqfvjc"),t(u,"class","svelte-1fqfvjc"),t(r,"id","play"),t(r,"class","svelte-1fqfvjc"),t(d,"id","intro"),t(d,"class","svelte-1fqfvjc"),t(e,"class","svelte-1fqfvjc")},m(s,f){z(s,l,f),z(s,p,f),z(s,e,f),a(e,i),a(e,m),a(e,r),a(r,y),a(r,q),a(r,u),a(u,o),k(o,c[0]),a(u,j),a(u,h),a(e,_),a(e,d),g||(b=[C(o,"input",c[2]),C(u,"submit",P(c[1]))],g=!0)},p(s,[f]){f&1&&o.value!==s[0]&&k(o,s[0])},i:H,o:H,d(s){s&&w(l),s&&w(p),s&&w(e),g=!1,T(b)}}}function S(c,l,p){let e="";const i=()=>{L("/game/"+e.toLowerCase().replaceAll(/[^a-z-]/g,""))};G.set("Puzzad");function m(){e=this.value,p(0,e)}return[e,i,m]}class Y extends I{constructor(l){super(),x(this,l,S,$,A,{})}}export{Y as default};
