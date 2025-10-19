import { UserConfig, ConfigEnv, loadEnv } from 'vite'
import path from 'path'
import XEUtils from 'xe-utils'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import { createHtmlPlugin } from 'vite-plugin-html'
import autoprefixer from 'autoprefixer'

// https://vitejs.dev/config/
export default ({ mode }: ConfigEnv): UserConfig => {
  const env = loadEnv(mode, process.cwd(), 'VITE_')
  return {
    base: '/',
    plugins: [
      vue(),
      vueJsx(),
      createHtmlPlugin({
        inject: {
          data: {
            VITE_APP_BUILD_TIME: XEUtils.toDateString(new Date(), 'yyyy-MM-dd HH:mm:ss'),
            ...env
          }
        }
      })
    ],
    resolve: {
      alias: {
        '@': path.join(__dirname, './src')
      },
      extensions: ['.js', '.vue', '.json', '.ts', '.tsx']
    },
    server: {
      port: 8084
    },
    css: {
      postcss: {
        plugins: [autoprefixer()]
      }
    },
    build: {
      assetsDir: 'assets',
      emptyOutDir: true,
      outDir: env.VITE_OUT_DIR || 'dist'
    }
  }
}
