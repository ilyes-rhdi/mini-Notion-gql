import { Navigate } from "react-router-dom";

import useAuth from "../../hooks/useAuth";

export default function AppHome() {
  const { isAuthed } = useAuth();

  if (isAuthed) return <Navigate to="/app" replace />;

  return (
    <div className="page">
      <div className="container">
        <div className="ui-panel pad-6 auth-card">
          <div className="row-between">
            <div>
              <div className="auth-title">Mini Notion</div>
              <div className="auth-sub">Login to open your workspaces</div>
            </div>
            <div className="ui-tag">frontend</div>
          </div>

          <div className="auth-links">
            <a className="ui-btn-primary full" href="/auth/login">
              login
            </a>
            <a className="ui-btn-secondary full" href="/auth/signup">
              create account
            </a>
          </div>

          <div className="ui-card pad-4 muted small">
            <div>REST auth: {import.meta.env.VITE_API_BASE || "http://localhost:8080"}</div>
            <div>GraphQL: {import.meta.env.VITE_GQL_URL || "http://localhost:8080/graphql"}</div>
          </div>
        </div>
      </div>
    </div>
  );
}
