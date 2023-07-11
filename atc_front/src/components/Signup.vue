<template>
    <div id="app" style="height:100%" class="login-container">
        <el-form :model="userinfo" status-icon ref="ruleForm" label-position="left" label-width="0px"
            class="demo-ruleForm login-page">
            <h2 class="title" style="margin-bottom: 20px; margin-top: 3px;">注册</h2>
            <el-form-item prop="userid">
                <p style="float: left; margin-bottom: 0px; font-size: small;">请输入账号:</p>
                <el-input type="text" v-model="userinfo.userid" auto-complete="off" placeholder="账号"></el-input>
            </el-form-item>
            <el-form-item prop="password">
                <p style="float: left; margin-bottom: 0px; font-size: small;">请输入密码:</p>
                <el-input type="password" v-model="userinfo.password" auto-complete="off" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item prop="companies">
                <p style="float: left; margin-bottom: 0px; font-size: small;">请选择公司:</p>
                <el-cascader :placeholder="''" style="width: 90%" v-model="userinfo.company"
                    :props="com_cas_props"></el-cascader>
            </el-form-item>
            <el-form-item style="width:100%;">
                <el-button type="primary" style="width:100%;" @click="handleSubmit" :loading="logining">注册</el-button>
            </el-form-item>
            <el-form-item>
                <div>{{ msg }}</div>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
//import 'bootstrap/dist/css/bootstrap.min.css'
import { getCompanies, postSignup } from '../api/axios'
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
            com_cas_props: {
                    lazy: true,
                    lazyLoad(node, resolve) {
                        const level = node.level;
                        //请求参数
                        const requestData = {};
                        if (level === 0) {  
                            resolve([
                                {
                                    value: 0,
                                    label: "信息提供商",
                                },
                                {
                                    value: 1,
                                    label: "信息获取终端",
                                }
                            ])
                        } else if (level === 1) { 
                            requestData.role = node.value;
                            //省市区接口

                            getCompanies(requestData).then(res => {
                                resolve(res.data.data)
                            })
                        }
                    }
                },
        }
    },
    mounted: function () {
    },
    methods: {
        handleSubmit() {
            var that = this
            this.userinfo.company = this.userinfo.company[1]
            postSignup(this.userinfo).then(res => {
                if (res.data.code === 1000){
                    this.$router.push({ path: '/login' });
                }else{
                    that.msg = res.data.data
                }

            })
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