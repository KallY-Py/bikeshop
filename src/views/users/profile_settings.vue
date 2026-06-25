<template>
  <UserNav>
    <div class="max-w-5xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Profile Settings</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-1">Manage your account settings and preferences.</p>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
          <p class="text-on-surface-variant">Loading settings...</p>
        </div>
      </div>

      <!-- Content -->
      <div v-else class="grid grid-cols-1 lg:grid-cols-12 gap-gutter">
        <!-- Settings Tabs -->
        <aside class="lg:col-span-3">
          <div class="sticky top-[100px]">
            <h2 class="font-headline-md text-headline-md text-on-surface mb-6">Settings</h2>
            <nav class="flex flex-row lg:flex-col gap-2 overflow-x-auto lg:overflow-visible pb-4 lg:pb-0">
              <button v-for="tab in settingsTabs" :key="tab.id" @click="activeTab = tab.id"
                class="flex-shrink-0 text-left px-4 py-3 rounded-lg w-full flex items-center justify-between group transition-all"
                :class="activeTab === tab.id ? 'bg-surface-container border border-surface-container-high text-primary' : 'text-on-surface-variant hover:bg-surface-container hover:text-on-surface'">
                <span class="font-label-md text-label-md">{{ tab.label }}</span>
                <span class="material-symbols-outlined" :class="activeTab === tab.id ? 'text-primary' : 'opacity-50'">{{ tab.icon }}</span>
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
                  <img v-if="profile.avatar" :src="profile.avatar" alt="Avatar" class="w-full h-full object-cover" />
                  <div v-else class="w-full h-full bg-primary/20 flex items-center justify-center text-primary text-4xl font-bold">{{ userInitial }}</div>
                  <button @click="triggerAvatarUpload" class="absolute inset-0 bg-black/60 flex flex-col items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer">
                    <span class="material-symbols-outlined text-white mb-1">photo_camera</span>
                    <span class="font-label-sm text-white">Change</span>
                  </button>
                </div>
                <input ref="avatarInput" type="file" accept="image/jpeg,image/png" class="hidden" @change="handleAvatarChange" />
              </div>
              <div class="relative z-10 flex-1">
                <h3 class="font-headline-lg text-on-surface mb-2">{{ profile.first_name }} {{ profile.last_name }}</h3>
                <p class="font-body-md text-on-surface-variant mb-4">{{ profile.bio || 'Add a bio to tell people about yourself' }}</p>
                <div class="flex flex-wrap gap-3">
                  <button @click="triggerAvatarUpload" class="bg-surface-container border border-outline-variant hover:border-primary text-on-surface px-4 py-2 rounded-lg font-label-md flex items-center gap-2">
                    <span class="material-symbols-outlined text-lg">upload</span> Upload New
                  </button>
                  <button @click="removeAvatar" class="text-error px-4 py-2 rounded-lg font-label-md hover:bg-error/10">Remove</button>
                </div>
              </div>
            </section>

            <!-- Personal Form -->
            <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8 mt-8">
              <div class="mb-6 border-b border-surface-container-high pb-4 flex items-center justify-between">
                <h3 class="font-headline-md text-headline-md text-on-surface">Personal Information</h3>
                <span class="material-symbols-outlined text-primary">edit_document</span>
              </div>

              <div v-if="saveSuccess" class="mb-6 p-4 bg-secondary-container/10 border border-secondary-container/30 rounded-lg flex items-center gap-2">
                <span class="material-symbols-outlined text-secondary-container">check_circle</span>
                <span class="text-secondary-container font-label-md">Profile updated successfully!</span>
              </div>
              <div v-if="saveError" class="mb-6 p-4 bg-error/10 border border-error/30 rounded-lg flex items-center gap-2">
                <span class="material-symbols-outlined text-error">error</span>
                <span class="text-error font-label-md">{{ saveError }}</span>
              </div>

              <form @submit.prevent="saveProfile" class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="space-y-2">
                    <label class="block font-label-sm text-on-surface-variant uppercase">First Name</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all">
                      <input v-model="form.first_name" type="text" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0" />
                    </div>
                  </div>
                  <div class="space-y-2">
                    <label class="block font-label-sm text-on-surface-variant uppercase">Last Name</label>
                    <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container transition-all">
                      <input v-model="form.last_name" type="text" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0" />
                    </div>
                  </div>
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant uppercase">Username</label>
                  <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container flex items-center">
                    <span class="material-symbols-outlined text-on-surface-variant pl-4">alternate_email</span>
                    <input v-model="form.username" type="text" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0" />
                  </div>
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant uppercase">Email</label>
                  <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container flex items-center">
                    <span class="material-symbols-outlined text-on-surface-variant pl-4">mail</span>
                    <input v-model="form.email" type="email" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0" />
                  </div>
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant uppercase">Phone</label>
                  <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container flex items-center">
                    <span class="material-symbols-outlined text-on-surface-variant pl-4">call</span>
                    <input v-model="form.phone" type="tel" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0" />
                  </div>
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant uppercase">Bio</label>
                  <div class="form-input-focus rounded-lg border border-surface-container-high bg-surface-container">
                    <textarea v-model="form.bio" rows="4" maxlength="500" class="w-full bg-transparent border-none text-on-surface py-3 px-4 focus:ring-0 resize-none" placeholder="Tell other riders about yourself..."></textarea>
                  </div>
                  <p class="text-right text-on-surface-variant text-xs mt-1">{{ form.bio ? form.bio.length : 0 }} / 500</p>
                </div>
                <div class="pt-4 flex justify-end gap-4">
                  <button type="button" @click="resetForm" class="px-6 py-2 rounded-lg text-on-surface-variant hover:bg-surface-container-high font-label-md">Cancel</button>
                  <button type="submit" :disabled="saving" class="bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black px-8 py-2 rounded-lg font-label-md hover:brightness-110 flex items-center gap-2 disabled:opacity-50">
                    <span v-if="saving" class="animate-spin rounded-full h-4 w-4 border-b-2 border-black"></span>
                    <span v-else class="material-symbols-outlined text-lg">save</span>
                    {{ saving ? 'Saving...' : 'Save Changes' }}
                  </button>
                </div>
              </form>
            </section>
          </div>

          <!-- Security Tab -->
          <div v-if="activeTab === 'security'">
            <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
              <h3 class="font-headline-md text-on-surface mb-6">Change Password</h3>
              <form @submit.prevent="changePassword" class="space-y-4 max-w-md">
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant">Current Password</label>
                  <input v-model="passwordForm.current" type="password" class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary" />
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant">New Password</label>
                  <input v-model="passwordForm.newPassword" type="password" class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary" />
                </div>
                <div class="space-y-2">
                  <label class="block font-label-sm text-on-surface-variant">Confirm Password</label>
                  <input v-model="passwordForm.confirm" type="password" class="w-full bg-surface border border-outline-variant rounded-lg py-3 px-4 text-on-surface focus:border-primary" />
                </div>
                <button type="submit" class="bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black px-6 py-2 rounded-lg font-label-md">Update Password</button>
              </form>
            </section>
          </div>

          <!-- Notifications Tab -->
          <div v-if="activeTab === 'notifications'">
            <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
              <h3 class="font-headline-md text-on-surface mb-6">Notification Preferences</h3>
              <div class="space-y-4">
                <div v-for="notif in notificationSettings" :key="notif.id" class="flex items-center justify-between py-3 border-b border-surface-container-high">
                  <div>
                    <p class="font-label-md text-on-surface">{{ notif.label }}</p>
                    <p class="text-sm text-on-surface-variant">{{ notif.description }}</p>
                  </div>
                  <button @click="notif.enabled = !notif.enabled" class="relative w-12 h-6 rounded-full transition-colors" :class="notif.enabled ? 'bg-primary' : 'bg-surface-variant'">
                    <span class="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full transition-transform" :class="notif.enabled ? 'translate-x-6' : ''"></span>
                  </button>
                </div>
              </div>
            </section>
          </div>

          <!-- Payment Tab -->
          <div v-if="activeTab === 'payment'">
            <section class="bg-surface-container-low border border-surface-container-high rounded-xl p-6 md:p-8">
              <h3 class="font-headline-md text-on-surface mb-6">Payment Methods</h3>
              <p class="text-on-surface-variant">Payment method management coming soon.</p>
            </section>
          </div>
        </div>
      </div>
    </div>
  </UserNav>
</template>

<script>
import { userService } from '@/services/userService'
import UserNav from '@/components/User_Nav.vue'

export default {
  name: 'ProfileSettingsView',
  components: { UserNav },
  data() {
    return {
      loading: true, saving: false, saveSuccess: false, saveError: null,
      activeTab: 'personal', showLocationInput: false,
      profile: { first_name: '', last_name: '', username: '', email: '', email_verified: false, phone: '', bio: '', location: '', avatar: null, status: 'Rider' },
      form: { first_name: '', last_name: '', username: '', email: '', phone: '', bio: '', location: '' },
      passwordForm: { current: '', newPassword: '', confirm: '' },
      preferences: { distance: 'km', weight: 'lbs' },
      notificationSettings: [
        { id: 1, label: 'New Messages', description: 'Get notified when you receive a new message', enabled: true },
        { id: 2, label: 'Listing Updates', description: 'Status changes on your listings', enabled: true }
      ],
      settingsTabs: [
        { id: 'personal', label: 'Personal Info', icon: 'person' },
        { id: 'security', label: 'Security', icon: 'shield_lock' },
        { id: 'notifications', label: 'Notifications', icon: 'notifications_active' },
        { id: 'payment', label: 'Payment Methods', icon: 'credit_card' }
      ]
    }
  },
  computed: {
    userInitial() { return (this.profile.first_name || 'U').charAt(0).toUpperCase() }
  },
  mounted() { this.loadProfile() },
  methods: {
    async loadProfile() {
      this.loading = true
      try {
        const userId = localStorage.getItem('userId') || '1'
        try {
          const res = await userService.getById(userId)
          const d = res.data
          this.profile = { first_name: d.first_name || '', last_name: d.last_name || '', username: d.username || '', email: d.email || '', phone: d.phone_number || '', bio: d.bio || '', location: d.location || '', avatar: d.profile_image || null, status: 'Rider' }
        } catch {
          this.profile = { first_name: 'Mark', last_name: 'Capillan', username: 'mark', email: 'ksapareto@gmail.com', phone: '', bio: '', location: '', avatar: null, status: 'Rider' }
        }
        this.form = { ...this.profile }
      } catch (err) { console.error(err) }
      finally { this.loading = false }
    },
    async saveProfile() {
      this.saving = true; this.saveSuccess = false; this.saveError = null
      try {
        const userId = localStorage.getItem('userId') || '1'
        await userService.updateProfile(userId, this.form)
        this.profile = { ...this.profile, ...this.form }
        this.saveSuccess = true
        setTimeout(() => { this.saveSuccess = false }, 3000)
      } catch (err) {
        this.saveError = err.response?.data?.error || 'Failed to save.'
      } finally { this.saving = false }
    },
    resetForm() { this.form = { ...this.profile } },
    triggerAvatarUpload() { this.$refs.avatarInput.click() },
    handleAvatarChange(e) {
      const file = e.target.files[0]
      if (file) { const r = new FileReader(); r.onload = (ev) => { this.profile.avatar = ev.target.result }; r.readAsDataURL(file) }
    },
    removeAvatar() { this.profile.avatar = null },
    changePassword() {
      if (this.passwordForm.newPassword !== this.passwordForm.confirm) { alert('Passwords do not match'); return }
      this.passwordForm = { current: '', newPassword: '', confirm: '' }
      alert('Password updated!')
    }
  }
}
</script>

<style scoped>
.form-input-focus:focus-within { box-shadow: inset 0 0 4px 1px rgba(255,107,0,0.4); border-color: #ff6b00; }
</style>