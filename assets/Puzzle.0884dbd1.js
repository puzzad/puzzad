import{S as le,i as ie,s as oe,b as se,c as z,n as M,f as k,F as Me,e as p,w,H as U,m as De,x as v,o as E,a as R,y as S,p as L,I as Oe,J as Ae,t as T,d as P,q as W,A as J,K as _e,B as te,G as ne,E as X,h as He,u as Fe,L as Be,k as qe,l as Ge,j as ge,M as Ee,N as ee,O as Le,r as fe,v as he,P as ze,Q as We,z as ke,R as Ye}from"./index.910c6cb7.js";import{R as Re}from"./RandomText.12dd4b21.js";import{S as Te}from"./Spinner.ac0ad4fa.js";import{t as je,T as Qe,F as Ue}from"./FlatToast.f9a4e95b.js";function be(l,e,n){const t=l.slice();return t[18]=e[n],t}function pe(l){let e,n={length:l[6]},t=[];for(let i=0;i<n.length;i+=1)t[i]=Ce(be(l,n,i));return{c(){e=p("div");for(let i=0;i<t.length;i+=1)t[i].c();w(e,"class","confetti-holder svelte-io58ff"),U(e,"rounded",l[9]),U(e,"cone",l[10]),U(e,"no-gravity",l[11])},m(i,s){z(i,e,s);for(let o=0;o<t.length;o+=1)t[o].m(e,null)},p(i,s){if(s&20991){n={length:i[6]};let o;for(o=0;o<n.length;o+=1){const r=be(i,n,o);t[o]?t[o].p(r,s):(t[o]=Ce(r),t[o].c(),t[o].m(e,null))}for(;o<t.length;o+=1)t[o].d(1);t.length=n.length}s&512&&U(e,"rounded",i[9]),s&1024&&U(e,"cone",i[10]),s&2048&&U(e,"no-gravity",i[11])},d(i){i&&k(e),De(t,i)}}}function Ce(l){let e;return{c(){e=p("div"),w(e,"class","confetti svelte-io58ff"),v(e,"--fall-distance",l[8]),v(e,"--size",l[0]+"px"),v(e,"--color",l[14]()),v(e,"--skew",N(-45,45)+"deg,"+N(-45,45)+"deg"),v(e,"--rotation-xyz",N(-10,10)+", "+N(-10,10)+", "+N(-10,10)),v(e,"--rotation-deg",N(0,360)+"deg"),v(e,"--translate-y-multiplier",N(l[2][0],l[2][1])),v(e,"--translate-x-multiplier",N(l[1][0],l[1][1])),v(e,"--scale",.1*N(2,10)),v(e,"--transition-duration",l[4]?`calc(${l[3]}ms * var(--scale))`:`${l[3]}ms`),v(e,"--transition-delay",N(l[5][0],l[5][1])+"ms"),v(e,"--transition-iteration-count",l[4]?"infinite":l[7]),v(e,"--x-spread",1-l[12])},m(n,t){z(n,e,t)},p(n,t){t&256&&v(e,"--fall-distance",n[8]),t&1&&v(e,"--size",n[0]+"px"),t&4&&v(e,"--translate-y-multiplier",N(n[2][0],n[2][1])),t&2&&v(e,"--translate-x-multiplier",N(n[1][0],n[1][1])),t&24&&v(e,"--transition-duration",n[4]?`calc(${n[3]}ms * var(--scale))`:`${n[3]}ms`),t&32&&v(e,"--transition-delay",N(n[5][0],n[5][1])+"ms"),t&144&&v(e,"--transition-iteration-count",n[4]?"infinite":n[7]),t&4096&&v(e,"--x-spread",1-n[12])},d(n){n&&k(e)}}}function Je(l){let e,n=!l[13]&&pe(l);return{c(){n&&n.c(),e=se()},m(t,i){n&&n.m(t,i),z(t,e,i)},p(t,[i]){t[13]?n&&(n.d(1),n=null):n?n.p(t,i):(n=pe(t),n.c(),n.m(e.parentNode,e))},i:M,o:M,d(t){n&&n.d(t),t&&k(e)}}}function N(l,e){return Math.random()*(e-l)+l}function Ke(l,e,n){let{size:t=10}=e,{x:i=[-.5,.5]}=e,{y:s=[.25,1]}=e,{duration:o=2e3}=e,{infinite:r=!1}=e,{delay:f=[0,50]}=e,{colorRange:d=[0,360]}=e,{colorArray:u=[]}=e,{amount:m=50}=e,{iterationCount:_=1}=e,{fallDistance:C="100px"}=e,{rounded:y=!1}=e,{cone:b=!1}=e,{noGravity:I=!1}=e,{xSpread:D=.15}=e,{destroyOnComplete:A=!0}=e,G=!1;Me(()=>{!A||r||_=="infinite"||setTimeout(()=>n(13,G=!0),(o+f[1])*_)});function j(){return u.length?u[Math.round(Math.random()*(u.length-1))]:`hsl(${Math.round(N(d[0],d[1]))}, 75%, 50%`}return l.$$set=c=>{"size"in c&&n(0,t=c.size),"x"in c&&n(1,i=c.x),"y"in c&&n(2,s=c.y),"duration"in c&&n(3,o=c.duration),"infinite"in c&&n(4,r=c.infinite),"delay"in c&&n(5,f=c.delay),"colorRange"in c&&n(15,d=c.colorRange),"colorArray"in c&&n(16,u=c.colorArray),"amount"in c&&n(6,m=c.amount),"iterationCount"in c&&n(7,_=c.iterationCount),"fallDistance"in c&&n(8,C=c.fallDistance),"rounded"in c&&n(9,y=c.rounded),"cone"in c&&n(10,b=c.cone),"noGravity"in c&&n(11,I=c.noGravity),"xSpread"in c&&n(12,D=c.xSpread),"destroyOnComplete"in c&&n(17,A=c.destroyOnComplete)},[t,i,s,o,r,f,m,_,C,y,b,I,D,G,j,d,u,A]}class Ve extends le{constructor(e){super(),ie(this,e,Ke,Je,oe,{size:0,x:1,y:2,duration:3,infinite:4,delay:5,colorRange:15,colorArray:16,amount:6,iterationCount:7,fallDistance:8,rounded:9,cone:10,noGravity:11,xSpread:12,destroyOnComplete:17})}}function ye(l,e,n){const t=l.slice();return t[7]=e[n],t}function Xe(l){let e=l[7].message+"",n;return{c(){n=J(e)},m(t,i){z(t,n,i)},p(t,i){i&1&&e!==(e=t[7].message+"")&&te(n,e)},d(t){t&&k(n)}}}function Ze(l){let e;return{c(){e=J(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                    labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                    laboris nisi ut aliquip ex ea commodo consequat.`)},m(n,t){z(n,e,t)},p:M,d(n){n&&k(e)}}}function ve(l){let e,n,t;function i(){return l[6](l[7])}return{c(){e=p("button"),e.textContent="Reveal this hint",w(e,"class","svelte-6e5e0o")},m(s,o){z(s,e,o),n||(t=X(e,"click",i),n=!0)},p(s,o){l=s},d(s){s&&k(e),n=!1,t()}}}function we(l,e){let n,t=e[7].title+"",i,s,o,r,f,d,u;function m(b,I){return b[7].locked?Ze:Xe}let _=m(e),C=_(e),y=e[7].locked&&ve(e);return{key:l,first:null,c(){n=p("h4"),i=J(t),s=R(),o=p("div"),r=p("p"),C.c(),d=R(),y&&y.c(),u=R(),w(n,"class","svelte-6e5e0o"),w(r,"class",f=_e(e[7].locked?"locked":"")+" svelte-6e5e0o"),w(o,"class","svelte-6e5e0o"),this.first=n},m(b,I){z(b,n,I),S(n,i),z(b,s,I),z(b,o,I),S(o,r),C.m(r,null),S(o,d),y&&y.m(o,null),S(o,u)},p(b,I){e=b,I&1&&t!==(t=e[7].title+"")&&te(i,t),_===(_=m(e))&&C?C.p(e,I):(C.d(1),C=_(e),C&&(C.c(),C.m(r,null))),I&1&&f!==(f=_e(e[7].locked?"locked":"")+" svelte-6e5e0o")&&w(r,"class",f),e[7].locked?y?y.p(e,I):(y=ve(e),y.c(),y.m(o,u)):y&&(y.d(1),y=null)},d(b){b&&k(n),b&&k(s),b&&k(o),C.d(),y&&y.d()}}}function xe(l){let e,n,t,i,s=[],o=new Map,r;t=new Re({props:{options:l[2]}});let f=l[0];const d=u=>u[7].id;for(let u=0;u<f.length;u+=1){let m=ye(l,f,u),_=d(m);o.set(_,s[u]=we(_,m))}return{c(){e=p("section"),n=p("p"),E(t.$$.fragment),i=R();for(let u=0;u<s.length;u+=1)s[u].c();w(e,"class","svelte-6e5e0o")},m(u,m){z(u,e,m),S(e,n),L(t,n,null),S(e,i);for(let _=0;_<s.length;_+=1)s[_].m(e,null);r=!0},p(u,[m]){m&3&&(f=u[0],s=Oe(s,m,d,1,u,f,o,e,Ae,we,null,ye))},i(u){r||(T(t.$$.fragment,u),r=!0)},o(u){P(t.$$.fragment,u),r=!1},d(u){u&&k(e),W(t);for(let m=0;m<s.length;m+=1)s[m].d()}}}function $e(l,e,n){let{gameCode:t=""}=e,{puzzleId:i=0}=e;const s=async()=>{const u=await ne(t),{data:m,error:_}=await u.rpc("gethints",{puzzleid:i,gamecode:t});if(_)throw _;n(0,r=m)},o=async function(u){await(await ne(t)).rpc("requesthint",{puzzleid:i,gamecode:t,hintid:u})};let r=[];const f=["Need a hint? Browse our extensive collection below!","Not sure where to go? Try one of our finest hints!","Get your hints here! Freshly plucked from the hint tree!","Psst... Can I interest you in a hint?","We've got the finest hints in all the land. Don't believe me? Try one for free!","Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?"],d=u=>o(u.id);return l.$$set=u=>{"gameCode"in u&&n(3,t=u.gameCode),"puzzleId"in u&&n(4,i=u.puzzleId)},l.$$.update=()=>{l.$$.dirty&24&&t&&i&&s()},[r,o,f,t,i,s,d]}class et extends le{constructor(e){super(),ie(this,e,$e,xe,oe,{gameCode:3,puzzleId:4,refresh:5})}get refresh(){return this.$$.ctx[5]}}function tt(l){let e;return{c(){e=p("p"),e.textContent="Oops! Something went wrong trying to render the puzzle."},m(n,t){z(n,e,t)},p:M,i:M,o:M,d(n){n&&k(e)}}}function nt(l){let e,n=l[0]+"",t;return{c(){e=new Be(!1),t=se(),e.a=t},m(i,s){e.m(n,i,s),z(i,t,s)},p:M,i:M,o:M,d(i){i&&k(t),i&&e.d()}}}function lt(l){let e,n;return e=new Te({}),{c(){E(e.$$.fragment)},m(t,i){L(e,t,i),n=!0},p:M,i(t){n||(T(e.$$.fragment,t),n=!0)},o(t){P(e.$$.fragment,t),n=!1},d(t){W(e,t)}}}function it(l){let e,n,t={ctx:l,current:null,token:null,hasCatch:!0,pending:lt,then:nt,catch:tt,value:0,error:4,blocks:[,,,]};return He(l[0],t),{c(){e=se(),t.block.c()},m(i,s){z(i,e,s),t.block.m(i,t.anchor=s),t.mount=()=>e.parentNode,t.anchor=e,n=!0},p(i,[s]){l=i,Fe(t,l,s)},i(i){n||(T(t.block),n=!0)},o(i){for(let s=0;s<3;s+=1){const o=t.blocks[s];P(o)}n=!1},d(i){i&&k(e),t.block.d(i),t.token=null,t=null}}}const Se=/\$[^$]+?\$/g;function ot(l,e,n){let{gameCode:t=""}=e,{storageSlug:i=""}=e,{content:s=""}=e,o=Promise.all([ne(t),i,s]).then(([r,f,d])=>{const u=d.match(Se)||[];return Promise.all([d,...u.map(m=>m.substring(1,m.length-1)).map(m=>r.storage.from("puzzles").createSignedUrl(`${f}/${m}`,60*60).then(({data:{signedUrl:_},error:C})=>{if(C)throw C;return _}))])}).then(([r,...f])=>r.replace(Se,()=>f.shift()));return l.$$set=r=>{"gameCode"in r&&n(1,t=r.gameCode),"storageSlug"in r&&n(2,i=r.storageSlug),"content"in r&&n(3,s=r.content)},[o,t,i,s]}class st extends le{constructor(e){super(),ie(this,e,ot,it,oe,{gameCode:1,storageSlug:2,content:3})}}function at(l){let e,n=l[6].adventure.name+"",t,i,s=l[6].title+"",o,r,f,d,u,m,_,C,y,b,I,D,A,G,j,c,K,O,Q,Y,Z,V,g,H,B,ae,ce;f=new st({props:{gameCode:l[0].code,storageSlug:l[6].storage_slug,content:l[6].content}});function Ne(a){l[13](a)}function Pe(a){l[14](a)}let ue={};l[0].code!==void 0&&(ue.gameCode=l[0].code),l[0].puzzle!==void 0&&(ue.puzzleId=l[0].puzzle),G=new et({props:ue}),fe.push(()=>he(G,"gameCode",Ne)),fe.push(()=>he(G,"puzzleId",Pe)),l[15](G),Y=new Re({props:{options:l[10]}});function me(a,h){return a[6].next?ft:rt}let x=me(l),F=x(l),q=l[2]&&Ie();return H=new Qe({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[ct,({data:a})=>({6:a}),({data:a})=>a?64:0]},$$scope:{ctx:l}}}),{c(){e=p("h2"),t=J(n),i=J(": "),o=J(s),r=R(),E(f.$$.fragment),d=R(),u=p("section"),m=p("form"),_=p("fieldset"),C=p("legend"),C.textContent="Enter a guess",y=R(),b=p("input"),I=R(),D=p("input"),A=R(),E(G.$$.fragment),K=R(),O=p("dialog"),Q=p("h3"),E(Y.$$.fragment),Z=R(),F.c(),V=R(),q&&q.c(),g=R(),E(H.$$.fragment),w(b,"type","text"),b.disabled=l[3],w(b,"class","svelte-6zlkam"),w(D,"type","submit"),D.value="Submit",D.disabled=l[3],w(_,"class","svelte-6zlkam"),w(u,"class","answer svelte-6zlkam"),w(Q,"class","svelte-6zlkam"),O.open=l[2],w(O,"class","svelte-6zlkam")},m(a,h){z(a,e,h),S(e,t),S(e,i),S(e,o),z(a,r,h),L(f,a,h),z(a,d,h),z(a,u,h),S(u,m),S(m,_),S(_,C),S(_,y),S(_,b),ze(b,l[1]),S(_,I),S(_,D),z(a,A,h),L(G,a,h),z(a,K,h),z(a,O,h),S(O,Q),L(Y,Q,null),S(O,Z),F.m(O,null),z(a,V,h),q&&q.m(a,h),z(a,g,h),L(H,a,h),B=!0,ae||(ce=[X(b,"input",l[11]),X(m,"submit",We(l[12]))],ae=!0)},p(a,h){(!B||h&64)&&n!==(n=a[6].adventure.name+"")&&te(t,n),(!B||h&64)&&s!==(s=a[6].title+"")&&te(o,s);const $={};h&1&&($.gameCode=a[0].code),h&64&&($.storageSlug=a[6].storage_slug),h&64&&($.content=a[6].content),f.$set($),(!B||h&8)&&(b.disabled=a[3]),h&2&&b.value!==a[1]&&ze(b,a[1]),(!B||h&8)&&(D.disabled=a[3]);const re={};!j&&h&1&&(j=!0,re.gameCode=a[0].code,ke(()=>j=!1)),!c&&h&1&&(c=!0,re.puzzleId=a[0].puzzle,ke(()=>c=!1)),G.$set(re),x===(x=me(a))&&F?F.p(a,h):(F.d(1),F=x(a),F&&(F.c(),F.m(O,null))),(!B||h&4)&&(O.open=a[2]),a[2]?q?h&4&&T(q,1):(q=Ie(),q.c(),T(q,1),q.m(g.parentNode,g)):q&&(qe(),P(q,1,1,()=>{q=null}),Ge());const de={};h&33554496&&(de.$$scope={dirty:h,ctx:a}),H.$set(de)},i(a){B||(T(f.$$.fragment,a),T(G.$$.fragment,a),T(Y.$$.fragment,a),T(q),T(H.$$.fragment,a),B=!0)},o(a){P(f.$$.fragment,a),P(G.$$.fragment,a),P(Y.$$.fragment,a),P(q),P(H.$$.fragment,a),B=!1},d(a){a&&k(e),a&&k(r),W(f,a),a&&k(d),a&&k(u),a&&k(A),l[15](null),W(G,a),a&&k(K),a&&k(O),W(Y),F.d(),a&&k(V),q&&q.d(a),a&&k(g),W(H,a),ae=!1,Ye(ce)}}}function ut(l){let e,n;return e=new Te({}),{c(){E(e.$$.fragment)},m(t,i){L(e,t,i),n=!0},p:M,i(t){n||(T(e.$$.fragment,t),n=!0)},o(t){P(e.$$.fragment,t),n=!1},d(t){W(e,t)}}}function rt(l){let e,n,t,i,s;return{c(){e=p("p"),e.textContent="You have completed the adventure!",n=R(),t=p("button"),t.textContent="Continue \xBB",w(e,"class","svelte-6zlkam"),w(t,"class","svelte-6zlkam")},m(o,r){z(o,e,r),z(o,n,r),z(o,t,r),i||(s=X(t,"click",l[17]),i=!0)},p:M,d(o){o&&k(e),o&&k(n),o&&k(t),i=!1,s()}}}function ft(l){let e,n,t,i,s;return{c(){e=p("p"),e.textContent="You have completed this step in the adventure!",n=R(),t=p("button"),t.textContent="Next puzzle \xBB",w(e,"class","svelte-6zlkam"),w(t,"class","svelte-6zlkam")},m(o,r){z(o,e,r),z(o,n,r),z(o,t,r),i||(s=X(t,"click",l[16]),i=!0)},p:M,d(o){o&&k(e),o&&k(n),o&&k(t),i=!1,s()}}}function Ie(l){let e,n,t;return n=new Ve({props:{x:[-5,5],y:[0,.1],delay:[0,2e3],infinite:!0,duration:"5000",amount:"400",fallDistance:"110vh"}}),{c(){e=p("div"),E(n.$$.fragment),w(e,"class","confetti-bg svelte-6zlkam")},m(i,s){z(i,e,s),L(n,e,null),t=!0},i(i){t||(T(n.$$.fragment,i),t=!0)},o(i){P(n.$$.fragment,i),t=!1},d(i){i&&k(e),W(n)}}}function ct(l){let e,n;return e=new Ue({props:{data:l[6]}}),{c(){E(e.$$.fragment)},m(t,i){L(e,t,i),n=!0},p(t,i){const s={};i&64&&(s.data=t[6]),e.$set(s)},i(t){n||(T(e.$$.fragment,t),n=!0)},o(t){P(e.$$.fragment,t),n=!1},d(t){W(e,t)}}}function mt(l){let e,n,t,i;const s=[ut,at],o=[];function r(f,d){return f[4]?0:1}return e=r(l),n=o[e]=s[e](l),{c(){n.c(),t=se()},m(f,d){o[e].m(f,d),z(f,t,d),i=!0},p(f,[d]){let u=e;e=r(f),e===u?o[e].p(f,d):(qe(),P(o[u],1,1,()=>{o[u]=null}),Ge(),n=o[e],n?n.p(f,d):(n=o[e]=s[e](f),n.c()),T(n,1),n.m(t.parentNode,t))},i(f){i||(T(n),i=!0)},o(f){P(n),i=!1},d(f){o[e].d(f),f&&k(t)}}}function dt(l,e,n){let{params:t={}}=e,i={},s="",o=!1,r=!1,f=!0,d=null,u=null,m;ge.set("Puzzad: Loading...");const _=()=>{try{C(),ne(t.code).then(g=>(u=g,g.from("puzzles").select("title, content, next, storage_slug, adventure (name)").eq("id",t.puzzle))).then(y).then(g=>n(6,i=g)).then(()=>{ge.set("Puzzad: "+i.adventure.name+" :"+i.title),b(),n(4,f=!1)}).catch(()=>ee("/game/"+t.code))}catch{ee("/game/"+t.code)}},C=()=>{console.log("Resetting..."),n(4,f=!0),n(2,o=!1),n(1,s=""),n(3,r=!1)},y=({data:g,error:H})=>{if(H)throw H;if(g.length===0)throw"Puzzle not found";return g[0]};Ee(function(){d&&d.disconnect()});const b=async function(){d||(d=await Le(t.code),await d.channel("public:guesses:game=eq."+t.code).on("postgres_changes",{event:"INSERT",schema:"public",table:"guesses",filter:"game=eq."+t.code},I).subscribe())},I=function(g){g.record.puzzle.toString()===t.puzzle&&(g.record.content==="*hint"?m.refresh():g.record.correct?n(2,o=!0):je.add({title:"Incorrect guess",description:g.record.content,duration:1e4,type:"error"}))},D=async function(){n(3,r=!0),await u.from("guesses").insert({content:s,puzzle:t.puzzle,game:t.code}),n(3,r=!1),n(1,s="")},A=async function(){await ee("/game/"+t.code+"/"+i.next)},G=async function(){await ee("/game/"+t.code)},j=["Congratulations!","Well done!","You did it!","Huzzah!","Good job!","Nice work!"];function c(){s=this.value,n(1,s)}const K=()=>D();function O(g){l.$$.not_equal(t.code,g)&&(t.code=g,n(0,t))}function Q(g){l.$$.not_equal(t.puzzle,g)&&(t.puzzle=g,n(0,t))}function Y(g){fe[g?"unshift":"push"](()=>{m=g,n(5,m)})}const Z=()=>A(),V=()=>G();return l.$$set=g=>{"params"in g&&n(0,t=g.params)},l.$$.update=()=>{l.$$.dirty&1&&t.code&&t.puzzle&&_()},[t,s,o,r,f,m,i,D,A,G,j,c,K,O,Q,Y,Z,V]}class kt extends le{constructor(e){super(),ie(this,e,dt,mt,oe,{params:0})}}export{kt as default};
