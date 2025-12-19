import { VxeIconPropTypes } from 'vxe-pc-ui/types/components/icon'

export type OptionMethod = 'post' | 'put' | 'delete'

export interface HomeDataVO {
  id: string,
  icon: VxeIconPropTypes.Name,
  name: string,
  sort: number,
  iconType: 0 | 1,
  position: 0 | 1,
  createdAt: string,
  updatedAt: string
}

/* ==================== Enums ==================== */
export interface EnumItemVO {
  Key: string;
  Name: string;
  Desc: string;
}

export interface EnumsVO {
  code: Record<number, EnumItemVO>;
  icon: Record<number, EnumItemVO>;
  position: Record<number, EnumItemVO>;
  status: Record<number, EnumItemVO>;
}

/* ==================== Dns ==================== */
export interface DnsVO {
  id: string
  protocol: string
  domain: string
  ip: string
  status: string
  port: string
  updatedAt: string
  createdAt: string
}

/* ==================== Category ==================== */
export interface CategoryVO {
  id: string
  name: string
  title: string
  docSort: string
  articleCount: number
  pid: string
  sort: number
  status: number
  isExpand: boolean
  createdAt: string
  updatedAt: string
  children?: CategoryVO[]
}

export interface CategoryAddVO {
  id: string
  title: string
  pid: string
  sort: number
  docSort: string
  isExpand: boolean
}

export interface ArticleListVO {
  id: string
  docSort: string
  categoryID: string
  categoryName: string
  categoryTitle: string
  tags: string
  title: string
  summary: string
  content: string
  createdAt: string
  updatedAt: string
}
