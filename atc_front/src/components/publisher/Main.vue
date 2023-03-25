<template>
<div>
  <el-container>
    <el-header style="text-align: center; font-size: 12px">
    <h3>欢迎您,用户{{ userid }}</h3>
    </el-header>
  </el-container>
</div>
</template>
<script>
export default {
    name:'Main',
    data(){
        return{
            userid:'',
        }
    },
    mounted: function(){
        //vue刚加载时执行的
        //事件需要在vue刚加载时被拉起
        this.initPage();
    },    
    methods:{
        initPage(){
            this.userid=this.$store.state.userid
        }
    }
}
</script>

<style scoped>

</style>