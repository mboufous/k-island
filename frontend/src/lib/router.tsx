import { Navigate, createBrowserRouter } from 'react-router-dom';
import App from '../App';
import ErrorPage from '../pages/ErrorPage';
import NamespaceDetailsPage from '../pages/NamespaceDetailsPage';
import NamespaceListPage, {
  namespacesLoader,
} from '../pages/NamespaceListPage';
import { queryClient } from './api';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <NamespaceListPage />,
        loader: namespacesLoader(queryClient),
      },
      {
        path: 'namespaces',
        element: <Navigate to='/' />,
      },
      {
        path: 'namespaces/:name',
        element: <NamespaceDetailsPage />,
      },
    ],
  },
]);

export default router;
