<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <div class="introduction">
      <p>
        This is a demo app showcasing Redis' caching capabilities.
        You can register, log in and join the chatrooms!
        Check out the implementation at
        <a
          href="https://github.com/ProchazkaDavid/pa195"
        >David's GitHub repository</a>.
      </p>
    </div>
  </div>
</template>

<style scoped>
a {
  color: aliceblue;
}

.introduction {
  width: 60%;
  text-align: center;
  margin: 0 auto;
}
</style>

<script>
export default {
  name: "home"
};
</script>
