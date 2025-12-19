import { CategoryVO } from '@/api/user'
import XEUtils from 'xe-utils'

export const unDraggable = (...tagNames: string[]) => {
  tagNames.forEach(tag => {
    document.querySelectorAll(tag).forEach(el => {
      el.setAttribute('draggable', 'false')
    })
  })
}

export const buildCategoryTree = (flat: CategoryVO[]): CategoryVO[] => {
  // 按 id 建索引，方便父子查找
  const map = new Map<string, CategoryVO>()
  flat.forEach(c => map.set(c.id, { ...c, children: [] }))

  const roots: CategoryVO[] = []

  // 建立父子关系
  flat.forEach(c => {
    const node = map.get(c.id)!
    if (c.name === 'all') {
      roots.splice(0, 0, node)
    } else if (c.name === 'default') {
      roots.splice(1, 0, node)
    } else if (c.pid === '0') {
      roots.push(node)
    } else {
      const parent = map.get(c.pid)
      if (parent) parent.children!.push(node)
    }
  })

  // 自底向上累加 articleCount
  const accumulate = (node: CategoryVO): number => {
    let total = node.articleCount
    if (node.children && node.children.length > 0) {
      for (const child of node.children) {
        total += accumulate(child)
      }
    }
    node.articleCount = total
    return total
  }

  roots.forEach(accumulate)
  return roots
}

export function formatTime (value: string) {
  return value ? XEUtils.toDateString(value, 'yyyy-MM-dd HH:mm:ss') : '-'
}

export function newTime () {
  return XEUtils.toDateString(XEUtils.now(), 'HH:mm:ss')
}

export function timeAgo (isoTime: string): string {
  const date = new Date(isoTime)
  const now = new Date()
  const diff = now.getTime() - date.getTime() // 毫秒差

  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour
  const month = 30 * day
  const year = 365 * day

  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return Math.floor(diff / minute) + '分钟前'
  } else if (diff < day) {
    return Math.floor(diff / hour) + '小时前'
  } else if (diff < month) {
    return Math.floor(diff / day) + '天前'
  } else if (diff < year) {
    return Math.floor(diff / month) + '月前'
  } else {
    return Math.floor(diff / year) + '年前'
  }
}
