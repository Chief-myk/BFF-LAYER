import { SafeAreaView } from "react-native-safe-area-context";

export default function SafeAreaNode({ node, renderNode }: any) {
  return (
    <SafeAreaView style={{ flex: 1 }}>
      {node.children?.map(renderNode)}
    </SafeAreaView>
  );
}
