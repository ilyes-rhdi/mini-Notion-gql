import { useApolloClient } from "@apollo/client/react";
import { useNavigate } from "react-router-dom";

import useAuth from "../../hooks/useAuth";

export default function Sidebar({
  me,
  workspaces,
  activeWorkspaceId,
  onSelectWorkspace,
  onCreateWorkspace,
  pages,
  activePageId,
  onSelectPage,
  onCreatePage,
  onOpenWorkspaceSettings,
}) {
  const navigate = useNavigate();
  const apollo = useApolloClient();
  const { logout } = useAuth();

  const doLogout = async () => {
    logout();
    try {
      await apollo.clearStore();
    } catch {
      // ignore
    }
    navigate("/auth/login", { replace: true });
  };

  return (
    <aside className="sidebar">
      <div className="p-4">
        <div className="row-between">
          <div>
            <div className="ui-tag">mini-notion</div>
            <div className="small muted2">GraphQL workspace</div>
          </div>
          <button className="ui-btn-danger" type="button" onClick={doLogout}>
            logout
          </button>
        </div>

        <button
          type="button"
          className="ui-btn-secondary full btn-between"
          onClick={() => navigate("/app/profile")}
        >
          <span className="truncate">{me?.name || me?.email || "profile"}</span>
          <span className="muted">›</span>
        </button>
      </div>

      <div >
        <div className="ui-panel p-3">
          <div className="row-between">
            <div className="h2">workspaces</div>
            <div className="row">
              <button className="ui-btn-secondary" type="button" onClick={onCreateWorkspace}>
                + ws
              </button>
              <button
                className="ui-btn-secondary"
                type="button"
                onClick={onOpenWorkspaceSettings}
                disabled={!activeWorkspaceId}
                title={activeWorkspaceId ? "Manage members & roles" : "Select a workspace first"}
              >
                settings
              </button>
            </div>
          </div>

          <div className="list">
            {(workspaces || []).map((ws) => {
              const active = ws.id === activeWorkspaceId;
              return (
                <button
                  key={ws.id}
                  type="button"
                  onClick={() => onSelectWorkspace(ws.id)}
                  className={
                    "list-item full " + (active ? "active" : "")
                  }
                >
                  <span className="truncate">{ws.name || "untitled"}</span>
                  {active ? <span className="ui-tag">active</span> : null}
                </button>
              );
            })}
            {!workspaces?.length ? <div className="small muted2">no workspace yet</div> : null}
          </div>
        </div>
      </div>

      <div className="sidebar-footer">
        <div className="ui-panel p-3">
          <div className="row-between">
            <div className="h2">pages</div>
            <button
              className="ui-btn-secondary"
              type="button"
              onClick={onCreatePage}
              disabled={!activeWorkspaceId}
              title={!activeWorkspaceId ? "Select a workspace first" : ""}
            >
              + page
            </button>
          </div>

          <div className="list">
            {(pages || []).map((p) => {
              const active = p.id === activePageId;
              return (
                <button
                  key={p.id}
                  type="button"
                  onClick={() => onSelectPage(p.id)}
                  className={
                    "list-item full " + (active ? "active" : "")
                  }
                >
                  <span className="truncate">{p.title || "untitled"}</span>
                  {p.archived ? <span className="ui-tag">archived</span> : active ? <span className="ui-tag">open</span> : null}
                </button>
              );
            })}
            {activeWorkspaceId && !pages?.length ? (
              <div className="small muted2">no pages yet</div>
            ) : null}
            {!activeWorkspaceId ? <div className="small muted2">select a workspace</div> : null}
          </div>
        </div>
      </div>
    </aside>
  );
}
