<template>
  <Layout>
    <div class="page">
      <h1>Pesanan Saya</h1>
      <div class="bookings-list">
        <div v-for="booking in bookings" :key="booking.id" class="booking-card">
          <div class="booking-header">
            <h3>Pesanan #{{ booking.id }}</h3>
            <span class="status-badge" :class="getStatusClass(booking.status)">
              {{ booking.status }}
            </span>
          </div>
        </div>
        <div v-if="bookings.length === 0" class="empty-state">
          <p>Belum ada pesanan</p>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '@/components/Layout.vue'
import api from '@/api'

const bookings = ref([])

const formatPrice = (num) => new Intl.NumberFormat('id-ID').format(num)
const getStatusClass = (status) => status.replace('_', '-').toLowerCase()

const fetchBookings = async () => {
  try {
    const res = await api.get('/penyewa/bookings')
    bookings.value = res.data.data || []
  } catch (err) {
    console.error(err)
  }
}

onMounted(fetchBookings)
</script>

<style scoped>
.page h1 {
  color: #1e3a5f;
  margin-bottom: 2rem;
  font-weight: 800;
  font-size: 2rem;
}

.bookings-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.booking-card {
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  border: 1px solid #e2e8f0;
}

.booking-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.booking-header h3 {
  color: #1e3a5f;
  font-weight: 700;
}

.status-badge {
  padding: 0.4rem 1rem;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 700;
}

.pending-payment { background-color: #fef3c7; color: #92400e; }
.paid { background-color: #d1fae5; color: #065f46; }
.active { background-color: #dbeafe; color: #1e40af; }
.completed { background-color: #e0e7ff; color: #3730a3; }
.cancelled { background-color: #fee2e2; color: #991b1b; }

.empty-state {
  text-align: center;
  padding: 3rem;
  color: #64748b;
}
</style>
