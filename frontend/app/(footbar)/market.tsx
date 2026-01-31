// market.tsx - Market Screen with Backend Integration
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
  TextInput,
  RefreshControl,
  SafeAreaView,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import axios from 'axios';

const API_URL = 'http://192.168.1.3:8080/bff/driver/market';

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
  UI: UISnippet[];
  message: string;
}

const MarketScreen: React.FC = () => {
  const [uiData, setUiData] = useState<UISnippet[]>([]);
  const [screenData, setScreenData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshing, setRefreshing] = useState<boolean>(false);
  const [searchQuery, setSearchQuery] = useState<string>('');

  const fetchMarketData = useCallback(async (forceRefresh = false) => {
    try {
      if (!forceRefresh) {
        setLoading(true);
      } else {
        setRefreshing(true);
      }

      console.log('Fetching market data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Market response:', response.data.status);

      if (response.data.status === 'success') {
        console.log('Market UI Data length:', response.data.UI?.length);
        setScreenData(response.data.data);
        setUiData(response.data.UI || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch market data');
      }
    } catch (err: any) {
      console.error('Error fetching market data:', err.message || err);
      setError(err.message || 'Network error. Please check your connection.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchMarketData();
  }, [fetchMarketData]);

  const handleAction = async (action: ActionData) => {
    console.log('Market action triggered:', action);
    
    try {
      switch (action.type) {
        case 'ACTION':
          if (action.value === 'open_filter') {
            Alert.alert('Filter', 'Filter options would open here');
          } else if (action.url) {
            const response = await axios.post(action.url, {
              action: action.value,
              data: action.data || {},
            });
            
            if (response.data.status === 'success') {
              Alert.alert('Success', response.data.message || 'Action completed');
              fetchMarketData(true);
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

  // Renderer for UI snippets
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    switch (snippet.type) {
      case 'SAFE_AREA':
        return (
          <SafeAreaView key={key} style={parseStyles(snippet.data)}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </SafeAreaView>
        );
        
      case 'VIEW':
        return (
          <View key={key} style={parseStyles(snippet.data)}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'TEXT':
        return (
          <TouchableOpacity
            key={key}
            activeOpacity={snippet.data.onPress ? 0.7 : 1}
            onPress={() => snippet.data.onPress && handleAction(snippet.data.onPress)}
          >
            <Text style={[styles.baseText, parseTextStyles(snippet.data)]}>
              {snippet.data.text || ''}
            </Text>
          </TouchableOpacity>
        );
        
      case 'ICON':
        return (
          <View key={key} style={parseStyles(snippet.data)}>
            <Ionicons
              name={snippet.data.name || 'help-circle-outline'}
              size={snippet.data.size || 24}
              color={snippet.data.color || '#000'}
            />
          </View>
        );
        
      case 'ICON_BUTTON':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.iconButton, parseStyles(snippet.data)]}
            onPress={() => snippet.data.onPress && handleAction(snippet.data.onPress)}
          >
            <Ionicons
              name={snippet.data.icon || 'help-circle-outline'}
              size={snippet.data.size || 24}
              color={snippet.data.color || '#000'}
            />
          </TouchableOpacity>
        );
        
      case 'INPUT':
        return (
          <View key={key} style={parseStyles(snippet.data.style)}>
            <TextInput
              style={styles.input}
              placeholder={snippet.data.placeholder || ''}
              placeholderTextColor="#999"
              value={searchQuery}
              onChangeText={setSearchQuery}
            />
          </View>
        );
        
      case 'SCROLL':
        return (
          <ScrollView
            key={key}
            style={parseStyles(snippet.data)}
            showsVerticalScrollIndicator={false}
            horizontal={snippet.data.flexDirection === 'row'}
            refreshControl={
              snippet.data.flexDirection !== 'row' ? (
                <RefreshControl
                  refreshing={refreshing}
                  onRefresh={() => fetchMarketData(true)}
                  colors={['#FF3B30']}
                  tintColor="#FF3B30"
                />
              ) : undefined
            }
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
      case 'TEXT_BUTTON':
        return (
          <TouchableOpacity
            key={key}
            style={styles.textButton}
            onPress={() => snippet.data.onPress && handleAction(snippet.data.onPress)}
          >
            <Text style={[
              styles.textButtonText,
              { color: snippet.data.color || '#FF3B30' }
            ]}>
              {snippet.data.text || ''}
            </Text>
          </TouchableOpacity>
        );
        
      case 'STATUS_BAR':
        return null; // Handled separately
        
      default:
        console.log(`Unknown snippet type: ${snippet.type}`);
        return null;
    }
  };

  // Fallback UI for when backend data fails to load
  const renderFallbackUI = () => {
    return (
      <SafeAreaView style={styles.fallbackContainer}>
        <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
        <ScrollView
          style={styles.container}
          refreshControl={
            <RefreshControl
              refreshing={refreshing}
              onRefresh={() => fetchMarketData(true)}
              colors={['#FF3B30']}
              tintColor="#FF3B30"
            />
          }
        >
          {/* Header */}
          <View style={styles.fallbackHeader}>
            <View style={styles.fallbackSearchContainer}>
              <Ionicons name="search-outline" size={20} color="#666" />
              <TextInput
                style={styles.fallbackSearchInput}
                placeholder="Search by city or route"
                placeholderTextColor="#999"
                value={searchQuery}
                onChangeText={setSearchQuery}
              />
            </View>
            <TouchableOpacity style={styles.fallbackFilterButton}>
              <Ionicons name="options-outline" size={24} color="#000" />
            </TouchableOpacity>
          </View>

          {/* Tabs */}
          <ScrollView horizontal showsHorizontalScrollIndicator={false} style={styles.fallbackTabs}>
            {['All Loads', 'Recommended', 'Nearby', 'High Paying', 'Urgent', 'My Bids'].map((tab, index) => (
              <TouchableOpacity key={index} style={styles.fallbackTab}>
                <Text style={styles.fallbackTabText}>{tab}</Text>
              </TouchableOpacity>
            ))}
          </ScrollView>

          {/* Sample Load Cards */}
          <View style={styles.fallbackContent}>
            <Text style={styles.fallbackTitle}>Available Loads</Text>
            
            {/* Load Card 1 */}
            <View style={styles.fallbackCard}>
              <View style={styles.cardHeader}>
                <View style={styles.matchBadge}>
                  <Text style={styles.matchText}>95% Match</Text>
                </View>
                <Text style={styles.timeLeft}>4 hrs left</Text>
              </View>
              
              <Text style={styles.route}>Mumbai → Delhi</Text>
              <Text style={styles.price}>₹28,000 – ₹32,000</Text>
              
              <Text style={styles.detail}>2024-09-20 • 10:00 AM</Text>
              <Text style={styles.detail}>Container Truck</Text>
              <Text style={styles.detail}>Electronics • 8 Tons</Text>
              <Text style={styles.detail}>1,416 km</Text>
              
              <TouchableOpacity style={styles.viewButton}>
                <Text style={styles.viewButtonText}>View Load Details</Text>
              </TouchableOpacity>
            </View>

            {/* Load Card 2 */}
            <View style={styles.fallbackCard}>
              <View style={styles.cardHeader}>
                <View style={styles.matchBadge}>
                  <Text style={styles.matchText}>88% Match</Text>
                </View>
                <Text style={styles.timeLeft}>2 hrs left</Text>
              </View>
              
              <Text style={styles.route}>Chennai → Bangalore</Text>
              <Text style={styles.price}>₹35,000 – ₹40,000</Text>
              
              <Text style={styles.detail}>2024-09-19 • 02:00 PM</Text>
              <Text style={styles.detail}>Trailer Truck</Text>
              <Text style={styles.detail}>FMCG Goods • 12 Tons</Text>
              <Text style={styles.detail}>346 km</Text>
              
              <TouchableOpacity style={styles.viewButton}>
                <Text style={styles.viewButtonText}>View Load Details</Text>
              </TouchableOpacity>
            </View>

            {/* No Results */}
            {searchQuery.trim() !== '' && (
              <View style={styles.noResults}>
                <Ionicons name="search-outline" size={48} color="#ccc" />
                <Text style={styles.noResultsText}>No loads found for "{searchQuery}"</Text>
                <Text style={styles.noResultsSubtext}>Try a different search term</Text>
              </View>
            )}
          </View>
        </ScrollView>

        <TouchableOpacity style={styles.refreshButton} onPress={() => fetchMarketData(true)}>
          <Ionicons name="refresh" size={20} color="#fff" />
          <Text style={styles.refreshButtonText}>Retry Backend</Text>
        </TouchableOpacity>
      </SafeAreaView>
    );
  };

  // Find status bar config
  const statusBarSnippet = uiData.find(s => s.type === 'STATUS_BAR');
  const statusBarConfig = statusBarSnippet?.data || { 
    backgroundColor: '#ffffff', 
    style: 'dark' 
  };

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar {...statusBarConfig} />
        <ActivityIndicator size="large" color="#FF3B30" />
        <Text style={styles.loadingText}>Loading market...</Text>
      </SafeAreaView>
    );
  }

  if (error || uiData.length === 0) {
    return renderFallbackUI();
  }

  // Render UI from backend
  return (
    <>
      <StatusBar {...statusBarConfig} />
      {uiData.map((snippet, index) => renderSnippet(snippet, index))}
    </>
  );
};

// Style parsing functions
const parseStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  // Map basic style properties
  if (data.backgroundColor) style.backgroundColor = data.backgroundColor;
  if (data.padding) style.padding = data.padding;
  if (data.paddingHorizontal) style.paddingHorizontal = data.paddingHorizontal;
  if (data.paddingVertical) style.paddingVertical = data.paddingVertical;
  if (data.margin) style.margin = data.margin;
  if (data.marginBottom) style.marginBottom = data.marginBottom;
  if (data.borderRadius) style.borderRadius = data.borderRadius;
  if (data.borderWidth) style.borderWidth = data.borderWidth;
  if (data.borderColor) style.borderColor = data.borderColor;
  if (data.borderBottomWidth) style.borderBottomWidth = data.borderBottomWidth;
  if (data.width) style.width = data.width;
  if (data.height) style.height = data.height;
  if (data.flex) style.flex = data.flex;
  if (data.flexGrow) style.flexGrow = data.flexGrow;
  
  // Layout properties
  if (data.flexDirection) style.flexDirection = data.flexDirection;
  if (data.justifyContent) style.justifyContent = data.justifyContent;
  if (data.alignItems) style.alignItems = data.alignItems;
  if (data.gap) style.gap = data.gap;
  
  // Shadow/elevation
  if (data.elevation !== undefined) {
    if (Platform.OS === 'android') {
      style.elevation = data.elevation;
    } else if (data.shadowColor) {
      style.shadowColor = data.shadowColor;
      style.shadowOffset = data.shadowOffset || { width: 0, height: 2 };
      style.shadowOpacity = data.shadowOpacity || 0.1;
      style.shadowRadius = data.shadowRadius || 4;
    }
  }
  
  return style;
};

const parseTextStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  if (data.fontSize) style.fontSize = parseInt(data.fontSize, 10);
  if (data.color) style.color = data.color;
  if (data.fontWeight) style.fontWeight = data.fontWeight === 'bold' ? 'bold' : 
    data.fontWeight === '600' ? '600' : 
    data.fontWeight === '700' ? '700' : 'normal';
  if (data.textAlign) style.textAlign = data.textAlign;
  if (data.paddingHorizontal) style.paddingHorizontal = parseInt(data.paddingHorizontal, 10);
  if (data.paddingVertical) style.paddingVertical = parseInt(data.paddingVertical, 10);
  if (data.backgroundColor) style.backgroundColor = data.backgroundColor;
  if (data.borderRadius) style.borderRadius = parseInt(data.borderRadius, 10);
  
  return style;
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f8f9fa',
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#f8f9fa',
  },
  loadingText: {
    marginTop: 16,
    fontSize: 16,
    color: '#666',
  },
  baseText: {
    fontSize: 14,
    color: '#000',
  },
  iconButton: {
    padding: 8,
  },
  input: {
    flex: 1,
    fontSize: 16,
    color: '#000',
    padding: 0,
  },
  textButton: {
    paddingVertical: 8,
  },
  textButtonText: {
    fontSize: 14,
    fontWeight: '600',
  },
  // Fallback UI Styles
  fallbackContainer: {
    flex: 1,
    backgroundColor: '#f8f9fa',
  },
  fallbackHeader: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 16,
    backgroundColor: '#ffffff',
    borderBottomWidth: 1,
    borderBottomColor: '#f0f0f0',
    gap: 12,
  },
  fallbackSearchContainer: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#f5f5f5',
    borderRadius: 10,
    paddingHorizontal: 12,
    paddingVertical: 12,
  },
  fallbackSearchInput: {
    flex: 1,
    fontSize: 16,
    color: '#000',
    marginLeft: 8,
  },
  fallbackFilterButton: {
    padding: 8,
  },
  fallbackTabs: {
    padding: 16,
  },
  fallbackTab: {
    backgroundColor: '#f5f5f5',
    borderRadius: 20,
    paddingHorizontal: 16,
    paddingVertical: 8,
    marginRight: 8,
  },
  fallbackTabText: {
    fontSize: 14,
    fontWeight: '500',
    color: '#666',
  },
  fallbackContent: {
    padding: 16,
  },
  fallbackTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 16,
  },
  fallbackCard: {
    backgroundColor: '#ffffff',
    borderRadius: 12,
    padding: 16,
    marginBottom: 12,
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
      },
      android: {
        elevation: 3,
      },
    }),
  },
  cardHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 12,
  },
  matchBadge: {
    backgroundColor: '#4CAF50',
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 6,
  },
  matchText: {
    color: '#ffffff',
    fontSize: 12,
    fontWeight: '500',
  },
  timeLeft: {
    color: '#FF3B30',
    fontSize: 12,
    fontWeight: '500',
  },
  route: {
    fontSize: 16,
    fontWeight: '600',
    color: '#333',
    marginBottom: 4,
  },
  price: {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#FF3B30',
    marginBottom: 12,
  },
  detail: {
    fontSize: 14,
    color: '#666',
    marginBottom: 4,
  },
  viewButton: {
    marginTop: 12,
    paddingVertical: 8,
  },
  viewButtonText: {
    fontSize: 14,
    fontWeight: '600',
    color: '#FF3B30',
  },
  noResults: {
    alignItems: 'center',
    padding: 40,
  },
  noResultsText: {
    fontSize: 16,
    color: '#333',
    marginTop: 16,
    marginBottom: 8,
  },
  noResultsSubtext: {
    fontSize: 14,
    color: '#666',
  },
  refreshButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#FF3B30',
    margin: 16,
    padding: 16,
    borderRadius: 12,
    gap: 8,
  },
  refreshButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
});

export default MarketScreen;