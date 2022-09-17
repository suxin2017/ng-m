<template>
  <div class="account">
    <el-row>
      <el-col>
        <el-row>
          <el-col>
            <el-button type="primary" @click="domainModal.openModal()"
              >新建域名</el-button
            >
          </el-col>
        </el-row>
      </el-col>
      <el-col>
        <el-table :data="tableData.data">
          <el-table-column prop="id" label="id"></el-table-column>
          <el-table-column prop="domain" label="域名"></el-table-column>
          <el-table-column prop="desc" label="描述"></el-table-column>
          <el-table-column
            prop="createdUser"
            label="创建者"
            :formatter="renderUserName"
          ></el-table-column>
          <el-table-column
            prop="updatedUser"
            label="更新者"
            :formatter="renderUserName"
          ></el-table-column>
          <el-table-column
            prop="createdAt"
            label="创建时间"
            :formatter="formatTime"
          ></el-table-column>
          <el-table-column
            prop="updatedAt"
            label="更新时间"
            :formatter="formatTime"
          ></el-table-column>
          <el-table-column prop="a" label="操作">
            <template #default="scope">
              <el-button link type="primary" size="small" @click="handleClick">
                <router-link
                  :to="{ name: 'domain', params: { id: scope.row.id } }"
                  >配置</router-link
                ></el-button
              >
              <el-button
                link
                type="primary"
                size="small"
                @click="handleEdit(scope.row)"
              >
                编辑
              </el-button>
              <el-button link type="primary" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <DomainModal @ok="getDomainList" ref="domainModal" />
  </div>
</template>
<script setup>
import { domainListUseGet } from "@/api/domain";
import { formatTime } from "@/utils";
import { onMounted, reactive, ref } from "vue";
import DomainModal from "@/components/domain/DomainModal.vue";
const domainModal = ref(null);
const tableData = reactive({
  data: [],
  total: 0,
});

const getDomainList = async () => {
  const { data, total } = await domainListUseGet();
  console.log(data, total);
  tableData.data = data;
  tableData.total = total;
};
onMounted(async () => {
  getDomainList();
});

const renderUserName = (row, column) => {
  const value = row[column.property];
  if (value == undefined) {
    return "";
  }
  return value.name ?? value.email;
};

const handleEdit = (row) => {
  domainModal.value.openModal(row);
};
</script>

<style></style>
