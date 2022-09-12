<template>
  <el-dialog v-model="dialogVisible" title="新建账户">
    <el-form
      ref="formRef"
      :label-position="labelPosition"
      label-width="100px"
      :model="formModal"
      style="max-width: 460px"
    >
      <el-form-item label="头像">
        <el-avatar :src="formModal.avatar"></el-avatar>
      </el-form-item>
      <el-form-item
        label="用户名"
        prop="name"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-input v-model="formModal.name" />
      </el-form-item>
      <el-form-item
        label="邮箱"
        prop="email"
        :rules="[
          { required: true, message: '请填写字段' },
          {
            type: 'email',
            message: '请填写正确邮箱',
          },
        ]"
      >
        <el-input v-model="formModal.email" />
      </el-form-item>
      <el-form-item
        label="密码"
        prop="password"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-input v-model="formModal.password" show-password />
      </el-form-item>
      <el-form-item label="二次确认" prop="password2" :rules="passwordRule">
        <el-input v-model="formModal.password2" show-password />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">创建</el-button>
        <el-button @click="resetForm(formRef)">返回</el-button>
      </el-form-item></el-form
    >
  </el-dialog>
</template>

<script setup>
import { ref, reactive } from "vue";
const formRef = ref();
const formModal = reactive({
  name: "",
  password: "",
  type: "",
  password2: "",
  avatar: `https://api.multiavatar.com/${Date.now()}.svg`,
});
const dialogVisible = ref(false);
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log("submit!");
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
const passwordRule = reactive([
  {
    validator: (rule, value, cb) => {
      if (value !== formModal.password) {
        console.log(value, "=========");
        cb(new Error("两次密码输入不一致"));
      }
    },
  },
  { required: true, message: "请填写字段" },
]);
defineExpose({ openModal });
</script>
