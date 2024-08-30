<template>
  <div id="app">
    <nav>
      <router-link to="/">首頁</router-link>
      <router-link to="/activities">查看活動</router-link>
    </nav>
    <router-view />
  </div>
</template>


<script>


export default {
  name: 'App',
  components: {
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', 'Helvetica', 'Arial', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

nav {
  background-color: #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 20px 0;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
}

nav a {
  text-decoration: none;
  color: #2c3e50;
  margin: 0 15px;
  font-size: 16px;
  font-weight: 500;
  transition: color 0.3s ease;
}

nav a:hover {
  color: #0056b3;
}

.router-link-exact-active {
  color: #0056b3;
  border-bottom: 2px solid #0056b3;
}
</style>
