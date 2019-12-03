<template>
  <div class="rooms">
    <Rooms />
    <CreateChatRoom v-if="isLoggedIn" />
  </div>
</template>

<script>
import Rooms from "@/components/Rooms.vue";
import CreateChatRoom from "@/components/CreateChatRoom.vue";

export default {
  name: "rooms",
  computed: {
    isLoggedIn() {
      return this.$store.state.loggedIn;
    }
  },
  components: {
    Rooms,
    CreateChatRoom
  }
};
</script>
