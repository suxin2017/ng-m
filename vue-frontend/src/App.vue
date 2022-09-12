<script setup>
import { useRouter } from "vue-router";

const route = useRouter();
console.log(route);
</script>

<template>
  <el-config-provider>
    <router-view></router-view>
  </el-config-provider>
</template>

<style>
body {
  padding: 0;
  margin: 0;
}
</style>
