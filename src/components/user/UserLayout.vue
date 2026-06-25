<template>
  <div class="bg-background text-on-background font-body-md min-h-screen flex flex-col md:flex-row">
    <!-- TopNavBar (Mobile Only) -->
    <header class="md:hidden flex justify-between items-center px-margin-mobile h-16 w-full bg-background/80 backdrop-blur-md border-b border-outline-variant shadow-none sticky top-0 z-50">
      <div class="font-display text-headline-lg-mobile text-primary uppercase tracking-tighter">VeloHub</div>
      <div class="flex gap-4 items-center">
        <span class="material-symbols-outlined text-primary cursor-pointer">notifications</span>
        <!-- Avatar Dropdown (Mobile) -->
        <div class="relative">
          <button @click="toggleDropdown" class="w-8 h-8 rounded-full overflow-hidden border border-outline-variant">
            <img :src="userAvatar" alt="Profile" class="w-full h-full object-cover" />
          </button>
          <div v-if="dropdownOpen" @click="dropdownOpen = false" class="fixed inset-0 z-40"></div>
          <div v-if="dropdownOpen" class="absolute right-0 top-10 w-48 bg-surface-container-high border border-outline-variant rounded-xl shadow-lg z-50 py-2">
            <router-link to="/settings" @click="dropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">person</span>
              <span class="font-label-md text-label-md">Profile</span>
            </router-link>
            <router-link to="/settings" @click="dropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">settings</span>
              <span class="font-label-md text-label-md">Settings</span>
            </router-link>
            <hr class="border-outline-variant my-1" />
            <button @click="handleLogout" class="flex items-center gap-3 px-4 py-2.5 text-error hover:bg-error/10 transition-colors w-full text-left">
              <span class="material-symbols-outlined text-lg">logout</span>
              <span class="font-label-md text-label-md">Logout</span>
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- SideNavBar (Desktop) -->
    <nav class="hidden md:flex flex-col h-screen py-8 fixed left-0 top-0 w-[280px] bg-surface-container-low border-r border-outline-variant shadow-sm z-40">
      <!-- Brand -->
      <div class="px-6 mb-4">
        <h2 class="font-display text-[28px] text-primary uppercase tracking-tighter">VeloHub</h2>
      </div>

      <!-- User Profile Section -->
      <div class="px-6 mb-6 flex items-center gap-4">
        <div class="w-12 h-12 rounded-full overflow-hidden bg-surface border border-outline-variant">
          <img :src="userAvatar" alt="User profile" class="w-full h-full object-cover" />
        </div>
        <div>
          <div class="font-label-md text-label-md text-on-surface-variant">Welcome back</div>
          <div class="font-headline-md text-headline-md text-primary">{{ userName }}</div>
        </div>
      </div>

      <!-- Navigation Links -->
      <div class="flex-1 overflow-y-auto mt-2">
        <router-link 
          v-for="item in navItems" 
          :key="item.name"
          :to="item.path"
          class="flex items-center gap-3 py-3 px-6 font-label-md text-label-md transition-all"
          :class="$route.path === item.path 
            ? 'text-primary border-l-4 border-primary bg-primary/10' 
            : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container-high active:translate-x-1'"
        >
          <span class="material-symbols-outlined">{{ item.icon }}</span>
          <span>{{ item.name }}</span>
          <span v-if="item.badge" class="ml-auto bg-primary text-background text-xs px-2 py-0.5 rounded-full">{{ item.badge }}</span>
        </router-link>
      </div>

      <!-- Post New Bike Button -->
      <div class="px-6 mt-4">
        <button 
          @click="$router.push('/dashboard/listings')"
          class="w-full bg-primary-container text-on-primary-container font-label-md text-label-md py-3 rounded-lg flex items-center justify-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] active:scale-95 transition-transform shadow-lg"
        >
          <span class="material-symbols-outlined">add</span>
          Post New Bike
        </button>
      </div>

      <!-- Bottom Links -->
      <div class="px-6 mt-4 border-t border-outline-variant pt-4 space-y-1">
        <router-link to="/settings" class="flex items-center gap-3 text-on-surface-variant hover:text-on-surface py-2 font-label-md text-label-md transition-colors">
          <span class="material-symbols-outlined">settings</span>
          <span>Settings</span>
        </router-link>
        <button @click="handleLogout" class="flex items-center gap-3 text-on-surface-variant hover:text-error py-2 font-label-md text-label-md transition-colors w-full text-left">
          <span class="material-symbols-outlined">logout</span>
          <span>Logout</span>
        </button>
      </div>
    </nav>

    <!-- Main Content Slot -->
    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <!-- Desktop Top Bar with Avatar Dropdown -->
      <div class="hidden md:flex justify-end items-center mb-gutter">
        <div class="relative">
          <button @click="toggleDropdown" class="flex items-center gap-3 hover:bg-surface-container-high rounded-lg p-2 transition-colors">
            <span class="font-label-sm text-label-sm text-on-surface-variant">{{ userName }}</span>
            <div class="w-9 h-9 rounded-full overflow-hidden border-2 border-primary/30 hover:border-primary transition-colors">
              <img :src="userAvatar" alt="Profile" class="w-full h-full object-cover" />
            </div>
            <span class="material-symbols-outlined text-on-surface-variant text-lg">expand_more</span>
          </button>
          
          <!-- Dropdown Menu -->
          <div v-if="dropdownOpen" @click="dropdownOpen = false" class="fixed inset-0 z-40"></div>
          <div v-if="dropdownOpen" class="absolute right-0 top-12 w-56 bg-surface-container-high border border-outline-variant rounded-xl shadow-lg z-50 py-2">
            <div class="px-4 py-3 border-b border-outline-variant">
              <p class="font-label-md text-label-md text-on-surface">{{ userName }}</p>
              <p class="font-label-sm text-label-sm text-on-surface-variant">{{ userEmail }}</p>
            </div>
            <router-link to="/settings" @click="dropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">person</span>
              <span class="font-label-md text-label-md">View Profile</span>
            </router-link>
            <router-link to="/settings" @click="dropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">settings</span>
              <span class="font-label-md text-label-md">Settings</span>
            </router-link>
            <hr class="border-outline-variant my-1" />
            <button @click="handleLogout" class="flex items-center gap-3 px-4 py-2.5 text-error hover:bg-error/10 transition-colors w-full text-left">
              <span class="material-symbols-outlined text-lg">logout</span>
              <span class="font-label-md text-label-md">Logout</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Page Content -->
      <slot></slot>
    </main>

    <!-- BottomNavBar (Mobile Only) -->
    <nav class="fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-4 py-3 md:hidden bg-surface-container-lowest/95 backdrop-blur-lg border-t border-outline-variant shadow-[0_-4px_10px_rgba(0,0,0,0.3)] rounded-t-xl">
      <router-link 
        v-for="item in mobileNavItems" 
        :key="item.name"
        :to="item.path"
        class="flex flex-col items-center justify-center rounded-xl px-4 py-1 active:scale-90 transition-transform duration-150"
        :class="$route.path === item.path ? 'text-primary bg-primary/10' : 'text-on-surface-variant'"
      >
        <span class="material-symbols-outlined">{{ item.icon }}</span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </router-link>
    </nav>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'

export default {
  name: 'UserLayout',
  
  data() {
    return {
      dropdownOpen: false,
      
      navItems: [
        { name: 'Overview', icon: 'dashboard', path: '/dashboard' },
        { name: 'My Listings', icon: 'directions_bike', path: '/dashboard/listings' },
        { name: 'Sales', icon: 'payments', path: '/dashboard/sales' },
        { name: 'Analytics', icon: 'leaderboard', path: '/dashboard/analytics' },
        { name: 'Messages', icon: 'chat', path: '/dashboard/messages', badge: '2' }
      ],
      
      mobileNavItems: [
        { name: 'Home', icon: 'home', path: '/dashboard' },
        { name: 'Listings', icon: 'storefront', path: '/dashboard/listings' },
        { name: 'Chat', icon: 'forum', path: '/dashboard/messages' },
        { name: 'Profile', icon: 'person', path: '/settings' }
      ]
    }
  },
  
  computed: {
    userName() {
      const user = localStorage.getItem('user')
      if (user) {
        try {
          const u = JSON.parse(user)
          return (u.first_name || '') + ' ' + (u.last_name || '') || u.username || 'Rider'
        } catch { return 'Rider' }
      }
      return 'Rider'
    },
    userEmail() {
      const user = localStorage.getItem('user')
      if (user) {
        try {
          return JSON.parse(user).email || ''
        } catch { return '' }
      }
      return ''
    },
    userAvatar() {
      const user = localStorage.getItem('user')
      if (user) {
        try {
          return JSON.parse(user).profile_image || 'https://via.placeholder.com/48'
        } catch { return 'https://via.placeholder.com/48' }
      }
      return 'https://via.placeholder.com/48'
    }
  },
  
  methods: {
    toggleDropdown() {
      this.dropdownOpen = !this.dropdownOpen
    },
    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('userId')
      localStorage.removeItem('userRole')
      localStorage.removeItem('user')
      this.$router.push('/register-login')
    }
  }
}
</script>