<template>
  <el-tree-select
    ref="tree"
    :default-expanded-keys="['/']"
    node-key="id"
    :props="props"
    :load="loadNode"
    lazy
    v-model="propsData.value"
    check-strictly
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
                node?.data?.id?.replace('//', `/domain-preview/${domainId}/`)
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
    </template></el-tree-select
  >
</template>
<script setup>
import { domainResource } from "@/api/domain";
import { useRoute } from "vue-router";
import prettyBytes from "pretty-bytes";
import { nextTick, ref } from "vue";

const propsData = defineProps({
  value: String,
});
const route = useRoute();
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
  let data = await domainResource(route.params.id, node?.data?.id ?? "");
  data = data.filter((item) => {
    return item.isDir;
  });
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
