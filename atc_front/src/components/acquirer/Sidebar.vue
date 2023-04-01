<template>
  <el-menu default-active="2" class="el-menu-vertical-demo" style=" margin-top: 1px;" router :class="side_menu">
    <el-menu-item>
      <i class="el-icon-setting"></i>
      <!-- <span slot="title">查看信息</span>-->
      <router-link to="/acquirer/checkatc">查看信息</router-link>
    </el-menu-item>
    <el-submenu>
      <template slot="title">
        <i class="el-icon-location"></i>
        <span>数字认证</span>
      </template>
      <el-menu-item>
        <i class="el-icon-setting"></i>
        <router-link to="/acquirer/prikey">我的</router-link>
      </el-menu-item>
      <el-menu-item>
        <i class="el-icon-setting"></i>
        <router-link to="/acquirer/certificate">查看证书</router-link>
      </el-menu-item>
    </el-submenu>

  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      routes: [
        { url: '/', title: '我的首页' },
        { url: '/note', title: '我的笔记' },
      ],
      currentIndex: 0,
    }
  },
  methods: {
    activeHandler(index) {
      //当点击导航栏不同按钮的时候，监听到v-for中的index值，并设置其为currentIndex
      //这样的设置可以用来做class="active"功能
      this.currentIndex = index;
    }
  },

  created() {
    console.log(this.$route) //在console中打印当前的路由信息，

    //刷新页面时，和路由信息中的path属性做比对，如果一致，则将当前遍历赋给currentIndex
    //用以做路由保持
    for (var i = 0; i < this.routes.length; i++) {
      if (this.routes[i].url == this.$route.path) {
        this.currentIndex = i;
        return;
      }
    }
  }
}
</script>

<style scoped>
.router-link-active {
  text-decoration: none;
  color: dodgerblue;
}

a {
  text-decoration: none;
  color: black;
}
</style>