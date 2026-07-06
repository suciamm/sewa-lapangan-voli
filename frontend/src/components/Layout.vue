<template>
  <div class="layout">
    <!-- Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-header">
        <div class="logo">
          <span class="logo-icon">🏐</span>
          <span class="logo-text">Sewa Voli</span>
        </div>
      </div>

      <nav class="sidebar-nav">
        <!-- Owner Navigation -->
        <template v-if="userRole === 'owner'">
          <div class="nav-section">
            <h4 class="nav-section-title">Menu Utama</h4>
            <router-link to="/owner/dashboard" class="nav-link" active-class="active">
              <span class="nav-icon">📊</span>
              <span class="nav-text">Dashboard</span>
            </router-link>
            <router-link to="/owner/courts" class="nav-link" active-class="active">
              <span class="nav-icon">🏟️</span>
              <span class="nav-text">Kelola Lapangan</span>
            </router-link>
          </div>
          <div class="nav-section">
            <h4 class="nav-section-title">Laporan</h4>
            <router-link to="/owner/bookings" class="nav-link" active-class="active">
              <span class="nav-icon">📋</span>
              <span class="nav-text">Booking</span>
            </router-link>
            <router-link to="/owner/report" class="nav-link" active-class="active">
              <span class="nav-icon">📈</span>
              <span class="nav-text">Rating & Pendapatan</span>
            </router-link>
          </div>
        </template>

        <!-- Penyewa Navigation -->
        <template v-else-if="userRole === 'penyewa'">
          <div class="nav-section">
            <h4 class="nav-section-title">Menu Utama</h4>
            <router-link to="/penyewa/dashboard" class="nav-link" active-class="active">
              <span class="nav-icon">📊</span>
              <span class="nav-text">Dashboard</span>
            </router-link>
            <router-link to="/courts" class="nav-link" active-class="active">
              <span class="nav-icon">🔍</span>
              <span class="nav-text">Cari Lapangan</span>
            </router-link>
          </div>
          <div class="nav-section">
            <h4 class="nav-section-title">Aktivitas</h4>
            <router-link to="/penyewa/bookings" class="nav-link" active-class="active">
              <span class="nav-icon">✅</span>
              <span class="nav-text">Booking Aktif</span>
            </router-link>
            <router-link to="/penyewa/ratings" class="nav-link" active-class="active">
              <span class="nav-icon">⭐</span>
              <span class="nav-text">Beri Rating</span>
            </router-link>
            <router-link to="/notifications" class="nav-link" active-class="active">
              <span class="nav-icon">🔔</span>
              <span class="nav-text">Notifikasi</span>
            </router-link>
          </div>
        </template>

        <!-- Superuser Navigation -->
        <template v-else-if="userRole === 'superuser'">
          <div class="nav-section">
            <h4 class="nav-section-title">Dashboard</h4>
            <router-link to="/superuser/dashboard" class="nav-link" active-class="active">
              <span class="nav-icon">📊</span>
              <span class="nav-text">Ringkasan Sistem</span>
            </router-link>
          </div>
          <div class="nav-section">
            <h4 class="nav-section-title">Manajemen</h4>
            <router-link to="/superuser/kelola-owner" class="nav-link" active-class="active">
              <span class="nav-icon">👥</span>
              <span class="nav-text">Kelola Owner</span>
            </router-link>
            <router-link to="/superuser/data-lapangan" class="nav-link" active-class="active">
              <span class="nav-icon">📋</span>
              <span class="nav-text">Data Lapangan</span>
            </router-link>
          </div>
          <div class="nav-section">
            <h4 class="nav-section-title">Laporan</h4>
            <router-link to="/superuser/report" class="nav-link" active-class="active">
              <span class="nav-icon">📈</span>
              <span class="nav-text">Report & Analitik</span>
            </router-link>
          </div>
        </template>
      </nav>

      <div class="sidebar-footer">
        <div class="user-profile">
          <div class="user-avatar">{{ getInitials(userName) }}</div>
          <div class="user-details">
            <p class="user-name">{{ userName }}</p>
            <p class="user-role">{{ getRoleLabel(userRole) }}</p>
          </div>
        </div>
        <button @click="logout" class="btn-logout">
          <span class="logout-icon">🚪</span>
          <span>Keluar</span>
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="main-content">
      <header class="content-header">
        <div class="header-left">
          <h1 class="page-title">{{ pageTitle }}</h1>
        </div>
        <div class="header-right">
          <div class="user-welcome">
            Selamat datang, <strong>{{ userName }}</strong>
          </div>
        </div>
      </header>
      <div class="content-body">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const userRole = computed(() => authStore.user?.role)
const userName = computed(() => authStore.user?.name || 'User')

const pageTitle = computed(() => {
  const titles = {
    '/owner/dashboard': 'Dashboard Pemilik Lapangan',
    '/owner/courts': 'Kelola Lapangan',
    '/owner/bookings': 'Booking Masuk',
    '/owner/report': 'Report Rating & Pendapatan',
    '/penyewa/dashboard': 'Dashboard Penyewa',
    '/courts': 'Cari Lapangan',
    '/penyewa/bookings': 'Booking Aktif',
    '/penyewa/ratings': 'Beri Rating',
    '/notifications': 'Notifikasi',
    '/superuser/dashboard': 'Dashboard Superuser',
    '/superuser/kelola-owner': 'Kelola Owner',
    '/superuser/data-lapangan': 'Data Lapangan',
    '/superuser/report': 'Report & Analitik',
  }
  return titles[route.path] || 'Dashboard'
})

const getInitials = (name) => {
  if (!name) return '?'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

const getRoleLabel = (role) => {
  const labels = {
    'owner': 'Pemilik Lapangan',
    'penyewa': 'Penyewa',
    'superuser': 'Admin Sistem'
  }
  return labels[role] || role
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
/* Layout */
.layout {
  display: flex;
  min-height: 100vh;
}

/* Sidebar */
.sidebar {
  width: 280px;
  background: linear-gradient(180deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  display: flex;
  flex-direction: column;
  box-shadow: 4px 0 12px rgba(0,0,0,0.15);
  position: fixed;
  height: 100vh;
  overflow-y: auto;
  z-index: 1000;
}

/* Sidebar Header */
.sidebar-header {
  padding: 1.5rem;
  border-bottom: 1px solid rgba(255,255,255,0.1);
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.5rem;
  font-weight: 800;
}

.logo-icon {
  font-size: 2rem;
}

.logo-text {
  letter-spacing: 0.5px;
}

/* Navigation */
.sidebar-nav {
  flex: 1;
  padding: 1.5rem 0;
  overflow-y: auto;
}

.nav-section {
  margin-bottom: 2rem;
}

.nav-section-title {
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  color: rgba(255, 255, 255, 0.5);
  padding: 0.5rem 1.5rem 1rem;
  letter-spacing: 1px;
  margin: 0;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: rgba(255, 255, 255, 0.85);
  text-decoration: none;
  padding: 0.75rem 1.5rem;
  font-weight: 500;
  font-size: 0.95rem;
  transition: all 0.2s ease;
  border-left: 4px solid transparent;
  margin: 0.25rem 0;
}

.nav-icon {
  font-size: 1.2rem;
  min-width: 1.5rem;
  text-align: center;
}

.nav-link:hover {
  color: white;
  background-color: rgba(255, 255, 255, 0.08);
  border-left-color: #ff6b35;
}

.nav-link.active {
  color: white;
  background-color: rgba(255, 107, 53, 0.2);
  border-left-color: #ff6b35;
  font-weight: 700;
}

/* User Profile Section */
.sidebar-footer {
  padding: 1.5rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.user-profile {
  display: flex;
  gap: 0.75rem;
  align-items: center;
  padding: 0.75rem;
  background-color: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.9rem;
  flex-shrink: 0;
}

.user-details {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: 0.9rem;
  font-weight: 700;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-role {
  font-size: 0.75rem;
  color: rgba(255, 255, 255, 0.7);
  margin: 0.25rem 0 0 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Logout Button */
.btn-logout {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  width: 100%;
  background-color: rgba(230, 57, 70, 0.15);
  color: #ff8a8a;
  border: 1px solid rgba(230, 57, 70, 0.3);
  padding: 0.75rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.btn-logout:hover {
  background-color: #e63946;
  color: white;
  border-color: transparent;
  transform: translateY(-2px);
}

/* Main Content */
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #f8fafc;
  margin-left: 280px;
}

/* Content Header */
.content-header {
  background-color: white;
  padding: 1.5rem 2rem;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.header-left {
  flex: 1;
}

.page-title {
  font-size: 1.75rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0;
  letter-spacing: -0.5px;
}

.header-right {
  text-align: right;
}

.user-welcome {
  font-size: 0.95rem;
  color: #64748b;
  font-weight: 500;
}

.user-welcome strong {
  color: #1e3a5f;
  font-weight: 700;
}

/* Content Body */
.content-body {
  flex: 1;
  padding: 2rem;
  overflow-y: auto;
}

/* Scrollbar Styling */
.sidebar-nav::-webkit-scrollbar,
.content-body::-webkit-scrollbar {
  width: 6px;
}

.sidebar-nav::-webkit-scrollbar-track,
.content-body::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.05);
}

.sidebar-nav::-webkit-scrollbar-thumb,
.content-body::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 3px;
}

.sidebar-nav::-webkit-scrollbar-thumb:hover,
.content-body::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.3);
}

/* Responsive */
@media (max-width: 768px) {
  .sidebar {
    width: 250px;
  }

  .main-content {
    margin-left: 250px;
  }

  .content-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 1rem;
  }

  .header-right {
    width: 100%;
    text-align: left;
  }
}

@media (max-width: 640px) {
  .sidebar {
    position: fixed;
    left: -280px;
    transition: left 0.3s ease;
  }

  .main-content {
    margin-left: 0;
  }

  .layout.menu-open .sidebar {
    left: 0;
  }
}
</style>
