<script setup lang="ts">
import { CategoryVO } from '@/api/user'
import { computed } from 'vue'
import ArticleNavContent from '@/views/article/components/ArticleNavContent.vue'

// 自己引用自己（递归组件的关键）
defineOptions({ name: 'ArticleNav' })

const props = defineProps<{
  modelValue: CategoryVO[]
}>()

const emit = defineEmits<{(e: 'update:modelValue', value: CategoryVO[]): void
  (e: 'add', item: CategoryVO): void
  (e: 'edit', item: CategoryVO): void
  (e: 'delete', item: CategoryVO): void
}>()

// 用 computed 包装 v-model
const navList = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

function onAdd (item: CategoryVO) {
  emit('add', item)
}
function onEdit (item: CategoryVO) {
  emit('edit', item)
}
function onDelete (item: CategoryVO) {
  emit('delete', item)
}
</script>

<template>
  <div class="vxe-menu">
    <div class="vxe-menu--item-list">
      <ArticleNavContent @add="onAdd" @edit="onEdit" @delete="onDelete" :navList="item" v-for="item in navList" :key="item.id" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.vxe-menu {
  overflow-y: auto;
}
</style>
