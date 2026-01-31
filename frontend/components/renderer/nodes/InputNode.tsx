import { TextInput, StyleSheet } from "react-native";
import { useFormStore } from "../store/formStore";
import { useEffect } from "react";

export default function InputNode({ node }: any) {
  const { values, setValue } = useFormStore();
  
  const inputId = node.data.id || "input";
  const value = values[inputId] || node.data.value || "";
  
  // Auto-fill phone number if this is a phone field
  useEffect(() => {
    const isPhoneField = inputId === "phone" || 
                        inputId === "mobileNumber" ||
                        inputId === "Phone";
    
    if (isPhoneField && !value && values["phone"]) {
      // If this is a phone field and empty, try to get from global store
      setValue(inputId, values["phone"]);
    }
  }, [inputId, value]);

  return (
    <TextInput
      style={[styles.input, node.data.style]}
      placeholder={node.data.placeholder}
      placeholderTextColor="#999"
      keyboardType={node.data.keyboardType || "default"}
      maxLength={node.data.maxLength}
      value={value}
      onChangeText={(v) => setValue(inputId, v)}
      secureTextEntry={node.data.secureTextEntry}
      autoCapitalize={node.data.autoCapitalize || "none"}
      autoCorrect={node.data.autoCorrect || false}
    />
  );
}

const styles = StyleSheet.create({
  input: {
    fontSize: 16,
    color: "#000",
    padding: 0,
    margin: 0,
    flex: 1,
  },
});