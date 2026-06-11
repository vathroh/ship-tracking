<script setup lang="ts">
import type { TrackingDetail } from '../services/api';
import { MapPin, User } from 'lucide-vue-next';

defineProps<{
  detail: TrackingDetail;
}>();
</script>

<template>
  <div class="detail-card card animate-fade-in">
    <h3 class="section-title">Shipment Details</h3>
    
    <div class="detail-grid">
      <div class="detail-group origin">
        <div class="icon-wrapper">
          <MapPin :size="18" />
        </div>
        <div class="info">
          <span class="label">Origin</span>
          <span class="value">{{ detail.origin || 'Unknown' }}</span>
          <span class="sub-value"><User :size="12" class="inline-icon"/> {{ detail.shipper || 'Unknown' }}</span>
        </div>
      </div>

      <div class="connector"></div>

      <div class="detail-group destination">
        <div class="icon-wrapper accent">
          <MapPin :size="18" />
        </div>
        <div class="info">
          <span class="label">Destination</span>
          <span class="value">{{ detail.destination || 'Unknown' }}</span>
          <span class="sub-value"><User :size="12" class="inline-icon"/> {{ detail.receiver || 'Unknown' }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.section-title {
  font-size: 1.1rem;
  margin-bottom: 20px;
  color: var(--color-text-secondary);
}

.detail-grid {
  display: flex;
  flex-direction: column;
  gap: 16px;
  position: relative;
}

.detail-group {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  position: relative;
  z-index: 2;
}

.icon-wrapper {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-secondary);
}

.icon-wrapper.accent {
  border-color: var(--color-accent);
  color: var(--color-accent);
  background: rgba(6, 182, 212, 0.1);
}

.info {
  display: flex;
  flex-direction: column;
}

.label {
  font-size: 0.75rem;
  text-transform: uppercase;
  color: var(--color-text-secondary);
  margin-bottom: 4px;
}

.value {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--color-text-primary);
  margin-bottom: 4px;
}

.sub-value {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 0.85rem;
  color: var(--color-text-secondary);
}

.inline-icon {
  opacity: 0.7;
}

.connector {
  position: absolute;
  left: 19px;
  top: 40px;
  bottom: 40px;
  width: 2px;
  background: var(--color-border);
  z-index: 1;
}

@media (min-width: 640px) {
  .detail-grid {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }

  .connector {
    top: 19px;
    left: 40px;
    right: 40px;
    bottom: auto;
    height: 2px;
    width: auto;
  }
}
</style>
