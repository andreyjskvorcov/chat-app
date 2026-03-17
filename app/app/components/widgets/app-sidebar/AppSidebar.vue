<script setup lang="ts">
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from '@/components/ui/sidebar';

import { MessageCircle } from 'lucide-vue-next';

const chats = [
  {
    id: 1,
    name: 'Maxim',
    lastMessage: 'Привет, как дела?',
    time: '12:45',
    online: true,
  },
  {
    id: 2,
    name: 'Серега Мартынов',
    lastMessage: 'Привет',
    time: '11:20',
    online: false,
  },
  {
    id: 3,
    name: 'Володя',
    lastMessage: 'Го кс',
    time: '11:20',
    online: false,
  },
  {
    id: 4,
    name: 'Артем Чистяков',
    lastMessage: 'Го кс',
    time: '11:20',
    online: false,
  },
  {
    id: 5,
    name: 'Евгений Простак Джексон Брага',
    lastMessage: 'Кофе?',
    time: '11:20',
    online: false,
  },
  {
    id: 6,
    name: 'Артем Куклин',
    lastMessage: 'Здарова бразе',
    time: '11:20',
    online: false,
  },
  {
    id: 7,
    name: 'Stolenair',
    lastMessage: 'Где взять ты говоришь re?',
    time: '11:20',
    online: false,
  },
  {
    id: 8,
    name: 'Sega Mega',
    lastMessage: 'Здаров го д2?',
    time: '11:20',
    online: false,
  },
];
</script>

<template>
  <SidebarProvider>
    <Sidebar>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg">
              <div
                class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
              >
                <GalleryVerticalEnd class="size-4" />
              </div>
              <div class="grid flex-1 text-left text-sm leading-tight">
                <span class="truncate font-semibold">Андрей Скворцов</span>
                <span class="truncate text-xs">Set emoji Status</span>
              </div>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Чаты</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem v-for="chat in chats" :key="chat.id">
                <SidebarMenuButton as-child>
                  <a href="#" class="flex items-center gap-3 w-full h-12">
                    <!-- Аватар -->
                    <div class="relative">
                      <div
                        class="w-9 h-9 rounded-full bg-muted flex items-center justify-center"
                      >
                        <MessageCircle class="w-4 h-4" />
                      </div>

                      <!-- Онлайн статус -->
                      <span
                        v-if="chat.online"
                        class="absolute bottom-0 right-0 w-2.5 h-2.5 bg-green-500 rounded-full border border-background"
                      />
                    </div>

                    <!-- Текст -->
                    <div class="flex-1 overflow-hidden text-left">
                      <div class="flex justify-between text-sm">
                        <span class="font-medium truncate">
                          {{ chat.name }}
                        </span>
                        <span class="text-xs text-muted-foreground">
                          {{ chat.time }}
                        </span>
                      </div>

                      <div class="text-xs text-muted-foreground truncate">
                        {{ chat.lastMessage }}
                      </div>
                    </div>
                  </a>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter />
      <SidebarRail />
    </Sidebar>
    <SidebarInset>
      <header
        class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12"
      >
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
        </div>
      </header>
      <div class="flex flex-1 flex-col gap-4 p-4 pt-0">
        <slot />
      </div>
    </SidebarInset>
  </SidebarProvider>
</template>
