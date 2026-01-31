import AsyncStorage from '@react-native-async-storage/async-storage';
import { useEffect  , useState} from "react";

export const TokenProvider = ({ children }) => {
  const [accessToken, SetAccessToken] = useState(null);
  const [refreshToken, SetRefreshToken] = useState(null);

  useEffect(() => {
    AsyncStorage.getItem('accessToken').then(SetAccessToken);
    AsyncStorage.getItem('refreshToken').then(SetRefreshToken);
  }, []);

  const saveAccessToken = async (token) => {
    SetAccessToken(token);
    await AsyncStorage.setItem('accessToken', token);
  };

  const saveRefreshToken = async (token) => {
    SetRefreshToken(token);
    await AsyncStorage.setItem('refreshToken', token);
  };

  return (
    <TokenContext.Provider
      value={{
        accessToken,
        refreshToken,
        SetAccessToken: saveAccessToken,
        SetRefreshToken: saveRefreshToken,
      }}
    >
      {children}
    </TokenContext.Provider>
  );
};
