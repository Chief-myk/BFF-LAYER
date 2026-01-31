import { View } from "react-native";
import { MaterialCommunityIcons } from "@expo/vector-icons";

export default function IconNode({ node }: any) {
  const { data } = node; // lowercase

  return (
    <View
      style={{
        width: data?.containerSize, // lowercase
        height: data?.containerSize, // lowercase
        borderRadius: data?.borderRadius, // lowercase
        backgroundColor: data?.backgroundColor, // lowercase
        justifyContent: "center",
        alignItems: "center",
        marginRight: data?.marginRight, // lowercase
      }}
    >
      <MaterialCommunityIcons
        name={data?.name} // lowercase
        size={data?.size} // lowercase
        color={data?.color} // lowercase
      />
    </View>
  );
}