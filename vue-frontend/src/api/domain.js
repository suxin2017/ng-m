import { createGlobalState, useMounted } from "@vueuse/core";
import { onMounted, reactive } from "vue";
import { r } from "./r";

export async function currentDomainUseGet() {
  return await r.get("/domain/info", {
    maxRedirects: 0,
  });
}

export async function domainListUseGet(params) {
  return await r.get("/domain/list", { params });
}

export async function addDomain(data) {
  return await r.post("/domain/add", data);
}

export async function infoDomain(id) {
  return await r.get("/domain/info", { params: { id } });
}

export async function domainResource(id, path = "") {
  return await r.get("/domain/resource", { params: { id, path } });
}

export async function addPath(params) {
  return await r.post("/domain/path/add", params);
}
export async function addPathRoot(params) {
  return await r.post("/domain/path/add/root", params);
}
export async function getPathInfo(params) {
  return await r.get("/domain/path/info", { params });
}
export async function getPathList(params) {
  return await r.get("/domain/path/list", { params });
}

export async function getNginxConfig(params) {
  return await r.get("/domain/nginxConfig", {
    params: {
      domainId: 10,
    },
  });
}

export async function previewNginxConfig(domainId){
  return await r.get("/domain/previewNginxConfig",{
    params:{
      domainId
    }
  })
}

export async function getUserListByDomain(params) {
  return await r.get("/domain/users", {
    params,
  });
}

export async function addUserToDomain(params) {
  return await r.post("/domain/user/add", params);
}
// getNginxConfig();
