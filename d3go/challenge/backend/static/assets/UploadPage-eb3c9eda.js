/* empty css             */import{h as m,i as w,j as b,u as v,o as k,k as C,w as a,b as t,a as l,t as n,l as d,d as T,m as s,n as x,p as E,q as P,s as U,v as y,_ as N}from"./index-29e6e13c.js";const z={class:"el-upload__text"},B={class:"el-upload__tip tip"},I=m({__name:"UploadPage",setup(V){const{t:o}=w.global;let c=b();const i=v(),u={children:"children",label:"label"},r=e=>{e.children||(window.location.href=e.url)},p=async e=>{if(e.status_code===-1){s.warning(o("upload.unauthorized")),setTimeout(()=>{i.push({path:"/admin"})},1e3),console.log("unauthorized");return}if(e.status_code===-2){s.warning(o("upload.forbidden")),console.log("forbidden");return}if(e.status_code!==0){s.warning(o("upload.failed")+":"+e.status_msg),console.log(e);return}s.success(o("upload.success")),console.log(e.data),c.value=e.data};return(e,$)=>{const _=x,f=E,g=P,h=U;return k(),C(h,{class:"upload-container position-center"},{default:a(()=>[t(f,{accept:".zip",drag:"",action:"/admin/upload","show-file-list":!1,"on-success":p},{tip:a(()=>[l("div",B,n(e.$t("upload.hint")),1)]),default:a(()=>[t(_,{class:"el-icon--upload"},{default:a(()=>[t(d(y))]),_:1}),l("div",z,[T(n(e.$t("upload.drag")),1),l("em",null,n(e.$t("upload.click")),1)])]),_:1}),t(g,{data:d(c),onNodeClick:r,props:u},null,8,["data"])]),_:1})}}});const D=N(I,[["__scopeId","data-v-33bcc611"]]);export{D as default};