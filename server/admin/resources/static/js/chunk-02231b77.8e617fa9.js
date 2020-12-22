(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-02231b77"],{"0520":function(e,t,o){},"234c":function(e,t,o){"use strict";o.r(t);var n=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",[o("Form",{ref:"menuForm",attrs:{model:e.menu,rules:e.ruleValidate,"label-position":"right","label-width":130}},[o("Row",{attrs:{gutter:2}},[o("h4",[e._v("菜单信息")]),o("Col",{attrs:{span:"12",push:"3"}},[o("FormItem",{attrs:{prop:"path",label:"路径："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入路径名称",clearable:""},model:{value:e.menu.path,callback:function(t){e.$set(e.menu,"path",t)},expression:"menu.path"}})],1),o("FormItem",{attrs:{prop:"name",label:"路径名称："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入路径名称",clearable:""},model:{value:e.menu.name,callback:function(t){e.$set(e.menu,"name",t)},expression:"menu.name"}})],1),o("FormItem",{attrs:{label:"模块："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",disabled:!e.menu.isSub,placeholder:"请输入模块名称",clearable:""},model:{value:e.menu.modular,callback:function(t){e.$set(e.menu,"modular",t)},expression:"menu.modular"}})],1),o("FormItem",{attrs:{label:"组件："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",disabled:!e.menu.isSub,placeholder:"请输入组件名称",clearable:""},model:{value:e.menu.component,callback:function(t){e.$set(e.menu,"component",t)},expression:"menu.component"}})],1)],1),o("Col",{attrs:{span:"12"}},[o("FormItem",{attrs:{label:"子菜单："}},[o("i-switch",{attrs:{size:"large"},on:{"on-change":e.isSubChange},model:{value:e.menu.isSub,callback:function(t){e.$set(e.menu,"isSub",t)},expression:"menu.isSub"}},[o("span",{attrs:{slot:"open"},slot:"open"},[e._v("是")]),o("span",{attrs:{slot:"close"},slot:"close"},[e._v("否")])])],1),o("FormItem",{attrs:{label:"父级菜单："}},[o("Select",{staticStyle:{width:"188px"},attrs:{disabled:!e.menu.isSub},on:{"on-open-change":e.loadParentMenus},model:{value:e.menu.parentId,callback:function(t){e.$set(e.menu,"parentId",t)},expression:"menu.parentId"}},e._l(e.parentMenuList,(function(t){return o("Option",{key:t.value,attrs:{value:t.value}},[e._v(e._s(t.label))])})),1),e._v("\n          "+e._s(e.menu.parentId)+"\n        ")],1),o("FormItem",{attrs:{label:"菜单顺序："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"number",number:"",placeholder:"请输入菜单顺序"},model:{value:e.menu.sort,callback:function(t){e.$set(e.menu,"sort",t)},expression:"menu.sort"}})],1)],1)],1),o("h4",[e._v("菜单元信息(meta)")]),o("Row",{attrs:{gutter:2}},[o("Col",{attrs:{span:"12",push:"3"}},[o("FormItem",{attrs:{prop:"meta.title",label:"菜单名称："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入菜单名称",clearable:""},model:{value:e.menu.meta.title,callback:function(t){e.$set(e.menu.meta,"title",t)},expression:"menu.meta.title"}})],1),o("FormItem",{attrs:{prop:"meta.icon",label:"菜单图标："},nativeOn:{click:function(t){e.icon.modal=!e.icon.modal}}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入菜单图标",readonly:"",prefix:e.menu.meta.icon},model:{value:e.menu.meta.icon,callback:function(t){e.$set(e.menu.meta,"icon",t)},expression:"menu.meta.icon"}})],1),o("FormItem",{attrs:{label:"钩子函数："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入钩子函数",clearable:""},model:{value:e.menu.meta.beforeCloseName,callback:function(t){e.$set(e.menu.meta,"beforeCloseName",t)},expression:"menu.meta.beforeCloseName"}})],1),o("FormItem",{attrs:{label:"外部连接："}},[o("Input",{staticStyle:{width:"auto"},attrs:{type:"text",placeholder:"请输入钩子函数",clearable:""},model:{value:e.menu.meta.href,callback:function(t){e.$set(e.menu.meta,"href",t)},expression:"menu.meta.href"}})],1),o("FormItem",{attrs:{label:"权限数组："}},[o("Input",{staticStyle:{width:"auto"},attrs:{disabled:"",type:"text",placeholder:"请输入权限数组",clearable:""},model:{value:e.menu.meta.access,callback:function(t){e.$set(e.menu.meta,"access",t)},expression:"menu.meta.access"}})],1)],1),o("Col",{attrs:{span:"12"}},[o("FormItem",{attrs:{label:"父级菜单显示："}},[o("i-switch",{attrs:{size:"large"},model:{value:e.menu.meta.showAlways,callback:function(t){e.$set(e.menu.meta,"showAlways",t)},expression:"menu.meta.showAlways"}},[o("span",{attrs:{slot:"open"},slot:"open"},[e._v("ON")]),o("span",{attrs:{slot:"close"},slot:"close"},[e._v("OFF")])])],1),o("FormItem",{attrs:{label:"关闭面包屑导航："}},[o("i-switch",{attrs:{size:"large"},model:{value:e.menu.meta.hideInBread,callback:function(t){e.$set(e.menu.meta,"hideInBread",t)},expression:"menu.meta.hideInBread"}},[o("span",{attrs:{slot:"open"},slot:"open"},[e._v("ON")]),o("span",{attrs:{slot:"close"},slot:"close"},[e._v("OFF")])])],1),o("FormItem",{attrs:{label:"不显示菜单："}},[o("i-switch",{attrs:{size:"large"},model:{value:e.menu.meta.hideInMenu,callback:function(t){e.$set(e.menu.meta,"hideInMenu",t)},expression:"menu.meta.hideInMenu"}},[o("span",{attrs:{slot:"open"},slot:"open"},[e._v("ON")]),o("span",{attrs:{slot:"close"},slot:"close"},[e._v("OFF")])])],1),o("FormItem",{attrs:{label:"不缓存页面："}},[o("i-switch",{attrs:{size:"large"},model:{value:e.menu.meta.notCache,callback:function(t){e.$set(e.menu.meta,"notCache",t)},expression:"menu.meta.notCache"}},[o("span",{attrs:{slot:"open"},slot:"open"},[e._v("ON")]),o("span",{attrs:{slot:"close"},slot:"close"},[e._v("OFF")])])],1)],1)],1),o("FormItem",{style:{padding:"0 400px"}},[o("Row",[o("Col",{attrs:{span:"8"}},[o("Button",{attrs:{type:"success",icon:"md-checkmark"},on:{click:e.handleSubmit}},[e._v("保存")])],1),o("Col",{attrs:{span:"8"}},[o("Button",{attrs:{type:"error",icon:"md-return-left"},on:{click:e.handleReset}},[e._v("重置")])],1),o("Col",{attrs:{span:"8"}},[o("Button",{attrs:{type:"warning",icon:"md-close"},on:{click:e.handleCloseTag}},[e._v("关闭")])],1)],1)],1)],1),o("Modal",{attrs:{draggable:"",width:"1200",title:"图标集合"},model:{value:e.icon.modal,callback:function(t){e.$set(e.icon,"modal",t)},expression:"icon.modal"}},[e._l(e.icon.icons,(function(t,n){return o("Row",{key:n,attrs:{gutter:0}},[e._l(t,(function(t,a){return o("Col",{key:""+n+a,attrs:{span:"4"}},[o("Button",{attrs:{type:"text"},on:{click:function(o){return e.selectIcon(t.icon)}}},[o("Icon",{attrs:{size:"20",type:t.icon}}),e._v(" "+e._s(t.icon))],1)],1)})),o("br")],2)})),o("span",{attrs:{slot:"footer"},slot:"footer"})],2)],1)},a=[],r=(o("8e6e"),o("456d"),o("ac6a"),o("bd86")),i=o("2f62"),s=[[{icon:"ios-add"},{icon:"md-add"},{icon:"ios-add-circle"},{icon:"ios-add-circle-outline"},{icon:"md-add-circle"},{icon:"ios-alarm"}],[{icon:"ios-alarm-outline"},{icon:"md-alarm"},{icon:"ios-albums"},{icon:"ios-albums-outline"},{icon:"md-albums"},{icon:"ios-alert"}],[{icon:"ios-alert-outline"},{icon:"md-alert"},{icon:"ios-american-football"},{icon:"ios-american-football-outline"},{icon:"md-american-football"},{icon:"ios-analytics"}],[{icon:"ios-analytics-outline"},{icon:"md-analytics"},{icon:"logo-android"},{icon:"logo-angular"},{icon:"ios-aperture"},{icon:"ios-aperture-outline"}],[{icon:"md-aperture"},{icon:"logo-apple"},{icon:"ios-apps"},{icon:"ios-apps-outline"},{icon:"md-apps"},{icon:"ios-appstore"}],[{icon:"ios-appstore-outline"},{icon:"md-appstore"},{icon:"ios-archive"},{icon:"ios-archive-outline"},{icon:"md-archive"},{icon:"ios-arrow-back"}],[{icon:"md-arrow-back"},{icon:"ios-arrow-down"},{icon:"md-arrow-down"},{icon:"ios-arrow-dropdown"},{icon:"md-arrow-dropdown"},{icon:"ios-arrow-dropdown-circle"}],[{icon:"md-arrow-dropdown-circle"},{icon:"ios-arrow-dropleft"},{icon:"md-arrow-dropleft"},{icon:"ios-arrow-dropleft-circle"},{icon:"md-arrow-dropleft-circle"},{icon:"ios-arrow-dropright"}],[{icon:"md-arrow-dropright"},{icon:"ios-arrow-dropright-circle"},{icon:"md-arrow-dropright-circle"},{icon:"ios-arrow-dropup"},{icon:"md-arrow-dropup"},{icon:"ios-arrow-dropup-circle"}],[{icon:"md-arrow-dropup-circle"},{icon:"ios-arrow-forward"},{icon:"md-arrow-forward"},{icon:"ios-arrow-round-back"},{icon:"md-arrow-round-back"},{icon:"ios-arrow-round-down"}],[{icon:"md-arrow-round-down"},{icon:"ios-arrow-round-forward"},{icon:"md-arrow-round-forward"},{icon:"ios-arrow-round-up"},{icon:"md-arrow-round-up"},{icon:"ios-arrow-up"}],[{icon:"ios-backspace"},{icon:"ios-backspace-outline"},{icon:"md-backspace"},{icon:"ios-barcode"},{icon:"ios-barcode-outline"},{icon:"md-barcode"}],[{icon:"ios-baseball"},{icon:"ios-baseball-outline"},{icon:"md-baseball"},{icon:"ios-basket"},{icon:"ios-basket-outline"},{icon:"md-basket"}]];function c(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),o.push.apply(o,n)}return o}function l(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?c(Object(o),!0).forEach((function(t){Object(r["a"])(e,t,o[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):c(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t))}))}return e}var u={props:{tMenu:{type:Object,default:function(){return null}},routeName:String},name:"OneMenuPage",data:function(){return{menu:{meta:{}},ruleValidate:{path:[{required:!0,message:"路径不能为空",trigger:"blur"}],name:[{required:!0,message:"路径名称不能为空",trigger:"blur"}],"meta.title":[{required:!0,message:"菜单名称不能为空",trigger:"blur"}],"meta.icon":[{required:!0,message:"图标不能为空",trigger:"blur"}]},icon:{icons:s,modal:!1},parentMenuList:[]}},methods:l({},Object(i["d"])(["closeTag"]),{},Object(i["b"])(["handleParentMenus"]),{handleSubmit:function(){var e=this;this.$refs.menuForm.validate((function(t){if(t){if(e.menu.isSub){if(!e.menu.modular)return void e.$Message.error("模块不能为空");if(!e.menu.component)return void e.$Message.error("组件不能为空");if(!e.menu.parentId)return void e.$Message.error("父级菜单不能为空")}e.$emit("submit",e.menu)}}))},handleReset:function(){this.tMenu?this.menu=l({},this.tMenu):this.menu={meta:{}},this.parentMenuList=[]},handleCloseTag:function(){this.closeTag({name:this.routeName}),sessionStorage.removeItem("selectionMenu")},selectIcon:function(e){this.icon.modal=!1,this.$set(this.menu.meta,"icon",e)},isSubChange:function(e){e||(delete this.menu.modular,delete this.menu.component)},loadParentMenus:function(e){var t=this;e&&this.parentMenuList.length<1&&this.handleParentMenus().then((function(e){e&&(t.menu.isSub&&(t.parentMenuList=[]),e.forEach((function(e){var o="未知";e.meta&&(o=e.meta.title),t.parentMenuList.push({value:e.id,label:o})})))})).catch((function(e){return console.log("获取所有父菜单数据错误：",e)}))}}),created:function(){this.tMenu&&(this.menu=l({},this.tMenu))}},m=u,d=(o("528a"),o("2877")),p=Object(d["a"])(m,n,a,!1,null,"2ad39166",null);t["default"]=p.exports},"528a":function(e,t,o){"use strict";var n=o("0520"),a=o.n(n);a.a},"98b7":function(e,t,o){"use strict";o.r(t);var n=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("OneMenuPage",{ref:"orp",attrs:{"route-name":"add_menu_page"},on:{submit:e.save}})},a=[],r=(o("8e6e"),o("ac6a"),o("456d"),o("bd86")),i=o("234c"),s=o("2f62");function c(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),o.push.apply(o,n)}return o}function l(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?c(Object(o),!0).forEach((function(t){Object(r["a"])(e,t,o[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):c(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t))}))}return e}var u={components:{OneMenuPage:i["default"]},name:"add_menu_page",data:function(){return{}},methods:l({},Object(s["b"])(["handleCreateMenu"]),{save:function(e){var t=this;this.handleCreateMenu(e).then((function(e){t.$refs.orp.handleReset(),t.$Message.success("创建菜单成功")})).catch((function(){return t.$Message.error("创建菜单失败")}))}}),created:function(){}},m=u,d=o("2877"),p=Object(d["a"])(m,n,a,!1,null,"54afeb56",null);t["default"]=p.exports}}]);