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
