import { ScrollView, StyleSheet } from "react-native";

export default function ScrollNode({ node, renderNode }: any) {
  const { data, children = [] } = node;
  
  // Extract style properties for the ScrollView
  const scrollViewStyle = {
    flex: data?.flexGrow || data?.flex || 1,
    backgroundColor: data?.backgroundColor,
  };

  // Extract style properties for the content container
  const contentContainerStyle = {
    flexGrow: 1,
    paddingHorizontal: data?.paddingHorizontal,
    paddingVertical: data?.paddingVertical,
    justifyContent: data?.justifyContent,
    alignItems: data?.alignItems,
    backgroundColor: data?.backgroundColor,
  };

  return (
    <ScrollView
      style={[styles.scrollView, scrollViewStyle]}
      contentContainerStyle={[styles.contentContainer, contentContainerStyle]}
      showsVerticalScrollIndicator={true}
      showsHorizontalScrollIndicator={false}
      bounces={true}
      scrollEnabled={true}
    >
      {children.map((child: any, index: number) => 
        renderNode ? renderNode(child, index) : null
      )}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  scrollView: {
    flex: 1,
  },
  contentContainer: {
    flexGrow: 1,
  },
});