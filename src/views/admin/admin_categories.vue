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
        <button @click="toggleSidebar" class="hidden md:flex items-center text-gray-400 hover:text-white transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>

        <h2 class="text-xl font-bold text-white capitalize absolute left-1/2 transform -translate-x-1/2 hidden md:block">
          Category & Listing Management
        </h2>

        <div class="flex items-center space-x-4 ml-auto">
          <div class="relative">
            <button @click="isProfileDropdownOpen = !isProfileDropdownOpen" class="flex items-center space-x-2 focus:outline-none">
              <img src="https://ui-avatars.com/api/?name=Admin+User&background=f97316&color=fff" alt="Admin" class="h-9 w-9 rounded-full border border-gray-600 hover:border-orange-500 transition-colors">
            </button>
            <div v-if="isProfileDropdownOpen" class="absolute right-0 mt-2 w-48 bg-gray-800 rounded-lg shadow-xl py-2 border border-gray-700 z-50">
              <button @click="handleLogout" class="w-full text-left px-4 py-2 text-sm text-red-400 hover:bg-gray-700">Sign out</button>
            </div>
          </div>
        </div>
      </header>

      <!-- Scrollable Content Area -->
      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-900 p-6">
        
        <div class="grid grid-cols-1 lg:grid-cols-4 gap-6 h-full">
          
          <!-- Left Sidebar: Categories -->
          <div class="lg:col-span-1 bg-gray-800 rounded-xl border border-gray-700 p-4 h-fit">
            <h3 class="text-lg font-bold text-white mb-4 flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 text-orange-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
              </svg>
              Categories
            </h3>
            
            <div class="space-y-2">
              <button 
                @click="selectedCategory = null"
                :class="['w-full text-left px-4 py-3 rounded-lg transition-colors flex items-center justify-between', selectedCategory === null ? 'bg-orange-500 text-white' : 'bg-gray-700/50 text-gray-300 hover:bg-gray-700']"
              >
                <span>All Listings</span>
                <span class="bg-gray-900/50 px-2 py-0.5 rounded text-xs">{{ listings.length }}</span>
              </button>

              <button 
                v-for="cat in categories" 
                :key="cat.id"
                @click="selectedCategory = cat.id"
                :class="['w-full text-left px-4 py-3 rounded-lg transition-colors flex items-center justify-between', selectedCategory === cat.id ? 'bg-orange-500 text-white' : 'bg-gray-700/50 text-gray-300 hover:bg-gray-700']"
              >
                <span>{{ cat.name }}</span>
                <span class="bg-gray-900/50 px-2 py-0.5 rounded text-xs">{{ getListingsByCategory(cat.id).length }}</span>
              </button>
            </div>

            <div class="mt-8 pt-6 border-t border-gray-700">
              <h4 class="text-sm font-semibold text-gray-400 mb-3 uppercase tracking-wider">Filter by Status</h4>
              <select v-model="statusFilter" class="w-full bg-gray-900 border border-gray-700 text-white px-3 py-2 rounded-lg focus:outline-none focus:border-orange-500">
                <option value="all">All Statuses</option>
                <option value="pending">Pending Approval</option>
                <option value="approved">Approved</option>
                <option value="rejected">Rejected</option>
              </select>
            </div>
          </div>

          <!-- Right Content: Listings Grid -->
          <div class="lg:col-span-3 space-y-6">
            
            <!-- Stats Summary -->
            <div class="grid grid-cols-3 gap-4">
              <div class="bg-gray-800 p-4 rounded-lg border border-gray-700">
                <p class="text-gray-400 text-xs uppercase">Pending</p>
                <p class="text-2xl font-bold text-yellow-500">{{ stats.pending }}</p>
              </div>
              <div class="bg-gray-800 p-4 rounded-lg border border-gray-700">
                <p class="text-gray-400 text-xs uppercase">Approved</p>
                <p class="text-2xl font-bold text-green-500">{{ stats.approved }}</p>
              </div>
              <div class="bg-gray-800 p-4 rounded-lg border border-gray-700">
                <p class="text-gray-400 text-xs uppercase">Total</p>
                <p class="text-2xl font-bold text-white">{{ stats.total }}</p>
              </div>
            </div>

            <!-- Listings Table -->
            <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
              <div class="p-4 border-b border-gray-700 flex justify-between items-center">
                <h3 class="font-bold text-white">Listing Management</h3>
                <button @click="loadData" class="text-sm text-orange-500 hover:text-orange-400">Refresh Data</button>
              </div>
              
              <div class="overflow-x-auto">
                <table class="w-full text-left text-sm text-gray-400">
                  <thead class="bg-gray-900/50 text-gray-200 uppercase">
                    <tr>
                      <th class="px-6 py-4">Image</th>
                      <th class="px-6 py-4">Title / Category</th>
                      <th class="px-6 py-4">Seller</th>
                      <th class="px-6 py-4">Price</th>
                      <th class="px-6 py-4">Status</th>
                      <th class="px-6 py-4 text-right">Actions</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-700">
                    <tr v-for="listing in filteredListings" :key="listing.id" class="hover:bg-gray-700/30 transition-colors">
                      <td class="px-6 py-4">
                        <img :src="(listing.images && listing.images.length > 0) ? listing.images[0] : 'https://via.placeholder.com/50'" class="w-12 h-12 rounded object-cover border border-gray-600">
                      </td>
                      <td class="px-6 py-4">
                        <div class="font-medium text-white">{{ listing.title }}</div>
                        <div class="text-xs text-gray-500">{{ listing.category_name }}</div>
                      </td>
                      <td class="px-6 py-4">
                        <div class="text-white">{{ listing.seller_name }}</div>
                        <div class="text-xs text-gray-500">{{ listing.seller_email }}</div>
                      </td>
                      <td class="px-6 py-4 font-mono text-white">
                        ₱{{ Number(listing.price).toLocaleString() }}
                      </td>
                      <td class="px-6 py-4">
                        <span :class="[
                          'px-2 py-1 rounded-full text-xs font-semibold border',
                          listing.status === 'approved' ? 'bg-green-500/10 text-green-500 border-green-500/20' :
                          listing.status === 'rejected' ? 'bg-red-500/10 text-red-500 border-red-500/20' :
                          'bg-yellow-500/10 text-yellow-500 border-yellow-500/20'
                        ]">
                          {{ listing.status.toUpperCase() }}
                        </span>
                      </td>
                      <td class="px-6 py-4 text-right space-x-2">
                        <button 
                          v-if="listing.status !== 'approved'"
                          @click="updateStatus(listing.id, 'approved')"
                          class="text-green-500 hover:text-green-400 p-1 rounded hover:bg-green-500/10 transition-colors"
                          title="Approve Listing"
                        >
                          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                          </svg>
                        </button>
                        <button 
                          v-if="listing.status !== 'rejected'"
                          @click="updateStatus(listing.id, 'rejected')"
                          class="text-red-500 hover:text-red-400 p-1 rounded hover:bg-red-500/10 transition-colors"
                          title="Reject Listing"
                        >
                          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                          </svg>
                        </button>
                      </td>
                    </tr>
                    <tr v-if="filteredListings.length === 0">
                      <td colspan="6" class="px-6 py-12 text-center text-gray-500">
                        No listings found matching your criteria.
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
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
  name: 'AdminCategory',
  components: { Admin_nav },
  data() {
    return {
      currentView: 'categories',
      isSidebarCollapsed: false,
      isProfileDropdownOpen: false,
      
      categories: [],
      listings: [],
      
      selectedCategory: null,
      statusFilter: 'all'
    }
  },
  computed: {
    stats() {
      return {
        total: this.listings.length,
        pending: this.listings.filter(l => l.status === 'pending').length,
        approved: this.listings.filter(l => l.status === 'approved').length
      }
    },
    filteredListings() {
      let filtered = [...this.listings]
      
      // Filter by Category
      if (this.selectedCategory) {
        filtered = filtered.filter(l => l.category_id === this.selectedCategory)
      }
      
      // Filter by Status
      if (this.statusFilter !== 'all') {
        filtered = filtered.filter(l => l.status === this.statusFilter)
      }
      
      return filtered
    }
  },
  async mounted() {
    await this.loadData()
    document.addEventListener('click', this.handleClickOutside)
  },
  beforeUnmount() {
    document.removeEventListener('click', this.handleClickOutside)
  },
  methods: {
    toggleSidebar() { this.isSidebarCollapsed = !this.isSidebarCollapsed },
    
    handleClickOutside(event) {
      const dropdown = this.$el.querySelector('.relative')
      if (dropdown && !dropdown.contains(event.target)) this.isProfileDropdownOpen = false
    },

    async loadData() {
      try {
        const token = localStorage.getItem('token')
        const headers = { Authorization: `Bearer ${token}` }
        
        // Fetch Categories
        const catRes = await axios.get('http://localhost:3000/api/categories')
        this.categories = catRes.data || []
        
        // Fetch All Listings (Admin endpoint)
        const listRes = await axios.get('http://localhost:3000/api/admin/listings', { headers })
        this.listings = listRes.data || []
        
      } catch (err) {
        console.error('Failed to load data:', err)
        alert('Failed to connect to server. Please ensure backend is running.')
      }
    },

    getListingsByCategory(catId) {
      return this.listings.filter(l => l.category_id === catId)
    },

    async updateStatus(id, newStatus) {
      if (!confirm(`Are you sure you want to mark this listing as ${newStatus}?`)) return
      
      try {
        const token = localStorage.getItem('token')
        await axios.patch(
          `http://localhost:3000/api/admin/listings/${id}/status`,
          { status: newStatus },
          { headers: { Authorization: `Bearer ${token}` } }
        )
        
        // Update local state immediately
        const index = this.listings.findIndex(l => l.id === id)
        if (index !== -1) {
          this.listings[index].status = newStatus
        }
        
        alert(`Listing marked as ${newStatus}`)
      } catch (err) {
        console.error('Failed to update status:', err)
        alert('Failed to update status. Please try again.')
      }
    },

    handleLogout() {
      localStorage.removeItem('token')
      this.$router.push('/register-login')
    }
  }
}
</script>

<style scoped>
/* Custom Scrollbar */
::-webkit-scrollbar { width: 8px; height: 8px; }
::-webkit-scrollbar-track { background: #111827; }
::-webkit-scrollbar-thumb { background: #374151; border-radius: 4px; }
::-webkit-scrollbar-thumb:hover { background: #f97316; }
</style>