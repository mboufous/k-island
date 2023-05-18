import NamespaceList from "./components/namespace/NamespaceList";
import { NamespaceItem } from "./lib/types";

const App = () => {
  const namespaces: NamespaceItem[] = [];
  return (
    <>
      <NamespaceList
        emptyHeading={"No Namespace was found"}
        namespaces={namespaces}
      />
    </>
  );
};
export default App;
