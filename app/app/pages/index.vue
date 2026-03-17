<script setup lang="ts">
const messages = [
  {
    id: 1,
    text: 'Привет!',
    time: '12:01',
    user: {
      name: 'Андрей Скворцов',
      avatar: 'https://i.pravatar.cc/100?img=1',
      isMe: true,
    },
  },
  {
    id: 2,
    text: 'Привет 👋',
    time: '12:02',
    user: {
      name: 'Сергей',
      avatar: 'https://i.pravatar.cc/100?img=3',
      isMe: false,
    },
  },
];
</script>

<template>
  <div class="flex h-full flex-col border rounded-xl bg-background">
    <!-- Header -->
    <div class="border-b px-4 py-3 font-semibold">Дотики</div>

    <!-- Messages -->
    <div class="flex-1 overflow-y-auto p-4 space-y-4">
      <div
        v-for="msg in messages"
        :key="msg.id"
        class="flex gap-3"
        :class="msg.user.isMe ? 'justify-end' : 'justify-start'"
      >
        <!-- Аватар -->
        <img :src="msg.user.avatar" class="w-8 h-8 rounded-full object-cover" />

        <!-- Контент -->
        <div class="max-w-[70%]">
          <!-- Сообщение -->
          <div
            class="rounded-xl px-4 py-2 text-sm shadow"
            :class="msg.user.isMe ? 'bg-blue-700' : 'bg-muted'"
          >
            <!-- Ник -->
            <div
              class="text-xs mb-1 text-orange-300"
              :class="msg.user.isMe ? 'text-right' : 'text-left'"
            >
              {{ msg.user.name }}
            </div>

            <div>{{ msg.text }}</div>

            <!-- Время -->
            <div class="mt-1 text-xs opacity-60 text-right">
              {{ msg.time }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Input -->
    <div class="border-t p-3 flex gap-2">
      <input
        type="text"
        placeholder="Написать сообщение..."
        class="flex-1 rounded-md border px-3 py-2 text-sm bg-background"
      />
      <button
        class="px-4 py-2 rounded-md bg-primary text-primary-foreground text-sm"
      >
        Отправить
      </button>
    </div>
  </div>
</template>
