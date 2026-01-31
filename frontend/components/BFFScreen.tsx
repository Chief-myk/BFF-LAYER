// components/BFFScreen.tsx
import { useEffect, useState } from "react";
import { ActivityIndicator, View, Text } from "react-native";
import Renderer from "./renderer/Renderer";
import { useFormStore } from "./renderer/store/formStore";

interface BFFScreenProps {
  endpoint?: string;
  screen?: string;
  data?: any;
  onAction?: (actionType: string, actionData: any) => void;
}

export default function BFFScreen({ 
  endpoint, 
  screen, 
  data = {}, 
  onAction 
}: BFFScreenProps) {
  const [ui, setUi] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { getPhone } = useFormStore();

  useEffect(() => {
    const fetchUI = async () => {
      try {
        let apiEndpoint = endpoint;
        
        if (!apiEndpoint && screen) {
          const endpointMap: Record<string, string> = {
            "home": "/bff/driver/home",
            "market": "/bff/driver/market", 
            "profile": "/bff/driver/profile",
            "payment": "/bff/driver/payment",
            "myTrip": "/bff/driver/mytrip",
            "tripCompleted": "/bff/driver/tripCompleted"
          };
          
          apiEndpoint = endpointMap[screen] || `/bff/${screen}`;
        }
        
        if (!apiEndpoint) {
          throw new Error("No endpoint or screen provided");
        }

        const phone = getPhone();
        const baseUrl = "http://192.168.1.3:8080";
        const url = phone 
          ? `${baseUrl}${apiEndpoint}?phone=${phone}`
          : `${baseUrl}${apiEndpoint}`;
        
        console.log("Fetching UI from:", url);
        
        const response = await fetch(url);
        
        if (!response.ok) {
          throw new Error(`Server error: ${response.status}`);
        }
        
        const json = await response.json();
        
        if (json.status !== "success") {
          throw new Error(json.message || "Failed to load UI");
        }
        
        setUi(json.ui || []);
        setError(null);
      } catch (err: any) {
        console.error("Fetch error:", err.message);
        setError(err.message);
        setUi([]);
      } finally {
        setLoading(false);
      }
    };

    fetchUI();
  }, [endpoint, screen, getPhone]);

  if (loading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" color="#DC2626" />
      </View>
    );
  }

  if (error) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center", padding: 20 }}>
        <Text style={{ color: "#DC2626", fontSize: 16, fontWeight: "bold" }}>Error</Text>
        <Text style={{ color: "#6B7280", marginTop: 8, textAlign: "center" }}>{error}</Text>
      </View>
    );
  }

  return <Renderer ui={ui} contextData={data} onAction={onAction} />;
}