import{d as e,a as l,g as a,w as t,c as o,F as s,j as u,u as d,e as n,f as r,E as i,s as c,o as p,b as m,l as f,h as g,k as v,p as V,n as h,J as y,q as b,t as _,B as j,H as U,y as k,x as w}from"./index-DgLzXXeY.js";/* empty css                   *//* empty css                    *//* empty css                       *//* empty css                     */import{R as x}from"./RichText-CEoOT6vx.js";import{p as Y}from"./plus-BBYjcoxm.js";const C={class:"split-row-module"},A={class:"split-1"},M=m("label",null,"项目名称",-1),T={class:"split-1"},N=m("label",null,"参与角色",-1),z={class:"split-1"},E=m("label",null,"项目时间",-1),F={class:"split-1",style:{"justify-content":"flex-end"}},P={class:"split-1"},R=e({__name:"FProjectExperience",setup(e){let R=0;const q=l(!1),B=a("getProject");function H(){for(const e of B.value)if(e.title||e.name||e.content){if(e.title.trim().length<=0)return void V({message:"请先填写完表单关键信息！",offset:55,grouping:!0,type:"error"});if(e.name.trim().length<=0)return void V({message:"请先填写完表单关键信息！",offset:55,grouping:!0,type:"error"});if(e.content.trim().length<=0)return void V({message:"请先填写完表单关键信息！",offset:55,grouping:!0,type:"error"})}h.post("/project").then((e=>{200===e.data.code?(B.value.push(e.data.data),q.value=!0):V({message:e.data.message,offset:55,grouping:!0,type:"error"})}))}function J(e,l){const a=l?e-1:e+1;y(B.value,e,a),q.value=!0;let t=[];for(let o=0;o<B.value.length;o++)q.value=!0,t.push(B.value[o].id);h.put("/project/sort",t).then((e=>{V({message:e.data.message,offset:55,grouping:!0,type:"success"})}))}return t(B.value,(e=>{clearTimeout(R),q.value?q.value=!1:R=setTimeout((function(){c("/project",e),clearTimeout(R)}),1e3)})),(e,l)=>{const a=b,t=_,c=j,y=U,R=k,$=i,D=w;return p(),o("div",null,[(p(!0),o(s,null,u(d(B),((e,l)=>(p(),o("div",C,[n($,{class:"split-row-1"},{default:r((()=>[n(t,{span:6},{default:r((()=>[m("div",A,[M,n(a,{modelValue:e.name,"onUpdate:modelValue":l=>e.name=l,placeholder:"请输入项目名称",clearable:""},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:5},{default:r((()=>[m("div",T,[N,n(a,{modelValue:e.title,"onUpdate:modelValue":l=>e.title=l,placeholder:"请输入角色名称",clearable:""},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:8},{default:r((()=>[m("div",z,[E,n(c,{clearable:!1,modelValue:e.startAt,"onUpdate:modelValue":l=>e.startAt=l,"value-format":"YYYY-MM",style:{width:"130px"},type:"month",placeholder:"开始时间"},null,8,["modelValue","onUpdate:modelValue"]),f("- "),e.toNow?v("",!0):(p(),g(c,{key:0,clearable:!1,modelValue:e.endAt,"onUpdate:modelValue":l=>e.endAt=l,"value-format":"YYYY-MM",style:{width:"130px"},type:"month",placeholder:"结束时间"},null,8,["modelValue","onUpdate:modelValue"])),n(y,{modelValue:e.toNow,"onUpdate:modelValue":l=>e.toNow=l,label:"至今"},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,{span:5},{default:r((()=>[m("div",F,[n(R,{onClick:e=>J(l,!0),class:"el-c-button",type:"success",round:""},{default:r((()=>[f("上移")])),_:2},1032,["onClick"]),n(R,{onClick:e=>J(l,!1),class:"el-c-button",type:"success",round:""},{default:r((()=>[f("下移")])),_:2},1032,["onClick"]),n(R,{onClick:l=>{return a=e.id,void h.delete(`/project?id=${a}`).then((e=>{if(200===e.data.code)for(let l=0;l<B.value.length;l++)B.value[l].id===a&&(q.value=!0,B.value.splice(l,1));else V({message:e.data.message,offset:55,grouping:!0,type:"error"})}));var a},class:"el-c-button",type:"danger",round:""},{default:r((()=>[f("删除")])),_:2},1032,["onClick"])])])),_:2},1024)])),_:2},1024),n($,{class:"split-row-1"},{default:r((()=>[n(t,null,{default:r((()=>[m("div",P,[n(x,{modelValue:e.content,"onUpdate:modelValue":l=>e.content=l,placeholder:"请输入项目内容、成果和感悟，简洁突出重点。"},null,8,["modelValue","onUpdate:modelValue"])])])),_:2},1024),n(t,null,{default:r((()=>[n(D,{"border-style":"dashed",style:{"border-color":"#ffc69f"}})])),_:1})])),_:2},1024)])))),256)),n($,null,{default:r((()=>[n(t,null,{default:r((()=>[n(R,{onClick:H,type:"success",size:"large",icon:Y,round:""},{default:r((()=>[f("新增一条项目经验")])),_:1})])),_:1})])),_:1})])}}});export{R as default};