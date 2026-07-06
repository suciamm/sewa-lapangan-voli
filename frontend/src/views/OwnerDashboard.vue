<template>
  <Layout>
    <div class="dashboard">
      <h1>Dashboard Pemilik Lapangan</h1>
      <div class="cards">
        <div class="card" @click="$router.push('/owner/courts')">
          <div class="card-icon">🏟️</div>
          <h3>Kelola Lapangan</h3>
          <p>Lihat dan kelola semua lapangan Anda</p>
        </div>
        <div class="card" @click="$router.push('/owner/bookings')">
          <div class="card-icon">📋</div>
          <h3>Pesanan Masuk</h3>
          <p>Lihat semua pesanan yang masuk</p>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import Layout from '@/components/Layout.vue'
</script>

<style scoped>
.dashboard h1 {
  color: #1e3a5f;
  margin-bottom: 2rem;
  font-weight: 800;
  font-size: 2.2rem;
}

.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 1.5rem;
}

.card {
  background: white;
  padding: 2.5rem 2rem;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
  text-align: center;
  cursor: pointer;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.12);
  border-color: #ff6b35;
}

.card-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.card h3 {
  color: #1e3a5f;
  margin-bottom: 0.75rem;
  font-weight: 700;
  font-size: 1.35rem;
}

.card p {
  color: #64748b;
}
</style>
