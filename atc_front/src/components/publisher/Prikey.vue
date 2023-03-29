<template>
	<el-container>
		<el-main>
			<h3 align="left">数字证书</h3>
			<el-input type="textarea" :rows="10" placeholder="公钥证书" v-model="certtext"
				style="margin-top: 10px; width: 45%;">
			</el-input>
			<el-input type="textarea" :rows="10" placeholder="私钥" v-model="privtext"
				style="margin-top: 10px; width: 45%; margin-left: 30px;">
			</el-input>
			<el-divider></el-divider>
			<h3 align="left">Issuer</h3>
			<el-table :data="issueData" style="width: 50%;">
				<el-table-column prop="key" label="对象">
				</el-table-column>
				<el-table-column prop="value" label="内容">
				</el-table-column>
			</el-table>
			<el-divider></el-divider>
			<h3 align="left">公钥</h3>
			<el-table :data="keyData" style="width: 100%;">
				<el-table-column prop="key" label="" width="80">
				</el-table-column>
				<el-table-column prop="value" label="值">
					<template slot-scope="scope">
						<el-tag :type="scope.row.key === 'X' ? 'warning' : 'success'" disable-transitions>{{ scope.row.value
						}}</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="length" label="长度">
				</el-table-column>
			</el-table>
		</el-main>
	</el-container>
</template>

<script>
import { getCertificate, getPrivateKey } from '../../api/axios'
export default {
	name: "tree1",
	data() {
		return {
			userid: "",
			certtext: "",
			privtext:"",
			issueData: [{
				key: 'C',
				value: '-',
			}, {
				key: 'CN',
				value: '-',
			}, {
				key: 'L',
				value: '-',
			}, {
				key: 'O',
				value: '-',
			}, {
				key: 'ST',
				value: '-',
			}],
			keyData: [{
				key: 'X',
				value: '-',
				length: '',
			}, {
				key: 'Y',
				value: '-',
				length: '',
			}]
		}
	},
	mounted() {
		// this.$refs.tree.getNode(1).expanded = true
		// this.$refs.tree.getNode(11).expanded = true
		// this.$refs.tree.getNode(111).isLeaf = false
        this.userid = this.$store.state.userid;
		this.getCert()
		this.getPriv()

	},
	methods: {
		getCert() {
			getCertificate({ "user": this.userid }).then(res => {
				this.cert_ret = res.data.data
				this.certtext = this.cert_ret.content
				this.issueData[0].value = this.cert_ret.issmap.C
				this.issueData[1].value = this.cert_ret.issmap.CN
				this.issueData[2].value = this.cert_ret.issmap.L
				this.issueData[3].value = this.cert_ret.issmap.O
				this.issueData[4].value = this.cert_ret.issmap.ST

				this.keyData[0].value = this.cert_ret.x
				this.keyData[0].length = this.cert_ret.xlength
				this.keyData[1].value = this.cert_ret.y
				this.keyData[1].length = this.cert_ret.ylength
			})
		},
		getPriv(){
			getPrivateKey({ "user": this.userid }).then(res => {
				this.privtext = res.data.data
			})
		}
	}
}
</script>
