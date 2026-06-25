<template>
  <div class="min-h-screen bg-gray-900 text-white flex overflow-hidden">
    
    <aside 
      class="w-64 bg-gray-800 border-r border-gray-700 flex-shrink-0 hidden md:flex flex-col transition-all duration-300"
    >
      <div class="h-16 flex items-center justify-center border-b border-gray-700">
        <h1 class="text-2xl font-bold text-orange-500 tracking-wider">
          PADYAK<span class="text-white"> ADMIN</span>
        </h1>
      </div>

      <nav class="flex-1 px-4 py-6 space-y-2 overflow-y-auto">
        <a 
          href="#" 
          @click.prevent="currentView = 'dashboard'"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors', currentView === 'dashboard' ? 'bg-orange-500 text-white' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
          Dashboard
        </a>
        
        <a 
          href="#" 
          @click.prevent="currentView = 'users'"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors', currentView === 'users' ? 'bg-orange-500 text-white' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          Users
        </a>

        <a 
          href="#" 
          @click.prevent="currentView = 'products'"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors', currentView === 'products' ? 'bg-orange-500 text-white' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
          </svg>
          Products
        </a>

        <a 
          href="#" 
          @click.prevent="currentView = 'settings'"
          :class="['flex items-center px-4 py-3 rounded-lg transition-colors', currentView === 'settings' ? 'bg-orange-500 text-white' : 'text-gray-400 hover:bg-gray-700 hover:text-white']"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          Settings
        </a>
      </nav>

      <div class="p-4 border-t border-gray-700">
        <button @click="handleLogout" class="flex items-center w-full px-4 py-2 text-gray-400 hover:text-white hover:bg-gray-700 rounded-lg transition-colors">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          Logout
        </button>
      </div>
    </aside>

    <div class="flex-1 flex flex-col h-screen overflow-hidden">
      
      <header class="h-16 bg-gray-800 border-b border-gray-700 flex items-center justify-between px-6 z-10">
        <div class="flex items-center md:hidden">
          <button class="text-gray-400 hover:text-white">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
        
        <h2 class="text-xl font-semibold text-white capitalize">{{ currentView }}</h2>

        <div class="flex items-center space-x-4">
          <div class="relative">
            <span class="absolute top-0 right-0 h-2 w-2 bg-red-500 rounded-full"></span>
            <button class="text-gray-400 hover:text-white">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
            </button>
          </div>
          
          <div class="flex items-center space-x-2">
            <img src="https://ui-avatars.com/api/?name=Admin+User&background=f97316&color=fff" alt="Admin" class="h-8 w-8 rounded-full border border-gray-600">
            <span class="text-sm font-medium text-gray-300 hidden sm:block">Admin User</span>
          </div>
        </div>
      </header>

      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-900 p-6">
        
        <div v-if="currentView === 'dashboard'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-gray-400 text-sm font-medium">Total Users</h3>
                <span class="p-2 bg-orange-500/10 rounded-lg text-orange-500">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0z" /></svg>
                </span>
              </div>
              <p class="text-3xl font-bold text-white">1,245</p>
              <p class="text-green-400 text-sm mt-1 flex items-center">
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path></svg>
                +12% this month
              </p>
            </div>

            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-gray-400 text-sm font-medium">Total Listings</h3>
                <span class="p-2 bg-blue-500/10 rounded-lg text-blue-500">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" /></svg>
                </span>
              </div>
              <p class="text-3xl font-bold text-white">856</p>
              <p class="text-green-400 text-sm mt-1 flex items-center">
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path></svg>
                +5% this week
              </p>
            </div>

            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-gray-400 text-sm font-medium">Revenue</h3>
                <span class="p-2 bg-green-500/10 rounded-lg text-green-500">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                </span>
              </div>
              <p class="text-3xl font-bold text-white">$45,230</p>
              <p class="text-red-400 text-sm mt-1 flex items-center">
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6"></path></svg>
                -2% this month
              </p>
            </div>

            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-gray-400 text-sm font-medium">Pending Reports</h3>
                <span class="p-2 bg-red-500/10 rounded-lg text-red-500">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
                </span>
              </div>
              <p class="text-3xl font-bold text-white">12</p>
              <p class="text-gray-400 text-sm mt-1">Requires attention</p>
            </div>
          </div>

          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg p-6">
            <h3 class="text-lg font-bold text-white mb-4">Recent User Registrations</h3>
            <div class="overflow-x-auto">
              <table class="w-full text-left text-sm text-gray-400">
                <thead class="bg-gray-700/50 text-gray-200 uppercase">
                  <tr>
                    <th class="px-4 py-3 rounded-l-lg">User</th>
                    <th class="px-4 py-3">Email</th>
                    <th class="px-4 py-3">Date</th>
                    <th class="px-4 py-3 rounded-r-lg">Status</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-700">
                  <tr v-for="i in 5" :key="i" class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-4 py-3 font-medium text-white">John Doe {{ i }}</td>
                    <td class="px-4 py-3">john{{ i }}@example.com</td>
                    <td class="px-4 py-3">June {{ 20 + i }}, 2026</td>
                    <td class="px-4 py-3">
                      <span class="px-2 py-1 rounded-full text-xs font-semibold bg-green-500/10 text-green-500 border border-green-500/20">Active</span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <div v-if="currentView === 'users'" class="space-y-6">
          <div class="flex justify-between items-center">
            <h2 class="text-2xl font-bold text-white">User Management</h2>
            <div class="flex space-x-2">
              <input type="text" placeholder="Search users..." class="bg-gray-800 border border-gray-700 text-white px-4 py-2 rounded-lg focus:outline-none focus:border-orange-500">
              <button class="bg-orange-500 hover:bg-orange-600 text-white px-4 py-2 rounded-lg font-medium transition-colors">Add User</button>
            </div>
          </div>

          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
            <table class="w-full text-left text-sm text-gray-400">
              <thead class="bg-gray-700/50 text-gray-200 uppercase">
                <tr>
                  <th class="px-6 py-4">ID</th>
                  <th class="px-6 py-4">Name</th>
                  <th class="px-6 py-4">Username</th>
                  <th class="px-6 py-4">Role</th>
                  <th class="px-6 py-4">Status</th>
                  <th class="px-6 py-4 text-right">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-700">
                <tr v-for="user in users" :key="user.id" class="hover:bg-gray-700/30 transition-colors">
                  <td class="px-6 py-4">#{{ user.id }}</td>
                  <td class="px-6 py-4">
                    <div class="flex items-center">
                      <div class="h-8 w-8 rounded-full bg-gray-600 flex items-center justify-center text-xs font-bold text-white mr-3">
                        {{ user.first_name.charAt(0) }}{{ user.last_name.charAt(0) }}
                      </div>
                      <div>
                        <div class="font-medium text-white">{{ user.first_name }} {{ user.last_name }}</div>
                        <div class="text-xs text-gray-500">{{ user.email }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4">@{{ user.username }}</td>
                  <td class="px-6 py-4">
                    <span :class="user.role === 'admin' ? 'text-purple-400' : 'text-blue-400'">{{ user.role }}</span>
                  </td>
                  <td class="px-6 py-4">
                    <span :class="[
                      'px-2 py-1 rounded-full text-xs font-semibold border',
                      user.status === 'active' ? 'bg-green-500/10 text-green-500 border-green-500/20' : 
                      user.status === 'suspended' ? 'bg-yellow-500/10 text-yellow-500 border-yellow-500/20' : 
                      'bg-red-500/10 text-red-500 border-red-500/20'
                    ]">
                      {{ user.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-right">
                    <button class="text-gray-400 hover:text-white mx-1">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" /></svg>
                    </button>
                    <button class="text-gray-400 hover:text-red-500 mx-1">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div v-if="currentView === 'products'" class="flex flex-col items-center justify-center h-96 text-gray-500">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-4 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <h3 class="text-xl font-semibold">Product Management Module</h3>
          <p>This section is under construction.</p>
        </div>

        <div v-if="currentView === 'settings'" class="flex flex-col items-center justify-center h-96 text-gray-500">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-4 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <h3 class="text-xl font-semibold">System Settings</h3>
          <p>Configure global application settings here.</p>
        </div>

      </main>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'AdminDashboard',
  
  data() {
    return {
      currentView: 'dashboard', // dashboard, users, products, settings
      users: []
    }
  },

  async mounted() {
    await this.loadUsers()
  },

  methods: {
    async loadUsers() {
      try {
        const response = await axios.get('http://localhost:3000/api/users')
        this.users = response.data || []
      } catch (err) {
        console.error('Failed to load users for admin:', err)
        this.users = [
          { id: 1, first_name: 'Admin', last_name: 'User', username: 'admin', email: 'admin@padyak.com', role: 'admin', status: 'active' },
          { id: 2, first_name: 'John', last_name: 'Doe', username: 'johndoe', email: 'john@example.com', role: 'user', status: 'active' },
          { id: 3, first_name: 'Jane', last_name: 'Smith', username: 'janesmith', email: 'jane@example.com', role: 'user', status: 'suspended' }
        ]
      }
    },

    handleLogout() {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      
      alert('Logged out successfully')
      
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
</style>