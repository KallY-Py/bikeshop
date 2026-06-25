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
          User Management
        </h2>

        <!-- Right Side Icons -->
        <div class="flex items-center space-x-4 ml-auto">
          <!-- Notification Bell -->
          <div class="relative">
            <span class="absolute top-0 right-0 h-2.5 w-2.5 bg-red-500 rounded-full border-2 border-gray-800"></span>
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
              <a href="#" class="block px-4 py-2 text-sm text-gray-300 hover:bg-gray-700 hover:text-white transition-colors">
                Settings
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
        
        <!-- User Management Content -->
        <div class="space-y-6 animate-fade-in">
          
          <!-- Header with Search and Filters -->
          <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
            <h2 class="text-2xl font-bold text-white">User Management</h2>
            
            <div class="flex flex-col sm:flex-row gap-3 w-full md:w-auto">
              <!-- Search Bar -->
              <div class="relative">
                <input 
                  v-model="searchQuery"
                  type="text" 
                  placeholder="Search users..." 
                  class="bg-gray-800 border border-gray-700 text-white px-4 py-2 pl-10 rounded-lg focus:outline-none focus:border-orange-500 w-full sm:w-64"
                />
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400 absolute left-3 top-2.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>

              <!-- Filter Dropdown -->
              <select 
                v-model="sortBy"
                class="bg-gray-800 border border-gray-700 text-white px-4 py-2 rounded-lg focus:outline-none focus:border-orange-500"
              >
                <option value="newest">Newest to Oldest</option>
                <option value="oldest">Oldest to Newest</option>
                <option value="alphabetical">Alphabetical (A-Z)</option>
                <option value="alphabetical-desc">Alphabetical (Z-A)</option>
              </select>
            </div>
          </div>

          <!-- Users Table -->
          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
            <div class="overflow-x-auto">
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
                  <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-gray-700/30 transition-colors">
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
                      <select 
                        v-model="user.status"
                        @change="updateUserStatus(user)"
                        :class="[
                          'px-2 py-1 rounded-lg text-xs font-semibold border cursor-pointer',
                          user.status === 'active' ? 'bg-green-500/10 text-green-500 border-green-500/20' : 
                          user.status === 'suspended' ? 'bg-yellow-500/10 text-yellow-500 border-yellow-500/20' : 
                          'bg-red-500/10 text-red-500 border-red-500/20'
                        ]"
                      >
                        <option value="active">Active</option>
                        <option value="suspended">Suspended</option>
                        <option value="banned">Banned</option>
                      </select>
                    </td>
                    <td class="px-6 py-4 text-right">
                      <button @click="openEditModal(user)" class="text-gray-400 hover:text-orange-500 mx-1 transition-colors" title="Edit User">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                        </svg>
                      </button>
                      <button @click="openDeleteModal(user)" class="text-gray-400 hover:text-red-500 mx-1 transition-colors" title="Delete User">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <!-- Empty State -->
            <div v-if="filteredUsers.length === 0" class="p-12 text-center text-gray-500">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-4 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              <p class="text-lg font-medium">No users found</p>
              <p class="text-sm mt-1">Try adjusting your search or filters</p>
            </div>
          </div>
        </div>

      </main>
    </div>

    <!-- Edit User Modal -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="bg-gray-800 rounded-2xl shadow-2xl w-full max-w-md border border-gray-700 animate-fade-in">
        <div class="flex items-center justify-between p-6 border-b border-gray-700">
          <h3 class="text-xl font-bold text-white">Edit User Profile</h3>
          <button @click="closeEditModal" class="text-gray-400 hover:text-white transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <form @submit.prevent="saveUser" class="p-6 space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">First Name</label>
              <input v-model="editingUser.first_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Last Name</label>
              <input v-model="editingUser.last_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-300 mb-1">Username</label>
            <input v-model="editingUser.username" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-300 mb-1">Email</label>
            <input v-model="editingUser.email" type="email" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-300 mb-1">Phone Number</label>
            <input v-model="editingUser.phone_number" type="tel" class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-300 mb-1">Role</label>
            <select v-model="editingUser.role" class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors">
              <option value="user">User</option>
              <option value="admin">Admin</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-300 mb-1">Status</label>
            <select v-model="editingUser.status" class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors">
              <option value="active">Active</option>
              <option value="suspended">Suspended</option>
              <option value="banned">Banned</option>
            </select>
          </div>

          <div class="flex space-x-3 pt-4">
            <button type="button" @click="closeEditModal" class="flex-1 px-4 py-2 border border-gray-600 text-gray-300 rounded-lg hover:bg-gray-700 transition-colors">
              Cancel
            </button>
            <button type="submit" class="flex-1 px-4 py-2 bg-orange-500 hover:bg-orange-600 text-white rounded-lg font-semibold transition-colors">
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center p-4">
      <div class="bg-gray-800 rounded-2xl shadow-2xl w-full max-w-md border border-gray-700 animate-fade-in">
        <div class="p-6 text-center">
          <div class="w-16 h-16 bg-red-500/20 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          
          <h3 class="text-xl font-bold text-white mb-2">Delete User Account</h3>
          <p class="text-gray-400 mb-6">
            Are you sure you want to delete <span class="text-white font-semibold">{{ deletingUser?.first_name }} {{ deletingUser?.last_name }}</span>? 
            This action cannot be undone and all associated data will be permanently removed.
          </p>

          <div class="flex space-x-3">
            <button @click="closeDeleteModal" class="flex-1 px-4 py-2 border border-gray-600 text-gray-300 rounded-lg hover:bg-gray-700 transition-colors">
              Cancel
            </button>
            <button @click="confirmDelete" class="flex-1 px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg font-semibold transition-colors">
              Delete Account
            </button>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios'
import Admin_nav from '@/components/Admin_nav.vue'

export default {
  name: 'AdminManagement',
  components: {
    Admin_nav
  },
  
  data() {
    return {
      currentView: 'management',
      isSidebarCollapsed: false,
      isProfileDropdownOpen: false,
      users: [],
      searchQuery: '',
      sortBy: 'newest',
      showEditModal: false,
      showDeleteModal: false,
      editingUser: {},
      deletingUser: null
    }
  },

  computed: {
    filteredUsers() {
      let filtered = [...this.users]

      // Search filter
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase()
        filtered = filtered.filter(user => 
          user.first_name.toLowerCase().includes(query) ||
          user.last_name.toLowerCase().includes(query) ||
          user.username.toLowerCase().includes(query) ||
          user.email.toLowerCase().includes(query)
        )
      }

      // Sorting
      switch (this.sortBy) {
        case 'newest':
          filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
          break
        case 'oldest':
          filtered.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
          break
        case 'alphabetical':
          filtered.sort((a, b) => a.first_name.localeCompare(b.first_name))
          break
        case 'alphabetical-desc':
          filtered.sort((a, b) => b.first_name.localeCompare(a.first_name))
          break
      }

      return filtered
    }
  },

  async mounted() {
    await this.loadUsers()
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
      // Handled by Admin_nav component
    },

    handleClickOutside(event) {
      const dropdown = this.$el.querySelector('.relative')
      if (dropdown && !dropdown.contains(event.target)) {
        this.isProfileDropdownOpen = false
      }
    },

    async loadUsers() {
      try {
        const response = await axios.get('http://localhost:3000/api/users')
        this.users = response.data || []
      } catch (err) {
        console.error('Failed to load users:', err)
        // Fallback dummy data
        this.users = [
          { id: 1, first_name: 'Admin', last_name: 'User', username: 'admin', email: 'admin@padyak.com', phone_number: '+63 123 456 7890', role: 'admin', status: 'active', created_at: '2026-01-01T00:00:00Z' },
          { id: 2, first_name: 'John', last_name: 'Doe', username: 'johndoe', email: 'john@example.com', phone_number: '+63 987 654 3210', role: 'user', status: 'active', created_at: '2026-01-15T00:00:00Z' },
          { id: 3, first_name: 'Jane', last_name: 'Smith', username: 'janesmith', email: 'jane@example.com', phone_number: '', role: 'user', status: 'suspended', created_at: '2026-02-01T00:00:00Z' },
          { id: 4, first_name: 'Mike', last_name: 'Johnson', username: 'mikej', email: 'mike@example.com', phone_number: '+63 555 123 4567', role: 'user', status: 'active', created_at: '2026-02-10T00:00:00Z' },
          { id: 5, first_name: 'Sarah', last_name: 'Williams', username: 'sarahw', email: 'sarah@example.com', phone_number: '', role: 'user', status: 'banned', created_at: '2026-03-01T00:00:00Z' }
        ]
      }
    },

    openEditModal(user) {
      this.editingUser = { ...user }
      this.showEditModal = true
    },

    closeEditModal() {
      this.showEditModal = false
      this.editingUser = {}
    },

    async saveUser() {
        try {
            const token = localStorage.getItem('token')
            await axios.put(
            `http://localhost:3000/api/admin/users/${this.editingUser.id}`, 
            this.editingUser,
            { headers: { Authorization: `Bearer ${token}` } }
            )
            
            const index = this.users.findIndex(u => u.id === this.editingUser.id)
            if (index !== -1) {
            this.users[index] = { ...this.editingUser }
            }
            
            alert('User updated successfully!')
            this.closeEditModal()
        } catch (err) {
            console.error('Failed to update user:', err)
            alert(err.response?.data?.error || 'Failed to update user. Please try again.')
        }
    },

    async updateUserStatus(user) {
        try {
            const token = localStorage.getItem('token')
            await axios.patch(
            `http://localhost:3000/api/admin/users/${user.id}/status`, 
            { status: user.status },
            { headers: { Authorization: `Bearer ${token}` } }
            )
            
            const index = this.users.findIndex(u => u.id === user.id)
            if (index !== -1) {
            this.users[index].status = user.status
            }
        } catch (err) {
            console.error('Failed to update user status:', err)
            alert(err.response?.data?.error || 'Failed to update status.')
            await this.loadUsers()
        }
    },

    openDeleteModal(user) {
      this.deletingUser = user
      this.showDeleteModal = true
    },

    closeDeleteModal() {
      this.showDeleteModal = false
      this.deletingUser = null
    },

    async confirmDelete() {
        if (!this.deletingUser) return

        try {
            const token = localStorage.getItem('token')
            await axios.delete(
            `http://localhost:3000/api/admin/users/${this.deletingUser.id}`,
            { headers: { Authorization: `Bearer ${token}` } }
            )
            
            this.users = this.users.filter(u => u.id !== this.deletingUser.id)
            alert('User account deleted successfully!')
            this.closeDeleteModal()
        } catch (err) {
            console.error('Failed to delete user:', err)
            alert(err.response?.data?.error || 'Failed to delete user. Please try again.')
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

.animate-fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>