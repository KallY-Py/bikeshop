<template>
  <div class="min-h-screen bg-gray-900 text-white flex overflow-hidden">
    
    <!-- Admin Navigation Component -->
    <Admin_nav 
      :current-view="currentView" 
      :is-collapsed="isSidebarCollapsed"
      @update-view="currentView = $event" 
      @logout="handleLogout" 
    />

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col h-screen overflow-hidden relative">
      
      <!-- Top Header -->
      <header class="h-16 bg-gray-800 border-b border-gray-700 flex items-center justify-between px-6 z-10 shadow-sm">
        <!-- Toggle Sidebar Button (Desktop) -->
        <button @click="toggleSidebar" class="hidden md:flex items-center text-gray-400 hover:text-white transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>

        <!-- Mobile Menu Button -->
        <button @click="toggleMobileMenu" class="md:hidden text-gray-400 hover:text-white">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>

        <!-- Centered Title -->
        <h2 class="text-xl font-bold text-white capitalize absolute left-1/2 transform -translate-x-1/2 hidden md:block">
          Dashboard Overview
        </h2>

        <!-- Right Side Icons -->
        <div class="flex items-center space-x-4 ml-auto">
          <!-- Notification Bell -->
          <div class="relative">
            <span v-if="pendingReportsCount > 0" class="absolute top-0 right-0 h-2.5 w-2.5 bg-red-500 rounded-full border-2 border-gray-800"></span>
            <button class="text-gray-400 hover:text-orange-500 transition-colors p-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
            </button>
          </div>
          
          <!-- Profile Dropdown -->
          <div class="relative">
            <button @click="isProfileDropdownOpen = !isProfileDropdownOpen" class="flex items-center space-x-2 focus:outline-none">
              <img 
                src="https://ui-avatars.com/api/?name=Admin+User&background=f97316&color=fff" 
                alt="Admin" 
                class="h-9 w-9 rounded-full border border-gray-600 hover:border-orange-500 transition-colors"
              >
            </button>

            <!-- Dropdown Menu -->
            <div 
              v-if="isProfileDropdownOpen" 
              class="absolute right-0 mt-2 w-48 bg-gray-800 rounded-lg shadow-xl py-2 border border-gray-700 z-50 origin-top-right transform transition-all"
            >
              <div class="px-4 py-2 border-b border-gray-700">
                <p class="text-sm font-medium text-white">Admin User</p>
                <p class="text-xs text-gray-400 truncate">admin@padyak.com</p>
              </div>
              
              <a href="#" class="block px-4 py-2 text-sm text-gray-300 hover:bg-gray-700 hover:text-white transition-colors">
                Your Profile
              </a>
              <div class="border-t border-gray-700 my-1"></div>
              <button @click="handleLogout" class="w-full text-left px-4 py-2 text-sm text-red-400 hover:bg-gray-700 hover:text-red-300 transition-colors">
                Sign out
              </button>
            </div>
          </div>
        </div>
      </header>

      <!-- Scrollable Content Area -->
      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-900 p-6">
        
        <!-- VIEW: DASHBOARD -->
        <div v-if="currentView === 'dashboard'" class="space-y-6 animate-fade-in">
          
          <!-- Stats Cards Row -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <!-- Total Users -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">TOTAL USERS</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-orange-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">{{ stats.totalUsers }}</p>
                <!-- Static percentage for now as we don't have historical data in this simple query -->
                <span class="text-gray-500 text-sm font-medium">Registered</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-orange-500 w-3/4"></div>
              </div>
            </div>

            <!-- Active Listings -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">ACTIVE LISTINGS</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">{{ stats.activeListings }}</p>
                <span class="text-gray-500 text-sm font-medium">Approved</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-blue-500 w-1/2"></div>
              </div>
            </div>

            <!-- Pending Reports -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">PENDING REPORTS</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">{{ stats.pendingReports }}</p>
                <span class="text-red-400 text-sm font-medium">Action Needed</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-red-500" :style="{ width: Math.min((stats.pendingReports / (stats.totalUsers || 1)) * 100, 100) + '%' }"></div>
              </div>
            </div>
          </div>

          <!-- Chart and Moderation Queue -->
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Revenue Growth Chart (Placeholder for visualization) -->
            <div class="lg:col-span-2 bg-gray-800 rounded-xl border border-gray-700 shadow-lg p-6">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-lg font-semibold text-white">Platform Activity</h3>
                <div class="flex space-x-2">
                  <button class="px-3 py-1 text-xs rounded bg-orange-500 text-white">Recent</button>
                </div>
              </div>
              <div class="h-64 flex items-center justify-center bg-gray-900/50 rounded-lg relative overflow-hidden">
                 <!-- Simple visual representation of data presence -->
                 <div v-if="stats.totalUsers > 0" class="text-center">
                    <p class="text-gray-400 mb-2">Database Connected</p>
                    <p class="text-sm text-gray-500">Total Records: {{ stats.totalUsers + stats.activeListings }}</p>
                 </div>
                 <p v-else class="text-gray-500">Loading Data...</p>
              </div>
            </div>

            <!-- Moderation Queue (Real Data) -->
            <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg p-6">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-lg font-semibold text-white flex items-center">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 text-orange-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3" />
                  </svg>
                  Moderation Queue
                </h3>
                <span class="px-2 py-1 text-xs rounded bg-orange-500/20 text-orange-400 border border-orange-500/30">
                  {{ pendingReports.length }} Pending
                </span>
              </div>
              
              <div class="space-y-4 max-h-[300px] overflow-y-auto pr-2 custom-scrollbar">
                <div v-if="pendingReports.length === 0" class="text-center text-gray-500 py-4">
                  <p>No pending reports.</p>
                </div>

                <div v-for="report in pendingReports" :key="report.id" class="p-3 bg-gray-900/50 rounded-lg border border-gray-700">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="text-sm font-medium text-white truncate max-w-[150px]" :title="report.target_name">
                      {{ report.target_name || 'Unknown Target' }}
                    </h4>
                    <span class="text-xs text-gray-500">{{ formatDate(report.created_at) }}</span>
                  </div>
                  <p class="text-xs text-gray-400 mb-3 line-clamp-2" :title="report.reason">
                    "{{ report.reason }}"
                  </p>
                  <div class="flex space-x-2">
                    <button @click="$router.push('/admin/reported')" class="px-3 py-1 text-xs rounded bg-gray-700 text-gray-300 hover:bg-gray-600 transition-colors">Review</button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent User Activity Table (Real Data) -->
          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
            <div class="p-6 border-b border-gray-700 flex items-center justify-between">
              <h3 class="text-lg font-semibold text-white">Recent User Registrations</h3>
              <button @click="$router.push('/admin/users')" class="text-orange-500 hover:text-orange-400 text-sm font-medium">View All Users</button>
            </div>
            
            <div class="overflow-x-auto">
              <table class="w-full text-left text-sm">
                <thead class="bg-gray-900/50 text-gray-400">
                  <tr>
                    <th class="px-6 py-4 font-medium">USER</th>
                    <th class="px-6 py-4 font-medium">STATUS</th>
                    <th class="px-6 py-4 font-medium">ROLE</th>
                    <th class="px-6 py-4 font-medium">JOINED</th>
                    <th class="px-6 py-4 font-medium text-right">ACTIONS</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-700">
                  <tr v-if="recentUsers.length === 0">
                    <td colspan="5" class="px-6 py-8 text-center text-gray-500">No users found in database.</td>
                  </tr>
                  <tr v-for="user in recentUsers" :key="user.id" class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-6 py-4">
                      <div class="flex items-center">
                        <div class="h-10 w-10 rounded-full bg-gradient-to-br from-orange-500 to-red-500 flex items-center justify-center text-white font-bold text-sm mr-3">
                          {{ user.first_name.charAt(0) }}{{ user.last_name.charAt(0) }}
                        </div>
                        <div>
                          <p class="font-medium text-white">{{ user.first_name }} {{ user.last_name }}</p>
                          <p class="text-xs text-gray-400">{{ user.email }}</p>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <span :class="[
                        'px-2 py-1 rounded-full text-xs font-medium border',
                        user.status === 'active' ? 'bg-green-500/10 text-green-400 border-green-500/20' : 
                        user.status === 'suspended' ? 'bg-yellow-500/10 text-yellow-400 border-yellow-500/20' : 
                        'bg-red-500/10 text-red-400 border-red-500/20'
                      ]">
                        {{ user.status.toUpperCase() }}
                      </span>
                    </td>
                    <td class="px-6 py-4 text-gray-300 capitalize">{{ user.role }}</td>
                    <td class="px-6 py-4 text-gray-400">{{ formatDate(user.created_at) }}</td>
                    <td class="px-6 py-4 text-right">
                      <button @click="$router.push('/admin/users')" class="text-gray-400 hover:text-white transition-colors">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Other views placeholders (kept minimal as requested) -->
        <div v-if="currentView !== 'dashboard'" class="flex flex-col items-center justify-center h-96 text-gray-500">
           <p>Select Dashboard view to see stats.</p>
        </div>

      </main>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Admin_nav from '@/components/Admin_nav.vue'

export default {
  name: 'AdminDashboard',
  components: {
    Admin_nav
  },
  
  data() {
    return {
      currentView: 'dashboard',
      isSidebarCollapsed: false,
      isProfileDropdownOpen: false,
      
      // Real Data State
      stats: {
        totalUsers: 0,
        activeListings: 0,
        pendingReports: 0
      },
      recentUsers: [],
      pendingReports: []
    }
  },

  async mounted() {
    await this.fetchDashboardData()
    document.addEventListener('click', this.handleClickOutside)
  },

  beforeUnmount() {
    document.removeEventListener('click', this.handleClickOutside)
  },

  methods: {
    toggleSidebar() {
      this.isSidebarCollapsed = !this.isSidebarCollapsed
    },
    
    toggleMobileMenu() {
      this.$emit('toggle-mobile-menu')
    },

    handleClickOutside(event) {
      const dropdown = this.$el.querySelector('.relative')
      if (dropdown && !dropdown.contains(event.target)) {
        this.isProfileDropdownOpen = false
      }
    },

    async fetchDashboardData() {
      try {
        const token = localStorage.getItem('token')
        const headers = { Authorization: `Bearer ${token}` }

        // Fetch all required data in parallel
        const [usersRes, listingsRes, reportsRes] = await Promise.all([
          axios.get('http://localhost:3000/api/admin/users', { headers }).catch(() => ({ data: [] })), // Fallback to public if admin fails
          axios.get('http://localhost:3000/api/listings', { headers }).catch(() => ({ data: [] })),
          axios.get('http://localhost:3000/api/admin/reports', { headers }).catch(() => ({ data: [] }))
        ])

        const users = usersRes.data || []
        const listings = listingsRes.data || []
        const reports = reportsRes.data || []

        // Calculate Stats
        this.stats.totalUsers = users.length
        
        // Count only 'approved' listings as active
        this.stats.activeListings = listings.filter(l => l.status === 'approved').length
        
        // Count only 'pending' reports
        const pending = reports.filter(r => r.status === 'pending')
        this.stats.pendingReports = pending.length
        this.pendingReports = pending.slice(0, 5) // Show top 5 in queue

        // Get Recent Users (Last 5 registered)
        // Sort by created_at descending and take first 5
        this.recentUsers = users
          .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
          .slice(0, 5)

      } catch (err) {
        console.error('Failed to load dashboard data:', err)
        alert('Could not connect to database. Please ensure backend is running.')
      }
    },

    formatDate(dateString) {
      if (!dateString) return 'N/A'
      const date = new Date(dateString)
      return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
    },

    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      this.$router.push('/register-login')
    }
  }
}
</script>

<style scoped>
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: #111827;
}

::-webkit-scrollbar-thumb {
  background: #374151;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #f97316;
}

.animate-fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #4b5563;
  border-radius: 2px;
}
</style>