import { useEffect } from "react";
import { View, Text, Image } from "react-native";
import Animated, {
  Easing,
  useAnimatedStyle,
  useSharedValue,
  withRepeat,
  withSequence,
  withTiming,
} from "react-native-reanimated";
import { router } from "expo-router";

export default function SplashScreenRenderer({ ui }: { ui: any[] }) {
  const fade = useSharedValue(0);
  const scale = useSharedValue(0.8);

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

  const renderNode = (node: any, index: number) => {
    switch (node.type) {
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
