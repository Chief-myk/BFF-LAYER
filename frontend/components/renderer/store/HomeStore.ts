import { create } from 'zustand';

interface HomeState {
  // State
  isTripStarted: boolean;
  locationSharing: boolean;
  tripStatus: string;
  documentsUploaded: {
    eWayBill: boolean;
    invoice: boolean;
    vehicleRC: boolean;
    driverLicense: boolean;
    insurance: boolean;
    pollutionCert: boolean;
  };
  activeTrip: any;
  uploadProgress: number;
  showUploadModal: boolean;
  
  // Actions
  setTripStarted: (value: boolean) => void;
  setLocationSharing: (value: boolean) => void;
  setTripStatus: (status: string) => void;
  updateDocument: (type: string, uploaded: boolean) => void;
  setUploadProgress: (progress: number) => void;
  setShowUploadModal: (show: boolean) => void;
  
  // API Actions
  startTrip: () => Promise<void>;
  uploadDocument: (documentType: string) => Promise<void>;
  updateTripStatus: (status: string) => Promise<void>;
  uploadPOD: () => Promise<void>;
  fetchHomeData: () => Promise<void>;
}

export const useHomeStore = create<HomeState>((set, get) => ({
  // Initial state
  isTripStarted: false,
  locationSharing: false,
  tripStatus: 'not_started',
  documentsUploaded: {
    eWayBill: false,
    invoice: false,
    vehicleRC: false,
    driverLicense: false,
    insurance: false,
    pollutionCert: false
  },
  activeTrip: null,
  uploadProgress: 0,
  showUploadModal: false,

  // State setters
  setTripStarted: (value) => set({ isTripStarted: value }),
  setLocationSharing: (value) => set({ locationSharing: value }),
  setTripStatus: (status) => set({ tripStatus: status }),
  updateDocument: (type, uploaded) => 
    set((state) => ({
      documentsUploaded: {
        ...state.documentsUploaded,
        [type]: uploaded
      }
    })),
  setUploadProgress: (progress) => set({ uploadProgress: progress }),
  setShowUploadModal: (show) => set({ showUploadModal: show }),

  // API Actions
  fetchHomeData: async () => {
    try {
      const response = await fetch('http://192.168.1.3:8080/bff/driver/home');
      const data = await response.json();
      
      // Update state with backend data
      const documents = data.data.documentsUploaded || {};
      set({
        isTripStarted: data.data.isTripStarted,
        locationSharing: data.data.locationSharing,
        tripStatus: data.data.tripStatus,
        documentsUploaded: {
          eWayBill: documents.eWayBill || false,
          invoice: documents.invoice || false,
          vehicleRC: documents.vehicleRC || false,
          driverLicense: documents.driverLicense || false,
          insurance: documents.insurance || false,
          pollutionCert: documents.pollutionCert || false
        },
        activeTrip: data.data.activeTrip
      });
    } catch (error) {
      console.error('Error fetching home data:', error);
    }
  },

  startTrip: async () => {
    const state = get();
    
    // Check documents first
    if (!state.documentsUploaded.eWayBill || !state.documentsUploaded.invoice) {
      set({ showUploadModal: true });
      return;
    }

    try {
      const response = await fetch('http://192.168.1.3:8080/bff/driver/home/action', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          action: 'START_TRIP',
          data: {}
        })
      });

      const result = await response.json();
      
      if (result.status === 'success') {
        set({
          isTripStarted: true,
          locationSharing: true,
          tripStatus: 'reached_origin'
        });
      }
    } catch (error) {
      console.error('Error starting trip:', error);
    }
  },

  uploadDocument: async (documentType: string) => {
    try {
      const response = await fetch('http://192.168.1.3:8080/bff/driver/home/action', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          action: 'UPLOAD_DOCUMENT',
          data: { documentType }
        })
      });

      const result = await response.json();
      
      if (result.status === 'success') {
        get().updateDocument(documentType as 'eWayBill' | 'invoice', true);
        
        // Check if all documents uploaded
        const state = get();
        if (state.documentsUploaded.eWayBill && state.documentsUploaded.invoice) {
          set({ showUploadModal: false });
        }
      }
    } catch (error) {
      console.error('Error uploading document:', error);
    }
  },

  updateTripStatus: async (status: string) => {
    try {
      const response = await fetch('http://192.168.1.3:8080/bff/driver/home/action', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          action: 'UPDATE_STATUS',
          data: { status }
        })
      });

      const result = await response.json();
      
      if (result.status === 'success') {
        set({ tripStatus: status });
        
        if (status === 'completed') {
          set({ locationSharing: false });
        }
      }
    } catch (error) {
      console.error('Error updating status:', error);
    }
  },

  uploadPOD: async () => {
    // Simulate upload progress
    let progress = 0;
    const interval = setInterval(() => {
      progress += 10;
      set({ uploadProgress: progress });
      
      if (progress >= 100) {
        clearInterval(interval);
        // Call backend API to confirm upload
        // router.push("/(pages)/TripCompleted");
      }
    }, 200);
  }
}));