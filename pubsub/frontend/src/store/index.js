import Vue from "vue";
import Vuex from "vuex";
import { socket } from "../main";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loggedIn: false,
    currentUser: undefined,
    rooms: []
  },
  mutations: {
    LOG_IN(state, nickname) {
      state.currentUser = { nickname };
      state.loggedIn = true;
    },
    LOG_OUT: state => {
      state.loggedIn = false;
      state.currentUser = undefined;
    },
    SEND_MSG(state, data) {
      socket.send(
        JSON.stringify({
          event: "send_msg",
          room: data.roomName,
          date: data.date,
          sender: state.currentUser.nickname,
          text: data.msg
        })
      );
      state.rooms
        .find(room => room.name === data.roomName)
        .messages.push({
          text: data.msg,
          sender: state.currentUser.nickname,
          date: data.date
        });
    },
    ADD_MSG(state, data) {
      state.rooms
        .find(room => room.name === data.room)
        .messages.push({
          text: data.text,
          sender: data.sender,
          date: data.date
        });
    },
    ADD_ROOM(state, data) {
      if (data.socket) {
        socket.send(
          JSON.stringify({
            event: "create_room",
            room: data.name
          })
        );
      }
      state.rooms.push({
        name: data.name,
        messages: data.messages
      });
    }
  },
  getters: {
    currentUser: state => {
      return state.currentUser;
    },
    chatRooms: state => {
      return state.rooms;
    },
    chatRoomName: state => {
      return roomName => state.rooms.find(room => room.name === roomName).name;
    },
    messages: state => {
      return roomName =>
        state.rooms.find(room => room.name === roomName).messages;
    },
    chatRoomMembers: state => {
      return roomName => [
        ...new Set(
          state.rooms
            .find(room => room.name === roomName)
            .messages.map(msg => msg.sender)
        )
      ];
    }
  }
});
