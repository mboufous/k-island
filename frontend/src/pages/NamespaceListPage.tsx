import { QueryClient } from '@tanstack/react-query';
import React from 'react';
import { Await, useLoaderData } from 'react-router-dom';
import NamespaceCard from '../components/namespace/NamespaceCard';
import { ensureNamespacesQueryData } from '../lib/api';
import { NamespaceItem, NamespacesHTTPResponse } from '../lib/types';

// Router Loader
type DeferNamespacesLoader = {
  namespaces: Promise<NamespacesHTTPResponse | undefined>;
};
export const loader = (queryClient: QueryClient) => async () =>
  ensureNamespacesQueryData(queryClient);

const NamespaceListPage = () => {
  const { namespaces } = useLoaderData() as DeferNamespacesLoader;

  return (
    <>
      <div className='container mx-auto m-4'>
        <h1 className='text-3xl'>Namespaces</h1>
      </div>
      <React.Suspense fallback={<div>Loading ...</div>}>
        <Await
          resolve={namespaces}
          errorElement={<div>Error loding namespaces...</div>}
        >
          {(namespaces) => (
            <section className='p-10 box-border'>
              <div className='grid grid-cols-1 sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 w-100'>
                {namespaces?.length ? (
                  namespaces?.map((namespace: NamespaceItem) => (
                    <NamespaceCard key={namespace.name} namespace={namespace} />
                  ))
                ) : (
                  <span className='text-gray-500 italic'>No namespaces!</span>
                )}
              </div>
            </section>
          )}
        </Await>
      </React.Suspense>
    </>
  );
};
export default NamespaceListPage;
