<script setup lang="ts">
import { AlertCircle } from 'lucide-vue-next';

defineProps<{
  message: string;
}>();
</script>

<template>
  <div class="error-state animate-fade-in card">
    <AlertCircle class="icon" :size="24" />
    <div class="content">
      <h3>Error Tracking Package</h3>
      <p>{{ message }}</p>
    </div>
  </div>
</template>

<style scoped>
.error-state {
  display: flex;
  align-items: flex-start;
  gap: var(--space-gap);
  border-left: 4px solid var(--color-error);
  background: rgba(239, 68, 68, 0.1);
}

.icon {
  color: var(--color-error);
  flex-shrink: 0;
}

h3 {
  color: var(--color-error);
  margin-bottom: 4px;
  font-size: 1.1rem;
}

p {
  color: var(--color-text-secondary);
  font-size: 0.95rem;
}
</style>
