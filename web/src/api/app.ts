import { requestAjax } from '@/api/http'

export function getEnums () {
  return requestAjax({
    url: '/enums',
    method: 'get'
  })
}
