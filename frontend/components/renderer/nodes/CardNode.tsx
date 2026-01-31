
import React from 'react';
import { View, ViewStyle } from 'react-native';

export interface CardData {
  backgroundColor?: string;
  padding?: number;
  borderRadius?: number;
  shadow?: boolean;
}

interface CardNodeProps {
  data: CardData;
  children?: React.ReactNode;
}

export const CardNode: React.FC<CardNodeProps> = ({ data, children }) => {
  const cardStyle: ViewStyle = {
    backgroundColor: data.backgroundColor || '#FFFFFF',
    padding: data.padding || 16,
    borderRadius: data.borderRadius || 12,
    ...(data.shadow && {
      shadowColor: '#000',
      shadowOffset: { width: 0, height: 2 },
      shadowOpacity: 0.1,
      shadowRadius: 4,
      elevation: 3,
    }),
  };

  return <View style={cardStyle}>{children}</View>;
};