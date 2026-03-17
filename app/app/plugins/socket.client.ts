import { io } from 'socket.io-client';

export default defineNuxtPlugin(() => {
  const socket = io('https://b16a73f475b1.vps.myjino.ru:3001');

  socket.on('connect', () => {
    console.log('SOCKET CONNECTED', socket.id);
  });

  return {
    provide: { socket },
  };
});
