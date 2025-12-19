<script setup lang="ts">
import ArticleContent from '@/views/article/ArticleContent.vue'
import { ref, reactive } from 'vue'
import { deleteCategory, postCategory } from '@/api/article'
import { CategoryAddVO, CategoryVO } from '@/api/user'
import ArticleNav from '@/views/article/components/ArticleNav.vue'
import { VxeFormEvents, VxeFormPropTypes, VxeModalPropTypes, VxeTreeSelectPropTypes, VxeUI } from 'vxe-pc-ui'
import { useRouter, useRoute } from 'vue-router'
import PageView from '@/views/layout/PageView.vue'
import { useArticleStore } from '@/store/article'

const route = useRoute()
const router = useRouter()
const isAdd = ref(false)
const showPopup = ref(false)
const showWinTitle = ref('添加文档分类')
const formData = ref<CategoryAddVO>({ docSort: '', id: '', pid: '', sort: 0, title: '', isExpand: true })

const defaultCategoryId = ref<string>('')
const treeConfig: VxeTreeSelectPropTypes.TreeConfig = {
  keyField: 'value',
  expandAll: true,
  parentField: 'pid'
}

const useArticle = useArticleStore()

const submitEvent: VxeFormEvents.Submit = () => {
  postCategory(isAdd.value, formData.value).then((res) => {
    showPopup.value = false
    VxeUI.modal.message({ content: res.message, status: 'success' })
    if (route.query.sort !== formData.value.docSort) {
      router.push({ name: 'Article', query: { cid: route.query.cid, type: route.query.type, sort: formData.value.docSort } })
    }
    useArticle.fetchCategoryTreeList(route.query, router)
  })
}

const resetEvent: VxeFormEvents.Reset = () => {
  VxeUI.modal.message({ content: '重置事件', status: 'info' })
}

function onAdd (item: CategoryVO) {
  isAdd.value = true
  formData.value.id = '0'
  formData.value.pid = (item.name !== 'all' && item.name !== 'default') ? item.id : '0'
  formData.value.sort = 0
  formData.value.isExpand = true
  formData.value.docSort = 'updatedAt'
  showWinTitle.value = '添加文档分类'
  showPopup.value = true
}

function onEdit (item: CategoryVO) {
  isAdd.value = false
  formData.value.id = item.id
  formData.value.pid = item.pid
  formData.value.title = item.title
  formData.value.sort = item.sort
  formData.value.isExpand = item.isExpand
  formData.value.docSort = item.docSort
  showWinTitle.value = `编辑[${item.title}]分类信息`
  showPopup.value = true
}

function onDelete (item: CategoryVO) {
  VxeUI.modal.confirm({
    title: '删除提示',
    content: `是否删除[${item.title}]分类信息？`,
    escClosable: true,
    status: 'warning',
    maskClosable: true,
    draggable: false,
    destroyOnClose: true
  }).then(rest => {
    if (rest !== 'confirm') return
    deleteCategory(item).then(result => {
      VxeUI.modal.message({
        content: result.message,
        status: 'success'
      })
      if (item.pid === '0') {
        if (!defaultCategoryId.value) {
          for (const item of useArticle.categoryList) {
            if (item.name === 'default') {
              defaultCategoryId.value = item.id
              break
            }
          }
        }
        router.push({ name: 'Article', query: { cid: defaultCategoryId.value, type: 'default' } })
      } else {
        router.push({ name: 'Article', query: { cid: item.pid, type: item.name } })
      }
      useArticle.fetchCategoryTreeList(route.query, router)
    })
  })
}

const beforeHideMethod: VxeModalPropTypes.BeforeHideMethod = () => {
  return new Promise((resolve) => {
    formData.value = {} as CategoryVO
    resolve(null)
  })
}

const formRules = reactive<VxeFormPropTypes.Rules>({
  title: [
    { required: true, content: '分类名称不能为空' }
  ],
  docSort: [
    { required: true, content: '必须选择文档排序' }
  ]
})

useArticle.fetchCategoryTreeList(route.query, router)
</script>

<template>
  <div class="article-box">
    <PageView class="article-category-box">
      <div class="article-category-title"><vxe-icon name="file-word"></vxe-icon><span>文章分类列表</span></div>
      <ArticleNav @add="onAdd" @edit="onEdit" @delete="onDelete" v-model="useArticle.categoryList" />
    </PageView>

    <ArticleContent />

    <vxe-modal :before-hide-method="beforeHideMethod" :title="showWinTitle" v-model="showPopup" :width="400" lock-scroll :destroyOnClose="true">
      <vxe-form vertical title-colon title-bold className="article-add-form" ref="formRef" :rules="formRules" :data="formData" @submit="submitEvent" @reset="resetEvent">
        <vxe-form-item title="名称" field="title" span="24" :item-render="{}">
          <template #default>
            <vxe-input placeholder="请输入分类名称" show-word-count :clearable="true" max-length="20" v-model="formData.title"></vxe-input>
          </template>
        </vxe-form-item>
        <vxe-form-item title="节点属于" field="pid" span="24" :item-render="{}">
          <template #default>
            <vxe-tree-select v-model="formData.pid" :tree-config="treeConfig" :options="useArticle.rootTreeList" :option-props="{label: 'title', value: 'id'}"></vxe-tree-select>
          </template>
        </vxe-form-item>
        <vxe-form-item title="文档排序" field="docSort" span="24" :item-render="{}">
          <template #default>
            <vxe-radio-group v-model="formData.docSort">
              <vxe-radio checked-value="title" content="标题[升序]"></vxe-radio>
              <vxe-radio checked-value="updatedAt" content="更新时间[降序]"></vxe-radio>
              <vxe-radio checked-value="createdAt" content="创建时间[降序]"></vxe-radio>
            </vxe-radio-group>
          </template>
        </vxe-form-item>
        <vxe-form-item title="子节点展开模式" field="isExpand" span="24" :item-render="{}">
          <template #default>
            <vxe-radio-group v-model="formData.isExpand">
              <vxe-radio :checkedValue="true" content="展开(默认)"></vxe-radio>
              <vxe-radio :checkedValue="false" content="收缩"></vxe-radio>
            </vxe-radio-group>
          </template>
        </vxe-form-item>
        <vxe-form-item title="菜单排序" field="sort" span="24" :item-render="{}">
          <template #default>
            <vxe-number-input placeholder="排序" show-word-count type="number" v-model="formData.sort"></vxe-number-input>
          </template>
        </vxe-form-item>
        <vxe-form-item className="article-add-submit-form" align="right" span="24">
          <vxe-button type="submit" status="primary" content="保存分类"></vxe-button>
          <vxe-button type="reset" content="重置表单"></vxe-button>
        </vxe-form-item>
      </vxe-form>
    </vxe-modal>
  </div>
</template>

<style scoped lang="scss">
.article-box {
  height: 100%;
  display: flex;
  gap: 20px;
  max-width: 1400px;
  margin: 0 auto;
  box-sizing: border-box;
  --brand-blue: #00AEEC;

  .article-add-form {
    :deep(.vxe-form--item-title) {
      height: unset;
      line-height: unset;
      user-select: none;
    }

    :deep(.vxe-radio--label) {
      user-select: none;
    }
  }

  .article-category-box {
    height: 100%;
    width: 340px;
    display: flex;
    min-width: unset;
    flex-direction: column;
    padding: unset;

    :deep(.vxe-menu) {
      padding: 5px;
    }

    .article-category-title {
      cursor: default;
      user-select: none;
      transition: .3s;
      gap: 5px;
      display: flex;
      align-items: center;
      padding: 16px 10px 10px 16px;

      i {
        font-size: 25px;
      }

      span {
        font-size: 18px;
        font-weight: bold;

        &:hover {
          color: var(--brand-blue);
        }
      }
    }
  }
}
</style>
