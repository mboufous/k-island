import { useMatches } from 'react-router-dom';

type MatchWithCrumb = ReturnType<typeof useMatches>[number] & {
  handle:
    | {
        crumb?: (data: any) => any;
      }
    | any;
};

const Breadcrumbs = () => {
  const matches = useMatches();

  const crumbs = matches
    .filter((match: MatchWithCrumb) => match.handle && match.handle.crumb)
    .map((match: MatchWithCrumb) => match.handle.crumb(match.data));

  return (
    <ol>
      {crumbs.map((crumb, index) => (
        <li key={index}>{crumb}</li>
      ))}
    </ol>
  );
};

export default Breadcrumbs;
