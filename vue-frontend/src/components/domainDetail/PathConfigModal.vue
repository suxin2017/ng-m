<template>
  <el-dialog v-model="dialogVisible" title="路径">
    <el-form
      ref="formRef"
      :label-position="labelPosition"
      label-width="100px"
      :model="formModal"
      style="max-width: 460px"
    >
      <el-form-item
        label="类型"
        prop="type"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-radio-group v-model="formModal.type">
          <el-radio :label="1">静态资源</el-radio>
          <el-radio :label="2">反向代理</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item
        v-if="formModal.type === 1"
        label="资源路径"
        prop="matchRule"
      >
        <el-input v-model="formModal.matchRule" />
      </el-form-item>
      <el-form-item v-else label="upstream配置" prop="matchRule">
        <el-input
          v-for="value in Array(10).fill(1)"
          v-model="formModal.matchRule"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">创建</el-button>
        <el-button @click="resetForm(formRef)">返回</el-button>
      </el-form-item></el-form
    >
  </el-dialog>
</template>

<script setup>
import { addPath } from "@/api/domain";
import { ref, reactive, defineEmits } from "vue";
import { useRoute } from "vue-router";
const emit = defineEmits(["ok"]);
const formRef = ref();
const formModal = reactive({
  type: 1,
  matchRule: "",
});
const dialogVisible = ref(false);
const route = useRoute();
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      await addPath({ domainId: +route.params["id"], ...formModal });
      emit("ok");
    } else {
      console.log("error submit!", fields);
    }
  });
};
const resetForm = (formEl) => {
  if (!formEl) return;
  formEl.resetFields();
  dialogVisible.value = false;
};

const openModal = () => {
  dialogVisible.value = true;
};

defineExpose({ openModal });
</script>
