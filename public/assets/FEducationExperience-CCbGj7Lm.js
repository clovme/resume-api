import{d as e,a as l,g as a,w as t,c as o,F as s,j as u,u as d,e as n,f as r,E as i,s as p,o as c,b as m,l as f,h as v,k as g,C as V,p as h,n as y,K as b,q as _,t as U,B as j,H as k,G as w,y as x,x as C,J as Y}from"./index-CXQ-BXo9.js";/* empty css                   *//* empty css                  *//* empty css                     *//* empty css                    *//* empty css                       */import{R as A}from"./RichText-B_AA1Rrk.js";import{p as E}from"./plus-QJit_dXA.js";const M={class:"split-row-module"},T={class:"split-1"},N=m("label",null,"学校名称",-1),F={class:"split-1"},R=m("label",null,"所学专业",-1),q={class:"split-1",style:{width:"330px"}},z=m("label",null,"就读时间",-1),B={class:"split-1"},G=m("label",null,"学历",-1),H={class:"split-1"},J={class:"split-1"},K=e({__name:"FEducationExperience",setup(e){let K=0;const $=l(!1),D=a("getEducation");function I(){for(const e of D.value)if(e.major||e.name||e.content){if(e.major.trim().length<=0)return void h({message:"请先填写完表单关键信息！",grouping:!0,type:"error"});if(e.name.trim().length<=0)return void h({message:"请先填写完表单关键信息！",grouping:!0,type:"error"});if(e.content.trim().length<=0)return void h({message:"请先填写完表单关键信息！",grouping:!0,type:"error"})}y.post("/education").then((e=>{200===e.data.code?(D.value.push(e.data.data),$.value=!0):h({message:e.data.message,grouping:!0,type:"error"})}))}function L(e,l){const a=l?e-1:e+1;b(D.value,e,a),$.value=!0;let t=[];for(let o=0;o<D.value.length;o++)$.value=!0,t.push(D.value[o].id);y.put("/education/sort",t).then((e=>{h({message:e.data.message,grouping:!0,type:"success"})}))}return t(D.value,(e=>{clearTimeout(K),$.value?$.value=!1:K=setTimeout((function(){p("/education",e),clearTimeout(K)}),1e3)})),(e,l)=>{const a=_,t=U,p=j,b=k,K=Y,O=w,P=x,Q=i,S=C;return c(),o("div",null,[(c(!0),o(s,null,u(d(D),((e,l)=>(c(),o("div",M,[n(Q,{class:"split-row-1"},{default:r((()=>[n(t,{span:5},{default:r((()=>[m("div",T,[N,n(a,{modelValue:e.name,"onUpdate:modelValue":l=>e.name=l,placeholder:"请输入学校名称",clearable:""},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:4},{default:r((()=>[m("div",F,[R,n(a,{modelValue:e.major,"onUpdate:modelValue":l=>e.major=l,placeholder:"请输入专业名称",clearable:""},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:6},{default:r((()=>[m("div",q,[z,n(p,{"value-format":"YYYY-MM",clearable:!1,modelValue:e.startAt,"onUpdate:modelValue":l=>e.startAt=l,style:{width:"120px"},type:"month",placeholder:"开始时间"},null,8,["modelValue","onUpdate:modelValue"]),f("- "),e.toNow?g("",!0):(c(),v(p,{key:0,"value-format":"YYYY-MM",clearable:!1,modelValue:e.endAt,"onUpdate:modelValue":l=>e.endAt=l,style:{width:"120px"},type:"month",placeholder:"结束时间"},null,8,["modelValue","onUpdate:modelValue"])),n(b,{modelValue:e.toNow,"onUpdate:modelValue":l=>e.toNow=l,label:"至今"},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:4},{default:r((()=>[m("div",B,[G,n(O,{style:{width:"100px"},modelValue:e.degree,"onUpdate:modelValue":l=>e.degree=l,placeholder:"请选择学历"},{default:r((()=>[(c(!0),o(s,null,u(d(V).xueli,(e=>(c(),v(K,{key:e,label:e,value:e},null,8,["label","value"])))),128))])),_:2},1032,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:2},{default:r((()=>[m("div",H,[n(P,{onClick:e=>L(l,!0),class:"el-c-button",type:"success",round:""},{default:r((()=>[f("上移")])),_:2},1032,["onClick"]),n(P,{onClick:e=>L(l,!1),class:"el-c-button",type:"success",round:""},{default:r((()=>[f("下移")])),_:2},1032,["onClick"]),n(P,{onClick:l=>{return a=e.id,void y.delete(`/education?id=${a}`).then((e=>{if(200===e.data.code)for(let l=0;l<D.value.length;l++)D.value[l].id===a&&($.value=!0,D.value.splice(l,1));else h({message:e.data.message,grouping:!0,type:"error"})}));var a},class:"el-c-button",type:"danger",round:""},{default:r((()=>[f("删除")])),_:2},1032,["onClick"])])])),_:2},1024)])),_:2},1024),n(Q,{class:"split-row-1"},{default:r((()=>[n(t,null,{default:r((()=>[m("div",J,[n(A,{modelValue:e.content,"onUpdate:modelValue":l=>e.content=l,placeholder:"所修课程、成绩排名、在校的职务、参赛获奖情况等有利于突出个人优势的信息。尽量简洁，突出重点。"},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,null,{default:r((()=>[n(S,{"border-style":"dashed",style:{"border-color":"#ffc69f"}})])),_:1})])),_:2},1024)])))),256)),n(Q,null,{default:r((()=>[n(t,null,{default:r((()=>[n(P,{onClick:I,type:"success",size:"large",icon:E,round:""},{default:r((()=>[f("新增一条教育经历")])),_:1})])),_:1})])),_:1})])}}});export{K as default};