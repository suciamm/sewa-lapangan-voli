<template>
  <div class="auth-container">
    <div class="auth-card">
      <h2>Masuk</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="form.email" required>
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" v-model="form.password" required>
        </div>
        <button type="submit" class="btn-submit" :disabled="loading">
          {{ loading ? 'Loading...' : 'Masuk' }}
        </button>
      </form>
      <p class="auth-link">
        Belum punya akun? <router-link to="/register">Daftar</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import api from '@/api'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  email: '',
  password: ''
})
const loading = ref(false)

const handleLogin = async () => {
  loading.value = true
  try {
    const response = await api.post('/auth/login', form.value)
    // Note: backend uses "access_token", not "token"!
    const token = response.data.data.access_token
    const user = response.data.data.user
    authStore.setAuth(token, user)
    
    // Redirect based on role
    if (user.role === 'owner') {
      router.push('/owner/dashboard')
    } else if (user.role === 'penyewa') {
      router.push('/penyewa/dashboard')
    } else if (user.role === 'superuser') {
      router.push('/superuser/dashboard')
    } else {
      router.push('/')
    }
  } catch (error) {
    alert('Login gagal: ' + (error.response?.data?.message || error.message))
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 65vh;
}

.auth-card {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  width: 100%;
  max-width: 440px;
  border: 1px solid #e2e8f0;
}

.auth-card h2 {
  text-align: center;
  margin-bottom: 2rem;
  color: #1e3a5f;
  font-weight: 800;
  font-size: 1.8rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.75rem;
  color: #475569;
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 1rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #ff6b35;
  box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
}

.btn-submit {
  width: 100%;
  background: linear-gradient(135deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  border: none;
  padding: 1rem;
  border-radius: 10px;
  cursor: pointer;
  font-size: 1.05rem;
  font-weight: 700;
  margin-top: 0.5rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(30, 58, 95, 0.2);
}

.btn-submit:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(30, 58, 95, 0.3);
}

.btn-submit:disabled {
  background: #94a3b8;
  cursor: not-allowed;
  box-shadow: none;
  transform: none;
}

.auth-link {
  text-align: center;
  margin-top: 1.5rem;
}

.auth-link a {
  color: #ff6b35;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.3s ease;
}

.auth-link a:hover {
  color: #e63946;
}
</style>
