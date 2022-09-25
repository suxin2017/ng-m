<template>
  <div>
    <el-row>
      <el-col>
        <el-button type="primary" @click="addPath">添加路径</el-button>
      </el-col>
      <el-col style="margin-top: 16px">
        <el-collapse v-model="collapseKeys">
          <el-collapse-item v-for="detail in tableData.data" :name="detail.id">
            <template #title>
              <div class="title">
                <div>
                  <el-space>
                    <span>id:</span>
                    <span>{{ detail.id }}</span>
                    <span>path:</span>
                    <span>{{ detail.path }}</span>
                  </el-space>
                </div>
                <div>
                  <el-button
                    @click.stop="addConfig(detail.id)"
                    link
                    type="primary"
                    size="small"
                    >配置路径</el-button
                  >
                  <el-button link type="primary" size="small">删除</el-button>
                </div>
              </div>
            </template>
            <el-descriptions :column="3">
              <el-descriptions-item label="id">{{
                detail.id
              }}</el-descriptions-item>
              <el-descriptions-item label="路径">{{
                detail.path
              }}</el-descriptions-item>
              <el-descriptions-item label="匹配规则">{{
                detail.matchRule
              }}</el-descriptions-item>
              <el-descriptions-item label="创建时间">{{
                formatDate(detail.createdAt ?? 0)
              }}</el-descriptions-item>
              <el-descriptions-item label="更新时间">{{
                formatDate(detail.updatedAt ?? 0)
              }}</el-descriptions-item>
              <el-descriptions-item label="类型">{{
                detail.type === 1 ? "静态资源" : "反向代理"
              }}</el-descriptions-item>
              <template v-if="detail.type === 1">
                <el-descriptions-item label="匹配路径">{{
                  detail.root
                }}</el-descriptions-item>
              </template>
              <template v-else>
                <el-descriptions-item label="服务地址"
                  >http://{{ detail.serverHost }}:{{
                    detail.serverPort
                  }}</el-descriptions-item
                >
              </template>
            </el-descriptions>
            <keep-alive>
              <PathMapper v-if="collapseKeys.includes(detail.id)"></PathMapper>
            </keep-alive>
          </el-collapse-item>
        </el-collapse>
      </el-col>
    </el-row>
    <PathConfigModal ref="modalRef"></PathConfigModal>
    <DomainPathModal ref="addPathModalRef" @ok="getList"></DomainPathModal>
  </div>
</template>
<script setup>
import { getPathList } from "@/api/domain";
import { formatDate } from "@/utils";
import { onMounted, reactive, ref } from "vue";
import { useRoute } from "vue-router";
import PathConfigModal from "./PathConfigModal.vue";
import PathMapper from "./PathMapper.vue";
import DomainPathModal from "./DomainPathModal.vue";
const addPathModalRef = ref(null);
const modalRef = ref(null);
const route = useRoute();
const collapseKeys = ref([]);
const tableData = reactive({
  total: 0,
  data: [],
});

const getList = async () => {
  const data = await getPathList({ domainId: route.params["id"] });
  tableData.data = data.data ?? [];
  tableData.total = data.total;
};
onMounted(() => {
  getList();
});
const addConfig = (pathId) => {
  modalRef.value.openModal(pathId);
};
const addPath = () => {
  addPathModalRef.value.openModal();
};
</script>
<style scoped>
.title {
  display: inline-flex;
  justify-content: space-between;
  width: 100%;
}
</style>
