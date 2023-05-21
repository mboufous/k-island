import { isRouteErrorResponse, useRouteError } from 'react-router-dom';

const ErrorPage = () => {
  const error = useRouteError();

  if (isRouteErrorResponse(error)) {
    // if (error.status === 401) {

    // }

    return (
      <div className='error-page'>
        <h1>
          {error.status} {error.statusText}
        </h1>
        {error.data?.message && (
          <p>
            <i>{error.data.message}</i>
          </p>
        )}
      </div>
    );
  } else if (error instanceof Error) {
    return (
      <div className='error-page'>
        <h1>Oops! Unexpected Error</h1>
        <p>Something went wrong.</p>
        <p>
          <i>{error.message}</i>
        </p>
      </div>
    );
  } else {
    return (
      <div className='error-page'>
        <h1>Oops! Unexpected Error</h1>
        <p>Something went wrong.</p>
      </div>
    );
  }
};

export default ErrorPage;
