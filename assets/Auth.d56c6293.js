import{S as M,i as Y,s as j,a as h,k as P,c as v,m as q,t as D,d as E,f as A,l as F,Y as B,ad as N,j as T,ae as H,g as z,P as R,e as m,A as L,w as f,y as u,C as k,D as I,n as U,R as y,B as W,T as G}from"./index.1653b910.js";import{T as J,F as K,t as C}from"./FlatToast.4277c5f5.js";function O(l){let e,s,a,n,r,i;return{c(){e=m("p"),s=L("You're already logged in, "),a=m("a"),a.textContent="Logout",n=L("?"),f(a,"href","")},m(t,o){v(t,e,o),u(e,s),u(e,a),u(e,n),r||(i=k(a,"click",I(H)),r=!0)},p:U,d(t){t&&A(e),r=!1,i()}}}function Q(l){let e,s,a,n,r,i,t,o,c,_,w,$,p,g,S;return{c(){e=m("form"),s=m("label"),s.textContent="Email:",a=h(),n=m("input"),r=h(),i=m("label"),i.textContent="Password:",t=h(),o=m("input"),c=h(),_=m("button"),w=L(l[4]),$=h(),p=m("section"),f(s,"for","email"),f(n,"id","email"),f(n,"type","email"),f(n,"name","email"),n.required=!0,f(i,"for","password"),f(o,"id","password"),f(o,"type","password"),f(o,"name","password"),o.required=!0,f(_,"type","submit"),_.disabled=l[6],f(e,"class","basic")},m(d,b){v(d,e,b),u(e,s),u(e,a),u(e,n),y(n,l[1]),u(e,r),u(e,i),u(e,t),u(e,o),y(o,l[2]),u(e,c),u(e,_),u(_,w),u(e,$),u(e,p),p.innerHTML=l[3],g||(S=[k(n,"input",l[8]),k(o,"input",l[9]),k(e,"submit",I(l[0]))],g=!0)},p(d,b){b&2&&n.value!==d[1]&&y(n,d[1]),b&4&&o.value!==d[2]&&y(o,d[2]),b&16&&W(w,d[4]),b&8&&(p.innerHTML=d[3])},d(d){d&&A(e),g=!1,G(S)}}}function V(l){let e,s;return e=new K({props:{data:l[10]}}),{c(){P(e.$$.fragment)},m(a,n){q(e,a,n),s=!0},p(a,n){const r={};n&1024&&(r.data=a[10]),e.$set(r)},i(a){s||(D(e.$$.fragment,a),s=!0)},o(a){E(e.$$.fragment,a),s=!1},d(a){F(e,a)}}}function X(l){let e,s,a;function n(t,o){var c;return(c=t[5])!=null&&c.user?O:Q}let r=n(l),i=r(l);return s=new J({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[V,({data:t})=>({10:t}),({data:t})=>t?1024:0]},$$scope:{ctx:l}}}),{c(){i.c(),e=h(),P(s.$$.fragment)},m(t,o){i.m(t,o),v(t,e,o),q(s,t,o),a=!0},p(t,[o]){r===(r=n(t))&&i?i.p(t,o):(i.d(1),i=r(t),i&&(i.c(),i.m(e.parentNode,e)));const c={};o&3072&&(c.$$scope={dirty:o,ctx:t}),s.$set(c)},i(t){a||(D(s.$$.fragment,t),a=!0)},o(t){E(s.$$.fragment,t),a=!1},d(t){i.d(t),t&&A(e),F(s,t)}}}function Z(l,e,s){let a;B(l,N,p=>s(5,a=p));let{type:n="login"}=e,r,i,t,o,c;const _=async()=>{let p,g;switch(n){case"signup":({error:g}=await z.auth.signUp({email:r,password:i})),p="An email has been sent to you for verification!";break;case"login":({error:g}=await z.auth.signInWithPassword({email:r,password:i})),p="Login success, redirecting.",await R("/");break}g?C.add({title:"Error",description:g.message,duration:1e4,type:"error"}):C.add({title:"Success",description:p,duration:1e4,type:"success"})};switch(n){case"signup":T.set("Puzzad: Signup"),o="Already have an account? <a href='#/Login'>Login!</a>",c="Sign up";break;case"login":T.set("Puzzad: Login"),o="Don't have an account? <a href='#/signup'>Sign up!</a>",c="Sign In";break;case"logout":T.set("Puzzad: Logout"),H()}function w(){r=this.value,s(1,r)}function $(){i=this.value,s(2,i)}return l.$$set=p=>{"type"in p&&s(7,n=p.type)},[_,r,i,o,c,a,t,n,w,$]}class te extends M{constructor(e){super(),Y(this,e,Z,X,j,{type:7,authAction:0})}get authAction(){return this.$$.ctx[0]}}export{te as default};
