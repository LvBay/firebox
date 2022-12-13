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
    <n-space justify="space-between" style="margin-top: 30px">
      <n-space vertical>
        <div v-for="item in matchDetail.participantIdentities.slice(0,5)">
          <n-space>
            <n-space vertical>
              <n-space>
                <div>英雄1</div>
                <div>装备1</div>
                <div>装备2</div>
                <n-button ghost type="success" size="tiny" @click="likeReport(item,matchDetail.gameId,2)">
                  <n-icon><Heart /></n-icon>
                </n-button>
                <n-button ghost type="error" size="tiny" @click="likeReport(item,matchDetail.gameId,1)">
                  <n-icon><Exclamation /></n-icon>
                </n-button>
              </n-space>
              <n-space>
                <div>闪现</div>
                <div>点燃</div>
                <div>{{item.player.summonerName}} {{item.stats.kills}}-{{item.stats.deaths}}-{{item.stats.assists}}</div>
              </n-space>
            </n-space>
          </n-space>
        </div>
      </n-space>
      <n-space vertical style="margin-left: 30px">
        <div v-for="item in matchDetail.participantIdentities.slice(5,10)">
          <n-space>
            <n-space vertical>
              <n-space>
                <div>英雄1</div>
                <div>装备1</div>
                <div>装备2</div>
                <n-button ghost type="success" size="tiny" @click="likeReport(item,matchDetail.gameId,2)">
                  <n-icon><Heart /></n-icon>
                </n-button>
                <n-button ghost type="error" size="tiny" @click="likeReport(item,matchDetail.gameId,1)">
                  <n-icon><Exclamation /></n-icon>
                </n-button>
              </n-space>
              <n-space>
                <div>闪现</div>
                <div>点燃</div>
                <div>{{item.player.summonerName}} {{item.stats.kills}}-{{item.stats.deaths}}-{{item.stats.assists}}</div>
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
import {OpenFileDialog,GetMatchInfo,CreateReport} from "../../wailsjs/go/main/App.js";
import { Exclamation } from '@vicons/fa'
import { Heart } from '@vicons/ionicons5'
import { useMessage, useDialog } from 'naive-ui'

const props = defineProps({
  currentGameId: {
    type: Number
  }
});

const matchDetail = ref({participantIdentities:[]})

const message = useMessage()
const dialog = useDialog()

function openDialog() {
  OpenFileDialog()
}

watch(props,function (val) {
  console.log('currentGameId',val.currentGameId)
  getMatchDetail(val.currentGameId)
})

function likeReport(item,gameId,reportType){
  let r = {
    reason:"firebox",
    reportedName:item.player.summonerName,
    extInfo:JSON.stringify({gameId:gameId}),
    reportType:reportType,
  }
  let title = '举报'
  if (reportType==2){
    title='点赞'
  }
  dialog.warning({
    title: title,
    content: title+' '+ r.reportedName+'?',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      CreateReport(r).then(result=>{
        console.log('举报成功')
      })
    },
    onNegativeClick: () => {
    }
  })
}

function getMatchDetail(id){
  if (!id){
    id = 'test'
  }
  GetMatchInfo(id.toString()).then(result=>{
    console.log('match info:',result)
    matchDetail.value= result
    // result.participants
  })
  // if (!id){
  //   id = props.currentGameId
  // }
  // matchDetail.value = {
  //   isWin:props.currentGameId==1,
  //   spendTime:'34分钟',
  //   team1:[
  //     {name:'aafdasdfasdf'+props.currentGameId,kda:'1/3/3'},
  //     {name:'b'+props.currentGameId,kda:'2/3/3'},
  //     {name:'c'+props.currentGameId,kda:'3/3/3'},
  //     {name:'d'+props.currentGameId,kda:'4/3/3'},
  //     {name:'e'+props.currentGameId,kda:'5/3/3'},
  //   ],
  //   team2:[
  //     {name:'aa',kda:'1/3/3'},
  //     {name:'bb',kda:'2/3/3'},
  //     {name:'cc',kda:'3/3/3'},
  //     {name:'dd',kda:'4/3/3'},
  //     {name:'ee',kda:'5/3/3'},
  //   ]
  // }
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