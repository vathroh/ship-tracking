import { defineStore } from 'pinia';
import { ref } from 'vue';
import { trackingService, type PaginatedHistory } from '../services/api';

export const useHistoryStore = defineStore('history', () => {
  const historyData = ref<PaginatedHistory | null>(null);
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  const fetchHistory = async (page = 1, limit = 10, courier?: string) => {
    isLoading.value = true;
    error.value = null;

    try {
      const response = await trackingService.getHistory(page, limit, courier);
      if (response.status === 'success' && response.data) {
        historyData.value = response.data;
      } else {
        error.value = response.message || 'Failed to fetch history';
      }
    } catch (err: any) {
      error.value = err.response?.data?.message || err.message || 'An unexpected error occurred';
    } finally {
      isLoading.value = false;
    }
  };

  return {
    historyData,
    isLoading,
    error,
    fetchHistory
  };
});
