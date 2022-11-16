<template>
<div>
  <n-space vertical>
    <n-space>
      <div>对局id {{currentGameId}}</div>
      <div v-if="matchDetail.isWin">胜 </div>
      <div v-else>败 </div>
      <div>{{matchDetail.spendTime}}</div>
    </n-space>
    <n-space>
      <n-space vertical>
        <div v-for="item in matchDetail.team1">
          <div>{{item.name}} {{item.kda}}</div>
        </div>
      </n-space>
      <n-space vertical>
        <div v-for="item in matchDetail.team2">
          <div>{{item.name}} {{item.kda}}</div>
        </div>
      </n-space>
    </n-space>
<!--    <n-button type="primary" @click="openDialog">选取lol启动文件</n-button>-->
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
      {name:'a',kda:'1/3/3'},
      {name:'b',kda:'2/3/3'},
      {name:'c',kda:'3/3/3'},
      {name:'d',kda:'4/3/3'},
      {name:'e',kda:'5/3/3'},
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