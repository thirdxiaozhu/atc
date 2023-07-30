<template>
    <el-container>
        <el-header>
            <el-button type="primary" @click="dialogFormVisible=true" v-loading.fullscreen.lock="fullscreenLoading">注册执飞信息</el-button>
        </el-header>
        <el-main>
            <el-table :data="tableData" style="width: 100%" v-loading="loading">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" inline class="demo-table-expand">
                            <el-form-item label="始发地">
                                <span>{{ props.row.origin }}</span>
                            </el-form-item>
                            <el-form-item label="目的地">
                                <span>{{ props.row.dest }}</span>
                            </el-form-item>
                            <br>
                            <el-form-item label="权限">
                                <span>{{ props.row.auths }}</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column label="飞机注册号" prop="arn">
                </el-table-column>
                <el-table-column label="航班号" prop="flight">
                </el-table-column>
                <el-table-column label="航空公司" prop="company">
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button size="mini" @click="dialogFormVisible = true; current_row = scope.row"
                            :disabled="scope.row.disabled">编辑</el-button>
                        <el-button size="mini" type="info" @click="handleHistory(scope.$index, scope.row)">历史</el-button>
                        <el-button size="mini" type="success" @click="handleDetail(scope.row)">详情</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-main>
        <el-dialog title="新增" :visible.sync="dialogFormVisible">
            <el-form :model="editform">
                <el-form-item label="飞机注册号">
            <el-select v-model="opt_form.opt_arn" filterable placeholder="请选择">
                <el-option v-for="item in options_arn" :key="item.ID" :label="item.Flight_Arn" :value="item.ID">
                </el-option>
            </el-select>
                </el-form-item>
                <el-form-item label="航线">
            <el-select v-model="opt_form.opt_route" filterable placeholder="请选择">
                <el-option v-for="item in options_route" :key="item.ID" :label="item.Flight_num" :value="item.ID">
                </el-option>
            </el-select>
                </el-form-item>
                <el-form-item label="权限">
            <el-select v-model="opt_form.opt_auth" multiple    collapse-tags filterable placeholder="请选择">
                <el-option v-for="item in options_auth" :key="item.ID" :label="item.Auth_name" :value="item.ID">
                </el-option>
            </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="handleRegist()">确定</el-button>
                <el-button @click="dialogFormVisible = false">取消</el-button>
            </div>
        </el-dialog>
    </el-container>
</template>

<script>
import {  getLinkOptions, postRegLink, getRegistLink } from '../../api/axios'
export default {
    name: 'Vmessagedraw',
    data() {
        return {
            tableData: [],
            options_arn: [],
            options_route: [],
            options_auth: [],
			dialogFormVisible: false,

            opt_form:{
                opt_arn:'',
                opt_auth:'',
                opt_route:'',
            }
        }
    },
    mounted: function () {
        this.userid = this.$store.state.userid
        this.getOptions()
        this.getLinks()
    },
    methods: {
        //点击保存之后
        onSave() {
            this.fullscreenLoading = true;
            this.form.timestamp = new Date().getTime()
            postAtc(this.form).then(res => {
                this.fullscreenLoading = false;

                if (res.data.code == "1000") {
                    this.$notify({
                        title: '成功',
                        message: '上传成功',
                        type: 'success'
                    });
                }
            })
        },
        getOptions(){
            getLinkOptions().then(res => {
                var opts = res.data.data
                console.log(opts)
                this.options_arn = opts.arns
                this.options_route = opts.routes
                this.options_auth = opts.auths
            })
        },
        getLinks(){
            getRegistLink().then(res =>{
                    this.tableData = res.data.data
            })
        },

        handleRegist(){
            postRegLink(this.opt_form).then(res => {
                if (res.data.data === true){
                    this.$notify({
                        title: '成功',
                        message: '上传成功',
                        type: 'success'
                    });

                }

                this.getLinks()
            })
        },
    },
    components: {
    },
    couputed: {
    }
}
</script>

<style scoped>
.mainclass {
    margin: 5%;
}

.upload-content {
    margin: 40px auto;
    width: 400px;
    text-align: center;
}
</style>