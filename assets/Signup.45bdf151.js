import{S as E,i as L,s as N,h as c,j as b,z as O,l as U,u as i,a as F,x as l,N as w,m as j,F as y,O as B,b as k,t as z,d as T,p as A,P as D,f as G}from"./index.7183b5fd.js";import{T as H,F as I,t as P}from"./FlatToast.1a101e46.js";function J(r){let t,n;return t=new I({props:{data:r[4]}}),{c(){U(t.$$.fragment)},m(s,a){j(t,s,a),n=!0},p(s,a){const u={};a&16&&(u.data=s[4]),t.$set(u)},i(s){n||(k(t.$$.fragment,s),n=!0)},o(s){z(t.$$.fragment,s),n=!1},d(s){A(t,s)}}}function K(r){let t,n,s,a,u,d,g,o,h,f,_,S,m,$,v,C;return m=new H({props:{placement:"bottom-right",theme:"dark",$$slots:{default:[J,({data:e})=>({4:e}),({data:e})=>e?16:0]},$$scope:{ctx:r}}}),{c(){t=c("form"),n=c("label"),n.textContent="Email:",s=b(),a=c("input"),u=b(),d=c("label"),d.textContent="Password:",g=b(),o=c("input"),h=b(),f=c("button"),_=O("Sign Up"),S=b(),U(m.$$.fragment),i(n,"for","email"),i(a,"id","email"),i(a,"type","email"),i(a,"name","email"),a.required=!0,i(d,"for","password"),i(o,"id","password"),i(o,"type","password"),i(o,"name","password"),o.required=!0,i(f,"type","submit"),f.disabled=r[2],f.value="Sign up",i(t,"class","basic")},m(e,p){F(e,t,p),l(t,n),l(t,s),l(t,a),w(a,r[0]),l(t,u),l(t,d),l(t,g),l(t,o),w(o,r[1]),l(t,h),l(t,f),l(f,_),F(e,S,p),j(m,e,p),$=!0,v||(C=[y(a,"input",r[5]),y(o,"input",r[6]),y(t,"submit",B(r[3]))],v=!0)},p(e,[p]){p&1&&a.value!==e[0]&&w(a,e[0]),p&2&&o.value!==e[1]&&w(o,e[1]),(!$||p&4)&&(f.disabled=e[2]);const q={};p&144&&(q.$$scope={dirty:p,ctx:e}),m.$set(q)},i(e){$||(k(m.$$.fragment,e),$=!0)},o(e){z(m.$$.fragment,e),$=!1},d(e){e&&T(t),e&&T(S),A(m,e),v=!1,D(C)}}}function M(r,t,n){let s="",a="",u=!1,d={};const g=async()=>{n(2,u=!0);const{_:f,error:_}=await G.auth.signUp({email:s,password:a});n(2,u=!1),_?P.add({title:"Signup",description:_.message,duration:1e4,type:"error"}):P.add({title:"Signup",description:"An email has been sent to you for verification!",duration:1e4,type:"success"})};function o(){s=this.value,n(0,s)}function h(){a=this.value,n(1,a)}return[s,a,u,g,d,o,h]}class V extends E{constructor(t){super(),L(this,t,M,K,N,{})}}export{V as default};