import{b as e,q as a,n as l,v as t,f as s,g as o,i as n,w as u,j as d,o as i,x as r,e as p,U as c,a6 as m,F as f,r as g,z as v,y as b,Y as V,m as _,M as y,E as h,a7 as U,G as x,p as w,k as j}from"./vendor-CO9sYFt5.js";import{g as k,s as T,i as C,_ as I}from"./index-BmuxSv1O.js";import{p as A}from"./plus-CnM3zOg6.js";const F=e=>(w("data-v-238ccf56"),e=e(),j(),e),G={class:"split-row-module"},q={class:"split-1"},z=F((()=>o("label",null,"报考院校",-1))),E={class:"split-1"},M=F((()=>o("label",null,"报考专业",-1))),Y={class:"split-1"},$={class:"baokao-table"},B={rowspan:"2"},D=["onUpdate:modelValue"],H=["onUpdate:modelValue"],J=F((()=>o("td",null,"删除",-1))),K=["onClick"],L=F((()=>o("b",null,"表格数据可任意编辑",-1))),N={class:"split-1"},O=I(e({__name:"FApplicationInformation",setup(e){const w=a({name:"",score:""});let j=0,I=0;const F=l(!1),O=l(!1),P=k("getApplicationInfo"),Q=k("getCourseGrade");function R(){!w.name||w.name.trim().length<=0?V({title:"提示信息",message:"专业课名称不能为空。",type:"warning"}):!w.score||w.score.trim().length<=0?V({title:"提示信息",message:"专业课成绩不能为空。",type:"warning"}):C.post("/applicationinfo/grade",w).then((e=>{200===e.data.code?(F.value=!0,O.value=!0,Q.value.push(e.data.data),w.name="",w.score=""):_({message:e.data.message,offset:55,grouping:!0,type:"error"})}))}return t(P.value,(e=>{clearTimeout(j),F.value&&(F.value=!1),j=setTimeout((function(){T("/applicationinfo",[e]),clearTimeout(j)}),1e3)})),t(Q.value,(e=>{clearTimeout(I),O.value&&(O.value=!1),I=setTimeout((function(){T("/applicationinfo/grade",e),clearTimeout(I)}),1e3)})),(e,a)=>{const l=y,t=h,V=d,j=U,k=x;return i(),s("div",null,[o("div",G,[n(V,{class:"split-row-1",style:{"margin-bottom":"15px"}},{default:u((()=>[n(t,{span:5},{default:u((()=>[o("div",q,[z,n(l,{modelValue:r(P).name,"onUpdate:modelValue":a[0]||(a[0]=e=>r(P).name=e),placeholder:"请输入报考院校",clearable:""},null,8,["modelValue"])])])),_:1}),n(t,{span:5},{default:u((()=>[o("div",E,[M,n(l,{modelValue:r(P).major,"onUpdate:modelValue":a[1]||(a[1]=e=>r(P).major=e),placeholder:"请输入专业名称",clearable:""},null,8,["modelValue"])])])),_:1}),n(t,{span:14})])),_:1}),n(V,{class:"split-row-1"},{default:u((()=>[r(Q).length>0?(i(),p(t,{key:0,style:{"margin-bottom":"15px"}},{default:u((()=>[o("div",Y,[n(V,null,{default:u((()=>[n(t,{span:24},{default:u((()=>[o("table",$,[o("tr",null,[o("td",B,[c(o("input",{type:"text","onUpdate:modelValue":a[2]||(a[2]=e=>r(P).cname=e)},null,512),[[m,r(P).cname]])]),(i(!0),s(f,null,g(r(Q),(e=>(i(),s("td",null,[c(o("input",{type:"text","onUpdate:modelValue":a=>e.name=a},null,8,D),[[m,e.name]])])))),256))]),o("tr",null,[(i(!0),s(f,null,g(r(Q),(e=>(i(),s("td",null,[c(o("input",{type:"text","onUpdate:modelValue":a=>e.score=a},null,8,H),[[m,e.score]])])))),256))]),o("tr",null,[J,(i(!0),s(f,null,g(r(Q),(e=>(i(),s("td",null,[o("i",{onClick:a=>{return l=e.id,void C.delete(`/applicationinfo/grade?id=${l}`).then((e=>{if(200===e.data.code)for(let a=0;a<Q.value.length;a++)Q.value[a].id===l&&(F.value=!0,O.value=!0,Q.value.splice(a,1));else _({message:e.data.message,offset:55,grouping:!0,type:"error"})}));var l},class:"icon-circle-with-minus"},null,8,K)])))),256))])])])),_:1})])),_:1})])])),_:1})):v("",!0),n(t,{style:{"margin-bottom":"15px"}},{default:u((()=>[n(j,{"border-style":"dashed",style:{"border-color":"#ffc69f"}},{default:u((()=>[L])),_:1})])),_:1}),n(t,{span:11},{default:u((()=>[o("div",N,[n(l,{modelValue:w.name,"onUpdate:modelValue":a[3]||(a[3]=e=>w.name=e),placeholder:"专业课名称，例如：英语",clearable:""},null,8,["modelValue"]),b(": "),n(l,{modelValue:w.score,"onUpdate:modelValue":a[4]||(a[4]=e=>w.score=e),placeholder:"专业课成绩，例如：80",clearable:""},null,8,["modelValue"]),n(k,{onClick:R,class:"el-c-button",icon:A,plain:""},{default:u((()=>[b("添加专业课成绩")])),_:1})])])),_:1})])),_:1})])])}}}),[["__scopeId","data-v-238ccf56"]]);export{O as default};