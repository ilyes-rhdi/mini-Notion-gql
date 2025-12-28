import { ApolloClient, InMemoryCache, HttpLink, from } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

const GQL_URL = import.meta.env.VITE_GQL_URL || "http://localhost:8080/graphql";

const httpLink = new HttpLink({
  uri: GQL_URL,
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem("token");
  return {
    headers: {
      ...headers,
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
  };
});

const client = new ApolloClient({
  link: from([authLink, httpLink]),
  cache: new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          blocks: {
            keyArgs: ["pageId", "parentblockId"],
          },
          pages: {
            keyArgs: ["workspaceId"],
          },
        },
      },
    },
  }),
});

export default client;
