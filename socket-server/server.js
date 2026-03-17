const http = require('http');
const { Server } = require('socket.io');

const server = http.createServer();

const io = new Server(server, {
  cors: {
    origin: '*',
  },
});

io.on('connection', (socket) => {
  console.log('User connected:', socket.id);

  socket.on('message', (msg) => {
    io.emit('message', msg);
  });
});

server.listen(3001, '0.0.0.0', () => {
  console.log('Socket running on http://localhost:3001');
});
