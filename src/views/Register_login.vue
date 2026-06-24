<template>
  <div class="min-h-screen bg-gray-900 text-white p-4 md:p-8">
    <div class="max-w-4xl mx-auto">
      
      <!-- Flip Card Section -->
      <div class="w-full max-w-md mx-auto mb-12 [perspective:1000px]">
        <div 
          class="relative w-full transition-transform duration-700 ease-in-out [transform-style:preserve-3d]" 
          :class="{ '[transform:rotateY(180deg)]': isFlipped }"
        >
          <!-- Front: Register -->
          <div class="w-full bg-gray-800 rounded-2xl shadow-2xl p-8 border border-gray-700 [backface-visibility:hidden]">
            <h2 class="text-2xl font-bold text-center mb-6 text-orange-500">Create Account</h2>
            <form @submit.prevent="handleRegister" class="space-y-4">
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">First Name</label>
                  <input v-model="registerForm.first_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-300 mb-1">Last Name</label>
                  <input v-model="registerForm.last_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Username</label>
                <input v-model="registerForm.username" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Email</label>
                <input v-model="registerForm.email" type="email" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Phone Number</label>
                <input v-model="registerForm.phone_number" type="tel" class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Password</label>
                <input v-model="registerForm.password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Confirm Password</label>
                <input v-model="registerForm.confirm_password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-2.5 rounded-lg transition-colors mt-2">
                Create Account
              </button>
            </form>
            
            <p class="text-center text-sm text-gray-400 mt-6">
              Already have an account? 
              <button @click="isFlipped = true" class="text-orange-500 font-semibold hover:underline">Login account</button>
            </p>
          </div>

          <!-- Back: Login -->
          <div class="absolute inset-0 w-full bg-gray-800 rounded-2xl shadow-2xl p-8 border border-gray-700 [backface-visibility:hidden] [transform:rotateY(180deg)]">
            <h2 class="text-2xl font-bold text-center mb-6 text-orange-500">Login Account</h2>
            <form @submit.prevent="handleLogin" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Email or Username</label>
                <input v-model="loginForm.identifier" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Password</label>
                <input v-model="loginForm.password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>

              <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-2.5 rounded-lg transition-colors mt-2">
                Log In
              </button>
            </form>
            
            <p class="text-center text-sm text-gray-400 mt-6">
              Don't have an account? 
              <button @click="isFlipped = false" class="text-orange-500 font-semibold hover:underline">Sign up</button>
            </p>
          </div>
        </div>
      </div>

      <!-- API Testing Section -->
      <div class="bg-gray-800 rounded-2xl p-6 border border-gray-700">
        <h1 class="text-2xl font-bold mb-6 text-center">BikeShop - API Test</h1>
        
        <!-- Connection Status -->
        <div class="mb-6 p-4 rounded-lg border" :class="connected ? 'bg-green-900/30 border-green-800 text-green-400' : 'bg-red-900/30 border-red-800 text-red-400'">
          {{ connected ? '✅ Connected to API' : '❌ Not Connected' }}
          <button @click="checkConnection" class="ml-4 underline text-blue-400 hover:text-blue-300">Retry</button>
        </div>

        <!-- Users List -->
        <div class="mb-8">
          <h2 class="text-xl font-semibold mb-4">Users</h2>
          <button @click="loadUsers" class="bg-blue-600 text-white px-4 py-2 rounded-lg mb-4 hover:bg-blue-700 transition-colors">
            Load Users
          </button>
          
          <div v-if="error" class="bg-red-900/30 text-red-400 border border-red-800 p-4 rounded-lg mb-4">
            {{ error }}
          </div>
          
          <div v-if="users && users.length">
            <div v-for="user in users" :key="user.id" class="bg-gray-700/50 border border-gray-600 p-4 mb-2 rounded-lg">
              <p class="font-bold text-white">{{ user.first_name }} {{ user.last_name }}</p>
              <p class="text-gray-400">{{ user.email }}</p>
              <p class="text-sm text-gray-500">@{{ user.username }} | {{ user.role }}</p>
            </div>
          </div>
          <p v-else-if="!error && users && users.length === 0" class="text-gray-500">No users loaded. Click the button above.</p>
          <p v-else-if="!error && (!users || users.length === 0)" class="text-gray-500">Click "Load Users" to fetch data.</p>
        </div>

        <!-- Categories -->
        <div>
          <h2 class="text-xl font-semibold mb-4">Categories</h2>
          <button @click="loadCategories" class="bg-green-600 text-white px-4 py-2 rounded-lg mb-4 hover:bg-green-700 transition-colors">
            Load Categories
          </button>
          
          <div v-if="categories && categories.length" class="flex gap-4 flex-wrap">
            <div v-for="cat in categories" :key="cat.id" 
                 class="border border-gray-600 rounded-lg px-4 py-2 bg-gray-700/50 text-gray-300">
              {{ cat.name }} <span class="text-sm text-gray-500">({{ cat.type }})</span>
            </div>
          </div>
          <p v-else-if="!error && (!categories || categories.length === 0)" class="text-gray-500">Click "Load Categories" to fetch data.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterLoginView',
  
  data() {
    return {
      isFlipped: false,
      registerForm: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
        phone_number: '',
        password: '',
        confirm_password: ''
      },
      loginForm: {
        identifier: '',
        password: ''
      },
      connected: false,
      users: [], // Initialize as empty array
      categories: [], // Initialize as empty array
      error: null,
      loading: false
    }
  },
  
  async mounted() {
    await this.checkConnection()
  },
  
  methods: {
    handleRegister() {
      if (this.registerForm.password !== this.registerForm.confirm_password) {
        alert('Passwords do not match!')
        return
      }
      // TODO: Implement API call to register user
      console.log('Registering user:', this.registerForm)
      alert('Registration logic not implemented yet. Check console for data.')
    },
    
    handleLogin() {
      // TODO: Implement API call to login user
      console.log('Logging in user:', this.loginForm)
      alert('Login logic not implemented yet. Check console for data.')
    },

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
      this.loading = true
      try {
        const response = await axios.get('http://localhost:3000/api/users')
        this.users = response.data || [] // Ensure it's always an array
        console.log('Users loaded:', this.users)
      } catch (err) {
        this.error = `Failed to load users: ${err.message}`
        this.users = [] // Reset to empty array on error
        console.error('Users error:', err)
      } finally {
        this.loading = false
      }
    },
    
    async loadCategories() {
      this.error = null
      this.loading = true
      try {
        const response = await axios.get('http://localhost:3000/api/categories')
        this.categories = response.data || [] // Ensure it's always an array
        console.log('Categories loaded:', this.categories)
      } catch (err) {
        this.error = `Failed to load categories: ${err.message}`
        this.categories = [] // Reset to empty array on error
        console.error('Categories error:', err)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>