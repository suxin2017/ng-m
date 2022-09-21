<template>
  <div>
    <el-row>
      <el-col>
        <el-button type="primary" @click="addPath">添加路径</el-button>
      </el-col>
      <el-col> </el-col>
    </el-row>
    <DomainPathModal ref="modalRef"></DomainPathModal>
  </div>
</template>
<script setup>
import { ref } from "vue";
import DomainPathModal from "./DomainPathModal.vue";
const modalRef = ref(null);
const addPath = () => {
  modalRef.value.openModal();
};
</script>
