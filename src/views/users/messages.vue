<template>
  <UserNav>
    <!-- Main Content Area (Messaging Interface) -->
    <div class="flex flex-col md:flex-row" style="height: calc(100vh - 200px); min-height: 500px;">
      <!-- Conversations List Sidebar -->
      <div class="w-full md:w-[350px] border-r border-surface-container-high flex flex-col bg-surface-container-low/50 shrink-0"
           :class="{ 'hidden md:flex': activeConversation }">
        <!-- Inbox Header -->
        <div class="p-4 border-b border-surface-container-high bg-surface-container-low flex justify-between items-center shrink-0">
          <h2 class="font-headline-md text-headline-md text-on-surface">Messages</h2>
          <button @click="showNewMessage = true" class="p-2 rounded-full bg-primary/10 text-primary hover:bg-primary/20 transition-colors" title="New Message">
            <span class="material-symbols-outlined">edit_square</span>
          </button>
        </div>

        <!-- Search -->
        <div class="p-3 border-b border-surface-container-high shrink-0">
          <div class="relative">
            <span class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant text-lg">search</span>
            <input v-model="conversationSearch" type="text" placeholder="Search conversations..." class="w-full bg-surface border border-outline-variant rounded-lg py-2 pl-10 pr-4 text-sm text-on-surface focus:border-primary transition-colors" />
          </div>
        </div>

        <!-- Loading -->
        <div v-if="loadingConversations" class="flex-1 flex items-center justify-center">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
        </div>

        <!-- Conversations List -->
        <div v-else class="overflow-y-auto flex-1 custom-scrollbar">
          <div v-if="filteredConversations.length === 0" class="p-8 text-center text-on-surface-variant">
            <span class="material-symbols-outlined text-4xl mb-2">forum</span>
            <p>No conversations yet</p>
          </div>

          <div v-for="conv in filteredConversations" :key="conv.id" @click="openConversation(conv)"
            class="p-4 border-l-4 cursor-pointer flex gap-3 transition-colors hover:bg-surface-container-high"
            :class="activeConversation?.id === conv.id ? 'border-primary bg-primary/5' : conv.unread ? 'border-transparent' : 'border-transparent opacity-70'">
            <div class="relative shrink-0">
              <div class="w-12 h-12 rounded-lg bg-surface-variant flex items-center justify-center text-on-surface-variant text-lg font-bold">
                {{ getInitials(conv.participant_name) }}
              </div>
              <div v-if="conv.unread" class="absolute -top-1 -right-1 w-3 h-3 bg-primary rounded-full border-2 border-surface-container-low"></div>
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex justify-between items-baseline mb-0.5">
                <h3 class="font-label-md text-label-md truncate" :class="conv.unread ? 'text-on-surface font-bold' : 'text-on-surface'">{{ conv.participant_name }}</h3>
                <span class="text-xs shrink-0 ml-2" :class="conv.unread ? 'text-primary' : 'text-on-surface-variant'">{{ formatMessageTime(conv.last_message_time) }}</span>
              </div>
              <p class="text-xs text-on-surface-variant truncate">{{ conv.listing_title }}</p>
              <p class="text-sm mt-0.5 truncate" :class="conv.unread ? 'text-on-surface font-semibold' : 'text-on-surface-variant'">{{ conv.last_message }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Chat Window -->
      <div class="flex-1 flex flex-col bg-background" :class="{ 'hidden md:flex': !activeConversation }">
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
          <div class="h-14 px-4 border-b border-surface-container-high bg-surface-container-low/80 flex justify-between items-center shrink-0">
            <div class="flex items-center gap-3">
              <button @click="closeConversation" class="md:hidden p-1 text-on-surface-variant">
                <span class="material-symbols-outlined">arrow_back</span>
              </button>
              <div class="w-8 h-8 rounded-full bg-primary/20 flex items-center justify-center text-primary font-bold text-sm">{{ getInitials(activeConversation.participant_name) }}</div>
              <div>
                <h2 class="font-label-md text-label-md text-on-surface">{{ activeConversation.participant_name }}</h2>
                <span class="text-xs text-on-surface-variant">{{ activeConversation.listing_title }}</span>
              </div>
            </div>
            <span class="font-label-sm text-label-sm text-primary font-bold">${{ formatCurrency(activeConversation.listing_price) }}</span>
          </div>

          <!-- Messages Area -->
          <div ref="messagesContainer" class="flex-1 overflow-y-auto p-4 space-y-4 custom-scrollbar">
            <div v-if="loadingMessages" class="flex justify-center py-8">
              <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
            </div>
            <template v-else>
              <div v-for="msg in messages" :key="msg.id" class="flex gap-2" :class="msg.sender_id === currentUserId ? 'justify-end' : 'justify-start'">
                <div v-if="msg.sender_id !== currentUserId" class="w-7 h-7 rounded-full bg-surface-container-high flex-shrink-0 flex items-center justify-center text-on-surface-variant text-xs font-bold">{{ getInitials(msg.sender_name) }}</div>
                <div class="max-w-[75%] p-3 text-sm text-on-surface" :class="msg.sender_id === currentUserId ? 'bg-primary/10 border border-primary/20 rounded-2xl rounded-tr-sm' : 'bg-surface-container-low border border-surface-container-high rounded-2xl rounded-tl-sm'">
                  {{ msg.message }}
                </div>
              </div>
            </template>
          </div>

          <!-- Input Area -->
          <div class="p-3 bg-surface-container-low border-t border-surface-container-high shrink-0">
            <div class="flex items-end gap-2 bg-background border border-surface-container-high rounded-xl p-2">
              <textarea ref="messageInput" v-model="newMessage" @keydown.enter.exact.prevent="sendMessage" @input="autoResize" class="flex-1 bg-transparent border-none text-on-surface text-sm resize-none max-h-32 min-h-[40px] py-2 focus:ring-0 placeholder:text-on-surface-variant/50" placeholder="Type a message..." rows="1"></textarea>
              <button @click="sendMessage" :disabled="!newMessage.trim()" class="p-2 bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black rounded-lg hover:shadow-primary/20 transition-all active:scale-95 flex-shrink-0 h-10 w-10 flex items-center justify-center disabled:opacity-50">
                <span class="material-symbols-outlined text-lg">send</span>
              </button>
            </div>
          </div>
        </template>
      </div>
    </div>
  </UserNav>
</template>

<script setup>
import { ref, computed, onMounted, nextTick } from 'vue'
import UserNav from '@/components/User_Nav.vue'

const loadingConversations = ref(true)
const loadingMessages = ref(false)
const conversationSearch = ref('')
const activeConversation = ref(null)
const newMessage = ref('')
const showNewMessage = ref(false)
const messagesContainer = ref(null)
const messageInput = ref(null)
const currentUserId = ref(parseInt(localStorage.getItem('userId') || '1'))
const conversations = ref([])
const messages = ref([])

const filteredConversations = computed(() => {
  if (!conversationSearch.value) return conversations.value
  const q = conversationSearch.value.toLowerCase()
  return conversations.value.filter(c => c.participant_name.toLowerCase().includes(q) || c.listing_title.toLowerCase().includes(q) || c.last_message.toLowerCase().includes(q))
})

const loadConversations = async () => {
  loadingConversations.value = true
  conversations.value = [
    { id: 1, participant_name: 'Alex Chen (Buyer)', participant_id: 2, listing_title: 'S-Works Tarmac SL7', listing_price: 5200, last_message: 'Can we do 5k flat if I pick it up today?', last_message_time: new Date(Date.now() - 3600000).toISOString(), unread: true },
    { id: 2, participant_name: 'Sarah Jenkins', participant_id: 3, listing_title: 'Fox Factory 36 Fork', listing_price: 850, last_message: 'Is the steerer tube cut?', last_message_time: new Date(Date.now() - 86400000).toISOString(), unread: false },
    { id: 3, participant_name: 'Marcus Vance', participant_id: 4, listing_title: 'SRAM Red AXS Groupset', listing_price: 1800, last_message: 'Tracking number sent.', last_message_time: new Date(Date.now() - 172800000).toISOString(), unread: false }
  ]
  loadingConversations.value = false
}

const loadMessages = async () => {
  loadingMessages.value = true
  messages.value = [
    { id: 1, sender_id: 2, sender_name: 'Alex Chen', message: 'Hey! Really interested in the Tarmac. Are there any scratches?', sent_at: new Date(Date.now() - 7200000).toISOString(), is_read: true },
    { id: 2, sender_id: currentUserId.value, sender_name: 'You', message: 'Hi Alex! Frame is pristine. One tiny clear-coat scuff but no carbon damage.', sent_at: new Date(Date.now() - 3600000).toISOString(), is_read: true },
    { id: 3, sender_id: 2, sender_name: 'Alex Chen', message: 'Great, can we do 5k if I pick up today?', sent_at: new Date(Date.now() - 1800000).toISOString(), is_read: false }
  ]
  loadingMessages.value = false
  scrollToBottom()
}

const openConversation = (conv) => { activeConversation.value = conv; if (conv.unread) conv.unread = false; loadMessages(conv.id) }
const closeConversation = () => { activeConversation.value = null; messages.value = [] }

const sendMessage = async () => {
  if (!newMessage.value.trim() || !activeConversation.value) return
  messages.value.push({ id: Date.now(), sender_id: currentUserId.value, sender_name: 'You', message: newMessage.value.trim(), sent_at: new Date().toISOString(), is_read: false })
  newMessage.value = ''
  if (messageInput.value) messageInput.value.style.height = '40px'
  await nextTick(); scrollToBottom()
}

const autoResize = () => { if (messageInput.value) { messageInput.value.style.height = '40px'; messageInput.value.style.height = messageInput.value.scrollHeight + 'px' } }
const scrollToBottom = () => { nextTick(() => { if (messagesContainer.value) messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight }) }
const formatCurrency = (v) => (v || 0).toLocaleString('en-US')
const formatMessageTime = (d) => { if (!d) return ''; const m = Math.floor((new Date() - new Date(d)) / 60000); if (m < 1) return 'Now'; if (m < 60) return m + 'm'; const h = Math.floor(m / 60); if (h < 24) return h + 'h'; return Math.floor(h / 24) + 'd' }
const getInitials = (n) => n ? n.split(' ').map(w => w[0]).join('').toUpperCase().slice(0, 2) : '??'

onMounted(() => { loadConversations() })
</script>