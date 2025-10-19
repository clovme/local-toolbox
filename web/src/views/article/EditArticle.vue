<script setup lang="ts">
import CryptoJS from 'crypto-js'
import { ref, onMounted, reactive, watch } from 'vue'
import { config, PreviewThemes, Themes, ToolbarNames, MdEditor } from 'md-editor-v3'

import type { CompletionSource } from '@codemirror/autocomplete'
import Emoji from './components/Emoji'

import 'md-editor-v3/lib/style.css'
import '@/style/Emoji.scss'
import { useRoute, useRouter } from 'vue-router'
import { VxeFormEvents, VxeFormInstance, VxeFormPropTypes, VxeTreeSelectPropTypes, VxeUI } from 'vxe-pc-ui'
import { ArticleDataVO } from '@/utils/interface'
import { getArticle, getCategory, getReadme, optionArticle, postUploadImg } from '@/api/article'
import { CategoryVO } from '@/api/user'
import { buildCategoryTree, newTime } from '@/utils'

const route = useRoute()
const router = useRouter()

const showPopup = ref(false)
const formDataHash = ref('')
const modelTitle = ref<string>(route.params.type === 'add' ? '添加 - 保存文档' : '编辑 - 保存文档')
const treeList = ref<CategoryVO[]>([])
const formRef = ref<VxeFormInstance<ArticleDataVO>>()
const formData = ref<ArticleDataVO>({
  id: '0',
  title: '',
  tags: '',
  categoryID: route.query.cid as string,
  categoryName: route.query.type as string,
  summary: '',
  content: ''
})
const editorRef = ref<InstanceType<typeof MdEditor> | null>(null)
const theme = ref<Themes>((localStorage.getItem('APP_THEME') || 'light') as Themes)

async function initArticleInfo (query) {
  getCategory().then(rest => {
    treeList.value = []
    buildCategoryTree(rest.data).forEach((item) => {
      if (item.name === 'all') {
        return
      }
      treeList.value.push(item)
    })
  })

  if (route.params.type === 'add') {
    const res = await getReadme()
    formData.value.content = res.data
  } else if (route.params.type === 'edit') {
    const res = await getArticle(query)
    if (res.data) {
      formData.value = res.data
    } else {
      formData.value = {
        id: '0',
        title: '',
        tags: '',
        categoryID: route.query.cid as string,
        categoryName: route.query.type as string,
        summary: '',
        content: ''
      }
    }

    formDataHash.value = CryptoJS.SHA256(formData.value.content.trim()).toString()

    let mdEditorFooterLabel: HTMLLabelElement|null = null
    let mdEditorFooterItem: HTMLDivElement|null = null
    let mdEditorFooterRight: HTMLDivElement|null = null

    setInterval(() => {
      const originalHash = CryptoJS.SHA256(formData.value.content.trim()).toString()
      if (formDataHash.value === originalHash) {
        return
      }
      formDataHash.value = originalHash
      optionArticle('put', formData.value).then(() => {
        if (!mdEditorFooterRight) {
          mdEditorFooterRight = document.querySelector('.md-editor-footer-right')
          mdEditorFooterItem = document.createElement('div')
          mdEditorFooterItem.classList.add('md-editor-footer-item')
          mdEditorFooterLabel = document.createElement('label')
          mdEditorFooterLabel.classList.add('md-editor-footer-label')
          mdEditorFooterLabel.style.color = '#f56c6c'
          mdEditorFooterItem.append(mdEditorFooterLabel)
          mdEditorFooterRight?.insertBefore(mdEditorFooterItem, mdEditorFooterRight.firstChild)
        }
        if (mdEditorFooterLabel) {
          mdEditorFooterLabel.textContent = `${newTime()}自动保存`
        }
      })
    }, 3000)
  }
}

initArticleInfo(route.query)

watch(() => route.query, (value) => {
  initArticleInfo(value)
})

const toolbars: ToolbarNames[] = [
  0, // 自定义菜单
  '-',
  'bold',
  'underline',
  'italic',
  'strikeThrough',
  '-',
  'title',
  'sub',
  'sup',
  'quote',
  'unorderedList',
  'orderedList',
  'task',
  '-',
  'codeRow',
  'code',
  'link',
  'image',
  'table',
  'mermaid',
  'katex',
  1, // 自定义菜单
  '-',
  'save',
  '=',
  'preview',
  'previewOnly',
  'catalog'
]
const treeConfig: VxeTreeSelectPropTypes.TreeConfig = {
  keyField: 'value',
  expandAll: true,
  parentField: 'pid'
}
const previewTheme = ref<PreviewThemes>('default')
const completions = ref<Array<CompletionSource>>([
  (context) => {
    const word = context.matchBefore(/@\w*/)

    if (word === null || (word.from === word.to && context.explicit)) {
      return null
    }

    return {
      from: word.from,
      options: [
        {
          label: '@imzbf',
          type: 'text'
        }
      ]
    }
  }
])

const formRules = reactive<VxeFormPropTypes.Rules>({
  title: [
    { required: true, content: '文档标题不能为空' }
  ],
  tags: [
    { required: true, content: '文档标签不能为空' }
  ],
  categoryID: [
    { required: true, content: '文档分类不能为空' }
  ],
  summary: [
    { required: true, content: '文档摘要不能为空' }
  ]
})

config({
  markdownItPlugins (plugins) {
    return plugins.map((item) => {
      if (item.type === 'taskList') {
        return {
          ...item,
          options: {
            ...item.options,
            enabled: true
          }
        }
      }
      return item
    })
  }
})

onMounted(() => {
  const mediaQuery: MediaQueryList = window.matchMedia('(prefers-color-scheme: dark)')
  mediaQuery.addEventListener('change', (e) => {
    if (e.matches) {
      theme.value = 'dark'
    } else {
      theme.value = 'light'
    }
  })
})

const backArticle = () => {
  router.push({ name: 'Article', query: { cid: route.query.cid, type: route.query.type } })
}

const onSave = () => {
  showPopup.value = true
}

const submitEvent: VxeFormEvents.Submit = () => {
  optionArticle(route.params.type === 'add' ? 'post' : 'put', formData.value).then(res => {
    showPopup.value = false
    VxeUI.modal.message({ content: res.message, status: 'success' })
    if (route.params.type === 'add') {
      router.push({ name: 'EditArticle', params: { type: 'edit' }, query: { id: res.data.id, ...route.query } })
    }
  })
}

const resetEvent: VxeFormEvents.Reset = () => {
  VxeUI.modal.message({ content: '重置事件', status: 'info' })
}

const onUploadImg = async (files: File[], callback: (urls: string[]) => void) => {
  const res = await Promise.all(
    files.map((file) => {
      // eslint-disable-next-line promise/param-names
      return new Promise((rev, rej) => {
        const form = new FormData()
        form.append('file', file)

        postUploadImg(form)
          .then((res) => rev(res.data))
          .catch((error) => rej(error))
      })
    })
  )
  callback(res.map((item: any) => item))
}
</script>

<template>
  <div>
    <MdEditor
        :theme="theme"
        ref=editorRef
        v-model="formData.content"
        id="md-editor-id"
        codeTheme="github"
        catalogLayout="flat"
        :previewTheme="previewTheme"
        placeholder="请输入文章内容"
        :autoFocus="true"
        :autoDetectCode="true"
        :pageFullscreen="false"
        :showToolbarName="false"
        :codeFoldable="false"
        :noUploadImg="false"
        :completions="completions"
        :toolbars="toolbars"
        :toolbarsExclude="['github', 'htmlPreview']"
        :tableShape="[6, 4, 10, 8]"
        @onSave="onSave"
        @onUploadImg="onUploadImg">
      <template #defToolbars>
        <vxe-button title="返回文章列表" @click="backArticle" mode="text" class="md-editor-toolbar-item back-article" icon="vxe-icon-arrows-left"></vxe-button>
        <Emoji />
      </template>
    </MdEditor>

    <vxe-modal :title="modelTitle" v-model="showPopup" :width="650" :height="450" :draggable="false" lock-scroll>
      <vxe-form className="md-edit-form" ref="formRef" :rules="formRules" :data="formData" @submit="submitEvent" @reset="resetEvent">
        <vxe-form-item title="标题" field="title" span="24" :item-render="{}">
          <template #default>
            <vxe-input placeholder="请输入文档标题" show-word-count max-length="100" v-model="formData.title"></vxe-input>
          </template>
        </vxe-form-item>
        <vxe-form-item title="属于" field="categoryID" span="24" :item-render="{}">
          <template #default>
            <vxe-tree-select v-model="formData.categoryID" :tree-config="treeConfig" :options="treeList" :option-props="{label: 'title', value: 'id'}"></vxe-tree-select>
          </template>
        </vxe-form-item>
        <vxe-form-item title="标签" field="tags" span="24" :item-render="{}">
          <template #default>
            <vxe-input placeholder="标签使用顿号分割，eg：人工智能、AI、ChatGPT、DeepSeek" show-word-count max-length="100" v-model="formData.tags"></vxe-input>
          </template>
        </vxe-form-item>
        <vxe-form-item title="摘要" field="summary" span="24" :item-render="{}">
          <template #default>
            <vxe-textarea placeholder="请输入文档摘要(简介)" show-word-count resize="vertical" max-length="300" rows="5" v-model="formData.summary"></vxe-textarea>
          </template>
        </vxe-form-item>
        <vxe-form-item className="md-edit-submit-form" align="right" span="24">
          <vxe-button status="error" content="使用 AI 总结"></vxe-button>
          <vxe-button type="submit" status="primary" content="保存内容"></vxe-button>
          <vxe-button type="reset" content="重置表单"></vxe-button>
        </vxe-form-item>
      </vxe-form>
    </vxe-modal>
  </div>
</template>

<style scoped lang="scss">
.md-edit-form {
  height: 100%;

  :deep(.md-edit-submit-form) {
    position: absolute;
    bottom: 0;
  }
}

#md-editor-id {
  height: 100vh;

  .back-article {
    :deep(i) {
      font-weight: bold;
      font-size: 16px;
    }
  }

  :deep(.md-editor-code-head) {
    z-index: 1000;
  }

  :deep(.md-editor-toolbar) {
    .md-editor-toolbar-item {
      margin: 0 2px;
      padding: unset;
      width: 32px;
      height: 32px;
      display: flex;
      justify-content: center;
      align-items: center;

      svg.md-editor-icon {
        width: 20px;
        height: 20px;
      }
    }
    .md-editor-divider {
      margin: 0 2px;
    }
  }
}
</style>
