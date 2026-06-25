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
        </router-link>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="flex-1 md:ml-[280px] p-margin-mobile md:p-margin-desktop pb-24 md:pb-margin-desktop">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
        <div>
          <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Sales Dashboard</h1>
          <p class="font-body-md text-body-md text-on-surface-variant mt-1">Track your revenue, transactions, and sales performance.</p>
        </div>
        <div class="flex gap-3">
          <select 
            v-model="dateRange"
            class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors"
          >
            <option value="7d">Last 7 days</option>
            <option value="30d">Last 30 days</option>
            <option value="90d">Last 90 days</option>
            <option value="1y">This year</option>
            <option value="all">All time</option>
          </select>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
          <p class="text-on-surface-variant">Loading sales data...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
        <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
        <p class="text-error font-label-md mb-4">{{ error }}</p>
        <button 
          @click="loadSalesData"
          class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors"
        >
          Retry
        </button>
      </div>

      <!-- Sales Content -->
      <div v-else>
        <!-- Revenue Overview Cards -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-gutter mb-gutter">
          <!-- Total Revenue -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <div class="flex justify-between items-start mb-4">
              <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Total Revenue</h3>
              <span class="material-symbols-outlined text-primary">payments</span>
            </div>
            <div class="font-display text-display text-on-surface">${{ formatCurrency(salesStats.total_revenue) }}</div>
            <div class="flex items-center gap-1 mt-2">
              <span class="material-symbols-outlined text-secondary-container text-sm">trending_up</span>
              <span class="font-label-sm text-label-sm text-secondary-container">+{{ salesStats.revenue_growth }}% vs previous</span>
            </div>
          </div>

          <!-- Total Sales -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <div class="flex justify-between items-start mb-4">
              <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Items Sold</h3>
              <span class="material-symbols-outlined text-primary">shopping_cart</span>
            </div>
            <div class="font-display text-display text-on-surface">{{ salesStats.total_sales }}</div>
            <div class="flex items-center gap-1 mt-2">
              <span class="material-symbols-outlined text-secondary-container text-sm">trending_up</span>
              <span class="font-label-sm text-label-sm text-secondary-container">+{{ salesStats.sales_growth }}% vs previous</span>
            </div>
          </div>

          <!-- Average Price -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <div class="flex justify-between items-start mb-4">
              <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Avg. Sale Price</h3>
              <span class="material-symbols-outlined text-primary">trending_up</span>
            </div>
            <div class="font-display text-display text-on-surface">${{ formatCurrency(salesStats.avg_price) }}</div>
            <div class="font-label-sm text-label-sm text-on-surface-variant mt-2">Per transaction</div>
          </div>

          <!-- Active Buyers -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <div class="flex justify-between items-start mb-4">
              <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Unique Buyers</h3>
              <span class="material-symbols-outlined text-primary">group</span>
            </div>
            <div class="font-display text-display text-on-surface">{{ salesStats.unique_buyers }}</div>
            <div class="font-label-sm text-label-sm text-on-surface-variant mt-2">Active customers</div>
          </div>
        </div>

        <!-- Charts Section -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-gutter mb-gutter">
          <!-- Revenue Chart -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Revenue Over Time</h3>
            <div class="h-64 flex items-end gap-2">
              <div 
                v-for="(bar, index) in revenueChart" 
                :key="index"
                class="flex-1 bg-primary/20 hover:bg-primary/40 rounded-t transition-all relative group"
                :style="{ height: `${(bar.value / maxRevenue) * 100}%` }"
              >
                <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-surface-container-high text-on-surface text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                  ${{ bar.value }}
                </div>
              </div>
            </div>
            <div class="flex justify-between mt-3 text-xs text-on-surface-variant">
              <span v-for="(bar, index) in revenueChart" :key="index">{{ bar.label }}</span>
            </div>
          </div>

          <!-- Sales by Category -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Sales by Category</h3>
            <div class="space-y-4">
              <div v-for="cat in categorySales" :key="cat.name" class="flex items-center gap-3">
                <span class="text-sm text-on-surface w-24 truncate">{{ cat.name }}</span>
                <div class="flex-1 h-2 bg-surface-variant rounded-full overflow-hidden">
                  <div 
                    class="h-full bg-primary rounded-full" 
                    :style="{ width: `${(cat.count / maxCategoryCount) * 100}%` }"
                  ></div>
                </div>
                <span class="text-sm text-on-surface-variant w-12 text-right">{{ cat.count }}</span>
                <span class="text-sm text-primary w-20 text-right">${{ formatCurrency(cat.revenue) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Recent Transactions Table -->
        <div class="bg-surface-container-low border border-outline-variant rounded-xl shadow-sm overflow-hidden">
          <div class="p-6 border-b border-outline-variant flex justify-between items-center">
            <h3 class="font-headline-md text-headline-md text-on-surface">Recent Transactions</h3>
            <div class="flex items-center gap-2">
              <span class="material-symbols-outlined text-on-surface-variant text-lg">search</span>
              <input 
                v-model="transactionSearch"
                type="text"
                placeholder="Search transactions..."
                class="bg-transparent border-none text-sm text-on-surface focus:ring-0 placeholder-on-surface-variant"
              />
            </div>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="bg-surface-container text-on-surface-variant font-label-sm text-label-sm uppercase">
                  <th class="p-4 font-semibold">Item</th>
                  <th class="p-4 font-semibold">Buyer</th>
                  <th class="p-4 font-semibold">Date</th>
                  <th class="p-4 font-semibold">Price</th>
                  <th class="p-4 font-semibold">Status</th>
                  <th class="p-4 font-semibold text-right">Actions</th>
                </tr>
              </thead>
              <tbody class="font-body-md text-body-md text-on-surface divide-y divide-outline-variant">
                <tr 
                  v-for="transaction in filteredTransactions" 
                  :key="transaction.id"
                  class="hover:bg-surface-container-high transition-colors"
                >
                  <td class="p-4">
                    <div class="flex items-center gap-3">
                      <div class="w-10 h-10 rounded bg-surface overflow-hidden">
                        <img 
                          :src="transaction.image || 'https://via.placeholder.com/40'"
                          :alt="transaction.item_name"
                          class="w-full h-full object-cover"
                        />
                      </div>
                      <div>
                        <div class="font-medium text-on-surface">{{ transaction.item_name }}</div>
                        <div class="text-sm text-on-surface-variant">{{ transaction.category }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="p-4">
                    <div class="flex items-center gap-2">
                      <div class="w-6 h-6 rounded-full bg-primary/20 flex items-center justify-center text-xs text-primary font-semibold">
                        {{ transaction.buyer_initials }}
                      </div>
                      <span>{{ transaction.buyer_name }}</span>
                    </div>
                  </td>
                  <td class="p-4 text-on-surface-variant">{{ formatDate(transaction.date) }}</td>
                  <td class="p-4 font-semibold text-primary">${{ formatCurrency(transaction.price) }}</td>
                  <td class="p-4">
                    <span 
                      class="inline-block px-2 py-1 rounded-full text-xs font-semibold uppercase tracking-wider"
                      :class="getStatusClass(transaction.status)"
                    >
                      {{ transaction.status }}
                    </span>
                  </td>
                  <td class="p-4 text-right">
                    <div class="flex gap-2 justify-end">
                      <button 
                        @click="viewTransaction(transaction.id)"
                        class="text-on-surface-variant hover:text-primary transition-colors"
                        title="View Details"
                      >
                        <span class="material-symbols-outlined text-lg">visibility</span>
                      </button>
                      <button 
                        v-if="transaction.status === 'pending'"
                        @click="contactBuyer(transaction)"
                        class="text-on-surface-variant hover:text-primary transition-colors"
                        title="Contact Buyer"
                      >
                        <span class="material-symbols-outlined text-lg">chat</span>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="filteredTransactions.length === 0">
                  <td colspan="6" class="p-8 text-center text-on-surface-variant">
                    <span class="material-symbols-outlined text-4xl mb-2 block">receipt_long</span>
                    No transactions found
                  </td>
                </tr>
              </tbody>
            </table>
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
        :class="$route.path === item.path 
          ? 'text-primary bg-primary/10' 
          : 'text-on-surface-variant active:bg-surface-container-high'"
      >
        <span class="material-symbols-outlined">{{ item.icon }}</span>
        <span class="font-label-sm text-label-sm mt-1">{{ item.name }}</span>
      </router-link>
    </nav>

    <!-- Transaction Detail Modal -->
    <div v-if="showTransactionModal" class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4">
      <div class="bg-surface-container-high border border-outline-variant rounded-xl p-6 max-w-lg w-full">
        <div class="flex justify-between items-center mb-6">
          <h3 class="font-headline-md text-headline-md text-on-surface">Transaction Details</h3>
          <button 
            @click="showTransactionModal = false"
            class="text-on-surface-variant hover:text-on-surface transition-colors"
          >
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>
        
        <div v-if="selectedTransaction" class="space-y-4">
          <div class="flex justify-between py-2 border-b border-outline-variant">
            <span class="text-on-surface-variant">Item</span>
            <span class="text-on-surface font-medium">{{ selectedTransaction.item_name }}</span>
          </div>
          <div class="flex justify-between py-2 border-b border-outline-variant">
            <span class="text-on-surface-variant">Buyer</span>
            <span class="text-on-surface">{{ selectedTransaction.buyer_name }}</span>
          </div>
          <div class="flex justify-between py-2 border-b border-outline-variant">
            <span class="text-on-surface-variant">Date</span>
            <span class="text-on-surface">{{ formatDate(selectedTransaction.date) }}</span>
          </div>
          <div class="flex justify-between py-2 border-b border-outline-variant">
            <span class="text-on-surface-variant">Price</span>
            <span class="text-primary font-semibold">${{ formatCurrency(selectedTransaction.price) }}</span>
          </div>
          <div class="flex justify-between py-2 border-b border-outline-variant">
            <span class="text-on-surface-variant">Status</span>
            <span :class="getStatusClass(selectedTransaction.status)">{{ selectedTransaction.status }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { userDashboardService } from '@/services/userService'

// State
const loading = ref(true)
const error = ref(null)
const dateRange = ref('30d')
const transactionSearch = ref('')
const showTransactionModal = ref(false)
const selectedTransaction = ref(null)

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

// Sales stats
const salesStats = ref({
  total_revenue: 0,
  revenue_growth: 0,
  total_sales: 0,
  sales_growth: 0,
  avg_price: 0,
  unique_buyers: 0
})

// Revenue chart data
const revenueChart = ref([
  { label: 'Mon', value: 0 },
  { label: 'Tue', value: 0 },
  { label: 'Wed', value: 0 },
  { label: 'Thu', value: 0 },
  { label: 'Fri', value: 0 },
  { label: 'Sat', value: 0 },
  { label: 'Sun', value: 0 }
])

// Category sales data
const categorySales = ref([
  { name: 'Road Bikes', count: 0, revenue: 0 },
  { name: 'Mountain Bikes', count: 0, revenue: 0 },
  { name: 'Parts', count: 0, revenue: 0 },
  { name: 'Accessories', count: 0, revenue: 0 }
])

// Transactions
const transactions = ref([])

// Computed
const maxRevenue = computed(() => {
  const max = Math.max(...revenueChart.value.map(bar => bar.value))
  return max > 0 ? max : 1
})

const maxCategoryCount = computed(() => {
  const max = Math.max(...categorySales.value.map(cat => cat.count))
  return max > 0 ? max : 1
})

const filteredTransactions = computed(() => {
  if (!transactionSearch.value) return transactions.value
  const query = transactionSearch.value.toLowerCase()
  return transactions.value.filter(t => 
    t.item_name.toLowerCase().includes(query) ||
    t.buyer_name.toLowerCase().includes(query)
  )
})

// Methods
const loadSalesData = async () => {
  try {
    loading.value = true
    error.value = null
    
    const userId = localStorage.getItem('userId') || '1'
    
    // Try to fetch from Go backend
    try {
      const response = await userDashboardService.getSales(userId)
      const data = response.data
      
      salesStats.value = {
        total_revenue: data.total_revenue || 24500,
        revenue_growth: data.revenue_growth || 15,
        total_sales: data.total_sales || 48,
        sales_growth: data.sales_growth || 12,
        avg_price: data.avg_price || 510,
        unique_buyers: data.unique_buyers || 32
      }
      
      revenueChart.value = data.revenue_chart || [
        { label: 'Mon', value: 3200 },
        { label: 'Tue', value: 2800 },
        { label: 'Wed', value: 4100 },
        { label: 'Thu', value: 3600 },
        { label: 'Fri', value: 5200 },
        { label: 'Sat', value: 4800 },
        { label: 'Sun', value: 3800 }
      ]
      
      categorySales.value = data.category_sales || [
        { name: 'Road Bikes', count: 18, revenue: 12400 },
        { name: 'Mountain Bikes', count: 12, revenue: 8200 },
        { name: 'Parts', count: 10, revenue: 2400 },
        { name: 'Accessories', count: 8, revenue: 1500 }
      ]
      
      transactions.value = data.transactions || []
      
    } catch (apiError) {
      console.log('API not available, using demo data:', apiError)
      // Use demo data if API isn't ready
      loadDemoData()
    }
    
  } catch (err) {
    console.error('Failed to load sales data:', err)
    error.value = 'Failed to load sales data. Please try again.'
  } finally {
    loading.value = false
  }
}

const loadDemoData = () => {
  // Demo transactions
  transactions.value = [
    {
      id: 1,
      item_name: 'Aero VeloX Pro 2024',
      category: 'Road Bike',
      buyer_name: 'Marcus Johnson',
      buyer_initials: 'MJ',
      date: '2026-06-23',
      price: 4200,
      status: 'completed',
      image: 'https://via.placeholder.com/40'
    },
    {
      id: 2,
      item_name: 'TrailBlazer MTB XT',
      category: 'Mountain Bike',
      buyer_name: 'Sarah Chen',
      buyer_initials: 'SC',
      date: '2026-06-21',
      price: 2850,
      status: 'pending',
      image: 'https://via.placeholder.com/40'
    },
    {
      id: 3,
      item_name: 'Carbon Handlebar Set',
      category: 'Parts',
      buyer_name: 'Alex Thompson',
      buyer_initials: 'AT',
      date: '2026-06-20',
      price: 350,
      status: 'completed',
      image: 'https://via.placeholder.com/40'
    },
    {
      id: 4,
      item_name: 'Shimano XTR Groupset',
      category: 'Parts',
      buyer_name: 'David Wilson',
      buyer_initials: 'DW',
      date: '2026-06-18',
      price: 1200,
      status: 'completed',
      image: 'https://via.placeholder.com/40'
    },
    {
      id: 5,
      item_name: 'Specialized S-Works',
      category: 'Road Bike',
      buyer_name: 'Emily Brown',
      buyer_initials: 'EB',
      date: '2026-06-15',
      price: 5800,
      status: 'completed',
      image: 'https://via.placeholder.com/40'
    }
  ]
}

const formatCurrency = (value) => {
  return (value || 0).toLocaleString('en-US', { minimumFractionDigits: 0, maximumFractionDigits: 0 })
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

const getStatusClass = (status) => {
  switch (status) {
    case 'completed':
      return 'bg-secondary-container/20 text-secondary-container'
    case 'pending':
      return 'bg-primary/20 text-primary'
    case 'cancelled':
      return 'bg-error/20 text-error'
    case 'refunded':
      return 'bg-tertiary/20 text-tertiary'
    default:
      return 'bg-surface-variant text-on-surface-variant'
  }
}

const viewTransaction = (id) => {
  selectedTransaction.value = transactions.value.find(t => t.id === id)
  showTransactionModal.value = true
}

const contactBuyer = (transaction) => {
  console.log('Contact buyer:', transaction.buyer_name)
  // Navigate to messages with this buyer
}
</script>