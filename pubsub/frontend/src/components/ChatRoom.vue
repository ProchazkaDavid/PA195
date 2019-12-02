<template>
  <div class="chat-room">
    <h1>Welcome to {{ name }}</h1>
    <p>Members: {{ members }}</p>
    <ul id="messages-list">
      <li v-for="(message, i) in messages" :key="`${i}-${message}`">
        <div class="message">
          <span class="sender">{{ message.sender }}</span>
          <!-- fujky! -->
          <span class="date">&nbsp; ({{ message.date }})</span>:
          <!-- end-fujky! -->
          <span class="text">{{ message.text }}</span>
        </div>
      </li>
    </ul>
    <div class="input" v-if="isLoggedIn()">
      <form @submit="send">
        <input type="text" id="message" v-model="message" placeholder="write some cute msg" />
        <button type="send">Send</button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.sender {
  color: gray;
}

.date {
  color: gray;
}
</style>

<script>
import moment from "moment";

export default {
  name: "Chat-Room",
  methods: {
    isLoggedIn() {
      return this.$store.state.loggedIn;
    },
    send() {
      this.$store.commit("SEND_MSG", {
        msg: this.message,
        roomId: this.$route.params.id,
        date: moment().format("LT")
      });
      this.message = "";
    }
  },
  data() {
    return {
      messages: this.$store.getters.messages(this.$route.params.id)
    };
  },
  computed: {
    name() {
      return this.$store.getters.chatRoomName(this.$route.params.id);
    },
    members() {
      return this.$store.getters
        .chatRoomMembers(this.$route.params.id)
        .join(", ");
    }
  }
};
</script>
