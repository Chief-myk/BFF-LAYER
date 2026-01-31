import { Tabs } from 'expo-router';
import React from 'react';
import { TouchableOpacity, Text, StyleSheet, Platform, ViewStyle } from 'react-native';
import * as Haptics from 'expo-haptics';
import CustomHeader from '@/ReusableCode/customHeader';

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
    <Tabs
      screenOptions={{
        // tabBarActiveTintColor: Colors[colorScheme ?? 'light'].tint,
        // tabBarInactiveTintColor: Colors[colorScheme ?? 'light'].tabIconDefault,
        tabBarActiveTintColor:  'red',
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
        name="home"
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
        name="myTrip"
        options={{
          title: 'My Trips',
          tabBarIcon: ({ color, focused }) => (
            <Ionicons 
              name={focused ? 'cube' : 'cube-outline'}
              size={22}
              color={color}
            />
          ),
        }}
      />

      <Tabs.Screen
        name="market"
        options={{
          title: 'market',
          tabBarIcon: ({ color, focused }) => (
            <Ionicons 
              name={focused ? 'paper-plane' : 'paper-plane-outline'}
              size={22}
              color={color}
            />
          ),
        }}
      />

      <Tabs.Screen
        name="payment"
        options={{
          title: 'payment',
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
  );
}