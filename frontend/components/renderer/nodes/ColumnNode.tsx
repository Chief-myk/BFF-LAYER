import React from 'react';
import { View, ViewStyle } from 'react-native';

interface ColumnNodeProps {
  data?: {
    gap?: number;
    alignItems?: 'flex-start' | 'center' | 'flex-end' | 'stretch';
    justifyContent?: 'flex-start' | 'center' | 'flex-end' | 'space-between' | 'space-around';
  };
  children?: React.ReactNode;
}

export const ColumnNode: React.FC<ColumnNodeProps> = ({ data = {}, children }) => {
  const columnStyle: ViewStyle = {
    flexDirection: 'column',
    gap: data.gap || 0,
    alignItems: data.alignItems || 'stretch',
    justifyContent: data.justifyContent || 'flex-start',
  };

  return <View style={columnStyle}>{children}</View>;
};