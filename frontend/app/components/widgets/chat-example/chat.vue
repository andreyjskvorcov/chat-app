<script setup>
const { messages, connect, send } = useChat();

const input = ref('');

onMounted(() => {
  connect();
});

const sendMessage = () => {
  send({
    username: 'user1',
    text: input.value,
  });

  input.value = '';
};
</script>

<template>
  <div>
    <h1>Chat</h1>

    <div v-for="(msg, i) in messages" :key="i">
      <b>{{ msg.username }}:</b> {{ msg.text }}
    </div>

    <input v-model="input" placeholder="Message..." />
    <button @click="sendMessage">Send</button>
  </div>
</template>
