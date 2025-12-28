import { gql } from "@apollo/client";

export const BLOCKS = gql`
  query Blocks($pageId: String!, $parentblockId: String) {
    blocks(pageId: $pageId, parentblockId: $parentblockId) {
      id
      pageId
      parentblockId
      Type
      data
    }
  }
`;
