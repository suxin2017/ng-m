<template>
  <el-row>
    <el-col>
      <el-button type="primary" @click="uploadResourecRef?.openModal"
        >上传资源</el-button
      >
    </el-col>
    <el-col>
      <el-tree
        v-if="reload"
        ref="tree"
        :default-expanded-keys="['/']"
        node-key="id"
        :props="props"
        :load="loadNode"
        lazy
      >
        <template #default="{ node, data }">
          <span class="custom-tree-node">
            <span>
              <span class="node-icon">
                <el-icon v-if="node?.data?.isDir"><Folder /></el-icon>
                <el-icon v-else><Document /></el-icon>
              </span>
              <el-button
                link
                type="primary"
                size="small"
                v-if="!node?.data?.isDir"
                @click="
                  showPreview(
                    node?.data?.id?.replace(
                      '//',
                      `/domain-preview/${domainId}/`
                    )
                  )
                "
              >
                {{ node.label }}
              </el-button>
              <span v-else>
                {{ node.label }}
              </span>
            </span>
            <span v-if="!node?.data?.isDir">{{
              prettyBytes(node?.data?.size ?? 0)
            }}</span>
          </span>
        </template></el-tree
      >
    </el-col>
  </el-row>
  <UplodaResourceModal
    @ok="reloadTree"
    ref="uploadResourecRef"
  ></UplodaResourceModal>
  <el-dialog v-model="previewRef" :width="800">
    <iframe :src="src" frameborder="0" width="100%" height="400px"></iframe>
    <a :href="src" target="_blank">新标签页打开</a>
  </el-dialog>
</template>
<script setup>
import { domainResource } from "@/api/domain";
import { useRoute } from "vue-router";
import prettyBytes from "pretty-bytes";
import UplodaResourceModal from "./UplodaResourceModal.vue";
import { nextTick, ref } from "vue";
const route = useRoute();
const uploadResourecRef = ref(null);
const tree = ref(null);
const reload = ref(true);
const domainId = route.params["id"];
const previewRef = ref(false);
const src = ref("");
const showPreview = (url) => {
  console.log(url);
  previewRef.value = true;
  src.value = url;
};

const loadNode = async (node, resolve) => {
  const data = await domainResource(route.params.id, node?.data?.id ?? "");
  data.map((item) => {
    item.id = `${node?.data?.id ?? node?.data?.path ?? ""}/${item.path}`;
  });
  data.sort((a, b) => {
    return a.isDir ? -1 : 1;
  });
  if (!node.data.path) {
    return resolve([
      {
        path: "/",
        id: "/",
        isDir: true,
        children: [data],
      },
    ]);
  }

  return resolve(data);
};
const reloadTree = () => {
  console.log("ok");
  reload.value = false;
  nextTick(() => {
    reload.value = true;
  });
};

const props = {
  label: "path",
  isLeaf: (_, node) => {
    return !node.data.isDir;
  },
};
</script>
<style scoped>
.custom-tree-node {
  display: inline-flex;
  width: 400px;
  justify-content: space-between;
}
.node-icon {
  display: inline-block;

  margin-right: 8px;
}
</style>
