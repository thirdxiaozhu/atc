<template>
  <el-menu class="el-menu-demo" mode="horizontal">
    <el-menu-item style="font-size: 25px;color:black; margin-left:0%">信息发布商</el-menu-item>
    <el-menu-item style="font-size: 15px; float:right; margin-right:1%" @click="logout">退出登录</el-menu-item>
    <el-menu-item style=" float:right; margin-right: 50px;">
      <el-tag :type=" 'success'" disable-transitions style="font-size: 20px;" @click="showProfile">{{ userid }}
      </el-tag>
    </el-menu-item>
    <el-menu-item style="font-size: 15px; float:right;">
      <span>欢迎您，来自{{ usercompany }} 的 </span>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      userid: this.$store.state.userid,
      usercompany: this.$store.state.cncompany
    }
  },
  methods: {
    activeHandler(index) {
      //当点击导航栏不同按钮的时候，监听到v-for中的index值，并设置其为currentIndex
      //这样的设置可以用来做class="active"功能
      this.currentIndex = index;
    },
    logout() {
      this.$store.commit('clearToken')
      this.$router.push({ path: '/login' });
    },
    showProfile(){
      console.log("!!!!!!!")
    }
  },

  created() {
    console.log(this.$route) //在console中打印当前的路由信息，

    //刷新页面时，和路由信息中的path属性做比对，如果一致，则将当前遍历赋给currentIndex
    //用以做路由保持
    //for(var i = 0; i < this.routes.length ; i++){
    //  if(this.routes[i].url == this.$route.path){
    //    this.currentIndex = i;
    //    return;
    //  }
    //}
  }
}
</script>
