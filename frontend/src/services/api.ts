import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export interface StandardResponse<T> {
  status: string;
  message?: string;
  data?: T;
}

export interface TrackingSummary {
  awb: string;
  courier: string;
  service: string;
  status: string;
  date: string;
  desc: string;
  amount: string;
  weight: string;
}

export interface TrackingDetail {
  origin: string;
  destination: string;
  shipper: string;
  receiver: string;
}

export interface TrackingHistory {
  date: string;
  desc: string;
  location: string;
}

export interface TrackingData {
  summary: TrackingSummary;
  detail: TrackingDetail;
  history: TrackingHistory[];
}

export interface SearchHistoryRecord {
  id: number;
  tracking_number: string;
  courier: string;
  searched_at: string;
}

export interface PaginatedHistory {
  data: SearchHistoryRecord[];
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface Courier {
  code: string;
  name: string;
}

export const trackingService = {
  trackShipment: async (courier: string, awb: string) => {
    const response = await api.get<StandardResponse<TrackingData>>('/tracking', {
      params: { courier, tracking_number: awb }
    });
    return response.data;
  },
  
  getHistory: async (page = 1, limit = 10, courier?: string) => {
    const params: Record<string, any> = { page, limit };
    if (courier) params.courier = courier;
    
    const response = await api.get<StandardResponse<PaginatedHistory>>('/history', { params });
    return response.data;
  },

  getCouriers: async () => {
    const response = await api.get<StandardResponse<Courier[]>>('/couriers');
    return response.data;
  }
};

export default api;
