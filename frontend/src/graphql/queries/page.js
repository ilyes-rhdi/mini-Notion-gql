import { gql } from "@apollo/client";

export const PAGE = gql`
  query Page($id: String!) {
    page(id: $id) {
      id
      workspaceId
      title
      archived
    }
  }
`;

export const PAGES = gql`
  query Pages($workspaceId: String!) {
    pages(workspaceId: $workspaceId) {
      id
      title
      archived
      workspaceId
    }
  }
`;
