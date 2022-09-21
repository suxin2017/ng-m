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
        label="路径"
        prop="path"
        :rules="[{ required: true, message: '请填写字段' }]"
      >
        <el-input v-model="formModal.path" />
      </el-form-item>
      <el-form-item label="匹配规则" prop="matchRule">
        <el-input v-model="formModal.matchRule" />
      </el-form-item>
      <el-form-item>
        <div>
          <div>~ 正则匹配，区分大小写</div>
          <div>~* 正则匹配，不区分大小写</div>
          <div>
            ^~
            普通字符匹配，如果该选项匹配，则，只匹配改选项，不再向下匹配其他选项
          </div>
          <div>= 普通字符匹配，精确匹配</div>
        </div>
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
import { ref, reactive } from "vue";
import { useRoute } from "vue-router";
const formRef = ref();
const formModal = reactive({
  path: "/",
  matchRule: "",
});
const dialogVisible = ref(false);
const route = useRoute();
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      addPath({ domainId: +route.params["id"], ...formModal });
      console.log(formModal);
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
