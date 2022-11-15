<template>
  <n-space class="body" >
    <div>
      <n-space vertical>
        <div v-for="(data,index) in matchList" @click="changeCurrentGameId(data)">
          {{ data.gameId }}
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

const currentGameId = ref(0)

function changeCurrentGameId(data){
  currentGameId.value = data.gameId
}

</script>

<style>
.body{
  margin-top: 30px;
}
</style>