import{S as q,i as H,s as M,C as g,e as v,a as f,D as S,b as $,t as w,d as p,f as k,E as N,h as d,z as b,w as T,x as h,n as s,l as A,j as y,m as C,p as L,F as j}from"./index.7183b5fd.js";import{S as z}from"./Spinner.40955dfc.js";import{A as D}from"./AdventureLogo.d7d4bfe2.js";function E(r){let e,a=r[6].message+"",t;return{c(){e=d("p"),t=b(a),T(e,"color","red")},m(n,l){f(n,e,l),h(e,t)},p:s,i:s,o:s,d(n){n&&p(e)}}}function F(r){let e,a,t,n,l=r[2].description+"",i,m,o;a=new D({props:{name:r[2].name}});let u={ctx:r,current:null,token:null,hasCatch:!1,pending:J,then:Y,catch:I,value:5};return g(r[0],u),{c(){e=d("h2"),A(a.$$.fragment),t=y(),n=d("section"),i=y(),m=v(),u.block.c()},m(c,_){f(c,e,_),C(a,e,null),f(c,t,_),f(c,n,_),n.innerHTML=l,f(c,i,_),f(c,m,_),u.block.m(c,u.anchor=_),u.mount=()=>m.parentNode,u.anchor=m,o=!0},p(c,_){r=c,S(u,r,_)},i(c){o||($(a.$$.fragment,c),o=!0)},o(c){w(a.$$.fragment,c),o=!1},d(c){c&&p(e),L(a),c&&p(t),c&&p(n),c&&p(i),c&&p(m),u.block.d(c),u.token=null,u=null}}}function I(r){return{c:s,m:s,p:s,d:s}}function Y(r){let e;function a(l,i){return l[5]?G:B}let n=a(r)(r);return{c(){n.c(),e=v()},m(l,i){n.m(l,i),f(l,e,i)},p(l,i){n.p(l,i)},d(l){n.d(l),l&&p(e)}}}function B(r){let e;return{c(){e=d("p"),e.innerHTML='You must <a href="/#/login">login</a> before you can start an adventure.'},m(a,t){f(a,e,t)},p:s,d(a){a&&p(e)}}}function G(r){let e,a,t=r[2].price===0?"free":"\xA3"+r[2].price,n,l,i;function m(){return r[4](r[2])}return{c(){e=d("button"),a=b("Start this adventure for "),n=b(t)},m(o,u){f(o,e,u),h(e,a),h(e,n),l||(i=j(e,"click",m),l=!0)},p(o,u){r=o},d(o){o&&p(e),l=!1,i()}}}function J(r){return{c:s,m:s,p:s,d:s}}function K(r){let e,a;return e=new z({}),{c(){A(e.$$.fragment)},m(t,n){C(e,t,n),a=!0},p:s,i(t){a||($(e.$$.fragment,t),a=!0)},o(t){w(e.$$.fragment,t),a=!1},d(t){L(e,t)}}}function O(r){let e,a,t={ctx:r,current:null,token:null,hasCatch:!0,pending:K,then:F,catch:E,value:2,error:6,blocks:[,,,]};return g(r[2],t),{c(){e=v(),t.block.c()},m(n,l){f(n,e,l),t.block.m(n,t.anchor=l),t.mount=()=>e.parentNode,t.anchor=e,a=!0},p(n,[l]){r=n,S(t,r,l)},i(n){a||($(t.block),a=!0)},o(n){for(let l=0;l<3;l+=1){const i=t.blocks[l];w(i)}a=!1},d(n){n&&p(e),t.block.d(n),t.token=null,t=null}}}function P(r,e,a){let{params:t={}}=e,n=k.from("adventures").select("id,name,description,price").eq("name",t.name).then(({data:o,error:u})=>{if(u)throw u;return o[0]}),l=k.auth.getSession().then(({data:{session:o}})=>!!(o!=null&&o.user));const i=async o=>{if(o.price===0){let{data:u,error:c}=await k.rpc("addfreeadventure",{adventureid:o.id});u&&await N("/game/"+u)}},m=o=>i(o);return r.$$set=o=>{"params"in o&&a(3,t=o.params)},[l,i,n,t,m]}class V extends q{constructor(e){super(),H(this,e,P,O,M,{params:3})}}export{V as default};