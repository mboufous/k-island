import { NamespaceItem } from "../../lib/types";

interface NamespaceProps {
  namespace: NamespaceItem;
}

const Namespace = ({ namespace }: NamespaceProps) => {
  return <div>{namespace.name}</div>;
};
export default Namespace;
