<template>
  <div id="app">
    <nav class="navbar">
      <div class="container">
        <router-link to="/" class="logo">Sewa Lapangan Voli</router-link>
        <div class="nav-links">
          <router-link to="/">Beranda</router-link>
          <router-link to="/courts">Lapangan</router-link>
          <router-link to="/login" v-if="!isAuthenticated">Masuk</router-link>
          <router-link to="/register" v-if="!isAuthenticated">Daftar</router-link>
          <button @click="logout" v-if="isAuthenticated">Keluar</button>
        </div>
      </div>
    </nav>
    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)

const logout = () => {
  authStore.logout()
  router.push('/')
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f8fafc;
  color: #1e293b;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.navbar {
  background: linear-gradient(135deg, #1e3a5f 0%, #0f2744 100%);
  padding: 1.25rem 0;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.navbar .container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  color: white;
  font-size: 1.6rem;
  font-weight: 700;
  text-decoration: none;
  letter-spacing: 0.5px;
}

.nav-links {
  display: flex;
  gap: 2rem;
  align-items: center;
}

.nav-links a {
  color: rgba(255,255,255,0.9);
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
}

.nav-links a:hover {
  color: #ff6b35;
  transform: translateY(-1px);
}

.nav-links button {
  background: linear-gradient(135deg, #ff6b35 0%, #e63946 100%);
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  font-size: 0.95rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(230, 57, 70, 0.3);
}

.nav-links button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(230, 57, 70, 0.4);
}

.main-content {
  flex: 1;
  max-width: 1280px;
  margin: 0 auto;
  padding: 2.5rem 2rem;
  width: 100%;
}
</style>
