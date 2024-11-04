import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import "/node_modules/primeflex/primeflex.css";
import App from "./App.tsx";

const client = new ApolloClient({
  uri: "http://localhost:8080/api",
  cache: new InMemoryCache(),
});

createRoot(document.getElementById("root")!).render(
  <ApolloProvider client={client}>
    <StrictMode>
      <App />
    </StrictMode>
    ,
  </ApolloProvider>
);
