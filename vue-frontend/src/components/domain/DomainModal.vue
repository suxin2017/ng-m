<template>
  <el-dialog v-model="dialogVisible" title="新建域名">
    <el-form
      ref="formRef"
      :label-position="labelPosition"
      label-width="100px"
      :model="formModal"
      style="max-width: 460px"
    >
      <el-form-item
        label="域名"
        prop="domain"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-input v-model="formModal.domain" />
      </el-form-item>
      <el-form-item
        label="描述"
        prop="desc"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-input v-model="formModal.desc" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">创建</el-button>
        <el-button @click="resetForm(formRef)">返回</el-button>
      </el-form-item></el-form
    >
  </el-dialog>
</template>

<script setup>
import { addDomain } from "@/api/domain";
import { ref, reactive } from "vue";
const formRef = ref();
const formModal = reactive({
  domain: "",
  desc: "",
});
const emit = defineEmits(["ok"]);
const dialogVisible = ref(false);
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      await addDomain(formModal);
      dialogVisible.value = false;
      resetForm();

      emit("ok");
      console.log("submit!");
    } else {
      console.log("error submit!", fields);
    }
  });
};

const resetForm = (formEl) => {
  if (!formEl) return;
  Object.assign(formModal, { id: 0, domain: "", desc: "" });
  formEl.resetFields();
  dialogVisible.value = false;
};

const openModal = (row) => {
  console.log(row);
  if (row) {
    Object.assign(formModal, row);
    console.log(formModal);
  }
  dialogVisible.value = true;
};

defineExpose({ openModal });
</script>
