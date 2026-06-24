<template>
  <div class="max-w-4xl mx-auto p-8">
    <h1 class="text-3xl font-bold mb-6">BikeShop - API Test</h1>
    
    <!-- Connection Status -->
    <div class="mb-6 p-4 rounded" :class="connected ? 'bg-green-100' : 'bg-red-100'">
      {{ connected ? '✅ Connected to API' : '❌ Not Connected' }}
      <button @click="checkConnection" class="ml-4 underline text-blue-600">Retry</button>
    </div>

    <!-- Users List -->
    <div class="mb-8">
      <h2 class="text-xl font-semibold mb-4">Users</h2>
      <button @click="loadUsers" class="bg-blue-500 text-white px-4 py-2 rounded mb-4 hover:bg-blue-600">
        Load Users
      </button>
      
      <div v-if="error" class="bg-red-100 text-red-700 p-4 rounded mb-4">
        {{ error }}
      </div>
      
      <div v-if="users.length">
        <div v-for="user in users" :key="user.id" class="border p-4 mb-2 rounded">
          <p class="font-bold">{{ user.first_name }} {{ user.last_name }}</p>
          <p class="text-gray-600">{{ user.email }}</p>
          <p class="text-sm text-gray-500">@{{ user.username }} | {{ user.role }}</p>
        </div>
      </div>
      <p v-else-if="!error" class="text-gray-500">No users loaded. Click the button above.</p>
    </div>

    <!-- Categories -->
    <div>
      <h2 class="text-xl font-semibold mb-4">Categories</h2>
      <button @click="loadCategories" class="bg-green-500 text-white px-4 py-2 rounded mb-4 hover:bg-green-600">
        Load Categories
      </button>
      
      <div v-if="categories.length" class="flex gap-4 flex-wrap">
        <div v-for="cat in categories" :key="cat.id" 
             class="border rounded px-4 py-2 bg-gray-50">
          {{ cat.name }} <span class="text-sm text-gray-500">({{ cat.type }})</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HomeView',
  
  data() {
    return {
      connected: false,
      users: [],
      categories: [],
      error: null
    }
  },
  
  async mounted() {
    await this.checkConnection()
  },
  
  methods: {
    async checkConnection() {
      this.error = null
      try {
        const response = await axios.get('http://localhost:3000/api/health')
        this.connected = response.data.status === 'ok'
        console.log('Connection check:', response.data)
      } catch (err) {
        this.connected = false
        this.error = `Connection failed: ${err.message}`
        console.error('Connection error:', err)
      }
    },
    
    async loadUsers() {
      this.error = null
      try {
        const response = await axios.get('http://localhost:3000/api/users')
        this.users = response.data
        console.log('Users loaded:', this.users)
      } catch (err) {
        this.error = `Failed to load users: ${err.message}`
        console.error('Users error:', err)
      }
    },
    
    async loadCategories() {
      this.error = null
      try {
        const response = await axios.get('http://localhost:3000/api/categories')
        this.categories = response.data
        console.log('Categories loaded:', this.categories)
      } catch (err) {
        this.error = `Failed to load categories: ${err.message}`
        console.error('Categories error:', err)
      }
    }
  }
}
</script>