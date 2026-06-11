<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useRoute } from 'vue-router';
import { useTrackingStore } from '../store/tracking';
import TrackingForm from '../components/TrackingForm.vue';
import TrackingSummary from '../components/TrackingSummary.vue';
import TrackingDetail from '../components/TrackingDetail.vue';
import TrackingTimeline from '../components/TrackingTimeline.vue';
import LoadingSkeleton from '../components/LoadingSkeleton.vue';
import ErrorState from '../components/ErrorState.vue';

const store = useTrackingStore();
const route = useRoute();

const handleTrack = ({ courier, awb }: { courier: string; awb: string }) => {
  store.trackPackage(courier, awb);
};

onMounted(() => {
  const courier = route.query.courier as string;
  const awb = route.query.awb as string;
  
  if (courier && awb) {
    handleTrack({ courier, awb });
  }
});

onUnmounted(() => {
  store.clearTracking();
});
</script>

<template>
  <div class="tracking-view">
    <div class="hero-section">
      <TrackingForm @submit="handleTrack" />
    </div>

    <div class="results-section">
      <LoadingSkeleton v-if="store.isLoading" />
      
      <ErrorState v-else-if="store.error" :message="store.error" />
      
      <div v-else-if="store.currentTracking" class="tracking-results">
        <TrackingSummary :summary="store.currentTracking.summary" />
        <TrackingDetail :detail="store.currentTracking.detail" />
        <TrackingTimeline :history="store.currentTracking.history" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.tracking-view {
  display: flex;
  flex-direction: column;
  gap: var(--space-section-padding);
}

.results-section {
  max-width: 800px;
  width: 100%;
  margin: 0 auto;
}

.tracking-results {
  display: flex;
  flex-direction: column;
  gap: var(--space-card-padding);
}
</style>
