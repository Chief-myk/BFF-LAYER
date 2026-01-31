import { Text, TouchableOpacity, View } from "react-native";
import { useEffect, useState } from "react";

export default function ResendOtpNode({ node, renderNode }: any) {
  const [timer, setTimer] = useState(node.data.timer);
  const { marginTop = 20 } = node.data;

  useEffect(() => {
    if (timer > 0) {
      const t = setTimeout(() => setTimer(timer - 1), 1000);
      return () => clearTimeout(t);
    }
  }, [timer]);

  return (
    <View style={{ marginTop }}>
      <TouchableOpacity disabled={timer > 0}>
        <Text style={{ color: timer > 0 ? "#999" : "#FF0000", fontWeight: "600" }}>
          {timer > 0 ? `Resend in ${timer}s` : "Resend OTP"}
        </Text>
      </TouchableOpacity>
      {/* Render children if they exist */}
      {node.children && node.children.map((child: any, index: number) => 
        renderNode ? renderNode(child, index) : null
      )}
    </View>
  );
}