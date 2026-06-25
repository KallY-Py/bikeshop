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
          {{ currentView.replace('-', ' ') }}
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
        
        <!-- VIEW: DASHBOARD -->
        <div v-if="currentView === 'dashboard'" class="space-y-6 animate-fade-in">
          
          <!-- Stats Cards Row -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <!-- Total Users -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">TOTAL USERS</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-orange-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">124,592</p>
                <span class="text-green-400 text-sm font-medium">+12.4%</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-orange-500 w-3/4"></div>
              </div>
            </div>

            <!-- Active Listings -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">ACTIVE LISTINGS</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">8,430</p>
                <span class="text-green-400 text-sm font-medium">+15.2%</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-blue-500 w-1/2"></div>
              </div>
            </div>

            <!-- 30D Volume -->
            <div class="bg-gray-800 rounded-xl p-6 border border-gray-700 shadow-lg">
              <div class="flex items-center justify-between mb-2">
                <h3 class="text-gray-400 text-sm font-medium">30D VOLUME</h3>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <div class="flex items-end justify-between">
                <p class="text-4xl font-bold text-white">$2.4M</p>
                <span class="text-red-400 text-sm font-medium">-1.8%</span>
              </div>
              <div class="mt-2 h-1 bg-gray-700 rounded-full overflow-hidden">
                <div class="h-full bg-green-500 w-2/3"></div>
              </div>
            </div>
          </div>

          <!-- Chart and Moderation Queue -->
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Revenue Growth Chart -->
            <div class="lg:col-span-2 bg-gray-800 rounded-xl border border-gray-700 shadow-lg p-6">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-lg font-semibold text-white">Revenue Growth</h3>
                <div class="flex space-x-2">
                  <button class="px-3 py-1 text-xs rounded bg-gray-700 text-gray-300 hover:bg-gray-600">1W</button>
                  <button class="px-3 py-1 text-xs rounded bg-orange-500 text-white">3M</button>
                  <button class="px-3 py-1 text-xs rounded bg-gray-700 text-gray-300 hover:bg-gray-600">1Y</button>
                </div>
              </div>
              <div class="h-64 flex items-center justify-center bg-gray-900/50 rounded-lg">
                <p class="text-gray-500">Chart Data Visualization</p>
              </div>
            </div>

            <!-- Moderation Queue -->
            <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg p-6">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-lg font-semibold text-white flex items-center">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 text-orange-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3" />
                  </svg>
                  Moderation Queue
                </h3>
                <span class="px-2 py-1 text-xs rounded bg-orange-500/20 text-orange-400 border border-orange-500/30">12 Pending</span>
              </div>
              
              <div class="space-y-4">
                <!-- Item 1 -->
                <div class="p-3 bg-gray-900/50 rounded-lg border border-gray-700">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="text-sm font-medium text-white">Suspicious Pricing</h4>
                    <span class="text-xs text-gray-500">2h ago</span>
                  </div>
                  <p class="text-xs text-gray-400 mb-3">"Trek Madone SLR 9 listed for $500. Seems too good to be true, likely a..."</p>
                  <div class="flex space-x-2">
                    <button class="px-3 py-1 text-xs rounded bg-gray-700 text-gray-300 hover:bg-gray-600">Review</button>
                    <button class="px-3 py-1 text-xs rounded bg-red-500/20 text-red-400 border border-red-500/30 hover:bg-red-500/30">Suspend</button>
                  </div>
                </div>

                <!-- Item 2 -->
                <div class="p-3 bg-gray-900/50 rounded-lg border border-gray-700">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="text-sm font-medium text-white">Inappropriate Content</h4>
                    <span class="text-xs text-gray-500">5h ago</span>
                  </div>
                  <p class="text-xs text-gray-400 mb-3">"Images included in the listing contain offensive material."</p>
                  <div class="flex space-x-2">
                    <button class="px-3 py-1 text-xs rounded bg-gray-700 text-gray-300 hover:bg-gray-600">Review</button>
                    <button class="px-3 py-1 text-xs rounded bg-red-500/20 text-red-400 border border-red-500/30 hover:bg-red-500/30">Suspend</button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent User Activity Table -->
          <div class="bg-gray-800 rounded-xl border border-gray-700 shadow-lg overflow-hidden">
            <div class="p-6 border-b border-gray-700 flex items-center justify-between">
              <h3 class="text-lg font-semibold text-white">Recent User Activity</h3>
              <div class="relative">
                <input 
                  type="text" 
                  placeholder="Search users..." 
                  class="bg-gray-900 border border-gray-700 text-white text-sm rounded-lg pl-10 pr-4 py-2 focus:outline-none focus:border-orange-500 w-64"
                />
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400 absolute left-3 top-2.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
            </div>
            
            <div class="overflow-x-auto">
              <table class="w-full text-left text-sm">
                <thead class="bg-gray-900/50 text-gray-400">
                  <tr>
                    <th class="px-6 py-4 font-medium">USER</th>
                    <th class="px-6 py-4 font-medium">STATUS</th>
                    <th class="px-6 py-4 font-medium">LISTINGS</th>
                    <th class="px-6 py-4 font-medium">LAST LOGIN</th>
                    <th class="px-6 py-4 font-medium text-right">ACTIONS</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-700">
                  <tr class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-6 py-4">
                      <div class="flex items-center">
                        <div class="h-10 w-10 rounded-full bg-orange-500 flex items-center justify-center text-white font-bold text-sm mr-3">
                          JD
                        </div>
                        <div>
                          <p class="font-medium text-white">John Doe</p>
                          <p class="text-xs text-gray-400">john.d@example.com</p>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <span class="px-2 py-1 rounded-full text-xs font-medium bg-green-500/10 text-green-400 border border-green-500/20">
                        ACTIVE
                      </span>
                    </td>
                    <td class="px-6 py-4 text-white">12</td>
                    <td class="px-6 py-4 text-gray-400">Just now</td>
                    <td class="px-6 py-4 text-right">
                      <button class="text-gray-400 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                      </button>
                    </td>
                  </tr>

                  <tr class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-6 py-4">
                      <div class="flex items-center">
                        <div class="h-10 w-10 rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-sm mr-3">
                          AS
                        </div>
                        <div>
                          <p class="font-medium text-white">Alice Smith</p>
                          <p class="text-xs text-gray-400">alice.s@example.com</p>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <span class="px-2 py-1 rounded-full text-xs font-medium bg-yellow-500/10 text-yellow-400 border border-yellow-500/20">
                        SUSPENDED
                      </span>
                    </td>
                    <td class="px-6 py-4 text-white">0</td>
                    <td class="px-6 py-4 text-gray-400">2 days ago</td>
                    <td class="px-6 py-4 text-right">
                      <button class="text-gray-400 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                      </button>
                    </td>
                  </tr>

                  <tr class="hover:bg-gray-700/30 transition-colors">
                    <td class="px-6 py-4">
                      <div class="flex items-center">
                        <div class="h-10 w-10 rounded-full bg-purple-500 flex items-center justify-center text-white font-bold text-sm mr-3">
                          MR
                        </div>
                        <div>
                          <p class="font-medium text-white">Mike Rider</p>
                          <p class="text-xs text-gray-400">mike.r@example.com</p>
                        </div>
                      </div>
                    </td>
                    <td class="px-6 py-4">
                      <span class="px-2 py-1 rounded-full text-xs font-medium bg-green-500/10 text-green-400 border border-green-500/20">
                        ACTIVE
                      </span>
                    </td>
                    <td class="px-6 py-4 text-white">5</td>
                    <td class="px-6 py-4 text-gray-400">1 hr ago</td>
                    <td class="px-6 py-4 text-right">
                      <button class="text-gray-400 hover:text-white">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                        </svg>
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <div class="p-4 border-t border-gray-700 text-center">
              <button class="text-orange-500 hover:text-orange-400 text-sm font-medium">View All Users</button>
            </div>
          </div>
        </div>

        <!-- Other views remain the same... -->
        <div v-if="currentView === 'users'" class="space-y-6 animate-fade-in">
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

        <!-- Placeholder views for other sections -->
        <div v-if="['reported', 'categories', 'settings'].includes(currentView)" class="flex flex-col items-center justify-center h-96 text-gray-500">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mb-4 opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
          </svg>
          <h3 class="text-xl font-semibold capitalize">{{ currentView.replace('-', ' ') }}</h3>
          <p>This section is under construction.</p>
        </div>

      </main>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Admin_nav from '@/components/Admin_nav.vue'

export default {
  name: 'AdminDashboard',
  components: {
    Admin_nav
  },
  
  data() {
    return {
      currentView: 'dashboard',
      isSidebarCollapsed: false,
      isProfileDropdownOpen: false,
      users: []
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
      // This will be handled by the Admin_nav component
      this.$emit('toggle-mobile-menu')
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

.animate-fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>