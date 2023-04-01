<template>
	<el-container>
		<el-header style="text-align: right; font-size: 12px">
			<el-select v-model="company" multiple collapse-tags style="margin-right: 50px;" placeholder="请选择发布方"
				@change="changeCompanyOption">
				<el-option v-for="item in company_options" :key="item.value" :label="item.label" :value="item.value">
				</el-option>
			</el-select>
			<el-select v-model="publisher" multiple collapse-tags style="margin-right: 50px;" placeholder="请选择发布者">
				<el-option v-for="item in publisher_options" :key="item.value" :label="item.label" :value="item.value">
				</el-option>
			</el-select>

			<el-date-picker v-model="value2" type="datetimerange" :picker-options="pickerOptions" range-separator="至"
				start-placeholder="开始日期" end-placeholder="结束日期" align="right" @change="saveTimeSpan"
				style="margin-right: 50px;">
			</el-date-picker>

			<el-button type="primary" @click="onSave" style="margin-right: 50px;">查询</el-button>
		</el-header>

		<el-main>
			<el-table :data="tableData" style="width: 100%" v-loading="loading">
				<el-table-column type="expand">
					<template slot-scope="props">
						<el-form label-position="left" inline class="demo-table-expand">
							<el-form-item label="发布时间">
								<span>{{ props.row.Time }}</span>
							</el-form-item>
							<el-form-item label="发布ID">
								<span>{{ props.row.ID }}</span>
							</el-form-item>
							<el-form-item label="发布方">
								<span>{{ props.row.Company }}</span>
							</el-form-item>
							<el-form-item label="发布者">
								<span>{{ props.row.Publisher }}</span>
							</el-form-item>
							<el-form-item label="签名值">
								<span>{{ props.row.Signature }}</span>
							</el-form-item>
							<el-form-item label="IPFS地址">
								<span>{{ props.row.Address }}</span>
							</el-form-item>
							<el-form-item label="报文类型">
								<span>{{ props.row.Type }}</span>
							</el-form-item>
							<el-form-item label="报文内容">
								<span>{{ props.row.Content }}</span>
							</el-form-item>
						</el-form>
					</template>
				</el-table-column>
				<el-table-column label="发布时间" prop="Time" :filters="filtertags" :filter-method="filterTag"
					filter-placement="bottom-end">
				</el-table-column>
				<el-table-column label="发布方" prop="Company">
				</el-table-column>
				<el-table-column label="发布者" prop="Publisher">
				</el-table-column>
				<el-table-column label="报文类型" prop="Type">
				</el-table-column>
				<el-table-column label="报文内容" prop="Content">
				</el-table-column>
				<el-table-column label="操作">
					<template slot-scope="scope">
						<!--
						<el-button size="mini" @click="dialogFormVisible=true; current_row=scope.row "
							:disabled="scope.row.disabled">编辑</el-button>-->
						<el-button size="mini" type="info" @click="handleHistory(scope.$index, scope.row)">历史</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-main>
		<el-dialog title="历史记录" :visible.sync="dialogTableVisible">
			<el-table :data="gridData">
				<el-table-column label="发布时间" prop="Atc.Time" :filters="filtertags" :filter-method="filterTag"
					filter-placement="bottom-end">
				</el-table-column>
				<el-table-column label="报文内容" prop="Atc.Content">
				</el-table-column>
				<el-table-column label="IPFS地址" prop="Atc.Address">
				</el-table-column>
			</el-table>
		</el-dialog>
		<el-dialog title="编辑" :visible.sync="dialogFormVisible">
			<el-form :model="editform">
				<el-form-item label="新报文">
					<el-input v-model="editform.content" autocomplete="off"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button type="primary" @click="handleEdit()">确定</el-button>
				<el-button  @click="dialogFormVisible = false">取消</el-button>
				<el-button type="danger" @click="handleDelete()">删除</el-button>
			</div>
		</el-dialog>
	</el-container>
</template>
  
<style>
.demo-table-expand {
	font-size: 0;
}

.demo-table-expand label {
	width: 90px;
	color: #99a9bf;
}

.demo-table-expand .el-form-item {
	margin-right: 0;
	margin-bottom: 0;
	width: 50%;
}

.el-header {
	background-color: #FFF;
	color: #333;
	line-height: 60px;
	padding: 0px;
}


.el-main {
	padding: 0px;
}
</style>
  
<script>
import { getAtc, getCompanyOptions, getPublisherOptions, postEdit, postDelete } from '../../api/axios'
export default {
	data() {
		return {
			userid: "",
			form: {
				userid: this.$store.state.userid,
				publisher: "",
				company: "",
				starttime: "",
				endtime: "",
			},
			tableData: [],
			pickerOptions: {
				shortcuts: [{
					text: '最近一周',
					onClick(picker) {
						this.endtime = new Date();
						this.starttime = new Date();
						this.starttime.setTime(this.starttime.getTime() - 3600 * 1000 * 24 * 7);
						picker.$emit('pick', [this.starttime, this.endtime]);
					}
				}, {
					text: '最近一个月',
					onClick(picker) {
						const end = new Date();
						const start = new Date();
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
						picker.$emit('pick', [start, end]);
					}
				}, {
					text: '最近三个月',
					onClick(picker) {
						const end = new Date();
						const start = new Date();
						start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
						picker.$emit('pick', [start, end]);
					}
				}]
			},
			value2: "",
			filtertags: [{ text: '12987122', value: '12987122' }, { text: '12987133', value: '12987133' }],
			publisher_options: [],
			publisher: [],
			company_options: [],
			company: [],
			dialogTableVisible: false,
			dialogFormVisible: false,
			gridData: [],
			editform: {
				userid: '',
				ID: '',
				content: '',
				timestamp: '',
			},
			current_row: 0,
			loading: false,
		}
	},
	mounted: function () {
		this.userid = this.$store.state.userid
		//this.form.publisher = this.userid
		this.getCompanyOptions()
		this.loading = true
		this.getAllAtc()
		this.loading = false
	},
	methods: {
		filterTag(value, row) {
			return row.id === value;
		},

		getAllAtc() {
			this.form.company = this.company.toString()
			this.form.publisher = this.publisher.toString()
			getAtc(this.form).then(res => {
				console.log(res.data.data)
				var len = res.data.data.length
				for (var i = 0; i < len; i++) {
					var msg_date = new Date().setTime(res.data.data[i].Time)
					res.data.data[i].Time = this.getdate(msg_date)

					if (res.data.data[i].Publisher == this.userid) {
						res.data.data[i].disabled = false
					} else {
						res.data.data[i].disabled = true
					}
				}
				this.tableData = res.data.data
				this.loading = false
			})
		},

		onSave() {
			this.loading = true
			this.getAllAtc()
			this.loading = false
		},
		saveTimeSpan() {
			this.form.starttime = this.value2[0].getTime()
			this.form.endtime = this.value2[1].getTime()
		},

		getdate(timeStamp) {
			if (timeStamp) {
				var now = new Date(timeStamp);
			} else {
				var now = new Date();
			}
			var Y = now.getFullYear();
			var M = now.getMonth() + 1;
			var D = now.getDate();
			var m = M < 10 ? "0" + M : M;
			var d = D < 10 ? "0" + D : D;
			return Y + "-" + m + "-" + d + " " + now.toTimeString().substr(0, 8);
		},

		getCompanyOptions() {
			getCompanyOptions().then(res => {
				this.company_options = res.data.data
			})
		},

		changeCompanyOption() {
			this.form.publisher = []
			getPublisherOptions({ "company": this.company.toString() }).then(res => {
				this.publisher_options = res.data.data
			})
		},
		handleEdit() {
			this.dialogFormVisible = false
            this.editform.timestamp = new Date().getTime()
			this.editform.ID = this.current_row.ID
			this.editform.userid = this.userid
			this.loading = true
			postEdit(this.editform).then(res =>{
				if(res.data.code === 1000){
					this.getAllAtc()

                    this.$notify({
                        title: '成功',
                        message: '修改成功',
                        type: 'success'
                    });
				}
			})
		},
		handleHistory(index, row) {
			this.dialogTableVisible = true
			this.gridData = row.Historys
			for (var i = 0; i < row.Historys.length; i++) {
				var msg_date = new Date().setTime(row.Historys[i].Atc.Time)
				row.Historys[i].Atc.Time = this.getdate(msg_date)
			}
		},
		handleDelete(index, row) {
			this.dialogFormVisible = false
			this.loading = true
			postDelete({userid: this.userid, ID: this.current_row.ID}).then(res =>{
                if (res.data.code == "1000") {
					this.getAllAtc()
                    this.$notify({
                        title: '成功',
                        message: '删除成功',
                        type: 'success'
                    });
                }
			})
		},

	}
}
</script>