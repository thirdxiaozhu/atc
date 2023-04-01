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

import AcquirerView from '@/view/AcquirerView'
import A_Main from '@/components/acquirer/Main'
import A_Checkatc from '@/components/acquirer/Checkatc'
import A_Certificate from '@/components/acquirer/Certificate'
import A_Prikey  from '@/components/acquirer/Prikey'

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
		},
		{
			path: '/acquirer',
			name: 'Acquirer',
			component: AcquirerView,
			children: [
				{
					path: "/",
					name: "A_Main",
					component: A_Main,
				},
				{
					path: "main",
					name: "A_Main",
					component: A_Main,
				},
				{
					path: "checkatc",
					name: "A_CheckAtc",
					component: A_Checkatc,
				},
				{
					path: "certificate",
					name: "A_Certificate",
					component: A_Certificate,
				},
				{
					path: "prikey",
					name: "A_Prikey",
					component: A_Prikey,
				}
			]
		}
	]
})
