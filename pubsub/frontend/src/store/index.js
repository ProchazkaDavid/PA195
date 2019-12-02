import Vue from "vue";
import Vuex from "vuex";
import moment from "moment";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loggedIn: false,
    currentUser: undefined,
    rooms: [
      {
        id: 1,
        name: "ChatRoom rawr xD",
        messages: [
          {
            sender: "Pepsi",
            text: "testovaci zprava",
            date: moment().format("LT")
          },
          {
            sender: "Pepsi",
            text: "testovaci zprava",
            date: moment().format("LT")
          },
          {
            sender: "Pepsi",
            text: "testovaci zprava",
            date: moment().format("LT")
          }
        ]
      },
      {
        id: 2,
        name: "Anime club",
        messages: [
          {
            sender: "Pepsi",
            text: "testovaci zprava",
            date: moment().format("LT")
          },
          {
            sender: "Pepsi",
            text: "testovaci zprava",
            date: moment().format("LT")
          }
        ]
      }
    ],
    users: [
      { id: 1, nickname: "pepsi" },
      { id: 2, nickname: "deiv191" }
    ]
  },
  mutations: {
    SOCKET_CONNECT: (state, status) => {
      console.log(status);
      state.connect = true;
    },
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
        .find(room => room.id === data.roomId)
        .messages.push({
          text: data.msg,
          sender: state.currentUser.nickname,
          date: data.date
        });
    }
  },
  actions: {
    logIn(context, user) {
      context.commit("LOG_IN", user);
    },
    logOut(context) {
      context.commit("LOG_OUT");
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
      return roomId => state.rooms.find(room => room.id === roomId).name;
    },
    messages: state => {
      return roomId => state.rooms.find(room => room.id === roomId).messages;
    },
    chatRoomMembers: state => {
      return roomId => [
        ...new Set(
          state.rooms
            .find(room => room.id === roomId)
            .messages.map(msg => msg.sender)
        )
      ];
    }
  },
  methods: {},
  modules: {}
});
