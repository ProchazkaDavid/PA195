<template>
  <div class="login" v-if="!isLoggedIn()">
    <h1>Log in</h1>
    <div id="form">
      <form @submit="logIn">
        <input type="text" id="nickname" v-model="nickname" placeholder="nickname" />

        <input type="password" id="password" v-model="password" placeholder="password" />
        <button type="submit">Log in</button>
      </form>
    </div>
  </div>
  <div class="login" v-else>
    You're logged in!
    <button v-on:click="logOut()">Log Out</button>
  </div>
</template>

<script>
export default {
  name: "Login",
  methods: {
    isLoggedIn() {
      return this.$store.state.loggedIn;
    },
    logIn() {
      if (this.nickname === "" || this.password === "") {
        alert("nickname and password can not be empty!");
      } else {
        this.$store.commit("LOG_IN", this.nickname);
      }
    },
    logOut() {
      this.$store.commit("LOG_OUT");
    }
  }
};
</script>
