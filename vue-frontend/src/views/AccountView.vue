<template>
  <div class="account">
    <el-row>
      <el-col>
        <el-row>
          <el-col>
            <el-button @click="openModal" type="primary">新建账户</el-button>
          </el-col>
        </el-row>
      </el-col>
      <el-col>
        <el-table :data="tableData.data">
          <el-table-column prop="id" label="id"></el-table-column>
          <el-table-column prop="name" label="姓名"></el-table-column>
          <el-table-column prop="email" label="邮箱"></el-table-column>
          <el-table-column
            prop="createdAt"
            :formatter="formatTime"
            label="创建时间"
          ></el-table-column>
          <el-table-column
            prop="updatedAt"
            :formatter="formatTime"
            label="更新时间"
          ></el-table-column>
          <el-table-column prop="a" label="操作">
            <template #default>
              <el-button link type="primary" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
      <el-col>
        <el-row justify="center">
          <el-pagination layout="prev, pager, next" :total="tableData.total" />
        </el-row>
      </el-col>
    </el-row>
    <UserModal ref="usermodal" @ok="getList" />
  </div>
</template>
<script setup>
import { userListUseGet } from "@/api/user";
import { onMounted, reactive, ref } from "vue";
import UserModal from "../components/user/UserModal.vue";
import { formatTime } from "@/utils";
const tableData = reactive({ data: [], total: 0 });
const usermodal = ref(null);

const openModal = () => {
  usermodal.value.openModal();
};
const getList = async () => {
  const { data, total } = await userListUseGet({ pageNum: 0, pageSize: 15 });
  tableData.data = data;
  tableData.total = total;
};
onMounted(getList);
</script>

<style></style>
