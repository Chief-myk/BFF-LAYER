// app/(auth)/_layout.tsx (Auth Layout)
import { Stack } from 'expo-router';

export default function AuthLayout() {
  return (
      <Stack>
        <Stack.Screen
          name="addTruck"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="addLoad"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="TripCompleted"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="DriverProfile"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="BrokerProfile"
          options={{
            headerShown: false,
          }}
        />
      </Stack>
  );
}
