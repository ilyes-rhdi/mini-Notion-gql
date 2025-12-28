const API_BASE = import.meta.env.VITE_API_BASE || "http://localhost:8080";
const API_URL = `${API_BASE}/auth`;

export const register = async (data) => {
  const res = await fetch(`${API_URL}/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  let json = null;
  try {
    json = await res.json();
  } catch {
    // ignore
  }

  if (!res.ok) {
    throw new Error((json && (json.message || json.error)) || "Register failed");
  }
  return json;
};

export const login = async (data) => {
  const res = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  let json = null;
  try {
    json = await res.json();
  } catch {
    // ignore
  }

  if (!res.ok) {
    throw new Error((json && (json.message || json.error)) || "Login failed");
  }
  return json;
};

export const verifyOTP = async (userId, otp) => {
  const res = await fetch(`${API_URL}/verify?id=${encodeURIComponent(userId)}&otp=${encodeURIComponent(otp)}`);

  let json = null;
  try {
    json = await res.json();
  } catch {
    // ignore
  }

  if (!res.ok) {
    throw new Error((json && (json.message || json.error)) || "Verification failed");
  }
  return json;
};
