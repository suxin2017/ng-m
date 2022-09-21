<template>
  <el-tree :props="props" :load="loadNode" lazy>
    <template #default="{ node, data }">
      <span class="custom-tree-node">
        <span>
          <span class="node-icon">
            <el-icon v-if="node?.data?.isDir"><Folder /></el-icon>
            <el-icon v-else><Document /></el-icon>
          </span>
          {{ node.label }}</span
        >
        <span v-if="!node?.data?.isDir">{{
          prettyBytes(node?.data?.size ?? 0)
        }}</span>
      </span>
    </template></el-tree
  >
</template>
<script setup>
import { domainResource } from "@/api/domain";
import { useRoute } from "vue-router";
import prettyBytes from "pretty-bytes";

const route = useRoute();

const loadNode = async (node, resolve) => {
  return resolve(await domainResource(route.params.id, node?.data?.path ?? ""));
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
