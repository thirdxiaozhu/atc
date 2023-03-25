<template>
  <el-menu  class="el-menu-demo" mode="horizontal" >
    <el-menu-item index="0" style="font-size: 25px;color:black; margin-left:0%">信息发布商</el-menu-item>
    <el-menu-item index="1" style="font-size: 15px; float:right; margin-right:1%" @click="logout">退出登录</el-menu-item>
    <el-menu-item style="float:right;color:white;margin-right:5px">
      <el-input placeholder="请输入内容" v-model="search_text">
        <el-button slot="append" icon="el-icon-search"></el-button>
      </el-input>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
    data(){
        return{
          routes:[
            {url:'/', title:'我的首页'},
            {url:'/note', title:'我的笔记'},
          ],
          currentIndex:0,
          search_text: "",
        }
    },
    methods:{  
      activeHandler(index){   
        //当点击导航栏不同按钮的时候，监听到v-for中的index值，并设置其为currentIndex
        //这样的设置可以用来做class="active"功能
        this.currentIndex = index;
      },
      logout(){
        this.$store.commit('clearToken')
        this.$router.push({ path: '/login' });
      }
    },

    created(){
      console.log(this.$route) //在console中打印当前的路由信息，

      //刷新页面时，和路由信息中的path属性做比对，如果一致，则将当前遍历赋给currentIndex
      //用以做路由保持
      for(var i = 0; i < this.routes.length ; i++){
        if(this.routes[i].url == this.$route.path){
          this.currentIndex = i;
          return;
        }
      }
    }
}
</script>
