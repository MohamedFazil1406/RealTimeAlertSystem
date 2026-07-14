import client from "./client";

export async function getViolations(params?: any) {
  const response = await client.get("/violations/history", {
    params,
  });

  return response.data;
}
