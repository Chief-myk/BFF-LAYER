// myTrip.tsx - My Trips Screen with Backend Integration
import React, { useState, useEffect, useCallback } from 'react';
import {
  View,
  ScrollView,
  Text,
  TouchableOpacity,
  StyleSheet,
  Alert,
  ActivityIndicator,
  StatusBar,
  Platform,
  RefreshControl,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { SafeAreaView } from 'react-native-safe-area-context';
import axios from 'axios';

const API_URL = 'http://192.168.1.3:8080/bff/driver/mytrip';

// Define types based on your backend
interface ActionData {
  type: string;
  value?: string;
  url?: string;
  to?: string;
  data?: any;
}

interface UISnippet {
  type: string;
  data: any;
  children?: UISnippet[];
}

interface ScreenResponse {
  status: string;
  screen: string;
  data: any;
  ui: UISnippet[];
  message: string;
}

const MyTripScreen: React.FC = () => {
  const [uiData, setUiData] = useState<UISnippet[]>([]);
  const [screenData, setScreenData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshing, setRefreshing] = useState<boolean>(false);
  const [activeTab, setActiveTab] = useState<string>('Current');

  const fetchTripData = useCallback(async (forceRefresh = false) => {
    try {
      if (!forceRefresh) {
        setLoading(true);
      } else {
        setRefreshing(true);
      }

      console.log('Fetching trip data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Trip response:', response.data.status);

      if (response.data.status === 'success') {
        console.log('Trip UI Data length:', response.data.ui?.length);
        setScreenData(response.data.data);
        setUiData(response.data.ui || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch trip data');
      }
    } catch (err: any) {
      console.error('Error fetching trip data:', err.message || err);
      setError(err.message || 'Network error. Please check your connection.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchTripData();
  }, [fetchTripData]);

  const handleAction = async (action: ActionData) => {
    console.log('Trip action triggered:', action);
    
    try {
      switch (action.type) {
        case 'ACTION':
          if (action.value === 'switch_tab') {
            const tab = action.data?.tab;
            if (tab) {
              setActiveTab(tab);
              Alert.alert('Tab Changed', `Switched to ${tab} trips`);
            }
          } else if (action.value === 'view_trip_details') {
            const tripId = action.data?.tripId;
            Alert.alert('Trip Details', `Viewing details for trip: ${tripId || 'Unknown'}`);
          } else if (action.url) {
            const response = await axios.post(action.url, {
              action: action.value,
              data: action.data || {},
            });
            
            if (response.data.status === 'success') {
              Alert.alert('Success', response.data.message || 'Action completed');
              fetchTripData(true);
            }
          }
          break;
          
        case 'NAVIGATE':
          Alert.alert('Navigation', `Would navigate to: ${action.to}`);
          break;
          
        default:
          console.log('Unhandled action type:', action.type);
      }
    } catch (err) {
      console.error('Action error:', err);
      Alert.alert('Error', 'Failed to perform action');
    }
  };

  // Helper function to render icon based on name
  const getIconName = (iconName: string): any => {
    const iconMap: Record<string, string> = {
      'calendar-outline': 'calendar-outline',
      'cash-outline': 'cash-outline',
      'chevron-forward': 'chevron-forward',
      'map-outline': 'map-outline',
      'time-outline': 'time-outline',
      'checkmark-circle': 'checkmark-circle',
      'car-outline': 'car-outline',
      'location-outline': 'location-outline',
      'flag-outline': 'flag-outline',
    };
    return iconMap[iconName] || 'help-circle-outline';
  };

  // Renderer for UI snippets
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    console.log(`Rendering snippet type: ${snippet.type}`, snippet.data);
    
    switch (snippet.type) {
      case 'SAFE_AREA':
        return (
          <SafeAreaView key={key} style={[parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </SafeAreaView>
        );
        
      case 'ROW':
        return (
          <View key={key} style={[styles.row, parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'VIEW':
        return (
          <View key={key} style={[styles.view, parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'TEXT':
        return (
          <Text key={key} style={[styles.baseText, parseTextStyles(snippet.data)]}>
            {snippet.data?.text || ''}
          </Text>
        );
        
      case 'ICON':
        if (!snippet.data?.name) return null;
        return (
          <View key={key} style={[styles.iconContainer, parseStyles(snippet.data)]}>
            <Ionicons
              name={getIconName(snippet.data.name)}
              size={snippet.data.size || 20}
              color={snippet.data.color || '#000'}
            />
          </View>
        );
        
      case 'PRESSABLE_CARD':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.pressableCard, parseStyles(snippet.data)]}
            onPress={() => {
              // Handle tab switching
              const tabText = snippet.children?.find(child => child.type === 'TEXT')?.data?.text;
              if (tabText) {
                handleAction({
                  type: 'ACTION',
                  value: 'switch_tab',
                  data: { tab: tabText }
                });
              }
            }}
            activeOpacity={0.7}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </TouchableOpacity>
        );
        
      case 'SCROLL':
        return (
          <ScrollView
            key={key}
            style={[styles.scrollView, parseStyles(snippet.data)]}
            refreshControl={
              <RefreshControl
                refreshing={refreshing}
                onRefresh={() => fetchTripData(true)}
                colors={['#ff0000']}
                tintColor="#ff0000"
              />
            }
            showsVerticalScrollIndicator={false}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
      case 'STATUS_BAR':
        return null; // Handled by React Native StatusBar
        
      default:
        console.log(`Unknown snippet type: ${snippet.type}`);
        return null;
    }
  };

  // Fallback UI for when backend data fails to load
  const renderFallbackUI = () => {
    return (
      <SafeAreaView style={styles.fallbackContainer}>
        <StatusBar backgroundColor="#ff0000" barStyle="light-content" />
        <View style={styles.fallbackContent}>
          <View style={styles.fallbackHeader}>
            <Text style={styles.fallbackTitle}>My Trips</Text>
          </View>
          
          {/* Tab Navigation */}
          <View style={styles.fallbackTabs}>
            {['Current', 'Upcoming', 'Completed'].map((tab) => (
              <TouchableOpacity
                key={tab}
                style={[
                  styles.fallbackTab,
                  activeTab === tab && styles.fallbackTabActive
                ]}
                onPress={() => setActiveTab(tab)}
              >
                <Text style={[
                  styles.fallbackTabText,
                  activeTab === tab && styles.fallbackTabTextActive
                ]}>
                  {tab}
                </Text>
                {activeTab === tab && (
                  <View style={styles.fallbackTabIndicator} />
                )}
              </TouchableOpacity>
            ))}
          </View>
          
          <ScrollView
            style={styles.fallbackScrollView}
            showsVerticalScrollIndicator={false}
            contentContainerStyle={styles.fallbackScrollContent}
          >
            {/* Error Message */}
            <View style={styles.errorContainer}>
              <Ionicons name="warning-outline" size={64} color="#ff0000" />
              <Text style={styles.errorTitle}>Oops! Something went wrong</Text>
              <Text style={styles.errorMessage}>
                {error || 'Unable to load your trips. Please check your connection and try again.'}
              </Text>
              
              <TouchableOpacity 
                style={styles.retryButton}
                onPress={() => fetchTripData(true)}
              >
                <Ionicons name="refresh" size={20} color="#fff" />
                <Text style={styles.retryButtonText}>Try Again</Text>
              </TouchableOpacity>
            </View>
            
            {/* Sample Trip Cards for Reference */}
            {activeTab === 'Current' && (
              <View style={styles.sampleTripCard}>
                <View style={styles.tripCardHeader}>
                  <View>
                    <Text style={styles.tripCardRoute}>Mumbai → Delhi</Text>
                    <Text style={styles.tripCardDetails}>1,450 km • 28 hrs</Text>
                  </View>
                  <View style={[styles.statusBadge, { backgroundColor: '#3B82F615' }]}>
                    <Text style={[styles.statusText, { color: '#3B82F6' }]}>In Transit</Text>
                  </View>
                </View>
                
                <View style={styles.tripCardInfo}>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="calendar-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>15 Dec 2024, 08:00 AM</Text>
                  </View>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="cash-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>₹45,000</Text>
                  </View>
                </View>
                
                <TouchableOpacity 
                  style={styles.viewDetailsButton}
                  onPress={() => Alert.alert('Trip Details', 'Viewing trip details')}
                >
                  <Text style={styles.viewDetailsButtonText}>View Trip Details</Text>
                  <Ionicons name="chevron-forward" size={20} color="#ff0000" />
                </TouchableOpacity>
              </View>
            )}
            
            {activeTab === 'Upcoming' && (
              <View style={styles.sampleTripCard}>
                <View style={styles.tripCardHeader}>
                  <View>
                    <Text style={styles.tripCardRoute}>Chennai → Hyderabad</Text>
                    <Text style={styles.tripCardDetails}>1,200 km • 24 hrs</Text>
                  </View>
                  <View style={[styles.statusBadge, { backgroundColor: '#EA580C15' }]}>
                    <Text style={[styles.statusText, { color: '#EA580C' }]}>Assigned</Text>
                  </View>
                </View>
                
                <View style={styles.tripCardInfo}>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="calendar-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>18 Dec 2024, 09:00 AM</Text>
                  </View>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="cash-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>₹52,000</Text>
                  </View>
                </View>
                
                <TouchableOpacity 
                  style={styles.viewDetailsButton}
                  onPress={() => Alert.alert('Trip Details', 'Viewing trip details')}
                >
                  <Text style={styles.viewDetailsButtonText}>View Trip Details</Text>
                  <Ionicons name="chevron-forward" size={20} color="#ff0000" />
                </TouchableOpacity>
              </View>
            )}
            
            {activeTab === 'Completed' && (
              <View style={styles.sampleTripCard}>
                <View style={styles.tripCardHeader}>
                  <View>
                    <Text style={styles.tripCardRoute}>Ahmedabad → Mumbai</Text>
                    <Text style={styles.tripCardDetails}>530 km • 12 hrs</Text>
                  </View>
                  <View style={[styles.statusBadge, { backgroundColor: '#6B728015' }]}>
                    <Text style={[styles.statusText, { color: '#6B7280' }]}>Completed</Text>
                  </View>
                </View>
                
                <View style={styles.tripCardInfo}>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="calendar-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>12 Dec 2024, 07:00 AM</Text>
                  </View>
                  <View style={styles.tripCardInfoItem}>
                    <Ionicons name="cash-outline" size={20} color="#6B7280" />
                    <Text style={styles.tripCardInfoText}>₹35,000</Text>
                  </View>
                </View>
                
                <TouchableOpacity 
                  style={styles.viewDetailsButton}
                  onPress={() => Alert.alert('Trip Details', 'Viewing trip details')}
                >
                  <Text style={styles.viewDetailsButtonText}>View Trip Details</Text>
                  <Ionicons name="chevron-forward" size={20} color="#ff0000" />
                </TouchableOpacity>
              </View>
            )}
            
            {/* Empty State */}
            <View style={styles.emptyState}>
              <Ionicons name="car-outline" size={64} color="#e5e7eb" />
              <Text style={styles.emptyStateTitle}>No trips available</Text>
              <Text style={styles.emptyStateText}>
                You don't have any {activeTab.toLowerCase()} trips at the moment
              </Text>
            </View>
          </ScrollView>
        </View>
      </SafeAreaView>
    );
  };

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#ff0000" barStyle="light-content" />
        <View style={styles.loadingContent}>
          <ActivityIndicator size="large" color="#ff0000" />
          <Text style={styles.loadingText}>Loading your trips...</Text>
        </View>
      </SafeAreaView>
    );
  }

  if (error || uiData.length === 0) {
    console.log('Showing fallback UI due to:', error || 'No UI data');
    return renderFallbackUI();
  }

  console.log('Rendering backend UI with', uiData.length, 'snippets');
  
  // Find status bar config from UI data
  const statusBarSnippet = findStatusBarSnippet(uiData);
  const statusBarConfig = statusBarSnippet?.data || { 
    backgroundColor: '#ff0000', 
    style: 'light' 
  };

  // Render UI from backend
  return (
    <>
      <StatusBar {...statusBarConfig} />
      {uiData.map((snippet, index) => renderSnippet(snippet, index))}
    </>
  );
};

// Helper function to find STATUS_BAR snippet in nested structure
const findStatusBarSnippet = (snippets: UISnippet[]): UISnippet | null => {
  for (const snippet of snippets) {
    if (snippet.type === 'STATUS_BAR') {
      return snippet;
    }
    if (snippet.children) {
      const found = findStatusBarSnippet(snippet.children);
      if (found) return found;
    }
  }
  return null;
};

// Style parsing functions
const parseStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  // Handle camelCase and PascalCase properties
  if (data.backgroundColor || data.BackgroundColor) style.backgroundColor = data.backgroundColor || data.BackgroundColor;
  if (data.padding || data.Padding) style.padding = data.padding || data.Padding;
  if (data.paddingHorizontal || data.PaddingHorizontal) style.paddingHorizontal = data.paddingHorizontal || data.PaddingHorizontal;
  if (data.paddingVertical || data.PaddingVertical) style.paddingVertical = data.paddingVertical || data.PaddingVertical;
  if (data.borderRadius || data.BorderRadius) style.borderRadius = data.borderRadius || data.BorderRadius;
  if (data.margin || data.Margin) style.margin = data.margin || data.Margin;
  if (data.marginBottom || data.MarginBottom) style.marginBottom = data.marginBottom || data.MarginBottom;
  if (data.marginTop || data.MarginTop) style.marginTop = data.marginTop || data.MarginTop;
  if (data.marginHorizontal || data.MarginHorizontal) style.marginHorizontal = data.marginHorizontal || data.MarginHorizontal;
  if (data.height || data.Height) style.height = data.height || data.Height;
  if (data.width || data.Width) style.width = typeof data.width === 'string' && data.width.includes('%') ? data.width : data.width;
  if (data.flex || data.Flex) style.flex = data.flex || data.Flex;
  if (data.flexGrow || data.FlexGrow) style.flexGrow = data.flexGrow || data.FlexGrow;
  
  // Layout properties
  if (data.flexDirection) style.flexDirection = data.flexDirection;
  if (data.justifyContent) style.justifyContent = data.justifyContent;
  if (data.alignItems) style.alignItems = data.alignItems;
  if (data.alignSelf) style.alignSelf = data.alignSelf;
  if (data.gap) style.gap = data.gap;
  if (data.flexShrink) style.flexShrink = data.flexShrink;
  if (data.flexWrap) style.flexWrap = data.flexWrap;
  
  // Border properties
  if (data.borderWidth || data.BorderWidth) style.borderWidth = data.borderWidth || data.BorderWidth;
  if (data.borderColor || data.BorderColor) style.borderColor = data.borderColor || data.BorderColor;
  if (data.borderBottomWidth || data.BorderBottomWidth) style.borderBottomWidth = data.borderBottomWidth || data.BorderBottomWidth;
  if (data.borderTopWidth || data.BorderTopWidth) style.borderTopWidth = data.borderTopWidth || data.BorderTopWidth;
  if (data.borderTopColor || data.BorderTopColor) style.borderTopColor = data.borderTopColor || data.BorderTopColor;
  
  // Position properties
  if (data.position) style.position = data.position;
  if (data.top) style.top = data.top;
  if (data.left) style.left = data.left;
  if (data.right) style.right = data.right;
  if (data.bottom) style.bottom = data.bottom;
  
  // Shadow/elevation
  if (data.elevation !== undefined) {
    if (Platform.OS === 'android') {
      style.elevation = data.elevation;
    } else if (data.shadowColor || data.ShadowColor) {
      style.shadowColor = data.shadowColor || data.ShadowColor || '#000';
      style.shadowOffset = data.shadowOffset || data.ShadowOffset || { width: 0, height: 2 };
      style.shadowOpacity = data.shadowOpacity || data.ShadowOpacity || 0.1;
      style.shadowRadius = data.shadowRadius || data.ShadowRadius || 3;
    }
  }
  
  return style;
};

const parseTextStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  if (data.fontSize || data.FontSize) style.fontSize = data.fontSize || data.FontSize;
  if (data.color || data.Color) style.color = data.color || data.Color;
  if (data.bold || data.Bold) style.fontWeight = 'bold';
  if (data.fontWeight || data.FontWeight) {
    const weight = data.fontWeight || data.FontWeight;
    style.fontWeight = weight === 'bold' ? 'bold' : 
                      weight === '600' ? '600' : 
                      weight === '700' ? '700' : 
                      weight === 'normal' ? 'normal' : weight;
  }
  if (data.textAlign || data.TextAlign) style.textAlign = data.textAlign || data.TextAlign;
  if (data.marginLeft || data.MarginLeft) style.marginLeft = data.marginLeft || data.MarginLeft;
  if (data.marginRight || data.MarginRight) style.marginRight = data.marginRight || data.MarginRight;
  if (data.marginBottom || data.MarginBottom) style.marginBottom = data.marginBottom || data.MarginBottom;
  if (data.marginTop || data.MarginTop) style.marginTop = data.marginTop || data.MarginTop;
  if (data.paddingHorizontal || data.PaddingHorizontal) style.paddingHorizontal = data.paddingHorizontal || data.PaddingHorizontal;
  if (data.paddingVertical || data.PaddingVertical) style.paddingVertical = data.paddingVertical || data.PaddingVertical;
  
  return style;
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#F8FAFC',
  },
  scrollView: {
    flex: 1,
  },
  row: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  view: {
    // Default view style
  },
  baseText: {
    fontSize: 14,
    color: '#000',
  },
  iconContainer: {
    // Default icon container
  },
  pressableCard: {
    // Default pressable card
  },
  loadingContainer: {
    flex: 1,
    backgroundColor: '#F8FAFC',
  },
  loadingContent: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  loadingText: {
    marginTop: 16,
    fontSize: 16,
    color: '#666',
    fontWeight: '500',
  },
  // Fallback UI Styles
  fallbackContainer: {
    flex: 1,
    backgroundColor: '#F8FAFC',
  },
  fallbackContent: {
    flex: 1,
  },
  fallbackHeader: {
    backgroundColor: '#ff0000',
    paddingTop: Platform.OS === 'ios' ? 50 : 30,
    paddingBottom: 20,
    paddingHorizontal: 20,
  },
  fallbackTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#FFFFFF',
    textAlign: 'center',
  },
  fallbackTabs: {
    flexDirection: 'row',
    backgroundColor: '#FFFFFF',
    borderBottomWidth: 1,
    borderBottomColor: '#F1F5F9',
  },
  fallbackTab: {
    flex: 1,
    alignItems: 'center',
    paddingVertical: 16,
  },
  fallbackTabActive: {
    // Active tab styling
  },
  fallbackTabText: {
    fontSize: 16,
    fontWeight: '600',
    color: '#6B7280',
  },
  fallbackTabTextActive: {
    color: '#ff0000',
  },
  fallbackTabIndicator: {
    backgroundColor: '#ff0000',
    marginTop: 8,
    borderRadius: 2,
    width: '60%',
    height: 3,
  },
  fallbackScrollView: {
    flex: 1,
  },
  fallbackScrollContent: {
    padding: 16,
  },
  errorContainer: {
    backgroundColor: '#FFFFFF',
    borderRadius: 16,
    padding: 24,
    marginBottom: 16,
    alignItems: 'center',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 3,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  errorTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#1a237e',
    marginTop: 16,
    marginBottom: 8,
    textAlign: 'center',
  },
  errorMessage: {
    fontSize: 14,
    color: '#666',
    textAlign: 'center',
    marginBottom: 24,
    lineHeight: 20,
  },
  retryButton: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#ff0000',
    paddingHorizontal: 24,
    paddingVertical: 12,
    borderRadius: 12,
    gap: 8,
  },
  retryButtonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: '600',
  },
  sampleTripCard: {
    backgroundColor: '#FFFFFF',
    borderRadius: 16,
    padding: 20,
    marginBottom: 16,
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 3,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  tripCardHeader: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginBottom: 16,
  },
  tripCardRoute: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#1a237e',
    marginBottom: 4,
  },
  tripCardDetails: {
    fontSize: 14,
    color: '#666',
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
  tripCardInfo: {
    marginBottom: 16,
  },
  tripCardInfoItem: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 6,
    marginBottom: 8,
  },
  tripCardInfoText: {
    fontSize: 14,
    color: '#666',
  },
  viewDetailsButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    borderTopWidth: 1,
    paddingTop: 12,
    borderTopColor: '#F1F5F9',
    marginTop: 16,
  },
  viewDetailsButtonText: {
    fontSize: 16,
    fontWeight: '600',
    color: '#ff0000',
    marginRight: 4,
  },
  emptyState: {
    alignItems: 'center',
    padding: 40,
  },
  emptyStateTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#374151',
    marginTop: 16,
    marginBottom: 8,
  },
  emptyStateText: {
    fontSize: 14,
    color: '#6B7280',
    textAlign: 'center',
  },
});

export default MyTripScreen;