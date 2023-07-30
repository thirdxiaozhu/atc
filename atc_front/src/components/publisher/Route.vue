<template>
        <el-container>
        <el-header>
            <el-button type="primary" @click="addRoute()" v-loading.fullscreen.lock="fullscreenLoading">注册航线</el-button>
        </el-header>
        <el-main>
            <el-table :data="tableData" style="width: 100%" v-loading="loading">
                <el-table-column type="expand">
                    <template slot-scope="props">
                        <el-form label-position="left" inline class="demo-table-expand">
                            <el-form-item label="序号">
                                <span>{{ props.row.ID }}</span>
                            </el-form-item>
                            <el-form-item label="航班号">
                                <span>{{ props.row.Flight_num }}</span>
                            </el-form-item>
                            <el-form-item label="始发地">
                                <span>{{ props.row.Origin }}</span>
                            </el-form-item>
                            <el-form-item label="目的地">
                                <span>{{ props.row.Destination }}</span>
                            </el-form-item>
                        </el-form>
                    </template>
                </el-table-column>
                <el-table-column label="序号" prop="ID">
                </el-table-column>
                <el-table-column label="航班号" prop="Flight_num">
                </el-table-column>
                <el-table-column label="始发地" prop="Origin">
                </el-table-column>
                <el-table-column label="目的地" prop="Destination">
                </el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-button size="mini" @click="dialogFormVisible = true; current_row = scope.row"
                            :disabled="scope.row.disabled">编辑</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-main>
    </el-container>
</template>

<script>
import { getRoutes } from '../../api/axios'
export default {
    name: 'Vmessagedraw',
    data() {
        return {
            form: {
                userid: this.$store.state.userid,
                msgtype: "ADS-B",
                content: '',
                timestamp: '',
                flight: '',
            },
            tableData: [],
        }
    },
    mounted: function () {
        this.userid = this.$store.state.userid
        this.getAllRoutes()
        
    },
    methods: {
        //点击保存之后
        getAllRoutes() {
            getRoutes(this.form).then(res => {
                console.log(res.data)
                this.tableData = res.data.data
            })
        },
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