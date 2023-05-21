import { QueryClient } from '@tanstack/react-query';
import React from 'react';
import { Await, useLoaderData, useParams } from 'react-router-dom';
import { ensureNamespaceDetailsQueryData } from '../lib/api';
import {
  NamespaceDetailsHTTPResponse,
  NamespaceDetailsParams,
} from '../lib/types';

type DeferNamespaceDetailsLoader = {
  namespace: Promise<NamespaceDetailsHTTPResponse | undefined>;
};

export const loader =
  (queryClient: QueryClient) =>
  async ({ params }: { params: any }) =>
    ensureNamespaceDetailsQueryData(queryClient, params.name);

const NamespaceDetailsPage = () => {
  const { name } = useParams<NamespaceDetailsParams>();
  const { namespace } = useLoaderData() as DeferNamespaceDetailsLoader;

  return (
    <>
      <div className='container mx-auto m-4'>
        <h1 className='text-3xl'>NamespaceDetailsPage for {name}</h1>
      </div>
      <React.Suspense fallback={<div>Loading Details...</div>}>
        <Await
          resolve={namespace}
          errorElement={<div>Error loading namespace ...</div>}
        >
          {(namespace) => <div>{namespace.id}</div>}
        </Await>
      </React.Suspense>
    </>
  );
};
export default NamespaceDetailsPage;
