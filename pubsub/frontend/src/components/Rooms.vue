<template>
  <div class="rooms">
    <h1>Here are all the available chat rooms:</h1>
    <ul id="rooms-list">
      <li v-for="room in chatRooms" v-bind:key="room">
        <div class="room">
          <router-link
            :to="{name: 'chat-room', params: {id: room.id}}"
          >{{ room.name }} ({{ chatRoomMembers(room.id).length }} members)</router-link>
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
    chatRoomMembers(roomId) {
      return this.$store.getters.chatRoomMembers(roomId);
    }
  }
};
</script>
