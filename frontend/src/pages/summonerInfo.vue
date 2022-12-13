<template>
  <n-space class="body" >
    <div >
      <n-space vertical style="border-right: 1px solid white;height: 100%">
        <div class="list" v-for="(data,index) in matchList" @click="changeCurrentGameId(data)" style="width: 100px;cursor:pointer" >
            <n-space>
              <div>11-9</div>
              <div>胜利</div>
            </n-space>
            <n-space justify="center">
              <div>单双排</div>
            </n-space>
        </div>
      </n-space>
    </div>
    <n-dialog-provider>
      <MatchDetail :currentGameId=currentGameId> detail </MatchDetail>
    </n-dialog-provider>
  </n-space>
</template>

<script setup>
import {ref,onMounted} from 'vue'
import MatchDetail from '../components/matchDetail.vue'
import {GetCurrentMatchList} from '../../wailsjs/go/main/App'

const matchList= ref([
  {gameId:1,isWin:true,kills:1,deaths:2,assists:3, detail:{}},
  {gameId:2,isWin:true,kills:1,deaths:2,assists:3, detail:{}},
])

function getCurrentMatchList(){
  GetCurrentMatchList().then(result=>{
    console.log('result',result)
    matchList.value = result.games.games
    if (matchList.value.length>0){
      currentGameId.value = matchList.value[0].gameId
    }
  })
}

onMounted(getCurrentMatchList)

const currentGameId = ref(matchList.value[0].gameId)

function changeCurrentGameId(data){
  currentGameId.value = data.gameId
}

</script>

<style>
.body{
  margin-top: 30px;
}

.list{
  margin-left: 30px;
  margin-right: 30px;
}
</style>