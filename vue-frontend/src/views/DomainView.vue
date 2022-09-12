<template>
  <div class="account">
    <el-row>
      <el-col>
        <el-row>
          <el-col>
            <el-button type="primary">新建域名</el-button>
          </el-col>
        </el-row>
      </el-col>
      <el-col>
        <el-table :data="tableData">
          <el-table-column prop="id" label="id"></el-table-column>
          <el-table-column prop="name" label="域名"></el-table-column>
          <el-table-column prop="email" label="管理员"></el-table-column>
          <el-table-column prop="createdAt" label="创建时间"></el-table-column>
          <el-table-column prop="updatedAt" label="更新时间"></el-table-column>
          <el-table-column prop="a" label="操作">
            <template #default>
              <el-button link type="primary" size="small" @click="handleClick">
                <router-link :to="{ name: 'domain', params: { id: 123 } }"
                  >配置</router-link
                ></el-button
              >
              <el-button link type="primary" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
  </div>
</template>
<script setup>
import { useMounted } from "@vueuse/core";
import { domainListUseGet } from "@/api/domain";
const tableData = [
  {
    id: "2016-05-03",
    name: "Tom",
    email: "California",
    createdAt: "Los Angeles",
    updatedAt: "No. 189, Grove St, Los Angeles",
  },
];
useMounted(async () => {
  domainListUseGet();
});
</script>

<style></style>
