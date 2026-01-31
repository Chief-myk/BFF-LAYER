import { Tabs } from 'expo-router';
import React, { createContext, useState, useContext, ReactNode } from 'react';
import { TouchableOpacity, Text, StyleSheet, Platform, ViewStyle } from 'react-native';
import * as Haptics from 'expo-haptics';
import CustomHeader from '@/ReusableCode/customHeader';

// UserType Context
type UserType = 'driver' | 'shipper' | undefined;

interface UserTypeContextType {
  userType: UserType;
  setUserType: (type: UserType) => void;
}

const UserTypeContext = createContext<UserTypeContextType | undefined>(undefined);

export const useUserType = () => {
  const context = useContext(UserTypeContext);
  if (!context) {
    throw new Error('useUserType must be used within a UserTypeProvider');
  }
  return context;
};

interface UserTypeProviderProps {
  children: ReactNode;
}

export const UserTypeProvider = ({ children }: UserTypeProviderProps) => {
  const [userType, setUserType] = useState<UserType>(undefined);

  return (
    <UserTypeContext.Provider value={{ userType, setUserType }}>
      {children}
    </UserTypeContext.Provider>
  );
};

// Custom HapticTab component
const HapticTab = (props: any) => {
  const handlePress = () => {
    if (Platform.OS !== 'web') {
      Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
    }
    props.onPress?.();
  };

  return (
    <TouchableOpacity
      {...props}
      onPress={handlePress}
      style={[props.style, styles.tabButton]}
      activeOpacity={0.7}
    >
      {props.children}
    </TouchableOpacity>
  );
};

// Colors constant
const Colors = {
  light: {
    background: '#FFFFFF',
    tint: 'red',
    tabIconDefault: 'grey',
    tabBarBorder: '#E5E5E5',
  },
  dark: {
    background: '#000000',
    tint: 'red',
    tabIconDefault: '#9E9E9E',
    tabBarBorder: '#2C2C2C',
  },
};

// useColorScheme hook
const useColorScheme = () => {
  // For simplicity, returning 'light' as default
  // In a real app, you would implement actual theme detection
  return 'light';
};

// Styles for HapticTab
const styles = StyleSheet.create({
  tabButton: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});


// Import icons
import Ionicons from '@expo/vector-icons/Ionicons';

export default function TabLayout() {
  const colorScheme = useColorScheme();

  return (
    <UserTypeProvider>
      <Tabs
        screenOptions={{
          // tabBarActiveTintColor: Colors[colorScheme ?? 'light'].tint,
          // tabBarInactiveTintColor: Colors[colorScheme ?? 'light'].tabIconDefault,
          tabBarActiveTintColor: 'red',
          tabBarInactiveTintColor: 'grey',
          header: () => <CustomHeader />,
          tabBarButton: HapticTab,
          tabBarStyle: {
            backgroundColor: Colors[colorScheme ?? 'light'].background,
            borderTopWidth: 1,
            borderTopColor: Colors[colorScheme ?? 'light'].tabBarBorder,
            elevation: 8,
            shadowColor: '#000',
            shadowOffset: { width: 0, height: -2 },
            shadowOpacity: 0.08,
            shadowRadius: 8,
          },
          tabBarLabelStyle: {
            fontSize: 12,
            fontWeight: '500',
            marginTop: -2,
            marginBottom: 4,
          },
        }}
      >
        <Tabs.Screen
          name="index"
          options={{
            title: 'Home',
            tabBarIcon: ({ color, focused }) => (
              <Ionicons
                name={focused ? 'home' : 'home-outline'}
                size={22}
                color={color}
              />
            ),
          }}
        />

        <Tabs.Screen
          name="load"
          options={{
            title: 'Loads',
            tabBarIcon: ({ color, focused }) => (
              <Ionicons
                name={focused ? 'search' : 'search-outline'}
                size={22}
                color={color}
              />
            ),
          }}
        />

        <Tabs.Screen
          name="liveTrip"
          options={{
            title: 'LiveTrips',
            tabBarIcon: ({ color, focused }) => (
              <Ionicons
                name={focused ? 'document-text' : 'document-text-outline'}
                size={22}
                color={color}
              />
            ),
          }}
        />

        <Tabs.Screen
          name="money"
          options={{
            title: 'Payment',
            tabBarIcon: ({ color, focused }) => (
              <Ionicons
                name={focused ? 'wallet' : 'wallet-outline'}
                size={22}
                color={color}
              />
            ),
          }}
        />
      </Tabs>
    </UserTypeProvider>
  );
}