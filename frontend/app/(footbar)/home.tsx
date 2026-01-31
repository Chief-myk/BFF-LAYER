// home.tsx - Simplified working version
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
  Image,
  RefreshControl,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Ionicons } from '@expo/vector-icons';
import axios from 'axios';

const API_URL = 'http://192.168.1.3:8080/bff/driver/home';

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

const HomeScreen: React.FC = () => {
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

      console.log('Fetching data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Response received:', response.data.status);

      if (response.data.status === 'success') {
        console.log('UI Data length:', response.data.ui?.length);
        setScreenData(response.data.data);
        setUiData(response.data.ui || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch data');
      }
    } catch (err: any) {
      console.error('Error fetching home data:', err.message || err);
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
    console.log('Action triggered:', action);
    
    try {
      switch (action.type) {
        case 'ACTION':
          if (action.url) {
            console.log('Making API call to:', action.url);
            const response = await axios.post(action.url, {
              action: action.value,
              data: action.data || {},
            });
            
            console.log('API response:', response.data);
            
            if (response.data.status === 'success') {
              Alert.alert('Success', response.data.message || 'Action completed');
              // Refresh data
              fetchHomeData(true);
            } else {
              Alert.alert('Error', response.data.message || 'Action failed');
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

  // Simple renderer for basic UI components
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    switch (snippet.type) {
      case 'VIEW':
        return (
          <View key={key} style={parseStyles(snippet.data)}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'TEXT':
        return (
          <Text key={key} style={[styles.baseText, parseTextStyles(snippet.data)]}>
            {snippet.data.text || ''}
          </Text>
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
            style={parseStyles(snippet.data)}
            onPress={() => snippet.data.onPress && handleAction(snippet.data.onPress)}
          >
            <Ionicons
              name={snippet.data.icon || 'help-circle-outline'}
              size={snippet.data.size || 24}
              color={snippet.data.color || '#000'}
            />
          </TouchableOpacity>
        );
        
      case 'BUTTON':
        return (
          <TouchableOpacity
            key={key}
            style={[
              styles.buttonBase,
              parseStyles(snippet.data.style || {}),
              snippet.data.disabled && styles.buttonDisabled,
            ]}
            onPress={() => snippet.data.action && handleAction(snippet.data.action)}
            disabled={snippet.data.disabled}
          >
            <Text style={styles.buttonText}>
              {snippet.data.text || 'Button'}
            </Text>
          </TouchableOpacity>
        );
        
      case 'SCROLL':
        return (
          <ScrollView
            key={key}
            style={parseStyles(snippet.data)}
            showsVerticalScrollIndicator={false}
            refreshControl={
              <RefreshControl
                refreshing={refreshing}
                onRefresh={() => fetchHomeData(true)}
                colors={['#1a237e']}
                tintColor="#1a237e"
              />
            }
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
      case 'PRESSABLE_CARD':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.cardBase, parseStyles(snippet.data.cardData || {})]}
            onPress={() => snippet.data.cardData?.onPress && handleAction(snippet.data.cardData.onPress)}
            activeOpacity={0.8}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </TouchableOpacity>
        );
        
      default:
        console.log(`Unknown snippet type: ${snippet.type}`);
        return null;
    }
  };

  // Debug: Show what we're getting
  console.log('Current UI Data:', uiData);
  console.log('Loading:', loading);
  console.log('Error:', error);

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#1a237e" barStyle="light-content" />
        <ActivityIndicator size="large" color="#1a237e" />
        <Text style={styles.loadingText}>Loading your dashboard...</Text>
      </SafeAreaView>
    );
  }

  if (error) {
    return (
      <SafeAreaView style={styles.errorContainer}>
        <StatusBar backgroundColor="#1a237e" barStyle="light-content" />
        <Ionicons name="warning-outline" size={64} color="#F44336" />
        <Text style={styles.errorTitle}>Something went wrong</Text>
        <Text style={styles.errorMessage}>{error}</Text>
        <TouchableOpacity style={styles.retryButton} onPress={() => fetchHomeData(true)}>
          <Text style={styles.retryButtonText}>Try Again</Text>
        </TouchableOpacity>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar backgroundColor="#1a237e" barStyle="light-content" />
      
      {uiData.length > 0 ? (
        uiData.map((snippet, index) => renderSnippet(snippet, index))
      ) : (
        <View style={styles.fallbackContainer}>
          <Ionicons name="home-outline" size={64} color="#666" />
          <Text style={styles.fallbackText}>No UI data available</Text>
          <TouchableOpacity style={styles.refreshButton} onPress={() => fetchHomeData(true)}>
            <Text style={styles.refreshButtonText}>Refresh</Text>
          </TouchableOpacity>
        </View>
      )}
    </SafeAreaView>
  );
};

// Style parsing functions (simplified)
const parseStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  // Map basic style properties
  if (data.backgroundColor) style.backgroundColor = data.backgroundColor;
  if (data.padding) style.padding = data.padding;
  if (data.paddingHorizontal) style.paddingHorizontal = data.paddingHorizontal;
  if (data.paddingVertical) style.paddingVertical = data.paddingVertical;
  if (data.margin) style.margin = data.margin;
  if (data.marginTop) style.marginTop = data.marginTop;
  if (data.marginBottom) style.marginBottom = data.marginBottom;
  if (data.borderRadius) style.borderRadius = data.borderRadius;
  if (data.borderWidth) style.borderWidth = data.borderWidth;
  if (data.borderColor) style.borderColor = data.borderColor;
  if (data.width) style.width = typeof data.width === 'string' && data.width.includes('%') ? data.width : parseInt(data.width, 10);
  if (data.height) style.height = typeof data.height === 'string' && data.height.includes('%') ? data.height : parseInt(data.height, 10);
  
  // Layout properties
  if (data.flexDirection) style.flexDirection = data.flexDirection;
  if (data.justifyContent) style.justifyContent = data.justifyContent;
  if (data.alignItems) style.alignItems = data.alignItems;
  if (data.flexWrap) style.flexWrap = data.flexWrap;
  if (data.flexGrow) style.flexGrow = data.flexGrow;
  if (data.flex) style.flex = data.flex;
  if (data.gap) style.gap = data.gap;
  
  // Add shadow/elevation
  if (data.elevation !== undefined) {
    if (Platform.OS === 'android') {
      style.elevation = data.elevation;
    } else {
      style.shadowColor = data.shadowColor || '#000';
      style.shadowOffset = data.shadowOffset || { width: 0, height: 2 };
      style.shadowOpacity = data.shadowOpacity || 0.1;
      style.shadowRadius = data.shadowRadius || 3;
    }
  }
  
  return style;
};

const parseTextStyles = (data: any): any => {
  if (!data) return {};
  
  const style: any = {};
  
  if (data.fontSize) style.fontSize = parseInt(data.fontSize, 10);
  if (data.color) style.color = data.color;
  if (data.fontWeight) style.fontWeight = data.fontWeight === 'bold' ? 'bold' : data.fontWeight;
  if (data.textAlign) style.textAlign = data.textAlign;
  if (data.marginBottom) style.marginBottom = parseInt(data.marginBottom, 10);
  if (data.paddingHorizontal) style.paddingHorizontal = parseInt(data.paddingHorizontal, 10);
  if (data.paddingVertical) style.paddingVertical = parseInt(data.paddingVertical, 10);
  if (data.lineHeight) style.lineHeight = parseInt(data.lineHeight, 10);
  
  return style;
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f7fa',
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#f5f7fa',
  },
  loadingText: {
    marginTop: 16,
    fontSize: 16,
    color: '#666',
  },
  errorContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#f5f7fa',
    padding: 20,
  },
  errorTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#333',
    marginTop: 16,
    marginBottom: 8,
  },
  errorMessage: {
    fontSize: 14,
    color: '#666',
    textAlign: 'center',
    marginBottom: 24,
  },
  retryButton: {
    backgroundColor: '#1a237e',
    paddingHorizontal: 24,
    paddingVertical: 12,
    borderRadius: 8,
  },
  retryButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  fallbackContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#f5f7fa',
  },
  fallbackText: {
    fontSize: 18,
    color: '#666',
    marginTop: 16,
    marginBottom: 24,
  },
  refreshButton: {
    backgroundColor: '#1a237e',
    paddingHorizontal: 24,
    paddingVertical: 12,
    borderRadius: 8,
  },
  refreshButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  baseText: {
    fontSize: 14,
    color: '#000',
  },
  buttonBase: {
    backgroundColor: '#4CAF50',
    paddingVertical: 12,
    paddingHorizontal: 24,
    borderRadius: 8,
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonDisabled: {
    backgroundColor: '#cccccc',
    opacity: 0.6,
  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  cardBase: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 16,
    marginHorizontal: 20,
    marginVertical: 8,
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
});

export default HomeScreen;