import{S as D,i as E,s as H,h as _,j as c,x as j,r as l,a as h,w as p,a3 as S,C as q,a4 as M,n as P,d as k,_ as N,y as W,f as z,a2 as A,a6 as B}from"./index.bf5e7580.js";function T(n){let e,a=n[2].text+"",f,s;return{c(){e=_("div"),f=j(a),l(e,"role",s=n[2].error?"alert":null),l(e,"class","svelte-5nsm4v")},m(o,i){h(o,e,i),p(e,f)},p(o,i){i&4&&a!==(a=o[2].text+"")&&W(f,a),i&4&&s!==(s=o[2].error?"alert":null)&&l(e,"role",s)},d(o){o&&k(e)}}}function F(n){let e,a,f,s,o,i,w,u,C,m,y,b,v,L,g,I,r=!!n[2].text&&T(n);return{c(){e=_("form"),a=_("label"),a.textContent="Email:",f=c(),s=_("input"),o=c(),i=_("label"),i.textContent="Password:",w=c(),u=_("input"),C=c(),m=_("button"),y=j("Sign In"),b=c(),r&&r.c(),v=c(),L=_("div"),L.innerHTML='Don&#39;t have an account? <a href="#/signup">Sign up!</a>',l(a,"for","email"),l(a,"class","svelte-5nsm4v"),l(s,"id","email"),l(s,"type","email"),l(s,"name","email"),s.required=!0,l(s,"class","svelte-5nsm4v"),l(i,"for","password"),l(i,"class","svelte-5nsm4v"),l(u,"id","password"),l(u,"type","password"),l(u,"name","password"),u.required=!0,l(u,"class","svelte-5nsm4v"),l(m,"type","submit"),m.disabled=n[3],l(m,"class","svelte-5nsm4v"),l(e,"class","svelte-5nsm4v")},m(t,d){h(t,e,d),p(e,a),p(e,f),p(e,s),S(s,n[0]),p(e,o),p(e,i),p(e,w),p(e,u),S(u,n[1]),p(e,C),p(e,m),p(m,y),h(t,b,d),r&&r.m(t,d),h(t,v,d),h(t,L,d),g||(I=[q(s,"input",n[6]),q(u,"input",n[7]),q(m,"click",n[8]),q(e,"submit",M(n[5]))],g=!0)},p(t,[d]){d&1&&s.value!==t[0]&&S(s,t[0]),d&2&&u.value!==t[1]&&S(u,t[1]),d&8&&(m.disabled=t[3]),t[2].text?r?r.p(t,d):(r=T(t),r.c(),r.m(v.parentNode,v)):r&&(r.d(1),r=null)},i:P,o:P,d(t){t&&k(e),t&&k(b),r&&r.d(t),t&&k(v),t&&k(L),g=!1,N(I)}}}function G(n,e,a){let f="",s="",o={error:null,text:null},i=!1;const w=async()=>{a(3,i=!0);const{user:b,error:v}=await z.auth.signInWithPassword({email:f,password:s});a(3,i=!1),v?a(2,o={error:!0,text:v.message}):v||await A("#/")};function u(b){B.call(this,n,b)}function C(){f=this.value,a(0,f)}function m(){s=this.value,a(1,s)}return[f,s,o,i,w,u,C,m,()=>w()]}class K extends D{constructor(e){super(),E(this,e,G,F,H,{})}}export{K as default};
