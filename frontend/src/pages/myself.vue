<template>
  <div class="top">
    <n-space>
      <div style="margin-right: 30px">
        <n-icon-wrapper :size="5" :border-radius="10">
          <img width="50" src="../assets/images/apple-touch-icon.png">
        </n-icon-wrapper>
      </div>
      <n-space>
        <div>
          <span style="margin-right: 20px">进步而非保守的 </span>
          <span style="margin-right: 20px">Lv 105 </span>
          <n-button type="primary" ghost>查询战绩</n-button>
        </div>
      </n-space>
    </n-space>
  </div>
  <div class="mid">
    <n-tabs type="segment" >
      <n-tab-pane name="chap1" tab="个人生涯" style="margin-top: 20px;">
        <n-space vertical>
          <n-table striped>
            <tbody>
            <tr>
              <td>单双排位</td>
              <td>黄金</td>
            </tr>
            <tr>
              <td>灵活排位</td>
              <td>黄金</td>
            </tr>
            <tr>
              <td>云顶之弈</td>
              <td>黄金</td>
            </tr>
            <tr>
              <td>荣誉等级</td>
              <td>黄金</td>
            </tr>
            </tbody>
          </n-table>
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="chap2" tab="英雄数据">
      </n-tab-pane>
      <n-tab-pane name="chap3" tab="永恒星碑">
        但是忽然，公寓的烟味消失，火警也停了。我的女朋友走进了房间，让我震惊的是，她摘下了自己的假发，她是
        Jeff Bezos（Amazon 老板）假扮的！
      </n-tab-pane>
    </n-tabs>
  </div>
  <div class="footer">
    <n-space justify="space-around">
      <n-icon size="30">
        <GameControllerOutline />
      </n-icon>
      <n-icon size="30">
        <LogoGithub />
      </n-icon>
      <n-icon size="30">
        <Tools />
      </n-icon>
      <n-icon size="30">
        <Fire />
      </n-icon>
    </n-space>
  </div>
</template>

<script setup>
import {NSpace,NButton,NCard,NIconWrapper,NProgress,NTabs,NTable, NTabPane,NIcon} from 'naive-ui'
import { GameControllerOutline, LogoGithub } from '@vicons/ionicons5'
import { Fire } from '@vicons/carbon'
import { Tools } from '@vicons/carbon'
</script>

<style scoped>
.top{
  padding-top: 30px;
  padding-left: 100px;
  text-align: left;
}

.mid{
  padding:30px;
}

.footer{
  padding:60px;
}
</style>