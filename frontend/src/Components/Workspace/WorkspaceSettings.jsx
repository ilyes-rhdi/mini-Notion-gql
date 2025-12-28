import { useEffect, useMemo, useState } from "react";

export default function WorkspaceSettings({
  onClose,
  workspace,
  me,
  onAddMember,
  onUpdateRole,
  onRemoveMember,
  onTransferOwnership,
}) {
  const [userId, setUserId] = useState("");
  const [role, setRole] = useState("MEMBER");

  useEffect(() => {
    const onKeyDown = (e) => {
      if (e.key === "Escape") onClose?.();
    };
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, [onClose]);

  const { myRole, canManage, canTransfer } = useMemo(() => {
    const ownerId = workspace?.OwnerID || "";
    const meId = me?.id || "";
    const members = workspace?.Members || [];

    let r = "";
    if (meId && ownerId && String(meId) === String(ownerId)) r = "OWNER";
    if (!r) {
      const m = members.find((x) => String(x.userId) === String(meId));
      r = m?.role || "";
    }

    return {
      myRole: r || "MEMBER",
      canManage: r === "OWNER" || r === "ADMIN",
      canTransfer: r === "OWNER",
    };
  }, [workspace, me]);

  const members = workspace?.Members || [];
  const ownerId = workspace?.OwnerID || "";
  const meId = me?.id || "";

  const submitAdd = async () => {
    if (!workspace?.id || !userId.trim()) return;
    await onAddMember?.(workspace.id, userId.trim(), role);
    setUserId("");
    setRole("MEMBER");
  };

  const submitTransfer = async (newOwnerUserId) => {
    if (!workspace?.id) return;
    const ok = confirm("Transfer ownership? This cannot be undone.");
    if (!ok) return;
    await onTransferOwnership?.(workspace.id, newOwnerUserId);
  };

  return (
    <div className="modal-backdrop" onMouseDown={onClose}>
      <div
        className="ui-panel modal pad-6"
        onMouseDown={(e) => e.stopPropagation()}
      >
        <div className="row-between">
          <div>
            <div className="h1">Workspace settings</div>
            <div className="auth-sub">
              {workspace?.name || "workspace"} · your role: {String(myRole).toLowerCase()}
            </div>
          </div>

          <button className="ui-btn-ghost" type="button" onClick={onClose}>
            close
          </button>
        </div>

        <div className="stack spacer-lg">
          <section className="ui-card pad-5">
            <div className="row-between">
              <div className="h2">Members</div>
              <div className="ui-tag">{members.length}</div>
            </div>

            <div className="table-list">
              {members.map((m) => {
                const isOwner = String(m.userId) === String(ownerId);
                const isMe = String(m.userId) === String(meId);
                const canEdit = canManage && !isOwner;

                return (
                  <div key={m.id} className="ui-card pad-4">
                    <div className="min-w-0">
                      <div className="row">
                        <div className="truncate">
                          {m.user?.name || m.user?.email || m.userId}
                        </div>
                        {isMe ? <span className="ui-tag">you</span> : null}
                        {isOwner ? <span className="ui-tag">owner</span> : null}
                        {!isOwner ? <span className="ui-tag">{String(m.role).toLowerCase()}</span> : null}
                      </div>
                      <div className="truncate small muted2">
                        {m.user?.email || ""}
                      </div>
                      <div className="truncate small muted2">user id: {m.userId}</div>
                    </div>

                    <div className="row">
                      <select
                        className="ui-select ui-select-sm"
                        disabled={!canEdit}
                        value={m.role}
                        onChange={(e) => onUpdateRole?.(workspace.id, m.userId, e.target.value)}
                      >
                        <option value="ADMIN">ADMIN</option>
                        <option value="MEMBER">MEMBER</option>
                      </select>

                      {canEdit ? (
                        <button
                          className="ui-btn-danger"
                          type="button"
                          onClick={() => onRemoveMember?.(workspace.id, m.userId)}
                        >
                          remove
                        </button>
                      ) : null}

                      {canTransfer && !isOwner ? (
                        <button
                          className="ui-btn-secondary"
                          type="button"
                          onClick={() => submitTransfer(m.userId)}
                        >
                          make owner
                        </button>
                      ) : null}
                    </div>
                  </div>
                );
              })}
            </div>

            {!canManage ? (
              <div className="small muted2 spacer">only owner/admin can manage members</div>
            ) : null}
          </section>

          <section className="ui-card pad-5">
            <div className="h2">Add member</div>
            <div className="small muted2">
              Tip: ask them for their user id from Profile.
            </div>

            <div className="grid-2 spacer">
              <input
                className="ui-input"
                placeholder="user id"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
                disabled={!canManage}
              />

              <select
                className="ui-select"
                value={role}
                onChange={(e) => setRole(e.target.value)}
                disabled={!canManage}
              >
                <option value="MEMBER">MEMBER</option>
                <option value="ADMIN">ADMIN</option>
              </select>
            </div>

            <div className="row spacer">
              <button
                className="ui-btn-primary"
                type="button"
                onClick={submitAdd}
                disabled={!canManage || !userId.trim()}
              >
                add member
              </button>

              <button className="ui-btn-ghost" type="button" onClick={() => { setUserId(""); setRole("MEMBER"); }}>
                reset
              </button>
            </div>
          </section>
        </div>
      </div>
    </div>
  );
}
