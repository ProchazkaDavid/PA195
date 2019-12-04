import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

export const socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("SOCKET: Attempting Connection...");
  socket.onopen = () => console.log("SOCKET: Successfully Connected");
  socket.onmessage = msg => {
    const message = JSON.parse(msg.data);
    switch (message.event) {
      case "create_room":
        store.commit("ADD_ROOM", {
          socket: false,
          name: message.room,
          messages: []
        });
        break;
      case "send_msg":
        store.commit("ADD_MSG", message);
        break;
      case "fetch_all":
        for (const room of message.rooms) {
          store.commit("ADD_ROOM", {
            socket: false,
            name: room.room,
            messages: room.msgs
          });
        }
        break;
      default:
        console.log(`SOCKET: unknown event (message: ${message})`);
    }
  };
  socket.onclose = event =>
    console.log("SOCKET: Socket Closed Connection: ", event);
  socket.onerror = error => console.log("SOCKET: Socket Error: ", error);
};

Vue.config.productionTip = false;

new Vue({
  mode: "history",
  router,
  store,
  render: h => h(App),
  mounted() {
    connect();
  }
}).$mount("#app");
