<script setup lang="ts">
import { CategoryVO } from '@/api/user'
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'

// 自己引用自己（递归组件的关键）
defineOptions({ name: 'ArticleNavContent' })

const props = defineProps<{
  level?: number,
  navList: CategoryVO
}>()

const emit = defineEmits<{(e: 'add', item: CategoryVO): void
  (e: 'edit', item: CategoryVO): void
  (e: 'delete', item: CategoryVO): void
  (e: 'up', item: CategoryVO): void
  (e: 'down', item: CategoryVO): void
}>()

const route = useRoute()
const level = computed(() => props.level ?? 1)

// 判断当前节点或者子节点是否 active
const isActive = computed(() => {
  // 当前节点 active
  if (String(props.navList.id) === route.query.id) return true

  // 子节点 active
  if (props.navList.children?.length) {
    return props.navList.children.some(child => checkActive(child))
  }
  return false
})

function checkActive (node: CategoryVO): boolean {
  if (String(node.id) === route.query.id) return true
  if (node.children?.length) {
    return node.children.some(child => checkActive(child))
  }
  return false
}

function collapse (node: CategoryVO) {
  node.isExpand = !node.isExpand
}

const menuItemOption = ref<HTMLUListElement | null>(null)
const showTimer = ref<number | null>(null)
const hideTimer = ref<number | null>(null)

function onMenuItemMouseenterOptions (event: MouseEvent) {
  const target = event.currentTarget as HTMLElement
  const rect = target.getBoundingClientRect()

  clearTimeout(hideTimer.value as number)
  clearTimeout(showTimer.value as number)

  showTimer.value = window.setTimeout(() => {
    if (!menuItemOption.value) {
      const ul = document.createElement('ul')
      ul.className = 'menu-item-options'
      ul.style.position = 'fixed'
      ul.style.padding = '5px'

      const add = document.createElement('li')
      add.style.color = '#409eff'
      add.innerHTML = '添加新分类'
      add.addEventListener('click', (e) => {
        e.stopPropagation()
        emit('add', props.navList)
        removeMenuOption()
      })

      if (props.navList.name !== 'all' && props.navList.name !== 'default') {
        const edit = document.createElement('li')
        edit.style.color = '#67c23a'
        edit.innerHTML = `编辑 [<b>${props.navList.title}</b>] 信息`
        edit.addEventListener('click', (e) => {
          e.stopPropagation()
          emit('edit', props.navList)
          removeMenuOption()
        })

        const del = document.createElement('li')
        del.style.color = '#f56c6c'
        del.innerHTML = `删除 [<b>${props.navList.title}</b>] 分类`
        del.addEventListener('click', (e) => {
          e.stopPropagation()
          emit('delete', props.navList)
          removeMenuOption()
        })

        ul.append(add, edit, del)
      } else {
        ul.append(add)
      }

      ul.addEventListener('mouseenter', () => {
        clearTimeout(hideTimer.value as number)
      })
      ul.addEventListener('mouseleave', () => {
        hideTimer.value = window.setTimeout(() => {
          removeMenuOption()
        }, 100)
      })

      menuItemOption.value = ul
    }

    const el = menuItemOption.value!
    document.body.appendChild(el) // 先 append 获取真实高度

    const elHeight = el.offsetHeight
    const viewportHeight = window.innerHeight

    let top = rect.top - elHeight / 2

    // 纵向边界判断
    if (top < 0) {
      top = 5 // 顶部留 5px 间距
    } else if (top + elHeight > viewportHeight) {
      top = viewportHeight - elHeight - 5 // 底部留 5px 间距
    }

    el.style.left = `${rect.right}px`
    el.style.top = `${top}px`
  }, 200)
}

function onMenuItemMouseleaveOptions () {
  clearTimeout(showTimer.value as number)
  hideTimer.value = window.setTimeout(() => {
    removeMenuOption()
  }, 100)
}

function removeMenuOption () {
  if (menuItemOption.value && menuItemOption.value.parentNode) {
    menuItemOption.value.parentNode.removeChild(menuItemOption.value)
    menuItemOption.value = null
  }
}

function onAdd (item: CategoryVO) {
  emit('add', item)
}
function onEdit (item: CategoryVO) {
  emit('edit', item)
}
function onDelete (item: CategoryVO) {
  emit('delete', item)
}
function onUp (item: CategoryVO) {
  emit('up', item)
}
function onDown (item: CategoryVO) {
  emit('down', item)
}
</script>

<template>
  <div class="vxe-menu--item-wrapper" :class="{['vxe-menu--item-level' + level]: true, 'is--expand': navList.isExpand, 'is--exact-active': route.query.cid===`${navList.id}`, 'is--active': isActive}">
    <router-link :to="{name: 'Article', query: {cid: navList.id, type: navList.name, sort: navList.docSort}}" class="vxe-menu--item-link">
      <div :class="{'vxe-menu--item-link-icon': true, 'vxe-menu--item-link-icon--is-option': navList.name !== 'all' && navList.name !== 'default' }">
        <vxe-icon name="file-markdown"></vxe-icon>
        <vxe-icon name="information"></vxe-icon>
      </div>
      <div class="vxe-menu--item-link-title">
        <div class="vxe-menu--item-link-option-title">{{ navList.title }}</div>
        <div :class="{ 'vxe-menu--item-link-option-children': !navList.children?.length }">
          <div @mouseenter="onMenuItemMouseenterOptions" @mouseleave="onMenuItemMouseleaveOptions" class="vxe-menu--item-link-option">
            <span>{{ navList.articleCount < 100 ? navList.articleCount : '99+' }}</span>
            <vxe-icon name="ellipsis-v"></vxe-icon>
          </div>
        </div>
      </div>
      <div v-if="navList.children?.length" class="vxe-menu--item-link-collapse" @click.stop.prevent="collapse(navList)">
        <vxe-icon v-if="navList.isExpand" name="arrow-down"></vxe-icon>
        <vxe-icon v-else name="arrow-left"></vxe-icon>
      </div>
    </router-link>
    <div v-if="navList.children?.length" class="vxe-menu--item-group">
      <ArticleNavContent @up="onUp" @down="onDown" @add="onAdd" @edit="onEdit" @delete="onDelete" :navList="item" v-for="item in navList.children" :key="item.id" :level="level + 1" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.vxe-menu--item-link {
  -webkit-user-drag: none;
  user-select: none !important;

  .vxe-menu--item-link-collapse {
    cursor: pointer;
  }

  .vxe-menu--item-link-icon {
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;

    i {
      font-size: 18px;
    }

    i:last-child {
      display: none;
    }
  }

  .vxe-menu--item-link-title {
    display: flex;

    .vxe-menu--item-link-option-title {
      flex: 1;
      display: flex;
      align-items: center;
    }

    .vxe-menu--item-link-option {
      width: 30px;
      height: 25px;
      display: flex;
      align-items: center;
      justify-content: center;
      position: relative;

      span {
        line-height: 15.41px;
      }

      i {
        display: none;
      }
    }

    .vxe-menu--item-link-option-children {
      margin-right: 25px;
    }
  }

  &:hover {
    .vxe-menu--item-link-icon--is-option {
      i:first-child {
        display: none;
      }

      i:last-child {
        display: inline-block;
      }
    }

    .vxe-menu--item-link-title {
      .vxe-menu--item-link-option {
        span {
          display: none;
        }

        i {
          display: inline-block;
        }
      }
    }
  }
}
</style>
