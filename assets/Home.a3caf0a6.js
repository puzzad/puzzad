import{S as A,i as q,s as L,h as n,j as v,r as t,a as z,w as a,a3 as C,C as H,a4 as P,n as I,d as b,_ as T,B as $}from"./index.50bbf899.js";function j(p){let l,h,e,i,y,r,f,w,u,o,x,c,_,d,g,k;return{c(){l=n("h2"),l.textContent="Puzzad: Puzzle Adventures",h=v(),e=n("main"),i=n("section"),i.innerHTML=`<h3 class="svelte-zxgmb4">Ready to start an adventure?</h3> 
        <a href="/#/Adventures" class="svelte-zxgmb4">Let&#39;s go \xBB</a>`,y=v(),r=n("section"),f=n("h3"),f.textContent="Got a game code?",w=v(),u=n("form"),o=n("input"),x=v(),c=n("input"),_=v(),d=n("section"),d.innerHTML=`<h3 class="svelte-zxgmb4">What is a puzzle adventure?</h3> 
        <p>A puzzle adventure is part escape room, part puzzle hunt.
            Each adventure is a story interleaved with a variety of puzzles.
            To progress from one part of the story to the next you have to
            solve the puzzle you&#39;re presented with \u2014 but don&#39;t worry,
            there are ample hints available if you get stuck!</p> 
        <h3 class="svelte-zxgmb4">Is there a time limit?</h3> 
        <p>Nope! Puzzle adventures are completely self-driven. If you want
            to stop and carry on later that&#39;s not a problem. Want to treat
            it like an advent calendar and open one door a day? Go ahead!</p> 
        <p>For those who are more competitive we&#39;ll show your total solve
            time when you finish the adventure (along with number of hints
            used) so you can compare against your friends.</p> 
        <h3 class="svelte-zxgmb4">Can I play with friends?</h3> 
        <p>Yes, and we actively encourage it! When you start an adventure
            you&#39;ll get a code that looks something like <code>revolving-magenta-ocelot</code>.
            Simply give that code to your friends, they put it in the box
            above and you can all play together.</p> 
        <p>If you&#39;re not physically present, we recommend jumping on a voice
            chat with your team (take your pick from Discord, Zoom, Teams,
            and so on). You can see your teammate&#39;s guesses in real time,
            but it helps if you can ask them <em>why</em> they keep guessing
            &quot;elephant&quot;.</p> 
        <h3 class="svelte-zxgmb4">How difficult is it?</h3> 
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
            to rope some less puzzle-inclined friends into the hobby!</p>`,t(i,"id","hero"),t(i,"class","svelte-zxgmb4"),t(f,"class","svelte-zxgmb4"),t(o,"type","text"),t(o,"placeholder","revolving-magenta-ocelot"),t(o,"class","svelte-zxgmb4"),t(c,"type","submit"),c.value="Play!",t(c,"class","svelte-zxgmb4"),t(u,"class","svelte-zxgmb4"),t(r,"id","play"),t(r,"class","svelte-zxgmb4"),t(d,"id","intro"),t(d,"class","svelte-zxgmb4"),t(e,"class","svelte-zxgmb4")},m(s,m){z(s,l,m),z(s,h,m),z(s,e,m),a(e,i),a(e,y),a(e,r),a(r,f),a(r,w),a(r,u),a(u,o),C(o,p[0]),a(u,x),a(u,c),a(e,_),a(e,d),g||(k=[H(o,"input",p[2]),H(u,"submit",P(p[1]))],g=!0)},p(s,[m]){m&1&&o.value!==s[0]&&C(o,s[0])},i:I,o:I,d(s){s&&b(l),s&&b(h),s&&b(e),g=!1,T(k)}}}function G(p,l,h){let e="";const i=()=>{$("/game/"+e.toLowerCase().replaceAll(/[^a-z-]/g,""))};function y(){e=this.value,h(0,e)}return[e,i,y]}class W extends A{constructor(l){super(),q(this,l,G,j,L,{})}}export{W as default};
