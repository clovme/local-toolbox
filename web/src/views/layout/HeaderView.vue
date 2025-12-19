<script lang="ts" setup>
import { ref, computed } from 'vue'
import { VxeGlobalI18nLocale, VxePulldownEvents } from 'vxe-pc-ui'
import { useAppStore } from '@/store/app'
import router from '@/router'
import { useRoute } from 'vue-router'

const route = useRoute()
const appStore = useAppStore()

const langPullList = ref([
  { label: '中文', value: 'zh-CN' },
  { label: '英文', value: 'en-US' }
])

const langLabel = computed(() => {
  const item = langPullList.value.find(item => item.value === appStore.language)
  return item ? item.label : appStore.language
})

const currTheme = computed({
  get () {
    return appStore.theme
  },
  set (name) {
    appStore.setTheme(name)
  }
})

const routers = computed(() => {
  return router.getRoutes().filter(route => route.meta.isShow)
})

const currPrimaryColor = computed({
  get () {
    return appStore.primaryColor
  },
  set (color) {
    appStore.setPrimaryColor(color || '')
  }
})

const currCompSize = computed({
  get () {
    return appStore.componentsSize
  },
  set (size) {
    appStore.setComponentsSize(size)
  }
})

const colorList = ref([
  '#409eff', '#29D2F8', '#31FC49', '#3FF2B3', '#B52DFE', '#FC3243', '#FA3077', '#D1FC44', '#FEE529', '#FA9A2C'
])

const sizeOptions = ref([
  { label: '默认', value: '' },
  { label: '中', value: 'medium' },
  { label: '小', value: 'small' },
  { label: '迷你', value: 'mini' }
])

const langOptionClickEvent: VxePulldownEvents.OptionClick = ({ option }) => {
  appStore.setLanguage(option.value as VxeGlobalI18nLocale)
}
</script>

<template>
  <div class="page-header">
    <div class="header-left">
      <div class="aside-logo">
        <img class="logo-img" draggable="false" :src="`/assets/icon-${appStore.icon}.png`" alt="logo" />
        <vxe-text class="logo-title">{{ appStore.webTitle }}</vxe-text>
      </div>
      <template v-for="(r, index) in routers">
        <vxe-link :key="index" :underline="false" :router-link="{path: r.path}" v-if="route.name !== r.name">{{ r.meta.title }}</vxe-link>
      </template>
    </div>
    <div class="header-right">
      <span class="right-item">
        <vxe-switch class="right-item-comp" v-model="currTheme" size="mini" open-value="light" open-label="白天" close-value="dark" close-label="夜间"></vxe-switch>
      </span>

      <span class="right-item">
        <vxe-color-picker class="switch-primary-color" v-model="currPrimaryColor" :colors="colorList" size="mini"></vxe-color-picker>
      </span>

      <span class="right-item">
        <vxe-radio-group class="switch-size" v-model="currCompSize" :options="sizeOptions" type="button" size="mini"></vxe-radio-group>
      </span>

      <span class="right-item">
        <vxe-pulldown :options="langPullList" trigger="click" class="right-item-comp" show-popup-shadow transfer  @option-click="langOptionClickEvent">
          <vxe-button mode="text" icon="vxe-icon-language-switch" :content="langLabel"></vxe-button>
        </vxe-pulldown>
      </span>

      <span>
        <div class="user-avatar">
          <img class="user-picture" draggable="false" src="@/assets/default-picture.gif">
        </div>
      </span>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.page-header {
  display: flex;
  flex-direction: row;
  align-items: center;
  height: 50px;
  padding: 0 16px;
  border-bottom: 1px solid var(--page-layout-border-color);

  .header-left {
    gap: 5px;
    flex-grow: 1;
    display: flex;
    align-items: center;

    .aside-logo {
      display: flex;
      flex-direction: row;
      align-items: center;
      flex-shrink: 0;
      padding: 8px 0 8px 16px;
      user-select: none;

      .logo-img {
        display: block;
        width: 30px;
        height: 30px;
      }
      .logo-title {
        padding-left: 8px;
        font-weight: 700;
        font-size: 18px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        user-select: none;
        -webkit-user-drag: none;
      }
    }
  }

  .header-right {
    display: flex;
    flex-direction: row;
    flex-shrink: 0;
    align-items: center;
  }

  .right-item {
    cursor: pointer;
    margin-left: 24px;
  }
  .right-item-title {
    vertical-align: middle;
  }

  .right-item-comp {
    vertical-align: middle;
  }

  .user-avatar {
    display: inline-flex;
    flex-direction: row;
    align-items: center;
    cursor: pointer;
  }

  .user-picture {
    width: 35px;
    height: 35px;
    margin: 0 2px;
  }

  .collapseBtn {
    font-size: 18px;
  }
}
</style>
