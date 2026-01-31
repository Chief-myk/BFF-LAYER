// profile.tsx - Profile Screen with Backend Integration
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
  Image,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { SafeAreaView } from 'react-native-safe-area-context';
import axios from 'axios';

const API_URL = 'http://192.168.1.3:8080/bff/driver/profile';

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

const ProfileScreen: React.FC = () => {
  const [uiData, setUiData] = useState<UISnippet[]>([]);
  const [screenData, setScreenData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshing, setRefreshing] = useState<boolean>(false);

  const fetchProfileData = useCallback(async (forceRefresh = false) => {
    try {
      if (!forceRefresh) {
        setLoading(true);
      } else {
        setRefreshing(true);
      }

      console.log('Fetching profile data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Profile response:', response.data.status);

      if (response.data.status === 'success') {
        console.log('Profile UI Data length:', response.data.ui?.length);
        setScreenData(response.data.data);
        setUiData(response.data.ui || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch profile data');
      }
    } catch (err: any) {
      console.error('Error fetching profile data:', err.message || err);
      setError(err.message || 'Network error. Please check your connection.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchProfileData();
  }, [fetchProfileData]);

  const handleAction = async (action: ActionData) => {
    console.log('Profile action triggered:', action);
    
    try {
      switch (action.type) {
        case 'NAVIGATE':
          if (action.to === 'EditProfile') {
            Alert.alert('Edit Profile', 'Would navigate to edit profile screen');
          }
          break;
          
        case 'LOGOUT':
          Alert.alert(
            'Logout',
            'Are you sure you want to logout?',
            [
              { text: 'Cancel', style: 'cancel' },
              { 
                text: 'Logout', 
                style: 'destructive',
                onPress: () => {
                  Alert.alert('Success', 'Logged out successfully');
                }
              },
            ]
          );
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
      'pencil-outline': 'pencil-outline',
      'checkmark-circle': 'checkmark-circle',
      'wallet-outline': 'wallet-outline',
      'document-text-outline': 'document-text-outline',
      'car-outline': 'car-outline',
      'phone-portrait-outline': 'phone-portrait-outline',
      'business-outline': 'business-outline',
      'log-out-outline': 'log-out-outline',
      'refresh-outline': 'refresh-outline',
    };
    return iconMap[iconName] || 'help-circle-outline';
  };

  // Renderer for UI snippets
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    console.log(`Rendering snippet type: ${snippet.type}`, snippet.data);
    
    switch (snippet.type) {
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
        
      case 'ICON_BUTTON':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.iconButton, parseStyles(snippet.data)]}
            onPress={() => snippet.data?.onPress && handleAction(snippet.data.onPress)}
            activeOpacity={0.7}
          >
            <Ionicons
              name={getIconName(snippet.data?.icon)}
              size={24}
              color="#666"
            />
          </TouchableOpacity>
        );
        
      case 'CARD':
        return (
          <View key={key} style={[styles.card, parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'IMAGE':
        if (!snippet.data?.url) return null;
        return (
          <Image
            key={key}
            source={{ uri: snippet.data.url }}
            style={[styles.image, parseStyles(snippet.data)]}
          />
        );
        
      case 'BUTTON':
        return (
          <TouchableOpacity
            key={key}
            style={[styles.button, parseStyles(snippet.data?.style)]}
            onPress={() => snippet.data?.action && handleAction(snippet.data.action)}
            activeOpacity={0.7}
          >
            <Text style={styles.buttonText}>
              {snippet.data?.text || 'Button'}
            </Text>
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
                onRefresh={() => fetchProfileData(true)}
                colors={['#ff0000']}
                tintColor="#ff0000"
              />
            }
            showsVerticalScrollIndicator={false}
            contentContainerStyle={styles.scrollContent}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
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
          style={styles.fallbackScrollView}
          refreshControl={
            <RefreshControl
              refreshing={refreshing}
              onRefresh={() => fetchProfileData(true)}
              colors={['#ff0000']}
              tintColor="#ff0000"
            />
          }
          showsVerticalScrollIndicator={false}
          contentContainerStyle={styles.fallbackScrollContent}
        >
          {/* Header */}
          <View style={styles.fallbackHeader}>
            <Text style={styles.fallbackHeaderTitle}>Profile</Text>
            <TouchableOpacity 
              style={styles.fallbackEditButton}
              onPress={() => Alert.alert('Edit Profile', 'Would navigate to edit profile screen')}
            >
              <Ionicons name="pencil-outline" size={24} color="#666" />
            </TouchableOpacity>
          </View>

          {/* Profile Info Card */}
          <View style={styles.fallbackProfileCard}>
            <View style={styles.fallbackProfileHeader}>
              <Image
                source={{ uri: 'https://i.pravatar.cc/150?img=12' }}
                style={styles.fallbackProfileImage}
              />
              <View style={styles.fallbackProfileInfo}>
                <View style={styles.fallbackNameContainer}>
                  <Text style={styles.fallbackProfileName}>Rajesh Kumar</Text>
                  <View style={styles.fallbackVerifiedBadge}>
                    <Ionicons name="checkmark-circle" size={16} color="#4CAF50" />
                    <Text style={styles.fallbackVerifiedText}>Verified</Text>
                  </View>
                </View>
                
                <View style={styles.fallbackProfileDetails}>
                  <View style={styles.fallbackDetailItem}>
                    <Ionicons name="phone-portrait-outline" size={16} color="#666" />
                    <Text style={styles.fallbackDetailText}>+919123456789</Text>
                  </View>
                  
                  <View style={styles.fallbackDetailItem}>
                    <Ionicons name="car-outline" size={16} color="#666" />
                    <Text style={styles.fallbackDetailText}>MH01AB1234</Text>
                  </View>
                  
                  <View style={styles.fallbackDetailItem}>
                    <Ionicons name="business-outline" size={16} color="#666" />
                    <Text style={styles.fallbackDetailText}>Kumar Logistics</Text>
                  </View>
                </View>
              </View>
            </View>
          </View>

          {/* Wallet Card */}
          <View style={styles.fallbackWalletCard}>
            <Text style={styles.fallbackWalletTitle}>Wallet Balance</Text>
            <Text style={styles.fallbackWalletAmount}>₹12,500</Text>
          </View>

          {/* Documents Section */}
          <View style={styles.fallbackDocumentsCard}>
            <Text style={styles.fallbackSectionTitle}>Documents</Text>
            
            <View style={styles.fallbackDocumentsList}>
              <View style={styles.fallbackDocumentItem}>
                <View style={styles.fallbackDocumentInfo}>
                  <Ionicons name="document-text-outline" size={24} color="#4CAF50" />
                  <View style={styles.fallbackDocumentText}>
                    <Text style={styles.fallbackDocumentName}>Aadhar Card</Text>
                    <Text style={styles.fallbackDocumentStatusVerified}>Verified</Text>
                  </View>
                </View>
                <Ionicons name="checkmark-circle" size={20} color="#4CAF50" />
              </View>
              
              <View style={styles.fallbackDocumentItem}>
                <View style={styles.fallbackDocumentInfo}>
                  <Ionicons name="document-text-outline" size={24} color="#4CAF50" />
                  <View style={styles.fallbackDocumentText}>
                    <Text style={styles.fallbackDocumentName}>PAN Card</Text>
                    <Text style={styles.fallbackDocumentStatusVerified}>Verified</Text>
                  </View>
                </View>
                <Ionicons name="checkmark-circle" size={20} color="#4CAF50" />
              </View>
              
              <View style={styles.fallbackDocumentItem}>
                <View style={styles.fallbackDocumentInfo}>
                  <Ionicons name="document-text-outline" size={24} color="#FF9800" />
                  <View style={styles.fallbackDocumentText}>
                    <Text style={styles.fallbackDocumentName}>Driving License</Text>
                    <Text style={styles.fallbackDocumentStatusPending}>Pending</Text>
                  </View>
                </View>
                <Ionicons name="time-outline" size={20} color="#FF9800" />
              </View>
              
              <View style={styles.fallbackDocumentItem}>
                <View style={styles.fallbackDocumentInfo}>
                  <Ionicons name="document-text-outline" size={24} color="#4CAF50" />
                  <View style={styles.fallbackDocumentText}>
                    <Text style={styles.fallbackDocumentName}>RC & Insurance</Text>
                    <Text style={styles.fallbackDocumentStatusVerified}>Verified</Text>
                  </View>
                </View>
                <Ionicons name="checkmark-circle" size={20} color="#4CAF50" />
              </View>
            </View>
          </View>

          {/* Error Message */}
          {error && (
            <View style={styles.fallbackErrorContainer}>
              <Ionicons name="warning-outline" size={48} color="#ff0000" />
              <Text style={styles.fallbackErrorTitle}>Oops! Something went wrong</Text>
              <Text style={styles.fallbackErrorMessage}>
                {error || 'Unable to load profile data. Please check your connection and try again.'}
              </Text>
              
              <TouchableOpacity 
                style={styles.fallbackRetryButton}
                onPress={() => fetchProfileData(true)}
              >
                <Ionicons name="refresh-outline" size={20} color="#fff" />
                <Text style={styles.fallbackRetryButtonText}>Try Again</Text>
              </TouchableOpacity>
            </View>
          )}

          {/* Logout Button */}
          <TouchableOpacity 
            style={styles.fallbackLogoutButton}
            onPress={() => {
              Alert.alert(
                'Logout',
                'Are you sure you want to logout?',
                [
                  { text: 'Cancel', style: 'cancel' },
                  { 
                    text: 'Logout', 
                    style: 'destructive',
                    onPress: () => {
                      Alert.alert('Success', 'Logged out successfully');
                    }
                  },
                ]
              );
            }}
          >
            <Ionicons name="log-out-outline" size={20} color="#fff" />
            <Text style={styles.fallbackLogoutButtonText}>Logout</Text>
          </TouchableOpacity>

          <View style={styles.fallbackSpacer} />
        </ScrollView>
      </SafeAreaView>
    );
  };

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
        <View style={styles.loadingContent}>
          <ActivityIndicator size="large" color="#ff0000" />
          <Text style={styles.loadingText}>Loading profile...</Text>
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
      <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
      <ScrollView
        style={styles.scrollView}
        refreshControl={
          <RefreshControl
            refreshing={refreshing}
            onRefresh={() => fetchProfileData(true)}
            colors={['#ff0000']}
            tintColor="#ff0000"
          />
        }
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.scrollContent}
      >
        {uiData.map((snippet, index) => renderSnippet(snippet, index))}
      </ScrollView>
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
  if (data.width || data.Width) style.width = data.width || data.Width;
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
  scrollContent: {
    padding: 16,
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
  iconButton: {
    padding: 8,
  },
  card: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
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
  image: {
    borderRadius: 35,
  },
  button: {
    backgroundColor: '#ff0000',
    padding: 16,
    borderRadius: 12,
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 16,
  },
  buttonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: '600',
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
  fallbackScrollView: {
    flex: 1,
  },
  fallbackScrollContent: {
    padding: 16,
  },
  fallbackHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    backgroundColor: '#FFFFFF',
    paddingHorizontal: 20,
    paddingVertical: 20,
    marginBottom: 16,
    borderRadius: 12,
  },
  fallbackHeaderTitle: {
    fontSize: 24,
    fontWeight: '700',
    color: '#1A1A1A',
  },
  fallbackEditButton: {
    padding: 8,
  },
  fallbackProfileCard: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
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
  fallbackProfileHeader: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 20,
  },
  fallbackProfileImage: {
    width: 70,
    height: 70,
    borderRadius: 35,
    marginRight: 16,
  },
  fallbackProfileInfo: {
    flex: 1,
  },
  fallbackNameContainer: {
    marginBottom: 12,
  },
  fallbackProfileName: {
    fontSize: 20,
    fontWeight: '600',
    color: '#1A1A1A',
    marginBottom: 8,
  },
  fallbackVerifiedBadge: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#F8F8F8',
    paddingHorizontal: 12,
    paddingVertical: 6,
    borderRadius: 16,
    alignSelf: 'flex-start',
    gap: 4,
  },
  fallbackVerifiedText: {
    fontSize: 12,
    fontWeight: '600',
    color: '#4CAF50',
  },
  fallbackProfileDetails: {
    gap: 8,
  },
  fallbackDetailItem: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  fallbackDetailText: {
    fontSize: 14,
    color: '#666',
  },
  fallbackWalletCard: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
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
  fallbackWalletTitle: {
    fontSize: 18,
    fontWeight: '600',
    color: '#1A1A1A',
    marginBottom: 8,
  },
  fallbackWalletAmount: {
    fontSize: 28,
    fontWeight: '700',
    color: '#1A1A1A',
  },
  fallbackDocumentsCard: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
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
  fallbackSectionTitle: {
    fontSize: 18,
    fontWeight: '600',
    color: '#1A1A1A',
    marginBottom: 16,
  },
  fallbackDocumentsList: {
    gap: 12,
  },
  fallbackDocumentItem: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 12,
    backgroundColor: '#F8F8F8',
    borderRadius: 8,
  },
  fallbackDocumentInfo: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 12,
  },
  fallbackDocumentText: {
    gap: 2,
  },
  fallbackDocumentName: {
    fontSize: 14,
    fontWeight: '600',
    color: '#1A1A1A',
  },
  fallbackDocumentStatusVerified: {
    fontSize: 12,
    color: '#4CAF50',
  },
  fallbackDocumentStatusPending: {
    fontSize: 12,
    color: '#FF9800',
  },
  fallbackErrorContainer: {
    backgroundColor: '#FFFFFF',
    borderRadius: 12,
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
  fallbackLogoutButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#ff0000',
    padding: 16,
    borderRadius: 12,
    gap: 8,
    marginTop: 8,
  },
  fallbackLogoutButtonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: '600',
  },
  fallbackSpacer: {
    height: 32,
  },
});

export default ProfileScreen;