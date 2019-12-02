import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

const socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("Sock: Attempting Connection...");
  socket.onopen = () => console.log("Sock: Successfully Connected");
  socket.onmessage = msg => store.commit("ADD_MSG", JSON.parse(msg.data));
  socket.onclose = event =>
    console.log("Sock: Socket Closed Connection: ", event);
  socket.onerror = error => console.log("Sock: Socket Error: ", error);
};

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
  mounted() {
    connect();
  }
}).$mount("#app");
