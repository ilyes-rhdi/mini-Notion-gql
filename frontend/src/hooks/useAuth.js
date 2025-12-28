// IMPORTANT: file names are case-sensitive on Linux/macOS.
// The context hook lives in ../context/UseAuth.jsx
import useAuth from "../context/UseAuth";

export default function useAuthHook() {
  return useAuth();
}
