// app/(auth)/auth.tsx
import BFFScreen from "@/components/BFFScreen";

export default function Auth() {
  return <BFFScreen endpoint="/bff/auth/auth" />;
}

