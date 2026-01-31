import { createContext, useState } from "react";

export const UserPhoneContext = createContext();

export const UserPhoneProvider = ({ children }) => {
  const [userPhone, setUserPhone] = useState(null);

  return (
    <UserPhoneContext.Provider value={{ userPhone, setUserPhone }}>
      {children}
    </UserPhoneContext.Provider>
  );
};