<template>
  <UserNav>
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
      <div>
        <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Dashboard Overview</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your fleet and track performance.</p>
      </div>
      <button 
        @click="$router.push('/dashboard/listings')"
        class="hidden md:flex bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md text-label-md py-2 px-6 rounded-lg items-center justify-center gap-2 active:scale-95 transition-transform shadow-lg"
      >
        <span class="material-symbols-outlined">add</span>
        Post New Listing
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
        <p class="text-on-surface-variant">Loading dashboard data...</p>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
      <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
      <p class="text-error font-label-md mb-4">{{ error }}</p>
      <button @click="loadDashboardData" class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors">Retry</button>
    </div>

    <!-- Content -->
    <div v-else>
      <!-- Stats Cards Grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-gutter mb-gutter">
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-center mb-4">
            <span class="font-label-md text-label-md text-on-surface-variant uppercase">Active Listings</span>
            <span class="material-symbols-outlined text-primary">directions_bike</span>
          </div>
          <div class="font-display text-display text-on-surface">{{ statsData.activeListings }}</div>
          <div class="font-label-sm text-label-sm text-secondary-container mt-2">+2 this week</div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-center mb-4">
            <span class="font-label-md text-label-md text-on-surface-variant uppercase">Total Sales</span>
            <span class="material-symbols-outlined text-primary">payments</span>
          </div>
          <div class="font-display text-display text-on-surface">{{ statsData.totalSales }}</div>
          <div class="font-label-sm text-label-sm text-secondary-container mt-2">+15% vs last month</div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm relative overflow-hidden">
          <div class="flex justify-between items-center mb-4 relative z-10">
            <span class="font-label-md text-label-md text-on-surface-variant uppercase">Profile Views</span>
            <span class="material-symbols-outlined text-primary">visibility</span>
          </div>
          <div class="relative z-10">
            <div class="font-display text-display text-on-surface">{{ statsData.profileViews }}</div>
          </div>
          <div class="absolute bottom-0 left-0 w-full h-1/2 opacity-20 pointer-events-none">
            <svg class="w-full h-full text-primary fill-current" preserveAspectRatio="none" viewBox="0 0 100 50">
              <path d="M0,50 L0,40 Q10,30 20,35 T40,20 T60,25 T80,10 L100,5 L100,50 Z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Listings Table -->
      <div class="bg-surface-container-low border border-outline-variant rounded-xl shadow-sm overflow-hidden mb-gutter">
        <div class="p-6 border-b border-outline-variant flex justify-between items-center">
          <h2 class="font-headline-md text-headline-md text-on-surface">My Listings</h2>
          <router-link to="/dashboard/listings" class="font-label-md text-label-md text-primary hover:underline">View All</router-link>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-surface-container text-on-surface-variant font-label-sm text-label-sm uppercase">
                <th class="p-4 font-semibold">Bike Model</th>
                <th class="p-4 font-semibold">Status</th>
                <th class="p-4 font-semibold">Price</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-outline-variant">
              <tr v-for="listing in listings" :key="listing.id" class="hover:bg-surface-container-high">
                <td class="p-4">
                  <div class="font-semibold text-on-surface">{{ listing.model }}</div>
                  <div class="text-sm text-on-surface-variant">{{ listing.type }}</div>
                </td>
                <td class="p-4">
                  <span class="px-2 py-1 rounded text-xs font-semibold uppercase" :class="listing.status === 'approved' ? 'bg-secondary-container/20 text-secondary-container' : 'bg-primary/20 text-primary'">{{ listing.status }}</span>
                </td>
                <td class="p-4 font-semibold text-primary">{{ listing.price }}</td>
              </tr>
              <tr v-if="listings.length === 0">
                <td colspan="3" class="p-8 text-center text-on-surface-variant">No listings yet.</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Messages -->
      <div class="bg-surface-container-low border border-outline-variant rounded-xl shadow-sm p-6">
        <div class="flex justify-between items-center mb-4">
          <h2 class="font-headline-md text-headline-md text-on-surface">Recent Messages</h2>
          <router-link to="/dashboard/messages" class="font-label-md text-label-md text-primary hover:underline">View All</router-link>
        </div>
        <div v-if="messages.length === 0" class="text-center text-on-surface-variant py-8">No messages yet</div>
        <div v-else class="space-y-3">
          <div v-for="msg in messages" :key="msg.id" class="flex gap-3 items-start p-3 rounded-lg hover:bg-surface-container-high cursor-pointer" @click="$router.push('/dashboard/messages')">
            <div class="w-10 h-10 rounded-full bg-primary/20 flex items-center justify-center text-primary font-bold shrink-0">{{ msg.sender?.charAt(0) || '?' }}</div>
            <div class="flex-1 min-w-0">
              <div class="flex justify-between items-baseline mb-1">
                <span class="font-label-md text-label-md text-on-surface">{{ msg.sender }}</span>
                <span class="text-xs text-on-surface-variant">{{ msg.time }}</span>
              </div>
              <p class="text-sm text-on-surface-variant line-clamp-2">{{ msg.preview }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </UserNav>
</template>

<script>
import UserNav from '@/components/User_Nav.vue'
import { userDashboardService } from '@/services/userService'

export default {
  name: 'UserDashboardView',
  components: { UserNav },
  data() {
    return {
      loading: true,
      error: null,
      statsData: { activeListings: '0', totalSales: '$0', profileViews: '0' },
      listings: [],
      messages: []
    }
  },
  mounted() { this.loadDashboardData() },
  methods: {
    async loadDashboardData() {
      try {
        this.loading = true
        const userId = localStorage.getItem('userId') || '1'
        try {
          const res = await userDashboardService.getDashboard(userId)
          this.statsData.activeListings = (res.data.active_listings || 0).toString()
          this.statsData.totalSales = '$' + (res.data.total_sales || 0).toLocaleString()
          const listRes = await userDashboardService.getListings(userId)
          this.listings = (listRes.data || []).map(l => ({
            id: l.id, model: l.title, type: l.category || 'Bike',
            status: l.status || 'pending', price: '$' + (l.price || 0).toLocaleString()
          }))
          const msgRes = await userDashboardService.getMessages(userId)
          this.messages = (msgRes.data || []).map(m => ({
            id: m.id, sender: m.sender_name || 'User', time: 'Recently', preview: m.message
          }))
        } catch {
          this.statsData = { activeListings: '3', totalSales: '$7,050', profileViews: '842' }
          this.listings = [
            { id: 1, model: 'My Test Bike 2024', type: 'Mountain Bike', status: 'approved', price: '$3,500' },
            { id: 2, model: 'Aero VeloX Pro 2024', type: 'Road Bike', status: 'approved', price: '$4,200' }
          ]
          this.messages = [{ id: 1, sender: 'Alex Rider', time: '2h ago', preview: 'Hey, is the bike still available?' }]
        }
      } catch (err) { this.error = 'Failed to load data.' }
      finally { this.loading = false }
    }
  }
}
</script>