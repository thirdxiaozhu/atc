<template>
    <el-form ref="form" :model="form" label-width="100px" class="mainclass">
        <el-form-item label="收信人类型" prop="option">
            <el-select v-model="form.msgtype" placeholder="请选择" style="width: 90%" @change="typeChange()">
                <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="报文">
            <el-input type="textarea" rows="4" v-model="form.content" style="width: 90%;"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSave" v-loading.fullscreen.lock="fullscreenLoading">发送</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
import { postAtc } from '../../api/axios'
export default {
    name: 'Vmessagedraw',
    data() {
        return {
            form: {
                userid: this.$store.state.userid,
                msgtype: "ADS-B",
                content: '',
                timestamp: '',
            },
            fullscreenLoading: false,
            //rules: {
            //    option: [
            //        { required: true, message: '请选择收信人类型', trigger: 'change' }
            //    ],
            //    title: [
            //        { required: true, message: '请输入标题', trigger: 'blur' },
            //        { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
            //    ],
            //    towho: [
            //        { required: true, message: '请输入收信人', trigger: 'blur' },
            //    ],
            //},
            options: [{
                value: 'ADS-B',
                label: 'ADS-B'
            }, {
                value: 'Acars',
                label: 'Acars'
            }],
            userid: '',
        }
    },
    mounted: function () {
        this.userid = this.$store.state.userid;
    },
    methods: {
        //点击保存之后
        onSave() {
            this.fullscreenLoading = true;
            this.form.timestamp = new Date().getTime()
            console.log(this.form)
            postAtc(this.form).then(res => {
                console.log(res);
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
        typeChange() {
            console.log(this.form.type)
        }
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
</style>