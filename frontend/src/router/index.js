import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import Daftar from '../pages/Daftar.vue'
import Home from '../pages/Home.vue'
import Userlayout from '../Layouts/UserLayout.vue'
import RegisterForm from '../components/RegisterForm.vue'
import SuccessModal from '../components/SuccessModal.vue'
import History from '../components/History.vue'
import Sidebar from '../pages/admin/Sidebar.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/daftar',
    name: 'Daftar',
    component: Daftar
  },
  {
    path: '/RegisterForm',
    name: 'RegisterForm',
    component: RegisterForm
  },
  {
    path: '/success',
    name: 'SuccessModal',
    component: SuccessModal
  },
  {
    path: '/status',
    name: 'StatusPengajuan',
    component: History
  },
  {
    path: '/history',
    name: 'History',
    component: History
  },
  { 
    path: '/',
    component: Userlayout,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
