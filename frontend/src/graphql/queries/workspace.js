import { gql } from "@apollo/client";

export const ME = gql`
  query Me {
    me {
      id
      name
      email
    }
  }
`;

export const WORKSPACES = gql`
  query Workspaces {
    workspaces {
      id
      name
      OwnerID
    }
  }
`;

export const WORKSPACE = gql`
  query Workspace($id: String!) {
    workspace(id: $id) {
      id
      name
      OwnerID
      Members {
        id
        userId
        role
        user { id name email }
      }
      Pages {
        id
        title
        archived
      }
    }
  }
`;
