import { dev } from "$app/env";

export async function getIP() {
  const response = await fetch("https://www.cloudflare.com/cdn-cgi/trace");
  let data = await response.text();
  data = data
    .trim()
    .split("\n")
    .reduce(function (obj, pair) {
      pair = pair.split("=");
      return (obj[pair[0]] = pair[1]), obj;
    }, {});
  return data;
}

export function getByIP(ip) {
  const url = dev ? "http://localhost:8080" : "";
  const response = fetch(`${url}/${ip}.json`).then((r) => r.json());
  return response;
}
