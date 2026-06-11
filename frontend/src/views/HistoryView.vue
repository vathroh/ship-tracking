<script setup lang="ts">
import { onMounted } from 'vue';
import { useHistoryStore } from '../store/history';
import { Clock, ExternalLink } from 'lucide-vue-next';
import { useRouter } from 'vue-router';
import LoadingSkeleton from '../components/LoadingSkeleton.vue';
import ErrorState from '../components/ErrorState.vue';

const store = useHistoryStore();
const router = useRouter();

onMounted(() => {
  store.fetchHistory();
});

const reTrack = (courier: string, awb: string) => {
  router.push({ name: 'tracking', query: { courier, awb } });
};

const nextPage = () => {
  if (store.historyData && store.historyData.page < store.historyData.total_pages) {
    store.fetchHistory(store.historyData.page + 1);
  }
};

const prevPage = () => {
  if (store.historyData && store.historyData.page > 1) {
    store.fetchHistory(store.historyData.page - 1);
  }
};
</script>

<template>
  <div class="history-view">
    <div class="header">
      <h1 class="page-title">Search History</h1>
      <p class="subtitle">Your recent tracking queries across all couriers.</p>
    </div>

    <LoadingSkeleton v-if="store.isLoading" />
    
    <ErrorState v-else-if="store.error" :message="store.error" />

    <div v-else-if="store.historyData && store.historyData.data.length > 0" class="history-content">
      <div class="history-list">
        <div v-for="item in store.historyData.data" :key="item.id" class="history-card card animate-fade-in" @click="reTrack(item.courier, item.tracking_number)">
          <div class="card-left">
            <div class="icon-bg">
              <Clock :size="20" class="clock-icon" />
            </div>
            <div class="info">
              <h3 class="awb">{{ item.tracking_number }}</h3>
              <div class="meta">
                <span class="courier">{{ item.courier.toUpperCase() }}</span>
                <span class="dot">•</span>
                <span class="date">{{ new Date(item.searched_at).toLocaleString() }}</span>
              </div>
            </div>
          </div>
          <ExternalLink :size="18" class="nav-icon" />
        </div>
      </div>

      <div class="pagination">
        <button 
          @click="prevPage" 
          :disabled="store.historyData.page === 1"
          class="page-btn"
        >
          Previous
        </button>
        <span class="page-info">
          Page {{ store.historyData.page }} of {{ store.historyData.total_pages }}
        </span>
        <button 
          @click="nextPage" 
          :disabled="store.historyData.page === store.historyData.total_pages"
          class="page-btn"
        >
          Next
        </button>
      </div>
    </div>

    <div v-else class="empty-state card">
      <Clock :size="48" class="empty-icon" />
      <h3>No History Found</h3>
      <p>You haven't tracked any packages yet.</p>
      <button @click="router.push('/')" class="primary-btn mt-4">Track a Package</button>
    </div>
  </div>
</template>

<style scoped>
.history-view {
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.header {
  margin-bottom: 16px;
}

.page-title {
  font-size: 2rem;
  margin-bottom: 8px;
}

.subtitle {
  color: var(--color-text-secondary);
}

.history-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.history-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 24px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.history-card:hover {
  border-color: var(--color-primary);
  transform: translateX(4px);
}

.history-card:hover .nav-icon {
  color: var(--color-primary);
  transform: translateX(2px) translateY(-2px);
}

.card-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.icon-bg {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.05);
  display: flex;
  align-items: center;
  justify-content: center;
}

.clock-icon {
  color: var(--color-text-secondary);
}

.awb {
  font-family: var(--font-mono);
  font-size: 1.1rem;
  margin-bottom: 4px;
}

.meta {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.85rem;
  color: var(--color-text-secondary);
}

.courier {
  font-weight: 600;
  color: var(--color-text-primary);
}

.dot {
  opacity: 0.5;
}

.nav-icon {
  color: var(--color-border);
  transition: all 0.2s ease;
}

.pagination {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid var(--color-border);
}

.page-info {
  font-size: 0.9rem;
  color: var(--color-text-secondary);
}

.page-btn {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  color: var(--color-text-primary);
  padding: 8px 16px;
  border-radius: var(--radius-control);
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.page-btn:hover:not(:disabled) {
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 64px 24px;
}

.empty-icon {
  color: var(--color-border);
  margin-bottom: 16px;
}

.empty-state h3 {
  margin-bottom: 8px;
}

.empty-state p {
  color: var(--color-text-secondary);
  margin-bottom: 24px;
}

.primary-btn {
  background-color: var(--color-primary);
  color: white;
  padding: 12px 24px;
  border-radius: var(--radius-control);
  font-weight: 500;
  transition: all 0.2s ease;
}

.primary-btn:hover {
  background-color: #4338ca;
}
</style>
