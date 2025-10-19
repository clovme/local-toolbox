<script setup lang="ts">
import { optionArticle } from '@/api/article'
import { useRoute, useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { MdPreview, MdCatalog } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'
import { ArticleListVO } from '@/api/user'
import { formatTime, timeAgo } from '@/utils'
import 'md-editor-v3/lib/style.css'
import { VxeUI } from 'vxe-pc-ui'
import { useArticleStore } from '@/store/article'

const route = useRoute()
const router = useRouter()
const useArticle = useArticleStore()
const id = 'preview-only'

const scrollElement = ref<HTMLElement | null>(null)

onMounted(() => {
  scrollElement.value = document.documentElement
})

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
      router.push({ name: 'Article', query: { cid: item.categoryID, type: item.categoryName } })
    })
  })
}
useArticle.setArticleInfo(route.query)
</script>

<template>
  <div class="article-preview-box">
    <aside class="article-preview-left-box">
      <div class="info-card">
        <h3>文章信息</h3>
        <p class="meta">分类：<router-link class="title-category" :to="{ name: 'Article', query: { cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName } }">{{ useArticle.articleInfo.categoryTitle }}</router-link></p>
        <p class="meta tags">标签：<router-link :to="{ name: 'Article', query: { cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName, kw: tag } }" v-for="(tag, index) in useArticle.articleInfo.tags?.split('、')" :key="index">{{ tag }}</router-link></p>
        <p class="meta">发布时间：{{ formatTime(useArticle.articleInfo.createdAt) }}({{ timeAgo(useArticle.articleInfo.createdAt) }}前)</p>
        <p class="meta">更新时间：{{ formatTime(useArticle.articleInfo.updatedAt) }}({{ timeAgo(useArticle.articleInfo.updatedAt) }}前)</p>
        <p class="meta"></p>
        <p class="meta options">
          <router-link class="options-box-edit" :to="{name: 'EditArticle', query: {id: useArticle.articleInfo.id, cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName}, params: {type: 'edit'}}">编辑</router-link>
          <a @click="onDeleteArticle(useArticle.articleInfo)" class="options-box-delete" href="javascript:void(0)">删除</a>
        </p>
      </div>
    </aside>

    <main class="article-preview-content">
      <PageView>
        <h1 class="title">{{ useArticle.articleInfo.title }}</h1>
        <div class="article-body">
          <!-- ✅ 加 ref -->
          <MdPreview codeTheme="github" :codeFoldable="false" :noPrettier="true" :id="id" :modelValue="useArticle.articleInfo.content || ''" />
        </div>
      </PageView>
    </main>

    <aside class="article-preview-catalog-box">
      <div class="article-preview-catalog">
        <!-- ✅ scrollElement 绑定预览容器 -->
        <MdCatalog :editorId="id" :scrollElement="scrollElement as HTMLElement" />
      </div>
    </aside>
  </div>
</template>

<style scoped lang="scss">
.article-preview-box {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  gap: 20px;
  width: 100%;
  max-width: 1600px;
  margin: 0 auto;
  box-sizing: border-box;

  .article-preview-left-box {
    width: 300px;
    position: sticky;
    top: 0; // 距离顶部固定位置
    align-self: flex-start;
    display: flex;
    gap: 20px;
    flex-direction: column;

    .info-card {
      background: #fff;
      padding: 16px;
      border-radius: 8px;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);

      h3 {
        font-size: 16px;
        margin-bottom: 10px;
      }

      .meta {
        font-size: 13px;
        color: #666;
        margin-top: 4px;

        &.options {
          display: flex;
          gap: 10px;
          justify-content: flex-end;

          a {
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

        &.tags {
          display: flex;
          gap: 5px;
          flex-wrap: wrap;

          a {
            text-decoration: none;

            &:nth-child(1n) { color: #e74c3c; } /* 红 */
            &:nth-child(2n) { color: #3498db; } /* 蓝 */
            &:nth-child(3n) { color: #2ecc71; } /* 绿 */
            &:nth-child(4n) { color: #9b59b6; } /* 紫 */
            &:nth-child(5n) { color: #f1c40f; } /* 黄 */
            &:nth-child(6n) { color: #909399; } /* 黄 */
            &:nth-child(7n) { color: #519f2b; } /* 黄 */
            &:nth-child(8n) { color: #cc8218; } /* 黄 */
            &:nth-child(9n) { color: #f56c6c; } /* 黄 */
            &:nth-child(10n) { color: #f56c6c; } /* 黄 */

            &:hover {
              &:nth-child(10n) { color: #e74c3c; } /* 红 */
              &:nth-child(9n) { color: #3498db; } /* 蓝 */
              &:nth-child(8n) { color: #2ecc71; } /* 绿 */
              &:nth-child(7n) { color: #9b59b6; } /* 紫 */
              &:nth-child(6n) { color: #f1c40f; } /* 黄 */
              &:nth-child(5n) { color: #909399; } /* 黄 */
              &:nth-child(4n) { color: #519f2b; } /* 黄 */
              &:nth-child(3n) { color: #cc8218; } /* 黄 */
              &:nth-child(2n) { color: #f56c6c; } /* 黄 */
              &:nth-child(1n) { color: #f56c6c; } /* 黄 */
            }
          }
        }
      }
    }
  }

  .article-preview-content {
    flex: 1;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
    line-height: 1.8;

    .title {
      font-size: 28px;
      font-weight: bold;
    }

    .article-body {
      max-width: 1000px;
      height: calc(100vh - 185px);
      overflow-y: auto;          // ✅ 让它成为滚动容器
      scroll-behavior: smooth;   // ✅ 点击目录平滑滚动
      overflow-x: hidden;
    }
  }

  .article-preview-catalog-box {
    width: 300px;
    position: sticky;
    top: 0;
    height: calc(100vh - 134px);

    .article-preview-catalog {
      position: sticky;
      top: 0;
      background: #fff;
      border-radius: 8px;
      padding: 16px;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);

      .md-editor-catalog {
        max-height: calc(100vh - 134px);
        overflow-y: auto;
      }
    }
  }
}

/* 暗色模式支持 */
@media (prefers-color-scheme: dark) {
  .article-preview-box {
    .info-card,
    .article-preview-content,
    .article-preview-catalog {
      background: #1e1e1e;
      color: #ddd;
      box-shadow: 0 2px 8px rgba(255, 255, 255, 0.1);
    }
  }
}
</style>
