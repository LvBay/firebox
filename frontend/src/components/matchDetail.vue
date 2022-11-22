<template>
<div>
  <n-space vertical>
    <n-space>
      <n-space vertical>
        <div>对局日期</div>
        <div>9/1</div>
      </n-space>
      <n-space vertical>
        <div>对局类型</div>
        <div>单双排</div>
      </n-space>
      <n-space vertical>
        <div>开始时间</div>
        <div>20:01</div>
      </n-space>
      <n-space vertical>
        <div>胜/负</div>
        <div v-if="matchDetail.isWin">胜</div>
        <div v-else>负</div>
      </n-space>
      <n-space vertical>
        <div>数据显示</div>
        <div>输出伤害</div>
      </n-space>
      <n-space vertical>
        <div>双方比分</div>
        <div>48/64</div>
      </n-space>
      <n-space vertical>
        <div>经济差距</div>
        <div>48k/64k</div>
      </n-space>
      <n-space vertical>
        <div>对局时长</div>
        <div>48/64</div>
      </n-space>
    </n-space>
    <n-space justify="space-between">
      <n-space vertical>
        <div v-for="item in matchDetail.team1">
          <n-space>
            <div>英雄1</div>
            <n-space vertical>
              <n-space>
                <div>装备1</div>
                <div>装备2</div>
              </n-space>
              <n-space>
                <div>闪现</div>
                <div>点燃</div>
                <div>{{item.name}} {{item.kda}}</div>
              </n-space>
            </n-space>
          </n-space>
        </div>
      </n-space>
      <n-space vertical>
        <div v-for="item in matchDetail.team2">
          <n-space>
            <div>英雄1</div>
            <n-space vertical>
              <n-space>
                <div>装备1</div>
                <div>装备2</div>
              </n-space>
              <n-space>
                <div>闪现</div>
                <div>点燃</div>
                <div>{{item.name}} {{item.kda}}</div>
              </n-space>
            </n-space>
          </n-space>
        </div>
      </n-space>
    </n-space>
  </n-space>
</div>
</template>

<script setup>
import {onMounted, ref, watch} from "vue";
import {OpenFileDialog} from "../../wailsjs/go/main/App.js";

const props = defineProps({
  currentGameId: {
    type: Number
  }
});

const matchDetail = ref({type:"test"})

function openDialog() {
  OpenFileDialog()
}

watch(props,function (val) {
  getMatchDetail(val.currentGameId)
})

function getMatchDetail(id){
  if (!id){
    id = props.currentGameId
  }
  matchDetail.value = {
    isWin:props.currentGameId==1,
    spendTime:'34分钟',
    team1:[
      {name:'aafdasdfasdf'+props.currentGameId,kda:'1/3/3'},
      {name:'b'+props.currentGameId,kda:'2/3/3'},
      {name:'c'+props.currentGameId,kda:'3/3/3'},
      {name:'d'+props.currentGameId,kda:'4/3/3'},
      {name:'e'+props.currentGameId,kda:'5/3/3'},
    ],
    team2:[
      {name:'aa',kda:'1/3/3'},
      {name:'bb',kda:'2/3/3'},
      {name:'cc',kda:'3/3/3'},
      {name:'dd',kda:'4/3/3'},
      {name:'ee',kda:'5/3/3'},
    ]
  }
}

onMounted(getMatchDetail)

</script>

<style>

.spanTitle {
  color: #9AA4AF;
  font-size: 13px;
}

.alignCent {
  align-items: center;
}
</style>