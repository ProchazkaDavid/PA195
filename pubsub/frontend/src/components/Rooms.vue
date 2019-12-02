<template>
  <div class="rooms">
    <h1>Here are all the available chat rooms:</h1>
    <ul id="rooms-list">
      <li v-for="(room, i) in chatRooms" :key="`${i}-${room}`">
        <div class="room">
          <router-link
            :to="{name: 'chat-room', params: {name: room.name}}"
          >{{ room.name }} ({{ chatRoomMembers(room.name).length }} members)</router-link>
        </div>
      </li>
    </ul>
  </div>
</template>

<style>
ul {
  list-style-type: none;
}
</style>

<script>
export default {
  name: "Rooms",
  computed: {
    chatRooms() {
      return this.$store.getters.chatRooms;
    }
  },
  methods: {
    chatRoomMembers(roomName) {
      return this.$store.getters.chatRoomMembers(roomName);
    }
  }
};
</script>
