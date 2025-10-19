import { requestAjax } from './http'

export function getDnsListPage (params?: any) {
  return requestAjax({
    url: '/list',
    method: 'get',
    params
  })
}

export function getNetworkInterfaces () {
  return requestAjax({
    url: '/network/interfaces',
    method: 'get'
  })
}

export function getCopyright () {
  return requestAjax({
    url: '/copyright',
    method: 'get'
  })
}

export function postDnsSaveBatch (data?: any) {
  return requestAjax({
    url: '/save',
    method: 'post',
    data
  })
}

export function deleteDnsDelete (data?: any) {
  return requestAjax({
    url: '/delete',
    method: 'delete',
    data
  })
}

export function postDnsService (action: string, iface: string) {
  action = action === 'running' ? 'stop' : 'running'
  return requestAjax({
    url: `/service/${action}/${iface}`,
    method: 'post'
  })
}
