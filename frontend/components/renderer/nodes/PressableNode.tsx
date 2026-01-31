import React from 'react';
import { Pressable, ViewStyle } from 'react-native';

interface PressableData {
  flex?: number;
  alignItems?: string;
  paddingVertical?: number;
  paddingHorizontal?: number;
  justifyContent?: string;
  backgroundColor?: string;
  borderRadius?: number;
  marginBottom?: number;
  // Action data
  onPress?: {
    type: string;
    value?: string;
    to?: string;
  };
}

interface PressableNodeProps {
  node: {
    data: PressableData;
    children?: any[];
  };
  renderNode: (node: any, index: number) => React.ReactNode;
}

export const PressableNode: React.FC<PressableNodeProps> = ({ node, renderNode }) => {
  const pressableStyle: ViewStyle = {
    flex: node.data?.flex,
    alignItems: node.data?.alignItems as any,
    paddingVertical: node.data?.paddingVertical,
    paddingHorizontal: node.data?.paddingHorizontal,
    justifyContent: node.data?.justifyContent as any,
    backgroundColor: node.data?.backgroundColor,
    borderRadius: node.data?.borderRadius,
    marginBottom: node.data?.marginBottom,
  };

  const handlePress = () => {
    console.log('Pressable pressed');
    
    // Handle different action types
    if (node.data?.onPress) {
      const { type, value, to } = node.data.onPress;
      
      switch (type) {
        case 'NAVIGATE':
          console.log(`Navigate to: ${to}`);
          // You would dispatch navigation action here
          break;
        case 'ACTION':
          console.log(`Action: ${value}`);
          // Handle custom action
          break;
        case 'LOGOUT':
          console.log('Logout action');
          // Handle logout
          break;
        default:
          console.log(`Unknown action type: ${type}`);
      }
    }
  };

  return (
    <Pressable style={pressableStyle} onPress={handlePress}>
      {node.children?.map((child, index) => renderNode(child, index))}
    </Pressable>
  );
};