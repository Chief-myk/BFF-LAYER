// TextNode.tsx (if you haven't shared it yet)
import { Text, StyleSheet } from "react-native";

export default function TextNode({ node }: any) {
  const { data } = node;
  
  const styles = StyleSheet.create({
    text: {
      fontSize: data?.fontSize || data?.size || 14,
      fontWeight: data?.fontWeight || data?.weight || (data?.bold ? "bold" : "normal"),
      color: data?.color || data?.textColor || "#000000",
      textAlign: data?.textAlign || "left",
      marginTop: data?.marginTop,
      marginBottom: data?.marginBottom,
      marginLeft: data?.marginLeft,
      marginRight: data?.marginRight,
      letterSpacing: data?.letterSpacing,
      lineHeight: data?.lineHeight,
      backgroundColor: data?.backgroundColor,
      paddingHorizontal: data?.paddingHorizontal,
      paddingVertical: data?.paddingVertical,
      borderRadius: data?.borderRadius,
      alignSelf: data?.alignSelf,
    }
  });

  return <Text style={styles.text}>{data?.text || data?.value || ""}</Text>;
}