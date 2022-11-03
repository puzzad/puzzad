import{S as I,i as x,s as q,l as n,p as v,r as t,a as b,w as a,a1 as C,B as A,a2 as L,n as H,d as z,Y as P,A as T}from"./index.f6a43aeb.js";function Y(p){let l,h,e,i,y,r,f,k,u,o,w,c,j,d,g,_;return{c(){l=n("h2"),l.textContent="Puzzad: Puzzle Adventures",h=v(),e=n("main"),i=n("section"),i.innerHTML=`<h3 class="svelte-3kmbj0">Ready to start an adventure?</h3> 
        <a href="/#/Adventures" class="svelte-3kmbj0">Let&#39;s go \xBB</a>`,y=v(),r=n("section"),f=n("h3"),f.textContent="Got a game code?",k=v(),u=n("form"),o=n("input"),w=v(),c=n("input"),j=v(),d=n("section"),d.innerHTML=`<h3 class="svelte-3kmbj0">What is a puzzle adventure?</h3> 
        <p>A puzzle adventure is part escape room, part puzzle hunt.
            Each adventure is a story interleaved with a variety of puzzles.
            To progress from one part of the story to the next you have to
            solve the puzzle you&#39;re presented with \u2014 but don&#39;t worry,
            there are ample hints available if you get stuck!</p> 
        <h3 class="svelte-3kmbj0">Is there a time limit?</h3> 
        <p>Nope! Puzzle adventures are completely self-driven. If you want
            to stop and carry on later that&#39;s not a problem. Want to treat
            it like an advent calendar and open one door a day? Go ahead!</p> 
        <p>For those who are more competitive we&#39;ll show your total solve
            time when you finish the adventure (along with number of hints
            used) so you can compare against your friends.</p> 
        <h3 class="svelte-3kmbj0">Can I play with friends?</h3> 
        <p>Yes, and we actively encourage it! When you start an adventure
            you&#39;ll get a code that looks something like <code>revolving-magenta-ocelot</code>.
            Simply give that code to your friends, they put it in the box
            above and you can all play together.</p> 
        <p>If you&#39;re not physically present, we recommend jumping on a voice
            chat with your team (take your pick from Discord, Zoom, Teams,
            and so on). You can see your teammate&#39;s guesses in real time,
            but it helps if you can ask them <em>why</em> they keep guessing
            &quot;elephant&quot;.</p> 
        <h3 class="svelte-3kmbj0">How difficult is it?</h3> 
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
            to rope some less puzzle-inclined friends into the hobby!</p>`,t(i,"id","hero"),t(i,"class","svelte-3kmbj0"),t(f,"class","svelte-3kmbj0"),t(o,"type","text"),t(o,"placeholder","revolving-magenta-ocelot"),t(o,"class","svelte-3kmbj0"),t(c,"type","submit"),c.value="Play!",t(c,"class","svelte-3kmbj0"),t(u,"class","svelte-3kmbj0"),t(r,"id","play"),t(r,"class","svelte-3kmbj0"),t(d,"id","intro"),t(d,"class","svelte-3kmbj0"),t(e,"class","svelte-3kmbj0")},m(s,m){b(s,l,m),b(s,h,m),b(s,e,m),a(e,i),a(e,y),a(e,r),a(r,f),a(r,k),a(r,u),a(u,o),C(o,p[0]),a(u,w),a(u,c),a(e,j),a(e,d),g||(_=[A(o,"input",p[2]),A(u,"submit",L(p[1]))],g=!0)},p(s,[m]){m&1&&o.value!==s[0]&&C(o,s[0])},i:H,o:H,d(s){s&&z(l),s&&z(h),s&&z(e),g=!1,P(_)}}}function $(p,l,h){let e="";const i=()=>{T("/game/"+e.toLowerCase().replaceAll(/[^a-z-]/g,""))};function y(){e=this.value,h(0,e)}return[e,i,y]}class S extends I{constructor(l){super(),x(this,l,$,Y,q,{})}}export{S as default};
