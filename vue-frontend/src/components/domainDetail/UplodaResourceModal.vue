<template>
  <el-dialog v-model="dialogVisible" title="上传资源">
    <el-upload
      ref="uploadRef"
      class="upload-demo"
      drag
      :headers="headers"
      :data="data"
      :on-success="onSuccess"
      action="/api/domain/resource/upload"
      :auto-upload="false"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        Drop file here or <em>click to upload</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          jpg/png files with a size less than 500kb
        </div>
      </template>
    </el-upload>
    <el-button class="ml-3" type="success" @click="submitUpload">
      上传
    </el-button>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, defineEmits, nextTick, onMounted } from "vue";
import { useRoute } from "vue-router";
import { getTokenHeader } from "@/api/r";
const emit = defineEmits(["ok"]);
const formRef = ref();
const formModal = reactive({
  type: 1,
  matchRule: "",
});
const dialogVisible = ref(false);
const route = useRoute();
const headers = reactive(getTokenHeader());
const data = reactive({ domainId: route.params["id"] });
const uploadRef = ref(null);

const submitUpload = () => {
  uploadRef.value?.submit();
};

const openModal = () => {
  dialogVisible.value = true;
};
const onSuccess = () => {
  dialogVisible.value = false;
  emit("ok");
  uploadRef.clearFiles();
};

defineExpose({ openModal });
</script>
