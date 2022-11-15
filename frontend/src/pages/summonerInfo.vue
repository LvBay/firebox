<template>
  <n-space class="body" >
    <div>
      <n-space vertical>
        <div class="list" v-for="(data,index) in matchList" @click="changeCurrentGameId(data)">
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
import {OpenFileDialog} from "../../wailsjs/go/main/App.js";

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

.list{
  margin-left: 30px;
  margin-right: 30px;
}
</style>