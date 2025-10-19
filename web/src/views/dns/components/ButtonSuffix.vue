<script setup lang="ts">
import { ref } from 'vue'
import { getNetworkInterfaces, postDnsService } from '@/api/dns'
import { VxePulldownEvents } from 'vxe-pc-ui'
import { VxeComponentStatusType } from 'vxe-table'

const iface = ref('')
const nif = ref<{ label: string, value: string }[]>([])
const serviceStatus = ref<{ status: VxeComponentStatusType, content: string, icon: string, running: string }>({
  status: 'info',
  content: '未启动',
  icon: 'vxe-icon-caret-right',
  running: 'stop'
})

function setServiceStatus (status: string) {
  switch (status) {
    case 'running':
      serviceStatus.value = { status: 'danger', content: '运行中...', icon: 'vxe-icon-checkbox-unchecked-fill', running: status }
      break
    default:
      serviceStatus.value = { status: 'info', content: '未启动', icon: 'vxe-icon-caret-right', running: status }
  }
}

getNetworkInterfaces().then(res => {
  const items: { label: string, value: string }[] = []
  res.data.ifaces.forEach(item => {
    items.push({ label: item.name, value: item.name })
  })
  nif.value = items
  iface.value = res.data?.iface
  setServiceStatus(res.data.running)
})

const networkInterfacesClickEvent: VxePulldownEvents.OptionClick = ({ option }) => {
  iface.value = option.value as string
}

const onSwitchServiceStatus = () => {
  postDnsService(serviceStatus.value.running, iface.value).then(res => {
    setServiceStatus(res.data.running)
  })
}
</script>

<template>
  <div class="services-group">
    <span class="right-item">
        <span class="right-item-title">代理网卡：</span>
        <vxe-pulldown :options="nif" trigger="click" class="right-item-comp" show-popup-shadow transfer  @option-click="networkInterfacesClickEvent">
          <template #default>
            <vxe-text>{{ iface }}</vxe-text>
            <vxe-icon name="caret-down"></vxe-icon>
          </template>
        </vxe-pulldown>
      </span>

    <span class="right-item">
      <span class="right-item-title">服务状态：</span>
      <vxe-button @click="onSwitchServiceStatus" class-name="service-status-btn" mode="text" :status="serviceStatus.status" :prefix-icon="serviceStatus.icon" :content="serviceStatus.content"></vxe-button>
    </span>
  </div>
</template>

<style scoped lang="scss">
.services-group {
  display: flex;
  flex-direction: row;
  flex-shrink: 0;
  align-items: center;
  gap: 12px;
  margin-left: 48px;

  .right-item {
    cursor: pointer;
    display: flex;
    align-items: center;
  }

  .right-item-title {
    vertical-align: middle;
  }

  .right-item-comp {
    vertical-align: middle;
  }
  .service-status-btn {
    padding: .1em 0;
    :deep(i) {
      font-size: 20px;
      line-height: unset;
    }
  }
  .switch-service.vxe-switch.is--on {
    :deep(.vxe-switch--button) {
      background-color: var(--vxe-ui-status-error-color);
    }
  }
}
</style>
