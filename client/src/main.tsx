import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import App from './App.tsx'
import "/node_modules/primeflex/primeflex.css";
import "./assets/theme.css";
import "./assets/index.css";

const client = new ApolloClient({
  uri: 'http://localhost:8080/api',
  cache: new InMemoryCache(),
});

createRoot(document.getElementById('root')!).render(
  <ApolloProvider client={client}>
    <StrictMode>
      <App />
    </StrictMode>,
  </ApolloProvider>,
)
