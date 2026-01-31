import { Pressable, Text } from "react-native";
import { useRouter } from "expo-router";

export default function TextButtonNode({ node, contextData = {}, onAction }: any) {
  const router = useRouter();
  const { data } = node; // lowercase

  const onPress = () => {
    if (data?.onPress?.type === "NAVIGATE_BACK") { // lowercase
      router.back();
    }
  };

  return (
    <Pressable onPress={onPress}>
      <Text
        style={{
          color: data?.color, // lowercase
          fontSize: data?.fontSize, // lowercase
          fontWeight: data?.fontWeight, // lowercase
          textAlign: "center",
        }}
      >
        {data?.text} 
      </Text>
    </Pressable>
  );
}