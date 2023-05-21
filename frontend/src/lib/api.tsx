import { QueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { defer } from 'react-router-dom';
import { NamespaceDetailsHTTPResponse, NamespacesHTTPResponse } from './types';

const apiClient = axios.create({
  baseURL: 'http://localhost:3000/',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Namespaces List
const namespacesQuery = {
  queryKey: ['namespaces'],
  queryFn: async () => fetchNamespaces(),
};

const fetchNamespaces = async () => {
  const response = await apiClient.get<NamespacesHTTPResponse>('/namespaces');
  return response.data;
};

export async function ensureNamespacesQueryData(queryClient: QueryClient) {
  const namespacesPromise =
    queryClient.ensureQueryData<NamespacesHTTPResponse>(namespacesQuery);
  return defer({
    namespaces: namespacesPromise,
  });
}

// Namespace details
const namespaceDetailsQuery = (name: string) => ({
  queryKey: ['namespace', 'details', name],
  queryFn: async () => fetchNamespaceDetails(name),
});

const fetchNamespaceDetails = async (name: string) => {
  const response = await apiClient.get<NamespaceDetailsHTTPResponse>(
    `/namespace/${name}`
  );
  return response.data;
};

export async function ensureNamespaceDetailsQueryData(
  queryClient: QueryClient,
  name: string
) {
  const namespacePromise =
    queryClient.ensureQueryData<NamespaceDetailsHTTPResponse>(
      namespaceDetailsQuery(name)
    );
  return defer({
    namespace: namespacePromise,
  });
}
