import{d as l,g as e,c as a,e as s,f as d,E as t,o,b as u,u as r,F as i,j as n,h as p,C as c,q as m,t as b,G as V,J as f}from"./index-CB57c1PA.js";/* empty css                  *//* empty css                     */const v={class:"split-2"},_=u("label",null,"求职岗位",-1),h={class:"split-2"},y=u("label",null,"意向城市",-1),j={class:"split-2"},U=u("label",null,"期望薪资",-1),C={class:"split-2"},D=u("label",null,"入职时间",-1),P=l({__name:"FJobPosition",setup(l){const P=e("getBasicInfo");return(l,e)=>{const g=m,w=b,x=f,F=V,J=t;return o(),a("div",null,[s(J,{class:"split-row-2"},{default:d((()=>[s(w,{span:6},{default:d((()=>[u("div",v,[_,s(g,{modelValue:r(P).desiredPosition,"onUpdate:modelValue":e[0]||(e[0]=l=>r(P).desiredPosition=l),placeholder:"请输入岗位名称",clearable:""},null,8,["modelValue"])])])),_:1}),s(w,{span:6},{default:d((()=>[u("div",h,[y,s(g,{modelValue:r(P).desiredCity,"onUpdate:modelValue":e[1]||(e[1]=l=>r(P).desiredCity=l),placeholder:"请输入城市名称",clearable:""},null,8,["modelValue"])])])),_:1}),s(w,{span:6},{default:d((()=>[u("div",j,[U,s(g,{modelValue:r(P).desiredSalary,"onUpdate:modelValue":e[2]||(e[2]=l=>r(P).desiredSalary=l),placeholder:"请输入期望薪资",clearable:""},null,8,["modelValue"])])])),_:1}),s(w,{span:6},{default:d((()=>[u("div",C,[D,s(F,{modelValue:r(P).startDate,"onUpdate:modelValue":e[3]||(e[3]=l=>r(P).startDate=l),placeholder:"请选择入职时间"},{default:d((()=>[(o(!0),a(i,null,n(r(c).startDate,(l=>(o(),p(x,{key:l,label:l,value:l},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])])),_:1})])),_:1})])}}});export{P as default};