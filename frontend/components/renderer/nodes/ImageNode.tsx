import { Image } from "react-native";

export default function ImageNode({ node }: any) {
  return (
    <Image
      source={{ uri: node.data.url }}
      style={{
        width: node.data.width,
        height: node.data.height,
        resizeMode: node.data.resizeMode,
      }}
    />
  );
}
