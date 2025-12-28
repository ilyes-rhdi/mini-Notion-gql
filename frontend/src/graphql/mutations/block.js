import { gql } from "@apollo/client";

export const CREATE_BLOCK = gql`
  mutation CreateBlock(
    $pageId: String!
    $parentblockId: String
    $type: BlockType!
    $order: Int
    $data: JSON
  ) {
    createBlock(
      pageId: $pageId
      parentblockId: $parentblockId
      type: $type
      order: $order
      data: $data
    ) {
      id
      pageId
      parentblockId
      Type
      data
    }
  }
`;

export const UPDATE_BLOCK = gql`
  mutation UpdateBlock($id: String!, $type: BlockType!, $data: JSON) {
    updateBlock(id: $id, type: $type, data: $data) {
      id
      Type
      data
    }
  }
`;

export const DELETE_BLOCK_TREE = gql`
  mutation DeleteBlockTree($id: String!) {
    deleteBlockTree(id: $id)
  }
`;
