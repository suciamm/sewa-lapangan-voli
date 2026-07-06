<template>
  <Layout>
    <div class="dashboard">
      <!-- Key Metrics -->
      <div class="metrics-grid">
        <div class="metric-card">
          <div class="metric-icon" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%)">
            👥
          </div>
          <div class="metric-content">
            <p class="metric-label">Total Users</p>
            <p class="metric-value">1,245</p>
            <p class="metric-change">↑ 142 minggu ini</p>
          </div>
        </div>

        <div class="metric-card">
          <div class="metric-icon" style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%)">
            🏟️
          </div>
          <div class="metric-content">
            <p class="metric-label">Total Lapangan</p>
            <p class="metric-value">389</p>
            <p class="metric-change">↑ 23 lapangan bulan ini</p>
          </div>
        </div>

        <div class="metric-card">
          <div class="metric-icon" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)">
            📋
          </div>
          <div class="metric-content">
            <p class="metric-label">Total Booking</p>
            <p class="metric-value">5,892</p>
            <p class="metric-change">↑ 18% dari bulan lalu</p>
          </div>
        </div>

        <div class="metric-card">
          <div class="metric-icon" style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%)">
            💰
          </div>
          <div class="metric-content">
            <p class="metric-label">Total Revenue</p>
            <p class="metric-value">Rp 2.1M</p>
            <p class="metric-change">↑ 24% dari bulan lalu</p>
          </div>
        </div>
      </div>

      <!-- System Health -->
      <div class="system-status">
        <h3 class="section-title">Status Sistem</h3>
        <div class="status-grid">
          <div class="status-item">
            <span class="status-indicator healthy"></span>
            <span>Database</span>
            <span class="status-text">Normal</span>
          </div>
          <div class="status-item">
            <span class="status-indicator healthy"></span>
            <span>API Server</span>
            <span class="status-text">Normal</span>
          </div>
          <div class="status-item">
            <span class="status-indicator healthy"></span>
            <span>Payment Gateway</span>
            <span class="status-text">Connected</span>
          </div>
          <div class="status-item">
            <span class="status-indicator healthy"></span>
            <span>Email Service</span>
            <span class="status-text">Active</span>
          </div>
        </div>
      </div>

      <!-- Management Actions -->
      <div class="management-section">
        <h3 class="section-title">Manajemen Sistem</h3>
        <div class="management-grid">
          <button class="management-btn" @click="$router.push('/superuser/kelola-owner')">
            <span class="mgmt-icon">👥</span>
            <h4>Kelola Owner</h4>
            <p>Approve/reject owner baru & kelola akun</p>
            <span class="badge-count">28 pending</span>
          </button>

          <button class="management-btn" @click="$router.push('/superuser/data-lapangan')">
            <span class="mgmt-icon">📋</span>
            <h4>Data Lapangan</h4>
            <p>Lihat & manage semua lapangan di sistem</p>
            <span class="badge-count">389 total</span>
          </button>

          <button class="management-btn" @click="$router.push('/superuser/report')">
            <span class="mgmt-icon">📊</span>
            <h4>Report & Analitik</h4>
            <p>Lihat laporan komprehensif dan analytics</p>
            <span class="badge-count">Real-time</span>
          </button>
        </div>
      </div>

      <!-- Recent Activities -->
      <div class="recent-activities">
        <h3 class="section-title">Aktivitas Terbaru Sistem</h3>
        <div class="activities-list">
          <div class="activity-item">
            <div class="activity-icon" style="background: #d1fae5">✅</div>
            <div class="activity-content">
              <p class="activity-title">Owner Baru Terdaftar</p>
              <p class="activity-desc">PT. Sana Indah Sports telah mendaftar sebagai owner lapangan</p>
              <p class="activity-time">06 Juli 2026, 14:30</p>
            </div>
          </div>

          <div class="activity-item">
            <div class="activity-icon" style="background: #fef3c7">⚠️</div>
            <div class="activity-content">
              <p class="activity-title">Booking Dispute</p>
              <p class="activity-desc">Ada keluhan dari penyewa tentang kualitas lapangan</p>
              <p class="activity-time">05 Juli 2026, 10:15</p>
            </div>
          </div>

          <div class="activity-item">
            <div class="activity-icon" style="background: #dbeafe">📈</div>
            <div class="activity-content">
              <p class="activity-title">Revenue Milestone</p>
              <p class="activity-desc">Total revenue mencapai Rp 2 juta untuk bulan Juli</p>
              <p class="activity-time">04 Juli 2026, 23:45</p>
            </div>
          </div>

          <div class="activity-item">
            <div class="activity-icon" style="background: #c7d2fe">🆕</div>
            <div class="activity-content">
              <p class="activity-title">Lapangan Baru</p>
              <p class="activity-desc">Lapangan Premium Baru telah ditambahkan ke sistem</p>
              <p class="activity-time">03 Juli 2026, 16:20</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Top Courts -->
      <div class="top-courts">
        <h3 class="section-title">Lapangan Terpopuler</h3>
        <div class="table-container">
          <table class="data-table">
            <thead>
              <tr>
                <th>Lapangan</th>
                <th>Owner</th>
                <th>Rating</th>
                <th>Booking</th>
                <th>Revenue</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td><strong>Lapangan Bintang 5</strong></td>
                <td>PT. Sana Indah</td>
                <td>⭐ 4.9 (89)</td>
                <td>156</td>
                <td>Rp 780M</td>
              </tr>
              <tr>
                <td><strong>Lapangan Elit Alam</strong></td>
                <td>CV. Merdeka Sports</td>
                <td>⭐ 4.7 (76)</td>
                <td>134</td>
                <td>Rp 603M</td>
              </tr>
              <tr>
                <td><strong>Lapangan Merdeka Sport</strong></td>
                <td>PT. Indonesia Voli</td>
                <td>⭐ 4.8 (64)</td>
                <td>128</td>
                <td>Rp 448M</td>
              </tr>
              <tr>
                <td><strong>GOR Putra Jaya</strong></td>
                <td>Yayasan Olahraga</td>
                <td>⭐ 4.6 (52)</td>
                <td>98</td>
                <td>Rp 343M</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Pending Actions -->
      <div class="pending-actions">
        <h3 class="section-title">Tindakan yang Diperlukan</h3>
        <div class="pending-list">
          <div class="pending-item">
            <div class="pending-badge">28</div>
            <div class="pending-info">
              <h4>Owner Pending Approval</h4>
              <p>Ada 28 owner baru yang menunggu approval dari admin</p>
            </div>
            <button class="btn-action" @click="$router.push('/superuser/kelola-owner')">
              Review
            </button>
          </div>

          <div class="pending-item">
            <div class="pending-badge warning">3</div>
            <div class="pending-info">
              <h4>Lapangan Perlu Review</h4>
              <p>3 lapangan memiliki banyak komplain dari penyewa</p>
            </div>
            <button class="btn-action warning" @click="$router.push('/superuser/data-lapangan')">
              Review
            </button>
          </div>

          <div class="pending-item">
            <div class="pending-badge info">5</div>
            <div class="pending-info">
              <h4>Dispute Perlu Penanganan</h4>
              <p>5 dispute booking perlu ditangani oleh admin</p>
            </div>
            <button class="btn-action info">
              Resolve
            </button>
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

/* Metrics Grid */
.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
}

.metric-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  display: flex;
  gap: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.metric-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  border-color: #ff6b35;
}

.metric-icon {
  width: 70px;
  height: 70px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  flex-shrink: 0;
}

.metric-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.metric-label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin: 0;
}

.metric-value {
  font-size: 1.75rem;
  font-weight: 800;
  color: #1e3a5f;
  margin: 0.25rem 0 0 0;
}

.metric-change {
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

/* System Status */
.system-status {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.status-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.status-item {
  background: #f8fafc;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  border: 1px solid #e2e8f0;
}

.status-indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  flex-shrink: 0;
}

.status-indicator.healthy {
  background: #22c55e;
  box-shadow: 0 0 0 4px rgba(34, 197, 94, 0.2);
}

.status-text {
  margin-left: auto;
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 600;
}

/* Management Section */
.management-section {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.management-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.management-btn {
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  padding: 1.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.75rem;
  font-family: inherit;
}

.management-btn:hover {
  background: linear-gradient(135deg, #ff6b35 0%, #ff8a5b 100%);
  border-color: #ff6b35;
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(255, 107, 53, 0.2);
}

.management-btn:hover h4,
.management-btn:hover p {
  color: white;
}

.mgmt-icon {
  font-size: 2.5rem;
}

.management-btn h4 {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0;
  transition: color 0.3s ease;
}

.management-btn p {
  font-size: 0.85rem;
  color: #64748b;
  margin: 0;
  transition: color 0.3s ease;
}

.badge-count {
  display: inline-block;
  background: #ff6b35;
  color: white;
  padding: 0.35rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 700;
}

/* Recent Activities */
.recent-activities {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.activities-list {
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.activity-item {
  display: flex;
  gap: 1.5rem;
  padding: 1rem;
  border-radius: 8px;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  transition: all 0.2s ease;
}

.activity-item:hover {
  background: #f0f4f8;
}

.activity-icon {
  width: 50px;
  height: 50px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  flex-shrink: 0;
}

.activity-content {
  flex: 1;
}

.activity-title {
  font-weight: 700;
  color: #1e3a5f;
  margin: 0;
  font-size: 0.95rem;
}

.activity-desc {
  color: #64748b;
  margin: 0.25rem 0;
  font-size: 0.85rem;
}

.activity-time {
  color: #94a3b8;
  font-size: 0.8rem;
  margin: 0.5rem 0 0 0;
}

/* Table */
.top-courts {
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

/* Pending Actions */
.pending-actions {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  border: 1px solid #e2e8f0;
}

.pending-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-top: 1.5rem;
}

.pending-item {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  padding: 1.5rem;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  border-radius: 8px;
  border: 2px solid #e2e8f0;
}

.pending-badge {
  min-width: 50px;
  height: 50px;
  background: #ff6b35;
  color: white;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 1.25rem;
  flex-shrink: 0;
}

.pending-badge.warning {
  background: #f59e0b;
}

.pending-badge.info {
  background: #3b82f6;
}

.pending-info {
  flex: 1;
}

.pending-info h4 {
  font-size: 1rem;
  font-weight: 700;
  color: #1e3a5f;
  margin: 0 0 0.25rem 0;
}

.pending-info p {
  font-size: 0.85rem;
  color: #64748b;
  margin: 0;
}

.btn-action {
  background: white;
  color: #ff6b35;
  border: 2px solid #ff6b35;
  padding: 0.6rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  font-size: 0.9rem;
  transition: all 0.2s ease;
  white-space: nowrap;
}

.btn-action:hover {
  background: #ff6b35;
  color: white;
}

.btn-action.warning {
  color: #f59e0b;
  border-color: #f59e0b;
}

.btn-action.warning:hover {
  background: #f59e0b;
  color: white;
}

.btn-action.info {
  color: #3b82f6;
  border-color: #3b82f6;
}

.btn-action.info:hover {
  background: #3b82f6;
  color: white;
}

/* Responsive */
@media (max-width: 768px) {
  .metrics-grid {
    grid-template-columns: 1fr;
  }

  .status-grid {
    grid-template-columns: 1fr;
  }

  .management-grid {
    grid-template-columns: 1fr;
  }

  .pending-item {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
