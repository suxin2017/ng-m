import { createRouter, createWebHistory } from "vue-router";
import DomainView from "../views/DomainView.vue";
import DomainDetailView from "../views/DomainDetailView.vue";
import AccountView from "../views/AccountView.vue";
import DashBoard from "../views/DashBoard.vue";
import Login from "../views/login/Login.vue";
import Layout from "../layout/Layout.vue";
import Terminal from "../views/Terminal.vue";
export const adminRoutes = [
  {
    path: "/dashbodar",
    meta: {
      title: "机器概览",
    },
    component: DashBoard,
  },
  {
    path: "/domain",
    meta: {
      title: "域名管理",
    },
    component: DomainView,
    children: [
      {
        path: ":id",
        name: "domain",
        meta: {
          title: "域名详情",
        },
        component: DomainDetailView,
      },
    ],
  },
  {
    path: "/account",
    meta: {
      title: "账户管理",
    },
    component: AccountView,
  },
  {
    path: "/terminal",
    meta: {
      title: "终端访问",
    },
    component: Terminal,
  },
];
const routes = [
  {
    path: "/login",
    component: Login,
    meta: {
      title: "登录",
      requiresAuth: false,
    },
  },
  {
    path: "/",
    redirect: adminRoutes[0].path,
    component: Layout,
    children: [...adminRoutes],
  },
];
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});
let token = localStorage.getItem("token");
router.beforeEach((to, from, next) => {
  console.log(to.path);
  if (!token && to.meta.requiresAuth !== false) {
    next({
      path: "/login",
      query: { rediect: to.fullPath },
    });
  } else {
    next();
  }
});
export default router;
export { routes };
