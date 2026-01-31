import { 
  StyleSheet, 
  Text, 
  View, 
  TouchableOpacity, 
  ScrollView, 
  Modal,
  TouchableWithoutFeedback 
} from 'react-native';
import React from 'react';
import { MaterialCommunityIcons, Ionicons } from '@expo/vector-icons';

const NotificationPopup = ({ visible, onClose }) => {
  const notifications = [
    {
      id: '1',
      type: 'success',
      icon: 'truck-fast',
      title: 'New Load Available',
      description: 'A new load from Delhi → Jaipur has been posted',
      time: '2 min ago',
      color: '#D4EDDA'
    },
    {
      id: '2',
      type: 'payment',
      icon: 'cash-multiple',
      title: 'Payment Received',
      description: 'Commission of ₹2,500 received for load #TH001234',
      time: '1 hour ago',
      color: '#D4EDDA'
    },
    {
      id: '3',
      type: 'warning',
      icon: 'alert-circle',
      title: 'Order Update',
      description: 'Load #TH001235 status changed to In Transit',
      time: 'Today, 4:45 PM',
      color: '#FFF3CD'
    },
    {
      id: '4',
      type: 'info',
      icon: 'account-check',
      title: 'Account Verification',
      description: 'Your KYC documents are pending verification',
      time: 'Today, 2:30 PM',
      color: '#D1ECF1'
    },
    {
      id: '5',
      type: 'critical',
      icon: 'close-circle',
      title: 'Payment Failed',
      description: 'Withdrawal of ₹5,000 failed due to bank issues',
      time: 'Yesterday',
      color: '#F8D7DA'
    },
    {
      id: '6',
      type: 'success',
      icon: 'check-circle',
      title: 'Load Completed',
      description: 'Load #TH001236 has been successfully delivered',
      time: 'Yesterday',
      color: '#D4EDDA'
    }
  ];

  const getIconColor = (type) => {
    switch (type) {
      case 'success': return '#28A745';
      case 'payment': return '#28A745';
      case 'warning': return '#FFC107';
      case 'info': return '#17A2B8';
      case 'critical': return '#DC3545';
      default: return '#666';
    }
  };

  const NotificationItem = ({ notification }) => (
    <TouchableOpacity style={styles.notificationItem}>
      <View style={[styles.notificationContent, { backgroundColor: notification.color }]}>
        <View style={styles.notificationLeft}>
          <MaterialCommunityIcons 
            name={notification.icon} 
            size={20} 
            color={getIconColor(notification.type)} 
          />
        </View>
        
        <View style={styles.notificationMiddle}>
          <Text style={styles.notificationTitle}>{notification.title}</Text>
          <Text style={styles.notificationDescription}>
            {notification.description}
          </Text>
        </View>
        
        <View style={styles.notificationRight}>
          <Text style={styles.notificationTime}>{notification.time}</Text>
        </View>
      </View>
    </TouchableOpacity>
  );

  return (
    <Modal
      visible={visible}
      transparent={true}
      animationType="fade"
      onRequestClose={onClose}
    >
      <TouchableWithoutFeedback onPress={onClose}>
        <View style={styles.modalOverlay}>
          <TouchableWithoutFeedback>
            <View style={styles.modalContainer}>
              {/* Modal Header */}
              <View style={styles.modalHeader}>
                <Text style={styles.modalTitle}>Notifications</Text>
                <TouchableOpacity style={styles.closeButton} onPress={onClose}>
                  <Ionicons name="close" size={24} color="#666" />
                </TouchableOpacity>
              </View>

              {/* Notifications List */}
              <ScrollView 
                style={styles.notificationsList}
                showsVerticalScrollIndicator={false}
                contentContainerStyle={styles.scrollContent}
              >
                {notifications.map((notification) => (
                  <NotificationItem 
                    key={notification.id} 
                    notification={notification} 
                  />
                ))}
              </ScrollView>

              {/* Footer */}
              <View style={styles.modalFooter}>
                <TouchableOpacity style={styles.clearAllButton}>
                  <Text style={styles.clearAllText}>Clear All Notifications</Text>
                </TouchableOpacity>
              </View>
            </View>
          </TouchableWithoutFeedback>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

export default NotificationPopup;

const styles = StyleSheet.create({
  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  modalContainer: {
    backgroundColor: '#fff',
    borderRadius: 16,
    width: '100%',
    maxHeight: '80%',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 4,
    },
    shadowOpacity: 0.25,
    shadowRadius: 12,
    elevation: 8,
    borderWidth: 1,
    borderColor: '#f0f0f0',
    overflow: 'hidden',
  },
  modalHeader: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 20,
    paddingVertical: 16,
    borderBottomWidth: 1,
    borderBottomColor: '#e5e5e5',
    backgroundColor: '#fff',
  },
  modalTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#1a1a1a',
    textAlign: 'center',
    flex: 1,
  },
  closeButton: {
    padding: 4,
    position: 'absolute',
    right: 16,
  },
  notificationsList: {
    flex: 1,
  },
  scrollContent: {
    padding: 16,
  },
  notificationItem: {
    marginBottom: 12,
    borderRadius: 12,
    overflow: 'hidden',
  },
  notificationContent: {
    flexDirection: 'row',
    alignItems: 'flex-start',
    padding: 16,
    borderRadius: 12,
    borderWidth: 1,
    borderColor: 'rgba(0, 0, 0, 0.05)',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 1,
    },
    shadowOpacity: 0.1,
    shadowRadius: 3,
    elevation: 2,
  },
  notificationLeft: {
    marginRight: 12,
    paddingTop: 2,
  },
  notificationMiddle: {
    flex: 1,
    marginRight: 12,
  },
  notificationTitle: {
    fontSize: 16,
    fontWeight: '600',
    color: '#1a1a1a',
    marginBottom: 4,
  },
  notificationDescription: {
    fontSize: 14,
    color: '#666',
    lineHeight: 18,
  },
  notificationRight: {
    alignItems: 'flex-end',
    minWidth: 80,
  },
  notificationTime: {
    fontSize: 12,
    color: '#666',
    fontWeight: '500',
  },
  modalFooter: {
    padding: 16,
    borderTopWidth: 1,
    borderTopColor: '#e5e5e5',
    backgroundColor: '#f8f9fa',
  },
  clearAllButton: {
    alignItems: 'center',
  },
  clearAllText: {
    fontSize: 14,
    color: '#666',
    fontWeight: '500',
    textDecorationLine: 'underline',
  },
});