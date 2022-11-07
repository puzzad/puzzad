import{S as Y,i as Q,s as F,b as J,c as h,n as G,f as z,I as he,e as y,w as S,K as B,q as ze,y as w,k as E,a as I,z as P,m as O,L as be,M as ke,t as N,d as T,l as A,B as L,x as te,C as X,J as U,D as W,h as pe,u as Ce,N as ye,O as ne,E as ve,P as we,o as me,p as de,j as le,Q as Se,R as $,T as Pe,U as Re,r as Ne}from"./index.b67658b9.js";import{R as _e}from"./RandomText.d8cf2a9e.js";import{S as ge}from"./Spinner.c017af1e.js";import{E as qe}from"./Error.64a5909b.js";function ie(i,e,n){const t=i.slice();return t[18]=e[n],t}function oe(i){let e,n={length:i[6]},t=[];for(let l=0;l<n.length;l+=1)t[l]=se(ie(i,n,l));return{c(){e=y("div");for(let l=0;l<t.length;l+=1)t[l].c();S(e,"class","confetti-holder svelte-io58ff"),B(e,"rounded",i[9]),B(e,"cone",i[10]),B(e,"no-gravity",i[11])},m(l,s){h(l,e,s);for(let o=0;o<t.length;o+=1)t[o].m(e,null)},p(l,s){if(s&20991){n={length:l[6]};let o;for(o=0;o<n.length;o+=1){const a=ie(l,n,o);t[o]?t[o].p(a,s):(t[o]=se(a),t[o].c(),t[o].m(e,null))}for(;o<t.length;o+=1)t[o].d(1);t.length=n.length}s&512&&B(e,"rounded",l[9]),s&1024&&B(e,"cone",l[10]),s&2048&&B(e,"no-gravity",l[11])},d(l){l&&z(e),ze(t,l)}}}function se(i){let e;return{c(){e=y("div"),S(e,"class","confetti svelte-io58ff"),w(e,"--fall-distance",i[8]),w(e,"--size",i[0]+"px"),w(e,"--color",i[14]()),w(e,"--skew",q(-45,45)+"deg,"+q(-45,45)+"deg"),w(e,"--rotation-xyz",q(-10,10)+", "+q(-10,10)+", "+q(-10,10)),w(e,"--rotation-deg",q(0,360)+"deg"),w(e,"--translate-y-multiplier",q(i[2][0],i[2][1])),w(e,"--translate-x-multiplier",q(i[1][0],i[1][1])),w(e,"--scale",.1*q(2,10)),w(e,"--transition-duration",i[4]?`calc(${i[3]}ms * var(--scale))`:`${i[3]}ms`),w(e,"--transition-delay",q(i[5][0],i[5][1])+"ms"),w(e,"--transition-iteration-count",i[4]?"infinite":i[7]),w(e,"--x-spread",1-i[12])},m(n,t){h(n,e,t)},p(n,t){t&256&&w(e,"--fall-distance",n[8]),t&1&&w(e,"--size",n[0]+"px"),t&4&&w(e,"--translate-y-multiplier",q(n[2][0],n[2][1])),t&2&&w(e,"--translate-x-multiplier",q(n[1][0],n[1][1])),t&24&&w(e,"--transition-duration",n[4]?`calc(${n[3]}ms * var(--scale))`:`${n[3]}ms`),t&32&&w(e,"--transition-delay",q(n[5][0],n[5][1])+"ms"),t&144&&w(e,"--transition-iteration-count",n[4]?"infinite":n[7]),t&4096&&w(e,"--x-spread",1-n[12])},d(n){n&&z(e)}}}function Ie(i){let e,n=!i[13]&&oe(i);return{c(){n&&n.c(),e=J()},m(t,l){n&&n.m(t,l),h(t,e,l)},p(t,[l]){t[13]?n&&(n.d(1),n=null):n?n.p(t,l):(n=oe(t),n.c(),n.m(e.parentNode,e))},i:G,o:G,d(t){n&&n.d(t),t&&z(e)}}}function q(i,e){return Math.random()*(e-i)+i}function Te(i,e,n){let{size:t=10}=e,{x:l=[-.5,.5]}=e,{y:s=[.25,1]}=e,{duration:o=2e3}=e,{infinite:a=!1}=e,{delay:u=[0,50]}=e,{colorRange:m=[0,360]}=e,{colorArray:r=[]}=e,{amount:c=50}=e,{iterationCount:d=1}=e,{fallDistance:p="100px"}=e,{rounded:g=!1}=e,{cone:C=!1}=e,{noGravity:v=!1}=e,{xSpread:H=.15}=e,{destroyOnComplete:D=!0}=e,M=!1;he(()=>{!D||a||d=="infinite"||setTimeout(()=>n(13,M=!0),(o+u[1])*d)});function b(){return r.length?r[Math.round(Math.random()*(r.length-1))]:`hsl(${Math.round(q(m[0],m[1]))}, 75%, 50%`}return i.$$set=_=>{"size"in _&&n(0,t=_.size),"x"in _&&n(1,l=_.x),"y"in _&&n(2,s=_.y),"duration"in _&&n(3,o=_.duration),"infinite"in _&&n(4,a=_.infinite),"delay"in _&&n(5,u=_.delay),"colorRange"in _&&n(15,m=_.colorRange),"colorArray"in _&&n(16,r=_.colorArray),"amount"in _&&n(6,c=_.amount),"iterationCount"in _&&n(7,d=_.iterationCount),"fallDistance"in _&&n(8,p=_.fallDistance),"rounded"in _&&n(9,g=_.rounded),"cone"in _&&n(10,C=_.cone),"noGravity"in _&&n(11,v=_.noGravity),"xSpread"in _&&n(12,H=_.xSpread),"destroyOnComplete"in _&&n(17,D=_.destroyOnComplete)},[t,l,s,o,a,u,c,d,p,g,C,v,H,M,b,m,r,D]}class Ge extends Y{constructor(e){super(),Q(this,e,Te,Ie,F,{size:0,x:1,y:2,duration:3,infinite:4,delay:5,colorRange:15,colorArray:16,amount:6,iterationCount:7,fallDistance:8,rounded:9,cone:10,noGravity:11,xSpread:12,destroyOnComplete:17})}}function re(i,e,n){const t=i.slice();return t[7]=e[n],t}function Me(i){let e=i[7].message+"",n;return{c(){n=L(e)},m(t,l){h(t,n,l)},p(t,l){l&1&&e!==(e=t[7].message+"")&&X(n,e)},d(t){t&&z(n)}}}function je(i){let e;return{c(){e=L(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                    labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco
                    laboris nisi ut aliquip ex ea commodo consequat.`)},m(n,t){h(n,e,t)},p:G,d(n){n&&z(e)}}}function ae(i){let e,n,t;function l(){return i[6](i[7])}return{c(){e=y("button"),e.textContent="Reveal this hint",S(e,"class","svelte-6e5e0o")},m(s,o){h(s,e,o),n||(t=W(e,"click",l),n=!0)},p(s,o){i=s},d(s){s&&z(e),n=!1,t()}}}function ue(i,e){let n,t=e[7].title+"",l,s,o,a,u,m,r;function c(C,v){return C[7].locked?je:Me}let d=c(e),p=d(e),g=e[7].locked&&ae(e);return{key:i,first:null,c(){n=y("h4"),l=L(t),s=I(),o=y("div"),a=y("p"),p.c(),m=I(),g&&g.c(),r=I(),S(n,"class","svelte-6e5e0o"),S(a,"class",u=te(e[7].locked?"locked":"")+" svelte-6e5e0o"),S(o,"class","svelte-6e5e0o"),this.first=n},m(C,v){h(C,n,v),P(n,l),h(C,s,v),h(C,o,v),P(o,a),p.m(a,null),P(o,m),g&&g.m(o,null),P(o,r)},p(C,v){e=C,v&1&&t!==(t=e[7].title+"")&&X(l,t),d===(d=c(e))&&p?p.p(e,v):(p.d(1),p=d(e),p&&(p.c(),p.m(a,null))),v&1&&u!==(u=te(e[7].locked?"locked":"")+" svelte-6e5e0o")&&S(a,"class",u),e[7].locked?g?g.p(e,v):(g=ae(e),g.c(),g.m(o,r)):g&&(g.d(1),g=null)},d(C){C&&z(n),C&&z(s),C&&z(o),p.d(),g&&g.d()}}}function De(i){let e,n,t,l,s=[],o=new Map,a;t=new _e({props:{options:i[2]}});let u=i[0];const m=r=>r[7].id;for(let r=0;r<u.length;r+=1){let c=re(i,u,r),d=m(c);o.set(d,s[r]=ue(d,c))}return{c(){e=y("section"),n=y("p"),E(t.$$.fragment),l=I();for(let r=0;r<s.length;r+=1)s[r].c();S(e,"class","svelte-6e5e0o")},m(r,c){h(r,e,c),P(e,n),O(t,n,null),P(e,l);for(let d=0;d<s.length;d+=1)s[d].m(e,null);a=!0},p(r,[c]){c&3&&(u=r[0],s=be(s,c,m,1,r,u,o,e,ke,ue,null,re))},i(r){a||(N(t.$$.fragment,r),a=!0)},o(r){T(t.$$.fragment,r),a=!1},d(r){r&&z(e),A(t);for(let c=0;c<s.length;c+=1)s[c].d()}}}function Ee(i,e,n){let{gameCode:t=""}=e,{puzzleId:l=0}=e;const s=async()=>{const r=await U(t),{data:c,error:d}=await r.rpc("gethints",{puzzleid:l,gamecode:t});if(d)throw d;n(0,a=c)},o=async function(r){await(await U(t)).rpc("requesthint",{puzzleid:l,gamecode:t,hintid:r})};let a=[];const u=["Need a hint? Browse our extensive collection below!","Not sure where to go? Try one of our finest hints!","Get your hints here! Freshly plucked from the hint tree!","Psst... Can I interest you in a hint?","We've got the finest hints in all the land. Don't believe me? Try one for free!","Struggling? In a rush? Why not take one of our hand-picked, artisanal hints?"],m=r=>o(r.id);return i.$$set=r=>{"gameCode"in r&&n(3,t=r.gameCode),"puzzleId"in r&&n(4,l=r.puzzleId)},i.$$.update=()=>{i.$$.dirty&24&&t&&l&&s()},[a,o,u,t,l,s,m]}class Oe extends Y{constructor(e){super(),Q(this,e,Ee,De,F,{gameCode:3,puzzleId:4,refresh:5})}get refresh(){return this.$$.ctx[5]}}function Ae(i){let e,n;return e=new qe({props:{error:i[4]}}),{c(){E(e.$$.fragment)},m(t,l){O(e,t,l),n=!0},p:G,i(t){n||(N(e.$$.fragment,t),n=!0)},o(t){T(e.$$.fragment,t),n=!1},d(t){A(e,t)}}}function He(i){let e,n=i[0]+"",t;return{c(){e=new ye(!1),t=J(),e.a=t},m(l,s){e.m(n,l,s),h(l,t,s)},p:G,i:G,o:G,d(l){l&&z(t),l&&e.d()}}}function Be(i){let e,n;return e=new ge({}),{c(){E(e.$$.fragment)},m(t,l){O(e,t,l),n=!0},p:G,i(t){n||(N(e.$$.fragment,t),n=!0)},o(t){T(e.$$.fragment,t),n=!1},d(t){A(e,t)}}}function Le(i){let e,n,t={ctx:i,current:null,token:null,hasCatch:!0,pending:Be,then:He,catch:Ae,value:0,error:4,blocks:[,,,]};return pe(i[0],t),{c(){e=J(),t.block.c()},m(l,s){h(l,e,s),t.block.m(l,t.anchor=s),t.mount=()=>e.parentNode,t.anchor=e,n=!0},p(l,[s]){i=l,Ce(t,i,s)},i(l){n||(N(t.block),n=!0)},o(l){for(let s=0;s<3;s+=1){const o=t.blocks[s];T(o)}n=!1},d(l){l&&z(e),t.block.d(l),t.token=null,t=null}}}const fe=/\$[^$]+?\$/g;function Ue(i,e,n){let{gameCode:t=""}=e,{storageSlug:l=""}=e,{content:s=""}=e,o=Promise.all([U(t),l,s]).then(([a,u,m])=>{const r=m.match(fe)||[];return Promise.all([m,...r.map(c=>c.substring(1,c.length-1)).map(c=>a.storage.from("puzzles").createSignedUrl(`${u}/${c}`,60*60).then(({data:{signedUrl:d},error:p})=>{if(p)throw p;return d}))])}).then(([a,...u])=>a.replace(fe,()=>u.shift()));return i.$$set=a=>{"gameCode"in a&&n(1,t=a.gameCode),"storageSlug"in a&&n(2,l=a.storageSlug),"content"in a&&n(3,s=a.content)},[o,t,l,s]}class We extends Y{constructor(e){super(),Q(this,e,Ue,Le,F,{gameCode:1,storageSlug:2,content:3})}}function Ye(i){let e,n,t,l,s,o,a,u,m,r;return{c(){e=y("section"),n=y("form"),t=y("fieldset"),l=y("legend"),l.textContent="Enter a guess",s=I(),o=y("input"),a=I(),u=y("input"),S(o,"type","text"),o.disabled=i[1],S(o,"class","svelte-1z90x0"),S(u,"type","submit"),u.value="Submit",u.disabled=i[1],S(t,"class","svelte-1z90x0")},m(c,d){h(c,e,d),P(e,n),P(n,t),P(t,l),P(t,s),P(t,o),ne(o,i[0]),P(t,a),P(t,u),m||(r=[W(o,"input",i[5]),W(n,"submit",ve(i[2]))],m=!0)},p(c,[d]){d&2&&(o.disabled=c[1]),d&1&&o.value!==c[0]&&ne(o,c[0]),d&2&&(u.disabled=c[1])},i:G,o:G,d(c){c&&z(e),m=!1,we(r)}}}function Qe(i,e,n){let{puzzle:t=0}=e,{gameCode:l=""}=e,s="",o=!1;const a=async()=>{n(1,o=!0),await U(l).then(m=>m.from("guesses").insert({content:s,puzzle:t,game:l}).throwOnError()).then(()=>n(0,s="")).finally(()=>n(1,o=!1))};function u(){s=this.value,n(0,s)}return i.$$set=m=>{"puzzle"in m&&n(3,t=m.puzzle),"gameCode"in m&&n(4,l=m.gameCode)},[s,o,a,t,l,u]}class Fe extends Y{constructor(e){super(),Q(this,e,Qe,Ye,F,{puzzle:3,gameCode:4})}}function Je(i){let e,n=i[1].adventure.name+"",t,l,s=i[1].title+"",o,a,u,m,r,c,d,p,g,C,v,H,D,M,b;u=new We({props:{gameCode:i[0].code,storageSlug:i[1].storage_slug,content:i[1].content}}),r=new Fe({props:{gameCode:i[0].code,puzzle:i[0].puzzle}});let _={gameCode:i[0].code,puzzleId:i[0].puzzle};d=new Oe({props:_}),i[8](d),v=new _e({props:{options:i[7]}});function ee(f,k){return f[1].next?Xe:Ve}let K=ee(i),j=K(i),R=i[2]&&ce();return{c(){e=y("h2"),t=L(n),l=L(": "),o=L(s),a=I(),E(u.$$.fragment),m=I(),E(r.$$.fragment),c=I(),E(d.$$.fragment),p=I(),g=y("dialog"),C=y("h3"),E(v.$$.fragment),H=I(),j.c(),D=I(),R&&R.c(),M=J(),S(C,"class","svelte-j9ca2a"),g.open=i[2],S(g,"class","svelte-j9ca2a")},m(f,k){h(f,e,k),P(e,t),P(e,l),P(e,o),h(f,a,k),O(u,f,k),h(f,m,k),O(r,f,k),h(f,c,k),O(d,f,k),h(f,p,k),h(f,g,k),P(g,C),O(v,C,null),P(g,H),j.m(g,null),h(f,D,k),R&&R.m(f,k),h(f,M,k),b=!0},p(f,k){(!b||k&2)&&n!==(n=f[1].adventure.name+"")&&X(t,n),(!b||k&2)&&s!==(s=f[1].title+"")&&X(o,s);const V={};k&1&&(V.gameCode=f[0].code),k&2&&(V.storageSlug=f[1].storage_slug),k&2&&(V.content=f[1].content),u.$set(V);const Z={};k&1&&(Z.gameCode=f[0].code),k&1&&(Z.puzzle=f[0].puzzle),r.$set(Z);const x={};k&1&&(x.gameCode=f[0].code),k&1&&(x.puzzleId=f[0].puzzle),d.$set(x),K===(K=ee(f))&&j?j.p(f,k):(j.d(1),j=K(f),j&&(j.c(),j.m(g,null))),(!b||k&4)&&(g.open=f[2]),f[2]?R?k&4&&N(R,1):(R=ce(),R.c(),N(R,1),R.m(M.parentNode,M)):R&&(me(),T(R,1,1,()=>{R=null}),de())},i(f){b||(N(u.$$.fragment,f),N(r.$$.fragment,f),N(d.$$.fragment,f),N(v.$$.fragment,f),N(R),b=!0)},o(f){T(u.$$.fragment,f),T(r.$$.fragment,f),T(d.$$.fragment,f),T(v.$$.fragment,f),T(R),b=!1},d(f){f&&z(e),f&&z(a),A(u,f),f&&z(m),A(r,f),f&&z(c),i[8](null),A(d,f),f&&z(p),f&&z(g),A(v),j.d(),f&&z(D),R&&R.d(f),f&&z(M)}}}function Ke(i){let e,n;return e=new ge({}),{c(){E(e.$$.fragment)},m(t,l){O(e,t,l),n=!0},p:G,i(t){n||(N(e.$$.fragment,t),n=!0)},o(t){T(e.$$.fragment,t),n=!1},d(t){A(e,t)}}}function Ve(i){let e,n,t,l,s;return{c(){e=y("p"),e.textContent="You have completed the adventure!",n=I(),t=y("button"),t.textContent="Continue \xBB",S(e,"class","svelte-j9ca2a"),S(t,"class","svelte-j9ca2a")},m(o,a){h(o,e,a),h(o,n,a),h(o,t,a),l||(s=W(t,"click",i[10]),l=!0)},p:G,d(o){o&&z(e),o&&z(n),o&&z(t),l=!1,s()}}}function Xe(i){let e,n,t,l,s;return{c(){e=y("p"),e.textContent="You have completed this step in the adventure!",n=I(),t=y("button"),t.textContent="Next puzzle \xBB",S(e,"class","svelte-j9ca2a"),S(t,"class","svelte-j9ca2a")},m(o,a){h(o,e,a),h(o,n,a),h(o,t,a),l||(s=W(t,"click",i[9]),l=!0)},p:G,d(o){o&&z(e),o&&z(n),o&&z(t),l=!1,s()}}}function ce(i){let e,n,t;return n=new Ge({props:{x:[-5,5],y:[0,.1],delay:[0,2e3],infinite:!0,duration:"5000",amount:"400",fallDistance:"110vh"}}),{c(){e=y("div"),E(n.$$.fragment),S(e,"class","confetti-bg svelte-j9ca2a")},m(l,s){h(l,e,s),O(n,e,null),t=!0},i(l){t||(N(n.$$.fragment,l),t=!0)},o(l){T(n.$$.fragment,l),t=!1},d(l){l&&z(e),A(n)}}}function Ze(i){let e,n,t,l;const s=[Ke,Je],o=[];function a(u,m){return u[3]?0:1}return e=a(i),n=o[e]=s[e](i),{c(){n.c(),t=J()},m(u,m){o[e].m(u,m),h(u,t,m),l=!0},p(u,[m]){let r=e;e=a(u),e===r?o[e].p(u,m):(me(),T(o[r],1,1,()=>{o[r]=null}),de(),n=o[e],n?n.p(u,m):(n=o[e]=s[e](u),n.c()),N(n,1),n.m(t.parentNode,t))},i(u){l||(N(n),l=!0)},o(u){T(n),l=!1},d(u){o[e].d(u),u&&z(t)}}}function xe(i,e,n){let{params:t={}}=e,l={},s=!1,o=!0,a=null,u;le.set("Puzzad: Loading...");const m=()=>{r(),U(t.code).then(b=>b.from("puzzles").select("title, content, next, storage_slug, adventure (name)").eq("id",t.puzzle).throwOnError()).then(c).then(b=>n(1,l=b)).then(()=>{le.set("Puzzad: "+l.adventure.name+": "+l.title),d(),n(3,o=!1)}).catch(()=>$("/game/"+t.code))},r=()=>{console.log("Resetting..."),n(3,o=!0),n(2,s=!1)},c=({data:b,error:_})=>{if(_)throw _;if(b.length===0)throw"Puzzle not found";return b[0]};Se(function(){a&&a.disconnect()});const d=async function(){a||(a=await Pe(t.code),await a.channel("public:guesses:game=eq."+t.code).on("postgres_changes",{event:"INSERT",schema:"public",table:"guesses",filter:"game=eq."+t.code},p).subscribe())},p=function(b){b.record.puzzle.toString()===t.puzzle&&(b.record.content==="*hint"?u.refresh():b.record.correct?n(2,s=!0):Re.add({title:"Incorrect guess",description:b.record.content,duration:1e4,type:"error"}))},g=async function(){await $("/game/"+t.code+"/"+l.next)},C=async function(){await $("/game/"+t.code)},v=["Congratulations!","Well done!","You did it!","Huzzah!","Good job!","Nice work!"];function H(b){Ne[b?"unshift":"push"](()=>{u=b,n(4,u)})}const D=()=>g(),M=()=>C();return i.$$set=b=>{"params"in b&&n(0,t=b.params)},i.$$.update=()=>{i.$$.dirty&1&&t.code&&t.puzzle&&m()},[t,l,s,o,u,g,C,v,H,D,M]}class lt extends Y{constructor(e){super(),Q(this,e,xe,Ze,F,{params:0})}}export{lt as default};