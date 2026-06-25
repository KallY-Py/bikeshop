<template>
  <div class="bg-background text-on-background font-body-md min-h-screen flex flex-col md:flex-row">
    <!-- TopNavBar (Mobile Only) -->
    <header class="md:hidden flex justify-between items-center px-margin-mobile h-16 w-full bg-background/80 backdrop-blur-md border-b border-outline-variant shadow-none sticky top-0 z-50">
      <div class="font-display text-headline-lg-mobile text-primary uppercase tracking-tighter">VeloHub</div>
      <div class="flex gap-3 items-center">
        <span class="material-symbols-outlined text-primary cursor-pointer">notifications</span>
        <div class="relative">
          <button @click="isDropdownOpen = !isDropdownOpen" class="w-8 h-8 rounded-full overflow-hidden border border-outline-variant">
            <img v-if="userAvatar" :src="userAvatar" alt="Profile" class="w-full h-full object-cover" />
            <div v-else class="w-full h-full bg-primary/20 flex items-center justify-center text-primary font-bold">{{ userInitial }}</div>
          </button>
          <div v-if="isDropdownOpen" @click="isDropdownOpen = false" class="fixed inset-0 z-40"></div>
          <div v-if="isDropdownOpen" class="absolute right-0 top-10 w-48 bg-surface-container-high border border-outline-variant rounded-xl shadow-lg z-50 py-2">
            <div class="px-4 py-2 border-b border-outline-variant">
              <p class="font-label-sm text-on-surface truncate">{{ userName }}</p>
            </div>
            <router-link to="/settings" @click="isDropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">person</span>
              <span class="font-label-md text-label-md">Profile</span>
            </router-link>
            <router-link to="/settings" @click="isDropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
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
      <div class="px-6 mb-4">
        <h2 class="font-display text-[28px] text-primary uppercase tracking-tighter">VeloHub</h2>
      </div>

      <router-link to="/settings" class="px-6 mb-6 flex items-center gap-4 hover:bg-surface-container-high rounded-lg py-2 transition-colors group">
        <div class="w-12 h-12 rounded-full overflow-hidden bg-surface border-2 border-outline-variant group-hover:border-primary transition-colors">
          <img v-if="userAvatar" :src="userAvatar" alt="Profile" class="w-full h-full object-cover" />
          <div v-else class="w-full h-full bg-primary/20 flex items-center justify-center text-primary font-bold text-lg">{{ userInitial }}</div>
        </div>
        <div>
          <div class="font-label-md text-label-md text-on-surface-variant">Welcome back</div>
          <div class="font-headline-md text-headline-md text-primary">{{ userName }}</div>
        </div>
      </router-link>

      <div class="flex-1 overflow-y-auto mt-2 px-2">
        <router-link 
          v-for="item in navItems" 
          :key="item.name"
          :to="item.path"
          class="flex items-center gap-3 py-3 px-4 rounded-lg font-label-md text-label-md transition-all mb-1"
          :class="isActive(item.path) ? 'text-primary border-l-4 border-primary bg-primary/10' : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container-high'"
        >
          <span class="material-symbols-outlined">{{ item.icon }}</span>
          <span>{{ item.name }}</span>
        </router-link>
      </div>

      <div class="px-4 mt-4 border-t border-outline-variant pt-4 space-y-1">
        <router-link to="/settings" class="flex items-center gap-3 text-on-surface-variant hover:text-on-surface py-2 px-2 rounded-lg font-label-md text-label-md transition-colors hover:bg-surface-container-high">
          <span class="material-symbols-outlined">settings</span>
          <span>Settings</span>
        </router-link>
        <button @click="handleLogout" class="flex items-center gap-3 text-on-surface-variant hover:text-error py-2 px-2 rounded-lg font-label-md text-label-md transition-colors w-full text-left hover:bg-error/10">
          <span class="material-symbols-outlined">logout</span>
          <span>Logout</span>
        </button>
      </div>
    </nav>

    <!-- Main Content Area -->
    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <!-- Desktop Top Bar with Avatar -->
      <div class="hidden md:flex justify-end items-center mb-gutter">
        <div class="relative">
          <button @click="isDropdownOpen = !isDropdownOpen" class="flex items-center gap-3 hover:bg-surface-container-high rounded-lg p-2 transition-colors">
            <span class="font-label-sm text-label-sm text-on-surface-variant hidden lg:block">{{ userName }}</span>
            <div class="w-9 h-9 rounded-full overflow-hidden border-2 border-outline-variant hover:border-primary transition-colors">
              <img v-if="userAvatar" :src="userAvatar" alt="Profile" class="w-full h-full object-cover" />
              <div v-else class="w-full h-full bg-primary/20 flex items-center justify-center text-primary font-bold">{{ userInitial }}</div>
            </div>
            <span class="material-symbols-outlined text-on-surface-variant text-lg hidden sm:block">expand_more</span>
          </button>
          
          <div v-if="isDropdownOpen" @click="isDropdownOpen = false" class="fixed inset-0 z-40"></div>
          <div v-if="isDropdownOpen" class="absolute right-0 top-12 w-56 bg-surface-container-high border border-outline-variant rounded-xl shadow-lg z-50 py-2">
            <div class="px-4 py-3 border-b border-outline-variant">
              <p class="font-label-md text-on-surface">{{ userName }}</p>
              <p class="font-label-sm text-on-surface-variant truncate">{{ userEmail }}</p>
            </div>
            <router-link to="/settings" @click="isDropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
              <span class="material-symbols-outlined text-lg">person</span>
              <span class="font-label-md text-label-md">View Profile</span>
            </router-link>
            <router-link to="/settings" @click="isDropdownOpen = false" class="flex items-center gap-3 px-4 py-2.5 text-on-surface hover:bg-surface-container transition-colors">
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

      <!-- Page Content Slot -->
      <slot></slot>
    </main>

    <!-- BottomNavBar (Mobile) -->
    <nav class="fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-4 py-3 md:hidden bg-surface-container-lowest/95 backdrop-blur-lg border-t border-outline-variant shadow-[0_-4px_10px_rgba(0,0,0,0.3)] rounded-t-xl">
      <router-link 
        v-for="item in mobileNavItems" 
        :key="item.name"
        :to="item.path"
        class="flex flex-col items-center justify-center rounded-xl px-4 py-1 active:scale-90 transition-transform duration-150"
        :class="isActive(item.path) ? 'text-primary bg-primary/10' : 'text-on-surface-variant'"
      >
        <span class="material-symbols-outlined">{{ item.icon }}</span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </router-link>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'UserNav',
  data() {
    return {
      isDropdownOpen: false,
      navItems: [
        { name: 'Overview', icon: 'dashboard', path: '/dashboard' },
        { name: 'My Listings', icon: 'directions_bike', path: '/dashboard/listings' },
        { name: 'Sales', icon: 'payments', path: '/dashboard/sales' },
        { name: 'Analytics', icon: 'leaderboard', path: '/dashboard/analytics' },
        { name: 'Messages', icon: 'chat', path: '/dashboard/messages' }
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
      try { const u = JSON.parse(localStorage.getItem('user') || '{}'); return u.first_name ? `${u.first_name} ${u.last_name || ''}` : (u.username || 'Rider') } catch { return 'Rider' }
    },
    userEmail() {
      try { return JSON.parse(localStorage.getItem('user') || '{}').email || '' } catch { return '' }
    },
    userAvatar() {
      try { return JSON.parse(localStorage.getItem('user') || '{}').profile_image || '' } catch { return '' }
    },
    userInitial() {
      return (this.userName || 'U').charAt(0).toUpperCase()
    }
  },
  methods: {
    isActive(path) {
      return this.$route.path === path || (path !== '/dashboard' && this.$route.path.startsWith(path))
    },
    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('userId')
      localStorage.removeItem('userRole')
      localStorage.removeItem('user')
      this.isDropdownOpen = false
      this.$router.push('/register-login')
    }
  }
}
</script>