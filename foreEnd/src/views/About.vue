
<template>
  <div>
    <div class="container">
      <div class="col-md-10 col-md-offset-1">
        <div class="page-header">
          <h3>欧冶云商组织架构树</h3>
		  <h3>上传文件后，请按浏览器左上角的返回按钮，输入公司名称，进行查询</h3>
		  <h3>若输入公司名称后，下方未显示树状图，则公司名输入错误或者公司文件未上传</h3>
        </div>
        <div class="row">
          <div class="col-md-8 col-md-offset-2">
            <form class="form-horizontal row">
              <div class="col-md-4">
                <div class="checkbox">
                  <label>
                    <input
                      type="checkbox"
                      v-model="horizontal"
                    > 水平展示
                  </label>
                </div>
              </div>
              <div class="col-md-4">
                <div class="checkbox">
                  <label>
                    <input
                      type="checkbox"
                      v-model="collapsable"
                    > 展示根节点
                  </label>
                </div>
              </div>
              <div class="col-md-4">
                <div class="checkbox">
                  <label>
                    <input
                      type="checkbox"
                      v-model="expandAll"
                      @change="expandChange"
                    > 展示全部
                  </label>
                </div>
              </div>
              
              
            </form>
          </div>
        </div>
        <p><br></p>
		<div class="flex box1">
		<form action="http://localhost:3000/api/v1/excel/toLead " method="post" enctype="multipart/form-data" >
		          上传文件:<input type="file" name="file" >
		          <input type="submit" value="提交">
		</form>
		</div>
		<input v-model="message" placeholder="请输入公司名称">
		<p>公司名称: {{ message }}</p>
		<input type="button" value="确定" @click="text(message)"/>
        <div class="text-center">
          <vue2-org-tree name="test"
            :data="data"
            :horizontal="horizontal"
            :collapsable="collapsable"
            :label-class-name="labelClassName"
            :render-content="renderContent"
            @on-expand="onExpand"
           
          />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
	 import Axios from 'axios';
export default {
  data() {
    return {
	  message : '',
      data: {},
	  companyName: '',
      horizontal: false,
      collapsable: true,
      expandAll: false,
      labelClassName: "bg-white"
    };
  },
  methods: {
	text:function(cn){
		var url = "http://localhost:3000/api/v1/excel/getJSON";
		            Axios.get(url,{
						params: {
						      companyName: cn
							}
					})
					.then((response)=>{
		                    this.data=JSON.parse(response.data.data);
		                    console.log("请求到的数据66666："+response);
		                }).catch((error)=>{
		                    console.log(error);
		                })
	},
    renderContent(h, data) {
      return data.label;
    },
    onExpand(e,data) {
      if ("expand" in data) {
        data.expand = !data.expand;
        if (!data.expand && data.children) {
          this.collapse(data.children);
        }
      } else {
        this.$set(data, "expand", true);
      }
    },
    collapse(list) {
      var _this = this;
      list.forEach(function(child) {
        if (child.expand) {
          child.expand = false;
        }
        child.children && _this.collapse(child.children);
      });
    },
    expandChange() {
      this.toggleExpand(this.data, this.expandAll);
    },
    toggleExpand(data, val) {
      var _this = this;
      if (Array.isArray(data)) {
        data.forEach(function(item) {
          _this.$set(item, "expand", val);
          if (item.children) {
            _this.toggleExpand(item.children, val);
          }
        });
      } else {
        this.$set(data, "expand", val);
        if (data.children) {
          _this.toggleExpand(data.children, val);
        }
      }
    }
  }
};
</script>
<style type="text/css">
.org-tree-node-label {
  white-space: nowrap;
}
.bg-white {
  background-color: white;
}
.bg-orange {
  background-color: orange;
}
.bg-gold {
  background-color: gold;
}
.bg-gray {
  background-color: gray;
}
.bg-lightpink {
  background-color: lightpink;
}
.bg-chocolate {
  background-color: chocolate;
}
.bg-tomato {
  background-color: tomato;
}
.box1{
	justify-content : center;
}
</style>