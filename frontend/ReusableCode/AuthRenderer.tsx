import { useEffect, useState } from "react";
import {
  View,
  Text,
  Image,
  TextInput,
  TouchableOpacity,
  ScrollView,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import Animated, {
  Easing,
  useAnimatedStyle,
  useSharedValue,
  withRepeat,
  withSequence,
  withTiming,
} from "react-native-reanimated";
import { router } from "expo-router";

export default function AuthRenderer({ ui }: { ui: any[] }) {
  const fade = useSharedValue(0);
  const scale = useSharedValue(0.8);
  const [form, setForm] = useState<Record<string, string>>({});

  useEffect(() => {
    fade.value = withTiming(1, { duration: 600 });
    scale.value = withRepeat(
      withSequence(
        withTiming(1, { duration: 700, easing: Easing.inOut(Easing.ease) }),
        withTiming(0.92, { duration: 700, easing: Easing.inOut(Easing.ease) })
      ),
      -1,
      true
    );
  }, []);

  const animatedStyle = useAnimatedStyle(() => ({
    opacity: fade.value,
    transform: [{ scale: scale.value }],
  }));

  const renderNode = (node: any, index: number): any => {
    switch (node.type) {
      case "SAFE_AREA":
        return (
          <SafeAreaView key={index} style={{ flex: 1 }}>
            {node.children?.map(renderNode)}
          </SafeAreaView>
        );

      case "SCROLL":
        return (
          <ScrollView
            key={index}
            contentContainerStyle={node.data}
            showsVerticalScrollIndicator={false}
          >
            {node.children?.map(renderNode)}
          </ScrollView>
        );

      case "VIEW":
        return (
          <Animated.View
            key={index}
            style={node.data.animation ? animatedStyle : node.data}
          >
            {node.children?.map(renderNode)}
          </Animated.View>
        );

      case "IMAGE":
        return (
          <Image
            key={index}
            source={{ uri: node.data.url }}
            style={{
              width: node.data.width,
              height: node.data.height,
              resizeMode: node.data.resizeMode,
            }}
          />
        );

      case "TEXT":
        return (
          <Text key={index} style={node.data}>
            {node.data.text}
          </Text>
        );

      case "INPUT":
        return (
          <TextInput
            key={index}
            style={node.data.style}
            placeholder={node.data.placeholder}
            keyboardType={node.data.keyboardType}
            maxLength={node.data.maxLength}
            value={form[node.data.id] || ""}
            onChangeText={(v) =>
              setForm((prev) => ({ ...prev, [node.data.id]: v }))
            }
          />
        );

      case "BUTTON":
        return (
          <TouchableOpacity
            key={index}
            style={node.data.style}
            onPress={() => {
              if (node.data.action?.type === "API_CALL") {
                console.log("API CALL →", node.data.action.url, form);
              }
            }}
          >
            <Text style={{ color: "#FFFFFF", fontWeight: "600" }}>
              {node.data.text}
            </Text>
          </TouchableOpacity>
        );

      case "NAVIGATE":
        setTimeout(() => {
          router.replace(node.data.to);
        }, node.data.after);
        return null;

      default:
        return null;
    }
  };

  return <View style={{ flex: 1 }}>{ui.map(renderNode)}</View>;
}
