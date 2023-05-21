import { QueryClient } from '@tanstack/react-query';
import { LoaderFunctionArgs, useLoaderData } from 'react-router-dom';

type LoaderFunction<T> = (
  queryClient: QueryClient
) => ({ params }: LoaderFunctionArgs) => T;

export function getTypedInitialDataFromRouterLoader<T>(
  loader: LoaderFunction<T>
) {
  return useLoaderData() as Awaited<ReturnType<ReturnType<typeof loader>>>;
}
