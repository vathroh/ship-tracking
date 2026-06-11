import { defineStore } from 'pinia';
import { ref } from 'vue';
import { trackingService, type TrackingData } from '../services/api';

export const useTrackingStore = defineStore('tracking', () => {
  const currentTracking = ref<TrackingData | null>(null);
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  const trackPackage = async (courier: string, awb: string) => {
    isLoading.value = true;
    error.value = null;
    currentTracking.value = null;

    try {
      const response = await trackingService.trackShipment(courier, awb);
      if (response.status === 'success' && response.data) {
        currentTracking.value = response.data;
      } else {
        error.value = response.message || 'Failed to track package';
      }
    } catch (err: any) {
      error.value = err.response?.data?.message || err.message || 'An unexpected error occurred';
    } finally {
      isLoading.value = false;
    }
  };

  const clearTracking = () => {
    currentTracking.value = null;
    error.value = null;
  };

  return {
    currentTracking,
    isLoading,
    error,
    trackPackage,
    clearTracking
  };
});
