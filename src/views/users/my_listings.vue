<template>
  <UserNav>
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
      <div>
        <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">My Listings</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your bike listings and track their status.</p>
      </div>
      <button 
        @click="showAddModal = true"
        class="flex md:hidden w-full bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md text-label-md py-3 rounded-lg items-center justify-center gap-2"
      >
        <span class="material-symbols-outlined">add</span>
        Post New Listing
      </button>
    </div>

    <!-- Filters -->
    <div class="flex flex-col sm:flex-row gap-4 mb-6">
      <div class="flex-1 relative">
        <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant">search</span>
        <input v-model="searchQuery" type="text" placeholder="Search listings..." class="w-full bg-surface-container-low border border-outline-variant rounded-lg py-2.5 pl-10 pr-4 text-on-surface placeholder-on-surface-variant focus:border-primary transition-colors" />
      </div>
      <select v-model="statusFilter" class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors">
        <option value="">All Status</option>
        <option value="approved">Active</option>
        <option value="pending">Pending</option>
        <option value="sold">Sold</option>
        <option value="rejected">Rejected</option>
      </select>
      <select v-model="sortBy" class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors">
        <option value="newest">Newest First</option>
        <option value="oldest">Oldest First</option>
        <option value="price_high">Price: High to Low</option>
        <option value="price_low">Price: Low to High</option>
        <option value="views">Most Viewed</option>
      </select>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
        <p class="text-on-surface-variant">Loading your listings...</p>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
      <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
      <p class="text-error font-label-md mb-4">{{ error }}</p>
      <button @click="loadListings" class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg">Retry</button>
    </div>

    <!-- Listings Grid -->
    <div v-else>
      <div v-if="filteredListings.length === 0" class="text-center py-16">
        <span class="material-symbols-outlined text-6xl text-on-surface-variant mb-4">directions_bike</span>
        <h3 class="font-headline-md text-headline-md text-on-surface mb-2">No listings found</h3>
        <p class="text-on-surface-variant mb-6">{{ searchQuery || statusFilter ? 'Try adjusting your filters' : 'Start selling your bikes today!' }}</p>
        <button v-if="!searchQuery && !statusFilter" @click="showAddModal = true" class="bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md py-2 px-6 rounded-lg">Post Your First Listing</button>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-gutter">
        <ListingCard v-for="listing in filteredListings" :key="listing.id" :listing="listing" @edit="editListing" @delete="confirmDelete" />
      </div>
    </div>

    <!-- Add/Edit Modal -->
    <ListingModal v-if="showAddModal || showEditModal" :listing="editingListing" :categories="categories" :saving="saving" @close="closeModal" @save="saveListing" />

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4">
      <div class="bg-surface-container-high border border-outline-variant rounded-xl p-6 max-w-md w-full">
        <h3 class="font-headline-md text-headline-md text-on-surface mb-4">Delete Listing</h3>
        <p class="text-on-surface-variant mb-6">Are you sure you want to delete "{{ deletingListing?.title }}"?</p>
        <div class="flex gap-3 justify-end">
          <button @click="showDeleteModal = false" class="px-4 py-2 rounded-lg bg-surface-variant text-on-surface hover:bg-surface-bright transition-colors">Cancel</button>
          <button @click="deleteListing" class="px-4 py-2 rounded-lg bg-error text-on-error hover:bg-error-container transition-colors">Delete</button>
        </div>
      </div>
    </div>

    <!-- Floating Action Button -->
    <button @click="showAddModal = true" class="md:hidden fixed bottom-24 right-6 w-14 h-14 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black rounded-full shadow-lg hover:scale-110 active:scale-95 transition-all z-50 flex items-center justify-center">
      <span class="material-symbols-outlined text-2xl">add</span>
    </button>
    <button @click="showAddModal = true" class="hidden md:flex fixed bottom-8 right-8 items-center gap-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md text-label-md py-3 px-6 rounded-full shadow-lg hover:scale-105 active:scale-95 transition-all z-50">
      <span class="material-symbols-outlined">add</span>
      <span>Post New Bike</span>
    </button>
  </UserNav>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { userDashboardService, categoryService } from '@/services/userService'
import UserNav from '@/components/User_Nav.vue'
import ListingCard from '@/components/user/ListingCard.vue'
import ListingModal from '@/components/user/ListingModal.vue'
import axios from 'axios'

const API_URL = 'http://localhost:3000/api'

const loading = ref(true)
const error = ref(null)
const saving = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')
const sortBy = ref('newest')
const showAddModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const editingListing = ref(null)
const deletingListing = ref(null)
const listings = ref([])
const categories = ref([])

const filteredListings = computed(() => {
  let filtered = [...listings.value]
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    filtered = filtered.filter(l => (l.title || '').toLowerCase().includes(q) || (l.category || '').toLowerCase().includes(q))
  }
  if (statusFilter.value) filtered = filtered.filter(l => l.status === statusFilter.value)
  switch (sortBy.value) {
    case 'oldest': filtered.sort((a, b) => new Date(a.created_at) - new Date(b.created_at)); break
    case 'price_high': filtered.sort((a, b) => (b.price || 0) - (a.price || 0)); break
    case 'price_low': filtered.sort((a, b) => (a.price || 0) - (b.price || 0)); break
    case 'views': filtered.sort((a, b) => (b.views || 0) - (a.views || 0)); break
    default: filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }
  return filtered
})

  const loadListings = async () => {
    try {
      loading.value = true
      error.value = null
      const userId = localStorage.getItem('userId') || '4'
      
      const response = await userDashboardService.getListings(userId)
      
      // Axios puts the JSON response in response.data
      // The API returns a JSON array directly
      listings.value = response.data
      
      console.log('Listings loaded:', listings.value.length, listings.value)
      
    } catch (err) {
      console.error('Failed to load listings:', err)
      error.value = 'Failed to load listings.'
    } finally { 
      loading.value = false 
    }
  }

const loadCategories = async () => {
  try { 
    const r = await categoryService.getAll()
    categories.value = r.data || [] 
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
    const id = deletingListing.value.id
    await axios.delete(`${API_URL}/listings/${id}`)
    listings.value = listings.value.filter(l => l.id !== id)
    showDeleteModal.value = false
    deletingListing.value = null
  } catch (err) {
    console.error('Failed to delete listing:', err)
    alert('Failed to delete listing. Please try again.')
  }
}

const closeModal = () => { 
  showAddModal.value = false
  showEditModal.value = false
  editingListing.value = null
}

const saveListing = async (listingData) => {
  try {
    saving.value = true
    const userId = localStorage.getItem('userId') || '1'

    const payload = {
      user_id: parseInt(userId),
      category_id: parseInt(listingData.category_id) || 1,
      title: listingData.title,
      description: listingData.description || '',
      price: parseFloat(listingData.price) || 0,
      condition_type: listingData.condition_type || 'used',
      location: listingData.location || '',
      image_url: listingData.image_url || ''
    }

    console.log('Saving listing:', payload)

    if (editingListing.value?.id) {
      await axios.put(`${API_URL}/listings/${editingListing.value.id}`, payload)
      console.log('Listing updated')
    } else {
      const response = await axios.post(`${API_URL}/listings`, payload)
      console.log('Listing created:', response.data)
    }
    
    closeModal()
    await loadListings() // Refresh the list
    
  } catch (err) {
    console.error('Failed to save listing:', err)
    const errorMsg = err.response?.data?.error || 'Failed to save listing. Please try again.'
    alert(errorMsg)
  } finally { 
    saving.value = false 
  }
}

onMounted(() => { 
  loadListings()
  loadCategories() 
})
</script>