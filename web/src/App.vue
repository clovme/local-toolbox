<template>
  <div class="app-container">
    <RouterView />
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { useAppStore } from '@/store/app'

const appStore = useAppStore()

function updateTheme (e) {
  if (e.matches) {
    appStore.setTheme('dark')
  } else {
    appStore.setTheme('light')
  }
}

appStore.setEnums()

onMounted(() => {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  mediaQuery.addEventListener('change', updateTheme)
})
</script>
