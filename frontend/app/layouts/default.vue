<script setup lang="ts">
import AppSidebar from '@/components/widgets/app-sidebar/AppSidebar.vue';
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
</script>

<template>
  <SidebarProvider>
    <AppSidebar>
      <slot />
    </AppSidebar>
  </SidebarProvider>
</template>
