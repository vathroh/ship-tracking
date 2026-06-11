<script setup lang="ts">
import { BarChart3, TrendingUp, Package } from 'lucide-vue-next';
</script>

<template>
  <div class="dashboard-view">
    <div class="header">
      <h1 class="page-title">Analytics Dashboard</h1>
      <p class="subtitle">Platform usage and tracking volume statistics.</p>
    </div>

    <div class="stats-grid">
      <div class="stat-card card">
        <div class="stat-header">
          <span class="stat-title">Total API Requests</span>
          <div class="icon-wrapper primary">
            <Activity :size="20" />
          </div>
        </div>
        <div class="stat-value">12,458</div>
        <div class="stat-trend positive">
          <TrendingUp :size="16" />
          <span>+14% this week</span>
        </div>
      </div>

      <div class="stat-card card">
        <div class="stat-header">
          <span class="stat-title">Unique Packages</span>
          <div class="icon-wrapper accent">
            <Package :size="20" />
          </div>
        </div>
        <div class="stat-value">8,192</div>
        <div class="stat-trend positive">
          <TrendingUp :size="16" />
          <span>+5% this week</span>
        </div>
      </div>
    </div>

    <div class="charts-section">
      <div class="chart-card card">
        <div class="card-header">
          <h3>Popular Couriers</h3>
          <BarChart3 class="header-icon" :size="20" />
        </div>
        
        <div class="mock-chart">
          <div class="bar-row">
            <span class="bar-label">JNT</span>
            <div class="bar-container">
              <div class="bar" style="width: 85%; background: var(--color-primary)"></div>
            </div>
            <span class="bar-value">4.2k</span>
          </div>
          
          <div class="bar-row">
            <span class="bar-label">JNE</span>
            <div class="bar-container">
              <div class="bar" style="width: 65%; background: var(--color-accent)"></div>
            </div>
            <span class="bar-value">3.1k</span>
          </div>
          
          <div class="bar-row">
            <span class="bar-label">SICEPAT</span>
            <div class="bar-container">
              <div class="bar" style="width: 45%; background: var(--color-border)"></div>
            </div>
            <span class="bar-value">2.4k</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-view {
  display: flex;
  flex-direction: column;
  gap: 32px;
  width: 100%;
}

.header {
  margin-bottom: 8px;
}

.page-title {
  font-size: 2rem;
  margin-bottom: 8px;
}

.subtitle {
  color: var(--color-text-secondary);
}

.stats-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 24px;
}

.stat-card {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.stat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stat-title {
  font-size: 0.9rem;
  color: var(--color-text-secondary);
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.icon-wrapper {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-wrapper.primary {
  background: rgba(79, 70, 229, 0.1);
  color: var(--color-primary);
}

.icon-wrapper.accent {
  background: rgba(6, 182, 212, 0.1);
  color: var(--color-accent);
}

.stat-value {
  font-family: var(--font-mono);
  font-size: 2.5rem;
  font-weight: 600;
  line-height: 1;
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.85rem;
}

.stat-trend.positive {
  color: #10B981;
}

.chart-card {
  padding: 32px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
}

.header-icon {
  color: var(--color-text-secondary);
}

.mock-chart {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.bar-row {
  display: flex;
  align-items: center;
  gap: 16px;
}

.bar-label {
  font-family: var(--font-mono);
  font-size: 0.85rem;
  width: 70px;
  color: var(--color-text-secondary);
}

.bar-container {
  flex: 1;
  height: 12px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-pill);
  overflow: hidden;
}

.bar {
  height: 100%;
  border-radius: var(--radius-pill);
}

.bar-value {
  font-family: var(--font-mono);
  font-size: 0.9rem;
  width: 40px;
  text-align: right;
  font-weight: 600;
}

@media (min-width: 640px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
