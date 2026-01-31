import React, { useState, useEffect } from 'react';
import {
  View,
  Text,
  ScrollView,
  TouchableOpacity,
  StyleSheet,
  ActivityIndicator,
  RefreshControl,
  Modal,
  Dimensions,
  Alert,
  FlatList,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
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

// Icon name mapping from backend to Ionicons
const iconMap: Record<string, string> = {
  'search': 'search',
  'filter': 'filter',
  'chevron-right': 'chevron-forward',
  'close': 'close',
  'truck': 'truck-outline',
  'map-marker': 'location-outline',
  'person': 'person-outline',
  'call': 'call-outline',
  'chatbubble': 'chatbubble-outline',
  'location-on': 'location-outline',
  'refresh': 'refresh',
  'calendar-month': 'calendar-outline',
  'clock-outline': 'time-outline',
  'check-circle-outline': 'checkmark-circle-outline',
  'plus': 'add',
  'file-download': 'download-outline',
  'credit-card': 'card-outline',
  'cash-multiple': 'cash-outline',
  'alert-circle-outline': 'alert-circle-outline',
};

// Main Load Component
const Load: React.FC = () => {
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [screenData, setScreenData] = useState<ScreenResponse | null>(null);
  const [modalData, setModalData] = useState<ScreenResponse | null>(null);
  const [showModal, setShowModal] = useState(false);
  const [activeTab, setActiveTab] = useState('Active Loads');
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
      if (snippetData.fontWeight !== undefined) {
        textStyle.fontWeight = snippetData.fontWeight === 'bold' ? 'bold' : snippetData.fontWeight;
      }
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

    // Handle FlatList component
    if (type === 'FlatList') {
      // Since we can't dynamically render FlatList from data, we'll use a View with ScrollView
      return (
        <ScrollView style={parseStyle(snippetData)}>
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
          name={iconName as any}
          size={snippetData.size || 24}
          color={snippetData.color || '#333'}
          style={iconStyle}
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

  // Create fallback UI for when API fails
  const createFallbackUI = () => ({
    status: 'success',
    screen: 'load',
    ui: [{
      type: 'View',
      data: { flex: 1, backgroundColor: '#FFFFFF' },
      children: [
        // Header
        {
          type: 'View',
          data: { 
            flexDirection: 'row',
            justifyContent: 'space-between',
            alignItems: 'center',
            paddingHorizontal: 20,
            paddingTop: 10,
            paddingBottom: 20,
            backgroundColor: '#fff',
            borderBottomWidth: 1,
            borderColor: '#f0f0f0'
          },
          children: [
            {
              type: 'View',
              data: { flex: 1 },
              children: [
                {
                  type: 'Text',
                  data: {
                    text: 'Loads',
                    fontSize: 28,
                    fontWeight: 'bold',
                    color: '#1a1a1a',
                    marginBottom: 4,
                  }
                },
                {
                  type: 'View',
                  data: {
                    width: 60,
                    height: 4,
                    backgroundColor: '#ff0000',
                    borderRadius: 2,
                  }
                }
              ]
            },
            {
              type: 'View',
              data: { flexDirection: 'row', alignItems: 'center' },
              children: [
                {
                  type: 'TouchableOpacity',
                  data: {
                    style: { marginLeft: 16 },
                    onPress: { type: 'search' }
                  },
                  children: [{
                    type: 'Icon',
                    data: { name: 'search', size: 24, color: '#1a1a1a' }
                  }]
                },
                {
                  type: 'TouchableOpacity',
                  data: {
                    style: { marginLeft: 16 },
                    onPress: { type: 'filter' }
                  },
                  children: [{
                    type: 'Icon',
                    data: { name: 'filter', size: 24, color: '#1a1a1a' }
                  }]
                }
              ]
            }
          ]
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
                name: 'truck-outline',
                size: 64,
                color: '#ccc',
                style: { marginBottom: 20 }
              }
            },
            {
              type: 'Text',
              data: {
                text: 'Unable to load loads data',
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
    data: {
      tabs: ['Active Loads', 'Pending Loads', 'Completed Loads'],
      activeTab: 'Active Loads',
      loads: {}
    }
  });

  // Create custom tab renderer since we need to track active tab state
  const renderCustomTabs = () => {
    const tabs = screenData?.data?.tabs || ['Active Loads', 'Pending Loads', 'Completed Loads'];
    
    return (
      <View style={styles.customTabsContainer}>
        <View style={styles.tabsInnerContainer}>
          {tabs.map((tab: string) => {
            const isActive = tab === activeTab;
            return (
              <TouchableOpacity
                key={tab}
                style={[
                  styles.customTab,
                  isActive && styles.customTabActive
                ]}
                onPress={() => {
                  setActiveTab(tab);
                  handleAction({
                    type: 'updateScreen',
                    data: { activeTab: tab }
                  });
                }}
              >
                <Text style={[
                  styles.customTabText,
                  isActive && styles.customTabTextActive
                ]}>
                  {tab}
                </Text>
              </TouchableOpacity>
            );
          })}
        </View>
      </View>
    );
  };

  // Create custom load cards renderer for active tab
  const renderLoadCards = () => {
    const loadsData = screenData?.data?.loads || {};
    const currentLoads = loadsData[activeTab] || [];
    
    if (currentLoads.length === 0) {
      return (
        <View style={styles.noLoadsContainer}>
          <Ionicons name="truck-outline" size={48} color="#ccc" />
          <Text style={styles.noLoadsText}>No loads found</Text>
          <Text style={styles.noLoadsSubtext}>Try switching tabs or check back later</Text>
        </View>
      );
    }

    return (
      <View style={styles.loadsListContainer}>
        {currentLoads.map((load: any, index: number) => {
          const statusStyle = screenData?.data?.statusStyles?.[load.status] || {
            bgColor: '#D4EDDA',
            color: '#1a1a1a'
          };

          return (
            <TouchableOpacity
              key={load.id || index}
              style={styles.loadCard}
              onPress={() => handleAction({
                type: 'showModal',
                data: { modal: 'loadDetail', loadId: load.id }
              })}
              activeOpacity={0.7}
            >
              {/* Card Header */}
              <View style={styles.cardHeader}>
                <View style={styles.cardHeaderLeft}>
                  <Text style={styles.routeText}>{load.pickup} → {load.drop}</Text>
                  <Text style={styles.distanceText}>{load.distance}</Text>
                </View>
                <View style={styles.bidBadge}>
                  <Text style={styles.bidText}>{load.bids} Bids</Text>
                </View>
              </View>

              {/* Card Details */}
              <View style={styles.cardDetails}>
                <View style={styles.detailRow}>
                  <Text style={styles.detailLabel}>Cargo:</Text>
                  <Text style={styles.detailValue}>{load.cargoType}</Text>
                </View>
                <View style={styles.detailRow}>
                  <Text style={styles.detailLabel}>Vehicle:</Text>
                  <Text style={styles.detailValue}>{load.vehicleType}</Text>
                </View>
                <View style={styles.detailRow}>
                  <Text style={styles.detailLabel}>Budget:</Text>
                  <Text style={styles.budgetText}>{load.budget}</Text>
                </View>
              </View>

              {/* Card Footer */}
              <View style={styles.cardFooter}>
                <View style={[styles.statusBadge, { backgroundColor: statusStyle.bgColor }]}>
                  <Text style={[styles.statusText, { color: statusStyle.color }]}>
                    {load.status}
                  </Text>
                </View>
                <View style={styles.viewDetails}>
                  <Text style={styles.viewDetailsText}>View Details</Text>
                  <Ionicons name="chevron-forward" size={20} color="#ff0000" />
                </View>
              </View>
            </TouchableOpacity>
          );
        })}
      </View>
    );
  };

  // Fetch main load data
  const fetchLoadData = async (isRefreshing = false) => {
    try {
      if (!isRefreshing) {
        setLoading(true);
      }
      setError(null);

      // Replace with your actual API endpoint
      const response = await fetch('http://192.168.1.3:8080/bff/broker/load', {
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
        setActiveTab(data.data?.activeTab || 'Active Loads');
      } else {
        throw new Error('Failed to load loads data');
      }
    } catch (err) {
      console.error('Error fetching load data:', err);
      setError(err instanceof Error ? err.message : 'Network request failed');
      
      // Fallback UI
      if (!screenData) {
        setScreenData(createFallbackUI() as ScreenResponse);
      }
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  };

  // Fetch load detail modal data
  const fetchLoadDetail = async (loadId: string) => {
    try {
      const response = await fetch(`http://your-backend-url/broker/load-detail?loadId=${loadId}`, {
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
        throw new Error('Failed to load load details');
      }
    } catch (err) {
      console.error('Error fetching load details:', err);
      Alert.alert('Error', 'Failed to load load details');
    }
  };

  // Handle all actions from UI
  const handleAction = (action: any) => {
    console.log('Action triggered:', action);
    
    switch (action.type) {
      case 'retry':
        fetchLoadData();
        break;
        
      case 'search':
        Alert.alert('Search', 'Search functionality would open here');
        break;
        
      case 'filter':
        Alert.alert('Filter', 'Filter options would appear here');
        break;
        
      case 'updateScreen':
        console.log('Update screen with:', action.data);
        if (action.data.activeTab) {
          setActiveTab(action.data.activeTab);
        }
        break;
      
      case 'showModal':
        if (action.data && action.data.loadId) {
          fetchLoadDetail(action.data.loadId);
        }
        break;
      
      case 'closeModal':
        setShowModal(false);
        setModalData(null);
        break;
      
      case 'placeBid':
        Alert.alert('Place Bid', `Would place bid for load: ${action.data?.loadId}`);
        break;
      
      default:
        console.log('Unhandled action:', action);
    }
  };

  // Initial load
  useEffect(() => {
    fetchLoadData();
  }, []);

  // Pull to refresh
  const onRefresh = () => {
    setRefreshing(true);
    fetchLoadData(true);
  };

  // Loading state
  if (loading && !refreshing) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="#ff0000" />
        <Text style={styles.loadingText}>Loading Loads...</Text>
      </SafeAreaView>
    );
  }

  // Error state (no data loaded yet)
  if (error && !screenData) {
    return (
      <SafeAreaView style={styles.errorContainer}>
        <View style={styles.errorContent}>
          <Ionicons name="alert-circle-outline" size={64} color="#D32F2F" />
          <Text style={styles.errorTitle}>Connection Issue</Text>
          <Text style={styles.errorMessage}>
            {error || 'Unable to load loads data. Please check your connection.'}
          </Text>
          <TouchableOpacity 
            style={styles.retryButton}
            onPress={() => fetchLoadData()}
          >
            <Ionicons name="refresh" size={20} color="#fff" />
            <Text style={styles.retryButtonText}>Try Again</Text>
          </TouchableOpacity>
        </View>
      </SafeAreaView>
    );
  }

  // Main render
  return (
    <>
      <SafeAreaView style={styles.container} edges={['top']}>
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
            {/* Render header from backend UI */}
            {screenData?.ui.slice(0, 1).map((snippet, index) => (
              <UIRenderer 
                key={`header-${index}`} 
                snippet={snippet} 
                data={screenData.data}
                onAction={handleAction}
              />
            ))}
            
            {/* Custom Tabs Renderer */}
            {renderCustomTabs()}
            
            {/* Custom Load Cards Renderer */}
            {renderLoadCards()}
            
            {/* Debug info in development */}
            {__DEV__ && screenData && (
              <View style={styles.debugPanel}>
                <Text style={styles.debugText}>
                  Screen: {screenData.screen} | 
                  Active Tab: {activeTab} | 
                  Loads: {screenData.data?.loads?.[activeTab]?.length || 0}
                </Text>
              </View>
            )}
          </ScrollView>
        </RefreshControl>
      </SafeAreaView>

      {/* Load Detail Modal */}
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
  
  // Custom Tabs Styles
  customTabsContainer: {
    flexDirection: 'row',
    paddingHorizontal: 10,
    paddingVertical: 7,
    marginTop: 20,
    borderBottomWidth: 1,
    borderColor: '#f0f0f0',
  },
  tabsInnerContainer: {
    flexDirection: 'row',
    flex: 1,
  },
  customTab: {
    flex: 1,
    paddingVertical: 12,
    alignItems: 'center',
    borderRadius: 8,
    marginHorizontal: 4,
    backgroundColor: 'transparent',
  },
  customTabActive: {
    backgroundColor: '#ff0000',
  },
  customTabText: {
    fontSize: 12,
    fontWeight: '900',
    color: '#666',
  },
  customTabTextActive: {
    color: '#fff',
  },
  
  // Load Cards Styles
  loadsListContainer: {
    paddingHorizontal: 20,
    paddingVertical: 10,
  },
  loadCard: {
    backgroundColor: '#fff',
    borderRadius: 16,
    padding: 20,
    marginBottom: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.1,
    shadowRadius: 8,
    elevation: 4,
    borderWidth: 1,
    borderColor: '#f0f0f0',
  },
  cardHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: 16,
  },
  cardHeaderLeft: {
    flex: 1,
  },
  routeText: {
    fontSize: 16,
    color: '#1a1a1a',
    marginBottom: 4,
  },
  distanceText: {
    fontSize: 14,
    color: '#666',
  },
  bidBadge: {
    backgroundColor: '#ff0000',
    paddingHorizontal: 12,
    paddingVertical: 6,
    borderRadius: 12,
  },
  bidText: {
    color: '#fff',
    fontSize: 12,
    fontWeight: '600',
  },
  cardDetails: {
    marginBottom: 16,
  },
  detailRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 8,
  },
  detailLabel: {
    fontSize: 14,
    color: '#666',
  },
  detailValue: {
    fontSize: 14,
    color: '#1a1a1a',
    fontWeight: '500',
  },
  budgetText: {
    fontSize: 16,
    color: '#ff0000',
    fontWeight: 'bold',
  },
  cardFooter: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  statusBadge: {
    paddingHorizontal: 12,
    paddingVertical: 6,
    borderRadius: 12,
  },
  statusText: {
    fontSize: 12,
    fontWeight: '600',
  },
  viewDetails: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  viewDetailsText: {
    color: '#ff0000',
    fontSize: 14,
    fontWeight: '600',
    marginRight: 4,
  },
  noLoadsContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 40,
    minHeight: 300,
  },
  noLoadsText: {
    fontSize: 18,
    fontWeight: '600',
    color: '#666',
    marginTop: 16,
    marginBottom: 8,
  },
  noLoadsSubtext: {
    fontSize: 14,
    color: '#999',
    textAlign: 'center',
  },
});

export default Load;