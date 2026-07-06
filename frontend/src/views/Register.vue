<template>
  <div class="auth-container">
    <div class="auth-card">
      <h2>Daftar</h2>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label>Nama</label>
          <input type="text" v-model="form.name" required>
        </div>
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="form.email" required>
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" v-model="form.password" required>
        </div>
        <div class="form-group">
          <label>Konfirmasi Password</label>
          <input type="password" v-model="form.confirmPassword" required>
        </div>
        <button type="submit" class="btn-submit" :disabled="loading">
          {{ loading ? 'Loading...' : 'Daftar' }}
        </button>
      </form>
      <p class="auth-link">
        Sudah punya akun? <router-link to="/login">Masuk</router-link>
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
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
})
const loading = ref(false)

const handleRegister = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    alert('Password tidak cocok')
    return
  }

  loading.value = true
  try {
    const response = await api.post('/auth/register', {
      name: form.value.name,
      email: form.value.email,
      password: form.value.password
    })
    authStore.setAuth(response.data.data.token, response.data.data.user)
    router.push('/')
  } catch (error) {
    alert('Daftar gagal: ' + (error.response?.data?.message || error.message))
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
  min-height: 70vh;
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
  background: linear-gradient(135deg, #ff6b35 0%, #e63946 100%);
  color: white;
  border: none;
  padding: 1rem;
  border-radius: 10px;
  cursor: pointer;
  font-size: 1.05rem;
  font-weight: 700;
  margin-top: 0.5rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(230, 57, 70, 0.3);
}

.btn-submit:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(230, 57, 70, 0.4);
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
  color: #1e3a5f;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.3s ease;
}

.auth-link a:hover {
  color: #ff6b35;
}
</style>
