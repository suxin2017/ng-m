import { r } from "./r";

export async function nginxStatus() {
  return await r.get("/nginx/status");
}

export async function startNginx() {
  return await r.get("/nginx/start");
}
export async function stopNginx() {
  return await r.get("/nginx/stop");
}
export async function reloadNginx() {
  return await r.get("/nginx/reload");
}
