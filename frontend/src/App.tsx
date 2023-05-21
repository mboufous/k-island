import { Outlet, useNavigation } from 'react-router-dom';
import Breadcrumbs from './components/Breadcrumbs';

const App = () => {
  return (
    <main className='h-screen w-screen flex flex-col'>
      <div hidden={useNavigation().state !== 'loading'}>
        <div className='w-full flex items-center flex-col bg-white'>
          <div className='flex flex-col bg-white shadow-md  rounded-md items-center'>
            <div className='flex items-center p-4'>
              <div
                data-placeholder
                className='mr-2 h-10 w-10  rounded-full overflow-hidden relative bg-gray-200'
              ></div>
              <div className='flex flex-col justify-between items-center'>
                <div
                  data-placeholder
                  className='mb-2 h-5 w-40 overflow-hidden relative bg-gray-200'
                ></div>
              </div>
            </div>
            <div
              data-placeholder
              className='h-52 w-full overflow-hidden relative bg-gray-200'
            ></div>

            <div className='flex flex-col p-4'>
              <div className='flex'>
                <div
                  data-placeholder
                  className=' flex h-5 w-5 overflow-hidden relative bg-gray-200 mr-1'
                ></div>
                <div
                  data-placeholder
                  className='flex h-5 w-48 overflow-hidden relative bg-gray-200'
                ></div>
              </div>
              <div className='flex mt-1'>
                <div
                  data-placeholder
                  className='flex h-5 w-5 overflow-hidden relative bg-gray-200 mr-1'
                ></div>
                <div
                  data-placeholder
                  className='flex h-5 w-48 overflow-hidden relative bg-gray-200'
                ></div>
              </div>
            </div>
            <div className='w-full h-px  overflow-hidden relative bg-gray-200 m-4'></div>
            <div className='flex justify-between items-center p-4 w-full'>
              <div
                data-placeholder
                className='mr-2 h-10 w-16  overflow-hidden relative bg-gray-200'
              ></div>

              <div
                data-placeholder
                className='mb-2 h-5 w-20 overflow-hidden relative bg-gray-200'
              ></div>
            </div>
          </div>
        </div>
      </div>
      <Breadcrumbs />
      <Outlet />
    </main>
  );
};
export default App;
