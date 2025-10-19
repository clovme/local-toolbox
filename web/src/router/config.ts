import { RouteRecordRaw } from 'vue-router'

import UserLayout from '@/views/layout/UserLayout.vue'

export const routeConfigs: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: UserLayout,
    children: [
      {
        path: 'dns',
        name: 'DnsList',
        component: () => import('../views/dns/DnsList.vue'),
        meta: {
          icon: 'dns',
          title: '本地DNS管理',
          isHome: true,
          homeTitle: 'DNS管理'
        }
      },
      {
        path: 'article',
        name: 'Article',
        component: () => import('../views/article/ArticleView.vue'),
        meta: {
          icon: 'article',
          title: '文章管理',
          isHome: true,
          homeTitle: '文章管理'
        }
      },
      {
        path: 'preview',
        name: 'Preview',
        component: () => import('../views/article/ArticlePreview.vue'),
        meta: {
          icon: 'article',
          title: '文章浏览',
          isHome: false,
          homeTitle: '文章浏览'
        }
      }
    ]
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('../views/home/HomeView.vue'),
    meta: {
      icon: 'home',
      title: '首页导航',
      isHome: false,
      homeTitle: '首页导航'
    }
  },
  {
    path: '/article/:type',
    name: 'EditArticle',
    component: () => import('../views/article/EditArticle.vue'),
    meta: {
      icon: 'article',
      title: '文章编辑',
      isHome: false,
      homeTitle: '文章编辑'
    }
  },
  {
    path: '/404',
    name: 'PageError404',
    component: () => import('../views/error/PageError404.vue'),
    meta: {
      title: '404 找不到页面',
      isHome: false,
      homeTitle: '404 找不到页面'
    }
  },
  {
    path: '/403',
    name: 'PageError403',
    component: () => import('../views/error/PageError403.vue'),
    meta: {
      title: '403 无权限访问',
      isHome: false,
      homeTitle: '403 无权限访问'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: {
      name: 'PageError404'
    }
  }
]
