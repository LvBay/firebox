<template>
  <n-space class="body" >
    <div>
      <n-space vertical>
        <div class="list" v-for="(data,index) in matchList" @click="changeCurrentGameId(data)">
            <n-space>
              <div>logo</div>
              <div>5-3-16</div>
              <div>11-09</div>
            </n-space>
            <n-space justify="center">
              <div>单双排</div>
            </n-space>
        </div>
      </n-space>
    </div>
    <MatchDetail :currentGameId=currentGameId> detail </MatchDetail>
  </n-space>
</template>

<script setup>
import {ref} from 'vue'
import MatchDetail from '../components/matchDetail.vue'

const matchList= ref([
  {gameId:1,isWin:true,kills:1,deaths:2,assists:3, detail:{}},
  {gameId:2,isWin:true,kills:1,deaths:2,assists:3, detail:{}},
])

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