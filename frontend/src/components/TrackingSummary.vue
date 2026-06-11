<script setup lang="ts">
import type { TrackingSummary } from '../services/api';
import { Package, Calendar, Activity, Info } from 'lucide-vue-next';

defineProps<{
  summary: TrackingSummary;
}>();
</script>

<template>
  <div class="summary-card card animate-fade-in">
    <div class="header">
      <div class="badge" :class="summary.status.toLowerCase()">
        {{ summary.status }}
      </div>
      <span class="courier-badge">{{ summary.courier.toUpperCase() }}</span>
    </div>
    
    <h2 class="awb">{{ summary.awb }}</h2>
    
    <div class="meta-grid">
      <div class="meta-item">
        <Package :size="16" class="meta-icon" />
        <div class="meta-text">
          <span class="label">Service</span>
          <span class="value">{{ summary.service }}</span>
        </div>
      </div>
      
      <div class="meta-item">
        <Calendar :size="16" class="meta-icon" />
        <div class="meta-text">
          <span class="label">Date</span>
          <span class="value">{{ summary.date }}</span>
        </div>
      </div>
      
      <div class="meta-item">
        <Activity :size="16" class="meta-icon" />
        <div class="meta-text">
          <span class="label">Last Status</span>
          <span class="value">{{ summary.desc }}</span>
        </div>
      </div>

      <div class="meta-item">
        <Info :size="16" class="meta-icon" />
        <div class="meta-text">
          <span class="label">Weight</span>
          <span class="value">{{ summary.weight || '-' }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.summary-card {
  position: relative;
  overflow: hidden;
}

/* Neural glow effect */
.summary-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--color-primary), var(--color-accent));
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.awb {
  font-family: var(--font-mono);
  font-size: 2rem;
  letter-spacing: 0.05em;
  margin-bottom: 24px;
  word-break: break-all;
}

.badge {
  padding: 4px 12px;
  border-radius: var(--radius-pill);
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  background: rgba(255, 255, 255, 0.1);
}

.badge.delivered {
  background: rgba(16, 185, 129, 0.15);
  color: #10B981;
}

.badge.on_process, .badge.on-process {
  background: rgba(245, 158, 11, 0.15);
  color: #F59E0B;
}

.courier-badge {
  font-family: var(--font-mono);
  font-size: 0.85rem;
  color: var(--color-text-secondary);
}

.meta-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
}

.meta-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.meta-icon {
  color: var(--color-accent);
  margin-top: 2px;
}

.meta-text {
  display: flex;
  flex-direction: column;
}

.label {
  font-size: 0.75rem;
  text-transform: uppercase;
  color: var(--color-text-secondary);
  margin-bottom: 2px;
}

.value {
  font-size: 0.95rem;
  font-weight: 500;
}

@media (min-width: 640px) {
  .meta-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
