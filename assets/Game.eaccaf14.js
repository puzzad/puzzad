import{S as J,i as O,s as U,b as q,c as d,o as V,d as P,p as X,t as S,f as p,I as F,J as K,j as I,H,r as Q,v as W,e as k,k as $,a as z,m as B,A as Z,l as L,B as w,z as v,C as j,n as y,w as G,D as R,h as M,u as x,q as ee}from"./index.b67658b9.js";import{A as te}from"./AdventureLogo.fbf0a8bf.js";import{S as ne}from"./Spinner.c017af1e.js";function N(r,e,l){const n=r.slice();return n[13]=e[l],n}function le(r){let e,l,n,a,i,o;function u(t){r[8](t)}let f={};r[1].adventures.name!==void 0&&(f.name=r[1].adventures.name),l=new te({props:f}),Q.push(()=>W(l,"name",u));function s(t,m){if(t[1].status==="EXPIRED")return ae;if(t[1].status==="PAID")return ie;if(t[1].status==="ACTIVE")return re}let _=s(r),c=_&&_(r);return{c(){e=k("h1"),$(l.$$.fragment),a=z(),c&&c.c(),i=q()},m(t,m){d(t,e,m),B(l,e,null),d(t,a,m),c&&c.m(t,m),d(t,i,m),o=!0},p(t,m){const h={};!n&&m&2&&(n=!0,h.name=t[1].adventures.name,Z(()=>n=!1)),l.$set(h),_===(_=s(t))&&c?c.p(t,m):(c&&c.d(1),c=_&&_(t),c&&(c.c(),c.m(i.parentNode,i)))},i(t){o||(S(l.$$.fragment,t),o=!0)},o(t){P(l.$$.fragment,t),o=!1},d(t){t&&p(e),L(l),t&&p(a),c&&c.d(t),t&&p(i)}}}function oe(r){let e,l,n,a;return{c(){e=k("h1"),e.textContent="Error finding game",l=z(),n=k("p"),a=w(r[4])},m(i,o){d(i,e,o),d(i,l,o),d(i,n,o),v(n,a)},p(i,o){o&16&&j(a,i[4])},i:y,o:y,d(i){i&&p(e),i&&p(l),i&&p(n)}}}function se(r){let e,l;return e=new ne({}),{c(){$(e.$$.fragment)},m(n,a){B(e,n,a),l=!0},p:y,i(n){l||(S(e.$$.fragment,n),l=!0)},o(n){P(e.$$.fragment,n),l=!1},d(n){L(e,n)}}}function re(r){let e,l,n,a,i;return{c(){e=k("p"),e.textContent="This adventure is already in progress!",l=z(),n=k("button"),n.textContent="Go to the current puzzle",G(n,"class","svelte-11i0xjd")},m(o,u){d(o,e,u),d(o,l,u),d(o,n,u),a||(i=R(n,"click",r[10]),a=!0)},p:y,d(o){o&&p(e),o&&p(l),o&&p(n),a=!1,i()}}}function ie(r){let e,l,n,a=r[0].code+"",i,o,u,f,s,_,c;return{c(){e=k("p"),e.textContent=`You've not yet started your adventure! Remember, it's dangerous to go alone.
            Take this if you want to recruit others to help you:`,l=z(),n=k("code"),i=w(a),o=z(),u=k("p"),u.textContent="Once you're ready, press below to go to the first puzzle in the adventure!",f=z(),s=k("button"),s.textContent="Begin the adventure!",G(n,"class","svelte-11i0xjd"),G(s,"class","svelte-11i0xjd")},m(t,m){d(t,e,m),d(t,l,m),d(t,n,m),v(n,i),d(t,o,m),d(t,u,m),d(t,f,m),d(t,s,m),_||(c=R(s,"click",r[9]),_=!0)},p(t,m){m&1&&a!==(a=t[0].code+"")&&j(i,a)},d(t){t&&p(e),t&&p(l),t&&p(n),t&&p(o),t&&p(u),t&&p(f),t&&p(s),_=!1,c()}}}function ae(r){let e,l,n,a,i=r[7](r[1].startTime,r[1].endTime)+"",o,u,f,s,_,c={ctx:r,current:null,token:null,hasCatch:!1,pending:ce,then:fe,catch:ue,value:12};return M(_=r[3],c),{c(){e=k("p"),e.textContent="Congratulations! You finished the adventure!",l=z(),n=k("p"),a=w("You took "),o=w(i),u=w("!"),f=z(),s=q(),c.block.c()},m(t,m){d(t,e,m),d(t,l,m),d(t,n,m),v(n,a),v(n,o),v(n,u),d(t,f,m),d(t,s,m),c.block.m(t,c.anchor=m),c.mount=()=>s.parentNode,c.anchor=s},p(t,m){r=t,m&2&&i!==(i=r[7](r[1].startTime,r[1].endTime)+"")&&j(o,i),c.ctx=r,m&8&&_!==(_=r[3])&&M(_,c)||x(c,r,m)},d(t){t&&p(e),t&&p(l),t&&p(n),t&&p(f),t&&p(s),c.block.d(t),c.token=null,c=null}}}function ue(r){return{c:y,m:y,p:y,d:y}}function fe(r){let e,l,n,a,i=r[12],o=[];for(let u=0;u<i.length;u+=1)o[u]=Y(N(r,i,u));return{c(){e=k("table"),l=k("thead"),l.innerHTML=`<tr><th>Puzzle</th> 
                    <th>Time</th> 
                    <th>Hints</th></tr>`,n=z(),a=k("tbody");for(let u=0;u<o.length;u+=1)o[u].c();G(e,"class","stats")},m(u,f){d(u,e,f),v(e,l),v(e,n),v(e,a);for(let s=0;s<o.length;s+=1)o[s].m(a,null)},p(u,f){if(f&8){i=u[12];let s;for(s=0;s<i.length;s+=1){const _=N(u,i,s);o[s]?o[s].p(_,f):(o[s]=Y(_),o[s].c(),o[s].m(a,null))}for(;s<o.length;s+=1)o[s].d(1);o.length=i.length}},d(u){u&&p(e),ee(o,u)}}}function Y(r){let e,l,n=r[13].title+"",a,i,o,u=r[13].time+"",f,s,_,c=r[13].hints+"",t,m;return{c(){e=k("tr"),l=k("td"),a=w(n),i=z(),o=k("td"),f=w(u),s=z(),_=k("td"),t=w(c),m=z(),G(o,"class","time")},m(h,b){d(h,e,b),v(e,l),v(l,a),v(e,i),v(e,o),v(o,f),v(e,s),v(e,_),v(_,t),v(e,m)},p(h,b){b&8&&n!==(n=h[13].title+"")&&j(a,n),b&8&&u!==(u=h[13].time+"")&&j(f,u),b&8&&c!==(c=h[13].hints+"")&&j(t,c)},d(h){h&&p(e)}}}function ce(r){return{c:y,m:y,p:y,d:y}}function me(r){let e,l,n,a;const i=[se,oe,le],o=[];function u(f,s){return f[2]?0:f[4]||!f[1]?1:2}return e=u(r),l=o[e]=i[e](r),{c(){l.c(),n=q()},m(f,s){o[e].m(f,s),d(f,n,s),a=!0},p(f,[s]){let _=e;e=u(f),e===_?o[e].p(f,s):(V(),P(o[_],1,1,()=>{o[_]=null}),X(),l=o[e],l?l.p(f,s):(l=o[e]=i[e](f),l.c()),S(l,1),l.m(n.parentNode,n))},i(f){a||(S(l),a=!0)},o(f){P(l),a=!1},d(f){o[e].d(f),f&&p(n)}}}function _e(r,e,l){let{params:n={}}=e,a={},i=!0,o,u,f;F(async function(){var g,A;try{f=await K(n.code)}catch(T){l(4,u=T.message||T),l(2,i=!1);return}let{data:b,error:D}=await f.from("games").select("status, puzzle, startTime, endTime, adventures ( name, description)").eq("code",n.code);(g=b[0])!=null&&g.adventures?(l(1,a=b[0]),l(3,o=f.rpc("getstats",{gamecode:n.code}).then(({data:T,error:E})=>E?[]:T).then(T=>{let E=a.startTime;return T.forEach(C=>{C.time=c(E,C.solvetime),E=C.solvetime}),T})),I.set("Puzzad: "+((A=b[0])==null?void 0:A.adventures.name)+" - "+n.code)):(I.set("Puzzad: Game not found"),D="Unable to find game"),D&&l(4,u=D),l(2,i=!1)});const s=async function(){let{data:b}=await f.rpc("startadventure");b&&await H("/game/"+n.code+"/"+b)},_=async function(){await H("/game/"+n.code+"/"+a.puzzle)},c=function(b,D){const g=(new Date(D).getTime()-new Date(b).getTime())/1e3,A=(T,E)=>{const C=Math.floor(T);return C===1?`${C} ${E}`:C>1?`${C} ${E}s`:""};return[A(g%60,"second"),A(g/60%60,"minute"),A(g/3600%24,"hour"),A(g/86400,"day")].filter(T=>T!=="").reverse().join(", ")};function t(b){r.$$.not_equal(a.adventures.name,b)&&(a.adventures.name=b,l(1,a))}const m=()=>s(),h=()=>_();return r.$$set=b=>{"params"in b&&l(0,n=b.params)},[n,a,i,o,u,s,_,c,t,m,h]}class ke extends J{constructor(e){super(),O(this,e,_e,me,U,{params:0})}}export{ke as default};