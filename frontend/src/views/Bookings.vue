<template>
  <div class="bookings">
    <h1>Pemesanan Lapangan</h1>
    <div class="booking-form-card">
      <form @submit.prevent="handleBooking">
        <div class="form-group">
          <label>Lapangan</label>
          <select v-model="form.courtId" required>
            <option value="">Pilih Lapangan</option>
            <option v-for="court in courts" :key="court.id" :value="court.id">{{ court.name }}</option>
          </select>
        </div>
        <div class="form-group">
          <label>Tanggal</label>
          <input type="date" v-model="form.date" required>
        </div>
        <div class="form-group">
          <label>Jam Mulai</label>
          <input type="time" v-model="form.startTime" required>
        </div>
        <div class="form-group">
          <label>Jam Selesai</label>
          <input type="time" v-model="form.endTime" required>
        </div>
        <button type="submit" class="btn-submit" :disabled="loading">
          {{ loading ? 'Loading...' : 'Pesan' }}
        </button>
      </form>
    </div>

    <h2 style="margin-top: 2rem;">Riwayat Pemesanan</h2>
    <div class="bookings-list">
      <div v-for="booking in bookings" :key="booking.id" class="booking-card">
        <div class="booking-info">
          <h3>{{ booking.courtName }}</h3>
          <p>Tanggal: {{ booking.date }}</p>
          <p>Waktu: {{ booking.startTime }} - {{ booking.endTime }}</p>
          <p>Status: <span :class="booking.status">{{ booking.status }}</span></p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/api'

const route = useRoute()
const courts = ref([])
const bookings = ref([])
const loading = ref(false)

const form = ref({
  courtId: route.query.courtId || '',
  date: '',
  startTime: '',
  endTime: ''
})

const fetchCourts = async () => {
  try {
    const response = await api.get('/courts')
    courts.value = response.data.data || []
  } catch (error) {
    console.error('Error fetching courts:', error)
  }
}

const fetchBookings = async () => {
  try {
    const response = await api.get('/bookings')
    bookings.value = response.data.data || []
  } catch (error) {
    console.error('Error fetching bookings:', error)
  }
}

const handleBooking = async () => {
  loading.value = true
  try {
    await api.post('/bookings', form.value)
    alert('Pemesanan berhasil!')
    fetchBookings()
  } catch (error) {
    alert('Pemesanan gagal: ' + (error.response?.data?.message || error.message))
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchCourts()
  fetchBookings()
})
</script>

<style scoped>
.bookings h1 {
  color: #1e3a5f;
  margin-bottom: 2rem;
  font-weight: 800;
  font-size: 2.2rem;
}

.bookings h2 {
  color: #1e3a5f;
  margin-bottom: 1.75rem;
  font-weight: 700;
  font-size: 1.6rem;
}

.booking-form-card {
  background: white;
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  margin-bottom: 3rem;
  border: 1px solid #e2e8f0;
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

.form-group input, .form-group select {
  width: 100%;
  padding: 1rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: white;
}

.form-group input:focus, .form-group select:focus {
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

.bookings-list {
  display: grid;
  gap: 1.25rem;
}

.booking-card {
  background: white;
  padding: 1.75rem;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.08);
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.booking-card:hover {
  border-color: #ff6b35;
  box-shadow: 0 6px 20px rgba(0,0,0,0.1);
}

.booking-card h3 {
  color: #1e3a5f;
  margin-bottom: 0.75rem;
  font-weight: 700;
  font-size: 1.25rem;
}

.booking-card p {
  color: #64748b;
  margin-bottom: 0.5rem;
  font-size: 0.95rem;
}

.booking-card .pending {
  color: #f59e0b;
  font-weight: 700;
}

.booking-card .confirmed {
  color: #10b981;
  font-weight: 700;
}

.booking-card .cancelled {
  color: #ef4444;
  font-weight: 700;
}
</style>
