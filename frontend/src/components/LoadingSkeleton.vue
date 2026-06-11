<script setup lang="ts">
// Basic loading skeleton that pulses
</script>

<template>
  <div class="skeleton-container animate-fade-in">
    <div class="skeleton title"></div>
    <div class="skeleton text"></div>
    <div class="skeleton text short"></div>
    
    <div class="card skeleton-card mt-4">
      <div class="skeleton text mb-2"></div>
      <div class="skeleton text"></div>
    </div>
  </div>
</template>

<style scoped>
.skeleton-container {
  display: flex;
  flex-direction: column;
  gap: var(--space-base);
  width: 100%;
}

.title {
  height: 24px;
  width: 40%;
  margin-bottom: var(--space-base);
}

.text {
  height: 16px;
  width: 100%;
}

.short {
  width: 60%;
}

.skeleton-card {
  padding: var(--space-card-padding);
}

.mt-4 {
  margin-top: calc(var(--space-base) * 2);
}

.mb-2 {
  margin-bottom: var(--space-base);
}
</style>
