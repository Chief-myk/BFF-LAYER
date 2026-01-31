import { View, TextInput, StyleSheet } from "react-native";
import { useRef, useState } from "react";
import { useFormStore } from "../store/formStore";

export default function OtpInputNode({ node }: any) {
  const { setValue } = useFormStore();
  const { 
    length, 
    boxSpacing = 12, // Default to 12 if not provided
    boxWidth = 50,
    boxHeight = 60,
    boxBorderRadius = 12,
    boxBorderColor = "#E5E5E5",
    boxBackgroundColor = "#F8F8F8"
  } = node.data;
  
  const [otp, setOtp] = useState(Array(length).fill(""));
  const refs = useRef<TextInput[]>([]);

  const onChange = (v: string, i: number) => {
    const next = [...otp];
    next[i] = v.slice(-1);
    setOtp(next);
    setValue(node.data.id, next.join(""));
    if (v && i < length - 1) refs.current[i + 1]?.focus();
  };

  // Create styles with dynamic spacing
  const styles = StyleSheet.create({
    container: {
      flexDirection: "row",
      marginBottom: 30,
    },
    input: {
      width: boxWidth,
      height: boxHeight,
      borderWidth: 2,
      borderRadius: boxBorderRadius,
      textAlign: "center",
      fontSize: 20,
      borderColor: boxBorderColor,
      backgroundColor: boxBackgroundColor,
      marginRight: boxSpacing, // Apply spacing to the right
    },
    lastInput: {
      width: boxWidth,
      height: boxHeight,
      borderWidth: 2,
      borderRadius: boxBorderRadius,
      textAlign: "center",
      fontSize: 20,
      borderColor: boxBorderColor,
      backgroundColor: boxBackgroundColor,
      marginRight: 0, // No spacing for last item
    },
    activeInput: {
      borderColor: "#10B981",
    }
  });

  return (
    <View style={styles.container}>
      {otp.map((d, i) => (
        <TextInput
          key={i}
          ref={(r) => (refs.current[i] = r!)}
          value={d}
          onChangeText={(v) => onChange(v, i)}
          keyboardType="number-pad"
          maxLength={1}
          style={[
            i === length - 1 ? styles.lastInput : styles.input, // Last item gets no right margin
            d ? styles.activeInput : {} // Active styling when filled
          ]}
        />
      ))}
    </View>
  );
}