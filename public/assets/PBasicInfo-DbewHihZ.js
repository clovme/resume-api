import{d as e,g as l,h as a,f as t,o as i,b as s,u as n,c as r,l as o,D as u,k as p,j as d,F as h,C as c}from"./index-DGQPiGR6.js";import{_ as y}from"./ModuleTitle.vue_vue_type_script_setup_true_lang-CunCvM9H.js";const k={class:"basic-info-box"},g={class:"basic-info-list"},m={key:0},f=s("span",null,"姓名",-1),b={key:1},w=s("span",null,"年龄",-1),_={key:2},S=s("span",null,"生日",-1),x={key:3},v=s("span",null,"性别",-1),D={key:4},P=s("span",null,"婚姻状况",-1),j={key:5},A=s("span",null,"身高体重",-1),B={key:6},C=s("span",null,"身高",-1),E={key:7},I=s("span",null,"体重",-1),T={key:8},Y=s("span",null,"民族",-1),G={key:9},M=s("span",null,"籍贯",-1),N={key:10},F=s("span",null,"政治面貌",-1),q={key:11},z=s("span",null,"工作年限",-1),H={key:12},J=s("span",null,"求职岗位",-1),K={key:13},L=s("span",null,"意向城市",-1),O={key:14},Q=s("span",null,"期望薪资",-1),R={key:15},U=s("span",null,"入职时间",-1),V={key:16},W=s("span",null,"电话",-1),X={key:17},Z=s("span",null,"邮箱",-1),$={class:"photo-box"},ee=["src","alt"],le=e({__name:"PBasicInfo",props:{title:{}},setup(e){const le=l("getBasicInfo");function ae(e){const l=new Date,a=new Date(e),t=(l.getTime()-a.getTime())/315576e5;return Math.floor(t)}return(e,l)=>(i(),a(y,{title:e.title},{default:t((()=>[s("div",k,[s("ul",g,[n(le).name?(i(),r("li",m,[f,o(u(n(le).name),1)])):p("",!0),n(le).birthday&&n(le).isAge?(i(),r("li",b,[w,o(u(ae(n(le).birthday))+"岁",1)])):n(le).birthday?(i(),r("li",_,[S,o(u(n(le).birthday),1)])):p("",!0),n(le).gender&&"不填"!==n(le).gender?(i(),r("li",x,[v,o(u(n(le).gender),1)])):p("",!0),n(le).maritalStatus&&"不填"!==n(le).maritalStatus?(i(),r("li",D,[P,o(u(n(le).maritalStatus),1)])):p("",!0),n(le).height&&n(le).weight?(i(),r("li",j,[A,o(u(n(le).height)+"cm/"+u(n(le).weight)+"kg",1)])):n(le).height?(i(),r("li",B,[C,o(u(n(le).height)+"cm",1)])):n(le).weight?(i(),r("li",E,[I,o(u(n(le).weight)+"kg",1)])):p("",!0),n(le).ethnicGroup?(i(),r("li",T,[Y,o(u(n(le).ethnicGroup),1)])):p("",!0),n(le).nativePlace?(i(),r("li",G,[M,o(u(n(le).nativePlace),1)])):p("",!0),n(le).politicalStatus&&"不填"!==n(le).politicalStatus?(i(),r("li",N,[F,o(u(n(le).politicalStatus),1)])):p("",!0),n(le).workExperienceYears&&"不填"!==n(le).workExperienceYears?(i(),r("li",q,[z,o(u(n(le).workExperienceYears),1)])):p("",!0),n(le).desiredPosition?(i(),r("li",H,[J,o(u(n(le).desiredPosition),1)])):p("",!0),n(le).desiredCity?(i(),r("li",K,[L,o(u(n(le).desiredCity),1)])):p("",!0),n(le).desiredSalary?(i(),r("li",O,[Q,o(u(n(le).desiredSalary),1)])):p("",!0),n(le).startDate&&"不填"!==n(le).startDate?(i(),r("li",R,[U,o(u(n(le).startDate),1)])):p("",!0),n(le).phoneNumber?(i(),r("li",V,[W,o(u(n(le).phoneNumber),1)])):p("",!0),n(le).emailAddress?(i(),r("li",X,[Z,o(u(n(le).emailAddress),1)])):p("",!0),(i(!0),r(h,null,d(n(le).customInfo,((e,l)=>(i(),r("li",null,[s("span",null,u(l),1),o(u(e),1)])))),256))]),s("div",$,[n(le).iShowPhoto?(i(),r("img",{key:0,src:n(le).photo&&n(le).photo.length>0?n(le).photo:n(c).photo(n(le).gender),alt:n(le).gender},null,8,ee)):p("",!0)])])])),_:1},8,["title"]))}});export{le as default};