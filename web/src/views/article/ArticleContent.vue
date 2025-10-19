<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArticleListVO, CategoryVO } from '@/api/user'
import { buildCategoryTree, timeAgo } from '@/utils'
import { getCategory, optionArticle } from '@/api/article'
import PageView from '@/views/layout/PageView.vue'
import { VxeUI } from 'vxe-pc-ui'
import { useArticleStore } from '@/store/article'

const route = useRoute()
const router = useRouter()
const treeList = ref<CategoryVO[]>([])
const useArticle = useArticleStore()
const search = ref(route.query.kw as string)

const onDeleteArticle = (item: ArticleListVO) => {
  VxeUI.modal.confirm({
    title: '删除提示',
    content: `删除《${item.title}》？`,
    escClosable: true,
    status: 'warning',
    maskClosable: true,
    draggable: false,
    destroyOnClose: true
  }).then(rest => {
    if (rest !== 'confirm') return
    optionArticle('delete', item).then((res) => {
      VxeUI.modal.message({
        content: res.message,
        status: 'success'
      })
      search.value = route.query.kw as string
      useArticle.setTreeList()
      useArticle.setArticleList(route.query)
      useArticle.setCategory(route.query, router)
    })
  })
}

watch(() => route.query, (query) => {
  search.value = query.kw as string
  useArticle.setArticleList(query)
})

getCategory().then(rest => {
  treeList.value = []
  buildCategoryTree(rest.data).forEach((item) => {
    if (item.name === 'all') {
      return
    }
    treeList.value.push(item)
  })
})

const onCreateArticle = () => {
  let id = route.query.cid
  let type = route.query.type
  if (route.query.type === 'all') {
    for (const tree of treeList.value) {
      if (tree.name === 'default') {
        id = tree.id
        type = tree.name
        break
      }
    }
  }
  router.push({ name: 'EditArticle', params: { type: 'add' }, query: { cid: id, type: type } })
}

const onSearch = () => {
  router.push({ name: 'Article', query: { ...route.query, kw: search.value } })
}
useArticle.setArticleList(route.query)
</script>

<template>
  <div class="article-content-box">
    <PageView class="article-content-tools">
      <div class="article-content-search">
        <vxe-input @change="onSearch" max-length="50" show-word-count v-model="search" placeholder="请输入搜索内容" type="search"></vxe-input>
        <vxe-button @click="onSearch" content="搜索"></vxe-button>
      </div>
      <div class="article-content-tool">
        <vxe-button @click="onCreateArticle" status="danger" content="新的创作" prefix-icon="vxe-icon-add"></vxe-button>
      </div>
    </PageView>
    <div class="article-content-list-box">
      <PageView class="article-content-list-content">
        <div class="article-content-list-item" v-for="item in useArticle.articleTags.articleList" :key="item.id">
          <div class="title-box">
            <router-link v-if="route.query.type==='all'" class="title-category" :to="{name: 'Article', query: {cid: item.categoryID, type: item.categoryName}}">{{ item.categoryTitle }}</router-link>
            <router-link class="title-content" :to="{name: 'Preview', query: {id: item.id, cid: item.categoryID, type: item.categoryName}}">{{ item.title }}</router-link>
          </div>
          <vxe-text-ellipsis class="summary" line-clamp="5" :content="item.summary"></vxe-text-ellipsis>
          <div class="options-time">
            <div class="options-box">
              <router-link class="options-box-edit" :to="{name: 'EditArticle', query: {id: item.id, cid: item.categoryID, type: item.categoryName}, params: {type: 'edit'}}">编辑</router-link>
              <a @click="onDeleteArticle(item)" class="options-box-delete" href="javascript:void(0)">删除</a>
            </div>
            <div class="time">
              <span>{{ timeAgo(item.createdAt) }}前发表</span>
              <span>{{ timeAgo(item.updatedAt) }}前更新</span>
            </div>
          </div>
        </div>
      </PageView>
      <div class="article-content-list-right">
        <PageView class="article-content-list-right-card">
          <router-link :to="{name: 'Article', query: {...route.query, kw: tag}}" v-for="(tag, index) in useArticle.articleTags.tags" :key="index">{{ tag }}</router-link>
        </PageView>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.article-content-box {
  height: 100%;
  width: 100%;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  gap: 10px;

  .article-content-list-box {
    display: flex;
    gap: 20px;
    overflow-y: auto;
    height: 100%;

    .article-content-list-content {
      display: flex;
      gap: 20px;
      height: 100%;
      min-width: unset;
      flex: 1;
      flex-direction: column;

      .article-content-list-item {
        display: flex;
        gap: 5px;
        flex-direction: column;

        &:hover {
          .title-box {
            a.title-category {
              background-color: var(--vxe-ui-status-error-color);

              &:hover {
                background-color: var(--vxe-ui-status-success-color);
              }
            }

            a.title-content {
              color: #409eff;

              &:hover {
                text-decoration: underline !important;
              }
            }
          }
          .options-time {
            .options-box {
              a {
                display: block;
              }
            }
          }
        }

        .title-box {
          display: flex;
          gap: 5px;
          align-items: center;

          a {
            color: #606266;
            transition: all .2s ease-in-out;
            text-decoration: none !important;
          }

          a.title-category {
            background-color: var(--vxe-ui-status-info-color);
            border-radius: 5px;
            padding: 2px 5px;
            color: #fff;
          }

          a.title-content {
            font-size: 18px;
            font-weight: bold;
          }
        }

        .options-time {
          display: flex;

          .options-box {
            flex: 1;
            display: flex;
            gap: 5px;

            a {
              display: none;
              transition: all .1s ease-in-out;
              text-decoration: none;

              &:nth-child(1n) { color: #e74c3c; } /* 红 */
              &:nth-child(2n) { color: #3498db; } /* 蓝 */
              &:nth-child(3n) { color: #2ecc71; } /* 绿 */
              &:nth-child(4n) { color: #9b59b6; } /* 紫 */
              &:nth-child(5n) { color: #f1c40f; } /* 黄 */

              &:hover {
                &:nth-child(1n) { color: #519f2b; } /* 黄 */
                &:nth-child(2n) { color: #cc8218; } /* 黄 */
                &:nth-child(3n) { color: #f56c6c; } /* 黄 */
                &:nth-child(4n) { color: #f56c6c; } /* 黄 */
                &:nth-child(5n) { color: #909399; } /* 黄 */
              }
            }
          }

          .time {
            display: flex;
            gap: 10px;
          }
        }
      }
      }

    .article-content-list-right {
      display: flex;
      flex-direction: column;
      gap: 10px;
      width: 264px;

      .article-content-list-right-card {
        gap: 5px;
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 10px 5px;
        height: unset;
        min-width: unset;

        a {
          background-color: var(--vxe-ui-status-info-color);
          border-radius: 5px;
          padding: 2px 5px;
          color: #fff;
          text-decoration: none;
          transition: all .2s ease-in-out;

          &:nth-child(1n) { background-color: #e74c3c; } /* 红 */
          &:nth-child(2n) { background-color: #3498db; } /* 蓝 */
          &:nth-child(3n) { background-color: #2ecc71; } /* 绿 */
          &:nth-child(4n) { background-color: #9b59b6; } /* 紫 */
          &:nth-child(5n) { background-color: #f1c40f; } /* 黄 */
          &:nth-child(6n) { background-color: #909399; } /* 黄 */
          &:nth-child(7n) { background-color: #519f2b; } /* 黄 */
          &:nth-child(8n) { background-color: #cc8218; } /* 黄 */
          &:nth-child(9n) { background-color: #f56c6c; } /* 黄 */

          &:hover {
            &:nth-child(9n) { background-color: #e74c3c; } /* 红 */
            &:nth-child(8n) { background-color: #3498db; } /* 蓝 */
            &:nth-child(7n) { background-color: #2ecc71; } /* 绿 */
            &:nth-child(6n) { background-color: #9b59b6; } /* 紫 */
            &:nth-child(5n) { background-color: #f1c40f; } /* 黄 */
            &:nth-child(4n) { background-color: #909399; } /* 黄 */
            &:nth-child(3n) { background-color: #519f2b; } /* 黄 */
            &:nth-child(2n) { background-color: #cc8218; } /* 黄 */
            &:nth-child(1n) { background-color: #f56c6c; } /* 黄 */
            transform: scale(1.2);
          }
        }
      }
    }
  }

  .article-content-tools {
    position: relative;
    display: flex;
    gap: 10px;
    height: unset;
    padding: 10px;

    .article-content-search {
      width: 100%;
      flex: 1;
      display: flex;
      gap: 5px;
      align-items: center;
      justify-content: center;

      .vxe-input {
        min-width: 500px;
      }
    }
    .article-content-tool {
      width: 276px;
      display: flex;
      align-items: center;
      justify-content: flex-end;
    }
  }
}
</style>
