<script lang="ts" setup>
import { computed, ref } from 'vue'

interface Novel {
  bookName: string
  author: string
  subscriber: number
  coin: number
  ratio: number
  words: number
  chapter: number
  actualChapter: number
  freeChapter: number
  commission: number
}

const data = ref<Novel>({
  bookName: '完美世界',
  author: '辰东',
  subscriber: 100000000,
  coin: 30421,
  ratio: 0.01,
  words: 3580000,
  chapter: 2090,
  actualChapter: 2014,
  freeChapter: 92,
  commission: 0.5
})

const copyText = ref('复制信息')
const price = computed(() => data.value.coin * data.value.ratio * data.value.subscriber)
const payChapter = computed(() => data.value.actualChapter - data.value.freeChapter)

const unit = (num: number) => {
  if (num >= 10000 && num < 100000000) {
    return `${roundTo(num / 10000)}万`
  }
  if (num >= 100000000 && num < 1000000000000) {
    return `${roundTo(num / 100000000)}亿`
  }
  if (num >= 1000000000000) {
    return `${roundTo(num / 1000000000000)}万亿`
  }
  return num
}

function roundTo (num: number, decimals: number = 2) {
  if (isNaN(num)) return 0
  const factor = Math.pow(10, decimals)
  return Math.round(num * factor) / factor
}

function copyToClipboard (text: string) {
  if (!navigator.clipboard) {
    copyText.value = '复制失败'
    console.warn('Clipboard API 不支持')
    return
  }
  navigator.clipboard.writeText(text).then(() => {
    copyText.value = '复制成功'
  }).catch(() => {
    copyText.value = '复制失败'
  })
  const timer = setTimeout(() => {
    copyText.value = '复制信息'
    clearTimeout(timer)
  }, 3000)
}

const onCopyInfo = () => {
  const val = data.value
  const info = ['基础信息']
  info.push(`书名：${val.bookName}`)
  info.push(`作者：${val.author}`)
  info.push(`订阅全书：${val.coin}(起点币)`)
  info.push(`单价比例：${val.ratio}(1角:10起点币)`)
  info.push(`全书订阅单价：${val.coin * val.ratio}(元)`)
  info.push(`平台抽成：${val.commission}(1:1)`)
  info.push(`订阅数量：${unit(val.subscriber)}(人)`)
  info.push(`小说字数：${unit(val.words)}(字)`)
  info.push(`总章节数量：${val.chapter}(章)`)
  info.push(`实际章节数：${val.actualChapter}(章)`)
  info.push(`免费章节数：${val.freeChapter}(章)`)
  info.push(`付费章节数量：${payChapter.value}(章)`)
  info.push('字数信息')
  info.push(`(总)平均每章：${roundTo(val.words / val.chapter)}(字)`)
  info.push(`(实)平均每章：${roundTo(val.words / val.actualChapter)}(字)`)
  info.push(`(费)平均每章：${roundTo(val.words / payChapter.value)}(字)`)
  info.push('收益信息')
  info.push(`总收益：${unit(price.value)}(元)`)
  info.push(`平台收益：${unit(price.value * val.commission)}(元)`)
  info.push(`作者收益：${unit(price.value * val.commission)}(元)`)

  copyToClipboard(info.join('\n'))
}
</script>

<template>
  <PageView class="container-box">
    <div class="container">
      <div class="container-items">
        <div class="container-item">
          <ul class="base-info">
            <li>
              <span>书名：</span>
              <div>
                <vxe-input placeholder="请输入书名" max-length="20" v-model="data.bookName"></vxe-input>
              </div>
            </li>
            <li>
              <span>作者：</span>
              <div>
                <vxe-input placeholder="请输入作者" max-length="20" v-model="data.author"></vxe-input>
              </div>
            </li>
            <li>
              <span>订阅全书：</span>
              <div>
                <vxe-input placeholder="请输入起点币" max-length="6" v-model="data.coin"></vxe-input>
                <span>(起点币)</span>
              </div>
            </li>
            <li>
              <span>单价比例：</span>
              <div>
                <vxe-input placeholder="默认(0.01)" v-model="data.ratio"></vxe-input>
                <span>(1角:10起点币)</span>
              </div>
            </li>
            <li>
              <span>平台抽成：</span>
              <div>
                <vxe-input placeholder="默认(0.5)" v-model="data.commission"></vxe-input>
                <span>(1:1)</span>
              </div>
            </li>
            <li>
              <span>订阅数量：</span>
              <div>
                <vxe-input placeholder="订阅数量" v-model="data.subscriber"></vxe-input>
                <span>{{ unit(data.subscriber) }}(人)</span>
              </div>
            </li>
            <li>
              <span>小说字数：</span>
              <div>
                <vxe-input placeholder="小说字数" max-length="9" v-model="data.words"></vxe-input>
                <span>{{ unit(data.words) }}(字)</span>
              </div>
            </li>
            <li>
              <span>总章节数量：</span>
              <div>
                <vxe-input placeholder="总章节数量" max-length="5" v-model="data.chapter"></vxe-input>
                <span>(章)</span>
              </div>
            </li>
            <li>
              <span>实际章节数量：</span>
              <div>
                <vxe-input placeholder="实际章节数量" max-length="5" v-model="data.actualChapter"></vxe-input>
                <span>(章)</span>
              </div>
            </li>
            <li>
              <span>免费章节数量：</span>
              <div>
                <vxe-input placeholder="免费章节数量" max-length="3" v-model="data.freeChapter"></vxe-input>
                <span>(章)</span>
              </div>
            </li>
          </ul>
          <ul class="calc-result">
            <li>
              <div>
                <div>基础信息</div>
                <div @click="onCopyInfo">{{ copyText }}</div>
              </div>
              <div><span>书名：</span>{{ data.bookName }}</div>
              <div><span>作者：</span>{{ data.author }}</div>
              <div><span>订阅全书：</span>{{ data.coin }}(起点币)</div>
              <div><span>单价比例：</span>{{ data.ratio }}(1角:10起点币)</div>
              <div><span>全书订阅单价：</span>{{ roundTo(data.coin * data.ratio) }}(元)</div>
              <div><span>平台抽成：</span>{{ data.commission }}(1:1)</div>
              <div><span>订阅数量：</span>{{ unit(data.subscriber) }}(人)</div>
              <div><span>小说字数：</span>{{ unit(data.words) }}(字)</div>
              <div><span>总章节数量：</span>{{ unit(data.chapter) }}(章)</div>
              <div><span>实际章节数量：</span>{{ unit(data.actualChapter) }}(章)</div>
              <div><span>免费章节数量：</span>{{ unit(data.freeChapter) }}(章)</div>
              <div><span>付费章节数量：</span>{{ payChapter }}(章)</div>
            </li>
            <li>
              <div>字数信息</div>
              <div><span>(总)平均每章：</span>{{ roundTo(data.words / data.chapter) }}(字)</div>
              <div><span>(实)平均每章：</span>{{ roundTo(data.words / data.actualChapter) }}(字)</div>
              <div><span>(费)平均每章：</span>{{ roundTo(data.words / payChapter) }}(字)</div>
            </li>
            <li>
              <div>收益信息</div>
              <div><span>总收益：</span>{{ unit(price) }}(元)</div>
              <div><span>平台收益：</span> {{ unit(price * data.commission) }}(元)</div>
              <div><span>作者收益：</span> {{ unit(price * data.commission) }}(元)</div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </PageView>
</template>

<style scoped lang="scss">
.container-box {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff;

  .container {
    ul {
      list-style: none;
    }

    .container-items {
      .container-item {
        ul {
          border: 1px solid var(--vxe-ui-input-border-color);
          padding: 10px 20px;
          border-radius: 5px;
          background-color: #fff;
        }

        ul.base-info {
          width: 400px;
          user-select: none;

          li {
            display: flex;
            align-items: flex-end;

            > span {
              display: block;
              width: 98px;
            }

            > div {
              flex: 1;
              position: relative;

              > span {
                position: absolute;
                right: 0;
                bottom: 3px;
              }

              :deep(.vxe-input) {
                width: 100%;
                border-radius: unset;
                border-top: unset;
                border-right: unset;
                border-left: unset;

                input {
                  padding: 13px 0 0 0;
                }
              }
            }
          }
        }

        ul.calc-result {
          width: 270px;
          position: fixed;
          top: calc(50% - 227px);
          right: 15px;

          li {
            div {
              display: flex;

              span {
                display: block;
                width: 98px;
              }
            }

            div:first-child {
              font-weight: bold;
              font-size: 16px;

              > div {
                &:first-child {
                  flex: 1;
                }

                &:last-child {
                  font-size: 14px;
                  font-weight: normal;
                  cursor: pointer;

                  &:hover {
                    color: #007bff;
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>
