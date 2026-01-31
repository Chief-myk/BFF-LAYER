import i18n from "i18next";
import { initReactI18next } from "react-i18next";

import en from "../constants/en.json";
import hi from "../constants/hi.json";

i18n
  .use(initReactI18next)
  .init({
    compatibilityJSON: "v3",
    resources: {
      en: { translation: en },
      hi: { translation: hi },
    },
    lng: "en",        
    fallbackLng: "en",
    interpolation: { escapeValue: false },
  });

export default i18n;
