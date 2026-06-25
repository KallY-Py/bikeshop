<template>
  <div class="bg-background text-on-background font-body-md min-h-screen flex flex-col md:flex-row">
    <!-- TopNavBar (Mobile Only) -->
    <header class="md:hidden flex justify-between items-center px-margin-mobile h-16 w-full bg-background/80 backdrop-blur-md border-b border-outline-variant shadow-none sticky top-0 z-50">
      <div class="font-display text-headline-lg-mobile text-primary uppercase tracking-tighter">VeloHub</div>
      <div class="flex gap-4">
        <span class="material-symbols-outlined text-primary cursor-pointer">notifications</span>
        <router-link to="/dashboard">
          <span class="material-symbols-outlined text-primary cursor-pointer">account_circle</span>
        </router-link>
      </div>
    </header>

    <!-- SideNavBar (Desktop) -->
    <nav class="hidden md:flex flex-col h-screen py-8 fixed left-0 top-0 w-[280px] bg-surface-container-low border-r border-outline-variant shadow-sm z-40">
      <div class="px-6 mb-8 flex items-center gap-4">
        <div class="w-12 h-12 rounded-full overflow-hidden bg-surface border border-outline-variant">
          <img :src="profile.avatar || 'https://via.placeholder.com/48'" alt="User profile" class="w-full h-full object-cover" />
        </div>
        <div>
          <div class="font-label-md text-label-md text-on-surface-variant">Welcome back</div>
          <div class="font-headline-md text-headline-md text-primary">{{ profile.status || 'Rider' }}</div>
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
    </nav>

    <!-- Main Content -->
    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <div class="max-w-5xl mx-auto">
        <!-- Header -->
        <div class="mb-8">
          <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Profile Settings</h1>
          <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your account settings and preferences.</p>
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="flex justify-center items-center h-64">
          <div class="text-center">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
            <p class="text-on-surface-variant">Loading settings...</p>
          </div>
        </div>

        <!-- Content -->
        <div v-else class="grid grid-cols-1 lg:grid-cols-12 gap-gutter">
          <!-- Settings Navigation Sidebar -->
          <aside class="lg:col-span-3">
            <div class="sticky top-[100px]">
              <h2 class="font-headline-md text-headline-md text-on-surface mb-6">Settings</h2>
              <nav class="flex flex-row lg:flex-col gap-2 overflow-x-auto lg:overflow-visible pb-4 lg:pb-0">
                <button 
                  v-for="tab in settingsTabs" 
                  :key="tab.id"
                  @click="activeTab = tab.id"
                  class="flex-shrink-0 text-left px-4 py-3 rounded-lg w-full flex items-center justify-between group transition-all"
                  :class="activeTab === tab.id 
                    ? 'bg-surface-container border border-surface-container-high text-primary' 
                    : 'text-on-surface-variant hover:bg-surface-container hover:text-on-surface'"
                >
                  <span class="font-label-md text-label-md">{{ tab.label }}</span>
                  <span 
                    class="material-symbols-outlined"
                    :class="activeTab === tab.id ? 'text-primary' : 'opacity-50 group-hover:opacity-100'"
                  >{{ tab.icon }}</span>
                </button>
              </nav>
            </div>
          </aside>

          <!-- Settings Content -->
          <div class="lg:col-span-9 space-y-8">
            <!-- Personal Info Tab -->
            <div v-if="activeTab === 'personal'">
              <!-- Avatar Section -->
              <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8 flex flex-col md:flex-row items-start md:items-center gap-8 relative overflow-hidden">
                <div class="absolute inset-0 opacity-10 pointer-events-none">
                  <div class="absolute -right-20 -top-20 w-64 h-64 bg-primary rounded-full blur-[100px]"></div>
                </div>
                
                <div class="relative z-10 flex-shrink-0">
                  <div class="w-32 h-32 rounded-full overflow-hidden border-2 border-surface-container-high relative group">
                    <img 
                      :src="profile.avatar || 'https://via.placeholder.com/128'" 
                      alt="Profile Avatar" 
                      class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105"
                    />
                    <button 
                      @click="triggerAvatarUpload"
                      class="absolute inset-0 bg-black/60 flex flex-col items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer"
                    >
                      <span class="material-symbols-outlined text-white mb-1">photo_camera</span>
                      <span class="font-label-sm text-label-sm text-white">Change</span>
                    </button>
                  </div>
                  <input 
                    ref="avatarInput"
                    type="file" 
                    accept="image/*" 
                    class="hidden" 
                    @change="handleAvatarChange"
                  />
                </div>

                <div class="relative z-10 flex-1">
                  <h3 class="font-headline-lg text-on-surface mb-2">{{ profile.first_name }} {{ profile.last_name }}</h3>
                  <p class="font-body-md text-body-md text-on-surface-variant mb-4">{{ profile.bio || 'Add a bio to tell people about yourself' }}</p>
                  <div class="flex flex-wrap gap-3">
                    <button 
                      @click="triggerAvatarUpload"
                      class="bg-surface-container border border-outline-variant hover:border-primary text-on-surface px-4 py-2 rounded-lg font-label-md text-label-md transition-colors flex items-center gap-2"
                    >
                      <span class="material-symbols-outlined text-[18px]">upload</span>
                      Upload New
                    </button>
                    <button 
                      @click="removeAvatar"
                      class="text-error px-4 py-2 rounded-lg font-label-md text-label-md hover:bg-error/10 transition-colors"
                    >
                      Remove
                    </button>
                  </div>
                </div>
              </section>

              <!-- Personal Information Form -->
              <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
                <div class="mb-6 border-b border-surface-container-high pb-4 flex items-center justify-between">
                  <h3 class="font-headline-md text-headline-md text-on-surface">Personal Information</h3>
                  <span class="material-symbols-outlined text-primary">edit_document</span>
                </div>

                <!-- Success Message -->
                <div v-if="saveSuccess" class="mb-6 p-4 bg-secondary-container/10 border border-secondary-container/30 rounded-lg flex items-center gap-2">
                  <span class="material-symbols-outlined text-secondary-container">check_circle</span>
                  <span class="text-secondary-container font-label-md">Profile updated successfully!</span>
                </div>

                <!-- Error Message -->
                <div v-if="saveError" class="mb-6 p-4 bg-error/10 border border-error/30 rounded-lg flex items-center gap-2">
                  <span class="material-symbols-outlined text-error">error</span>
                  <span class="text-error font-label-md">{{ saveError }}</span>
                </div>

                <form @submit.prevent="saveProfile" class="space-y-6">
                  <!-- Name Fields -->
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div class="space-y-2">
                      <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">First Name</label>
                      <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all">
                        <input 
                          v-model="form.first_name"
                          class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md"
                          type="text"
                          placeholder="First name"
                        />
                      </div>
                    </div>
                    <div class="space-y-2">
                      <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">Last Name</label>
                      <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all">
                        <input 
                          v-model="form.last_name"
                          class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md"
                          type="text"
                          placeholder="Last name"
                        />
                      </div>
                    </div>
                  </div>

                  <!-- Username -->
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">Username</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all flex items-center">
                      <span class="material-symbols-outlined text-on-surface-variant pl-4">alternate_email</span>
                      <input 
                        v-model="form.username"
                        class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md"
                        type="text"
                        placeholder="Username"
                      />
                    </div>
                  </div>

                  <!-- Email -->
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">Email Address</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all flex items-center">
                      <span class="material-symbols-outlined text-on-surface-variant pl-4">mail</span>
                      <input 
                        v-model="form.email"
                        class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md"
                        type="email"
                        placeholder="Email address"
                      />
                      <span 
                        v-if="profile.email_verified"
                        class="bg-primary/10 text-primary font-label-sm text-label-sm px-2 py-1 rounded mr-3 uppercase"
                      >
                        Verified
                      </span>
                      <button 
                        v-else
                        @click="verifyEmail"
                        type="button"
                        class="text-primary font-label-sm text-label-sm px-3 py-1 rounded mr-3 hover:bg-primary/10 transition-colors"
                      >
                        Verify
                      </button>
                    </div>
                  </div>

                  <!-- Phone -->
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">Phone Number</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all flex items-center">
                      <span class="material-symbols-outlined text-on-surface-variant pl-4">call</span>
                      <input 
                        v-model="form.phone"
                        class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md"
                        type="tel"
                        placeholder="+1 (555) 000-0000"
                      />
                    </div>
                  </div>

                  <!-- Bio -->
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant uppercase tracking-wider">Bio / Rider Profile</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all">
                      <textarea 
                        v-model="form.bio"
                        class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 font-body-md resize-none"
                        rows="4"
                        placeholder="Tell other riders about yourself..."
                        maxlength="500"
                      ></textarea>
                    </div>
                    <p class="text-right text-on-surface-variant font-label-sm text-label-sm mt-1">
                      {{ form.bio ? form.bio.length : 0 }} / 500
                    </p>
                  </div>

                  <!-- Form Actions -->
                  <div class="pt-4 flex justify-end gap-4">
                    <button 
                      type="button"
                      @click="resetForm"
                      class="px-6 py-2 rounded-lg text-on-surface-variant hover:bg-surface-container-high font-label-md text-label-md transition-colors border border-transparent"
                    >
                      Cancel
                    </button>
                    <button 
                      type="submit"
                      :disabled="saving"
                      class="bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black px-8 py-2 rounded-lg font-label-md text-label-md hover:brightness-110 transition-all shadow-lg flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      <span v-if="saving" class="animate-spin rounded-full h-4 w-4 border-b-2 border-black"></span>
                      <span v-else class="material-symbols-outlined text-[18px]">save</span>
                      {{ saving ? 'Saving...' : 'Save Changes' }}
                    </button>
                  </div>
                </form>
              </section>

              <!-- Location & Preferences -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mt-8">
                <!-- Location -->
                <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 relative overflow-hidden group">
                  <div class="relative z-10 h-full flex flex-col">
                    <div class="mb-4 flex items-center gap-2">
                      <span class="material-symbols-outlined text-primary">location_on</span>
                      <h3 class="font-headline-md text-headline-md text-on-surface">Primary Location</h3>
                    </div>
                    <div class="mt-auto space-y-2 bg-background/80 backdrop-blur-sm p-4 rounded-lg border border-surface-container">
                      <p class="font-label-md text-label-md text-on-surface">{{ form.location || 'Not set' }}</p>
                      <p class="font-body-md text-body-md text-on-surface-variant text-sm">Used for local pickup estimates.</p>
                      <button 
                        @click="showLocationInput = !showLocationInput"
                        class="text-primary font-label-md text-label-md mt-2 hover:underline"
                      >
                        {{ showLocationInput ? 'Cancel' : 'Update Location' }}
                      </button>
                      <div v-if="showLocationInput" class="mt-3 flex gap-2">
                        <input 
                          v-model="form.location"
                          type="text"
                          placeholder="City, State"
                          class="flex-1 bg-surface border border-outline-variant rounded-lg py-2 px-3 text-sm text-on-surface focus:border-primary"
                        />
                        <button 
                          @click="showLocationInput = false"
                          class="bg-primary-container text-on-primary-container px-3 py-1 rounded-lg text-sm"
                        >
                          Save
                        </button>
                      </div>
                    </div>
                  </div>
                </section>

                <!-- Units Preferences -->
                <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 flex flex-col justify-between">
                  <div>
                    <div class="mb-6 flex items-center gap-2 border-b border-surface-container-high pb-4">
                      <span class="material-symbols-outlined text-on-surface">straighten</span>
                      <h3 class="font-headline-md text-headline-md text-on-surface">Measurement Units</h3>
                    </div>
                    <div class="space-y-4">
                      <div class="flex items-center justify-between">
                        <span class="font-body-md text-body-md text-on-surface">Distance</span>
                        <div class="bg-surface-container p-1 rounded-lg inline-flex">
                          <button 
                            @click="preferences.distance = 'km'"
                            class="px-3 py-1 rounded font-label-md text-label-md transition-colors"
                            :class="preferences.distance === 'km' ? 'bg-primary/20 text-primary' : 'text-on-surface-variant hover:text-on-surface'"
                          >km</button>
                          <button 
                            @click="preferences.distance = 'mi'"
                            class="px-3 py-1 rounded font-label-md text-label-md transition-colors"
                            :class="preferences.distance === 'mi' ? 'bg-primary/20 text-primary' : 'text-on-surface-variant hover:text-on-surface'"
                          >mi</button>
                        </div>
                      </div>
                      <div class="flex items-center justify-between">
                        <span class="font-body-md text-body-md text-on-surface">Weight</span>
                        <div class="bg-surface-container p-1 rounded-lg inline-flex">
                          <button 
                            @click="preferences.weight = 'kg'"
                            class="px-3 py-1 rounded font-label-md text-label-md transition-colors"
                            :class="preferences.weight === 'kg' ? 'bg-primary/20 text-primary' : 'text-on-surface-variant hover:text-on-surface'"
                          >kg</button>
                          <button 
                            @click="preferences.weight = 'lbs'"
                            class="px-3 py-1 rounded font-label-md text-label-md transition-colors"
                            :class="preferences.weight === 'lbs' ? 'bg-primary/20 text-primary' : 'text-on-surface-variant hover:text-on-surface'"
                          >lbs</button>
                        </div>
                      </div>
                    </div>
                  </div>
                </section>
              </div>
            </div>

            <!-- Security Tab -->
            <div v-if="activeTab === 'security'">
              <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
                <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Change Password</h3>
                <form @submit.prevent="changePassword" class="space-y-4 max-w-md">
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant">Current Password</label>
                    <input 
                      v-model="passwordForm.current"
                      type="password"
                      class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary"
                    />
                  </div>
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant">New Password</label>
                    <input 
                      v-model="passwordForm.newPassword"
                      type="password"
                      class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary"
                    />
                  </div>
                  <div class="space-y-2">
                    <label class="block font-label-sm text-label-sm text-on-surface-variant">Confirm New Password</label>
                    <input 
                      v-model="passwordForm.confirm"
                      type="password"
                      class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary"
                    />
                  </div>
                  <button 
                    type="submit"
                    class="bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black px-6 py-2 rounded-lg font-label-md"
                  >
                    Update Password
                  </button>
                </form>
              </section>
            </div>

            <!-- Notifications Tab -->
            <div v-if="activeTab === 'notifications'">
              <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
                <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Notification Preferences</h3>
                <div class="space-y-4">
                  <div v-for="notif in notificationSettings" :key="notif.id" 
                       class="flex items-center justify-between py-3 border-b border-surface-container-high">
                    <div>
                      <p class="font-label-md text-on-surface">{{ notif.label }}</p>
                      <p class="text-sm text-on-surface-variant">{{ notif.description }}</p>
                    </div>
                    <button 
                      @click="notif.enabled = !notif.enabled"
                      class="relative w-12 h-6 rounded-full transition-colors"
                      :class="notif.enabled ? 'bg-primary' : 'bg-surface-variant'"
                    >
                      <span class="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full transition-transform"
                            :class="notif.enabled ? 'translate-x-6' : ''"></span>
                    </button>
                  </div>
                </div>
              </section>
            </div>

            <!-- Payment Tab -->
            <div v-if="activeTab === 'payment'">
              <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
                <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Payment Methods</h3>
                <p class="text-on-surface-variant">Payment method management coming soon.</p>
              </section>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- BottomNavBar (Mobile Only) -->
    <nav class="fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-4 py-3 md:hidden bg-surface-container-lowest/95 backdrop-blur-lg border-t border-outline-variant shadow-[0_-4px_10px_rgba(0,0,0,0.3)] rounded-t-xl">
      <router-link 
        v-for="item in mobileNavItems" 
        :key="item.name"
        :to="item.path"
        class="flex flex-col items-center justify-center rounded-xl px-4 py-1 active:scale-90 transition-transform duration-150"
        :class="$route.path === item.path ? 'text-primary bg-primary/10' : 'text-on-surface-variant'"
      >
        <span class="material-symbols-outlined">{{ item.icon }}</span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </router-link>
    </nav>
  </div>
</template>

<script>
import { userService } from '@/services/userService'

export default {
  name: 'ProfileSettingsView',
  
  data() {
    return {
      loading: true,
      saving: false,
      saveSuccess: false,
      saveError: null,
      activeTab: 'personal',
      showLocationInput: false,
      
      profile: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
        email_verified: false,
        phone: '',
        bio: '',
        location: '',
        avatar: null,
        status: 'Rider Elite Status'
      },
      
      form: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
        phone: '',
        bio: '',
        location: ''
      },
      
      passwordForm: {
        current: '',
        newPassword: '',
        confirm: ''
      },
      
      preferences: {
        distance: 'km',
        weight: 'lbs'
      },
      
      notificationSettings: [
        { id: 1, label: 'New Messages', description: 'Get notified when you receive a new message', enabled: true },
        { id: 2, label: 'Listing Updates', description: 'Status changes on your listings', enabled: true },
        { id: 3, label: 'Price Alerts', description: 'When similar bikes are listed in your area', enabled: false },
        { id: 4, label: 'Marketing Emails', description: 'Promotions and marketplace updates', enabled: false }
      ],
      
      settingsTabs: [
        { id: 'personal', label: 'Personal Info', icon: 'person' },
        { id: 'security', label: 'Security', icon: 'shield_lock' },
        { id: 'notifications', label: 'Notifications', icon: 'notifications_active' },
        { id: 'payment', label: 'Payment Methods', icon: 'credit_card' }
      ],
      
      navItems: [
        { name: 'Overview', icon: 'dashboard', path: '/dashboard' },
        { name: 'My Listings', icon: 'directions_bike', path: '/dashboard/listings' },
        { name: 'Sales', icon: 'payments', path: '/dashboard/sales' },
        { name: 'Analytics', icon: 'leaderboard', path: '/dashboard/analytics' },
        { name: 'Messages', icon: 'chat', path: '/dashboard/messages' }
      ],
      
      mobileNavItems: [
        { name: 'Home', icon: 'home', path: '/dashboard' },
        { name: 'Listings', icon: 'storefront', path: '/dashboard/listings' },
        { name: 'Chat', icon: 'forum', path: '/dashboard/messages' },
        { name: 'Profile', icon: 'person', path: '/settings' }
      ]
    }
  },
  
  mounted() {
    this.loadProfile()
  },
  
  methods: {
    async loadProfile() {
      try {
        this.loading = true
        const userId = localStorage.getItem('userId') || '1'
        
        try {
          const response = await userService.getById(userId)
          const data = response.data
          
          this.profile = {
            first_name: data.first_name || '',
            last_name: data.last_name || '',
            username: data.username || '',
            email: data.email || '',
            email_verified: data.email_verified || false,
            phone: data.phone_number || '',
            bio: data.bio || '',
            location: data.location || 'Portland, Oregon',
            avatar: data.profile_image || null,
            status: data.role === 'admin' ? 'Admin' : 'Rider Elite Status'
          }
        } catch {
          // Demo data
          this.profile = {
            first_name: 'Alex',
            last_name: 'Mercer',
            username: 'alexmercer',
            email: 'alex.mercer@velohub.com',
            email_verified: true,
            phone: '+1 (555) 019-2834',
            bio: 'Specializing in high-end road and gravel bikes. 10+ years of mechanic experience.',
            location: 'Portland, Oregon',
            avatar: null,
            status: 'Rider Elite Status'
          }
        }
        
        // Populate form
        this.form = {
          first_name: this.profile.first_name,
          last_name: this.profile.last_name,
          username: this.profile.username,
          email: this.profile.email,
          phone: this.profile.phone,
          bio: this.profile.bio,
          location: this.profile.location
        }
        
      } catch (err) {
        console.error('Failed to load profile:', err)
      } finally {
        this.loading = false
      }
    },
    
  async saveProfile() {
    try {
      this.saving = true
      this.saveSuccess = false
      this.saveError = null
      
      const userId = localStorage.getItem('userId') || '1'
      
      console.log('Saving profile for user:', userId)
      console.log('Form data:', this.form)
      
      const response = await userService.updateProfile(userId, {
        first_name: this.form.first_name,
        last_name: this.form.last_name,
        username: this.form.username,
        email: this.form.email,
        phone_number: this.form.phone,
        bio: this.form.bio,
        location: this.form.location
      })
      
      console.log('Save response:', response.data)
      
      // Update the profile data to reflect changes
      this.profile.first_name = this.form.first_name
      this.profile.last_name = this.form.last_name
      this.profile.username = this.form.username
      this.profile.email = this.form.email
      this.profile.phone = this.form.phone
      this.profile.bio = this.form.bio
      this.profile.location = this.form.location
      
      this.saveSuccess = true
      setTimeout(() => { this.saveSuccess = false }, 3000)
      
    } catch (err) {
      console.error('Failed to save profile:', err)
      
      if (err.response) {
        // Server responded with error
        this.saveError = err.response.data?.error || 'Server error. Please try again.'
      } else if (err.request) {
        // No response from server
        this.saveError = 'Cannot connect to server. Make sure the Go backend is running on port 3000.'
      } else {
        this.saveError = 'Failed to save changes. Please try again.'
      }
    } finally {
      this.saving = false
    }
  },
    
    resetForm() {
      this.form = {
        first_name: this.profile.first_name,
        last_name: this.profile.last_name,
        username: this.profile.username,
        email: this.profile.email,
        phone: this.profile.phone,
        bio: this.profile.bio,
        location: this.profile.location
      }
    },
    
    triggerAvatarUpload() {
      this.$refs.avatarInput.click()
    },
    
    handleAvatarChange(event) {
      const file = event.target.files[0]
      if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
          this.profile.avatar = e.target.result
        }
        reader.readAsDataURL(file)
      }
    },
    
    removeAvatar() {
      this.profile.avatar = null
    },
    
    verifyEmail() {
      console.log('Verify email clicked')
    },
    
    async changePassword() {
      if (this.passwordForm.newPassword !== this.passwordForm.confirm) {
        alert('Passwords do not match')
        return
      }
      
      try {
        // await userService.changePassword(this.passwordForm)
        console.log('Password changed')
        this.passwordForm = { current: '', newPassword: '', confirm: '' }
        alert('Password updated successfully!')
      } catch (err) {
        console.error('Failed to change password:', err)
        alert('Failed to change password')
      }
    }
  }
}
</script>

<style scoped>
.form-input-focus:focus-within {
  box-shadow: inset 0 0 4px 1px rgba(255, 107, 0, 0.4);
  border-color: #ff6b00;
}
</style>