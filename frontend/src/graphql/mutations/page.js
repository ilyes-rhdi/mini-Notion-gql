import { gql } from "@apollo/client";

export const CREATE_PAGE = gql`
  mutation CreatePage($workspaceId: String!, $title: String) {
    createPage(workspaceId: $workspaceId, title: $title) {
      id
      workspaceId
      title
      archived
    }
  }
`;

export const UPDATE_PAGE = gql`
  mutation UpdatePage($id: String!, $title: String!) {
    updatePage(id: $id, title: $title) {
      id
      workspaceId
      title
      archived
    }
  }
`;
