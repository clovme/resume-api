import{b as e,n as a,v as s,o as l,f as t,i as o,w as u,y as n,t as d,A as c,G as m,ab as p}from"./vendor-CO9sYFt5.js";import{_ as r}from"./index-BmuxSv1O.js";const i={class:"check-button-box"},v=r(e({__name:"CheckButton",props:{type:{default:"success"},title:{},class:{},modelValue:{type:Boolean}},emits:["update:modelValue"],setup(e,{emit:r}){const v=e,V=a(v.modelValue),f=r;function y(){f("update:modelValue",V)}return s((()=>v.modelValue),(e=>{V.value=e})),(e,a)=>{const s=m,r=p;return l(),t("div",i,[o(s,{class:c(v.class),type:V.value?v.type:"",round:"",onClick:a[0]||(a[0]=()=>{V.value=!V.value})},{default:u((()=>[n(d(v.title),1)])),_:1},8,["class","type"]),o(r,{modelValue:V.value,"onUpdate:modelValue":a[1]||(a[1]=e=>V.value=e),onChange:y},null,8,["modelValue"])])}}}),[["__scopeId","data-v-60e0645f"]]);export{v as C};