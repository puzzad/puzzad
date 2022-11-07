import{S as I,i as P,s as A,e as i,a as m,b as e,d as z,f as o,a9 as k,l as C,p as H,n as x,g as w,r as L,h as T}from"./index.b92c290c.js";import{t as G}from"./title.a27f956a.js";import{g as S}from"./index.404ba498.js";function W(c){let r,p,t,s,y,n,d,q,u,l,b,h,j,f,g,_;return{c(){r=i("h2"),r.textContent="Puzzad: Puzzle Adventures",p=m(),t=i("main"),s=i("section"),s.innerHTML=`<h3 class="svelte-1fqfvjc">Ready to start an adventure?</h3> 
        <a href="/adventures" class="svelte-1fqfvjc">Let&#39;s go \xBB</a>`,y=m(),n=i("section"),d=i("h3"),d.textContent="Got a game code?",q=m(),u=i("form"),l=i("input"),b=m(),h=i("input"),j=m(),f=i("section"),f.innerHTML=`<h3 class="svelte-1fqfvjc">What is a puzzle adventure?</h3> 
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
            to rope some less puzzle-inclined friends into the hobby!</p>`,e(s,"id","hero"),e(s,"class","svelte-1fqfvjc"),e(d,"class","svelte-1fqfvjc"),e(l,"type","text"),e(l,"placeholder","revolving-magenta-ocelot"),e(l,"class","svelte-1fqfvjc"),e(h,"type","submit"),h.value="Play!",e(h,"class","svelte-1fqfvjc"),e(u,"class","svelte-1fqfvjc"),e(n,"id","play"),e(n,"class","svelte-1fqfvjc"),e(f,"id","intro"),e(f,"class","svelte-1fqfvjc"),e(t,"class","svelte-1fqfvjc")},m(a,v){z(a,r,v),z(a,p,v),z(a,t,v),o(t,s),o(t,y),o(t,n),o(n,d),o(n,q),o(n,u),o(u,l),k(l,c[0]),o(u,b),o(u,h),o(t,j),o(t,f),g||(_=[C(l,"input",c[2]),C(u,"submit",H(c[1]))],g=!0)},p(a,[v]){v&1&&l.value!==a[0]&&k(l,a[0])},i:x,o:x,d(a){a&&w(r),a&&w(p),a&&w(t),g=!1,L(_)}}}function Y(c,r,p){let t;T(c,S,d=>p(3,t=d));let s="";const y=()=>{t("/games/[code]",{code:s.toLowerCase().replaceAll(/[^a-z-]/g,"")})};G.set("Puzzad");function n(){s=this.value,p(0,s)}return[s,y,n]}class E extends I{constructor(r){super(),P(this,r,Y,W,A,{})}}export{E as default};
