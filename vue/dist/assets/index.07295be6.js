import{d as p,r as d,o as f,c as h,a as g,b as L,e as E,f as v,i as y,z as O}from"./vendor.e61da12b.js";const P=function(){const n=document.createElement("link").relList;if(n&&n.supports&&n.supports("modulepreload"))return;for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver(e=>{for(const r of e)if(r.type==="childList")for(const s of r.addedNodes)s.tagName==="LINK"&&s.rel==="modulepreload"&&t(s)}).observe(document,{childList:!0,subtree:!0});function i(e){const r={};return e.integrity&&(r.integrity=e.integrity),e.referrerpolicy&&(r.referrerPolicy=e.referrerpolicy),e.crossorigin==="use-credentials"?r.credentials="include":e.crossorigin==="anonymous"?r.credentials="omit":r.credentials="same-origin",r}function t(e){if(e.ep)return;e.ep=!0;const r=i(e);fetch(e.href,r)}};P();const I=p({setup(o){return(n,i)=>{const t=d("router-view");return f(),h(t)}}});const c=g({state:{collapse:!1,isLogin:!1,token:null,username:null,is_admin:0},mutations:{changeCollapse(o,n){console.log(n),o.collapse=n},loginSucc(o,n){o.token=n,o.isLogin=!0},setUser(o,n){o.username=n.username,o.is_admin=n.is_admin},logout(o,n){o.isLogin=!1}}}),A="modulepreload",l={},S="./",a=function(n,i){return!i||i.length===0?n():Promise.all(i.map(t=>{if(t=`${S}${t}`,t in l)return;l[t]=!0;const e=t.endsWith(".css"),r=e?'[rel="stylesheet"]':"";if(document.querySelector(`link[href="${t}"]${r}`))return;const s=document.createElement("link");if(s.rel=e?"stylesheet":A,e||(s.as="script",s.crossOrigin=""),s.href=t,document.head.appendChild(s),e)return new Promise((m,_)=>{s.addEventListener("load",m),s.addEventListener("error",()=>_(new Error(`Unable to preload CSS for ${t}`)))})})).then(()=>n())},D=[{path:"/",redirect:"/questionList"},{path:"/index",component:()=>a(()=>import("./index.a602eb1f.js"),["assets/index.a602eb1f.js","assets/index.c96025ae.css","assets/vendor.e61da12b.js","assets/plugin-vue_export-helper.21dcd24c.js","assets/login.7b310547.js","assets/login.69812301.css","assets/api.ea65fb67.js"]),children:[{path:"",name:"Home",component:()=>a(()=>import("./home.c39873e0.js"),["assets/home.c39873e0.js","assets/home.17df9f5a.css","assets/plugin-vue_export-helper.21dcd24c.js","assets/vendor.e61da12b.js"])},{path:"/questionList",name:"questionList",component:()=>a(()=>import("./index.1c7aca2e.js"),["assets/index.1c7aca2e.js","assets/index.6e2658c0.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])},{path:"/questionManage",name:"questionManage",component:()=>a(()=>import("./index.8f753a70.js"),["assets/index.8f753a70.js","assets/index.52ced1ad.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])},{path:"/questionDetail",name:"questionDetail",component:()=>a(()=>import("./details.36db549d.js").then(function(o){return o.d}),["assets/details.36db549d.js","assets/details.d95b943e.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])},{path:"/topList",name:"topList",component:()=>a(()=>import("./index.4baa82f8.js"),["assets/index.4baa82f8.js","assets/index.8d6b1428.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])},{path:"/submitList",name:"submitList",component:()=>a(()=>import("./index.7d36121b.js"),["assets/index.7d36121b.js","assets/index.9670f934.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])}]},{path:"/login",component:()=>a(()=>import("./login.7b310547.js"),["assets/login.7b310547.js","assets/login.69812301.css","assets/vendor.e61da12b.js","assets/api.ea65fb67.js","assets/plugin-vue_export-helper.21dcd24c.js"])}],u=L({history:E(),routes:D});u.beforeEach((o,n,i)=>{const t=localStorage.getItem("token");if(!t&&o.path!=="/login")i();else if(o.meta.permission)t==="admin"?i():i("/403");else{if(t&&!c.state.isLogin){const e=localStorage.getItem("username"),r=localStorage.getItem("is_admin");c.commit("loginSucc",t),c.commit("setUser",{username:e,is_admin:r})}i()}});v(I).use(u).use(y,{locale:O}).use(c).mount("#app");