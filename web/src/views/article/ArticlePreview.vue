<script setup lang="ts">
import { optionArticle } from '@/api/article'
import { useRoute, useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { MdEditor, PreviewThemes, Themes } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'
import { ArticleListVO } from '@/api/user'
import { formatTime, timeAgo } from '@/utils'
import 'md-editor-v3/lib/style.css'
import { VxeUI } from 'vxe-pc-ui'
import { useArticleStore } from '@/store/article'

const route = useRoute()
const router = useRouter()
const useArticle = useArticleStore()

const editorRef = ref<InstanceType<typeof MdEditor> | null>(null)
const theme = ref<Themes>((localStorage.getItem('APP_THEME') || 'light') as Themes)
const previewTheme = ref<PreviewThemes>('default')

onMounted(() => {
  // @ts-ignore
  editorRef.value?.togglePreviewOnly(true)
  // @ts-ignore
  editorRef.value?.toggleCatalog(true)
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
const copyEvent = (content: string) => {
  if (VxeUI.clipboard.copy(content)) {
    VxeUI.modal.message({
      status: 'success',
      content: '复制成功'
    })
  }
}
useArticle.fetchArticleInfo(route.query)
</script>

<template>
  <div class="article-preview-box">
    <aside class="article-preview-left-box">
      <div class="info-card">
        <h3>文章信息</h3>
        <p class="meta">分类：<router-link class="title-category" :to="{ name: 'Article', query: { cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName, sort: useArticle.articleInfo.docSort } }">{{ useArticle.articleInfo.categoryTitle }}</router-link></p>
        <p class="meta tags">标签：<router-link :to="{ name: 'Article', query: { cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName, sort: useArticle.articleInfo.docSort, kw: tag } }" v-for="(tag, index) in useArticle.articleInfo.tags?.split('、')" :key="index">{{ tag }}</router-link></p>
        <p class="meta">发布时间：{{ formatTime(useArticle.articleInfo.createdAt) }}({{ timeAgo(useArticle.articleInfo.createdAt) }})</p>
        <p class="meta">更新时间：{{ formatTime(useArticle.articleInfo.updatedAt) }}({{ timeAgo(useArticle.articleInfo.updatedAt) }})</p>
        <p class="meta"></p>
        <div class="meta options">
          <router-link class="options-box-edit" :to="{name: 'EditArticle', query: {id: useArticle.articleInfo.id, cid: useArticle.articleInfo.categoryID, type: useArticle.articleInfo.categoryName, sort: useArticle.articleInfo.docSort}, params: {type: 'edit'}}">编辑</router-link>
          <a @click="onDeleteArticle(useArticle.articleInfo)" class="options-box-delete" href="javascript:void(0)">删除</a>
          <div class="flex1"></div>
          <a @click="copyEvent(useArticle.articleInfo.content)" class="options-box-delete" href="javascript:void(0)">点击复制</a>
        </div>
      </div>
    </aside>

    <main class="article-preview-content">
      <PageView>
        <h1 class="title">{{ useArticle.articleInfo.title }}</h1>
        <div class="article-body">
          <!-- ✅ 加 ref -->
          <MdEditor
              :theme="theme"
              ref="editorRef"
              v-model="useArticle.articleInfo.content"
              codeTheme="github"
              catalogLayout="flat"
              :previewTheme="previewTheme"
              :readOnly="true"
              :footers="[]"
              :toolbars="[]">
          </MdEditor>
        </div>
      </PageView>
    </main>

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

          .flex1 {
            flex: 1;
          }

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
      height: calc(100vh - 185px);
      overflow-y: auto;          // ✅ 让它成为滚动容器
      scroll-behavior: smooth;   // ✅ 点击目录平滑滚动
      overflow-x: hidden;

      .md-editor {
        height: 100%;
        border: none;

        :deep(.md-editor-catalog-editor) {
          width: 300px;
        }
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
