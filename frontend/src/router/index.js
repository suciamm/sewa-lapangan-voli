import { createRouter, createWebHistory } from 'vue-router'
import { h } from 'vue'
import Home from '@/views/Home.vue'
import Courts from '@/views/Courts.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import OwnerDashboard from '@/views/OwnerDashboard.vue'
import PenyewaDashboard from '@/views/PenyewaDashboard.vue'
import SuperuserDashboard from '@/views/SuperuserDashboard.vue'
import OwnerCourts from '@/views/OwnerCourts.vue'
import OwnerBookings from '@/views/OwnerBookings.vue'
import PenyewaBookings from '@/views/PenyewaBookings.vue'
import Notifications from '@/views/Notifications.vue'
import Layout from '@/components/Layout.vue'
import { useAuthStore } from '@/stores/auth'

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
    path: '/owner/dashboard',
    name: 'OwnerDashboard',
    component: OwnerDashboard,
    meta: { requiresAuth: true, role: 'owner' }
  },
  {
    path: '/owner/courts',
    name: 'OwnerCourts',
    component: OwnerCourts,
    meta: { requiresAuth: true, role: 'owner' }
  },
  {
    path: '/owner/bookings',
    name: 'OwnerBookings',
    component: OwnerBookings,
    meta: { requiresAuth: true, role: 'owner' }
  },
  {
    path: '/owner/report',
    name: 'OwnerReport',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Report Rating & Pendapatan'),
            h('p', 'Halaman ini akan menampilkan laporan rating lapangan dan pendapatan.')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'owner' }
  },
  {
    path: '/penyewa/dashboard',
    name: 'PenyewaDashboard',
    component: PenyewaDashboard,
    meta: { requiresAuth: true, role: 'penyewa' }
  },
  {
    path: '/penyewa/booking-aktif',
    name: 'PenyewaBookingAktif',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Booking Aktif'),
            h('p', 'Halaman ini menampilkan booking yang sedang aktif.')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'penyewa' }
  },
  {
    path: '/penyewa/beri-rating',
    name: 'PenyewaBeriRating',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Beri Rating'),
            h('p', 'Halaman ini untuk memberikan rating setelah selesai menyewa.')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'penyewa' }
  },
  {
    path: '/notifications',
    name: 'Notifications',
    component: Notifications,
    meta: { requiresAuth: true }
  },
  {
    path: '/superuser/dashboard',
    name: 'SuperuserDashboard',
    component: SuperuserDashboard,
    meta: { requiresAuth: true, role: 'superuser' }
  },
  {
    path: '/superuser/kelola-owner',
    name: 'SuperuserKelolaOwner',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Kelola Owner'),
            h('p', 'Halaman ini untuk mengelola owner (approve/reject).')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'superuser' }
  },
  {
    path: '/superuser/data-lapangan',
    name: 'SuperuserDataLapangan',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Data Lapangan'),
            h('p', 'Halaman ini menampilkan semua data lapangan.')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'superuser' }
  },
  {
    path: '/superuser/report',
    name: 'SuperuserReport',
    component: {
      render() {
        return h(Layout, null, {
          default: () => h('div', { class: 'page' }, [
            h('h1', 'Report Rating & Booking'),
            h('p', 'Halaman ini menampilkan laporan rating dan booking.')
          ])
        })
      }
    },
    meta: { requiresAuth: true, role: 'superuser' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
    return
  }

  if (to.meta.role && authStore.user?.role !== to.meta.role) {
    // Redirect to appropriate dashboard based on role
    if (authStore.user?.role === 'owner') {
      next('/owner/dashboard')
    } else if (authStore.user?.role === 'penyewa') {
      next('/penyewa/dashboard')
    } else if (authStore.user?.role === 'superuser') {
      next('/superuser/dashboard')
    } else {
      next('/')
    }
    return
  }

  next()
})

export default router
