import { defineStore } from 'pinia'
import { VxeUI, VxeGlobalThemeName, VxeGlobalI18nLocale, VxeComponentSizeType } from 'vxe-pc-ui'
import tinycolor2 from 'tinycolor2'
import { RouteLocationNormalizedLoadedGeneric } from 'vue-router'
import { EnumsVO } from '@/api/user'
import { getEnums } from '@/api/app'

function updatePrimaryColor (color: string) {
  if (color) {
    document.documentElement.style.setProperty('--vxe-ui-font-primary-color', color)
    document.documentElement.style.setProperty('--vxe-ui-font-primary-tinge-color', tinycolor2(color).lighten(28).toString())
    document.documentElement.style.setProperty('--vxe-ui-font-primary-lighten-color', tinycolor2(color).lighten(6).toString())
    document.documentElement.style.setProperty('--vxe-ui-font-primary-darken-color', tinycolor2(color).darken(12).toString())
    document.documentElement.style.setProperty('--vxe-ui-font-primary-disabled-color', tinycolor2(color).lighten(15).toString())
  } else {
    document.documentElement.style.removeProperty('--vxe-ui-font-primary-color')
    document.documentElement.style.removeProperty('--vxe-ui-font-primary-tinge-color')
    document.documentElement.style.removeProperty('--vxe-ui-font-primary-lighten-color')
    document.documentElement.style.removeProperty('--vxe-ui-font-primary-darken-color')
    document.documentElement.style.removeProperty('--vxe-ui-font-primary-disabled-color')
  }
}

const currTheme = (localStorage.getItem('APP_THEME') || 'light') as VxeGlobalThemeName
const currPrimaryColor = localStorage.getItem('VXE_DOCS_PRIMARY_COLOR') || ''
const currComponentsSize = (localStorage.getItem('VXE_DOCS_COMPONENTS_SIZE') || '') as VxeComponentSizeType
const currLanguage = (localStorage.getItem('APP_LANGUAGE') || 'zh-CN') as VxeGlobalI18nLocale

VxeUI.setLanguage(currLanguage)
setTimeout(() => {
  VxeUI.setTheme(currTheme)
})

if (currPrimaryColor) {
  updatePrimaryColor(currPrimaryColor)
}

export const useAppStore = defineStore('app', {
  state: () => {
    return {
      icon: 'dns',
      theme: currTheme,
      enums: {} as EnumsVO,
      webTitle: window.WebTitle,
      primaryColor: currPrimaryColor,
      componentsSize: currComponentsSize,
      language: currLanguage
    }
  },
  getters: {
  },
  actions: {
    /**
     * 设置主题
     * @param theme
     */
    setTheme (theme: VxeGlobalThemeName) {
      this.theme = theme
      VxeUI.setTheme(theme)
      localStorage.setItem('APP_THEME', theme || '')
    },
    setPrimaryColor (color: string) {
      updatePrimaryColor(color)
      this.primaryColor = color
      localStorage.setItem('VXE_DOCS_PRIMARY_COLOR', color)
    },
    setComponentsSize (size: VxeComponentSizeType) {
      this.componentsSize = size
      localStorage.setItem('VXE_DOCS_COMPONENTS_SIZE', size || '')
    },
    async setEnums () {
      const res = await getEnums()
      this.enums = res.data as EnumsVO
    },
    /**
     * 设置语言
     * @param language
     */
    setLanguage (language: VxeGlobalI18nLocale) {
      if (language !== this.language) {
        this.language = language
        VxeUI.setLanguage(language)
        localStorage.setItem('APP_LANGUAGE', language)
      }
    },
    updatePageTitle (route: RouteLocationNormalizedLoadedGeneric) {
      this.icon = (route.meta?.icon as string) || 'dns'
      const favicon = document.querySelector<HTMLLinkElement>('link[rel="icon"]')
      if (favicon) {
        favicon.type = 'image/png'
        favicon.href = `/assets/icon-${this.icon}.png`
      }
      this.webTitle = String(route.meta?.title || '')
      document.title = [this.webTitle, window.WebTitle].join(' - ')
    }
  }
})
