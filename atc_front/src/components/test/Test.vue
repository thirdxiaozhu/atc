<template>
	<div class="test">
		<div class="create">
			<h1>添加用户</h1><br />
			<el-input v-model="name" style="width: 90%;"></el-input>
			<el-input v-model="password" style="width: 90%;"></el-input> <br/>
			<el-button type="success" @click="onCreate" >添加</el-button> <br/>
			<i>{{ create_res }}</i>
		</div>
		<div class="search">
			<h1>查询用户</h1><br />
			<el-input v-model="search_id" style="width: 90%;"></el-input>
			<el-button type="success" @click="onSave" >查询</el-button> <br/>
			<i>{{ name }}</i>
		</div>
	</div>
</template>

<script>
import { getResult, postCreate } from '../../api/axios'
export default {
	name: 'test',
	data() {
		return {
			msg: 'Welcome to Your Vue.js App',
			name: '',
			password: '',
			search_id: '',
			create_res: '',
		}
	},
	mounted: function () {
		//this.getUser()
	},
	methods: {
		getUser() {
			getResult().then(ret => {
				console.log(ret)
			})
		},
		onCreate(){
			var that = this
			const param = {
				username: this.name,
				password: this.password
			}

			postCreate(param).then(res => {
				console.log(res)
			})
		},
		onSave() {
			console.log(this.search_id)
			var that = this
			const param = {
				id: this.search_id
			}
			getResult(param).then(ret =>{
				console.log(ret.data)
				that.name = ret.data
			})
			
		}
	}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
	font-weight: normal;
}

ul {
	list-style-type: none;
	padding: 0;
}

li {
	display: inline-block;
	margin: 0 10px;
}

a {
	color: #42b983;
}
</style>
