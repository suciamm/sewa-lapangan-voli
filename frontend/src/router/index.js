import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Courts from '@/views/Courts.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import Bookings from '@/views/Bookings.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/courts',
    name: 'Courts',
    component: Courts
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/bookings',
    name: 'Bookings',
    component: Bookings
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
