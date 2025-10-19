import { defineStore } from 'pinia'
import { ArticleListVO, CategoryVO } from '@/api/user'
import { getArticle, getArticleList, getCategory } from '@/api/article'
import { buildCategoryTree } from '@/utils'

export const useArticleStore = defineStore('article', {
  state: () => {
    return {
      treeList: [] as CategoryVO[],
      articleInfo: {} as ArticleListVO,
      categoryList: [] as CategoryVO[],
      articleTags: {
        tags: [] as string[],
        articleList: [] as ArticleListVO[]
      }
    }
  },
  getters: {
  },
  actions: {
    setTreeList () {
      const _this = this
      _this.treeList = [{
        id: '0',
        title: '根节点',
        name: 'root',
        description: '根节点',
        articleCount: 0,
        pid: '0',
        sort: 0,
        status: 1,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        children: []
      }]
      getCategory().then(rest => {
        const categoryList = buildCategoryTree(rest.data)
        categoryList.forEach((item) => {
          if (item.name === 'all' || item.name === 'default') {
            return
          }
          _this.treeList.push(item)
        })
      })
    },
    setCategory (query, router) {
      const _this = this
      getCategory().then(rest => {
        _this.categoryList = buildCategoryTree(rest.data)
        if (!query.cid || !query.type) {
          const q = _this.categoryList[0]
          router.push({ name: 'Article', query: { cid: q.id, type: q.name } })
        }
      })
    },
    setArticleList (query) {
      const _this = this
      getArticleList(query).then((res) => {
        _this.articleTags.tags = []
        _this.articleTags.articleList = res.data
        if (res.data) {
          res.data.forEach(item => {
            for (const tag of item.tags.split('、')) {
              if (!_this.articleTags.tags.includes(tag)) {
                _this.articleTags.tags.push(tag)
              }
            }
          })
        }
      })
    },
    setArticleInfo (query) {
      const _this = this
      getArticle(query).then((res) => {
        document.title = [res.data.title, window.WebTitle].join(' - ')
        _this.articleInfo = res.data
        _this.articleInfo.content = `> **文章摘要**：${res.data.summary}\n\n${res.data.content}`
      })
    }
  }
})
