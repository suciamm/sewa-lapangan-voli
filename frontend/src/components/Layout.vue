<template>
  <div class="layout">
    <!-- Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-header">
        <h2>Sewa Lapangan</h2>
      </div>
      <nav class="sidebar-nav">
        <template v-if="userRole === 'owner'">
          <router-link to="/owner/dashboard" class="nav-link" active-class="active">
            <span>Dashboard Ringkasan Booking</span>
          </router-link>
          <router-link to="/owner/courts" class="nav-link" active-class="active">
            <span>Kelola Lapangan</span>
          </router-link>
          <router-link to="/owner/report" class="nav-link" active-class="active">
            <span>Report Rating & Pendapatan</span>
          </router-link>
        </template>
        <template v-else-if="userRole === 'penyewa'">
          <router-link to="/courts" class="nav-link" active-class="active">
            <span>Cari Lapangan</span>
          </router-link>
          <router-link to="/penyewa/booking-aktif" class="nav-link" active-class="active">
            <span>Booking Aktif</span>
          </router-link>
          <router-link to="/penyewa/beri-rating" class="nav-link" active-class="active">
            <span>Beri Rating</span>
          </router-link>
          <router-link to="/notifications" class="nav-link" active-class="active">
            <span>Notifikasi</span>
          </router-link>
        </template>
        <template v-else-if="userRole === 'superuser'">
          <router-link to="/superuser/dashboard" class="nav-link" active-class="active">
            <span>Dashboard Ringkasan Sistem</span>
          </router-link>
          <router-link to="/superuser/kelola-owner" class="nav-link" active-class="active">
            <span>Kelola Owner</span>
          </router-link>
          <router-link to="/superuser/data-lapangan" class="nav-link" active-class="active">
            <span>Data Lapangan</span>
          </router-link>
          <router-link to="/superuser/report" class="nav-link" active-class="active">
            <span>Report Rating & Booking</span>
          </router-link>
        </template>
      </nav>
      <div class="sidebar-footer">
        <button @click="logout" class="btn-logout">
          Keluar
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="main-content">
      <header class="content-header">
        <div class="user-info">
          <span>{{ userName }}</span>
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
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const userRole = computed(() => authStore.user?.role)
const userName = computed(() => authStore.user?.name)

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 280px;
  background: linear-gradient(180deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  display: flex;
  flex-direction: column;
  box-shadow: 4px 0 12px rgba(0,0,0,0.1);
}

.sidebar-header {
  padding: 2rem 1.5rem;
  border-bottom: 1px solid rgba(255,255,255,0.1);
}

.sidebar-header h2 {
  font-size: 1.4rem;
  font-weight: 800;
  letter-spacing: 0.5px;
}

.sidebar-nav {
  flex: 1;
  padding: 1rem 0;
}

.nav-link {
  display: block;
  color: rgba(255,255,255,0.85);
  text-decoration: none;
  padding: 1rem 1.5rem;
  font-weight: 500;
  transition: all 0.2s ease;
  border-left: 4px solid transparent;
}

.nav-link:hover {
  color: white;
  background-color: rgba(255,255,255,0.08);
  border-left-color: #ff6b35;
}

.nav-link.active {
  color: white;
  background-color: rgba(255,107,53,0.15);
  border-left-color: #ff6b35;
  font-weight: 700;
}

.sidebar-footer {
  padding: 1.5rem;
  border-top: 1px solid rgba(255,255,255,0.1);
}

.btn-logout {
  width: 100%;
  background-color: rgba(230, 57, 70, 0.2);
  color: #ff8a8a;
  border: 1px solid rgba(230, 57, 70, 0.3);
  padding: 0.8rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-logout:hover {
  background-color: #e63946;
  color: white;
  border-color: transparent;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #f8fafc;
}

.content-header {
  background-color: white;
  padding: 1rem 2rem;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.user-info {
  font-weight: 600;
  color: #1e3a5f;
}

.content-body {
  flex: 1;
  padding: 2rem;
  overflow-y: auto;
}
</style>
