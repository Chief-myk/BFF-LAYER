import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import 'react-native-reanimated';
import { useFonts } from 'expo-font';
import * as SplashScreen from 'expo-splash-screen';
import { useEffect } from 'react';
import { ColorSchemeName, useColorScheme as _useColorScheme } from 'react-native'; // Import from react-native directly
import { UserTypeProvider } from '@/context/UserTypeContext';
import { UserPhoneProvider } from "@/context/UserPhoneContent";
import { UserBrokerDataProvider } from '@/context/UserBrokerDataContext'; // Import the provider
import {UserDriverDataProvider} from "@/context/UserDriverDataContext";
import "@/service/i18n";


SplashScreen.preventAutoHideAsync();

export const unstable_settings = {
  anchor: '(tabs)',
};

// Define useColorScheme hook directly in this file
function useColorScheme(): NonNullable<ColorSchemeName> {
  const colorScheme = _useColorScheme();
  return colorScheme ?? 'light';
}

export default function RootLayout() {
  const [fontsLoaded, fontError] = useFonts({
    // 'Inter-Black': require('./assets/fonts/Inter-Black.otf'),
  });

  const colorScheme = useColorScheme();

  // useEffect(() => {
  //   if (fontsLoaded || fontError) {
  //     SplashScreen.hideAsync();
  //   }
  // }, [fontsLoaded, fontError]);

  // if (!fontsLoaded && !fontError) {
  //   return null;
  // }

  useEffect(() => {
    // Hide splash screen once app is ready
    SplashScreen.hideAsync();
  }, []);

  return (
    <UserTypeProvider>
      <UserPhoneProvider>
        <UserBrokerDataProvider>
          <UserDriverDataProvider>

          <ThemeProvider value={colorScheme === 'dark' ? DarkTheme : DefaultTheme}>
            <Stack>
              <Stack.Screen
                name="index"
                options={{
                  headerShown: false,
                }}
              />
              <Stack.Screen
                name="(tabs)"
                options={{
                  headerShown: false,
                }}
              />
              <Stack.Screen
                name="(auth)"
                options={{
                  headerShown: false,
                }}
              />
              {/* <Stack.Screen
                name="(profile)"
                options={{
                  headerShown: false,
                }}
              /> */}
              <Stack.Screen
                name="(footbar)"
                options={{
                  headerShown: false,
                }}
              />
              <Stack.Screen
                name="(pages)"
                options={{
                  headerShown: false,
                }}
              />
            </Stack>
            <StatusBar style="auto" />
          </ThemeProvider>

          </UserDriverDataProvider>
        </UserBrokerDataProvider>
      </UserPhoneProvider>
    </UserTypeProvider>

  );
}