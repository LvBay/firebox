<template>
<div>
  <n-space vertical>
    <div>对局日期 1</div>
    <div>对局类型 2</div>
    <div>id: {{ currentGameId }}</div>
  </n-space>
</div>
</template>

<script setup>
import {ref} from "vue";

const props = defineProps({
  currentGameId: {
    type: Number
  }
});

const matchList = ref({time:"1",type:"test"})
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