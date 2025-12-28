import { useEffect, useState } from "react";
import { AuthContext } from "./authContext";

export default function AuthProvider({ children }) {
  const [token, setToken] = useState(() => localStorage.getItem("token") || "");

  useEffect(() => {
    if (token) localStorage.setItem("token", token);
    else localStorage.removeItem("token");
  }, [token]);

  const logout = () => setToken("");

  const value = { token, setToken, isAuthed: !!token, logout };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
