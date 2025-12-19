import { createRouter, createWebHashHistory } from 'vue-router'
import { routeConfigs } from './config'
import NProgress from 'nprogress'
import { useAppStore } from '@/store/app'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: routeConfigs
})

router.beforeEach((to, __, next) => {
  NProgress.start()
  if (!to.name) {
    next({ name: 'Article' }) // 没有路由名，跳到首页
  } else {
    next() // 有路由名，正常放行
  }
})

router.afterEach((to) => {
  const appStore = useAppStore()
  // 更新标题
  appStore.updatePageTitle(to)
  NProgress.done()
})

export default router
