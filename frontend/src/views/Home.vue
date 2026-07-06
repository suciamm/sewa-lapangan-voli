<template>
  <div class="home">
    <section class="hero">
      <h1>Selamat Datang di Sewa Lapangan Voli</h1>
      <p>Sewa lapangan voli berkualitas dengan harga terjangkau</p>
      <router-link to="/courts" class="btn-primary">Lihat Lapangan</router-link>
    </section>
    <section class="features">
      <div class="feature">
        <h3>Lapangan Berkualitas</h3>
        <p>Lapangan dengan standar internasional</p>
      </div>
      <div class="feature">
        <h3>Pemesanan Mudah</h3>
        <p>Pesan lapangan secara online dengan mudah</p>
      </div>
      <div class="feature">
        <h3>Harga Terjangkau</h3>
        <p>Harga sewa yang kompetitif</p>
      </div>
    </section>
  </div>
</template>

<script setup>
</script>

<style scoped>
.hero {
  text-align: center;
  padding: 5rem 2rem;
  background: linear-gradient(135deg, #1e3a5f 0%, #0f2744 100%);
  color: white;
  border-radius: 16px;
  margin-bottom: 3rem;
  box-shadow: 0 10px 40px rgba(30, 58, 95, 0.2);
}

.hero h1 {
  font-size: 3rem;
  margin-bottom: 1rem;
  font-weight: 800;
  letter-spacing: 0.5px;
}

.hero p {
  font-size: 1.25rem;
  margin-bottom: 2.5rem;
  color: rgba(255,255,255,0.9);
  font-weight: 400;
}

.btn-primary {
  background: linear-gradient(135deg, #ff6b35 0%, #e63946 100%);
  color: white;
  padding: 1rem 2.5rem;
  border-radius: 10px;
  text-decoration: none;
  font-weight: 700;
  font-size: 1.1rem;
  transition: all 0.3s ease;
  display: inline-block;
  box-shadow: 0 4px 15px rgba(230, 57, 70, 0.3);
}

.btn-primary:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(230, 57, 70, 0.4);
}

.features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
}

.feature {
  background: white;
  padding: 2.5rem 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.08);
  text-align: center;
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.feature:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.12);
  border-color: #ff6b35;
}

.feature h3 {
  color: #1e3a5f;
  margin-bottom: 0.75rem;
  font-size: 1.35rem;
  font-weight: 700;
}

.feature p {
  color: #64748b;
  line-height: 1.6;
}
</style>
