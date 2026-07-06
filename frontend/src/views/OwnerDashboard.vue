<template>
  <Layout>
    <div class="dashboard">
      <!-- Stats Overview -->
      <div class="stats-grid">
        <div class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
            📊
          </div>
          <div class="stat-content">
            <p class="stat-label">Total Lapangan</p>
            <p class="stat-value">12</p>
            <p class="stat-change">↑ 2 lapangan bulan ini</p>
          </div>
        </div>

        <div class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%)">
            📋
          </div>
          <div class="stat-content">
            <p class="stat-label">Booking Bulan Ini</p>
            <p class="stat-value">48</p>
            <p class="stat-change">↑ 12% dari bulan lalu</p>
          </div>
        </div>

        <div class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
            💰
          </div>
          <div class="stat-content">
            <p class="stat-label">Pendapatan Bulan Ini</p>
            <p class="stat-value">Rp 24.5M</p>
            <p class="stat-change">↑ 8% dari bulan lalu</p>
          </div>
        </div>

        <div class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%)">
            ⭐
          </div>
          <div class="stat-content">
            <p class="stat-label">Rating Rata-rata</p>
            <p class="stat-value">4.8</p>
            <p class="stat-change">dari 5 bintang (142 review)</p>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="quick-actions">
        <h3 class="section-title">Akses Cepat</h3>
        <div class="action-grid">
          <button class="action-btn" @click="$router.push('/owner/courts')">
            <span class="action-icon">🏟️</span>
            <span class="action-label">Kelola Lapangan</span>
            <span class="action-desc">Tambah, edit, atau hapus lapangan</span>
          </button>
          <button class="action-btn" @click="$router.push('/owner/bookings')">
            <span class="action-icon">✅</span>
            <span class="action-label">Booking Masuk</span>
            <span class="action-desc">Lihat dan kelola pesanan</span>
          </button>
          <button class="action-btn" @click="$router.push('/owner/report')">
            <span class="action-icon">📈</span>
            <span class="action-label">Laporan & Rating</span>
            <span class="action-desc">Analisis pendapatan & rating</span>
          </button>
        </div>
      </div>

      <!-- Recent Bookings -->
      <div class="recent-section">
        <h3 class="section-title">Booking Terbaru</h3>
        <div class="table-container">
          <table class="data-table">
            <thead>
              <tr>
                <th>ID Booking</th>
                <th>Lapangan</th>
                <th>Penyewa</th>
                <th>Tanggal</th>
                <th>Status</th>
                <th>Total</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>#BK001</td>
                <td>Lapangan Utama</td>
                <td>Budi Santoso</td>
                <td>06 Juli 2026</td>
                <td><span class="badge badge-success">Selesai</span></td>
                <td>Rp 500.000</td>
              </tr>
              <tr>
                <td>#BK002</td>
                <td>Lapangan Cadangan</td>
                <td>Siti Rahmawati</td>
                <td>06 Juli 2026</td>
                <td><span class="badge badge-info">Menunggu</span></td>
                <td>Rp 400.000</td>
              </tr>
              <tr>
                <td>#BK003</td>
                <td>Lapangan Premium</td>
                <td>Ahmad Hidayat</td>
                <td>05 Juli 2026</td>
                <td><span class="badge badge-success">Selesai</span></td>
                <td>Rp 750.000</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="view-all">
          <router-link to="/owner/bookings">Lihat semua booking →</router-link>
        </div>
      </div>

      <!-- Best Performing Courts -->
      <div class="best-courts">
        <h3 class="section-title">Lapangan Terbaik</h3>
        <div class="courts-grid">
          <div class="court-card">
            <div class="court-header">
              <h4>Lapangan Utama</h4>
              <span class="rating">⭐ 4.9 (45 review)</span>
            </div>
            <p class="court-stat">Booking: 18 bulan ini</p>
            <p class="court-stat">Pendapatan: Rp 9M</p>
          </div>
          <div class="court-card">
            <div class="court-header">
              <h4>Lapangan Premium</h4>
              <span class="rating">⭐ 4.7 (32 review)</span>
            </div>
            <p class="court-stat">Booking: 15 bulan ini</p>
            <p class="court-stat">Pendapatan: Rp 11.25M</p>
          </div>
          <div class="court-card">
            <div class="court-header">
              <h4>Lapangan Cadangan</h4>
              <span class="rating">⭐ 4.8 (28 review)</span>
            </div>
            <p class="court-stat">Booking: 12 bulan ini</p>
            <p class="court-stat">Pendapatan: Rp 4.8M</p>
          </div>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import Layout from '@/components/Layout.vue'
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

/* Stats Grid */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
}

.stat-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  display: flex;
  gap: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  border-color: #ff6b35;
}

.stat-icon {
  width: 70px;
  height: 70px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  flex-shrink: 0;
}

.stat-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.stat-label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin: 0;
}

.stat-value {
  font-size: 1.75rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0.25rem 0 0 0;
}

.stat-change {
  font-size: 0.8rem;
  color: #22c55e;
  margin: 0.5rem 0 0 0;
  font-weight: 600;
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

/* Quick Actions */
.quick-actions {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.action-btn {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  padding: 1.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  text-align: center;
  font-family: inherit;
}

.action-btn:hover {
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  border-color: #ff6b35;
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(255, 107, 53, 0.2);
}

.action-btn:hover .action-icon,
.action-btn:hover .action-label {
  color: white;
}

.action-icon {
  font-size: 2.5rem;
  transition: color 0.3s ease;
}

.action-label {
  font-weight: 700;
  color: #1e3a5f;
  font-size: 1rem;
  transition: color 0.3s ease;
}

.action-desc {
  font-size: 0.8rem;
  color: #64748b;
  margin-top: 0.5rem;
}

/* Recent Section */
.recent-section {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.table-container {
  overflow-x: auto;
  margin-top: 1.5rem;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table thead {
  background-color: #f8fafc;
}

.data-table th {
  padding: 1rem;
  text-align: left;
  font-weight: 700;
  color: #1e3a5f;
  font-size: 0.9rem;
  border-bottom: 2px solid #e2e8f0;
}

.data-table td {
  padding: 1rem;
  border-bottom: 1px solid #e2e8f0;
  color: #64748b;
}

.data-table tbody tr:hover {
  background-color: #f8fafc;
}

.badge {
  display: inline-block;
  padding: 0.35rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.badge-success {
  background-color: #d1fae5;
  color: #065f46;
}

.badge-info {
  background-color: #dbeafe;
  color: #1e40af;
}

.badge-warning {
  background-color: #fef3c7;
  color: #92400e;
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

/* Best Courts */
.best-courts {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.courts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.court-card {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border-radius: 12px;
  padding: 1.5rem;
  border: 2px solid #e2e8f0;
  transition: all 0.3s ease;
}

.court-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
}

.court-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1rem;
}

.court-header h4 {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0;
}

.rating {
  font-size: 0.85rem;
  color: #ff6b35;
  font-weight: 600;
  white-space: nowrap;
}

.court-stat {
  font-size: 0.9rem;
  color: #1e3a5f;
  margin: 0.5rem 0;
  font-weight: 600;
}

/* Responsive */
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }

  .action-grid {
    grid-template-columns: 1fr;
  }

  .courts-grid {
    grid-template-columns: 1fr;
  }
}
</style>
