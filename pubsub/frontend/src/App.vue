<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link>&nbsp;|
      <router-link to="/rooms">Chat Rooms</router-link>&nbsp;|
      <router-link to="/login">Login</router-link>
      <span class="welcome" v-if="isLoggedIn()">&nbsp;Welcome, {{ getUser().nickname }}</span>
      <span class="welcome" v-else></span>
    </div>
    <router-view />
  </div>
</template>

<style>
body {
  background-color: #303031;
}

a {
  color: aliceblue;
}
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #e2e6eb;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #e2e6eb;
}

#nav a.router-link-exact-active {
  color: #a5a5a5;
}
</style>

<script>
export default {
  name: "login",
  methods: {
    isLoggedIn() {
      return this.$store.state.loggedIn;
    },
    getUser() {
      return this.$store.state.currentUser;
    }
  }
};
</script>