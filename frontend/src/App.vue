<template>
  <div id="app">
    <nav>
      <div class="nav-container">
        <router-link to="/" exact>首頁</router-link>
        <router-link to="/activities">查看活動</router-link>
        <router-link to="/report">回報問題</router-link>
        <router-link to="/detect">測試</router-link>
        <router-link to="/data-detect">資料檢測</router-link>
      </div>
    </nav>
    <div class="content">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  }
}
</script>

<style scoped>
#app {
  font-family: 'Avenir', 'Helvetica', 'Arial', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  background-color: #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 10px 0;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
}

.nav-container {
  display: flex;
  justify-content: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

nav a {
  text-decoration: none;
  color: #2c3e50;
  margin: 0 20px;
  font-size: 18px;
  font-weight: 600;
  padding: 10px 0;
  transition: color 0.3s ease, border-bottom 0.3s ease;
  border-bottom: 2px solid transparent;
}

nav a:hover {
  color: #0056b3;
  border-bottom: 2px solid #0056b3;
}

.router-link-exact-active {
  color: #0056b3;
  border-bottom: 2px solid #0056b3;
}

.content {
  margin-top: 80px;
  padding: 20px;
}
</style>
