<template>
  <el-row>
    <el-col>
      <el-descriptions title="域名详情" :column="2">
        <el-descriptions-item label="域名">kooriookami</el-descriptions-item>
        <el-descriptions-item label="创建人">kooriookami</el-descriptions-item>
        <el-descriptions-item label="更新时间"
          >18100000000</el-descriptions-item
        >
        <el-descriptions-item label="描述"
          >No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu
          Province</el-descriptions-item
        >
      </el-descriptions>
    </el-col>
    <el-col>
      <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane label="权限管理" name="first">User</el-tab-pane>
        <el-tab-pane label="路径映射" name="second">Config</el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>
<script setup>
import { useMounted } from '@vueuse/core';

const tableData = [
  {
    id: "2016-05-03",
    name: "Tom",
    email: "California",
    createdAt: "Los Angeles",
    updatedAt: "No. 189, Grove St, Los Angeles",
  },
];
useMounted(()=>{
  
})
console.log(123);
</script>

<style></style>
