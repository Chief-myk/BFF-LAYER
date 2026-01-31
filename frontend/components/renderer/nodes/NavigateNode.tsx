import { useEffect } from "react";
import { router } from "expo-router";

export default function NavigateNode({ node }: any) {
  useEffect(() => {
    const timer = setTimeout(() => {
      router.replace(node.data.to);
    }, node.data.after || 0);

    return () => clearTimeout(timer);
  }, []);

  return null;
}
