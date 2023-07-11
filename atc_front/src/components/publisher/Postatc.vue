<template>
    <el-form ref="form" :model="form" label-width="100px" class="mainclass">
        <el-divider content-position="left">单报文上传</el-divider>
        <el-form-item label="报文类型" prop="option">
            <el-select v-model="form.msgtype" placeholder="请选择" style="width: 90%" @change="typeChange()">
                <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="航班号">
            <el-input v-model="form.flight" placeholder="请输入航班号" style="width: 90%;"></el-input>
        </el-form-item>
        <el-form-item label="报文">
            <el-input type="textarea" rows="4" v-model="form.content" style="width: 90%;"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSave" v-loading.fullscreen.lock="fullscreenLoading">上传</el-button>
        </el-form-item>

        <el-divider content-position="left">多报文上传</el-divider>
        <el-form-item>
            <div class="upload-content">
                <el-upload class="upload-demo" drag action="https://jsonplaceholder.typicode.com/posts/" :on-success="handleChange" :file-list="fileList">
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                    <div class="el-upload__tip" slot="tip">只能上传xlsx文件，且不超过500kb</div>
                </el-upload>
            </div>
            <el-table v-if="tableHead.length" :data="tableData[0]" style="width: 100%">
                <el-table-column v-for="(data, key) in tableHead" :prop="data" :label="data" :key="key" width="180">
                </el-table-column>
            </el-table>
        </el-form-item>
    </el-form>
</template>

<script>
import { postAtc, postMultiAtc } from '../../api/axios'
import { read, utils } from "xlsx"; // 注意处理方法引入方式
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
            messages: [],
            fullscreenLoading: false,
            options: [{
                value: 'ADS-B',
                label: 'ADS-B'
            }, {
                value: 'ACARS',
                label: 'ACARS'
            }],
            userid: '',
            fileList: [],
            tableData: [],
            tableHead: [],
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
        onSaveMulti(){
            this.fullscreenLoading = true;
            var timestamp = new Date().getTime()
            postMultiAtc({userid: this.userid, msgs: JSON.stringify(this.tableData[0]), timestamp}).then(res => {
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
        handleChange(res, file, fileList) {
            // 将文件放入
            for (let i = 0; i < fileList.length; i++) {
                if (file.name != fileList[i].name) {
                    this.fileList.push({
                        name: file.name,
                        url: "",
                        uid: file.uid
                    });
                }
            }
            const files = { 0: file };
            this.readExcel(files);
        },
        readExcel(file) {
            const fileReader = new FileReader();
            fileReader.onload = ev => {
                this.tableData.length=0
                try {
                    const data = ev.target.result;
                    const workbook = read(data, { type: "binary" });
                    const params = [];
                    // 取对应表生成json表格内容
                    workbook.SheetNames.forEach(item => {
                        this.tableData.push(utils.sheet_to_json(workbook.Sheets[item]));
                    });
                    // 该算法仅针对表头无合并的情况
                    if (this.tableData.length > 0) {
                        // 获取excel中第一个表格数据tableData[0][0]，并且将表头提取出来
                        for (const key in this.tableData[0][0]) {
                            this.tableHead.push(key);
                        }
                    }
                    this.onSaveMulti()
                    // 重写数据
                } catch (e) {
                    console.log("error:" + e);
                    return false;
                }
            };
            fileReader.readAsBinaryString(file[0].raw);
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

.upload-content {
    margin: 40px auto;
    width: 400px;
    text-align: center;
}
</style>