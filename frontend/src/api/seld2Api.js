const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

async function request(path, options = {}) {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
    ...options,
  });

  const data = await response.json().catch(() => null);

  if (!response.ok) {
    throw new Error(data?.error ?? "No se pudo completar la solicitud.");
  }

  return data;
}

export function getSeld2Rules() {
  return request("/seld2/rules");
}

export function inferSeld2(payload) {
  return request("/seld2/infer", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}
