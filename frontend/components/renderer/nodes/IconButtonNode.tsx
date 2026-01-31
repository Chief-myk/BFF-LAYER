import React from "react";
import { TouchableOpacity, View, StyleProp, ViewStyle } from "react-native";
import { MaterialCommunityIcons, Ionicons, FontAwesome5 } from "@expo/vector-icons";
import { router } from "expo-router";

type IconButtonNodeProps = {
  node: any;
  renderNode?: (node: any, index: number) => React.ReactNode;
};

// Icon name mapping - convert backend names to MaterialCommunityIcons names
const ICON_NAME_MAP: Record<string, string> = {
  "arrow-back": "arrow-left", // Map "arrow-back" to "arrow-left"
  "arrow-left": "arrow-left",
  "arrow-right": "arrow-right",
  "close": "close",
  "menu": "menu",
  // Add more mappings as needed
};

export default function IconButtonNode({ node, contextData = {}, onAction }: IconButtonNodeProps) {
  const { data } = node;
  
  // Extract data with defaults
  const backendIconName = data?.icon || "help";
  // Use mapped name or fallback to original
  const iconName = ICON_NAME_MAP[backendIconName] || backendIconName;
  const size = data?.size || 24;
  const color = data?.color || "#000000";
  const style = data?.style || {};
  const onPress = data?.onPress;

  const handlePress = () => {
    if (!onPress) return;

    console.log("IconButton action:", onPress.type);

    switch (onPress.type) {
      case "NAVIGATE_BACK":
        router.back();
        break;
      case "NAVIGATE":
        if (onPress.navigate || onPress.to) {
          router.push(onPress.navigate || onPress.to);
        }
        break;
      default:
        console.log("IconButton action type:", onPress.type);
    }
  };

  return (
    <TouchableOpacity style={style as StyleProp<ViewStyle>} onPress={handlePress}>
      <MaterialCommunityIcons name={iconName} size={size} color={color} />
    </TouchableOpacity>
  );
}