<script setup lang="ts">
import { RouterView, RouterLink, useRoute } from 'vue-router';
import { PackageSearch } from 'lucide-vue-next';

const route = useRoute();
</script>

<template>
  <div class="app-container">
    <!-- Navbar -->
    <nav class="navbar">
      <div class="nav-content">
        <div class="logo">
          <PackageSearch class="logo-icon" :size="24" />
          <span>ShipTrack</span>
        </div>
        
        <div class="nav-links">
          <RouterLink to="/" class="nav-item" :class="{ active: route.name === 'tracking' }">
            <PackageSearch :size="18" />
            <span class="nav-text">Track</span>
          </RouterLink>
          
          <!-- History and Dashboard links are temporarily hidden
          <RouterLink to="/history" class="nav-item" :class="{ active: route.name === 'history' }">
            <Clock :size="18" />
            <span class="nav-text">History</span>
          </RouterLink>
          
          <RouterLink to="/dashboard" class="nav-item" :class="{ active: route.name === 'dashboard' }">
            <LayoutDashboard :size="18" />
            <span class="nav-text">Dashboard</span>
          </RouterLink>
          -->
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="main-content">
      <RouterView v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </RouterView>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.navbar {
  background-color: var(--color-surface);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid var(--color-border);
  border-radius: 16px;
  position: sticky;
  top: 16px;
  z-index: 50;
  margin: 16px 24px 0 24px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.2), inset 0 1px 0 rgba(255, 255, 255, 0.05);
}

.nav-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 12px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: var(--font-sans);
  font-weight: 600;
  font-size: 1.2rem;
  color: var(--color-text-primary);
}

.logo-icon {
  color: var(--color-primary);
}

.nav-links {
  display: flex;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 12px;
  border-radius: var(--radius-control);
  color: var(--color-text-secondary);
  font-size: 0.95rem;
  transition: all 0.2s ease;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-text-primary);
}

.nav-item.active {
  background: rgba(99, 102, 241, 0.15);
  color: var(--color-primary);
  font-weight: 500;
  box-shadow: inset 0 1px 0 rgba(255,255,255,0.05);
}

.nav-text {
  display: none;
}

.main-content {
  flex: 1;
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  padding: 40px 24px;
}

/* Page Transitions */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

@media (min-width: 640px) {
  .nav-text {
    display: inline;
  }
}
</style>
