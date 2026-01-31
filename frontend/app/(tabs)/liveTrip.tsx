import React, { useState, useEffect } from 'react';
import {
  View,
  Text,
  ScrollView,
  TouchableOpacity,
  StyleSheet,
  StatusBar,
  ActivityIndicator,
  RefreshControl,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
// Define TypeScript interfaces based on your backend structure
interface UISnippet {
  type: string;
  data: any;
  children?: UISnippet[];
}

interface ScreenResponse {
  status: string;
  screen: string;
  ui: UISnippet[];
  data: any;
}

// Component to render UI snippets recursively
const UIRenderer: React.FC<{ snippet: UISnippet; data?: any }> = ({ snippet, data }) => {
  if (!snippet) return null;

  const { type, data: snippetData, children } = snippet;

  const renderChildren = () => {
    if (!children || children.length === 0) return null;
    return children.map((child, index) => (
      <UIRenderer key={index} snippet={child} data={data} />
    ));
  };

  // Helper function to convert style object
  const parseStyle = (styleData: any) => {
    const style: any = {};
    
    // Map backend style properties to React Native style properties
    if (styleData) {
      // Layout
      if (styleData.flex !== undefined) style.flex = styleData.flex;
      if (styleData.flexDirection !== undefined) style.flexDirection = styleData.flexDirection;
      if (styleData.justifyContent !== undefined) style.justifyContent = styleData.justifyContent;
      if (styleData.alignItems !== undefined) style.alignItems = styleData.alignItems;
      if (styleData.gap !== undefined) style.gap = styleData.gap;
      
      // Spacing
      if (styleData.padding !== undefined) style.padding = styleData.padding;
      if (styleData.paddingHorizontal !== undefined) style.paddingHorizontal = styleData.paddingHorizontal;
      if (styleData.paddingVertical !== undefined) style.paddingVertical = styleData.paddingVertical;
      if (styleData.paddingTop !== undefined) style.paddingTop = styleData.paddingTop;
      if (styleData.paddingBottom !== undefined) style.paddingBottom = styleData.paddingBottom;
      if (styleData.paddingLeft !== undefined) style.paddingLeft = styleData.paddingLeft;
      if (styleData.paddingRight !== undefined) style.paddingRight = styleData.paddingRight;
      
      if (styleData.margin !== undefined) style.margin = styleData.margin;
      if (styleData.marginBottom !== undefined) style.marginBottom = styleData.marginBottom;
      if (styleData.marginTop !== undefined) style.marginTop = styleData.marginTop;
      if (styleData.marginLeft !== undefined) style.marginLeft = styleData.marginLeft;
      if (styleData.marginRight !== undefined) style.marginRight = styleData.marginRight;
      if (styleData.marginHorizontal !== undefined) style.marginHorizontal = styleData.marginHorizontal;
      if (styleData.marginVertical !== undefined) style.marginVertical = styleData.marginVertical;
      
      // Dimensions
      if (styleData.width !== undefined) style.width = styleData.width;
      if (styleData.height !== undefined) style.height = styleData.height;
      if (styleData.maxHeight !== undefined) style.maxHeight = styleData.maxHeight;
      
      // Borders
      if (styleData.borderRadius !== undefined) style.borderRadius = styleData.borderRadius;
      if (styleData.borderWidth !== undefined) style.borderWidth = styleData.borderWidth;
      if (styleData.borderColor !== undefined) style.borderColor = styleData.borderColor;
      if (styleData.borderBottomWidth !== undefined) style.borderBottomWidth = styleData.borderBottomWidth;
      
      // Colors
      if (styleData.backgroundColor !== undefined) style.backgroundColor = styleData.backgroundColor;
      if (styleData.color !== undefined) style.color = styleData.color;
      
      // Shadows (iOS)
      if (styleData.shadowColor !== undefined) style.shadowColor = styleData.shadowColor;
      if (styleData.shadowOffsetX !== undefined) style.shadowOffset = { width: styleData.shadowOffsetX, height: styleData.shadowOffsetY || 0 };
      if (styleData.shadowOpacity !== undefined) style.shadowOpacity = styleData.shadowOpacity;
      if (styleData.shadowRadius !== undefined) style.shadowRadius = styleData.shadowRadius;
      
      // Elevation (Android)
      if (styleData.elevation !== undefined) style.elevation = styleData.elevation;
      
      // ScrollView specific
      if (styleData.horizontal !== undefined) style.horizontal = styleData.horizontal;
      if (styleData.showsHorizontalScrollIndicator !== undefined) style.showsHorizontalScrollIndicator = styleData.showsHorizontalScrollIndicator;
      if (styleData.showsVerticalScrollIndicator !== undefined) style.showsVerticalScrollIndicator = styleData.showsVerticalScrollIndicator;
    }
    
    return style;
  };

  // Handle Text component
  if (type === 'Text') {
    const textStyle: any = {};
    if (snippetData.fontSize !== undefined) textStyle.fontSize = snippetData.fontSize;
    if (snippetData.fontWeight !== undefined) textStyle.fontWeight = snippetData.fontWeight;
    if (snippetData.color !== undefined) textStyle.color = snippetData.color;
    
    // Merge with any additional style from snippetData.style
    const containerStyle = parseStyle(snippetData.style || {});
    
    return (
      <Text style={[textStyle, containerStyle]}>
        {snippetData.text}
      </Text>
    );
  }

  // Handle View component
  if (type === 'View') {
    return (
      <View style={parseStyle(snippetData)}>
        {renderChildren()}
      </View>
    );
  }

  // Handle ScrollView component
  if (type === 'ScrollView') {
    const scrollViewProps: any = {
      style: parseStyle(snippetData),
      horizontal: snippetData.horizontal || false,
      showsHorizontalScrollIndicator: snippetData.showsHorizontalScrollIndicator || false,
      showsVerticalScrollIndicator: snippetData.showsVerticalScrollIndicator || false,
    };
    
    return (
      <ScrollView {...scrollViewProps}>
        {renderChildren()}
      </ScrollView>
    );
  }

  // Handle TouchableOpacity component
  if (type === 'TouchableOpacity') {
    const handlePress = () => {
      if (snippetData.onPress) {
        const action = snippetData.onPress;
        
        // Handle different action types
        switch (action.type) {
          case 'navigate':
            console.log(`Navigate to: ${action.to}`);
            // You would integrate with your navigation here
            // navigation.navigate(action.to);
            break;
          
          case 'updateScreen':
            console.log('Update screen with:', action.data);
            // You would update your state here
            break;
          
          case 'call':
            console.log('Call:', action.data.phone);
            // You would initiate a call here
            // Linking.openURL(`tel:${action.data.phone}`);
            break;
          
          case 'message':
            console.log('Message to:', action.data.driver);
            // You would open messaging here
            break;
          
          default:
            console.log('Action pressed:', action);
        }
      }
    };

    return (
      <TouchableOpacity 
        style={parseStyle(snippetData.style || snippetData)} 
        onPress={handlePress}
        activeOpacity={0.7}
      >
        {renderChildren()}
      </TouchableOpacity>
    );
  }

  // Handle StatusBar component
  if (type === 'StatusBar') {
    return (
      <StatusBar
        backgroundColor={snippetData.backgroundColor || '#FFFFFF'}
        barStyle={snippetData.style === 'dark' ? 'dark-content' : 'light-content'}
      />
    );
  }

  // Handle Icon component (simplified - you'd use your icon library)
  if (type === 'Icon') {
    // This is a placeholder for icon rendering
    // You would replace this with your actual icon component
    return (
      <View style={{
        width: snippetData.size || 24,
        height: snippetData.size || 24,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: snippetData.bgColor || 'transparent',
        borderRadius: snippetData.borderRadius || 0,
      }}>
        <Text style={{
          fontSize: (snippetData.size || 24) * 0.8,
          color: snippetData.color || '#333',
          fontWeight: 'bold',
        }}>
          {snippetData.name ? snippetData.name.charAt(0).toUpperCase() : 'I'}
        </Text>
      </View>
    );
  }

  // Default fallback
  console.warn(`Unknown UI component type: ${type}`);
  return null;
};

// Main LiveTrip Component
const LiveTrip: React.FC = () => {
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [screenData, setScreenData] = useState<ScreenResponse | null>(null);

  const fetchLiveTripData = async (isRefreshing = false) => {
    try {
      if (!isRefreshing) {
        setLoading(true);
      }
      setError(null);

      // Replace with your actual API endpoint
      const response = await fetch('http://192.168.1.3:8080/bff/broker/livetrip', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data: ScreenResponse = await response.json();
      
      if (data.status === 'success') {
        setScreenData(data);
      } else {
        throw new Error('Failed to load screen data');
      }
    } catch (err) {
      console.error('Error fetching live trip data:', err);
      setError(err instanceof Error ? err.message : 'Network request failed');
      
      // Fallback: Create a basic UI structure if API fails
      if (!screenData) {
        setScreenData({
          status: 'success',
          screen: 'liveTrip',
          ui: [{
            type: 'View',
            data: { flex: 1, backgroundColor: '#FFFFFF' },
            children: [
              {
                type: 'StatusBar',
                data: { backgroundColor: '#FFFFFF', style: 'dark' }
              },
              {
                type: 'View',
                data: { 
                  flex: 1, 
                  justifyContent: 'center', 
                  alignItems: 'center',
                  padding: 20
                },
                children: [
                  {
                    type: 'Text',
                    data: {
                      text: 'Live Trips',
                      fontSize: 24,
                      fontWeight: '700',
                      color: '#333',
                      marginBottom: 10
                    }
                  },
                  {
                    type: 'Text',
                    data: {
                      text: 'No trip data available',
                      fontSize: 16,
                      color: '#666',
                      textAlign: 'center'
                    }
                  }
                ]
              }
            ]
          }],
          data: {}
        });
      }
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  };

  useEffect(() => {
    fetchLiveTripData();
  }, []);

  const onRefresh = () => {
    setRefreshing(true);
    fetchLiveTripData(true);
  };

  if (loading && !refreshing) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
        <ActivityIndicator size="large" color="#ff0000" />
        <Text style={styles.loadingText}>Loading Live Trips...</Text>
      </SafeAreaView>
    );
  }

  if (error && !screenData) {
    return (
      <SafeAreaView style={styles.errorContainer}>
        <StatusBar backgroundColor="#FFEBEE" barStyle="dark-content" />
        <View style={styles.errorContent}>
          <Text style={styles.errorIcon}>⚠️</Text>
          <Text style={styles.errorTitle}>Network Issue</Text>
          <Text style={styles.errorMessage}>
            {error || 'Unable to load live trip data. Please check your connection.'}
          </Text>
          <TouchableOpacity 
            style={styles.retryButton}
            onPress={() => fetchLiveTripData()}
          >
            <Text style={styles.retryButtonText}>Retry</Text>
          </TouchableOpacity>
        </View>
      </SafeAreaView>
    );
  }

  return (
  <SafeAreaView style={styles.container}>
    <ScrollView
      refreshControl={
        <RefreshControl
          refreshing={refreshing}
          onRefresh={onRefresh}
          colors={['#ff0000']}
          tintColor="#ff0000"
        />
      }
      style={{ flex: 1 }}
    >
      <View style={styles.screenContainer}>
        {screenData?.ui.map((snippet, index) => (
          <UIRenderer key={index} snippet={snippet} data={screenData.data} />
        ))}
      </View>
    </ScrollView>
  </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#FFFFFF',
  },
  screenContainer: {
    flex: 1,
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#FFFFFF',
  },
  loadingText: {
    marginTop: 12,
    fontSize: 16,
    color: '#666',
  },
  errorContainer: {
    flex: 1,
    backgroundColor: '#FFEBEE',
  },
  errorContent: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  errorIcon: {
    fontSize: 48,
    marginBottom: 16,
  },
  errorTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#D32F2F',
    marginBottom: 8,
  },
  errorMessage: {
    fontSize: 16,
    color: '#666',
    textAlign: 'center',
    marginBottom: 24,
  },
  retryButton: {
    backgroundColor: '#ff0000',
    paddingHorizontal: 32,
    paddingVertical: 12,
    borderRadius: 8,
  },
  retryButtonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: '600',
  },
  debugPanel: {
    backgroundColor: '#F5F5F5',
    padding: 8,
    borderTopWidth: 1,
    borderTopColor: '#E0E0E0',
  },
  debugText: {
    fontSize: 12,
    color: '#666',
    textAlign: 'center',
  },
});

export default LiveTrip;