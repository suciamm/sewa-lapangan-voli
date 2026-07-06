<template>
  <div class="courts">
    <h1>Daftar Lapangan</h1>
    <div class="courts-grid">
      <div v-for="court in courts" :key="court.id" class="court-card">
        <div class="court-image">
          <img :src="court.image || 'https://coresg-normal.trae.ai/api/ide/v1/text_to_image?prompt=volleyball%20court%20indoor%20professional&image_size=square_hd'" :alt="court.name">
        </div>
        <div class="court-info">
          <h3>{{ court.name }}</h3>
          <p class="location">{{ court.location }}</p>
          <p class="price">Rp {{ court.pricePerHour }} / jam</p>
          <button class="btn-book" @click="bookCourt(court)">Pesan Sekarang</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '@/api'
import { useRouter } from 'vue-router'

const router = useRouter()
const courts = ref([])

const fetchCourts = async () => {
  try {
    const response = await api.get('/courts')
    courts.value = response.data.data || []
  } catch (error) {
    console.error('Error fetching courts:', error)
  }
}

const bookCourt = (court) => {
  router.push(`/bookings?courtId=${court.id}`)
}

onMounted(() => {
  fetchCourts()
})
</script>

<style scoped>
.courts h1 {
  margin-bottom: 2.5rem;
  color: #1e3a5f;
  font-weight: 800;
  font-size: 2.2rem;
}

.courts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2rem;
}

.court-card {
  background: white;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(0,0,0,0.08);
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.court-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 30px rgba(0,0,0,0.12);
  border-color: #ff6b35;
}

.court-image {
  height: 220px;
  overflow: hidden;
}

.court-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s ease;
}

.court-card:hover .court-image img {
  transform: scale(1.05);
}

.court-info {
  padding: 1.75rem;
}

.court-info h3 {
  color: #1e3a5f;
  margin-bottom: 0.75rem;
  font-weight: 700;
  font-size: 1.35rem;
}

.location {
  color: #64748b;
  margin-bottom: 0.75rem;
  font-size: 0.95rem;
}

.price {
  font-size: 1.4rem;
  font-weight: 800;
  color: #ff6b35;
  margin-bottom: 1.25rem;
}

.btn-book {
  width: 100%;
  background: linear-gradient(135deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  border: none;
  padding: 1rem;
  border-radius: 10px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(30, 58, 95, 0.2);
}

.btn-book:hover {
  background: linear-gradient(135deg, #0f2744 0%, #1e3a5f 100%);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(30, 58, 95, 0.3);
}
</style>
