import{b as e,e as t,w as s,o as a,x as l,z as n,f as r,r as o,g as c,t as u,i,F as p,ah as _,E as g,j as f}from"./vendor-CO9sYFt5.js";import{f as d,g as m}from"./index-BmuxSv1O.js";import{_ as v}from"./ModuleTitle.vue_vue_type_script_setup_true_lang-DpZQ9oYw.js";import{_ as k}from"./TitleContent.vue_vue_type_script_setup_true_lang-C3jvPOX9.js";const x=["textContent"],y=e({__name:"PSkillsExpertise",props:{title:{}},setup(e){const y=d.ProficiencyLevel.reduce(((e,t)=>(e[t.value]=t.label,e)),{}),j=m("getSkills");return(e,d)=>{const m=_,h=g,T=f;return a(),t(v,{title:e.title},{default:s((()=>[l(j).content?(a(),t(k,{key:0,content:l(j).content,margin:0},null,8,["content"])):n("",!0),Object.keys(l(j).checkedTags).length>0?(a(),t(T,{key:1,gutter:20,style:{"margin-top":"10px","font-size":"13px"}},{default:s((()=>[(a(!0),r(p,null,o(Object.keys(l(j).checkedTags),(e=>(a(),t(h,{span:8},{default:s((()=>{return[c("label",{textContent:u(e)},null,8,x),i(m,{percentage:l(j).checkedTags[e].level,format:(t=l(j).checkedTags[e],()=>t.isWord?y[t.level]:`${t.level}%`)},null,8,["percentage","format"])];var t})),_:2},1024)))),256))])),_:1})):n("",!0)])),_:1},8,["title"])}}});export{y as default};