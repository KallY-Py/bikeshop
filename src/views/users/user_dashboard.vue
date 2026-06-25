<!-- <template>
  <div class="bg-background text-on-background font-body-md min-h-screen flex flex-col md:flex-row">
    <header class="md:hidden flex justify-between items-center px-margin-mobile h-16 w-full bg-background/80 backdrop-blur-md border-b border-outline-variant shadow-none sticky top-0 z-50">
      <div class="font-display text-headline-lg-mobile text-primary uppercase tracking-tighter">VeloHub</div>
      <div class="flex gap-4">
        <span 
          class="material-symbols-outlined text-primary cursor-pointer active:scale-95 transition-transform"
          @click="toggleFavorite"
        >
          favorite
        </span>
        <span 
          class="material-symbols-outlined text-primary cursor-pointer active:scale-95 transition-transform"
          @click="goToProfile"
        >
          account_circle
        </span>
      </div>
    </header>

    <nav class="hidden md:flex flex-col h-screen py-8 fixed left-0 top-0 w-[280px] bg-surface-container-low border-r border-outline-variant shadow-sm z-40">
      <div class="px-6 mb-8 flex items-center gap-4">
        <div class="w-12 h-12 rounded-full overflow-hidden bg-surface border border-outline-variant">
          <img 
            alt="User profile photo" 
            class="w-full h-full object-cover" 
            :src="userProfile.image || 'https://via.placeholder.com/48'"
          />
        </div>
        <div>
          <div class="font-label-md text-label-md text-on-surface-variant">Welcome back</div>
          <div class="font-headline-md text-headline-md text-primary">{{ userProfile.status || 'Rider' }}</div>
        </div>
      </div>

      <div class="flex-1 overflow-y-auto mt-4">
        <a 
          v-for="item in navItems" 
          :key="item.name"
          :href="item.href"
          @click.prevent="setActiveNav(item.name)"
          class="flex items-center gap-3 py-3 px-6 font-label-md text-label-md transition-all"
          :class="activeNav === item.name 
            ? 'text-primary border-l-4 border-primary bg-primary/10' 
            : 'text-on-surface-variant hover:text-on-surface hover:bg-surface-container-high active:translate-x-1'"
        >
          <span class="material-symbols-outlined">{{ item.icon }}</span>
          <span>{{ item.name }}</span>
        </a>
      </div>

      <div class="px-6 mt-auto">
        <button 
          @click="postNewBike"
          class="w-full bg-primary-container text-on-primary-container font-label-md text-label-md py-3 rounded-lg flex items-center justify-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] active:scale-95 transition-transform shadow-lg"
        >
          <span class="material-symbols-outlined">add</span>
          Post New Bike
        </button>
        
        <div class="mt-6 border-t border-outline-variant pt-4 space-y-2">
          <a 
            v-for="item in bottomNavItems" 
            :key="item.name"
            :href="item.href"
            class="flex items-center gap-3 text-on-surface-variant hover:text-on-surface py-2 font-label-md text-label-md transition-colors"
          >
            <span class="material-symbols-outlined">{{ item.icon }}</span>
            <span>{{ item.name }}</span>
          </a>
        </div>
      </div>
    </nav>

    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
        <div>
          <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Dashboard Overview</h1>
          <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your fleet and track performance.</p>
        </div>
        <button 
          @click="postNewListing"
          class="hidden md:flex bg-primary-container text-on-primary-container font-label-md text-label-md py-2 px-6 rounded-lg items-center justify-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] active:scale-95 transition-transform shadow-lg"
        >
          <span class="material-symbols-outlined">add</span>
          Post New Listing
        </button>
      </div>

      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
          <p class="text-on-surface-variant">Loading dashboard data...</p>
        </div>
      </div>

      <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
        <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
        <p class="text-error font-label-md mb-4">{{ error }}</p>
        <button 
          @click="loadDashboardData"
          class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors"
        >
          Retry
        </button>
      </div>

      <div v-else>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-gutter mb-gutter">
          <StatCard 
            v-for="stat in statsData" 
            :key="stat.title"
            :title="stat.title"
            :value="stat.value"
            :change="stat.change"
            :icon="stat.icon"
            :chart="stat.chart"
          />
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-gutter">
          <div class="lg:col-span-2 bg-surface-container-low border border-outline-variant rounded-xl shadow-sm overflow-hidden flex flex-col">
            <div class="p-6 border-b border-outline-variant flex justify-between items-center">
              <h2 class="font-headline-md text-headline-md text-on-surface">My Listings</h2>
              <a 
                @click.prevent="viewAllListings"
                href="#" 
                class="font-label-md text-label-md text-primary hover:underline cursor-pointer"
              >
                View All
              </a>
            </div>
            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="bg-surface-container text-on-surface-variant font-label-sm text-label-sm uppercase">
                    <th class="p-4 font-semibold">Bike Model</th>
                    <th class="p-4 font-semibold">Status</th>
                    <th class="p-4 font-semibold">Price</th>
                    <th class="p-4 font-semibold text-right">Actions</th>
                  </tr>
                </thead>
                <tbody class="font-body-md text-body-md text-on-surface divide-y divide-outline-variant">
                  <ListingRow 
                    v-for="listing in listings" 
                    :key="listing.id"
                    :listing="listing"
                    @edit="editListing"
                  />
                  <tr v-if="listings.length === 0">
                    <td colspan="4" class="p-8 text-center text-on-surface-variant">
                      <span class="material-symbols-outlined text-4xl mb-2 block">directions_bike</span>
                      No listings yet. Start selling your bikes!
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <div class="bg-surface-container-low border border-outline-variant rounded-xl shadow-sm flex flex-col">
            <div class="p-6 border-b border-outline-variant flex justify-between items-center">
              <h2 class="font-headline-md text-headline-md text-on-surface">Messages</h2>
              <span class="material-symbols-outlined text-primary">chat</span>
            </div>
            <div class="p-4 space-y-4 flex-1 overflow-y-auto max-h-[300px]">
              <MessageItem 
                v-for="message in messages" 
                :key="message.id"
                :message="message"
                @click="openConversation(message.id)"
              />
              <div v-if="messages.length === 0" class="text-center text-on-surface-variant py-8">
                <span class="material-symbols-outlined text-3xl mb-2 block">forum</span>
                No messages yet
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <nav class="fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-4 py-3 md:hidden bg-surface-container-lowest/95 backdrop-blur-lg border-t border-outline-variant shadow-[0_-4px_10px_rgba(0,0,0,0.3)] rounded-t-xl">
      <a 
        v-for="item in mobileNavItems" 
        :key="item.name"
        :href="item.href"
        @click.prevent="setActiveNav(item.name)"
        class="flex flex-col items-center justify-center rounded-xl px-4 py-1 active:scale-90 transition-transform duration-150"
        :class="activeNav === item.name 
          ? 'text-primary bg-primary/10' 
          : 'text-on-surface-variant active:bg-surface-container-high'"
      >
        <span 
          class="material-symbols-outlined"
          :style="activeNav === item.name ? 'font-variation-settings: \'FILL\' 1;' : ''"
        >
          {{ item.icon }}
        </span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </a>
    </nav>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
// import { userDashboardService } from '@/services/userService'
// import StatCard from '@/components/user/StatCard.vue'
// import ListingRow from '@/components/user/ListingRow.vue'
// import MessageItem from '@/components/user/MessageItem.vue'

const router = useRouter()

const activeNav = ref('Overview')
const loading = ref(true)
const error = ref(null)
const userProfile = ref({
  image: null,
  status: 'Rider Elite Status'
})

const statsData = ref([
  {
    title: 'Active Listings',
    value: '0',
    change: null,
    icon: 'directions_bike',
    changePositive: true,
    chart: false
  },
  {
    title: 'Total Sales Volume',
    value: '$0',
    change: null,
    icon: 'payments',
    changePositive: true,
    chart: false
  },
  {
    title: 'Unread Messages',
    value: '0',
    change: null,
    icon: 'visibility',
    changePositive: true,
    chart: true
  }
])

const listings = ref([])
const messages = ref([])

const navItems = [
  { name: 'Overview', icon: 'dashboard', href: '#' },
  { name: 'My Listings', icon: 'directions_bike', href: '#' },
  { name: 'Sales', icon: 'payments', href: '#' },
  { name: 'Analytics', icon: 'leaderboard', href: '#' },
  { name: 'Messages', icon: 'chat', href: '#' }
]

const bottomNavItems = [
  { name: 'Settings', icon: 'settings', href: '#' },
  { name: 'Support', icon: 'help', href: '#' }
]

const mobileNavItems = [
  { name: 'Home', icon: 'home', href: '#' },
  { name: 'Shop', icon: 'storefront', href: '#' },
  { name: 'Chat', icon: 'forum', href: '#' },
  { name: 'Profile', icon: 'person', href: '#' }
]

const setActiveNav = (name) => {
  activeNav.value = name
}

const toggleFavorite = () => {
  console.log('Toggle favorite')
}

const goToProfile = () => {
  router.push('/profile')
}

const postNewBike = () => {
  console.log('Post new bike')
}

const postNewListing = () => {
  console.log('Post new listing')
}

const viewAllListings = () => {
  setActiveNav('My Listings')
}

const editListing = (listingId) => {
  console.log('Edit listing:', listingId)
}

const openConversation = (messageId) => {
  console.log('Open conversation:', messageId)
}

const loadDashboardData = async () => {
  try {
    loading.value = true
    error.value = null
    
    const userId = localStorage.getItem('userId') || '1'
    
    console.log('Fetching dashboard for user:', userId)
    
    // const dashboardResponse = await userDashboardService.getDashboard(userId)
    const dashboardData = dashboardResponse.data
    
    console.log('Dashboard data received:', dashboardData)
    
    statsData.value[0].value = (dashboardData.active_listings || 0).toString()
    
    const totalSales = dashboardData.total_sales || 0
    statsData.value[1].value = `$${totalSales.toLocaleString()}`
    
    statsData.value[2].value = (dashboardData.unread_messages || 0).toString()
    statsData.value[2].title = 'Unread Messages'
    statsData.value[2].icon = 'mail'
    
    console.log('Fetching listings...')
    // const listingsResponse = await userDashboardService.getListings(userId)
    const listingsData = listingsResponse.data || []
    console.log('Listings received:', listingsData)
    
    listings.value = listingsData.map(listing => ({
      id: listing.id,
      model: listing.title,
      type: `${listing.category || 'Bike'} • ${listing.condition_type || 'Used'}`,
      image: listing.image_url || 'https://via.placeholder.com/48',
      status: listing.status || 'pending',
      price: `$${(listing.price || 0).toLocaleString()}`
    }))
    
    console.log('Fetching messages...')
    // const messagesResponse = await userDashboardService.getMessages(userId)
    const messagesData = messagesResponse.data || []
    console.log('Messages received:', messagesData)
    
    messages.value = messagesData.map(msg => ({
      id: msg.id,
      sender: {
        name: msg.sender_name || 'Unknown User',
        avatar: msg.sender_image || 'https://via.placeholder.com/40'
      },
      time: formatTimeAgo(msg.sent_at),
      preview: msg.message
    }))
    
  } catch (err) {
    console.error('Failed to load dashboard data:', err)
    error.value = `Failed to connect to server: ${err.message}. Make sure the Go backend is running on port 3000.`
    
    statsData.value[0].value = '0'
    statsData.value[1].value = '$0'
    statsData.value[2].value = '0'
    listings.value = []
    messages.value = []
    
  } finally {
    loading.value = false
  }
}

const formatTimeAgo = (dateString) => {
  if (!dateString) return 'Recently'
  
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMs / 3600000)
  const diffDays = Math.floor(diffMs / 86400000)
  
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffHours < 24) return `${diffHours}h ago`
  if (diffDays < 7) return `${diffDays}d ago`
  return date.toLocaleDateString()
}

onMounted(() => {
  loadDashboardData()
})
</script> -->