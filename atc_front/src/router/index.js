import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Testing from '@/components/test/Test'
import TestView from '@/view/TestView'
import Login from '@/components/Login'
import SignUp from '@/components/Signup'

import PublisherView from '@/view/PublisherView'
import Main from '@/components/publisher/Main'
import Postatc from '@/components/publisher/Postatc'
import Checkatc from '@/components/publisher/Checkatc'
import Certificate from '@/components/publisher/Certificate'
import Prikey  from '@/components/publisher/Prikey'

Vue.use(Router)

export default new Router({
	mode: 'history',
	routes: [
		{
			path: '/test',
			name: 'HelloWorld',
			component: TestView,
			children: [
				{
					path: "/",
					name: "test",
					component: HelloWorld
				},
				{
					path: 'append',
					name: 'TestAppend',
					component: Testing
				}
			]
		},
		{
			path: '/',
			redirect: 'login'
		},
		{
			path: '/signup',
			name: 'SignUp',
			component: SignUp,
		},
		{
			path: '/login',
			name: 'Login',
			component: Login,
		},
		{
			path: '/publisher',
			name: 'Publisher',
			component: PublisherView,
			children: [
				{
					path: "/",
					name: "Main",
					component: Main,
				},
				{
					path: "main",
					name: "Main",
					component: Main,
				},
				{
					path: "postatc",
					name: "PostAtc",
					component: Postatc,
				},
				{
					path: "checkatc",
					name: "CheckAtc",
					component: Checkatc,
				},
				{
					path: "certificate",
					name: "Certificate",
					component: Certificate,
				},
				{
					path: "prikey",
					name: "Prikey",
					component: Prikey,
				}
			]
		}
	]
})
