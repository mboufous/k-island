import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  Link,
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
  useParams,
} from 'react-router-dom';
import App from './App';
import './index.css';
import { NamespaceDetailsParams } from './lib/types';
import ErrorPage from './pages/ErrorPage';
import NamespaceDetailsPage, {
  loader as namespaceDetailsLoader,
} from './pages/NamespaceDetailsPage';
import NamespaceListPage, {
  loader as namespacesLoader,
} from './pages/NamespaceListPage';

// Query Client
export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 10,
    },
  },
});

// Router
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      path='/'
      element={<App />}
      errorElement={<ErrorPage />}
      handle={{
        crumb: () => <Link to='/'>Namespaces</Link>,
      }}
    >
      <Route
        index
        element={<NamespaceListPage />}
        loader={namespacesLoader(queryClient)}
      />
      <Route
        path=':name'
        element={<NamespaceDetailsPage />}
        loader={namespaceDetailsLoader(queryClient)}
        handle={{
          crumb: () => <span>{useParams<NamespaceDetailsParams>().name}</span>,
        }}
      />
    </Route>
  )
);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>
);
