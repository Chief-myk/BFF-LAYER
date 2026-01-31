// app/(auth)/_layout.tsx (Auth Layout)
import { Stack } from 'expo-router';

export default function AuthLayout() {
  return (
      <Stack>
        <Stack.Screen
          name="auth"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="registration-role"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="otp"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r1"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r2"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r3"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r4"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r5"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r6"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r7"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="r8"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="g1"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="g2"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="g3"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="g4"
          options={{
            headerShown: false,
          }}
        />
        <Stack.Screen
          name="g5"
          options={{
            headerShown: false,
          }}
        />
      </Stack>
  );
}
