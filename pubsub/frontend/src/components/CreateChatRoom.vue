<template>
  <div class="create-room">
    <form @submit="create">
      <label for="newRoomName">Create a room:&nbsp;</label>
      <input type="text" id="newRoomName" v-model="newRoomName" placeholder="enter a room name" />
      <button type="create">Create</button>
    </form>
  </div>
</template>


<script>
export default {
  name: "CreateChatRoom",
  beforeMount() {
    this.newRoomName = "";
  },
  methods: {
    create() {
      this.$store.commit("ADD_ROOM", {
        name: this.newRoomName,
        messages: [],
        socket: true
      });
      this.newRoomName = "";
    }
  }
};
</script>