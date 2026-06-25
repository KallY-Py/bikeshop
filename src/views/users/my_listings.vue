<template>
  <div class="bg-background text-on-background font-body-md min-h-screen flex flex-col md:flex-row">
    <!-- TopNavBar (Mobile Only) -->
    <header class="md:hidden flex justify-between items-center px-margin-mobile h-16 w-full bg-background/80 backdrop-blur-md border-b border-outline-variant shadow-none sticky top-0 z-50">
      <div class="font-display text-headline-lg-mobile text-primary uppercase tracking-tighter">VeloHub</div>
      <div class="flex gap-4">
        <span class="material-symbols-outlined text-primary cursor-pointer">notifications</span>
        <span class="material-symbols-outlined text-primary cursor-pointer">account_circle</span>
      </div>
    </header>

    <!-- SideNavBar (Desktop) -->
    <nav class="hidden md:flex flex-col h-screen py-8 fixed left-0 top-0 w-[280px] bg-surface-container-low border-r border-outline-variant shadow-sm z-40">
      <div class="px-6 mb-8 flex items-center gap-4">
        <div class="w-12 h-12 rounded-full overflow-hidden bg-surface border border-outline-variant">
          <img alt="User profile" class="w-full h-full object-cover" src="https://via.placeholder.com/48" />
        </div>
        <div>
          <div class="font-label-md text-label-md text-on-surface-variant">Welcome back</div>
          <div class="font-headline-md text-headline-md text-primary">Rider Elite Status</div>
        </div>
      </div>

      <div class="flex-1 overflow-y-auto mt-4">
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
        </router-link>
      </div>

      <div class="px-6 mt-auto">
        <button 
          @click="showAddModal = true"
          class="w-full bg-primary-container text-on-primary-container font-label-md text-label-md py-3 rounded-lg flex items-center justify-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] active:scale-95 transition-transform shadow-lg"
        >
          <span class="material-symbols-outlined">add</span>
          Post New Bike
        </button>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
        <div>
          <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">My Listings</h1>
          <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your bike listings and track their status.</p>
        </div>
        <button 
          @click="showAddModal = true"
          class="flex md:hidden w-full bg-primary-container text-on-primary-container font-label-md text-label-md py-3 rounded-lg items-center justify-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] active:scale-95 transition-transform shadow-lg"
        >
          <span class="material-symbols-outlined">add</span>
          Post New Listing
        </button>
      </div>

      <!-- Filters and Search -->
      <div class="flex flex-col sm:flex-row gap-4 mb-6">
        <div class="flex-1 relative">
          <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant">search</span>
          <input 
            v-model="searchQuery"
            type="text"
            placeholder="Search listings..."
            class="w-full bg-surface-container-low border border-outline-variant rounded-lg py-2.5 pl-10 pr-4 text-on-surface placeholder-on-surface-variant focus:border-primary focus:ring-1 focus:ring-primary transition-colors"
          />
        </div>
        <select 
          v-model="statusFilter"
          class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors"
        >
          <option value="">All Status</option>
          <option value="active">Active</option>
          <option value="pending">Pending</option>
          <option value="sold">Sold</option>
          <option value="rejected">Rejected</option>
        </select>
        <select 
          v-model="sortBy"
          class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors"
        >
          <option value="newest">Newest First</option>
          <option value="oldest">Oldest First</option>
          <option value="price_high">Price: High to Low</option>
          <option value="price_low">Price: Low to High</option>
          <option value="views">Most Viewed</option>
        </select>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
          <p class="text-on-surface-variant">Loading your listings...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
        <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
        <p class="text-error font-label-md mb-4">{{ error }}</p>
        <button 
          @click="loadListings"
          class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors"
        >
          Retry
        </button>
      </div>

      <!-- Listings Grid -->
      <div v-else>
        <div v-if="filteredListings.length === 0" class="text-center py-16">
          <span class="material-symbols-outlined text-6xl text-on-surface-variant mb-4">directions_bike</span>
          <h3 class="font-headline-md text-headline-md text-on-surface mb-2">No listings found</h3>
          <p class="text-on-surface-variant mb-6">{{ searchQuery || statusFilter ? 'Try adjusting your filters' : 'Start selling your bikes today!' }}</p>
          <button 
            v-if="!searchQuery && !statusFilter"
            @click="showAddModal = true"
            class="bg-primary-container text-on-primary-container font-label-md text-label-md py-2 px-6 rounded-lg bg-gradient-to-b from-[#FF8A00] to-[#FF6B00]"
          >
            Post Your First Listing
          </button>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-gutter">
          <ListingCard 
            v-for="listing in filteredListings" 
            :key="listing.id"
            :listing="listing"
            @edit="editListing"
            @delete="confirmDelete"
            @view="viewListing"
          />
        </div>

        <!-- Pagination -->
        <div v-if="totalPages > 1" class="flex justify-center items-center gap-2 mt-8">
          <button 
            @click="currentPage--"
            :disabled="currentPage === 1"
            class="px-4 py-2 rounded-lg bg-surface-container-low border border-outline-variant text-on-surface hover:bg-surface-container-high disabled:opacity-50 transition-colors"
          >
            Previous
          </button>
          <span class="text-on-surface-variant px-4">
            Page {{ currentPage }} of {{ totalPages }}
          </span>
          <button 
            @click="currentPage++"
            :disabled="currentPage === totalPages"
            class="px-4 py-2 rounded-lg bg-surface-container-low border border-outline-variant text-on-surface hover:bg-surface-container-high disabled:opacity-50 transition-colors"
          >
            Next
          </button>
        </div>
      </div>
    </main>

    <!-- Add/Edit Listing Modal -->
    <ListingModal 
      v-if="showAddModal || showEditModal"
      :listing="editingListing"
      :categories="categories"
      @close="closeModal"
      @save="saveListing"
    />

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4">
      <div class="bg-surface-container-high border border-outline-variant rounded-xl p-6 max-w-md w-full">
        <h3 class="font-headline-md text-headline-md text-on-surface mb-4">Delete Listing</h3>
        <p class="text-on-surface-variant mb-6">Are you sure you want to delete "{{ deletingListing?.title }}"? This action cannot be undone.</p>
        <div class="flex gap-3 justify-end">
          <button 
            @click="showDeleteModal = false"
            class="px-4 py-2 rounded-lg bg-surface-variant text-on-surface hover:bg-surface-bright transition-colors"
          >
            Cancel
          </button>
          <button 
            @click="deleteListing"
            class="px-4 py-2 rounded-lg bg-error text-on-error hover:bg-error-container transition-colors"
          >
            Delete
          </button>
        </div>
      </div>
    </div>

    <!-- BottomNavBar (Mobile Only) -->
    <nav class="fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-4 py-3 md:hidden bg-surface-container-lowest/95 backdrop-blur-lg border-t border-outline-variant shadow-[0_-4px_10px_rgba(0,0,0,0.3)] rounded-t-xl">
      <router-link 
        v-for="item in mobileNavItems" 
        :key="item.name"
        :to="item.path"
        class="flex flex-col items-center justify-center rounded-xl px-4 py-1 active:scale-90 transition-transform duration-150"
        :class="$route.path === item.path 
          ? 'text-primary bg-primary/10' 
          : 'text-on-surface-variant active:bg-surface-container-high'"
      >
        <span class="material-symbols-outlined">{{ item.icon }}</span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </router-link>
    </nav>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { userDashboardService } from '@/services/userService'
import { categoryService } from '@/services/userService'
import ListingCard from '@/components/user/ListingCard.vue'
import ListingModal from '@/components/user/ListingModal.vue'

const router = useRouter()

// State
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')
const statusFilter = ref('')
const sortBy = ref('newest')
const currentPage = ref(1)
const itemsPerPage = 9

// Modal states
const showAddModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const editingListing = ref(null)
const deletingListing = ref(null)

// Data
const listings = ref([])
const categories = ref([])

// Navigation
const navItems = [
  { name: 'Overview', icon: 'dashboard', path: '/dashboard' },
  { name: 'My Listings', icon: 'directions_bike', path: '/dashboard/listings' },
  { name: 'Sales', icon: 'payments', path: '/dashboard/sales' },
  { name: 'Analytics', icon: 'leaderboard', path: '/dashboard/analytics' },
  { name: 'Messages', icon: 'chat', path: '/dashboard/messages' }
]

const mobileNavItems = [
  { name: 'Home', icon: 'home', path: '/dashboard' },
  { name: 'Listings', icon: 'storefront', path: '/dashboard/listings' },
  { name: 'Chat', icon: 'forum', path: '/dashboard/messages' },
  { name: 'Profile', icon: 'person', path: '/profile' }
]

// Computed
const filteredListings = computed(() => {
  let filtered = [...listings.value]
  
  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(listing => 
      listing.title.toLowerCase().includes(query) ||
      listing.category?.toLowerCase().includes(query)
    )
  }
  
  // Status filter
  if (statusFilter.value) {
    filtered = filtered.filter(listing => listing.status === statusFilter.value)
  }
  
  // Sort
  switch (sortBy.value) {
    case 'oldest':
      filtered.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
      break
    case 'price_high':
      filtered.sort((a, b) => b.price - a.price)
      break
    case 'price_low':
      filtered.sort((a, b) => a.price - b.price)
      break
    case 'views':
      filtered.sort((a, b) => (b.views || 0) - (a.views || 0))
      break
    default: // newest
      filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }
  
  return filtered
})

const totalPages = computed(() => Math.ceil(filteredListings.value.length / itemsPerPage))

// Methods
const loadListings = async () => {
  try {
    loading.value = true
    error.value = null
    
    const userId = localStorage.getItem('userId') || '1'
    const response = await userDashboardService.getListings(userId)
    listings.value = response.data || []
    
  } catch (err) {
    console.error('Failed to load listings:', err)
    error.value = 'Failed to load listings. Please make sure the server is running.'
  } finally {
    loading.value = false
  }
}

const loadCategories = async () => {
  try {
    const response = await categoryService.getAll()
    categories.value = response.data || []
  } catch (err) {
    console.error('Failed to load categories:', err)
  }
}

const editListing = (listing) => {
  editingListing.value = { ...listing }
  showEditModal.value = true
}

const confirmDelete = (listing) => {
  deletingListing.value = listing
  showDeleteModal.value = true
}

const deleteListing = async () => {
  try {
    // TODO: Add delete API endpoint
    listings.value = listings.value.filter(l => l.id !== deletingListing.value.id)
    showDeleteModal.value = false
    deletingListing.value = null
  } catch (err) {
    console.error('Failed to delete listing:', err)
  }
}

const saveListing = async (listingData) => {
  try {
    if (editingListing.value?.id) {
      // Update existing listing
      const index = listings.value.findIndex(l => l.id === editingListing.value.id)
      if (index !== -1) {
        listings.value[index] = { ...listings.value[index], ...listingData }
      }
    } else {
      // Add new listing
      const userId = localStorage.getItem('userId') || '1'
      const response = await userDashboardService.createListing({
        ...listingData,
        user_id: parseInt(userId)
      })
      listings.value.unshift(response.data)
    }
    closeModal()
    await loadListings() // Reload to get fresh data
  } catch (err) {
    console.error('Failed to save listing:', err)
  }
}

const closeModal = () => {
  showAddModal.value = false
  showEditModal.value = false
  editingListing.value = null
}

const viewListing = (listing) => {
  router.push(`/listings/${listing.id}`)
}

// Lifecycle
onMounted(() => {
  loadListings()
  loadCategories()
})
</script>