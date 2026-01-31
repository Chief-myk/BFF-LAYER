import { View, Text } from "react-native";
import SafeAreaNode from "./nodes/SafeAreaNode";
import ScrollNode from "./nodes/ScrollNode";
import ViewNode from "./nodes/ViewNode";
import TextNode from "./nodes/TextNode";
import ImageNode from "./nodes/ImageNode";
import InputNode from "./nodes/InputNode";
import ButtonNode from "./nodes/ButtonNode";
import NavigateNode from "./nodes/NavigateNode";
import OtpInputNode from "./nodes/OtpInputNode";
import PressableCardNode from "./nodes/PressableCardNode";
import IconNode from "./nodes/IconNode";
import TextButtonNode from "./nodes/TextButtonNode";
import ResendOtpNode from "./nodes/ResendOtpNode";
import IconButtonNode from "./nodes/IconButtonNode";

// Map node types from backend to React components
const NODE_MAP: Record<string, any> = {
  SAFE_AREA: SafeAreaNode,
  SCROLL: ScrollNode,
  VIEW: ViewNode,
  TEXT: TextNode,
  IMAGE: ImageNode,
  INPUT: InputNode,
  BUTTON: ButtonNode,
  NAVIGATE: NavigateNode,
  OTP_INPUT: OtpInputNode,
  RESEND_OTP: ResendOtpNode,
  PRESSABLE_CARD: PressableCardNode,
  ICON: IconNode,
  TEXT_BUTTON: TextButtonNode,
  ICON_BUTTON: IconButtonNode,
  FOOTER_NOTE: TextNode,
};

type RendererProps = {
  ui: any[];
  contextData?: any; // ADD THIS LINE ONLY
  onAction?: (actionType: string, actionData: any) => void; // ADD THIS LINE ONLY
};

export default function Renderer({ ui, contextData, onAction }: RendererProps) { // ADD contextData, onAction HERE
  const renderNode = (node: any, index: number) => {
    // Check if node is a primitive string
    if (typeof node === 'string' || typeof node === 'number') {
      // Wrap strings/numbers in Text component
      return <Text key={index}>{node}</Text>;
    }
    
    const Component = NODE_MAP[node.type];
    if (!Component) {
      console.warn(`Unknown node type: ${node.type}`);
      return null;
    }
    
    // ADD THESE 2 LINES - Pass contextData and onAction to nodes
    const props = { 
      node, 
      renderNode, 
      contextData: contextData || {}, 
      onAction: onAction 
    };
    
    return <Component key={index} {...props} />;
  };

  return <View style={{ flex: 1 }}>{ui.map(renderNode)}</View>;
}