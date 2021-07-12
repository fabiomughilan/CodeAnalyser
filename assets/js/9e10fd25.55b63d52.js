(self.webpackChunkcode_analyser=self.webpackChunkcode_analyser||[]).push([[764],{774:function(e,n,a){"use strict";a.r(n),a.d(n,{frontMatter:function(){return u},contentTitle:function(){return o},metadata:function(){return p},toc:function(){return d},default:function(){return m}});var t=a(2122),l=a(9756),i=(a(7294),a(3905)),r=["components"],u={sidebar_position:1},o="Add Plugins",p={unversionedId:"call-for-ontributions/add-plugin",id:"call-for-ontributions/add-plugin",isDocsHomePage:!1,title:"Add Plugins",description:"This page helps you to add your own plugin",source:"@site/docs/call-for-ontributions/add-plugin.md",sourceDirName:"call-for-ontributions",slug:"/call-for-ontributions/add-plugin",permalink:"/code-analyser/docs/call-for-ontributions/add-plugin",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/call-for-ontributions/add-plugin.md",version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Usage",permalink:"/code-analyser/docs/getting-started/usages"}},d=[{value:"Add a new language",id:"add-a-new-language",children:[]},{value:"Add a new plugin",id:"add-a-new-plugin",children:[{value:"Plugin Type",id:"plugin-type",children:[]}]}],s={toc:d};function m(e){var n=e.components,a=(0,l.Z)(e,r);return(0,i.kt)("wrapper",(0,t.Z)({},s,a,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"add-plugins"},"Add Plugins"),(0,i.kt)("p",null,"This page helps you to add your own plugin"),(0,i.kt)("h2",{id:"add-a-new-language"},"Add a new language"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Need to add new language plugin folder "),(0,i.kt)("li",{parentName:"ul"},"In ",(0,i.kt)("inlineCode",{parentName:"li"},"code-analyser/static/supportedLanguages.yaml")," ")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'languages:\n  - name: "Go"\n    path: "./plugin/go"\n')),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"If folder already exists can continue to the next step")),(0,i.kt)("h2",{id:"add-a-new-plugin"},"Add a new plugin"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"I you want your folder recognized as a plugin you need to have ",(0,i.kt)("inlineCode",{parentName:"li"},"pluginDetails.yaml")),(0,i.kt)("li",{parentName:"ul"},"Structure of ",(0,i.kt)("inlineCode",{parentName:"li"},"pluginDetails.yaml"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre"},'version: "0.1"\nauthor:\nname: "author name"\nemail: "author email"\ndate: "dd-mm-yyyy"\nplugindetails:\ntype: "pluginType"\nname: "name of plugin"\nversion: "v1.x"\ndescription: "plugin description"\nicon: ""\nsemver:\n  - ">=0.x,<1.x"\n  - ">=2.3,<4.x"\ncommand: "node alfredPlugin/server.js" (plugin run command)\n')))),(0,i.kt)("h3",{id:"plugin-type"},"Plugin Type"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Language Specific"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},"detectRuntime"),(0,i.kt)("li",{parentName:"ul"},"commands"),(0,i.kt)("li",{parentName:"ul"},"env"),(0,i.kt)("li",{parentName:"ul"},"preDetectCommand"),(0,i.kt)("li",{parentName:"ul"},"staticAssets"),(0,i.kt)("li",{parentName:"ul"},"buildDirectory"),(0,i.kt)("li",{parentName:"ul"},"testCasesCommands"),(0,i.kt)("li",{parentName:"ul"},"framework"),(0,i.kt)("li",{parentName:"ul"},"orm"),(0,i.kt)("li",{parentName:"ul"},"library"),(0,i.kt)("li",{parentName:"ul"},"database"),(0,i.kt)("li",{parentName:"ul"},"getDependencies"))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Global"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},"dockerFile"),(0,i.kt)("li",{parentName:"ul"},"procFile"),(0,i.kt)("li",{parentName:"ul"},"makeFile"))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Every plugin show follow their respective proto for rpc calls. ")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Server shoould start by chacking available ports.")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Plugin directory must be in it respective language folder."))))}m.isMDXComponent=!0}}]);