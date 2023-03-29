<template>
  <div id="app" style="height:100%">

      <!-- <Vheader></Vheader>
      <el-row class="tac" style="height:100%">
        <el-col :span="4" style="height:100%">
          <VSidebar></VSidebar>
        </el-col>
        <el-col :span="20">
          <router-view></router-view>
        </el-col>
      </el-row> -->

      
    <el-container>
      <el-header style="border: 1px"><Vheader></Vheader></el-header>
      <el-container>
        <el-aside width="250px" style="height: 100vh;"><VSidebar></VSidebar></el-aside>
        <el-container>
          <el-main><router-view></router-view></el-main>
        </el-container>
      </el-container>
      </el-container>

  </div>
</template>

<script>
//第一步：引入子组件
import Vheader from '../components/publisher/Header'
import VSidebar from '../components/publisher/Sidebar'
import 'bootstrap/dist/css/bootstrap.min.css'
export default {
  name: 'publisher_view',
  date() {
    return {
      usertype: '',
    }
  },

  //第二步：挂载
  components: {
    Vheader: Vheader,
    VSidebar,
  },
  created() {
    if (this.$store.state.usertype == "") {
      this.$alert('当前用户无法操作该页面，请返回登录', '错误', {
        confirmButtonText: '确定',
      })
      this.$router.push({ path: '/login' })
    }
  },
  methods: {
  },
  mounted: function () {
  }
}
</script>

<style>
html,
body,
#app {
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

  .el-header {
    background-color: #FFF;
    color: #333;
    line-height: 60px;
    padding: 0px;
  }
  
  .el-aside {
    color: #333;
  }
</style>
