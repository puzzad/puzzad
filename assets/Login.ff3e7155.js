import{S as B,i as D,s as E,l as m,p as b,x as T,r as s,a as k,w as p,a1 as S,B as q,a2 as H,n as I,d as x,Y as M,y as N,f as W,a0 as Y,a4 as j}from"./index.957abd27.js";function P(n){let e,a=n[2].text+"",f,l;return{c(){e=m("div"),f=T(a),s(e,"role",l=n[2].error?"alert":null),s(e,"class","svelte-lcxfu1")},m(o,i){k(o,e,i),p(e,f)},p(o,i){i&4&&a!==(a=o[2].text+"")&&N(f,a),i&4&&l!==(l=o[2].error?"alert":null)&&s(e,"role",l)},d(o){o&&x(e)}}}function z(n){let e,a,f,l,o,i,w,r,L,d,y,v,_,h,C,g,u=!!n[2].text&&P(n);return{c(){e=m("form"),a=m("label"),a.textContent="Email:",f=b(),l=m("input"),o=b(),i=m("label"),i.textContent="Password:",w=b(),r=m("input"),L=b(),d=m("button"),y=T("Sign In"),v=b(),u&&u.c(),_=b(),h=m("div"),h.innerHTML='Don&#39;t have an account? <a href="#/signup">Sign up!</a>',s(a,"for","email"),s(a,"class","svelte-lcxfu1"),s(l,"id","email"),s(l,"type","email"),s(l,"name","email"),l.required=!0,s(l,"class","svelte-lcxfu1"),s(i,"for","password"),s(i,"class","svelte-lcxfu1"),s(r,"id","password"),s(r,"type","password"),s(r,"name","password"),r.required=!0,s(r,"class","svelte-lcxfu1"),s(d,"type","submit"),d.disabled=n[3],s(d,"class","svelte-lcxfu1"),s(e,"class","svelte-lcxfu1"),s(h,"class","svelte-lcxfu1")},m(t,c){k(t,e,c),p(e,a),p(e,f),p(e,l),S(l,n[0]),p(e,o),p(e,i),p(e,w),p(e,r),S(r,n[1]),p(e,L),p(e,d),p(d,y),k(t,v,c),u&&u.m(t,c),k(t,_,c),k(t,h,c),C||(g=[q(l,"input",n[6]),q(r,"input",n[7]),q(d,"click",n[8]),q(e,"submit",H(n[5]))],C=!0)},p(t,[c]){c&1&&l.value!==t[0]&&S(l,t[0]),c&2&&r.value!==t[1]&&S(r,t[1]),c&8&&(d.disabled=t[3]),t[2].text?u?u.p(t,c):(u=P(t),u.c(),u.m(_.parentNode,_)):u&&(u.d(1),u=null)},i:I,o:I,d(t){t&&x(e),t&&x(v),u&&u.d(t),t&&x(_),t&&x(h),C=!1,M(g)}}}function A(n,e,a){let f="",l="",o={error:null,text:null},i=!1;const w=async()=>{a(3,i=!0);const{user:v,error:_}=await W.auth.signInWithPassword({email:f,password:l});a(3,i=!1),_?a(2,o={error:!0,text:_.message}):_||await Y("#/")};function r(v){j.call(this,n,v)}function L(){f=this.value,a(0,f)}function d(){l=this.value,a(1,l)}return[f,l,o,i,w,r,L,d,()=>w()]}class G extends B{constructor(e){super(),D(this,e,A,z,E,{})}}export{G as default};
