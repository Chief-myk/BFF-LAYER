import { MaterialCommunityIcons } from '@expo/vector-icons';
import { router } from 'expo-router';
import {
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
  Image,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useState, useContext } from 'react';
import NotificationPopup from './notification';
import { UserTypeContext } from "@/context/UserTypeContext";
import i18n from "@/service/i18n";

export default function CustomHeader() {
  const [notificationsVisible, setNotificationsVisible] = useState(false);
  const { userType } = useContext(UserTypeContext);
  const [isEnglish, setIsEnglish] = useState(i18n.language === 'en');

  const handleClick = () => {
    router.push(
      // userType === "driver"
      //   ? "/(pages)/DriverProfile"
      //   : "/(pages)/BrokerProfile"
        "/(pages)/DriverProfile"
    );
  };

  const toggleLanguage = () => {
    const newLanguage = isEnglish ? 'hi' : 'en';
    i18n.changeLanguage(newLanguage);
    setIsEnglish(!isEnglish);
  };

  return (
    <>
      <SafeAreaView style={styles.container}>
        {/* LEFT */}
        <View style={styles.leftSection}>
          <Image
            source={require('@/assets/images/logo.png')}
            style={styles.logo}
            resizeMode="contain"
            accessibilityLabel="CIVI Shield AI Logo"
          />
        </View>

        {/* CENTER */}
        <View style={styles.centerSection}>
          <TouchableOpacity
            style={styles.verifyButton}
            activeOpacity={0.8}
            accessibilityRole="button"
            accessibilityLabel="Verify user"
          >
            <View style={styles.statusDot} />
            <Text
              style={styles.verifyText}
              allowFontScaling={false}
              numberOfLines={1}
            >
              Verify
            </Text>
          </TouchableOpacity>

          {userType === "driver" && (
            <TouchableOpacity
              style={styles.languageToggle}
              onPress={toggleLanguage}
              activeOpacity={0.8}
              accessibilityRole="switch"
              accessibilityState={{ checked: isEnglish }}
            >
              <View
                style={[
                  styles.toggleTrack,
                  isEnglish
                    ? styles.toggleTrackEnglish
                    : styles.toggleTrackHindi,
                ]}
              >
                <View
                  style={[
                    styles.toggleThumb,
                    isEnglish
                      ? styles.toggleThumbEnglish
                      : styles.toggleThumbHindi,
                  ]}
                >
                  <Text
                    style={styles.toggleText}
                    allowFontScaling={false}
                  >
                    {isEnglish ? 'EN' : 'HI'}
                  </Text>
                </View>
              </View>
            </TouchableOpacity>
          )}

        </View>

        {/* RIGHT */}
        <View style={styles.rightSection}>
          <TouchableOpacity
            style={styles.iconButton}
            onPress={() => setNotificationsVisible(true)}
            activeOpacity={0.7}
            accessibilityLabel="Notifications"
          >
            <MaterialCommunityIcons
              name="bell-outline"
              size={24}
              color="#666"
            />
            <View style={styles.notificationBadge} />
          </TouchableOpacity>

          <TouchableOpacity
            style={styles.iconButton}
            onPress={handleClick}
            activeOpacity={0.7}
            accessibilityLabel="Profile"
          >
            <MaterialCommunityIcons
              name="account-outline"
              size={24}
              color="#FF0000"
            />
          </TouchableOpacity>
        </View>
      </SafeAreaView>

      <NotificationPopup
        visible={notificationsVisible}
        onClose={() => setNotificationsVisible(false)}
      />
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: 16,
    paddingTop: 4,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    backgroundColor: "#fff",
    borderBottomWidth: 1,
    borderBottomColor: "#f0f0f0",
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.05,
    shadowRadius: 3,
    elevation: 3,
  },

  leftSection: {},

  logo: {
    width: 110,
    height: 20,
  },

  centerSection: {
    gap: 6,
    paddingHorizontal: 2,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "center",
    flex: 1,
    flexShrink: 1, // 🔒 overflow safety (no visual change)
  },

  rightSection: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-end",
  },

  verifyButton: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
    paddingHorizontal: 6,
    paddingVertical: 4,
    backgroundColor: "#fff5f5",
    borderRadius: 16,
    borderColor: "#fed7d7",
    borderWidth: 1,
  },

  statusDot: {
    width: 8,
    height: 8,
    borderRadius: 4,
    backgroundColor: "#e53e3e",
  },

  verifyText: {
    fontSize: 14,
    fontWeight: "600",
    color: "#e53e3e",
  },

  languageToggle: {
    padding: 2,
  },

  toggleTrack: {
    width: 58,
    height: 30,
    borderRadius: 16,
    justifyContent: 'center',
    borderWidth: 1,
  },

  toggleTrackEnglish: {
    backgroundColor: '#f0f0f0',
    borderColor: '#ddd',
  },

  toggleTrackHindi: {
    backgroundColor: '#f8f8f8',
    borderColor: '#ddd',
  },

  toggleThumb: {
    width: 24,
    height: 20,
    borderRadius: 14,
    position: 'absolute',
    justifyContent: 'center',
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 1,
    elevation: 2,
  },

  toggleThumbEnglish: {
    backgroundColor: '#e53e3e',
    left: 2,
  },

  toggleThumbHindi: {
    backgroundColor: '#2f855a',
    right: 2,
  },

  toggleText: {
    fontSize: 12,
    fontWeight: 'bold',
    color: '#fff',
  },

  iconButton: {
    paddingVertical: 8,
    paddingHorizontal: 4,
    borderRadius: 8,
    position: 'relative',
  },

  notificationBadge: {
    position: 'absolute',
    top: 6,
    right: 6,
    width: 6,
    height: 6,
    borderRadius: 3,
    backgroundColor: "#e53e3e",
  },
});

