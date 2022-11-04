<script setup>
import {NButton} from 'naive-ui'
import {NSpace} from 'naive-ui'
</script>

<template>
  <div class="header">
  <n-space justify="center">
    <router-link class="rlink" to="/myself"><n-button type="primary" ghost @click="">我的</n-button></router-link>
    <router-link class="rlink" to="/beforeGame"><n-button type="info" ghost>对战前</n-button></router-link>
    <router-link class="rlink" to="/duringGame"><n-button type="success" ghost>对战中</n-button> </router-link>
    <router-link class="rlink"  to="/afterGame"><n-button type="warning" ghost>对战后</n-button></router-link>
  </n-space>
  </div>
  <router-view></router-view>
</template>

<style scoped>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}

.rlink{
  text-decoration: none;
}

.header{
  margin-top: 30px;
  margin-bottom: 20px;
}

</style>
