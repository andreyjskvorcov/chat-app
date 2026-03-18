import { ref } from 'vue';

const messages = ref([]);
let socket = null;

export const useChat = () => {
  const connect = () => {
    socket = new WebSocket('ws://localhost:8080/ws');

    socket.onmessage = (event) => {
      messages.value.push(JSON.parse(event.data));
    };
  };

  const send = (msg) => {
    socket.send(JSON.stringify(msg));
  };

  return {
    messages,
    connect,
    send,
  };
};
