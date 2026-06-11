<script setup lang="ts">
import type { TrackingHistory } from '../services/api';
import { CircleDot } from 'lucide-vue-next';

defineProps<{
  history: TrackingHistory[];
}>();
</script>

<template>
  <div class="timeline-container card animate-fade-in">
    <h3 class="section-title">Tracking History</h3>
    
    <div v-if="history.length === 0" class="empty-state">
      No history available for this shipment yet.
    </div>

    <div v-else class="timeline">
      <div v-for="(item, index) in history" :key="index" class="timeline-item">
        <div class="timeline-indicator">
          <div class="line" v-if="index !== history.length - 1"></div>
          <div class="dot" :class="{ 'latest': index === 0 }">
            <CircleDot v-if="index === 0" :size="14" />
          </div>
        </div>
        
        <div class="timeline-content">
          <div class="date">{{ item.date }}</div>
          <div class="desc">{{ item.desc }}</div>
          <div class="location" v-if="item.location">{{ item.location }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.section-title {
  font-size: 1.1rem;
  margin-bottom: 24px;
  color: var(--color-text-secondary);
}

.empty-state {
  color: var(--color-text-secondary);
  font-style: italic;
  padding: 20px 0;
  text-align: center;
}

.timeline {
  display: flex;
  flex-direction: column;
}

.timeline-item {
  display: flex;
  gap: 16px;
  padding-bottom: 24px;
}

.timeline-item:last-child {
  padding-bottom: 0;
}

.timeline-indicator {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 24px;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background-color: var(--color-border);
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 4px;
}

.dot.latest {
  width: 24px;
  height: 24px;
  background-color: rgba(79, 70, 229, 0.2);
  color: var(--color-primary);
  margin-top: 0;
}

.line {
  position: absolute;
  top: 16px;
  bottom: -24px;
  width: 2px;
  background-color: var(--color-border);
  z-index: 1;
}

.latest + .line {
  top: 24px;
}

.timeline-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.date {
  font-family: var(--font-mono);
  font-size: 0.85rem;
  color: var(--color-text-secondary);
}

.desc {
  font-size: 1.05rem;
  color: var(--color-text-primary);
  line-height: 1.4;
}

.location {
  font-size: 0.9rem;
  color: var(--color-accent);
  display: inline-block;
  margin-top: 2px;
}
</style>
