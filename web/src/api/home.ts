import { requestAjax } from '@/api/http'

export function getHomeDataList () {
  return requestAjax({
    url: '/home/data',
    method: 'get'
  })
}
