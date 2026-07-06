<template>
  <Layout>
    <div class="page">
      <div class="page-header">
        <h1>Kelola Lapangan</h1>
        <button class="btn-primary" @click="showCreateModal = true">
          + Tambah Lapangan
        </button>
      </div>

      <div class="courts-grid">
        <div v-for="court in courts" :key="court.id" class="court-card">
          <div class="court-image">
            <img :src="court.image || placeholderImg" :alt="court.name" />
          </div>
          <div class="court-info">
            <h3>{{ court.name }}</h3>
            <p class="location">{{ court.city || 'Lokasi tidak ditentukan' }}</p>
            <p class="price">Rp {{ formatPrice(court.price_per_hour) }} / jam</p>
            <div class="court-actions">
              <button class="btn-outline" @click="editCourt(court)">Edit</button>
              <button class="btn-danger" @click="deleteCourt(court.id)">Hapus</button>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal Create/Edit -->
      <div v-if="showCreateModal || showEditModal" class="modal-overlay" @click.self="closeModals">
        <div class="modal">
          <h2>{{ showCreateModal ? 'Tambah Lapangan Baru' : 'Edit Lapangan' }}</h2>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>Nama Lapangan</label>
              <input type="text" v-model="form.name" required />
            </div>
            <div class="form-group">
              <label>Deskripsi</label>
              <textarea v-model="form.description" rows="3"></textarea>
            </div>
            <div class="form-group">
              <label>Alamat</label>
              <textarea v-model="form.address" rows="2"></textarea>
            </div>
            <div class="form-group">
              <label>Kota</label>
              <input type="text" v-model="form.city" />
            </div>
            <div class="form-group">
              <label>Harga per Jam (Rp)</label>
              <input 
                type="text" 
                v-model="formattedPrice" 
                @input="updatePriceFromFormatted" 
                placeholder="Rp 100.000" 
                required 
              />
            </div>
            <div class="form-group">
              <label>Gambar Lapangan (opsional)</label>
              <div class="file-upload">
                <input type="file" ref="fileInput" @change="handleFileChange" accept="image/*" />
                <button type="button" @click="triggerFileInput" class="btn-upload">
                  Pilih File
                </button>
                <span v-if="selectedFileName" class="file-name">{{ selectedFileName }}</span>
              </div>
              <div v-if="imagePreview" class="image-preview">
                <img :src="imagePreview" alt="Preview" />
              </div>
            </div>
            <div class="modal-actions">
              <button type="button" class="btn-outline" @click="closeModals">Batal</button>
              <button type="submit" class="btn-primary" :disabled="loading">
                {{ loading ? 'Menyimpan...' : 'Simpan' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import Layout from '@/components/Layout.vue'
import api from '@/api'
import { formatRupiah, parseRupiah } from '@/utils/currency'

const placeholderImg = 'https://coresg-normal.trae.ai/api/ide/v1/text_to_image?prompt=volleyball%20court%20professional%20indoor&image_size=square_hd'

const courts = ref([])
const loading = ref(false)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const fileInput = ref(null)
const selectedFile = ref(null)
const selectedFileName = ref('')
const imagePreview = ref('')
const form = ref({
  name: '',
  description: '',
  address: '',
  city: '',
  price_per_hour: 0,
  image: ''
})

const formattedPrice = computed({
  get() {
    return formatRupiah(form.value.price_per_hour);
  },
  set(value) {
    form.value.price_per_hour = parseRupiah(value);
  }
})

const updatePriceFromFormatted = (e) => {
  formattedPrice.value = e.target.value;
}

const formatPrice = (num) => {
  return new Intl.NumberFormat('id-ID').format(num)
}

const fetchCourts = async () => {
  try {
    const res = await api.get('/my-courts')
    courts.value = res.data.data || []
  } catch (err) {
    console.error(err)
  }
}

const editCourt = (court) => {
  form.value = { ...court }
  imagePreview.value = court.image || ''
  showEditModal.value = true
}

const deleteCourt = async (id) => {
  if (confirm('Yakin ingin menghapus lapangan ini?')) {
    try {
      await api.delete(`/courts/${id}`)
      fetchCourts()
    } catch (err) {
      alert('Gagal menghapus lapangan')
    }
  }
}

const triggerFileInput = () => {
  fileInput.value.click();
}

const handleFileChange = (e) => {
  const file = e.target.files[0];
  if (file) {
    selectedFile.value = file;
    selectedFileName.value = file.name;
    
    // Create preview
    const reader = new FileReader();
    reader.onload = (e) => {
      imagePreview.value = e.target.result;
    };
    reader.readAsDataURL(file);
  }
}

const handleSubmit = async () => {
  loading.value = true
  try {
    const payload = { ...form.value }
    
    // If we have a selected file, we'd need to handle file upload
    // For now, just use the preview/base64 as a fallback
    if (imagePreview.value && !selectedFile.value) {
      payload.image = imagePreview.value;
    } else if (imagePreview.value) {
      payload.image = imagePreview.value;
    }

    if (showCreateModal.value) {
      await api.post('/courts', payload)
    } else {
      await api.put(`/courts/${form.value.id}`, payload)
    }
    closeModals()
    fetchCourts()
  } catch (err) {
    alert('Gagal menyimpan lapangan')
  } finally {
    loading.value = false
  }
}

const closeModals = () => {
  showCreateModal.value = false
  showEditModal.value = false
  selectedFile.value = null
  selectedFileName.value = ''
  imagePreview.value = ''
  form.value = {
    name: '',
    description: '',
    address: '',
    city: '',
    price_per_hour: 0,
    image: ''
  }
}

onMounted(() => {
  fetchCourts()
})
</script>

<style scoped>
.page {
  max-width: 100%;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.page-header h1 {
  color: #1e3a5f;
  font-weight: 800;
  font-size: 2rem;
}

.btn-primary {
  background: linear-gradient(135deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  border: none;
  padding: 0.8rem 1.5rem;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(30, 58, 95, 0.2);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(30, 58, 95, 0.3);
}

.courts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.5rem;
}

.court-card {
  background: white;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0,0,0,0.08);
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}

.court-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.12);
}

.court-image {
  height: 200px;
  overflow: hidden;
}

.court-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.court-info {
  padding: 1.5rem;
}

.court-info h3 {
  color: #1e3a5f;
  margin-bottom: 0.5rem;
  font-weight: 700;
}

.location {
  color: #64748b;
  font-size: 0.95rem;
  margin-bottom: 0.5rem;
}

.price {
  color: #ff6b35;
  font-weight: 800;
  font-size: 1.25rem;
  margin-bottom: 1rem;
}

.court-actions {
  display: flex;
  gap: 0.75rem;
}

.btn-outline {
  flex: 1;
  padding: 0.7rem;
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

.btn-danger {
  flex: 1;
  padding: 0.7rem;
  border: 2px solid #e63946;
  color: #e63946;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 700;
  transition: all 0.3s ease;
}

.btn-danger:hover {
  background: #e63946;
  color: white;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 2rem;
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal h2 {
  color: #1e3a5f;
  margin-bottom: 1.5rem;
  font-weight: 800;
}

.form-group {
  margin-bottom: 1.25rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #475569;
  font-weight: 600;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.9rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.3s ease;
  box-sizing: border-box;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #ff6b35;
  box-shadow: 0 0 0 3px rgba(255,107,53,0.1);
}

.file-upload {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.file-upload input[type="file"] {
  display: none;
}

.btn-upload {
  background: #f1f5f9;
  border: 2px dashed #cbd5e1;
  padding: 0.7rem 1.5rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s ease;
  color: #64748b;
}

.btn-upload:hover {
  border-color: #1e3a5f;
  color: #1e3a5f;
}

.file-name {
  color: #64748b;
  font-size: 0.95rem;
}

.image-preview {
  margin-top: 1rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  overflow: hidden;
}

.image-preview img {
  width: 100%;
  height: 150px;
  object-fit: cover;
}

.modal-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
}
</style>
