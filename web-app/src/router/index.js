import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import LoginPage from '../pages/login/login';
import CreateUserPage from '../pages/create-user/create-user';

Vue.use(Router)

export default new Router({
  routes: [
    {
        path: '/',
        name: 'Login',
        component: LoginPage
    },
    {
        path: '/create-user',
        name: '',
        component: CreateUserPage
    }
  ]
})
