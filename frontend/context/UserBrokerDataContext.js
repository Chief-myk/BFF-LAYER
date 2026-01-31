import { createContext, useState } from "react";

export const UserBrokerDataContext = createContext();

export const UserBrokerDataProvider = ({ children }) => {
    const [userFirstName, setUserFirstName] = useState(null);
    const [userLastName, setUserLastName] = useState(null);
    const [userPhone, setUserPhone] = useState(null);
    const [userEmail, setUserEmail] = useState(null);
    const [userCompanyName, setUserCompanyName] = useState(null);
    const [userGSTNumber, setUserGSTNumber] = useState(null);
    const [userCity, setUserCity] = useState(null);
    const [userRoutes, setUserRoutes] = useState([]);
    const [userPanCardImage, setUserPanCardImage] = useState(null);
    const [userAadharCardImage, setUserAadharCardImage] = useState(null);
    const [userBrokerLicenseCardImage, setUserBrokerLicenseCardImage] = useState(null);
    const [userCompanyDocumentImage, setUserCompanyDocumentImage] = useState(null);

    const value = {
        // Personal Info
        userFirstName, setUserFirstName,
        userLastName, setUserLastName,
        userPhone, setUserPhone,
        userEmail, setUserEmail,
        
        // Company Info
        userCompanyName, setUserCompanyName,
        userGSTNumber, setUserGSTNumber,
        userCity, setUserCity,
        userRoutes, setUserRoutes,
        
        // Documents
        userPanCardImage, setUserPanCardImage,
        userAadharCardImage, setUserAadharCardImage,
        userBrokerLicenseCardImage, setUserBrokerLicenseCardImage,
        userCompanyDocumentImage, setUserCompanyDocumentImage,
    };

    return (
        <UserBrokerDataContext.Provider value={value}>
            {children}
        </UserBrokerDataContext.Provider>
    );
};