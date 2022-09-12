<script setup>
import { loginUsePost } from "@/api/user";
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";

const formRef = ref();
const loginValue = reactive({
  email: "",
  password: "",
});
const router = useRouter();
const submitForm = async (formEl) => {
  formEl.validate(async (valid) => {
    if (valid) {
      const token = await loginUsePost(loginValue);
      if (token) {
        localStorage.setItem("token", token);

        window.location.replace(router.currentRoute.value.query.rediect ?? "/");
      }
    } else {
      return false;
    }
  });
};
</script>

<template>
  <div class="main">
    <el-card class="login-box">
      <el-form
        ref="formRef"
        label-position="top"
        label-width="100px"
        :model="loginValue"
        style="max-width: 460px"
      >
        <el-form-item
          label="邮箱"
          prop="email"
          :rules="[
            { required: true, message: '请填写邮箱' },
            { type: 'email', message: '邮箱格式不正确' },
          ]"
        >
          <el-input v-model="loginValue.email" />
        </el-form-item>
        <el-form-item
          label="密码"
          prop="password"
          :rules="[{ required: true, message: '请填写密码' }]"
        >
          <el-input show-password v-model="loginValue.password" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm(formRef)"
            >登录</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.main {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-box {
  width: 300px;
  border: 1px solid rgb(255, 255, 255);
  border-radius: 4px;
  padding: 24px;
}
</style>
