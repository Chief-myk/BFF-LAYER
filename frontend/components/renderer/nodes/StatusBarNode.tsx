import React from 'react';
import { StatusBar } from 'react-native';

export interface StatusBarData {
  backgroundColor?: string;
  style?: 'light' | 'dark';
}

interface StatusBarNodeProps {
  data: StatusBarData;
}

export const StatusBarNode: React.FC<StatusBarNodeProps> = ({ data }) => {
  const barStyle = data.style === 'light' ? 'light-content' : 'dark-content';
  
  return (
    <StatusBar
      backgroundColor={data.backgroundColor}
      barStyle={barStyle}
    />
  );
};