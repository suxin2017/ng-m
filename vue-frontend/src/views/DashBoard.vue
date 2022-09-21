<template>
  <div class="account">
    <el-button @click="start"> 启动 </el-button>
    <el-button @click="stop"> 停止 </el-button>
    <el-button @click="reload" :disabled="nginxInfo === 'stop'">
      重新加载配置
    </el-button>
    <div>
      <pre>{{ nginxInfo }}</pre>
    </div>
  </div>
</template>
<script setup>
import { nginxStatus, startNginx, stopNginx, reloadNginx } from "@/api/nginx";
import { onMounted, ref } from "vue";
const nginxInfo = ref("");

const getStatus = async () => {
  const status = await nginxStatus();
  console.log(status);
  nginxInfo.value = status;
};
onMounted(getStatus);
const sleep = () => {
  return new Promise((res) => {
    setTimeout(res, 500);
  });
};

const start = async () => {
  await startNginx();
  await sleep();
  await getStatus();
};
const stop = async () => {
  await stopNginx();
  await sleep();
  await getStatus();
};
const reload = async () => {
  await reloadNginx();
  await sleep();
  await getStatus();
};
</script>

<style></style>
