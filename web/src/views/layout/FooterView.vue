<template>
  <div class="page-footer">
    <span>名称: {{ copyright.name }}</span>
    <span class="split">|</span>
    <span>版本: {{ copyright.version }}</span>
    <span class="split">|</span>
    <span>Go版本: {{ copyright.goVersion }}</span>
    <span class="split">|</span>
    <span>平台: {{ copyright.platform }}</span>
    <span class="split">|</span>
    <span>PID: {{ copyright.pid }}</span>
    <span class="split">|</span>
    <span>编译时间: {{ copyright.buildTime }}</span>
    <span class="split">|</span>
    <span>运行时间: {{ copyright.runtime }}</span>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { getCopyright } from '@/api/dns'

let run: number = new Date().getTime()
let now: number = new Date().getTime()

const copyright = ref({
  name: '本地DNS代理',
  version: '1.0.0',
  platform: 'windows/amd64',
  goVersion: '1.0.0',
  pid: '0',
  buildTime: '2024-01-01 00:00:00',
  runtime: '0天0小时0分0秒'
})

getCopyright().then(res => {
  run = new Date(res.data.runTime).getTime()
  now = new Date(res.data.nowTime).getTime()
  copyright.value = res.data
})

function updateTimer () {
  const diff = now - run // 毫秒差
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((diff / (1000 * 60 * 60)) % 24)
  const minutes = Math.floor((diff / (1000 * 60)) % 60)
  const seconds = Math.floor((diff / 1000) % 60)

  copyright.value.runtime = `${days}天 ${hours}小时 ${minutes}分 ${seconds}秒`

  // 秒数自动增加
  now += 1000
}
updateTimer() // 立即执行一次
setInterval(updateTimer, 1000) // 每秒更新
</script>

<style lang="scss" scoped>
.page-footer {
  text-align: center;
  padding: 8px;
  display: flex;
  gap: 10px;
  align-items: center;
  justify-content: center;
  background-color: var(--page-layout-background-color);

  .split {
    user-select: none;
    color: var(--vxe-ui-status-info-color);
  }
}
</style>
