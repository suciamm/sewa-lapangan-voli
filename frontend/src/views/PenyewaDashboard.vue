<template>
  <Layout>
    <div class="dashboard">
      <!-- Welcome & Quick Stats -->
      <div class="welcome-section">
        <div class="welcome-content">
          <h2>Selamat datang, {{ userName }}! 👋</h2>
          <p>Kelola booking dan temukan lapangan voli terbaik untuk Anda</p>
        </div>
        <div class="quick-stats">
          <div class="quick-stat">
            <span class="stat-number">5</span>
            <span class="stat-label">Booking Aktif</span>
          </div>
          <div class="quick-stat">
            <span class="stat-number">12</span>
            <span class="stat-label">Total Booking</span>
          </div>
          <div class="quick-stat">
            <span class="stat-number">4.6⭐</span>
            <span class="stat-label">Rating Anda</span>
          </div>
        </div>
      </div>

      <!-- Primary Actions -->
      <div class="primary-actions">
        <button class="primary-btn" @click="$router.push('/courts')">
          <span class="primary-icon">🔍</span>
          <div class="primary-content">
            <h3>Cari Lapangan Baru</h3>
            <p>Temukan lapangan dengan harga terbaik</p>
          </div>
          <span class="arrow">→</span>
        </button>
      </div>

      <!-- Booking Status Overview -->
      <div class="booking-overview">
        <h3 class="section-title">Status Booking Saya</h3>
        <div class="status-cards">
          <div class="status-card" style="border-left: 4px solid #3b82f6">
            <div class="status-icon" style="background: #dbeafe">📅</div>
            <div class="status-info">
              <p class="status-count">3</p>
              <p class="status-text">Menunggu Konfirmasi</p>
            </div>
            <button class="status-btn" @click="$router.push('/penyewa/bookings')">
              Lihat
            </button>
          </div>

          <div class="status-card" style="border-left: 4px solid #10b981">
            <div class="status-icon" style="background: #d1fae5">✅</div>
            <div class="status-info">
              <p class="status-count">2</p>
              <p class="status-text">Booking Terkonfirmasi</p>
            </div>
            <button class="status-btn" @click="$router.push('/penyewa/bookings')">
              Lihat
            </button>
          </div>

          <div class="status-card" style="border-left: 4px solid #f59e0b">
            <div class="status-icon" style="background: #fef3c7">⭐</div>
            <div class="status-info">
              <p class="status-count">1</p>
              <p class="status-text">Menunggu Rating</p>
            </div>
            <button class="status-btn" @click="$router.push('/penyewa/ratings')">
              Beri Rating
            </button>
          </div>
        </div>
      </div>

      <!-- Active Bookings -->
      <div class="active-bookings">
        <h3 class="section-title">Booking Aktif Saya</h3>
        <div class="bookings-list">
          <div class="booking-card">
            <div class="booking-header">
              <h4>Lapangan Utama - GOR Merdeka</h4>
              <span class="booking-badge badge-active">Aktif</span>
            </div>
            <div class="booking-details">
              <div class="detail-item">
                <span class="detail-label">📅 Tanggal</span>
                <span class="detail-value">10 Juli 2026, 18:00 - 20:00</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">📍 Lokasi</span>
                <span class="detail-value">Jl. Merdeka No. 123, Jakarta Pusat</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">💰 Total Bayar</span>
                <span class="detail-value">Rp 600.000</span>
              </div>
            </div>
            <div class="booking-actions">
              <button class="btn-secondary">Lihat Detail</button>
              <button class="btn-danger">Batalkan Booking</button>
            </div>
          </div>

          <div class="booking-card">
            <div class="booking-header">
              <h4>Lapangan Premium - Sana Indah</h4>
              <span class="booking-badge badge-pending">Menunggu</span>
            </div>
            <div class="booking-details">
              <div class="detail-item">
                <span class="detail-label">📅 Tanggal</span>
                <span class="detail-value">12 Juli 2026, 19:00 - 21:00</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">📍 Lokasi</span>
                <span class="detail-value">Jl. Sanur No. 45, Jakarta Selatan</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">💰 Total Bayar</span>
                <span class="detail-value">Rp 800.000</span>
              </div>
            </div>
            <div class="booking-actions">
              <button class="btn-secondary">Lihat Detail</button>
              <button class="btn-primary">Bayar Sekarang</button>
            </div>
          </div>
        </div>
        <div class="view-all">
          <router-link to="/penyewa/bookings">Lihat semua booking →</router-link>
        </div>
      </div>

      <!-- Recommendations -->
      <div class="recommendations">
        <h3 class="section-title">Lapangan Rekomendasi Untuk Anda</h3>
        <div class="courts-carousel">
          <div class="court-item">
            <div class="court-image">🏟️</div>
            <h4>Lapangan Bintang 5</h4>
            <p class="rating">⭐⭐⭐⭐⭐ (89 review)</p>
            <p class="price">Rp 500.000/jam</p>
            <button class="btn-small">Booking</button>
          </div>

          <div class="court-item">
            <div class="court-image">🏟️</div>
            <h4>Lapangan Elit Alam</h4>
            <p class="rating">⭐⭐⭐⭐⭐ (76 review)</p>
            <p class="price">Rp 450.000/jam</p>
            <button class="btn-small">Booking</button>
          </div>

          <div class="court-item">
            <div class="court-image">🏟️</div>
            <h4>Lapangan Merdeka Sport</h4>
            <p class="rating">⭐⭐⭐⭐ (64 review)</p>
            <p class="price">Rp 350.000/jam</p>
            <button class="btn-small">Booking</button>
          </div>
        </div>
        <div class="view-all">
          <router-link to="/courts">Lihat semua lapangan →</router-link>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { computed } from 'vue'
import Layout from '@/components/Layout.vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const userName = computed(() => authStore.user?.name || 'Penyewa')
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

/* Welcome Section */
.welcome-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 2rem;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.3);
}

.welcome-content h2 {
  font-size: 1.75rem;
  font-weight: 800;
  margin: 0 0 0.5rem 0;
}

.welcome-content p {
  margin: 0;
  font-size: 1rem;
  opacity: 0.95;
}

.quick-stats {
  display: flex;
  gap: 2rem;
}

.quick-stat {
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.stat-number {
  font-size: 1.75rem;
  font-weight: 800;
}

.stat-label {
  font-size: 0.85rem;
  opacity: 0.9;
}

/* Primary Actions */
.primary-actions {
  display: flex;
  gap: 1rem;
}

.primary-btn {
  flex: 1;
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  color: white;
  border: none;
  border-radius: 12px;
  padding: 1.5rem 2rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(255, 107, 53, 0.3);
}

.primary-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(255, 107, 53, 0.4);
}

.primary-icon {
  font-size: 2.5rem;
}

.primary-content {
  flex: 1;
  text-align: left;
}

.primary-content h3 {
  margin: 0;
  font-size: 1.1rem;
}

.primary-content p {
  margin: 0.25rem 0 0 0;
  opacity: 0.95;
  font-size: 0.9rem;
}

.arrow {
  font-size: 1.5rem;
  transition: transform 0.3s ease;
}

.primary-btn:hover .arrow {
  transform: translateX(4px);
}

/* Section Title */
.section-title {
  font-size: 1.3rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0;
  padding-bottom: 1rem;
  border-bottom: 2px solid #e2e8f0;
}

/* Booking Overview */
.booking-overview {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.status-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.status-card {
  border-radius: 12px;
  padding: 1.5rem;
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: all 0.3s ease;
}

.status-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.status-icon {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.8rem;
  flex-shrink: 0;
}

.status-info {
  flex: 1;
}

.status-count {
  font-size: 1.5rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0;
}

.status-text {
  font-size: 0.9rem;
  color: #64748b;
  margin: 0.25rem 0 0 0;
}

.status-btn {
  background-color: white;
  color: #1e3a5f;
  border: 2px solid #1e3a5f;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  font-size: 0.85rem;
  transition: all 0.2s ease;
  white-space: nowrap;
}

.status-btn:hover {
  background-color: #1e3a5f;
  color: white;
}

/* Active Bookings */
.active-bookings {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.bookings-list {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.booking-card {
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  padding: 1.5rem;
  transition: all 0.3s ease;
}

.booking-card:hover {
  border-color: #ff6b35;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.booking-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #e2e8f0;
}

.booking-header h4 {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0;
}

.booking-badge {
  display: inline-block;
  padding: 0.35rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.badge-active {
  background-color: #d1fae5;
  color: #065f46;
}

.badge-pending {
  background-color: #fef3c7;
  color: #92400e;
}

.booking-details {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.detail-item {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.detail-label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 600;
}

.detail-value {
  font-size: 0.95rem;
  color: #1e3a5f;
  font-weight: 600;
}

.booking-actions {
  display: flex;
  gap: 1rem;
}

.btn-primary {
  flex: 1;
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.2s ease;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(255, 107, 53, 0.3);
}

.btn-secondary {
  flex: 1;
  background: white;
  color: #1e3a5f;
  border: 2px solid #1e3a5f;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.2s ease;
}

.btn-secondary:hover {
  background: #1e3a5f;
  color: white;
}

.btn-danger {
  flex: 1;
  background: rgba(239, 68, 68, 0.1);
  color: #dc2626;
  border: 2px solid #dc2626;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.2s ease;
}

.btn-danger:hover {
  background: #dc2626;
  color: white;
}

/* Recommendations */
.recommendations {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.courts-carousel {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.court-item {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border-radius: 12px;
  padding: 1.5rem;
  text-align: center;
  border: 2px solid #e2e8f0;
  transition: all 0.3s ease;
}

.court-item:hover {
  transform: translateY(-4px);
  border-color: #ff6b35;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
}

.court-image {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.court-item h4 {
  font-size: 1rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0 0 0.5rem 0;
}

.rating {
  color: #ff6b35;
  font-weight: 700;
  font-size: 0.9rem;
  margin: 0.25rem 0;
}

.price {
  font-size: 1.1rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0.5rem 0;
}

.btn-small {
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  color: white;
  border: none;
  padding: 0.6rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  width: 100%;
  transition: all 0.2s ease;
  margin-top: 1rem;
}

.btn-small:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(255, 107, 53, 0.3);
}

.view-all {
  margin-top: 1.5rem;
  text-align: center;
}

.view-all a {
  color: #ff6b35;
  text-decoration: none;
  font-weight: 700;
  transition: color 0.2s ease;
}

.view-all a:hover {
  color: #ff5220;
}

/* Responsive */
@media (max-width: 768px) {
  .welcome-section {
    flex-direction: column;
    text-align: center;
  }

  .quick-stats {
    justify-content: center;
  }

  .primary-btn {
    flex-direction: column;
  }

  .primary-content {
    text-align: center;
  }

  .courts-carousel {
    grid-template-columns: 1fr;
  }
}
</style>
