<template>
	<el-container>
		<el-aside width="200px">
			<!-- 
			<el-tree :data="treeData" :props="defaultProps" ref="tree" node-key="id" @node-click="handleNodeClick"
				@node-expand="expandClick" style="margin-top: 1px;"> -->
			<el-tree :data="treeData" :props="defaultProps" ref="tree" node-key="id" @node-click="handleNodeClick"
				style="margin-top: 1px;">
				<template #default="{ node, data }">
					<span class="custom-tree-node">
						<i v-if="data.loading" class="el-icon-loading"></i>
						<span class="node-label">{{ node.label }}</span>
					</span>
				</template>
			</el-tree>
		</el-aside>
		<el-main>
			<h3 align="left">数字证书</h3>
			<el-input type="textarea" :rows="10" placeholder="公钥证书" v-model="certtext"
				style="margin-top: 10px; width: 100%;">
			</el-input>
			<el-divider></el-divider>
			<el-row>
				<el-col :span="12">
					<h3 align="left">Issuer</h3>
					<el-table :data="issueData" style="width: 50%;">
						<el-table-column prop="key" label="对象">
						</el-table-column>
						<el-table-column prop="value" label="内容">
						</el-table-column>
					</el-table>
				</el-col>
				<el-col  :span="12">
					<h3 align="left">Algorithm</h3>
					<el-table :data="algoData" style="width: 50%;">
						<el-table-column prop="key" label="对象">
						</el-table-column>
						<el-table-column prop="value" label="内容">
						</el-table-column>
					</el-table>
				</el-col>
			</el-row>
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
import { getCompanyOptions, getPublisherOptions, getCertificate } from '../../api/axios'
export default {
	name: "tree1",
	data() {
		return {
			treeData: [{
				value: "0",
				level: 0,
				label: '选择用户',
				isLeaf: false,
				children: []
			}],
			defaultProps: {
				children: 'children',
				label: 'label',
			},
			certtext: "",
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
			algoData:[{
				key: "签名算法",
				value: '-',
			},{
				key: "公钥算法",
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
		this.treeDataDeal(this.treeData)
	},
	methods: {
		// 树数据处理
		treeDataDeal(treeData) {
			treeData.forEach(item => {
				if (!item.children || item.children.length <= 0) {
					// 设置为非叶子节点
					this.$refs.tree.getNode(item.value).isLeaf = false
				} else {
					// 默认展开子节点
					this.$refs.tree.getNode(item.value).expanded = true
				}
				this.treeDataDeal(item.children)
			})
		},
		// 点击节点
		handleNodeClick(data) {
			this.getData(data)
		},
		//// 节点展开
		//expandClick(data) {
		//	this.getData(data)
		//},
		getData(data) {
			if (data.level === 0 && data.children.length < 1) {
				// 开启loading
				this.$set(data, 'loading', true)

				getCompanyOptions({ role: -1 }).then(res => {
					for (var i = 0; i < res.data.data.length; i++) {
						res.data.data[i].level = 1
						res.data.data[i].children = []
					}
					data.children = res.data.data
					this.$set(data, 'loading', false)
				})

			} else if (data.level === 1 && data.children.length < 1) {
				getPublisherOptions({ "company": data.value }).then(res => {
					for (var i = 0; i < res.data.data.length; i++) {
						res.data.data[i].level = 2
						res.data.data[i].isLeaf = true
					}
					data.children = res.data.data
					this.$set(data, 'loading', false)
				})
			} else if (data.level === 2) {
				console.log(data.value)
				getCertificate({ "user": data.value }).then(res => {
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

					this.algoData[0].value = this.cert_ret.algomap.signature_algorithm
					this.algoData[1].value = this.cert_ret.algomap.publickey_algorithm
				})
			}
		},
	}
}
</script>
