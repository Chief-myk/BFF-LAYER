/**
 * 🚀 Enhanced Global Theme for Expo App (Light + Dark)
 * Includes:
 * - Full semantic color system
 * - Tab bar + header colors
 * - Spacing, radius, shadows, typography
 * - Platform-specific fonts
 */

import { Platform } from "react-native";

const tintLight = "#0a7ea4";
const tintDark = "#4AB8FF";

export const Colors = {
  light: {
    // Base
    text: "#11181C",
    background: "#FFFFFF",
    surface: "#F7F9FA",
    surfaceSecondary: "#ECEDEE",

    // Branding / Primary
    tint: tintLight,
    primary: tintLight,

    // Icons
    icon: "#687076",
    tabIconDefault: "#687076",
    tabIconSelected: tintLight,

    // Borders
    border: "#E5E7EB",
    borderStrong: "#D1D5DB",
    tabBarBorder: "#E5E7EB",

    // Tab / Navigation
    tabBarBackground: "#FFFFFF",
    navBarBackground: "#FFFFFF",
    navBarBorder: "#E5E5E5",

    // Status colors
    success: "#22C55E",
    warning: "#F59E0B",
    danger: "#EF4444",
    info: "#3B82F6",

    // Shadows
    shadow: "rgba(0,0,0,0.12)",
    shadowStrong: "rgba(0,0,0,0.20)",
  },

  dark: {
    // Base
    text: "#ECEDEE",
    background: "#151718",
    surface: "#1A1D1F",
    surfaceSecondary: "#222527",

    // Branding / Primary
    tint: tintDark,
    primary: tintDark,

    // Icons
    icon: "#9BA1A6",
    tabIconDefault: "#9BA1A6",
    tabIconSelected: tintDark,

    // Borders
    border: "#2D2F31",
    borderStrong: "#3B3D3F",
    tabBarBorder: "#2A2C2E",

    // Tab / Navigation
    tabBarBackground: "#1A1D1F",
    navBarBackground: "#1A1D1F",
    navBarBorder: "#2A2C2E",

    // Status colors
    success: "#4ADE80",
    warning: "#FBBF24",
    danger: "#F87171",
    info: "#60A5FA",

    // Shadows
    shadow: "rgba(0,0,0,0.45)",
    shadowStrong: "rgba(0,0,0,0.60)",
  },
};

//
// 🎨 Typography Scale
//
export const Typography = {
  fontSizes: {
    xs: 10,
    sm: 12,
    base: 14,
    md: 16,
    lg: 18,
    xl: 20,
    "2xl": 24,
    "3xl": 28,
    "4xl": 34,
  },

  fontWeights: {
    thin: "100",
    light: "300",
    regular: "400",
    medium: "500",
    semibold: "600",
    bold: "700",
    extrabold: "800",
  },
};

//
// 📦 Spacing Scale
//
export const Spacing = {
  xxs: 2,
  xs: 4,
  sm: 8,
  md: 12,
  lg: 16,
  xl: 24,
  "2xl": 32,
  "3xl": 40,
};

//
// 🔲 Radius Scale
//
export const Radius = {
  sm: 4,
  md: 8,
  lg: 12,
  xl: 16,
  "2xl": 24,
  full: 999,
};

//
// 🌫 Global Shadows (iOS + Android)
//
export const Shadows = {
  light: {
    default: {
      shadowColor: "#000",
      shadowOpacity: 0.08,
      shadowRadius: 6,
      shadowOffset: { width: 0, height: 2 },
      elevation: 4,
    },
    strong: {
      shadowColor: "#000",
      shadowOpacity: 0.15,
      shadowRadius: 12,
      shadowOffset: { width: 0, height: 4 },
      elevation: 10,
    },
  },

  dark: {
    default: {
      shadowColor: "#000",
      shadowOpacity: 0.4,
      shadowRadius: 6,
      shadowOffset: { width: 0, height: 2 },
      elevation: 4,
    },
    strong: {
      shadowColor: "#000",
      shadowOpacity: 0.6,
      shadowRadius: 12,
      shadowOffset: { width: 0, height: 6 },
      elevation: 12,
    },
  },
};

//
// 🅰 Platform-Specific Fonts
//
export const Fonts = Platform.select({
  ios: {
    sans: "system-ui",
    serif: "ui-serif",
    rounded: "ui-rounded",
    mono: "ui-monospace",
  },
  android: {
    sans: "normal",
    serif: "serif",
    rounded: "normal",
    mono: "monospace",
  },
  web: {
    sans: "system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif",
    serif: "Georgia, 'Times New Roman', serif",
    rounded: "'SF Pro Rounded', 'Hiragino Maru Gothic ProN', Meiryo, sans-serif",
    mono: "SFMono-Regular, Menlo, Consolas, 'Courier New', monospace",
  },
  default: {
    sans: "normal",
    serif: "serif",
    rounded: "normal",
    mono: "monospace",
  },
});
