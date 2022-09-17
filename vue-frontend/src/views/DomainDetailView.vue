<template>
  <el-row>
    <el-col>
      <el-collapse>
        <el-collapse-item title="域名详情" name="1">
          <el-descriptions :column="2">
            <el-descriptions-item label="域名">{{
              detail.domain
            }}</el-descriptions-item>
            <el-descriptions-item label="描述">{{
              detail.desc
            }}</el-descriptions-item>
            <el-descriptions-item label="创建人">{{
              detail.createdUser?.name ?? detail.createdUser?.email
            }}</el-descriptions-item>
            <el-descriptions-item label="更新人">{{
              detail.updatedUser?.name ?? detail.updatedUser?.email
            }}</el-descriptions-item>
            <el-descriptions-item label="创建时间">{{
              formatDate(detail.createdAt ?? 0)
            }}</el-descriptions-item>
            <el-descriptions-item label="更新时间">{{
              formatDate(detail.updatedAt ?? 0)
            }}</el-descriptions-item>
          </el-descriptions>
        </el-collapse-item>
      </el-collapse>
    </el-col>
    <el-col>
      <el-tabs v-model="activeName">
        <el-tab-pane label="路径映射" name="path"><domain-path /></el-tab-pane>
        <el-tab-pane label="权限管理" name="auth">User</el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>
<script setup>
import { infoDomain } from "@/api/domain";
import { formatDate } from "@/utils";
import { onMounted, reactive, ref, watch, watchEffect } from "vue";
import { useRoute, useRouter } from "vue-router";
import DomainPath from "../components/domainDetail/DomainPath.vue";
const router = useRouter();
const detail = reactive({});

const getDomainInfo = async () => {
  const data = await infoDomain(route.params.id);
  Object.assign(detail, data);
};

const activeName = ref("path");
watch(() => {
  router.replace({
    query: {
      key: activeName.value,
    },
  });
  console.log(activeName.value);
});
onMounted(() => {
  getDomainInfo();
  const key = router.currentRoute.value.query["key"];
  if (key) {
    activeName.value = key;
  }
});
</script>

<style></style>
