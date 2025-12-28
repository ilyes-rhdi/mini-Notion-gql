import { useQuery } from "@apollo/client/react";
import { useNavigate } from "react-router-dom";
import { useApolloClient } from "@apollo/client/react";

import useAuth from "../../hooks/useAuth";
import { ME, WORKSPACES } from "../../graphql/queries/workspace";

export default function ProfilePage() {
  const navigate = useNavigate();
  const apollo = useApolloClient();
  const { logout } = useAuth();

  const meQuery = useQuery(ME, { fetchPolicy: "cache-first" });
  const wsQuery = useQuery(WORKSPACES, { fetchPolicy: "cache-and-network" });

  const me = meQuery.data?.me || null;
  const workspaces = wsQuery.data?.workspaces || [];

  const handleLogout = async () => {
    try {
      await apollo.clearStore();
    } catch {
      // ignore
    }
    logout();
    navigate("/auth/login", { replace: true });
  };

  return (
    <div className="page">
      <div className="container">
        <div className="ui-panel pad-6">
          <div className="row-between">
            <div>
              <div className="h1">Profile</div>
              <div className="auth-sub">Your account & access</div>
            </div>

            <div className="row">
              <span className="ui-tag">GraphQL: /me</span>
              <button className="ui-btn-secondary" type="button" onClick={() => navigate("/app")}>
                back
              </button>
              <button className="ui-btn-danger" type="button" onClick={handleLogout}>
                logout
              </button>
            </div>
          </div>

          {meQuery.loading ? (
            <div className="spacer muted">loading...</div>
          ) : meQuery.error ? (
            <div className="spacer alert alert-danger">{meQuery.error.message}</div>
          ) : (
            <div className="stack spacer-lg">
              <div className="ui-card pad-5">
                <div className="row-between">
                  <div className="h2">Account</div>
                  <span className="ui-tag">id: {me?.id}</span>
                </div>
                <div className="spacer small muted2">name</div>
                <div className="truncate">{me?.name || "—"}</div>
                <div className="spacer small muted2">email</div>
                <div className="truncate">{me?.email || "—"}</div>
              </div>

              <div className="ui-card pad-5">
                <div className="row-between">
                  <div className="h2">Workspaces</div>
                  <span className="ui-tag">{wsQuery.loading ? "loading…" : workspaces.length}</span>
                </div>

                <div className="table-list">
                  {workspaces.map((w) => (
                    <button
                      key={w.id}
                      type="button"
                      className="list-item"
                      onClick={() => navigate(`/app/w/${w.id}`)}
                    >
                      <div className="truncate">{w.name || `workspace ${w.id}`}</div>
                      <span className="ui-tag">open</span>
                    </button>
                  ))}

                  {!wsQuery.loading && !workspaces.length ? (
                    <div className="table-row small muted2">no workspaces yet</div>
                  ) : null}
                </div>
              </div>

              <div className="ui-card pad-5">
                <div className="h2">Endpoints</div>
                <div className="spacer small muted2">
                  REST auth: {import.meta.env.VITE_API_BASE || "http://localhost:8080"}
                </div>
                <div className="small muted2">
                  GraphQL: {import.meta.env.VITE_GQL_URL || "http://localhost:8080/graphql"}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
