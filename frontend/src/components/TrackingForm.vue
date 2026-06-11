<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { Search } from 'lucide-vue-next';
import CourierSelect from './CourierSelect.vue';

const emit = defineEmits<{
  (e: 'submit', payload: { courier: string; awb: string }): void;
}>();

const route = useRoute();

const courier = ref((route.query.courier as string) || '');
const awb = ref((route.query.awb as string) || '');

const handleSubmit = () => {
  if (courier.value && awb.value) {
    emit('submit', { courier: courier.value, awb: awb.value });
  }
};
</script>

<template>
  <form @submit.prevent="handleSubmit" class="tracking-form card">
    <div class="form-header">
      <h2 class="title">Trace Shipment</h2>
      <p class="subtitle">Enter your tracking number and courier to get real-time updates.</p>
    </div>

    <div class="inputs">
      <CourierSelect v-model="courier" />
      
      <div class="input-group">
        <label for="awb">Airway Bill (AWB)</label>
        <div class="input-wrapper">
          <input 
            id="awb"
            v-model="awb" 
            type="text" 
            placeholder="e.g. JP123456789"
            class="neural-input"
            required
          />
        </div>
      </div>
    </div>

    <button type="submit" class="submit-btn" :disabled="!courier || !awb">
      <Search :size="18" />
      Track Package
    </button>
  </form>
</template>

<style scoped>
.tracking-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
  max-width: 500px;
  width: 100%;
  margin: 0 auto;
}

.form-header {
  text-align: center;
}

.title {
  font-family: var(--font-sans);
  font-size: 1.5rem;
  margin-bottom: 8px;
}

.subtitle {
  color: var(--color-text-secondary);
  font-size: 0.95rem;
}

.inputs {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.input-group {
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

.input-wrapper {
  position: relative;
}

.neural-input {
  width: 100%;
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

.neural-input:focus {
  border-color: var(--color-primary);
  background-color: #FFFFFF;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.15);
}

.submit-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-accent) 100%);
  color: white;
  padding: 14px 24px;
  border-radius: var(--radius-control);
  font-weight: 600;
  font-size: 1rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  margin-top: 8px;
  box-shadow: 0 4px 14px rgba(99, 102, 241, 0.4);
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-2px) scale(1.02);
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.6);
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

@media (min-width: 640px) {
  .inputs {
    flex-direction: row;
    align-items: flex-end;
  }
  
  .courier-select, .input-group {
    flex: 1;
  }
}
</style>
