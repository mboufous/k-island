import { Link } from 'react-router-dom';
import { NamespaceItem } from '../../lib/types';

interface NamespaceProps {
  namespace: NamespaceItem;
}

const NamespaceCard = ({ namespace }: NamespaceProps) => {
  return (
    <Link
      to={`${namespace.name}`}
      className='w-full m-auto border-2 border-white cursor-pointer rounded-lg drop-shadow-xl bg-white hover:bg-gray-100'
    >
      <div className=''>
        <div className='p-3'>
          <p className='text-base font-bold'>{namespace.name}</p>
          <div className='mt-4 overflow-hidden break-all text-xs text-gray-600'>
            <p className='whitespace-nowrap text-ellipsis overflow-hidden text-gray-600 font-semibold'>
              {namespace.spec.branch}
            </p>
            <p className='opacity-50'>{namespace.spec.commit}</p>
            <p className='opacity-50'>{namespace.spec.image_tag}</p>
          </div>
        </div>

        <div className='flex flex-row justify-end'>
          {namespace.blocked && (
            <span className='badge bg-red-200 text-red-600'>
              <svg
                xmlns='http://www.w3.org/2000/svg'
                fill='none'
                viewBox='0 0 24 24'
                strokeWidth='1.5'
                stroke='currentColor'
                className='-ms-1 me-1.5 h-4 w-4'
              >
                <path
                  strokeLinecap='round'
                  strokeLinejoin='round'
                  d='M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z'
                />
              </svg>

              <span className='whitespace-nowrap text-sm'>Blocked</span>
            </span>
          )}
          <span className='badge bg-emerald-100 text-emerald-700'>
            <svg
              xmlns='http://www.w3.org/2000/svg'
              fill='none'
              viewBox='0 0 24 24'
              strokeWidth='1.5'
              stroke='currentColor'
              className='-ms-1 me-1.5 h-4 w-4'
            >
              <path
                strokeLinecap='round'
                strokeLinejoin='round'
                d='M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z'
              />
            </svg>

            <span className='whitespace-nowrap text-sm'>
              {namespace.status}
            </span>
          </span>
        </div>
      </div>
    </Link>
  );
};
export default NamespaceCard;
