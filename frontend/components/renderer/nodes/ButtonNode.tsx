import { TouchableOpacity, Text, Alert } from "react-native";
import { useFormStore } from "../store/formStore";
import { router } from "expo-router";

export default function ButtonNode({ node, contextData = {}, onAction }: any) {
  const { values } = useFormStore();

  const onPress = async () => {
    const action = node.data.action;
    
    // If no action, just return
    if (!action) return;

    console.log("Button Action detected:", action.type, action.value);

    // Handle NAVIGATE action
    if (action.type === "NAVIGATE") {
      console.log("Navigating to:", action.to || action.navigate);
      const route = action.to || action.navigate;
      if (route) {
        router.push(route);
      }
      return;
    }

    // Handle ACTION type (this is what your new backend sends)
    if (action.type === "ACTION") {
      console.log("Dynamic action triggered:", action.value);
      
      // Call the onAction handler if provided
      if (onAction) {
        onAction(action.value, {
          ...action,
          contextData: contextData,
          nodeData: node.data
        });
      } else {
        // If no onAction handler, try to handle it directly
        handleDirectAction(action);
      }
      return;
    }

    // Handle non-API actions
    if (action.type !== "API_CALL") {
      console.log("Non-API action:", action);
      return;
    }

    // ... REST OF YOUR EXISTING CODE STAYS THE SAME ...
    try {
      const apiUrl = action.url.startsWith("http")
        ? action.url
        : "https://backend.truckhai.com" + action.url;

      // SPECIAL HANDLING FOR OTP REQUEST
      // Backend expects {"Phone": "9990538802"} with capital P
      let requestBody: any = values;
      
      // If this is an OTP request, format the data properly
      if (action.url.includes("request-otp") || action.url.includes("/otp")) {
        // Try to get phone number from different possible field names
        const phoneNumber = values.phone || values.mobileNumber || values.Phone;
        if (phoneNumber) {
          requestBody = { Phone: phoneNumber }; // Capital P as backend expects
          console.log("Formatted OTP request body:", requestBody);
        } else {
          console.error("No phone number found in form values");
          Alert.alert("Error", "Please enter your mobile number");
          return;
        }
      }

      console.log("API CALL →", apiUrl, requestBody);

      const response = await fetch(apiUrl, {
        method: action.method || "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(requestBody),
      });

      let data: any = {};
      try {
        data = await response.json();
      } catch (e) {
        console.log("No JSON response body");
      }

      console.log("API RESPONSE →", response.status, data);

      // FAILURE CASE
      if (!response.ok) {
        const errorMessage = data?.error || data?.message || "Something went wrong";
        
        console.log("FailureNavigate →", action.failureNavigate);

        Alert.alert(
          "Error",
          errorMessage,
          [
            {
              text: "OK",
              onPress: () => {
                if (action.failureNavigate) {
                  setTimeout(() => {
                    router.push(action.failureNavigate);
                  }, 100);
                }
              },
            },
          ],
          { cancelable: false }
        );
        return;
      }

      // SUCCESS CASE - Store phone number for next screen
      if (action.url.includes("request-otp") || action.url.includes("/otp")) {
        const phoneNumber = values.phone || values.mobileNumber;
        if (phoneNumber) {
          // Store phone number globally for next screen
          useFormStore.getState().setValue("phone", phoneNumber);
          console.log("Phone number stored for next screen:", phoneNumber);
        }
      }

      console.log("SUCCESS →", data?.message);
      console.log("SuccessNavigate →", action.successNavigate);

      if (action.successNavigate) {
        console.log("Navigating to:", action.successNavigate);
        setTimeout(() => {
          router.push(action.successNavigate);
        }, 100);
      }

    } catch (err) {
      console.error("API CALL ERROR →", err);
      Alert.alert("Network Error", "Please check your internet connection and try again");
    }
  };

  // Add this helper function
  const handleDirectAction = async (action: any) => {
    try {
      const response = await fetch(action.url || 'http://192.168.1.3:8080/bff/driver/home/action', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          action: action.value,
          data: action.data || {}
        })
      });

      const result = await response.json();
      console.log("Action response:", result);
      
      if (result.status === 'success') {
        // Handle success navigation if specified
        if (action.successNavigate) {
          router.push(action.successNavigate);
        }
      } else {
        Alert.alert("Error", result.message || "Action failed");
      }
    } catch (error) {
      console.error("Action error:", error);
      Alert.alert("Network Error", "Please check your connection");
    }
  };

  return (
    <TouchableOpacity style={node.data.style} onPress={onPress}>
      <Text
        style={{
          color: node.data.style?.textColor || "#FFFFFF",
          fontWeight: "600",
          fontSize: 16,
        }}
      >
        {node.data.text}
      </Text>
    </TouchableOpacity>
  );
}