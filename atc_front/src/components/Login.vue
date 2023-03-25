<template>
    <div id="app" style="height:100%" class="login-container">
        <el-form :model="userinfo" status-icon ref="ruleForm" label-position="left" label-width="0px"
            class="demo-ruleForm login-page">
            <h2 class="title" style="margin-bottom: 20px; margin-top: 3px;">登录</h2>
            <el-form-item prop="userid">
                <p style="float: left; margin-bottom: 0px; font-size: small;">请输入账号:</p>
                <el-input type="text" v-model="userinfo.userid" auto-complete="off" placeholder="账号"></el-input>
            </el-form-item>
            <el-form-item prop="password">
                <p style="float: left; margin-bottom: 0px; font-size: small;">请输入密码:</p>
                <el-input type="password" v-model="userinfo.password" auto-complete="off" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item style="width:100%;">
                <el-button type="primary" style="width:100%;" @click="handleSubmit" :loading="logining">登录</el-button>
            </el-form-item>
            <el-form-item style="width:100%;">
                <el-button type="primary" style="width:100%;" @click="handleToSign">注册</el-button>
            </el-form-item>
            <el-form-item>
                <div>{{ msg }}</div>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
//import 'bootstrap/dist/css/bootstrap.min.css'
import { postLogin } from '../api/axios'
export default {
    data() {
        return {
            logining: false,
            companies: [],
            companySelected: '',
            msg: '',
            userinfo: {
                userid: '',
                password: '',
                company: '',
            },
            
        }
    },
    mounted: function () {
        console.log(this.$store);
    },
    methods: {
        handleSubmit() {
            var that = this
            postLogin(this.userinfo).then(res => {
                console.log(res)
                this.json = res.data.data
                if (res.data.code === 1000){
                    that.$store.commit('saveToken',{token: this.json.UUID,userid: this.json.Userid,
                                        usercompany: this.json.Company})
                    console.log(that.$store)
                    this.$router.push({ path: '/publisher' });
                }else{
                    that.msg = res.data.data
                }
            })
        },
        handleToSign(){
            this.$router.push({ path: '/signup' });
        }
    }
};
</script>

<style>
html,
body,
#app {
    text-align: center;
    color: #2c3e50;
    height: 100%;
}

.login-container {
    width: 100%;
    height: 100%;
}

.login-page {
    -webkit-border-radius: 5px;
    border-radius: 20px;
    margin: 150px auto;
    width: 350px;
    height: auto;
    padding: 20px 35px 15px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
}

label.el-checkbox.rememberme {
    margin: 0px 0px 15px;
    text-align: left;
}
</style>