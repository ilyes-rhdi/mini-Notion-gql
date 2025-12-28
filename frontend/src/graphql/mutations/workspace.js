import { gql } from "@apollo/client";

export const CREATE_WORKSPACE = gql`
  mutation CreateWorkspace($name: String!) {
    createWorkspace(name: $name) {
      id
      name
      OwnerID
    }
  }
`;

export const ADD_WORKSPACE_MEMBER = gql`
  mutation AddWorkspaceMember($workspaceId: String!, $userId: String!, $role: WorkspaceRole) {
    addWorkspaceMember(workspaceId: $workspaceId, userId: $userId, role: $role) {
      id
      workspaceId
      userId
      role
      user { id name email }
    }
  }
`;

export const UPDATE_WORKSPACE_MEMBER_ROLE = gql`
  mutation UpdateWorkspaceMemberRole($workspaceId: String!, $userId: String!, $role: WorkspaceRole!) {
    updateWorkspaceMemberRole(workspaceId: $workspaceId, userId: $userId, role: $role) {
      id
      workspaceId
      userId
      role
      user { id name email }
    }
  }
`;

export const REMOVE_WORKSPACE_MEMBER = gql`
  mutation RemoveWorkspaceMember($workspaceId: String!, $userId: String!) {
    removeWorkspaceMember(workspaceId: $workspaceId, userId: $userId)
  }
`;

export const TRANSFER_WORKSPACE_OWNERSHIP = gql`
  mutation TransferWorkspaceOwnership($workspaceId: String!, $newOwnerUserId: String!) {
    transferWorkspaceOwnership(workspaceId: $workspaceId, newOwnerUserId: $newOwnerUserId) {
      id
      name
      OwnerID
    }
  }
`;
