import { createContext, useState } from "react";

export const UserDriverDataContext = createContext();

export const UserDriverDataProvider = ({ children }) => {
    const [userPhone, setUserPhone] = useState(null);
    const [userFirstName, setUserFirstName] = useState(null);
    const [userLastName, setUserLastName] = useState(null);
    const [userVehicleCategory, setUserVehicleCategory] = useState(null);
    const [userVehicleName, setUserVehicleName] = useState(null);
    const [userRoutes, setUserRoutes] = useState([]);
    const [userDriverLicenseCardImage, setUserDriverLicenseCardImage] = useState(null);
    const [userDriverImage, setUserDriverImage] = useState(null);
    const [userRCImage, setUserRCImage] = useState(null);
    const [userFitnessCertificateImage, setUserFitnessCertificateImage] = useState(null);
    const [userNationalPermitImage, setUserNationalPermitImage] = useState(null);
    const [userStatePermitImage, setUserStatePermitImage] = useState(null);
    const [panNumber, setPanNumber] = useState('');
    const [aadhaarNumber, setAadhaarNumber] = useState('');
    const [accountNumber, setAccountNumber] = useState('');
    const [ifscCode, setIfscCode] = useState('');
    const [accountHolderName, setAccountHolderName] = useState('');

    const [panVerified, setPanVerified] = useState(false);
    const [aadhaarVerified, setAadhaarVerified] = useState(false);
    const [accountVerified, setAccountVerified] = useState(false);

    const value = {
        // Personal Info
        userFirstName, setUserFirstName,
        userLastName, setUserLastName,
        userPhone, setUserPhone,
        // Vehicle Info        
        userVehicleCategory, setUserVehicleCategory,
        userVehicleName, setUserVehicleName,
        userRoutes, setUserRoutes,

        // Documents
        userDriverImage, setUserDriverImage,
        userDriverLicenseCardImage, setUserDriverLicenseCardImage,
        userRCImage, setUserRCImage,
        userFitnessCertificateImage, setUserFitnessCertificateImage,
        userNationalPermitImage, setUserNationalPermitImage,
        userStatePermitImage, setUserStatePermitImage,

        // Details
        panNumber, setPanNumber,
        aadhaarNumber, setAadhaarNumber,
        accountNumber, setAccountNumber,
        ifscCode, setIfscCode,
        accountHolderName, setAccountHolderName,
  
        panVerified, setPanVerified,
        aadhaarVerified, setAadhaarVerified,
        accountVerified, setAccountVerified,
    };

    return (
        <UserDriverDataContext.Provider value={value}>
            {children}
        </UserDriverDataContext.Provider>
    );
};