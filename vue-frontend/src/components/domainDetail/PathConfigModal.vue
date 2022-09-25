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
      <div v-if="formModal.type === 1">
        <el-form-item label="资源路径" prop="root">
          <resource-tree v-model="formModal.root"></resource-tree>
        </el-form-item>
      </div>
      <div v-else>
        <el-form-item label="ip" prop="serverHost">
          <el-input v-model="formModal.serverHost" />
        </el-form-item>
        <el-form-item label="port" prop="serverPort">
          <el-input v-model="formModal.serverPort" />
        </el-form-item>
      </div>

      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">创建</el-button>
        <el-button @click="resetForm(formRef)">返回</el-button>
      </el-form-item></el-form
    >
  </el-dialog>
</template>

<script setup>
import { getPathInfo, addPathRoot } from "@/api/domain";
import { ref, reactive, defineEmits } from "vue";
import { useRoute } from "vue-router";
import ResourceTree from "./ResourceTree.vue";
const emit = defineEmits(["ok"]);
const formRef = ref();
const formModal = reactive({
  type: 1,
  root: "/",
  serverHost: "",
  serverPort: "",
});
const dialogVisible = ref(false);
const route = useRoute();
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      let data = {
        pathId: formModal.id,
        domainId: +route.params["id"],
        type: formModal.type,
      };
      if (formModal.type === 1) {
        data.root = formModal.root.replace("//", "/");
      } else {
        data.ip = formModal.serverHost;
        data.prop = formModal.serverPort;
      }
      await addPathRoot(data);
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

const openModal = async (pathId) => {
  console.log(pathId);
  const data = await getPathInfo({ pathId });
  if (data) {
    data.root = "/" + data.root;
    Object.assign(formModal, data);
  } else {
    formEl.resetFields();
  }
  dialogVisible.value = true;
};

defineExpose({ openModal });
</script>
