import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import Home from '../view/Home.vue'
import Main from '../view/Main.vue'
import CatM from '../view/CatM.vue'
import UserM from '../view/UserM.vue'
import AdaptM from '../view/AdaptM.vue'
import AdaptU from '../view/AdaptU.vue'
import PersonalFile from '../view/PersonalFile.vue'
import Login from '../view/Login.vue'
import Register from '../view/Register.vue'

const routes = [
  {
    path: '/',
    redirect: '/home',
    component: Main,
    children: [
      {
        path: 'home',
        component: Home
      },
      {
        path: 'catm',
        component: CatM
      },
      {
        path: 'userm',
        component: UserM
      },
      {
        path: 'adaptm',
        component: AdaptM
      },
      {
        path: 'adaptu',
        component: AdaptU
      },
      {
        path: 'personalfile',
        component: PersonalFile
      },
    ]
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: '/register',
    component: Register
  }

]

const router = createRouter({
  // history: createWebHistory(),
  history: createWebHashHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  if (localStorage.getItem('id')) {
    next()
  } else {
    if (to.path === '/login' || to.path === '/register') {
      next()
    }
    next('/login')
  }
})

export default router;

