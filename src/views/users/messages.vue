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
          <span v-if="item.name === 'Messages' && unreadCount > 0" 
                class="ml-auto bg-primary text-background text-xs px-2 py-0.5 rounded-full">
            {{ unreadCount }}
          </span>
        </router-link>
      </div>
    </nav>

    <!-- Main Content Area (Messaging Interface) -->
    <main class="flex-1 md:ml-[280px] flex flex-col md:flex-row h-screen overflow-hidden">
      <!-- Conversations List Sidebar -->
      <div class="w-full md:w-[380px] border-r border-surface-container-high flex flex-col bg-surface-container-low/50"
           :class="{ 'hidden md:flex': activeConversation }">
        <!-- Inbox Header -->
        <div class="p-4 border-b border-surface-container-high bg-surface-container-low flex justify-between items-center">
          <h2 class="font-headline-md text-headline-md text-on-surface">Messages</h2>
          <div class="flex gap-2">
            <button 
              @click="showNewMessage = true"
              class="p-2 rounded-full bg-primary/10 text-primary hover:bg-primary/20 transition-colors"
              title="New Message"
            >
              <span class="material-symbols-outlined">edit_square</span>
            </button>
            <button 
              class="p-2 rounded-full bg-surface-container-high text-on-surface hover:text-primary transition-colors"
              title="Filter"
            >
              <span class="material-symbols-outlined">filter_list</span>
            </button>
          </div>
        </div>

        <!-- Search -->
        <div class="p-3 border-b border-surface-container-high">
          <div class="relative">
            <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant text-lg">search</span>
            <input 
              v-model="conversationSearch"
              type="text"
              placeholder="Search conversations..."
              class="w-full bg-surface border border-outline-variant rounded-lg py-2 pl-10 pr-4 text-sm text-on-surface focus:border-primary focus:ring-1 focus:ring-primary transition-colors"
            />
          </div>
        </div>

        <!-- Loading State -->
        <div v-if="loadingConversations" class="flex-1 flex items-center justify-center">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
        </div>

        <!-- Conversations List -->
        <div v-else class="overflow-y-auto flex-1 custom-scrollbar">
          <div v-if="filteredConversations.length === 0" class="p-8 text-center text-on-surface-variant">
            <span class="material-symbols-outlined text-4xl mb-2">forum</span>
            <p>No conversations yet</p>
          </div>

          <div 
            v-for="conversation in filteredConversations" 
            :key="conversation.id"
            @click="openConversation(conversation)"
            class="p-4 border-l-4 cursor-pointer flex gap-4 transition-colors hover:bg-surface-container-high"
            :class="activeConversation?.id === conversation.id 
              ? 'border-primary bg-primary/5' 
              : conversation.unread ? 'border-transparent' : 'border-transparent opacity-70'"
          >
            <!-- Avatar -->
            <div class="relative">
              <img 
                :src="conversation.listing_image || 'https://via.placeholder.com/56'"
                :alt="conversation.listing_title"
                class="w-14 h-14 rounded-lg object-cover border border-surface-container-high"
              />
              <div 
                v-if="conversation.unread"
                class="absolute -top-1 -right-1 w-3 h-3 bg-primary rounded-full border-2 border-surface-container-low"
              ></div>
            </div>

            <!-- Content -->
            <div class="flex-1 min-w-0">
              <div class="flex justify-between items-baseline mb-1">
                <h3 class="font-label-md text-label-md truncate"
                    :class="conversation.unread ? 'text-on-surface font-bold' : 'text-on-surface'">
                  {{ conversation.participant_name }}
                </h3>
                <span class="font-label-sm text-label-sm shrink-0 ml-2"
                      :class="conversation.unread ? 'text-primary' : 'text-on-surface-variant'">
                  {{ formatMessageTime(conversation.last_message_time) }}
                </span>
              </div>
              <p class="font-label-sm text-label-sm text-on-surface-variant truncate">
                {{ conversation.listing_title }}
              </p>
              <p class="font-body-md text-sm mt-1 truncate"
                 :class="conversation.unread ? 'text-on-surface font-semibold' : 'text-on-surface-variant'">
                {{ conversation.last_message }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Chat Window -->
      <div class="flex-1 flex flex-col bg-background relative"
           :class="{ 'hidden md:flex': !activeConversation }">
        
        <!-- No Conversation Selected -->
        <div v-if="!activeConversation" class="flex-1 flex items-center justify-center">
          <div class="text-center">
            <span class="material-symbols-outlined text-6xl text-on-surface-variant mb-4">chat</span>
            <h3 class="font-headline-md text-headline-md text-on-surface mb-2">Your Messages</h3>
            <p class="text-on-surface-variant">Select a conversation to start chatting</p>
          </div>
        </div>

        <!-- Active Chat -->
        <template v-else>
          <!-- Chat Header -->
          <div class="h-20 px-6 border-b border-surface-container-high bg-surface-container-low/80 backdrop-blur flex justify-between items-center z-10 shrink-0">
            <div class="flex items-center gap-4">
              <button 
                @click="closeConversation"
                class="md:hidden p-2 -ml-2 text-on-surface-variant"
              >
                <span class="material-symbols-outlined">arrow_back</span>
              </button>
              <div class="flex flex-col">
                <h2 class="font-headline-md text-headline-md text-on-surface">
                  {{ activeConversation.participant_name }}
                </h2>
                <span class="font-label-sm text-label-sm text-primary flex items-center gap-1">
                  <span class="w-2 h-2 rounded-full bg-primary inline-block"></span>
                  Online
                </span>
              </div>
            </div>
            <div class="flex items-center gap-4">
              <!-- Listing Info -->
              <div class="hidden sm:flex items-center gap-3 bg-surface-container px-3 py-1.5 rounded-lg border border-surface-container-high">
                <img 
                  :src="activeConversation.listing_image || 'https://via.placeholder.com/32'"
                  :alt="activeConversation.listing_title"
                  class="w-8 h-8 rounded object-cover"
                />
                <div class="flex flex-col">
                  <span class="font-label-sm text-label-sm text-on-surface truncate max-w-[150px]">
                    {{ activeConversation.listing_title }}
                  </span>
                  <span class="font-label-sm text-label-sm text-primary font-bold">
                    ${{ formatCurrency(activeConversation.listing_price) }}
                  </span>
                </div>
              </div>
              <button class="p-2 text-on-surface-variant hover:text-primary transition-colors">
                <span class="material-symbols-outlined">more_vert</span>
              </button>
            </div>
          </div>

          <!-- Chat Messages -->
          <div ref="messagesContainer" class="flex-1 overflow-y-auto p-6 space-y-6 custom-scrollbar bg-gradient-to-b from-background to-surface-container-lowest">
            <!-- Loading Messages -->
            <div v-if="loadingMessages" class="flex justify-center py-8">
              <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
            </div>

            <!-- Messages -->
            <template v-else>
              <div v-for="(group, groupIndex) in groupedMessages" :key="groupIndex">
                <!-- Date Separator -->
                <div class="flex justify-center mb-6">
                  <span class="px-3 py-1 bg-surface-container text-on-surface-variant font-label-sm text-label-sm rounded-full text-xs">
                    {{ group.date }}
                  </span>
                </div>

                <!-- Messages in group -->
                <div 
                  v-for="message in group.messages" 
                  :key="message.id"
                  class="flex gap-4 max-w-[85%]"
                  :class="message.sender_id === currentUserId ? 'ml-auto justify-end' : ''"
                >
                  <!-- Incoming Message Avatar -->
                  <div v-if="message.sender_id !== currentUserId" 
                       class="w-8 h-8 rounded-full bg-surface-container-high flex-shrink-0 flex items-center justify-center text-on-surface-variant font-label-md">
                    {{ getInitials(message.sender_name) }}
                  </div>

                  <div :class="message.sender_id === currentUserId ? 'flex flex-col items-end' : ''">
                    <!-- Message Bubble -->
                    <div 
                      class="p-4 rounded-2xl font-body-md shadow-sm relative overflow-hidden"
                      :class="message.sender_id === currentUserId 
                        ? 'bg-primary/10 border border-primary/20 rounded-tr-sm shadow-[0_0_15px_rgba(255,107,0,0.05)]' 
                        : 'bg-surface-container-low border border-surface-container-high rounded-tl-sm'"
                    >
                      <div v-if="message.sender_id === currentUserId" 
                           class="absolute top-0 right-0 w-20 h-20 bg-primary/5 blur-xl rounded-full"></div>
                      <p class="text-on-surface relative z-10">{{ message.message }}</p>
                    </div>

                    <!-- Timestamp and Status -->
                    <span class="font-label-sm text-label-sm text-on-surface-variant mt-1 flex items-center gap-1 px-1">
                      {{ formatMessageTime(message.sent_at) }}
                      <span 
                        v-if="message.sender_id === currentUserId"
                        class="material-symbols-outlined text-[14px]"
                        :class="message.is_read ? 'text-primary' : 'text-on-surface-variant'"
                      >
                        {{ message.is_read ? 'done_all' : 'check' }}
                      </span>
                    </span>
                  </div>

                  <!-- Outgoing Message (no avatar needed, shown at start for incoming) -->
                </div>
              </div>
            </template>

            <!-- Typing Indicator -->
            <div v-if="isTyping" class="flex gap-4 max-w-[85%]">
              <div class="w-8 h-8 rounded-full bg-surface-container-high flex-shrink-0 flex items-center justify-center text-on-surface-variant font-label-md">
                {{ getInitials(activeConversation.participant_name) }}
              </div>
              <div class="bg-surface-container-low border border-surface-container-high p-4 rounded-2xl rounded-tl-sm">
                <div class="flex gap-1">
                  <span class="w-2 h-2 bg-on-surface-variant rounded-full animate-bounce" style="animation-delay: 0ms"></span>
                  <span class="w-2 h-2 bg-on-surface-variant rounded-full animate-bounce" style="animation-delay: 150ms"></span>
                  <span class="w-2 h-2 bg-on-surface-variant rounded-full animate-bounce" style="animation-delay: 300ms"></span>
                </div>
              </div>
            </div>
          </div>

          <!-- Input Area -->
          <div class="p-4 bg-surface-container-low border-t border-surface-container-high shrink-0">
            <div class="flex items-end gap-2 bg-background border border-surface-container-high rounded-xl p-2 focus-within:border-primary/50 focus-within:shadow-[0_0_10px_rgba(255,107,0,0.1)] transition-all">
              <button 
                @click="triggerFileUpload"
                class="p-2 text-on-surface-variant hover:text-primary transition-colors flex-shrink-0"
                title="Attach image"
              >
                <span class="material-symbols-outlined">add_photo_alternate</span>
              </button>
              <input 
                ref="fileInput"
                type="file"
                accept="image/*"
                class="hidden"
                @change="handleFileUpload"
              />
              <textarea 
                ref="messageInput"
                v-model="newMessage"
                @keydown.enter.exact.prevent="sendMessage"
                @keydown.enter.shift.exact="newMessage += '\n'"
                @input="autoResize"
                class="flex-1 bg-transparent border-none text-on-surface font-body-md resize-none max-h-32 min-h-[44px] py-3 focus:ring-0 placeholder:text-on-surface-variant/50 custom-scrollbar"
                placeholder="Type a message..."
                rows="1"
              ></textarea>
              <button 
                @click="sendMessage"
                :disabled="!newMessage.trim()"
                class="p-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-background rounded-lg hover:shadow-primary/20 transition-all active:scale-95 flex-shrink-0 h-11 w-11 flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span class="material-symbols-outlined">send</span>
              </button>
            </div>
          </div>
        </template>
      </div>
    </main>

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
import { ref, computed, onMounted, nextTick, watch } from 'vue'
import { userDashboardService } from '@/services/userService'

// State
const loadingConversations = ref(true)
const loadingMessages = ref(false)
const conversationSearch = ref('')
const activeConversation = ref(null)
const newMessage = ref('')
const showNewMessage = ref(false)
const isTyping = ref(false)

// Refs
const messagesContainer = ref(null)
const messageInput = ref(null)
const fileInput = ref(null)

// User
const currentUserId = ref(parseInt(localStorage.getItem('userId') || '1'))

// Data
const conversations = ref([])
const messages = ref([])

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
const unreadCount = computed(() => {
  return conversations.value.filter(c => c.unread).length
})

const filteredConversations = computed(() => {
  if (!conversationSearch.value) return conversations.value
  const query = conversationSearch.value.toLowerCase()
  return conversations.value.filter(c => 
    c.participant_name.toLowerCase().includes(query) ||
    c.listing_title.toLowerCase().includes(query) ||
    c.last_message.toLowerCase().includes(query)
  )
})

const groupedMessages = computed(() => {
  const groups = []
  let currentDate = null
  let currentGroup = null

  messages.value.forEach(message => {
    const messageDate = new Date(message.sent_at).toLocaleDateString('en-US', {
      weekday: 'long',
      month: 'long',
      day: 'numeric'
    })

    if (messageDate !== currentDate) {
      currentDate = messageDate
      currentGroup = { date: messageDate, messages: [] }
      groups.push(currentGroup)
    }

    currentGroup.messages.push(message)
  })

  return groups
})

// Methods
const loadConversations = async () => {
  try {
    loadingConversations.value = true
    const userId = currentUserId.value
    
    try {
      const response = await userDashboardService.getConversations(userId)
      conversations.value = response.data || []
    } catch (apiError) {
      console.log('API not available, using demo data')
      // Demo data
      conversations.value = [
        {
          id: 1,
          participant_name: 'Alex Chen (Buyer)',
          participant_id: 2,
          listing_id: 1,
          listing_title: 'S-Works Tarmac SL7 - 56cm',
          listing_price: 5200,
          listing_image: 'https://via.placeholder.com/56',
          last_message: 'Can we do 5k flat if I pick it up today?',
          last_message_time: new Date(Date.now() - 3600000).toISOString(),
          unread: true
        },
        {
          id: 2,
          participant_name: 'Sarah Jenkins',
          participant_id: 3,
          listing_id: 2,
          listing_title: 'Fox Factory 36 Fork',
          listing_price: 850,
          listing_image: 'https://via.placeholder.com/56',
          last_message: 'Is the steerer tube cut?',
          last_message_time: new Date(Date.now() - 86400000).toISOString(),
          unread: true
        },
        {
          id: 3,
          participant_name: 'Marcus Vance (Seller)',
          participant_id: 4,
          listing_id: 3,
          listing_title: 'SRAM Red AXS Groupset',
          listing_price: 1800,
          listing_image: 'https://via.placeholder.com/56',
          last_message: 'Tracking number: 1Z999999999',
          last_message_time: new Date(Date.now() - 172800000).toISOString(),
          unread: false
        }
      ]
    }
  } catch (err) {
    console.error('Failed to load conversations:', err)
  } finally {
    loadingConversations.value = false
  }
}

const loadMessages = async (conversationId) => {
  try {
    loadingMessages.value = true
    
    try {
      const response = await userDashboardService.getConversationMessages(conversationId)
      messages.value = response.data || []
    } catch (apiError) {
      console.log('API not available, using demo messages')
      // Demo messages
      messages.value = [
        {
          id: 1,
          sender_id: 2,
          sender_name: 'Alex Chen',
          message: 'Hey! Really interested in the Tarmac. The condition looks great in the photos. Are there any scratches on the frame that aren\'t visible?',
          sent_at: new Date(Date.now() - 7200000).toISOString(),
          is_read: true
        },
        {
          id: 2,
          sender_id: currentUserId.value,
          sender_name: 'You',
          message: 'Hi Alex, thanks for reaching out. The frame is practically pristine. There\'s one tiny clear-coat scuff on the non-drive side chainstay from a dropped chain, but no carbon damage. I can send a close-up if you\'d like.',
          sent_at: new Date(Date.now() - 3600000).toISOString(),
          is_read: true
        },
        {
          id: 3,
          sender_id: 2,
          sender_name: 'Alex Chen',
          message: 'That would be great, thanks. If everything looks good, can we do 5k flat if I pick it up today? I\'m local to your area.',
          sent_at: new Date(Date.now() - 1800000).toISOString(),
          is_read: false
        }
      ]
    }
  } catch (err) {
    console.error('Failed to load messages:', err)
  } finally {
    loadingMessages.value = false
    scrollToBottom()
  }
}

const openConversation = (conversation) => {
  activeConversation.value = conversation
  
  // Mark as read
  if (conversation.unread) {
    conversation.unread = false
    // Call API to mark as read
    userDashboardService.markConversationRead(conversation.id).catch(console.error)
  }
  
  loadMessages(conversation.id)
}

const closeConversation = () => {
  activeConversation.value = null
  messages.value = []
}

const sendMessage = async () => {
  if (!newMessage.value.trim() || !activeConversation.value) return

  const message = {
    id: Date.now(),
    sender_id: currentUserId.value,
    sender_name: 'You',
    message: newMessage.value.trim(),
    sent_at: new Date().toISOString(),
    is_read: false
  }

  // Add to messages immediately (optimistic update)
  messages.value.push(message)
  newMessage.value = ''
  
  // Reset textarea height
  if (messageInput.value) {
    messageInput.value.style.height = '44px'
  }

  // Scroll to bottom
  await nextTick()
  scrollToBottom()

  // Update conversation preview
  if (activeConversation.value) {
    activeConversation.value.last_message = message.message
    activeConversation.value.last_message_time = message.sent_at
  }

  // Send to backend
  try {
    await userDashboardService.sendMessage({
      conversation_id: activeConversation.value.id,
      sender_id: currentUserId.value,
      message: message.message
    })
  } catch (err) {
    console.error('Failed to send message:', err)
    // Could add error indicator to message
  }
}

const autoResize = () => {
  const textarea = messageInput.value
  if (textarea) {
    textarea.style.height = '44px'
    textarea.style.height = textarea.scrollHeight + 'px'
  }
}

const scrollToBottom = () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}

const triggerFileUpload = () => {
  fileInput.value?.click()
}

const handleFileUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    console.log('File selected:', file.name)
    // TODO: Implement file upload to server
  }
}

const formatCurrency = (value) => {
  return (value || 0).toLocaleString('en-US')
}

const formatMessageTime = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMs / 3600000)
  const diffDays = Math.floor(diffMs / 86400000)

  if (diffMins < 1) return 'Just now'
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffHours < 24) return `${diffHours}h ago`
  if (diffDays < 7) return `${diffDays}d ago`
  
  return date.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit', hour12: true })
}

const getInitials = (name) => {
  if (!name) return '??'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

// Lifecycle
onMounted(() => {   
  loadConversations()
  
  // Poll for new messages every 30 seconds
  const pollInterval = setInterval(() => {
    loadConversations()
    if (activeConversation.value) {
      loadMessages(activeConversation.value.id)
    }
  }, 30000)
  
  // Clean up interval on component unmount
  return () => clearInterval(pollInterval)
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #353534;
  border-radius: 10px;
}
</style>