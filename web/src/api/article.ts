import { requestAjax } from '@/api/http'
import { ArticleDataVO } from '@/utils/interface'
import { OptionMethod } from '@/api/user'

export function getReadme () {
  return requestAjax({
    url: '/readme',
    method: 'get'
  })
}

export function getCategory () {
  return requestAjax({
    url: '/category',
    method: 'get'
  })
}

export function deleteCategory (params?: any) {
  return requestAjax({
    url: '/category',
    method: 'delete',
    data: { id: params.id }
  })
}

export function postCategory (isAdd: boolean, data?: any) {
  return requestAjax({
    url: '/category',
    method: isAdd ? 'post' : 'put',
    data
  })
}

export function postUploadImg (form: FormData) {
  return requestAjax({
    url: '/upload/images',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data: form
  })
}

export function optionArticle (method: OptionMethod, data: ArticleDataVO) {
  return requestAjax({
    url: '/article',
    method: method,
    data
  })
}

export function getArticleList (params?: any) {
  return requestAjax({
    url: '/article/list',
    method: 'get',
    params
  })
}

export function getArticle (params?: any) {
  return requestAjax({
    url: '/article',
    method: 'get',
    params
  })
}
