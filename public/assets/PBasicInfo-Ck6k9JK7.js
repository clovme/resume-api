import{d as e,g as l,h as t,f as a,o as i,b as s,D as n,u as o,c as r,l as u,H as p,k as d,j as h,F as y,J as c}from"./index-DBIdrAkX.js";import{_ as g}from"./ModuleTitle.vue_vue_type_script_setup_true_lang-CIdGROLw.js";const k={class:"basic-info-box"},m={key:0},f=s("span",null,"姓名",-1),S={key:1},b=s("span",null,"年龄",-1),w={key:2},_=s("span",null,"生日",-1),x={key:3},v=s("span",null,"性别",-1),D={key:4},P=s("span",null,"婚姻状况",-1),j={key:5},z=s("span",null,"身高体重",-1),A={key:6},B=s("span",null,"身高",-1),E={key:7},I=s("span",null,"体重",-1),T={key:8},Y=s("span",null,"民族",-1),$={key:9},C=s("span",null,"籍贯",-1),F={key:10},G=s("span",null,"政治面貌",-1),M={key:11},N=s("span",null,"工作年限",-1),H={key:12},J=s("span",null,"求职岗位",-1),q={key:13},K=s("span",null,"意向城市",-1),L={key:14},O=s("span",null,"期望薪资",-1),Q={key:15},R=s("span",null,"入职时间",-1),U={key:16},V=s("span",null,"电话",-1),W={key:17},X=s("span",null,"邮箱",-1),Z={class:"photo-box"},ee=["src","alt"],le=e({__name:"PBasicInfo",props:{title:{}},setup(e){const le=l("getSetting"),te=l("getBasicInfo");function ae(e){const l=new Date,t=new Date(e),a=(l.getTime()-t.getTime())/315576e5;return Math.floor(a)}return(e,l)=>(i(),t(g,{title:e.title},{default:a((()=>[s("div",k,[s("ul",{class:"basic-info-list",style:n(`font-family: ${o(le).fontFamily};font-size: ${o(le).fontSize}px;line-height: ${o(le).lines*o(le).fontSize*2}px`)},[o(te).name?(i(),r("li",m,[f,u(p(o(te).name),1)])):d("",!0),o(te).birthday&&o(te).isAge?(i(),r("li",S,[b,u(p(ae(o(te).birthday))+"岁",1)])):o(te).birthday?(i(),r("li",w,[_,u(p(o(te).birthday),1)])):d("",!0),o(te).gender&&"不填"!==o(te).gender?(i(),r("li",x,[v,u(p(o(te).gender),1)])):d("",!0),o(te).maritalStatus&&"不填"!==o(te).maritalStatus?(i(),r("li",D,[P,u(p(o(te).maritalStatus),1)])):d("",!0),o(te).height&&o(te).weight?(i(),r("li",j,[z,u(p(o(te).height)+"cm/"+p(o(te).weight)+"kg",1)])):o(te).height?(i(),r("li",A,[B,u(p(o(te).height)+"cm",1)])):o(te).weight?(i(),r("li",E,[I,u(p(o(te).weight)+"kg",1)])):d("",!0),o(te).ethnicGroup?(i(),r("li",T,[Y,u(p(o(te).ethnicGroup),1)])):d("",!0),o(te).nativePlace?(i(),r("li",$,[C,u(p(o(te).nativePlace),1)])):d("",!0),o(te).politicalStatus&&"不填"!==o(te).politicalStatus?(i(),r("li",F,[G,u(p(o(te).politicalStatus),1)])):d("",!0),o(te).workExperienceYears&&"不填"!==o(te).workExperienceYears?(i(),r("li",M,[N,u(p(o(te).workExperienceYears),1)])):d("",!0),o(te).desiredPosition?(i(),r("li",H,[J,u(p(o(te).desiredPosition),1)])):d("",!0),o(te).desiredCity?(i(),r("li",q,[K,u(p(o(te).desiredCity),1)])):d("",!0),o(te).desiredSalary?(i(),r("li",L,[O,u(p(o(te).desiredSalary),1)])):d("",!0),o(te).startDate&&"不填"!==o(te).startDate?(i(),r("li",Q,[R,u(p(o(te).startDate),1)])):d("",!0),o(te).phoneNumber?(i(),r("li",U,[V,u(p(o(te).phoneNumber),1)])):d("",!0),o(te).emailAddress?(i(),r("li",W,[X,u(p(o(te).emailAddress),1)])):d("",!0),(i(!0),r(y,null,h(o(te).customInfo,((e,l)=>(i(),r("li",null,[s("span",null,p(l),1),u(p(e),1)])))),256))],4),s("div",Z,[o(te).iShowPhoto?(i(),r("img",{key:0,src:o(te).photo&&o(te).photo.length>0?o(te).photo:o(c).photo(o(te).gender),alt:o(te).gender},null,8,ee)):d("",!0)])])])),_:1},8,["title"]))}});export{le as default};