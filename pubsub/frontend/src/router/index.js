import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Rooms from "../views/Rooms.vue";
import ChatRoom from "../views/ChatRoom.vue";
import Login from "../views/Login.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/rooms",
    name: "rooms",
    component: Rooms
  },
  {
    path: "/chat-room/:name",
    name: "chat-room",
    component: ChatRoom
  },
  {
    path: "/login",
    name: "login",
    component: Login
  }
];

const router = new VueRouter({
  routes
});

export default router;
