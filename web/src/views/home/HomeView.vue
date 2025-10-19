<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import router from '@/router'
import { getHomeDataList } from '@/api/home'
import { HomeDataVO } from '@/api/user'
import { unDraggable } from '@/utils'

const sidebar = ref<HomeDataVO[]>([])
const body = ref<HomeDataVO[]>([])

getHomeDataList().then(res => {
  res.data.forEach((item: HomeDataVO) => {
    if (item.position === 1) {
      sidebar.value.push(item)
    } else {
      body.value.push(item)
    }
  })
})

const routers = computed(() => {
  return router.getRoutes().filter(route => route.meta.isHome)
})

onMounted(() => {
  unDraggable('img', 'a')
})
</script>

<template>
  <PageView>
    <div class="home-box">
      <div class="home-sidebar">
        <div class="sidebar-body">
          <div class="sidebar-avatar">
            <vxe-avatar circle src="/assets/icon-home.png"></vxe-avatar>
          </div>
          <ul class="sidebar-group">
            <li class="sidebar-group-item" v-for="(item, i) in sidebar" :key="i" :class="{'active': i === 0}">
              <vxe-icon class="sidebar-icon f22" :name="item.icon"/>
              <vxe-text class="group-item-title">{{ item.name }}</vxe-text>
            </li>
            <li class="sidebar-group-item">
              <vxe-icon class="sidebar-icon sidebar-group-add f20" name="add"/>
            </li>
          </ul>
          <div class="sidebar-setting">
            <vxe-icon class="sidebar-icon f20" name="setting"/>
          </div>
        </div>
      </div>
      <div class="main-body-background"></div>
      <div class="home-main" style="background-image:url(https://files.codelife.cc/wallpaper/wallspic/20250310akbkk0.jpeg?x-oss-process=image/resize,limit_0,m_fill,w_1920,h_1080/quality,Q_93/format,webp);">
        <div class="time-wrap" style="display: none">
          <vxe-text class="time">{{ new Date().toLocaleString() }}</vxe-text>
          <vxe-text class="date">
            <vxe-text class="time-month">10月2号</vxe-text>
            <vxe-text class="time-week">星期四</vxe-text>
            <vxe-text class="time-lunar">八月十一</vxe-text>
          </vxe-text>
        </div>
        <div class="icon-body-wrap">
          <div class="icon-body-wrap-grid">
            <ul>
              <li class="icon-size-2x2" v-for="(route, index) in routers" :key="index">
                <vxe-link :underline="false" :router-link="{path: route.path, query: {id: 1}}">
                  <vxe-avatar style="--icon-fit: cover; --icon-bg-color: #ffffff;" :src="`/assets/icon-${route.meta.icon}.png`"></vxe-avatar>
                </vxe-link>
                <p>{{ route.meta.homeTitle }}</p>
              </li>
            </ul>
          </div>
        </div>
        <div class="footer-warp" style="display: none">
          博观而约取，厚积而薄发。 ---苏轼-
        </div>
      </div>
    </div>
  </PageView>
</template>

<style scoped lang="scss">
.home-box {
  --icon-opacity: 1;
  --icon-radius: 16px;
  --icon-gap-x: 30px;
  --icon-gap-y: 30px;
  --alpha-bg: 255, 255, 255;
  --spacing: .25rem;
  --img-bg: 199, 216, 229;
  --img-text: 34, 34, 34;
  --sidebar-width: 50px;
  --icon-size: 60px;
  --color-white: #ffffff;
  --avatar-size: 30px;
  --text-color-disabled: #c0c4cc;
  --icon-max-width: 1350px;
  --icon-transition-duration: .3s;
  --sidebar-opacity: 0.4;
  --time-size: 70px;
  --time-font: HarmonyOS_Sans;
  --time-color: #fff;
  --time-fontWeight: 400;
  --time-month: inline;
  --time-week: inline;
  --time-lunar: inline;
  --time-sec: inline;

  height: 100vh;

  .home-sidebar {
    position: fixed;
    left: 0;
    z-index: 2;
    animation: fadeIn .2s;
    width: var(--sidebar-width);
    bottom: calc(var(--spacing) * 0);
    top: calc(var(--spacing) * 0);
    box-shadow: 0 3px 0 rgba(12, 12, 12, 0.03);
    transition: transform .2s;
    color: rgba(var(--img-text), .6);
    backdrop-filter: blur(6px);
    background-color: rgba(var(--img-bg), var(--sidebar-opacity, .2));

    .sidebar-body {
      text-align: center;
      justify-content: center;
      flex-direction: column;
      width: 100%;
      height: 100%;
      font-size: 12px;
      display: flex;

      .sidebar-avatar {
        padding-top: 40px;
        padding-bottom: 40px;
        cursor: pointer;

        .vxe-avatar {
          --avatar-bg-color: var(--text-color-disabled);
          --avatar-size: 40px;
          align-items: center;
          background: var(--avatar-bg-color);
          box-sizing: border-box;
          color: var(--color-white);
          display: inline-flex;
          font-size: 14px;
          height: var(--avatar-size);
          justify-content: center;
          outline: none;
          overflow: hidden;
          text-align: center;
          width: var(--avatar-size);
        }
      }

      .f22 {
        font-size: 22px;
      }

      .f20 {
        font-size: 20px;
      }

      .sidebar-icon {
        height: 1em;
        width: 1em;
        line-height: 1em;
        display: inline-flex;
        justify-content: center;
        align-items: center;
        fill: currentColor;
        color: inherit;
        cursor: pointer;
        transition: transform .2s;
      }

      .sidebar-group {
        flex: 1;
        display: flex;
        flex-direction: column;
        overflow-y: auto;
        overflow-x: hidden;
        scrollbar-color: rgba(var(--alpha-bg), .4) transparent;
        scrollbar-width: none;

        .sidebar-group-item {
          display: flex;
          justify-content: center;
          align-items: center;
          height: 50px;
          min-height: 50px;
          transition-duration: .3s;
          user-select: none;
          flex-direction: column;

          &.active {
            background-color: #ffffff4d
          }

          &:hover {
            .vxe-icon {
              transform: scale(1.2);
            }
          }

          .vxe-icon {
            font-weight: 500;
          }

          .group-item-title {
            user-select: none;
            width: var(--sidebar-width);
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            font-weight: bold;
          }
        }
      }

      .sidebar-setting {
        height: var(--sidebar-width);
        display: flex;
        flex-flow: wrap;
        align-items: center;
        justify-content: center;

        .vxe-icon:hover {
          animation: rotate ease-in-out .4s;
        }
      }
    }
  }

  .main-body-background {
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;

    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    opacity: 0.5;
    background-color: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(0px);
  }

  .home-main {
    display: flex;
    height: 100%;
    flex-direction: column;

    .time-wrap {
      display: flex;
      justify-content: center;
      flex-direction: column;
      user-select: none;
      text-align: center;

      :deep(.vxe-text--content) {
        color: var(--time-color, "#fff") !important;
      }

      .time {
        font-size: var(--time-size);
        font-family: var(--time-font);
        user-select: none;
        font-weight: var(--time-fontWeight);
        text-shadow: 0 2px 6px rgb(0 0 0 / 16%);
        display: inline-block;
        line-height: var(--time-size);
        transition: font .2s;
      }

      .date {
        font-size: 14px;
        line-height: 26px;
        opacity: 0.88;
        margin-top: -3px;
        text-shadow: 0 2px 4px rgb(0 0 0 / 16%);

        .time-week {
          margin: 0 6px;
        }

        .time-week,
        .time-lunar,
        .time-month {
          display: var(--time-month);
        }
      }
    }

    .icon-body-wrap {
      transition: .1s ease;

      .icon-body-wrap-grid {
        max-width: var(--icon-max-width, 1350px);
        margin: 0 auto;
        padding: 0 var(--sidebar-width, 45px);

        > ul {
          position: relative;
          display: grid;
          padding-top: 2vh;
          user-select: none;
          grid-template-columns: repeat(auto-fill, var(--icon-size));
          grid-template-rows: repeat(auto-fill, var(--icon-size));
          grid-auto-flow: dense;
          grid-gap: var(--icon-gap-x) var(--icon-gap-y);
          box-sizing: border-box;
          justify-content: center;
          padding-bottom: 120px;

          .icon-size-1x1 {
            width: var(--icon-size);
            height: var(--icon-size)
          }

          .icon-size-1x2 {
            grid-column: span 2;
            width: calc(var(--icon-size) * 2 + var(--icon-gap-y) * 1);
            height: var(--icon-size)
          }

          .icon-size-2x1 {
            grid-row: span 2;
            width: var(--icon-size);
            height: calc(var(--icon-size) * 2 + var(--icon-gap-x) * 1)
          }

          .icon-size-2x2 {
            grid-column: span 2;
            grid-row: span 2;
            width: calc(var(--icon-size) * 2 + var(--icon-gap-y));
            height: calc(var(--icon-size) * 2 + var(--icon-gap-x))
          }

          .icon-size-2x4 {
            grid-column: span 4;
            grid-row: span 2;
            width: calc(var(--icon-size) * 4 + var(--icon-gap-y) * 3);
            height: calc(var(--icon-size) * 2 + var(--icon-gap-x))
          }

          > li {
            list-style-type: none;
            position: relative;
            grid-column: span 1;
            grid-row: span 1;
            user-select: none;
            -webkit-user-select: none;
            box-sizing: border-box;
            opacity: var(--icon-opacity);
            width: calc(var(--icon-size) + var(--icon-gap-y));
            height: calc(var(--icon-size) + var(--icon-gap-x));
            transition-duration: var(--icon-transition-duration);
            transition: transform .2s;

            :deep(.vxe-link--content) {
              padding: 0 !important;

              .vxe-avatar {
                width: 100% !important;
                height: 100% !important;
                overflow: hidden;
                border-radius: var(--icon-radius);
                box-shadow: 0 0 5px #0000001a;

                &:hover {
                  box-shadow: 0 0 10px #0000004d;
                }

                img {
                  display: block;
                  object-fit: var(--icon-fit, cover);
                  width: 100%;
                  height: 100%;
                  background-repeat: no-repeat;
                  background-size: cover;
                  background-color: var(--icon-bg-color);
                  user-select: none;
                }
              }
            }

            p {
              width: calc(100% + var(--icon-gap-y));
              margin-left: calc(var(--icon-gap-y) / 2 * -1);
              display: block;
              margin-top: 3px;
              text-align: center;
              color: #fff;
              font-size: 12px;
              line-height: 1.1;
              filter: drop-shadow(0px 2px 7px rgba(0, 0, 0, .8));
              text-overflow: ellipsis;
              overflow: hidden;
              white-space: nowrap;
            }
          }
        }
      }
    }
  }
}

@keyframes rotate {
  0% {
    transform: rotate(120deg)
  }
}

@keyframes fadeIn {
  0% {
    opacity: 0
  }
  to {
    opacity: 1
  }
}
</style>
