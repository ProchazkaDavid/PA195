<template>
  <div class="rooms">
    <Rooms />
    <CreateRoom v-if="isLoggedIn()" />
  </div>
</template>

<script>
import Rooms from "@/components/Rooms.vue";
import CreateRoom from "@/components/CreateRoom.vue";

export default {
  name: "rooms",
  components: {
    Rooms,
    CreateRoom
  },
  methods: {
    isLoggedIn() {
      return this.$store.state.loggedIn;
    }
  }
};
</script>
