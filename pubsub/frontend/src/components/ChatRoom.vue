<template>
  <div class="chat-room">
    <h1>Welcome to {{ name }}</h1>
    <p>Members: {{ members }}</p>
    <ul id="messages-list">
      <li v-for="(message, i) in getMessages()" :key="`${i}-${message}`">
        <div class="message">
          <span class="sender">{{message.sender}}</span>
          <span class="date">&nbsp; ({{ message.date }})</span>:
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
    getMessages() {
      return this.$store.getters.messages(this.$route.params.name);
    },
    send() {
      this.$store.commit("SEND_MSG", {
        msg: this.message,
        roomName: this.$route.params.name,
        date: moment().format("LT")
      });
      this.message = "";
    }
  },
  // data() {
  //  return {
  //    messages: this.$store.getters.messages(this.$route.params.name)
  //  };
  // },
  computed: {
    name() {
      return this.$store.getters.chatRoomName(this.$route.params.name);
    },
    members() {
      return this.$store.getters
        .chatRoomMembers(this.$route.params.name)
        .join(", ");
    }
  }
};
</script>
