import { createGlobalState, useMounted } from "@vueuse/core";
import { onMounted, reactive } from "vue";
import { r } from "./r";

export async function loginUsePost(data) {
  return await r.post("/login", data);
}

export async function logoutUseGet() {
  return await r.get("/logout");
}

export async function currentUserUseGet() {
  return await r.get("/user/info", {
    maxRedirects: 0,
  });
}

export async function userListUseGet(params) {
  return await r.get("/user/list", { params });
}

export const useGlobalUserState = createGlobalState(() => {
  const user = reactive({
    id: undefined,
    email: "",
    password: "",
    name: "",
    avatar: "",
  });

  onMounted(async () => {
    if (!user.id) {
      const remoteUser = await currentUserUseGet();
      console.log("get user", remoteUser);
      Object.assign(user, remoteUser);
    }
  });

  const logout = async () => {
    await logoutUseGet();
  };
  return { user, logout };
});
