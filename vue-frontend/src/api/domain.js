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
