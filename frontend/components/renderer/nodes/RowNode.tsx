import React from 'react';
import { View, ViewStyle } from 'react-native';

interface RowNodeProps {
  data?: {
    gap?: number;
    alignItems?: 'flex-start' | 'center' | 'flex-end' | 'stretch';
    justifyContent?: 'flex-start' | 'center' | 'flex-end' | 'space-between' | 'space-around';
  };
  children?: React.ReactNode;
}

export const RowNode: React.FC<RowNodeProps> = ({ data = {}, children }) => {
  const rowStyle: ViewStyle = {
    flexDirection: 'row',
    gap: data.gap || 0,
    alignItems: data.alignItems || 'center',
    justifyContent: data.justifyContent || 'flex-start',
  };

  return <View style={rowStyle}>{children}</View>;
};