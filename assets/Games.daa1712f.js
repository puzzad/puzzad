import{S as G,i as M,s as T,h,x as Y,a as d,w as S,g as L,t as _,c as A,b as m,d as g,o as z,f as D,e as E,k as F,l as C,m as j,n as q,p as y,j as H,r as w}from"./index.18d72213.js";import{A as I}from"./AdventureBanner.6a451c80.js";import{S as J}from"./Spinner.4d2ec8b1.js";function U(c,n,a){const t=c.slice();return t[2]=n[a],t}function K(c){let n,a,t=c[0],e=[];for(let o=0;o<t.length;o+=1)e[o]=N(U(c,t,o));const i=o=>_(e[o],1,1,()=>{e[o]=null});let l=null;return t.length||(l=B()),{c(){for(let o=0;o<e.length;o+=1)e[o].c();n=E(),l&&l.c()},m(o,u){for(let r=0;r<e.length;r+=1)e[r].m(o,u);d(o,n,u),l&&l.m(o,u),a=!0},p(o,u){if(u&1){t=o[0];let r;for(r=0;r<t.length;r+=1){const s=U(o,t,r);e[r]?(e[r].p(s,u),m(e[r],1)):(e[r]=N(s),e[r].c(),m(e[r],1),e[r].m(n.parentNode,n))}for(L(),r=t.length;r<e.length;r+=1)i(r);A(),!t.length&&l?l.p(o,u):t.length?l&&(l.d(1),l=null):(l=B(),l.c(),l.m(n.parentNode,n))}},i(o){if(!a){for(let u=0;u<t.length;u+=1)m(e[u]);a=!0}},o(o){e=e.filter(Boolean);for(let u=0;u<e.length;u+=1)_(e[u]);a=!1},d(o){F(e,o),o&&g(n),l&&l.d(o)}}}function O(c){let n,a;return n=new J({}),{c(){C(n.$$.fragment)},m(t,e){j(n,t,e),a=!0},p:q,i(t){a||(m(n.$$.fragment,t),a=!0)},o(t){_(n.$$.fragment,t),a=!1},d(t){y(n,t)}}}function B(c){let n;return{c(){n=h("p"),n.textContent="You haven't purchased any games."},m(a,t){d(a,n,t)},p:q,d(a){a&&g(n)}}}function N(c){var l,o,u,r;let n,a,t,e,i;return a=new I({props:{price:"",status:c[2].status,adventureName:(o=(l=c[2].adventures)==null?void 0:l.name)!=null?o:"Unknown",backgroundUrl:(u=c[2].adventures)==null?void 0:u.promoBackground,logoUrl:(r=c[2].adventures)==null?void 0:r.promoLogo,code:c[2].code}}),{c(){n=h("a"),C(a.$$.fragment),t=H(),w(n,"href",e="/#/game/"+c[2].id)},m(s,f){d(s,n,f),j(a,n,null),S(n,t),i=!0},p(s,f){var k,b,v,$;const p={};f&1&&(p.status=s[2].status),f&1&&(p.adventureName=(b=(k=s[2].adventures)==null?void 0:k.name)!=null?b:"Unknown"),f&1&&(p.backgroundUrl=(v=s[2].adventures)==null?void 0:v.promoBackground),f&1&&(p.logoUrl=($=s[2].adventures)==null?void 0:$.promoLogo),f&1&&(p.code=s[2].code),a.$set(p),(!i||f&1&&e!==(e="/#/game/"+s[2].id))&&w(n,"href",e)},i(s){i||(m(a.$$.fragment,s),i=!0)},o(s){_(a.$$.fragment,s),i=!1},d(s){s&&g(n),y(a)}}}function P(c){let n,a,t,e,i;const l=[O,K],o=[];function u(r,s){return r[1]?0:1}return t=u(c),e=o[t]=l[t](c),{c(){n=h("section"),a=Y(`This is a list of your games.
    `),e.c()},m(r,s){d(r,n,s),S(n,a),o[t].m(n,null),i=!0},p(r,[s]){let f=t;t=u(r),t===f?o[t].p(r,s):(L(),_(o[f],1,1,()=>{o[f]=null}),A(),e=o[t],e?e.p(r,s):(e=o[t]=l[t](r),e.c()),m(e,1),e.m(n,null))},i(r){i||(m(e),i=!0)},o(r){_(e),i=!1},d(r){r&&g(n),o[t].d()}}}function Q(c,n,a){let t=[],e=!0;return z(async function(){let{data:i,error:l}=await D.from("games").select(`
                    id,status, adventures ( name, promoBackground, promoLogo ), status, code
                `);l||(a(1,e=!1),a(0,t=i))}),[t,e]}class X extends G{constructor(n){super(),M(this,n,Q,P,T,{})}}export{X as default};
