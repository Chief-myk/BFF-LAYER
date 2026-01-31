// themed.tsx - Minimal version
import React, { ReactNode } from 'react';
import { Text, TextProps, View, ViewProps } from 'react-native';

interface ThemedViewProps extends ViewProps {
  children?: ReactNode;
}

interface ThemedTextProps extends TextProps {
  type?: 'default' | 'title' | 'defaultSemiBold' | 'subtitle' | 'link';
  children?: ReactNode;
}

export function ThemedView({ style, children, ...otherProps }: ThemedViewProps) {
  return (
    <View
      style={[
        { backgroundColor: '#FFFFFF' },
        style,
      ]}
      {...otherProps}
    >
      {children}
    </View>
  );
}

export function ThemedText({ style, type = 'default', children, ...otherProps }: ThemedTextProps) {
  const typeStyles: Record<string, TextProps['style']> = {
    default: {
      fontSize: 14,
      color: '#1A1A1A',
    },
    title: {
      fontSize: 24,
      fontWeight: '700',
      color: '#1A1A1A',
      lineHeight: 32,
    },
    defaultSemiBold: {
      fontSize: 16,
      fontWeight: '600',
      color: '#1A1A1A',
      lineHeight: 24,
    },
    subtitle: {
      fontSize: 18,
      fontWeight: '600',
      color: '#666666',
      lineHeight: 24,
    },
    link: {
      fontSize: 16,
      color: '#ff0000',
      fontWeight: '600',
    },
  };

  return (
    <Text
      style={[
        typeStyles[type],
        style,
      ]}
      {...otherProps}
    >
      {children}
    </Text>
  );
}