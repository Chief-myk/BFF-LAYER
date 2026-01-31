import { View, Pressable, Text } from "react-native"; // Add Text import
import { useRouter } from "expo-router";

export default function PressableCardNode({ node, renderNode, contextData = {}, onAction }: any) {
  const router = useRouter();
  const { data, children } = node;

  const onPress = () => {
    if (!data?.onPress) return;

    if (data.onPress.navigate) {
      router.push(data.onPress.navigate);
    }
  };

  // Helper function to properly render children
  const renderChildren = () => {
    if (!children) return null;
    
    return children.map((child: any, index: number) => {
      // If child is a string, wrap it in Text
      if (typeof child === 'string' || typeof child === 'number') {
        return <Text key={index}>{child}</Text>;
      }
      // Otherwise use the renderNode function
      return renderNode ? renderNode(child, index) : null;
    });
  };

  return (
    <Pressable
      onPress={onPress}
      style={{
        backgroundColor: data?.backgroundColor,
        padding: data?.padding,
        borderRadius: data?.borderRadius,
        borderWidth: data?.borderWidth,
        borderColor: data?.borderColor,
        marginTop: data?.marginTop,
        shadowColor: "#000",
        shadowOpacity: data?.shadow ? 0.15 : 0,
        shadowRadius: 8,
        elevation: data?.shadow ? 4 : 0,
        flexDirection: "row",
        alignItems: "center",
      }}
    >
      {renderChildren()}
    </Pressable>
  );
}
