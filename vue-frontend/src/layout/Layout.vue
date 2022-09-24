<script setup>
import { useGlobalUserState } from "@/api/user";
import { adminRoutes } from "@/router";
import { ref, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
const { user } = useGlobalUserState();
const router = useRouter();
const route = useRoute();
const activeIndex = ref(
  adminRoutes.map((item) => item.path).find((item) => route.path.includes(item))
);
const handleSelect = (index) => {
  activeIndex.value = index;
};
console.log(route.matched);
const routeMatched = computed(() => route.matched.slice(1));
const logout = () => {
  localStorage.removeItem("token");
  router.push({
    path: "/login",
    query: { rediect: route.fullPath },
  });
};
</script>

<template>
  <el-container class="main">
    <el-header class="header">
      <el-row class="header" justify="space-between" align="middle">
        <el-col class="logo" :span="12">
          <span>NG-M</span>
        </el-col>
        <el-col :span="12">
          <el-dropdown>
            <el-row :gutter="12" align="middle" class="user-box">
              <el-avatar size="small" :src="user.avatar"></el-avatar>
              <span style="margin-left: 16px">
                {{ user.name }}
              </span>
            </el-row>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="logout">
                  退出登录</el-dropdown-item
                ></el-dropdown-menu
              >
            </template>
          </el-dropdown>
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <el-aside class="aside">
        <el-menu
          v-for="route in adminRoutes"
          router="true"
          @select="handleSelect"
          :default-active="activeIndex"
        >
          <el-menu-item :index="route.path">{{
            route.meta.title
          }}</el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <el-row>
          <el-col class="breadcrumb">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item
                v-for="matched in routeMatched"
                :to="{ path: matched.path }"
                >{{ matched.meta.title }}</el-breadcrumb-item
              >
            </el-breadcrumb>
          </el-col>
          <el-col>
            <router-view v-if="route.matched.length === 3">
              <router-view></router-view>
            </router-view>
            <router-view v-slot="{ Component, route }" v-else>
              <transition name="fade" mode="out-in">
                <component :is="Component" />
              </transition>
            </router-view>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </el-container>
</template>
<style scoped>
.header {
  background-color: #001529;
  height: 48px;
  color: #fff;
}
.breadcrumb {
  margin-bottom: 20px;
}
.logo {
}
.aside {
  padding: 4px;
  width: 200px;
  min-height: calc(100vh - 48px);
}
.user-box {
  cursor: pointer;
  color: #fff;
}
.el-menu {
  border: none;
}
.el-menu-item {
  height: 48px;
  border-radius: 4px;
}
.el-menu-item.is-active {
  background-color: #e6f7ff;
}
.fade-enter {
  opacity: 0;
}
.fade-leave {
  opacity: 1;
}
.fade-enter-active {
  transition: opacity 0.2s;
}
.fade-leave-active {
  opacity: 0;
  transition: opacity 0.2s;
}
</style>
