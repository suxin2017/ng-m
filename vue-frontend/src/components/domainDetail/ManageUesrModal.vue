<template>
  <el-dialog v-model="dialogVisible" title="路径">
    <el-form
      ref="formRef"
      :label-position="labelPosition"
      label-width="100px"
      :model="formModal"
      style="max-width: 460px"
    >
      <el-form-item label="用户" prop="selectUser">
        <el-select
          v-model="formModal.selectUser"
          filterable
          placeholder="Select"
        >
          <el-option
            v-for="item in userList"
            :key="item.id"
            :label="item.name ?? item.email"
            :value="item.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">创建</el-button>
        <el-button @click="resetForm(formRef)">返回</el-button>
      </el-form-item></el-form
    >
  </el-dialog>
</template>

<script setup>
import { addUserToDomain } from "@/api/domain";
import { userListUseGet } from "@/api/user";
import { ref, reactive, defineEmits, onMounted } from "vue";
import { useRoute } from "vue-router";
const emit = defineEmits(["ok"]);
const formRef = ref();
const userList = ref([]);
const formModal = reactive({
  selectUser: "",
});
const dialogVisible = ref(false);
const route = useRoute();
const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      let data = {
        domainId: +route.params["id"],
        userId: formModal.selectUser,
      };
      await addUserToDomain(data);
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
  dialogVisible.value = true;
};

onMounted(async () => {
  const { data, total } = await userListUseGet({ pageNum: 0, pageSize: 1000 });

  userList.value = data;
});

defineExpose({ openModal });
</script>
