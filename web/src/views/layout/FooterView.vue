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
    <span>运行时间: {{ copyright.runtime.days }}天 {{ copyright.runtime.hours }}小时 {{ copyright.runtime.minutes }}分 {{ copyright.runtime.seconds }}秒</span>
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
  runtime: {
    days: '00',
    hours: '00',
    minutes: '00',
    seconds: '00'
  }
})

getCopyright().then(res => {
  run = new Date(res.data.runTime).getTime()
  now = new Date(res.data.nowTime).getTime()
  Object.assign(copyright.value, {
    ...res.data,
    runtime: copyright.value.runtime // 保留原 runtime
  })
  updateTimer() // 立即执行一次
})

function updateTimer () {
  const diff = now - run // 毫秒差
  copyright.value.runtime = copyright.value.runtime || {}
  Object.assign(copyright.value.runtime, {
    days: String(Math.floor(diff / (1000 * 60 * 60 * 24))).padStart(2, '0'),
    hours: String(Math.floor((diff / (1000 * 60 * 60)) % 24)).padStart(2, '0'),
    minutes: String(Math.floor((diff / (1000 * 60)) % 60)).padStart(2, '0'),
    seconds: String(Math.floor((diff / 1000) % 60)).padStart(2, '0')
  })
  // 秒数自动增加
  now += 1000
}
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
