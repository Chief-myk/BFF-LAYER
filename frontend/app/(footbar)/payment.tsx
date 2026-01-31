// payment.tsx - Enhanced Payment Screen with Better Styling
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

const API_URL = 'http://192.168.1.3:8080/bff/driver/payment';

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

const PaymentScreen: React.FC = () => {
  const [uiData, setUiData] = useState<UISnippet[]>([]);
  const [screenData, setScreenData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [refreshing, setRefreshing] = useState<boolean>(false);

  const fetchPaymentData = useCallback(async (forceRefresh = false) => {
    try {
      if (!forceRefresh) {
        setLoading(true);
      } else {
        setRefreshing(true);
      }

      console.log('Fetching payment data from:', API_URL);
      const response = await axios.get<ScreenResponse>(API_URL, {
        timeout: 10000,
        headers: {
          'Content-Type': 'application/json',
        },
      });

      console.log('Payment response:', response.data.status);

      if (response.data.status === 'success') {
        console.log('Payment UI Data length:', response.data.ui?.length);
        // console.log('First snippet:', response.data.ui?.[0]);
        setScreenData(response.data.data);
        setUiData(response.data.ui || []);
        setError(null);
      } else {
        throw new Error(response.data.message || 'Failed to fetch payment data');
      }
    } catch (err: any) {
      console.error('Error fetching payment data:', err.message || err);
      setError(err.message || 'Network error. Please check your connection.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, []);

  useEffect(() => {
    fetchPaymentData();
  }, [fetchPaymentData]);

  const handleAction = async (action: ActionData) => {
    console.log('Payment action triggered:', action);
    
    try {
      switch (action.type) {
        case 'ACTION':
          if (action.value === 'refresh' || action.value === 'button_refresh') {
            fetchPaymentData(true);
          } else if (action.value === 'see_all' || action.value === 'button_see all') {
            Alert.alert('View All', 'Would show all transactions');
          } else if (action.value?.startsWith('button_')) {
            const buttonName = action.value.replace('button_', '');
            Alert.alert(`${buttonName.charAt(0).toUpperCase() + buttonName.slice(1)}`, `Opening ${buttonName} options...`);
          } else if (action.url) {
            const response = await axios.post(action.url, {
              action: action.value,
              data: action.data || {},
            });
            
            if (response.data.status === 'success') {
              Alert.alert('Success', response.data.message || 'Action completed');
              fetchPaymentData(true);
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
      'checkmark-circle': 'checkmark-circle',
      'information-circle-outline': 'information-circle-outline',
      'wallet': 'wallet-outline',
      'cash-multiple': 'cash-outline',
      'clock-outline': 'time-outline',
      'truck-check': 'checkmark-done-outline',
      'bank-transfer': 'swap-horizontal-outline',
      'filter-variant': 'filter-outline',
      'magnify': 'search-outline',
      'file-export': 'download-outline',
      'refresh': 'refresh-outline',
      'calendar-outline': 'calendar-outline',
      'cash-outline': 'cash-outline',
      'chevron-forward': 'chevron-forward',
      'map-outline': 'map-outline',
      'business-outline': 'business-outline',
      'receipt-outline': 'receipt-outline',
      'card-outline': 'card-outline',
      'document-text-outline': 'document-text-outline',
    };
    return iconMap[iconName] || 'help-circle-outline';
  };

  // Enhanced renderer for UI snippets
  const renderSnippet = (snippet: UISnippet, index: number): JSX.Element | null => {
    const key = `${snippet.type}-${index}`;
    
    // console.log(`Rendering snippet type: ${snippet.type}`, snippet.data);
    
    switch (snippet.type) {
      case 'ROW':
        return (
          <View key={key} style={[styles.row, parseStyles(snippet.data)]}>
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </View>
        );
        
      case 'COLUMN':
        return (
          <View key={key} style={[styles.column, parseStyles(snippet.data)]}>
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
        
      case 'TOUCHABLE_OPACITY':
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
        
      case 'STATUS_BAR':
        return null; // Handled by React Native StatusBar
        
      case 'SCROLL':
        return (
          <ScrollView
            key={key}
            style={[styles.scrollView, parseStyles(snippet.data)]}
            contentContainerStyle={styles.scrollContent}
            refreshControl={
              <RefreshControl
                refreshing={refreshing}
                onRefresh={() => fetchPaymentData(true)}
                colors={['#28A745']}
                tintColor="#28A745"
              />
            }
            showsVerticalScrollIndicator={false}
          >
            {snippet.children?.map((child, idx) => renderSnippet(child, idx))}
          </ScrollView>
        );
        
      default:
        console.log(`Unknown snippet type: ${snippet.type}`);
        return null;
    }
  };

  // Fallback UI with enhanced styling
  const renderFallbackUI = () => {
    return (
      <SafeAreaView style={styles.fallbackContainer}>
        <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
        <ScrollView
          style={styles.container}
          refreshControl={
            <RefreshControl
              refreshing={refreshing}
              onRefresh={() => fetchPaymentData(true)}
              colors={['#28A745']}
              tintColor="#28A745"
            />
          }
          showsVerticalScrollIndicator={false}
          contentContainerStyle={styles.scrollContent}
        >
          {/* Enhanced Payment Alert */}
          <View style={styles.fallbackPaymentAlert}>
            <View style={styles.fallbackIconCircle}>
              <Ionicons name="checkmark-circle" size={28} color="#28A745" />
            </View>
            <View style={styles.fallbackAlertContent}>
              <Text style={styles.fallbackAlertTitle}>Payment received for trip 4580</Text>
              <Text style={styles.fallbackAlertAmount}>₹5,200 credited to wallet</Text>
              <Text style={styles.fallbackAlertTime}>2 hours ago</Text>
            </View>
            <Ionicons name="information-circle-outline" size={24} color="#007AFF" />
          </View>

          <View style={styles.sectionSpacer} />

          {/* Enhanced Wallet Summary */}
          <View style={styles.sectionContainer}>
            <View style={styles.sectionHeader}>
              <Text style={styles.sectionTitle}>Wallet Summary</Text>
              <TouchableOpacity 
                style={styles.refreshButtonSmall}
                onPress={() => fetchPaymentData(true)}
              >
                <Ionicons name="refresh" size={16} color="#007AFF" />
                <Text style={styles.refreshButtonTextSmall}>Refresh</Text>
              </TouchableOpacity>
            </View>

            <View style={styles.walletCardsContainer}>
              <View style={[styles.walletCard, { borderLeftColor: '#28A745' }]}>
                <View style={[styles.walletIconContainer, { backgroundColor: '#E8F5E8' }]}>
                  <Ionicons name="wallet-outline" size={28} color="#28A745" />
                </View>
                <Text style={styles.walletCardTitle}>Current Balance</Text>
                <Text style={styles.walletCardAmount}>₹45,200</Text>
                <Text style={styles.walletCardSubtitle}>Available for use</Text>
              </View>

              <View style={[styles.walletCard, { borderLeftColor: '#007AFF' }]}>
                <View style={[styles.walletIconContainer, { backgroundColor: '#E8F5E8' }]}>
                  <Ionicons name="cash-outline" size={28} color="#007AFF" />
                </View>
                <Text style={styles.walletCardTitle}>Total Earnings</Text>
                <Text style={styles.walletCardAmount}>₹78,500</Text>
                <Text style={styles.walletCardSubtitle}>This month</Text>
              </View>

              <View style={[styles.walletCard, { borderLeftColor: '#FF9500' }]}>
                <View style={[styles.walletIconContainer, { backgroundColor: '#FFF3E0' }]}>
                  <Ionicons name="time-outline" size={28} color="#FF9500" />
                </View>
                <Text style={styles.walletCardTitle}>Pending</Text>
                <Text style={styles.walletCardAmount}>₹12,30</Text>
                <Text style={styles.walletCardSubtitle}>Awaiting clearance</Text>
              </View>
            </View>
          </View>

          <View style={styles.sectionSpacer} />

          {/* Enhanced Recent Transactions */}
          <View style={styles.sectionContainer}>
            <View style={styles.sectionHeader}>
              <Text style={styles.sectionTitle}>Recent Transactions</Text>
              <TouchableOpacity>
                <Text style={styles.sectionAction}>See All</Text>
              </TouchableOpacity>
            </View>

            <View style={styles.transactionCard}>
              <View style={styles.transactionHeader}>
                <View style={[styles.transactionIconContainer, { backgroundColor: '#E8F5E8' }]}>
                  <Ionicons name="checkmark-done-outline" size={24} color="#28A745" />
                </View>
                <View style={styles.transactionInfo}>
                  <Text style={styles.transactionTitle}>Commission Earned</Text>
                  <Text style={styles.transactionSubtitle}>TRIP#4587 • Mumbai → Delhi</Text>
                </View>
                <Text style={styles.transactionAmountSuccess}>₹2,500</Text>
              </View>
              <View style={styles.transactionFooter}>
                <View style={[styles.statusBadge, { backgroundColor: '#E8F5E8' }]}>
                  <Text style={[styles.statusText, { color: '#28A745' }]}>Success</Text>
                </View>
                <Text style={styles.transactionTime}>2024-09-15 • 10:30 AM</Text>
              </View>
            </View>

            <View style={styles.transactionCard}>
              <View style={styles.transactionHeader}>
                <View style={[styles.transactionIconContainer, { backgroundColor: '#FFE5E5' }]}>
                  <Ionicons name="swap-horizontal-outline" size={24} color="#DC3545" />
                </View>
                <View style={styles.transactionInfo}>
                  <Text style={styles.transactionTitle}>Withdrawal</Text>
                  <Text style={styles.transactionSubtitle}>Bank Transfer</Text>
                </View>
                <Text style={styles.transactionAmountPending}>₹5,000</Text>
              </View>
              <View style={styles.transactionFooter}>
                <View style={[styles.statusBadge, { backgroundColor: '#FFE5E5' }]}>
                  <Text style={[styles.statusText, { color: '#DC3545' }]}>Pending</Text>
                </View>
                <Text style={styles.transactionTime}>2024-09-13 • 09:45 AM</Text>
              </View>
            </View>
          </View>

          <View style={styles.sectionSpacer} />

          {/* Enhanced Trip Wise Earnings */}
          <View style={styles.sectionContainer}>
            <View style={styles.sectionHeader}>
              <Text style={styles.sectionTitle}>Trip Wise Earnings</Text>
              <TouchableOpacity>
                <Text style={styles.sectionAction}>See All</Text>
              </TouchableOpacity>
            </View>

            <View style={styles.tripCard}>
              <View style={styles.tripHeader}>
                <View>
                  <Text style={styles.tripId}>TRIP#4587</Text>
                  <Text style={styles.tripRoute}>Mumbai → Delhi</Text>
                </View>
                <View style={[styles.statusBadge, { backgroundColor: '#E8F5E8' }]}>
                  <Text style={[styles.statusText, { color: '#28A745' }]}>Settled</Text>
                </View>
              </View>
              <View style={styles.tripDetails}>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="map-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>1,416 km</Text>
                </View>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="business-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>ABC Logistics</Text>
                </View>
              </View>
              <View style={styles.tripFooter}>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="calendar-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>2024-09-15</Text>
                </View>
                <View>
                  <Text style={styles.tripEarningsLabel}>Earnings</Text>
                  <Text style={styles.tripEarningsAmount}>₹15,200</Text>
                </View>
              </View>
            </View>

            <View style={styles.tripCard}>
              <View style={styles.tripHeader}>
                <View>
                  <Text style={styles.tripId}>TRIP#4498</Text>
                  <Text style={styles.tripRoute}>Delhi → Kolkata</Text>
                </View>
                <View style={[styles.statusBadge, { backgroundColor: '#FFF3E0' }]}>
                  <Text style={[styles.statusText, { color: '#FF9500' }]}>Pending</Text>
                </View>
              </View>
              <View style={styles.tripDetails}>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="map-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>1,533 km</Text>
                </View>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="business-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>PQR Freight</Text>
                </View>
              </View>
              <View style={styles.tripFooter}>
                <View style={styles.tripDetailItem}>
                  <Ionicons name="calendar-outline" size={18} color="#666666" />
                  <Text style={styles.tripDetailText}>2024-09-13</Text>
                </View>
                <View>
                  <Text style={styles.tripEarningsLabel}>Earnings</Text>
                  <Text style={styles.tripEarningsAmount}>₹16,800</Text>
                </View>
              </View>
            </View>
          </View>

          <View style={styles.sectionSpacer} />

          {/* Enhanced Settlement Details */}
          <View style={styles.sectionContainer}>
            <View style={styles.sectionHeader}>
              <Text style={styles.sectionTitle}>Settlement Details</Text>
              <TouchableOpacity>
                <Text style={styles.sectionAction}>View All</Text>
              </TouchableOpacity>
            </View>

            <View style={styles.settlementCard}>
              <View style={styles.settlementHeader}>
                <View style={styles.settlementIdContainer}>
                  <Ionicons name="receipt-outline" size={20} color="#666666" />
                  <Text style={styles.settlementId}>SET#789456</Text>
                </View>
                <Text style={styles.settlementAmount}>₹15,200</Text>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="business-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Broker</Text>
                  <Text style={styles.detailValue}>ABC Logistics</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="card-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Payment Mode</Text>
                  <Text style={styles.detailValue}>Bank Transfer</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="calendar-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Date</Text>
                  <Text style={styles.detailValue}>2024-09-15</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="document-text-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Remarks</Text>
                  <Text style={styles.detailValue}>Full payment received</Text>
                </View>
              </View>
            </View>

            <View style={styles.settlementCard}>
              <View style={styles.settlementHeader}>
                <View style={styles.settlementIdContainer}>
                  <Ionicons name="receipt-outline" size={20} color="#666666" />
                  <Text style={styles.settlementId}>SET#789455</Text>
                </View>
                <Text style={styles.settlementAmount}>₹8,500</Text>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="business-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Broker</Text>
                  <Text style={styles.detailValue}>XYZ Transport</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="card-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Payment Mode</Text>
                  <Text style={styles.detailValue}>UPI</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="calendar-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Date</Text>
                  <Text style={styles.detailValue}>2024-09-14</Text>
                </View>
              </View>
              <View style={styles.settlementDetailItem}>
                <View style={styles.detailIconContainer}>
                  <Ionicons name="document-text-outline" size={20} color="#007AFF" />
                </View>
                <View>
                  <Text style={styles.detailLabel}>Remarks</Text>
                  <Text style={styles.detailValue}>Commission deducted: ₹500</Text>
                </View>
              </View>
            </View>
          </View>

          <View style={styles.sectionSpacer} />

          {/* Enhanced Action Buttons */}
          <View style={styles.actionButtonsContainer}>
            <TouchableOpacity style={styles.actionButton}>
              <Ionicons name="filter-outline" size={20} color="#666666" />
              <Text style={styles.actionButtonText}>Filter</Text>
            </TouchableOpacity>
            
            <TouchableOpacity style={styles.actionButton}>
              <Ionicons name="search-outline" size={20} color="#666666" />
              <Text style={styles.actionButtonText}>Search</Text>
            </TouchableOpacity>
            
            <TouchableOpacity style={styles.actionButton}>
              <Ionicons name="download-outline" size={20} color="#666666" />
              <Text style={styles.actionButtonText}>Export</Text>
            </TouchableOpacity>
          </View>

          <View style={styles.finalSpacer} />
        </ScrollView>

        <TouchableOpacity style={styles.retryBackendButton} onPress={() => fetchPaymentData(true)}>
          <Ionicons name="refresh" size={20} color="#fff" />
          <Text style={styles.retryBackendButtonText}>Retry Backend</Text>
        </TouchableOpacity>
      </SafeAreaView>
    );
  };

  if (loading && !screenData) {
    return (
      <SafeAreaView style={styles.loadingContainer}>
        <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
        <View style={styles.loadingContent}>
          <ActivityIndicator size="large" color="#28A745" />
          <Text style={styles.loadingText}>Loading payments...</Text>
        </View>
      </SafeAreaView>
    );
  }

  if (error || uiData.length === 0) {
    console.log('Showing fallback UI due to:', error || 'No UI data');
    return renderFallbackUI();
  }

  // console.log('Rendering backend UI with', uiData.length, 'snippets');
  
  // Render UI from backend
  return (
    <SafeAreaView style={styles.container}>
      <StatusBar backgroundColor="#ffffff" barStyle="dark-content" />
      {uiData.map((snippet, index) => renderSnippet(snippet, index))}
    </SafeAreaView>
  );
};

// Enhanced Style parsing functions
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
  if (data.width || data.Width) style.width = data.width || data.Width;
  if (data.flex || data.Flex) style.flex = data.flex || data.Flex;
  
  // Layout properties
  if (data.flexDirection) style.flexDirection = data.flexDirection;
  if (data.justifyContent) style.justifyContent = data.justifyContent;
  if (data.alignItems) style.alignItems = data.alignItems;
  if (data.alignSelf) style.alignSelf = data.alignSelf;
  if (data.gap) style.gap = data.gap;
  if (data.flexGrow) style.flexGrow = data.flexGrow;
  if (data.flexShrink) style.flexShrink = data.flexShrink;
  if (data.flexWrap) style.flexWrap = data.flexWrap;
  
  // Border properties
  if (data.borderWidth || data.BorderWidth) style.borderWidth = data.borderWidth || data.BorderWidth;
  if (data.borderColor || data.BorderColor) style.borderColor = data.borderColor || data.BorderColor;
  if (data.borderTopWidth || data.BorderTopWidth) style.borderTopWidth = data.borderTopWidth || data.BorderTopWidth;
  if (data.borderTopColor || data.BorderTopColor) style.borderTopColor = data.borderTopColor || data.BorderTopColor;
  if (data.borderLeftWidth) style.borderLeftWidth = data.borderLeftWidth;
  if (data.borderLeftColor) style.borderLeftColor = data.borderLeftColor;
  
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
      style.shadowRadius = data.shadowRadius || data.ShadowRadius || 8;
    }
  }
  
  // Opacity
  if (data.opacity !== undefined) style.opacity = data.opacity;
  
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
  if (data.opacity !== undefined) style.opacity = data.opacity;
  if (data.lineHeight) style.lineHeight = data.lineHeight;
  
  return style;
};

// Enhanced Styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f7fa',
  },
  scrollView: {
    flex: 1,
  },
  scrollContent: {
    paddingVertical: 16,
  },
  loadingContainer: {
    flex: 1,
    backgroundColor: '#f5f7fa',
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
  row: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  column: {
    flexDirection: 'column',
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
  // Fallback UI Enhanced Styles
  fallbackContainer: {
    flex: 1,
    backgroundColor: '#f5f7fa',
  },
  fallbackPaymentAlert: {
    flexDirection: 'row',
    alignItems: 'center',
    backgroundColor: '#E8F5E8',
    borderRadius: 16,
    padding: 20,
    marginHorizontal: 16,
    marginBottom: 24,
    borderWidth: 1,
    borderColor: '#C8E6C9',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.05,
        shadowRadius: 8,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  fallbackIconCircle: {
    width: 48,
    height: 48,
    borderRadius: 24,
    backgroundColor: '#ffffff',
    justifyContent: 'center',
    alignItems: 'center',
    padding: 10,
  },
  fallbackAlertContent: {
    flex: 1,
    marginLeft: 16,
    marginRight: 16,
  },
  fallbackAlertTitle: {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#1a1a1a',
    marginBottom: 4,
  },
  fallbackAlertAmount: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#28A745',
    marginBottom: 4,
  },
  fallbackAlertTime: {
    fontSize: 14,
    color: '#666666',
    opacity: 0.8,
  },
  sectionContainer: {
    paddingHorizontal: 16,
    marginBottom: 24,
  },
  sectionHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  sectionTitle: {
    fontSize: 22,
    fontWeight: 'bold',
    color: '#1a237e',
  },
  sectionAction: {
    fontSize: 16,
    fontWeight: '600',
    color: '#007AFF',
  },
  sectionSpacer: {
    height: 8,
  },
  finalSpacer: {
    height: 32,
  },
  refreshButtonSmall: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 12,
    paddingVertical: 8,
    backgroundColor: '#f0f7ff',
    borderRadius: 20,
    borderWidth: 1,
    borderColor: '#007AFF20',
  },
  refreshButtonTextSmall: {
    fontSize: 14,
    fontWeight: '600',
    color: '#007AFF',
    marginLeft: 6,
  },
  walletCardsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    gap: 12,
  },
  walletCard: {
    flex: 1,
    backgroundColor: '#ffffff',
    padding: 20,
    borderRadius: 16,
    alignItems: 'center',
    borderLeftWidth: 4,
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.08,
        shadowRadius: 12,
      },
      android: {
        elevation: 3,
      },
    }),
  },
  walletIconContainer: {
    width: 60,
    height: 60,
    borderRadius: 30,
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 16,
  },
  walletCardTitle: {
    fontSize: 14,
    fontWeight: '500',
    color: '#666666',
    marginBottom: 8,
    textAlign: 'center',
  },
  walletCardAmount: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#1a237e',
    marginBottom: 4,
  },
  walletCardSubtitle: {
    fontSize: 12,
    color: '#999999',
    textAlign: 'center',
  },
  transactionCard: {
    backgroundColor: '#ffffff',
    borderRadius: 16,
    padding: 20,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: '#f0f0f0',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.05,
        shadowRadius: 8,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  transactionHeader: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 12,
  },
  transactionIconContainer: {
    width: 48,
    height: 48,
    borderRadius: 24,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 16,
  },
  transactionInfo: {
    flex: 1,
  },
  transactionTitle: {
    fontSize: 16,
    fontWeight: '600',
    color: '#1a237e',
    marginBottom: 4,
  },
  transactionSubtitle: {
    fontSize: 14,
    color: '#666666',
  },
  transactionAmountSuccess: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#28A745',
  },
  transactionAmountPending: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#DC3545',
  },
  transactionFooter: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  statusBadge: {
    paddingHorizontal: 12,
    paddingVertical: 6,
    borderRadius: 20,
    borderWidth: 1,
  },
  statusText: {
    fontSize: 12,
    fontWeight: '600',
  },
  transactionTime: {
    fontSize: 12,
    color: '#999999',
  },
  tripCard: {
    backgroundColor: '#ffffff',
    borderRadius: 20,
    padding: 24,
    marginBottom: 16,
    borderWidth: 1,
    borderColor: '#f0f0f0',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.08,
        shadowRadius: 12,
      },
      android: {
        elevation: 3,
      },
    }),
  },
  tripHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  tripId: {
    fontSize: 12,
    fontWeight: '600',
    color: '#666666',
    marginBottom: 4,
  },
  tripRoute: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#1a237e',
  },
  tripDetails: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: 16,
  },
  tripDetailItem: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  tripDetailText: {
    fontSize: 14,
    color: '#666666',
  },
  tripFooter: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingTop: 16,
    borderTopWidth: 1,
    borderTopColor: '#f0f0f0',
  },
  tripEarningsLabel: {
    fontSize: 12,
    fontWeight: '500',
    color: '#666666',
    marginBottom: 4,
  },
  tripEarningsAmount: {
    fontSize: 22,
    fontWeight: 'bold',
    color: '#28A745',
  },
  settlementCard: {
    backgroundColor: '#ffffff',
    borderRadius: 16,
    padding: 20,
    marginBottom: 12,
    borderWidth: 1,
    borderColor: '#f0f0f0',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.05,
        shadowRadius: 8,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  settlementHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  settlementIdContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  settlementId: {
    fontSize: 16,
    fontWeight: '600',
    color: '#1a237e',
  },
  settlementAmount: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#28A745',
  },
  settlementDetailItem: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 12,
    marginBottom: 16,
  },
  detailIconContainer: {
    width: 40,
    height: 40,
    borderRadius: 20,
    backgroundColor: '#f0f7ff',
    justifyContent: 'center',
    alignItems: 'center',
  },
  detailLabel: {
    fontSize: 12,
    fontWeight: '500',
    color: '#999999',
    marginBottom: 2,
  },
  detailValue: {
    fontSize: 16,
    fontWeight: '600',
    color: '#1a237e',
  },
  actionButtonsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    paddingHorizontal: 16,
    marginBottom: 32,
    gap: 12,
  },
  actionButton: {
    flex: 1,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#ffffff',
    paddingVertical: 16,
    paddingHorizontal: 8,
    borderRadius: 12,
    borderWidth: 1,
    borderColor: '#f0f0f0',
    ...Platform.select({
      ios: {
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.05,
        shadowRadius: 8,
      },
      android: {
        elevation: 2,
      },
    }),
  },
  actionButtonText: {
    fontSize: 14,
    fontWeight: '600',
    color: '#666666',
    marginLeft: 8,
  },
  retryBackendButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#28A745',
    margin: 16,
    padding: 16,
    borderRadius: 12,
    gap: 8,
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
  retryBackendButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
});

export default PaymentScreen;