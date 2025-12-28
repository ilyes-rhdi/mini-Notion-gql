import { Link } from "react-router-dom";

export default function ErrorPage() {
  return (
    <div className="page">
      <div className="ui-panel pad-6 auth-card">
        <div className="stack">
          <div>
            <div className="h1">Page not found</div>
            <div className="auth-sub">The route you opened doesn't exist.</div>
          </div>

          <div className="row">
            <Link className="ui-btn-primary full" to="/app">go to app</Link>
            <Link className="ui-btn-secondary full" to="/auth/login">login</Link>
          </div>
        </div>
      </div>
    </div>
  );
}
