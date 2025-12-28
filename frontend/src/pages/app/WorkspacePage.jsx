import { useEffect, useMemo, useState } from "react";
import { useMutation, useQuery } from "@apollo/client/react";

import Sidebar from "../../Components/Workspace/Sidebar";
import WorkspaceSettings from "../../Components/Workspace/WorkspaceSettings";
import Editor from "../../Components/Page/Editor";

import useWorkspace from "../../hooks/useWorkspace";

import { ME, WORKSPACE, WORKSPACES } from "../../graphql/queries/workspace";
import { PAGE } from "../../graphql/queries/page";
import { BLOCKS } from "../../graphql/queries/block";

import {
  ADD_WORKSPACE_MEMBER,
  CREATE_WORKSPACE,
  REMOVE_WORKSPACE_MEMBER,
  TRANSFER_WORKSPACE_OWNERSHIP,
  UPDATE_WORKSPACE_MEMBER_ROLE,
} from "../../graphql/mutations/workspace";
import { CREATE_PAGE, UPDATE_PAGE } from "../../graphql/mutations/page";
import { CREATE_BLOCK, DELETE_BLOCK_TREE, UPDATE_BLOCK } from "../../graphql/mutations/block";

export default function WorkspacePage() {
  const { workspaceId, pageId, openWorkspace, openPage } = useWorkspace();
  const [settingsOpen, setSettingsOpen] = useState(false);

  const wsQuery = useQuery(WORKSPACES, { fetchPolicy: "cache-and-network" });
  const workspaces = wsQuery.data?.workspaces || [];

  const meQuery = useQuery(ME, { fetchPolicy: "cache-first" });
  const me = meQuery.data?.me || null;

  const activeWorkspaceId = useMemo(() => {
    if (workspaceId) return workspaceId;
    return workspaces[0]?.id || "";
  }, [workspaceId, workspaces]);

  const workspaceQuery = useQuery(WORKSPACE, {
    variables: activeWorkspaceId ? { id: activeWorkspaceId } : undefined,
    skip: !activeWorkspaceId,
    fetchPolicy: "cache-and-network",
  });
  const workspace = workspaceQuery.data?.workspace || null;
  const pages = workspace?.Pages || [];

  const activePageId = useMemo(() => {
    if (pageId) return pageId;
    return pages[0]?.id || "";
  }, [pageId, pages]);

  const pageQuery = useQuery(PAGE, {
    variables: activePageId ? { id: activePageId } : undefined,
    skip: !activePageId,
    fetchPolicy: "cache-first",
  });

  const blocksQuery = useQuery(BLOCKS, {
    variables: activePageId ? { pageId: activePageId } : undefined,
    skip: !activePageId,
    fetchPolicy: "cache-and-network",
  });
  const blocks = blocksQuery.data?.blocks || [];

  const [createWorkspace] = useMutation(CREATE_WORKSPACE, {
    refetchQueries: [{ query: WORKSPACES }],
  });
  const [createPage] = useMutation(CREATE_PAGE, {
    refetchQueries: activeWorkspaceId ? [{ query: WORKSPACE, variables: { id: activeWorkspaceId } }] : [],
  });
  const [updatePage] = useMutation(UPDATE_PAGE);

  const [addWorkspaceMember] = useMutation(ADD_WORKSPACE_MEMBER, {
    refetchQueries: activeWorkspaceId ? [{ query: WORKSPACE, variables: { id: activeWorkspaceId } }] : [],
  });
  const [updateWorkspaceMemberRole] = useMutation(UPDATE_WORKSPACE_MEMBER_ROLE, {
    refetchQueries: activeWorkspaceId ? [{ query: WORKSPACE, variables: { id: activeWorkspaceId } }] : [],
  });
  const [removeWorkspaceMember] = useMutation(REMOVE_WORKSPACE_MEMBER, {
    refetchQueries: activeWorkspaceId ? [{ query: WORKSPACE, variables: { id: activeWorkspaceId } }] : [],
  });
  const [transferWorkspaceOwnership] = useMutation(TRANSFER_WORKSPACE_OWNERSHIP, {
    refetchQueries: activeWorkspaceId
      ? [{ query: WORKSPACE, variables: { id: activeWorkspaceId } }, { query: WORKSPACES }]
      : [{ query: WORKSPACES }],
  });

  const [createBlock] = useMutation(CREATE_BLOCK, {
    refetchQueries: activePageId ? [{ query: BLOCKS, variables: { pageId: activePageId } }] : [],
  });
  const [updateBlock] = useMutation(UPDATE_BLOCK);
  const [deleteBlockTree] = useMutation(DELETE_BLOCK_TREE, {
    refetchQueries: activePageId ? [{ query: BLOCKS, variables: { pageId: activePageId } }] : [],
  });

  useEffect(() => {
    if (!workspaceId && activeWorkspaceId) {
      openWorkspace(activeWorkspaceId);
      return;
    }
    if (workspaceId && !pageId && activeWorkspaceId && activePageId) {
      openPage(activeWorkspaceId, activePageId);
    }
  }, [workspaceId, pageId, activeWorkspaceId, activePageId, openWorkspace, openPage]);

  const handleCreateWorkspace = async () => {
    const name = prompt("workspace name") || "";
    if (!name.trim()) return;
    const res = await createWorkspace({ variables: { name: name.trim() } });
    const id = res.data?.createWorkspace?.id;
    if (id) openWorkspace(id);
  };

  const handleCreatePage = async () => {
    if (!activeWorkspaceId) return;
    const title = prompt("page title") || "";
    const res = await createPage({ variables: { workspaceId: activeWorkspaceId, title: title.trim() || null } });
    const id = res.data?.createPage?.id;
    if (id) openPage(activeWorkspaceId, id);
  };

  const handleAddMember = async (workspaceIdArg, userIdArg, roleArg) => {
    await addWorkspaceMember({ variables: { workspaceId: workspaceIdArg, userId: userIdArg, role: roleArg } });
  };

  const handleUpdateRole = async (workspaceIdArg, userIdArg, roleArg) => {
    await updateWorkspaceMemberRole({ variables: { workspaceId: workspaceIdArg, userId: userIdArg, role: roleArg } });
  };

  const handleRemoveMember = async (workspaceIdArg, userIdArg) => {
    const ok = confirm("Remove this member from the workspace?");
    if (!ok) return;
    await removeWorkspaceMember({ variables: { workspaceId: workspaceIdArg, userId: userIdArg } });
  };

  const handleTransferOwnership = async (workspaceIdArg, newOwnerUserId) => {
    await transferWorkspaceOwnership({ variables: { workspaceId: workspaceIdArg, newOwnerUserId } });
  };

  const handleRenamePage = async (id, title) => {
    if (!id) return;
    const next = (title || "").trim();
    if (!next) return;
    await updatePage({ variables: { id, title: next } });
  };

  const handleCreateParagraph = async () => {
    if (!activePageId) return;
    await createBlock({
      variables: {
        pageId: activePageId,
        parentblockId: null,
        type: "PARAGRAPH",
        data: { text: "" },
      },
    });
  };

  const handleUpdateBlock = async (id, type, data) => {
    if (!id) return;
    await updateBlock({ variables: { id, type: type || "PARAGRAPH", data } });
  };

  const handleDeleteBlock = async (id) => {
    if (!id) return;
    await deleteBlockTree({ variables: { id } });
  };

  const loading = wsQuery.loading || meQuery.loading || (activeWorkspaceId && workspaceQuery.loading);
  const error = wsQuery.error || meQuery.error || workspaceQuery.error || pageQuery.error || blocksQuery.error;

  return (
    <div className="app-shell">
      <Sidebar
        me={me}
        workspaces={workspaces}
        activeWorkspaceId={activeWorkspaceId}
        onSelectWorkspace={openWorkspace}
        onCreateWorkspace={handleCreateWorkspace}
        pages={pages}
        activePageId={activePageId}
        onSelectPage={(pid) => openPage(activeWorkspaceId, pid)}
        onCreatePage={handleCreatePage}
        onOpenWorkspaceSettings={() => setSettingsOpen(true)}
      />

      <main className="main">
        <div className="main-scroll">
          <div className="row-between">
            <div>
              <div className="h2">{workspace?.name || "Workspace"}</div>
              <div className="auth-sub">
                {activeWorkspaceId ? `workspace id: ${activeWorkspaceId}` : ""}
              </div>
            </div>
            <div className="row">
              {loading ? <div className="ui-tag">loading...</div> : null}
              <button
                className="ui-btn-secondary"
                type="button"
                onClick={() => setSettingsOpen(true)}
                disabled={!activeWorkspaceId}
              >
                workspace settings
              </button>
            </div>
          </div>

          {error ? (
            <div className="alert alert-danger">
              {error.message}
            </div>
          ) : null}

          {!activeWorkspaceId ? (
            <div className="ui-panel pad-6 spacer-lg">
              <div className="muted">Create your first workspace to start.</div>
              <button className="ui-btn-primary" type="button" onClick={handleCreateWorkspace}>
                create workspace
              </button>
            </div>
          ) : (
            <div className="spacer-lg">
              <Editor
                key={activePageId || "no-page"}
                page={pageQuery.data?.page || null}
                blocks={blocks}
                onRenamePage={handleRenamePage}
                onCreateParagraph={handleCreateParagraph}
                onUpdateBlock={handleUpdateBlock}
                onDeleteBlock={handleDeleteBlock}
              />
            </div>
          )}

            {settingsOpen && (
              <WorkspaceSettings
                key={activeWorkspaceId} // optionnel mais recommandé (reset si tu changes de workspace)
                onClose={() => setSettingsOpen(false)}
                workspace={workspace}
                me={me}
                onAddMember={handleAddMember}
                onUpdateRole={handleUpdateRole}
                onRemoveMember={handleRemoveMember}
                onTransferOwnership={handleTransferOwnership}
              />
            )}
        </div>
      </main>
    </div>
  );
}
