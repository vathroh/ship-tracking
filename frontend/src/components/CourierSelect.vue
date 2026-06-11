<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { trackingService, type Courier } from '../services/api';

defineProps<{
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const couriers = ref<Courier[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    const response = await trackingService.getCouriers();
    if (response.status === 'success' && response.data) {
      couriers.value = response.data;
    }
  } catch (error) {
    console.error('Failed to fetch couriers:', error);
  } finally {
    isLoading.value = false;
  }
});
</script>

<template>
  <div class="courier-select">
    <label for="courier">Courier</label>
    <div class="select-wrapper">
      <select 
        id="courier"
        :value="modelValue" 
        @change="emit('update:modelValue', ($event.target as HTMLSelectElement).value)"
        class="neural-input"
        required
        :disabled="isLoading"
      >
        <option value="" disabled>{{ isLoading ? 'Loading couriers...' : 'Select a courier' }}</option>
        <option v-for="c in couriers" :key="c.code" :value="c.code">
          {{ c.name }}
        </option>
      </select>
    </div>
  </div>
</template>

<style scoped>
.courier-select {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

label {
  font-family: var(--font-mono);
  font-size: 0.85rem;
  color: var(--color-text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.select-wrapper {
  position: relative;
}

.neural-input {
  width: 100%;
  appearance: none;
  background-color: rgba(0, 0, 0, 0.02);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-control);
  padding: 12px 16px;
  color: var(--color-text-primary);
  font-family: var(--font-sans);
  font-size: 1rem;
  transition: all 0.2s ease;
  outline: none;
}

.neural-input option {
  background-color: #FFFFFF;
  color: var(--color-text-primary);
}

.neural-input:focus {
  border-color: var(--color-primary);
  background-color: #FFFFFF;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.15);
}

.select-wrapper::after {
  content: '▼';
  font-size: 0.8rem;
  color: var(--color-text-secondary);
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
}
</style>
