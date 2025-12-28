import { useMemo } from "react";
import { useNavigate, useParams } from "react-router-dom";

export default function useWorkspace() {
  const navigate = useNavigate();
  const params = useParams();

  const workspaceId = params.workspaceId || "";
  const pageId = params.pageId || "";

  const actions = useMemo(
    () => ({
      openWorkspace: (id) => navigate(`/app/w/${id}`),
      openPage: (ws, pid) => navigate(`/app/w/${ws}/p/${pid}`),
    }),
    [navigate]
  );

  return { workspaceId, pageId, ...actions };
}
