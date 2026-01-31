// home.tsx - Broker Home Screen with Backend Integration
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

const API_URL = 'http://192.168.1.3:8080/bff/broker/home';

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
  UI?: UISnippet[];
  ui?: UISnippet[];
  message: string;
}

const BrokerHomeScreen: React.FC = () => {
  const [uiData, setUiData] = useState<UISnippet[]>([]);
  const [screenData, setScreenData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshing, setRefreshing] = useState<boolean>(false);

  const fetchHomeData = useCallback(async (forceRefresh = false) => {
    try {
      if (!forceRefresh) {
        setLoading(true);
      } else {
        setRefreshing(true);
      }

      console.log('Fetching broker home data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Broker Home response:', response.data.status);

      if (response.data.status === 'success') {
        console.log('Broker Home UI Data:', response.data.UI || response.data.ui);
        setScreenData(response.data.data);
        setUiData(response.data.UI || response.data.ui || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch broker home data');
      }
    } catch (err: any) {
      console.error('Error fetching broker home data:', err.message || err);
      setError(err.message || 'Network error. Please check your connection.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchHomeData();
  }, [fetchHomeData]);

  const handleAction = async (action: ActionData) => {
    console.log('Broker action triggered:', action);
    
    try {
      switch (action.type) {
        case 'navigate':
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
      'package-variant': 'cube-outline',
      'gavel': 'hammer-outline',
      'map-marker-path': 'map-outline',
      'cash': 'cash-outline',
      'package-variant-plus': 'add-circle-outline',
      'truck-plus': 'add-circle-outline',
      'truck': 'car-outline',
      'wallet-outline': 'wallet-outline',
      'file-document': 'document-text-outline',
      'analytics-outline': 'stats-chart-outline',
      'hammer': 'hammer-outline',
      'chat': 'chatbubble-outline',
      'notifications': 'notifications-outline',
      'settings': 'settings-outline',
      'help-circle': 'help-circle-outline',
      'person': 'person-outline',
      'refresh': 'refresh-outline',
    };
    return iconMap[iconName] || 'help-circle-outline';
  };

  // Renderer for UI snippets
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    console.log(`Rendering snippet type: ${snippet.type}`, snippet.data);
    
    switch (snippet.type) {
      case 'View':
        return (
          <View key={key} style={[styles.view, parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'Text':
        return (
          <Text key={key} style={[styles.baseText, parseTextStyles(snippet.data)]}>
            {snippet.data?.text || ''}
          </Text>
        );
        
      case 'Icon':
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
        
      case 'TouchableOpacity':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.touchableOpacity, parseStyles(snippet.data?.style)]}
            onPress={() => snippet.data?.onPress && handleAction(snippet.data.onPress)}
            activeOpacity={0.7}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </TouchableOpacity>
        );
        
      case 'ScrollView':
        return (
          <ScrollView
            key={key}
            style={[styles.scrollView, parseStyles(snippet.data)]}
            horizontal={snippet.data?.horizontal}
            showsHorizontalScrollIndicator={false}
            showsVerticalScrollIndicator={false}
            refreshControl={
              !snippet.data?.horizontal ? (
                <RefreshControl
                  refreshing={refreshing}
                  onRefresh={() => fetchHomeData(true)}
                  colors={['#ff0000']}
                  tintColor="#ff0000"
                />
              ) : undefined
            }
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
      case 'StatusBar':
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
        <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
        
        {/* Header */}
        <View style={styles.fallbackHeader}>
          <View style={styles.fallbackHeaderContent}>
            <Text style={styles.fallbackAppName}>LogiBrokerrrr</Text>
            <Text style={styles.fallbackWelcome}>Welcome back! 👋</Text>
          </View>
        </View>
        
        <ScrollView
          style={styles.fallbackScrollView}
          showsVerticalScrollIndicator={false}
          refreshControl={
            <RefreshControl
              refreshing={refreshing}
              onRefresh={() => fetchHomeData(true)}
              colors={['#ff0000']}
              tintColor="#ff0000"
            />
          }
          contentContainerStyle={styles.fallbackScrollContent}
        >
          {/* Stats Section */}
          <View style={styles.fallbackSection}>
            <ScrollView 
              horizontal 
              showsHorizontalScrollIndicator={false}
              contentContainerStyle={styles.fallbackStatsContainer}
            >
              {/* Active Loads Stat */}
              <TouchableOpacity 
                style={styles.fallbackStatCard}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /loads')}
              >
                <View style={[styles.fallbackStatCardInner, { backgroundColor: '#ff0000' }]}>
                  <Ionicons name="cube-outline" size={24} color="#fff" />
                  <Text style={styles.fallbackStatNumber}>12</Text>
                  <Text style={styles.fallbackStatLabel}>Active Loads</Text>
                </View>
              </TouchableOpacity>
              
              {/* Pending Bids Stat */}
              <TouchableOpacity 
                style={[styles.fallbackStatCard, { backgroundColor: '#FFFFFF', borderWidth: 1, borderColor: '#F0F0F0' }]}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /bids')}
              >
                <View style={[styles.fallbackStatCardInner, { backgroundColor: '#ff0000' }]}>
                  <Ionicons name="hammer-outline" size={24} color="#fff" />
                  <Text style={styles.fallbackStatNumber}>8</Text>
                  <Text style={styles.fallbackStatLabel}>Pending Bids</Text>
                  <View style={styles.fallbackBadge}>
                    <Text style={styles.fallbackBadgeText}>+2 new</Text>
                  </View>
                </View>
              </TouchableOpacity>
              
              {/* Live Trips Stat */}
              <TouchableOpacity 
                style={[styles.fallbackStatCard, { backgroundColor: '#FFFFFF', borderWidth: 1, borderColor: '#F0F0F0' }]}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /trips')}
              >
                <View style={[styles.fallbackStatCardInner, { backgroundColor: '#ff0000' }]}>
                  <Ionicons name="map-outline" size={24} color="#fff" />
                  <Text style={styles.fallbackStatNumber}>5</Text>
                  <Text style={styles.fallbackStatLabel}>Live Trips</Text>
                </View>
              </TouchableOpacity>
              
              {/* Pending Payments Stat */}
              <TouchableOpacity 
                style={[styles.fallbackStatCard, { backgroundColor: '#FFFFFF', borderWidth: 1, borderColor: '#F0F0F0' }]}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /payments')}
              >
                <View style={[styles.fallbackStatCardInner, { backgroundColor: '#DC2616' }]}>
                  <Ionicons name="cash-outline" size={24} color="#fff" />
                  <Text style={styles.fallbackStatNumber}>3</Text>
                  <Text style={styles.fallbackStatLabel}>Pending Payments</Text>
                  <Text style={styles.fallbackPaymentAmount}>₹78,300</Text>
                </View>
              </TouchableOpacity>
            </ScrollView>
          </View>

          {/* Quick Actions Section */}
          <View style={styles.fallbackSection}>
            <Text style={styles.fallbackSectionTitle}>Quick Actions</Text>
            <View style={styles.fallbackActionsRow}>
              <TouchableOpacity 
                style={styles.fallbackActionButtonPrimary}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /add-load')}
              >
                <View style={styles.fallbackActionButtonContent}>
                  <Ionicons name="add-circle-outline" size={28} color="#fff" />
                  <Text style={styles.fallbackActionButtonText}>Add Load</Text>
                </View>
              </TouchableOpacity>
              
              <TouchableOpacity 
                style={styles.fallbackActionButtonSecondary}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /add-truck')}
              >
                <Ionicons name="add-circle-outline" size={28} color="#ff0000" />
                <Text style={styles.fallbackActionButtonTextSecondary}>Add Truck</Text>
              </TouchableOpacity>
            </View>
          </View>

          {/* Quick Tools Section */}
          <View style={styles.fallbackSection}>
            <Text style={styles.fallbackSectionTitle}>Quick Tools</Text>
            <View style={styles.fallbackToolsGrid}>
              {/* My Trucks */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /my-trucks')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#ff0000' }]}>
                  <Ionicons name="car-outline" size={20} color="#fff" />
                </View>
                <Text style={styles.fallbackToolText}>My Trucks</Text>
              </TouchableOpacity>
              
              {/* Create Load */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /create-load')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#F8F9FA', borderWidth: 1, borderColor: '#F0F0F0' }]}>
                  <Ionicons name="add-circle-outline" size={20} color="#ff0000" />
                </View>
                <Text style={styles.fallbackToolText}>Create Load</Text>
              </TouchableOpacity>
              
              {/* Wallet */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /wallet')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#F8F9FA', borderWidth: 1, borderColor: '#F0F0F0' }]}>
                  <Ionicons name="wallet-outline" size={20} color="#ff0000" />
                </View>
                <Text style={styles.fallbackToolText}>Wallet</Text>
              </TouchableOpacity>
              
              {/* Documents */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /documents')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#F8F9FA', borderWidth: 1, borderColor: '#F0F0F0' }]}>
                  <Ionicons name="document-text-outline" size={20} color="#ff0000" />
                </View>
                <Text style={styles.fallbackToolText}>Documents</Text>
              </TouchableOpacity>
              
              {/* Analytics */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /analytics')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#F8F9FA', borderWidth: 1, borderColor: '#F0F0F0' }]}>
                  <Ionicons name="stats-chart-outline" size={20} color="#ff0000" />
                </View>
                <Text style={styles.fallbackToolText}>Analytics</Text>
              </TouchableOpacity>
              
              {/* Place Bid */}
              <TouchableOpacity 
                style={styles.fallbackToolItem}
                onPress={() => Alert.alert('Navigation', 'Would navigate to /place-bid')}
              >
                <View style={[styles.fallbackToolIcon, { backgroundColor: '#F8F9FA', borderWidth: 1, borderColor: '#F0F0F0' }]}>
                  <Ionicons name="hammer-outline" size={20} color="#ff0000" />
                </View>
                <Text style={styles.fallbackToolText}>Place Bid</Text>
              </TouchableOpacity>
            </View>
          </View>

          {/* Error Message */}
          {error && (
            <View style={styles.fallbackErrorContainer}>
              <Ionicons name="warning-outline" size={48} color="#ff0000" />
              <Text style={styles.fallbackErrorTitle}>Oops! Something went wrong</Text>
              <Text style={styles.fallbackErrorMessage}>
                {error || 'Unable to load broker dashboard. Please check your connection and try again.'}
              </Text>
              
              <TouchableOpacity 
                style={styles.fallbackRetryButton}
                onPress={() => fetchHomeData(true)}
              >
                <Ionicons name="refresh-outline" size={20} color="#fff" />
                <Text style={styles.fallbackRetryButtonText}>Try Again</Text>
              </TouchableOpacity>
            </View>
          )}

          <View style={styles.fallbackSpacer} />
        </ScrollView>
      </SafeAreaView>
    );
  };

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
        <View style={styles.loadingContent}>
          <ActivityIndicator size="large" color="#ff0000" />
          <Text style={styles.loadingText}>Loading broker dashboard...</Text>
        </View>
      </SafeAreaView>
    );
  }

  if (error || uiData.length === 0) {
    console.log('Showing fallback UI due to:', error || 'No UI data');
    return renderFallbackUI();
  }

  console.log('Rendering backend UI with', uiData.length, 'snippets');
  
  // Render UI from backend
  return (
    <SafeAreaView style={styles.container}>
      <StatusBar backgroundColor="#FFFFFF" barStyle="dark-content" />
      {uiData.map((snippet, index) => renderSnippet(snippet, index))}
    </SafeAreaView>
  );
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
  
  // Position properties
  if (data.position) style.position = data.position;
  if (data.top) style.top = data.top;
  if (data.left) style.left = data.left;
  if (data.right) style.right = data.right;
  if (data.bottom) style.bottom = data.bottom;
  
  // Overflow
  if (data.overflow) style.overflow = data.overflow;
  
  // Shadow/elevation
  if (data.elevation !== undefined) {
    if (Platform.OS === 'android') {
      style.elevation = data.elevation;
    } else if (data.shadowColor || data.ShadowColor) {
      style.shadowColor = data.shadowColor || data.ShadowColor || '#000';
      style.shadowOffset = {
        width: data.shadowOffsetX || data.ShadowOffsetX || 0,
        height: data.shadowOffsetY || data.ShadowOffsetY || 2,
      };
      style.shadowOpacity = data.shadowOpacity || data.ShadowOpacity || 0.1;
      style.shadowRadius = data.shadowRadius || data.ShadowRadius || 8;
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
    backgroundColor: '#FFFFFF',
  },
  scrollView: {
    flex: 1,
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
  touchableOpacity: {
    // Default touchable opacity
  },
  loadingContainer: {
    flex: 1,
    backgroundColor: '#FFFFFF',
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
    backgroundColor: '#FFFFFF',
  },
  fallbackHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 20,
    paddingVertical: 16,
    backgroundColor: '#FFFFFF',
    borderBottomWidth: 1,
    borderBottomColor: '#F0F0F0',
  },
  fallbackHeaderContent: {
    flex: 1,
  },
  fallbackAppName: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#ff0000',
    marginBottom: 4,
  },
  fallbackWelcome: {
    fontSize: 14,
    color: '#666',
  },
  fallbackScrollView: {
    flex: 1,
  },
  fallbackScrollContent: {
    paddingBottom: 20,
  },
  fallbackSection: {
    paddingHorizontal: 20,
    paddingVertical: 16,
  },
  fallbackSectionTitle: {
    fontSize: 22,
    fontWeight: 'bold',
    color: '#1A1A1A',
    marginBottom: 16,
  },
  fallbackStatsContainer: {
    flexDirection: 'row',
    paddingBottom: 8,
  },
  fallbackStatCard: {
    width: 160,
    height: 120,
    borderRadius: 16,
    marginRight: 12,
    overflow: 'hidden',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 8,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  fallbackStatCardInner: {
    flex: 1,
    padding: 16,
    alignItems: 'center',
    justifyContent: 'center',
  },
  fallbackStatNumber: {
    fontSize: 28,
    fontWeight: 'bold',
    color: '#FFFFFF',
    marginTop: 8,
    marginBottom: 4,
  },
  fallbackStatLabel: {
    fontSize: 12,
    color: '#FFFFFF',
    textAlign: 'center',
    fontWeight: '500',
  },
  fallbackBadge: {
    position: 'absolute',
    top: 12,
    right: 12,
    backgroundColor: '#FFF0F0',
    paddingHorizontal: 8,
    paddingVertical: 2,
    borderRadius: 10,
  },
  fallbackBadgeText: {
    fontSize: 10,
    color: '#ff0000',
    fontWeight: '600',
  },
  fallbackPaymentAmount: {
    fontSize: 12,
    color: '#ffffff',
    fontWeight: 'bold',
    marginTop: 4,
  },
  fallbackActionsRow: {
    flexDirection: 'row',
    gap: 12,
  },
  fallbackActionButtonPrimary: {
    flex: 1,
    borderRadius: 16,
    overflow: 'hidden',
    ...Platform.select({
      ios: {
        shadowColor: '#ff0000',
        shadowOffset: { width: 0, height: 4 },
        shadowOpacity: 0.3,
        shadowRadius: 8,
      },
      android: {
        elevation: 4,
      },
    }),
  },
  fallbackActionButtonContent: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 18,
    borderRadius: 16,
    gap: 8,
    backgroundColor: '#ff0000',
  },
  fallbackActionButtonText: {
    fontSize: 16,
    fontWeight: '600',
    color: '#FFFFFF',
  },
  fallbackActionButtonSecondary: {
    flex: 1,
    backgroundColor: '#FFFFFF',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 18,
    borderRadius: 16,
    gap: 8,
    borderWidth: 2,
    borderColor: '#ff0000',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  fallbackActionButtonTextSecondary: {
    fontSize: 16,
    fontWeight: '600',
    color: '#ff0000',
  },
  fallbackToolsGrid: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-between',
    gap: 12,
  },
  fallbackToolItem: {
    width: '23%',
    alignItems: 'center',
    marginBottom: 16,
  },
  fallbackToolIcon: {
    width: 60,
    height: 60,
    borderRadius: 16,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 8,
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.1,
        shadowRadius: 2,
      },
      android: {
        elevation: 1,
      },
    }),
  },
  fallbackToolText: {
    fontSize: 12,
    color: '#333',
    textAlign: 'center',
    fontWeight: '500',
  },
  fallbackErrorContainer: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
    padding: 24,
    marginHorizontal: 20,
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
  fallbackErrorTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#1A1A1A',
    marginTop: 16,
    marginBottom: 8,
    textAlign: 'center',
  },
  fallbackErrorMessage: {
    fontSize: 14,
    color: '#666',
    textAlign: 'center',
    marginBottom: 24,
    lineHeight: 20,
  },
  fallbackRetryButton: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#ff0000',
    paddingHorizontal: 24,
    paddingVertical: 12,
    borderRadius: 12,
    gap: 8,
  },
  fallbackRetryButtonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: '600',
  },
  fallbackSpacer: {
    height: 20,
  },
});

export default BrokerHomeScreen;