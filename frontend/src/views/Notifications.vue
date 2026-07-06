<template>
  <Layout>
    <div class="page">
      <div class="page-header">
        <h1>Notifikasi</h1>
        <button class="btn-outline" @click="markAllRead" v-if="unreadCount > 0">
          Tandai Semua Dibaca
        </button>
      </div>
      <div class="notif-list">
        <div v-for="notif in notifications" :key="notif.id" class="notif-card" :class="{ unread: !notif.is_read }" @click="markRead(notif.id)">
          <div class="notif-icon">
            {{ notif.type === 'booking_confirmed' ? '✅' : notif.type === 'payment_success' ? '💳' : '❌' }}
          </div>
          <div class="notif-content">
            <h4>{{ notif.title }}</h4>
            <p>{{ notif.body }}</p>
            <span class="notif-time">{{ formatDate(notif.created_at) }}</span>
          </div>
        </div>
        <div v-if="notifications.length === 0" class="empty-state">
          <p>Tidak ada notifikasi</p>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Layout from '@/components/Layout.vue'
import api from '@/api'

const notifications = ref([])

const unreadCount = computed(() => notifications.value.filter(n => !n.is_read).length)

const formatDate = (d) => new Date(d).toLocaleString('id-ID')

const fetchNotifications = async () => {
  try {
    const res = await api.get('/notifications')
    notifications.value = res.data.data || []
  } catch (err) {
    console.error(err)
  }
}

const markRead = async (id) => {
  try {
    await api.put(`/notifications/${id}/read`)
    fetchNotifications()
  } catch (err) {
    console.error(err)
  }
}

const markAllRead = async () => {
  try {
    await api.put('/notifications/read-all')
    fetchNotifications()
  } catch (err) {
    console.error(err)
  }
}

onMounted(fetchNotifications)
</script>

<style scoped>
.page h1 {
  color: #1e3a5f;
  margin-bottom: 1rem;
  font-weight: 800;
  font-size: 2rem;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.btn-outline {
  padding: 0.6rem 1.2rem;
  border: 2px solid #1e3a5f;
  color: #1e3a5f;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.3s ease;
}

.btn-outline:hover {
  background: #1e3a5f;
  color: white;
}

.notif-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.notif-card {
  display: flex;
  gap: 1rem;
  background: white;
  padding: 1.25rem;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  cursor: pointer;
  transition: all 0.2s ease;
}

.notif-card:hover {
  box-shadow: 0 2px 12px rgba(0,0,0,0.08);
}

.notif-card.unread {
  background: #eff6ff;
  border-color: #bfdbfe;
}

.notif-icon {
  font-size: 1.5rem;
  width: 40px;
  text-align: center;
}

.notif-content h4 {
  color: #1e3a5f;
  margin-bottom: 0.25rem;
  font-weight: 700;
}

.notif-content p {
  color: #475569;
  margin-bottom: 0.5rem;
}

.notif-time {
  color: #94a3b8;
  font-size: 0.85rem;
}

.empty-state {
  text-align: center;
  padding: 3rem;
  color: #64748b;
}
</style>
