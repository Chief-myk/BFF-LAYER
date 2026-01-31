import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";
import AsyncStorage from "@react-native-async-storage/async-storage";

type FormState = {
  values: Record<string, string>;
  setValue: (key: string, value: string) => void;
  reset: () => void;
  // Special method to set phone number
  setPhone: (phone: string) => void;
  getPhone: () => string;
};

export const useFormStore = create<FormState>()(
  persist(
    (set, get) => ({
      values: {},
      setValue: (key, value) =>
        set((state) => ({
          values: { ...state.values, [key]: value },
        })),
      reset: () => set({ values: {} }),
      setPhone: (phone) =>
        set((state) => ({
          values: { ...state.values, phone: phone },
        })),
      getPhone: () => get().values.phone || "",
    }),
    {
      name: "form-storage",
      storage: createJSONStorage(() => AsyncStorage),
    }
  )
);