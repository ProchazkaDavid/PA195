import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loggedIn: false,
    currentUser: undefined,
    rooms: [
      {
        name: "ChatRoom rawr xD",
        messages: []
      },
      {
        name: "Anime club",
        messages: []
      }
    ]
  },
  mutations: {
    LOG_IN(state, nickname) {
      state.currentUser = {
        id: 1000,
        nickname
      };
      state.loggedIn = true;
    },
    LOG_OUT: state => {
      state.loggedIn = false;
      state.currentUser = undefined;
    },
    SEND_MSG(state, data) {
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
