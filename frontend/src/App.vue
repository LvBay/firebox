<script setup>
import {NButton,NSpace,NIcon, NIconWrapper} from 'naive-ui'
import {Dismiss12Filled,ChevronDoubleDown16Regular} from '@vicons/fluent'
import {Fire} from '@vicons/carbon'
</script>

<template>
  <div>
    <div class="title">
      <n-space justify="space-between">
        <div>
<!--          <n-icon size="18" :component="Fire" />-->
        </div>
        <div>
          <span style="margin-right: 10px">
            <n-icon size="18" :component="ChevronDoubleDown16Regular" />
        </span>
          <span>
            <n-icon size="18" :component="Dismiss12Filled" />
          </span>
        </div>
      </n-space>
    </div>
    <div class="header">
    <n-space justify="center">
      <router-link class="rlink" to="/myself"><n-button type="primary" ghost @click="">我的</n-button></router-link>
      <router-link class="rlink" to="/beforeGame"><n-button type="info" ghost>对战前</n-button></router-link>
      <router-link class="rlink" to="/duringGame"><n-button type="success" ghost>对战中</n-button> </router-link>
      <router-link class="rlink"  to="/afterGame"><n-button type="warning" ghost>对战后</n-button></router-link>
    </n-space>
    </div>
    <router-view></router-view>
  </div>
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
}
.title{
  height: 30px;
  padding: 10px 0px 10px 10px;
  --wails-draggable:drag;
}

</style>
