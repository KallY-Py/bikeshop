<template>
  <div class="flex flex-col h-screen bg-gray-800 border-r border-gray-700 transition-all duration-300" :class="isCollapsed ? 'w-20' : 'w-64'">
    <!-- Mobile Header -->
    <div class="h-16 flex items-center justify-between px-4 border-b border-gray-700 md:hidden bg-gray-900">
      <button @click="toggleSidebar" class="text-gray-400 hover:text-white focus:outline-none">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
      <span class="text-orange-500 font-bold tracking-wider">PADYAK ALL</span>
      <div class="w-6"></div> 
    </div>

    <!-- Desktop Sidebar -->
    <aside 
      :class="[
        'hidden md:flex flex-col h-full bg-gray-800 transition-all duration-300',
        isMobileOpen ? 'flex fixed inset-y-0 left-0 z-50 w-64' : 'hidden',
        'md:relative md:flex'
      ]"
    >
      <!-- Logo Section -->
      <div class="h-16 flex items-center justify-center border-b border-gray-700 bg-gray-800">
        <h1 v-if="!isCollapsed" class="text-xl font-bold text-orange-500 tracking-wider">
          PADYAK<span class="text-white"> ADMIN</span>
        </h1>
        <h1 v-else class="text-2xl font-bold text-orange-500">PA</h1>
      </div>

      <!-- Profile Section (Only show when expanded) -->
      <div v-if="!isCollapsed" class="p-6 border-b border-gray-700 flex items-center space-x-3 bg-gray-800/50">
        <img 
          :src="adminProfileImage" 
          alt="Admin" 
          class="h-12 w-12 rounded-full border-2 border-orange-500"
        />
        <div>
          <p class="text-sm text-gray-400">Welcome back,</p>
          <p class="font-semibold text-white">{{ adminName }}</p>
        </div>
      </div>

      <!-- Navigation Links -->
      <nav class="flex-1 px-4 py-6 space-y-2 overflow-y-auto custom-scrollbar">
        <a 
          href="#" 
          @click.prevent="handleNav('dashboard')"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors group', currentView === 'dashboard' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">Dashboard</span>
        </a>
        
        <a 
          href="#" 
          @click.prevent="handleNav('management')"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors group', currentView === 'management' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">User Management</span>
        </a>

        <a 
          href="#" 
          @click.prevent="handleNav('reported')"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors group', currentView === 'reported' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">Reported Listings</span>
        </a>

        <a 
          href="#" 
          @click.prevent="handleNav('categories')"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors group', currentView === 'categories' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">Category Management</span>
        </a>

        <a 
          href="#" 
          @click.prevent="handleNav('settings')"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors group', currentView === 'settings' ? 'bg-orange-500 text-white shadow-lg shadow-orange-500/20' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">Platform Settings</span>
        </a>
      </nav>

      <!-- Logout Button -->
      <div class="p-4 border-t border-gray-700 bg-gray-800">
        <button @click="emitLogout" class="flex items-center w-full px-4 py-2 text-gray-400 hover:text-red-400 hover:bg-gray-700 rounded-lg transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          <span :class="{ 'hidden': isCollapsed }">Logout</span>
        </button>
      </div>
    </aside>

    <!-- Mobile Overlay -->
    <div 
      v-if="isMobileOpen" 
      class="fixed inset-0 bg-black/50 z-40 md:hidden" 
      @click="toggleSidebar"
    ></div>
  </div>
</template>

<script>
export default {
  name: 'AdminNav',
  props: {
    currentView: {
      type: String,
      default: 'dashboard'
    },
    isCollapsed: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      isMobileOpen: false,
      adminName: 'Admin',
      adminProfileImage: 'https://ui-avatars.com/api/?name=Admin+User&background=f97316&color=fff'
    }
  },
  mounted() {
    const userStr = localStorage.getItem('user')
    if (userStr) {
      try {
        const user = JSON.parse(userStr)
        this.adminName = user.first_name || 'Admin'
      } catch (e) {
        console.error('Error parsing user data', e)
      }
    }
  },
  methods: {
    toggleSidebar() {
      this.isMobileOpen = !this.isMobileOpen
    },
    handleNav(view) {
    const routes = {
            dashboard: '/admin/dashboard',
            management: '/admin/management',
            reported: '/admin/reported',
            categories: '/admin/categories',
            settings: '/admin/settings'
        }
        
        if (routes[view]) {
            this.$router.push(routes[view])
        } else {
            this.$emit('update-view', view)
        }
        
        if (window.innerWidth < 768) {
            this.isMobileOpen = false
        }
    },
    emitLogout() {
      this.$emit('logout')
    }
  }
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #374151;
  border-radius: 20px;
}
</style>