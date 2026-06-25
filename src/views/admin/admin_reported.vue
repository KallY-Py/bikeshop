<template>
  <div class="min-h-screen bg-gray-900 text-white flex overflow-hidden">
    <!-- Admin Navigation Component -->
    <Admin_nav 
      :current-view="currentView" 
      :is-collapsed="isSidebarCollapsed"
      @update-view="handleNavUpdate" 
      @logout="handleLogout" 
    />

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col h-screen overflow-hidden relative">
      
      <!-- Top Header -->
      <header class="h-16 bg-gray-800 border-b border-gray-700 flex items-center justify-between px-6 z-10 shadow-sm">
        <button @click="toggleSidebar" class="hidden md:flex items-center text-gray-400 hover:text-white transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
        <button @click="toggleMobileMenu" class="md:hidden text-gray-400 hover:text-white">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>

        <h2 class="text-xl font-bold text-white capitalize absolute left-1/2 transform -translate-x-1/2 hidden md:block">
          Reported Listings
        </h2>

        <div class="flex items-center space-x-4 ml-auto">
          <div class="relative">
            <span class="absolute top-0 right-0 h-2.5 w-2.5 bg-red-500 rounded-full border-2 border-gray-800"></span>
            <button class="text-gray-400 hover:text-orange-500 transition-colors p-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
            </button>
          </div>
          
          <div class="relative">
            <button @click="isProfileDropdownOpen = !isProfileDropdownOpen" class="flex items-center space-x-2 focus:outline-none">
              <img src="https://ui-avatars.com/api/?name=Admin+User&background=f97316&color=fff" alt="Admin" class="h-9 w-9 rounded-full border border-gray-600 hover:border-orange-500 transition-colors">
            </button>
            <div v-if="isProfileDropdownOpen" class="absolute right-0 mt-2 w-48 bg-gray-800 rounded-lg shadow-xl py-2 border border-gray-700 z-50 origin-top-right transform transition-all">
              <div class="px-4 py-2 border-b border-gray-700">
                <p class="text-sm font-medium text-white">Admin User</p>
                <p class="text-xs text-gray-400 truncate">admin@padyak.com</p>
              </div>
              <a href="#" class="block px-4 py-2 text-sm text-gray-300 hover:bg-gray-700 hover:text-white transition-colors">Your Profile</a>
              <a href="#" class="block px-4 py-2 text-sm text-gray-300 hover:bg-gray-700 hover:text-white transition-colors">Settings</a>
              <div class="border-t border-gray-700 my-1"></div>
              <button @click="handleLogout" class="w-full text-left px-4 py-2 text-sm text-red-400 hover:bg-gray-700 hover:text-red-300 transition-colors">Sign out</button>
            </div>
          </div>
        </div>
      </header>

      <!-- Scrollable Content Area -->
      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-900 p-6">
        <div class="space-y-6 animate-fade-in">
          
          <!-- Filters & Search -->
          <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
            <h2 class="text-2xl font-bold text-white">Report Management</h2>
            <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
              <div class="relative">
                <input v-model="searchQuery" type="text" placeholder="Search reports..." class="bg-gray-800 border border-gray-700 text-white px-4 py-2 pl-10 rounded-lg focus:outline-none focus:border-orange-500 w-full sm:w-64"/>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400 absolute left-3 top-2.5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
              </div>
              <select v-model="statusFilter" class="bg-gray-800 border border-gray-700 text-white px-4 py-2 rounded-lg focus:outline-none focus:border-orange-500">
                <option value="all">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="investigating">Investigating</option>
                <option value="resolved">Resolved</option>
              </select>
            </div>
          </div>

          <!-- Reports Table -->
          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
            <div class="overflow-x-auto">
              <table class="w-full text-left text-sm text-gray-400">
                <thead class="bg-gray-700/50 text-gray-200 uppercase">
                  <tr>
                    <th class="px-6 py-4">ID</th>
                    <th class="px-6 py-4">Reporter</th>
                    <th class="px-6 py-4">Listing / User</th>
                    <th class="px-6 py-4">Reason</th>
                    <th class="px-6 py-4">Status</th>
                    <th class="px-6 py-4">Date</th>
                    <th class="px-6 py-4 text-right">Actions</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-700">
                  <tr v-for="report in filteredReports" :key="report.id" class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-6 py-4">#{{ report.id }}</td>
                    <td class="px-6 py-4">
                      <div class="font-medium text-white">{{ report.reporter_name }}</div>
                      <div class="text-xs text-gray-500">{{ report.reporter_email }}</div>
                    </td>
                    <td class="px-6 py-4">
                      <div v-if="report.listing_id" class="font-medium text-white">Listing #{{ report.listing_id }}</div>
                      <div v-else class="font-medium text-white">User Report</div>
                      <div class="text-xs text-gray-500">{{ report.target_name || 'N/A' }}</div>
                    </td>
                    <td class="px-6 py-4 max-w-xs truncate" :title="report.reason">{{ report.reason }}</td>
                    <td class="px-6 py-4">
                      <select 
                        v-model="report.status" 
                        @change="updateReportStatus(report)"
                        :class="[
                          'px-2 py-1 rounded-lg text-xs font-semibold border cursor-pointer',
                          report.status === 'pending' ? 'bg-yellow-500/10 text-yellow-500 border-yellow-500/20' :
                          report.status === 'investigating' ? 'bg-blue-500/10 text-blue-500 border-blue-500/20' :
                          'bg-green-500/10 text-green-500 border-green-500/20'
                        ]"
                      >
                        <option value="pending">Pending</option>
                        <option value="investigating">Investigating</option>
                        <option value="resolved">Resolved</option>
                      </select>
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap">{{ formatDate(report.created_at) }}</td>
                    <td class="px-6 py-4 text-right">
                      <button v-if="report.listing_id && report.status !== 'resolved'" @click="suspendListing(report)" class="text-red-400 hover:text-red-300 mx-1 transition-colors" title="Suspend Listing">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" /></svg>
                      </button>
                      <button @click="viewDetails(report)" class="text-gray-400 hover:text-white mx-1 transition-colors" title="View Details">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" /></svg>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <div v-if="filteredReports.length === 0" class="p-12 text-center text-gray-500">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-4 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" /></svg>
              <p class="text-lg font-medium">No reports found</p>
              <p class="text-sm mt-1">Try adjusting your filters</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Admin_nav from '@/components/Admin_nav.vue'

export default {
  name: 'AdminReported',
  components: { Admin_nav },
  data() {
    return {
      currentView: 'reported',
      isSidebarCollapsed: false,
      isProfileDropdownOpen: false,
      searchQuery: '',
      statusFilter: 'all',
      reports: []
    }
  },
  computed: {
    filteredReports() {
      let filtered = [...this.reports]
      if (this.searchQuery) {
        const q = this.searchQuery.toLowerCase()
        filtered = filtered.filter(r => 
          r.reason.toLowerCase().includes(q) || 
          r.reporter_name.toLowerCase().includes(q) ||
          (r.target_name && r.target_name.toLowerCase().includes(q))
        )
      }
      if (this.statusFilter !== 'all') {
        filtered = filtered.filter(r => r.status === this.statusFilter)
      }
      return filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
    }
  },
  async mounted() {
    await this.loadReports()
    document.addEventListener('click', this.handleClickOutside)
  },
  beforeUnmount() {
    document.removeEventListener('click', this.handleClickOutside)
  },
  methods: {
    toggleSidebar() { this.isSidebarCollapsed = !this.isSidebarCollapsed },
    toggleMobileMenu() {},
    handleNavUpdate(view) { /* Handled by router */ },
    
    handleClickOutside(event) {
      const dropdown = this.$el.querySelector('.relative')
      if (dropdown && !dropdown.contains(event.target)) this.isProfileDropdownOpen = false
    },

    async loadReports() {
      try {
        const token = localStorage.getItem('token')
        const res = await axios.get('http://localhost:3000/api/admin/reports', {
          headers: { Authorization: `Bearer ${token}` }
        })
        this.reports = res.data || []
      } catch (err) {
        console.error('Failed to load reports:', err)
        this.reports = [] // Clear array instead of loading mock data
        
        // Alert the user so they know why the table is empty
        alert('Failed to load reports from the database. Please ensure the backend is running and you are logged in as an admin.')
      }
    },

    async updateReportStatus(report) {
      try {
        const token = localStorage.getItem('token')
        await axios.patch(`http://localhost:3000/api/admin/reports/${report.id}`, 
          { status: report.status },
          { headers: { Authorization: `Bearer ${token}` } }
        )
      } catch (err) {
        console.error('Failed to update report status:', err)
        alert('Failed to update status. Please try again.')
        await this.loadReports()
      }
    },

    async suspendListing(report) {
      if (!confirm(`Are you sure you want to suspend Listing #${report.listing_id}? This will hide it from the marketplace.`)) return
      
      try {
        const token = localStorage.getItem('token')
        await axios.patch(`http://localhost:3000/api/admin/listings/${report.listing_id}/status`, 
          { status: 'rejected' },
          { headers: { Authorization: `Bearer ${token}` } }
        )
        alert('Listing has been suspended successfully.')
        report.status = 'resolved'
        await this.updateReportStatus(report)
      } catch (err) {
        console.error('Failed to suspend listing:', err)
        alert('Failed to suspend listing. Please try again.')
      }
    },

    viewDetails(report) {
      alert(`Report Details:\n\nReason: ${report.reason}\nReporter: ${report.reporter_name}\nTarget: ${report.target_name || 'N/A'}\nStatus: ${report.status}`)
    },

    formatDate(dateString) {
      return new Date(dateString).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
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
::-webkit-scrollbar { width: 8px; height: 8px; }
::-webkit-scrollbar-track { background: #111827; }
::-webkit-scrollbar-thumb { background: #374151; border-radius: 4px; }
::-webkit-scrollbar-thumb:hover { background: #f97316; }
.animate-fade-in { animation: fadeIn 0.3s ease-in-out; }
@keyframes fadeIn { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: translateY(0); } }
</style>