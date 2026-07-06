<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f8fafc;
  color: #1e293b;
}

#app {
  min-height: 100vh;
}

.page {
  max-width: 100%;
}

.page h1 {
  color: #1e3a5f;
  margin-bottom: 2rem;
  font-weight: 800;
  font-size: 2rem;
}
</style>
