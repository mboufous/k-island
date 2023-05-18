import { NamespaceItem } from "../../lib/types";
import Namespace from "./Namespace";

interface NamespaceListProps {
  namespaces: NamespaceItem[];
  emptyHeading: string;
}

const NamespaceList = ({ namespaces, emptyHeading }: NamespaceListProps) => {
  const count = namespaces.length;
  let heading = emptyHeading;
  if (count > 0) {
    const noun = count > 0 ? "namespace" : "Namespace";
    heading = count + " " + noun;
  }
  return (
    <section>
      <h2>{heading}</h2>
      {namespaces.map((namespace: NamespaceItem) => (
        <Namespace key={namespace.name} namespace={namespace} />
      ))}
    </section>
  );
};
export default NamespaceList;
