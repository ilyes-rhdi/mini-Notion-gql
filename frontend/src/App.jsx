import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom";

import AuthProvider from "./context/AuthContext.jsx";
import useAuth from "./context/UseAuth";

import AppHome from "./pages/app/AppHome";
import WorkspacePage from "./pages/app/WorkspacePage";
import ProfilePage from "./pages/app/ProfilePage";

import Login from "./pages/auth/login";
import Signup from "./pages/auth/signup";
import Verify from "./pages/auth/verify";

import ErrorPage from "./pages/errors/ErrorPage";

function RequireAuth({ children }) {
  const { isAuthed } = useAuth();
  if (!isAuthed) return <Navigate to="/auth/login" replace />;
  return children;
}

const router = createBrowserRouter([
  { path: "/", element: <AppHome /> },

  { path: "/auth/signup", element: <Signup /> },
  { path: "/auth/login", element: <Login /> },
  { path: "/auth/verify", element: <Verify /> },

  { path: "/app", element: <RequireAuth> <WorkspacePage /></RequireAuth>   },
  { path: "/app/profile", element: <RequireAuth><ProfilePage /></RequireAuth> },
  { path: "/app/w/:workspaceId", element: <RequireAuth><WorkspacePage /></RequireAuth> },
  { path: "/app/w/:workspaceId/p/:pageId", element: <RequireAuth><WorkspacePage /></RequireAuth> },

  { path: "*", element: <ErrorPage /> },
]);

export default function App() {
  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
}
