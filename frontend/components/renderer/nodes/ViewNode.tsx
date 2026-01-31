import { useEffect } from "react";
import Animated, {
  useAnimatedStyle,
  useSharedValue,
  withTiming,
} from "react-native-reanimated";

export default function ViewNode({ node, renderNode }: any) {
  const fade = useSharedValue(0);

  useEffect(() => {
    fade.value = withTiming(1, { duration: 400 });
  }, []);

  const animatedStyle = useAnimatedStyle(() => ({
    opacity: fade.value,
  }));

  return (
    <Animated.View
      style={[node.data, node.data?.animation && animatedStyle]}
    >
      {node.children?.map(renderNode)}
    </Animated.View>
  );
}
