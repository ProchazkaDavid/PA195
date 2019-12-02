<template>
  <div class="create-room">
    <form @submit="addRoom">
      <label for="room">Create a room:&nbsp;</label>
      <input type="text" id="room" v-model="room" placeholder="room name" />
      <button type="send">Create</button>
    </form>
  </div>
</template>

<script>
export default {
  name: "Create-Room",
  methods: {
    addRoom() {
      this.$store.commit("ADD_ROOM", this.room);
      this.room = "";
    }
  }
};
</script>