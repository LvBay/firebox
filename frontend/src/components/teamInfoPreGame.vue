<template>
  <div>
    <n-space vertical>
      <n-space justify="space-evenly" style="margin-top: 30px">
        <n-space vertical>
          <div v-for="(player,index) in teamInfo.friendTeam">
            <n-space>
              <n-space vertical>
                <n-space>
                  <img class="itemClass" :src="getChampImgUrl(player.championId)">
                </n-space>
                <n-space>
                  <div>{{player.summonerName}}</div>
                </n-space>
              </n-space>
            </n-space>
          </div>
        </n-space>
        <n-space vertical style="margin-left: 30px">
          <div v-for="(player,index) in teamInfo.enemyTeam">
            <n-space>
              <n-space vertical>
                <n-space>
                  <img class="itemClass" :src="getChampImgUrl(player.championId)">
                </n-space>
                <n-space>
                  <div>{{player.summonerName}}</div>
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
import {onMounted, ref} from "vue";
import {GetEnemySummonerIdAndSummonerName,GetMatchInfo,CreateReport} from "../../wailsjs/go/main/App.js";
import {champDict} from "../utils/lolDataList";
const teamInfo = ref({friendTeam:[],enemyTeam:[]})

function getChampImgUrl(championId){
  if (!championId){
    championId=1
  }
  return `https://game.gtimg.cn/images/lol/act/img/champion/${champDict[String(championId)].alias}.png`
}

function getGamePhaseSession(id){
  GetEnemySummonerIdAndSummonerName().then(result=>{
    console.log('match info:',result)
    teamInfo.value= result
  })
}

onMounted(getGamePhaseSession)

</script>

<style>

.itemClass{
  width: 28px;
  height: 28px;
  border-radius: 3px;
}
</style>