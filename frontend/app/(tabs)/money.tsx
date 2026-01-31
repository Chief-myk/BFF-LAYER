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
  Modal,
  Dimensions,
  Alert,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
// import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import { Ionicons } from '@expo/vector-icons';
// TypeScript interfaces
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

// Icon name mapping from backend to MaterialCommunityIcons
const iconMap: Record<string, string> = {
  'clock-outline': 'clock-outline',
  'check-circle-outline': 'check-circle-outline',
  'calendar-month': 'calendar-month',
  'chevron-down': 'chevron-down',
  'chevron-right': 'chevron-right',
  'close': 'close',
  'file-download': 'file-download',
  'credit-card': 'credit-card',
  'truck': 'truck',
  'map-marker': 'map-marker',
  'plus': 'plus',
  'refresh': 'refresh',
  'call': 'phone',
  'chatbubble': 'message-text-outline',
  'location-on': 'map-marker',
  'person': 'account-circle',
  'filter': 'filter-variant',
  'cash-multiple': 'cash-multiple',
  'alert-circle-outline': 'alert-circle-outline',
};

// Main Money Component
const Money: React.FC = () => {
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [screenData, setScreenData] = useState<ScreenResponse | null>(null);
  const [modalData, setModalData] = useState<ScreenResponse | null>(null);
  const [showModal, setShowModal] = useState(false);
  const { width, height } = Dimensions.get('window');

  // Parse style object from backend data
  const parseStyle = (styleData: any) => {
    const style: any = {};
    
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
      if (styleData.borderTopWidth !== undefined) style.borderTopWidth = styleData.borderTopWidth;
      if (styleData.borderBottomWidth !== undefined) style.borderBottomWidth = styleData.borderBottomWidth;
      if (styleData.borderTopLeftRadius !== undefined) style.borderTopLeftRadius = styleData.borderTopLeftRadius;
      if (styleData.borderTopRightRadius !== undefined) style.borderTopRightRadius = styleData.borderTopRightRadius;
      
      // Colors
      if (styleData.backgroundColor !== undefined) style.backgroundColor = styleData.backgroundColor;
      if (styleData.color !== undefined) style.color = styleData.color;
      
      // Shadows
      if (styleData.shadowColor !== undefined) style.shadowColor = styleData.shadowColor;
      if (styleData.shadowOffsetX !== undefined || styleData.shadowOffsetY !== undefined) {
        style.shadowOffset = { 
          width: styleData.shadowOffsetX || 0, 
          height: styleData.shadowOffsetY || 0 
        };
      }
      if (styleData.shadowOpacity !== undefined) style.shadowOpacity = styleData.shadowOpacity;
      if (styleData.shadowRadius !== undefined) style.shadowRadius = styleData.shadowRadius;
      
      // Elevation (Android)
      if (styleData.elevation !== undefined) style.elevation = styleData.elevation;
      
      // Position
      if (styleData.position !== undefined) style.position = styleData.position;
      if (styleData.bottom !== undefined) style.bottom = styleData.bottom;
      if (styleData.left !== undefined) style.left = styleData.left;
      if (styleData.right !== undefined) style.right = styleData.right;
      if (styleData.top !== undefined) style.top = styleData.top;
    }
    
    return style;
  };

  // UI Renderer component
  const UIRenderer: React.FC<{ 
    snippet: UISnippet; 
    data?: any;
    onAction?: (action: any) => void;
  }> = ({ snippet, data, onAction }) => {
    if (!snippet) return null;

    const { type, data: snippetData, children } = snippet;

    const renderChildren = () => {
      if (!children || children.length === 0) return null;
      return children.map((child, index) => (
        <UIRenderer 
          key={index} 
          snippet={child} 
          data={data}
          onAction={onAction}
        />
      ));
    };

    const handlePress = () => {
      if (snippetData.onPress && onAction) {
        onAction(snippetData.onPress);
      }
    };

    // Handle Text component
    if (type === 'Text') {
      const textStyle: any = {};
      if (snippetData.fontSize !== undefined) textStyle.fontSize = snippetData.fontSize;
      if (snippetData.fontWeight !== undefined) textStyle.fontWeight = snippetData.fontWeight;
      if (snippetData.color !== undefined) textStyle.color = snippetData.color;
      if (snippetData.textAlign !== undefined) textStyle.textAlign = snippetData.textAlign;
      
      const containerStyle = parseStyle(snippetData.style || {});
      
      return (
        <Text style={[textStyle, containerStyle]}>
          {snippetData.text}
        </Text>
      );
    }

    // Handle View component
    if (type === 'View') {
      const style = parseStyle(snippetData);
      
      // Handle horizontal property for ScrollView inside
      if (snippetData.horizontal !== undefined) {
        return (
          <ScrollView
            horizontal={snippetData.horizontal}
            showsHorizontalScrollIndicator={snippetData.showsHorizontalScrollIndicator || false}
            showsVerticalScrollIndicator={snippetData.showsVerticalScrollIndicator || false}
            style={style}
          >
            {renderChildren()}
          </ScrollView>
        );
      }
      
      return (
        <View style={style}>
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

    // Handle Icon component
    if (type === 'Icon') {
      const iconName = iconMap[snippetData.name] || snippetData.name || 'help-circle';
      const iconStyle = parseStyle(snippetData.style || {});
      
      return (
        <Ionicons
          name={iconName}
          size={snippetData.size || 24}
          color={snippetData.color || '#333'}
          style={iconStyle}
        />
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

    // Handle Modal component (simplified)
    if (type === 'Modal') {
      // Modal is handled separately at top level
      return renderChildren();
    }

    // Default fallback
    return null;
  };

  // Fetch main money data
  const fetchMoneyData = async (isRefreshing = false) => {
    try {
      if (!isRefreshing) {
        setLoading(true);
      }
      setError(null);

      // Replace with your actual API endpoint
      const response = await fetch('http://192.168.1.3:8080/bff/broker/money', {
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
        throw new Error('Failed to load money data');
      }
    } catch (err) {
      console.error('Error fetching money data:', err);
      setError(err instanceof Error ? err.message : 'Network request failed');
      
      // Fallback: Create a basic UI structure if API fails
      if (!screenData) {
        setScreenData({
          status: 'success',
          screen: 'money',
          ui: [{
            type: 'View',
            data: { flex: 1, backgroundColor: '#FFFFFF' },
            children: [
              // Header
              {
                type: 'View',
                data: { 
                  paddingHorizontal: 20,
                  paddingTop: 10,
                  marginTop: 10,
                  paddingBottom: 20,
                  backgroundColor: '#fff',
                  borderBottomWidth: 1,
                  borderColor: '#f0f0f0'
                },
                children: [{
                  type: 'Text',
                  data: {
                    text: 'Payments & Settlements',
                    fontSize: 24,
                    fontWeight: 'bold',
                    color: '#1a1a1a',
                  }
                }]
              },
              // Fallback content
              {
                type: 'View',
                data: { 
                  flex: 1, 
                  justifyContent: 'center', 
                  alignItems: 'center',
                  padding: 40
                },
                children: [
                  {
                    type: 'Icon',
                    data: {
                      name: 'cash-multiple',
                      size: 64,
                      color: '#ccc',
                      style: { marginBottom: 20 }
                    }
                  },
                  {
                    type: 'Text',
                    data: {
                      text: 'Unable to load payment data',
                      fontSize: 18,
                      fontWeight: '600',
                      color: '#666',
                      textAlign: 'center',
                      marginBottom: 8
                    }
                  },
                  {
                    type: 'Text',
                    data: {
                      text: 'Please check your connection and try again',
                      fontSize: 14,
                      color: '#999',
                      textAlign: 'center',
                      marginBottom: 24
                    }
                  },
                  {
                    type: 'TouchableOpacity',
                    data: {
                      style: {
                        backgroundColor: '#ff0000',
                        paddingHorizontal: 32,
                        paddingVertical: 12,
                        borderRadius: 8
                      },
                      onPress: { type: 'retry' }
                    },
                    children: [{
                      type: 'Text',
                      data: {
                        text: 'Retry',
                        color: '#fff',
                        fontSize: 16,
                        fontWeight: '600',
                      }
                    }]
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

  // Fetch payment detail modal data
  const fetchPaymentDetail = async (paymentId: string) => {
    try {
      const response = await fetch(`http://your-backend-url/broker/payment-detail?paymentId=${paymentId}`, {
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
        setModalData(data);
        setShowModal(true);
      } else {
        throw new Error('Failed to load payment details');
      }
    } catch (err) {
      console.error('Error fetching payment details:', err);
      Alert.alert('Error', 'Failed to load payment details');
    }
  };

  // Handle all actions from UI
  const handleAction = (action: any) => {
    console.log('Action triggered:', action);
    
    switch (action.type) {
      case 'retry':
        fetchMoneyData();
        break;
        
      case 'dateFilter':
        Alert.alert('Date Filter', 'Select date range');
        break;
      
      case 'updateScreen':
        console.log('Update screen with:', action.data);
        if (action.data.selectedFilter) {
          Alert.alert('Filter Changed', `Selected: ${action.data.selectedFilter}`);
        }
        break;
      
      case 'showModal':
        if (action.data && action.data.paymentId) {
          fetchPaymentDetail(action.data.paymentId);
        }
        break;
      
      case 'closeModal':
        setShowModal(false);
        setModalData(null);
        break;
      
      case 'downloadInvoice':
        Alert.alert('Download', 'Invoice download started');
        break;
      
      case 'makePayment':
        Alert.alert('Make Payment', 'Redirecting to payment gateway');
        break;
      
      case 'navigate':
        console.log(`Navigate to: ${action.to}`);
        Alert.alert('Navigation', `Would navigate to: ${action.to}`);
        break;
      
      case 'call':
        Alert.alert('Call', `Calling: ${action.data.phone}`);
        break;
      
      case 'message':
        Alert.alert('Message', `Messaging driver: ${action.data.driver}`);
        break;
      
      default:
        console.log('Unhandled action:', action);
    }
  };

  // Initial load
  useEffect(() => {
    fetchMoneyData();
  }, []);

  // Pull to refresh
  const onRefresh = () => {
    setRefreshing(true);
    fetchMoneyData(true);
  };

  // Loading state
  if (loading && !refreshing) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
        <ActivityIndicator size="large" color="#ff0000" />
        <Text style={styles.loadingText}>Loading Payments...</Text>
      </SafeAreaView>
    );
  }

  // Error state (no data loaded yet)
  if (error && !screenData) {
    return (
      <SafeAreaView style={styles.errorContainer}>
        <StatusBar backgroundColor="#FFEBEE" barStyle="dark-content" />
        <View style={styles.errorContent}>
          <Icon name="alert-circle-outline" size={64} color="#D32F2F" />
          <Text style={styles.errorTitle}>Connection Issue</Text>
          <Text style={styles.errorMessage}>
            {error || 'Unable to load payment data. Please check your connection.'}
          </Text>
          <TouchableOpacity 
            style={styles.retryButton}
            onPress={() => fetchMoneyData()}
          >
            <Icon name="refresh" size={20} color="#fff" />
            <Text style={styles.retryButtonText}>Try Again</Text>
          </TouchableOpacity>
        </View>
      </SafeAreaView>
    );
  }

  // Main render
  return (
    <>
      <SafeAreaView style={styles.container}>
        <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
        <RefreshControl
          refreshing={refreshing}
          onRefresh={onRefresh}
          style={{ flex: 1 }}
        >
          <ScrollView 
            style={styles.scrollContainer}
            showsVerticalScrollIndicator={false}
            refreshControl={
              <RefreshControl
                refreshing={refreshing}
                onRefresh={onRefresh}
                colors={['#ff0000']}
                tintColor="#ff0000"
              />
            }
          >
            {screenData?.ui.map((snippet, index) => (
              <UIRenderer 
                key={index} 
                snippet={snippet} 
                data={screenData.data}
                onAction={handleAction}
              />
            ))}
            
            {/* Debug info in development */}
            {__DEV__ && screenData && (
              <View style={styles.debugPanel}>
                <Text style={styles.debugText}>
                  Screen: {screenData.screen} | 
                  Payments: {screenData.data?.payments?.length || 0}
                </Text>
              </View>
            )}
          </ScrollView>
        </RefreshControl>
      </SafeAreaView>

      {/* Payment Detail Modal */}
      <Modal
        animationType="slide"
        transparent={true}
        visible={showModal}
        onRequestClose={() => {
          setShowModal(false);
          setModalData(null);
        }}
      >
        <View style={styles.modalOverlay}>
          <View style={styles.modalContainer}>
            {modalData?.ui.map((snippet, index) => (
              <UIRenderer 
                key={index} 
                snippet={snippet} 
                data={modalData.data}
                onAction={(action) => {
                  if (action.type === 'closeModal') {
                    setShowModal(false);
                    setModalData(null);
                  } else {
                    handleAction(action);
                  }
                }}
              />
            ))}
          </View>
        </View>
      </Modal>
    </>
  );
};

// Styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#FFFFFF',
  },
  scrollContainer: {
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
  errorTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#D32F2F',
    marginTop: 16,
    marginBottom: 8,
  },
  errorMessage: {
    fontSize: 16,
    color: '#666',
    textAlign: 'center',
    marginBottom: 24,
    lineHeight: 22,
  },
  retryButton: {
    backgroundColor: '#ff0000',
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 32,
    paddingVertical: 12,
    borderRadius: 8,
    gap: 8,
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
  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'flex-end',
  },
  modalContainer: {
    backgroundColor: '#fff',
    borderTopLeftRadius: 24,
    borderTopRightRadius: 24,
    maxHeight: Dimensions.get('window').height * 0.9,
  },
});

export default Money;