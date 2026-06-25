<template>
  <UserNav>
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
      <button @click="loadSalesData" class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors">Retry</button>
    </div>

    <!-- Sales Content -->
    <div v-else>
      <!-- Revenue Overview Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-gutter mb-gutter">
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-start mb-4">
            <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Total Revenue</h3>
            <span class="material-symbols-outlined text-primary">payments</span>
          </div>
          <div class="font-display text-display text-on-surface">${{ formatCurrency(salesStats.total_revenue) }}</div>
          <div class="flex items-center gap-1 mt-2">
            <span class="material-symbols-outlined text-secondary-container text-sm">trending_up</span>
            <span class="font-label-sm text-label-sm text-secondary-container">+{{ salesStats.revenue_growth }}%</span>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-start mb-4">
            <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Items Sold</h3>
            <span class="material-symbols-outlined text-primary">shopping_cart</span>
          </div>
          <div class="font-display text-display text-on-surface">{{ salesStats.total_sales }}</div>
          <div class="flex items-center gap-1 mt-2">
            <span class="material-symbols-outlined text-secondary-container text-sm">trending_up</span>
            <span class="font-label-sm text-label-sm text-secondary-container">+{{ salesStats.sales_growth }}%</span>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-start mb-4">
            <h3 class="font-label-md text-label-md text-on-surface-variant uppercase">Avg. Sale Price</h3>
            <span class="material-symbols-outlined text-primary">trending_up</span>
          </div>
          <div class="font-display text-display text-on-surface">${{ formatCurrency(salesStats.avg_price) }}</div>
          <div class="font-label-sm text-label-sm text-on-surface-variant mt-2">Per transaction</div>
        </div>

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
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Revenue Over Time</h3>
          <div class="h-64 flex items-end gap-2">
            <div v-for="(bar, index) in revenueChart" :key="index" class="flex-1 bg-primary/20 hover:bg-primary/40 rounded-t transition-all relative group" :style="{ height: `${(bar.value / maxRevenue) * 100}%` }">
              <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-surface-container-high text-on-surface text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">${{ bar.value }}</div>
            </div>
          </div>
          <div class="flex justify-between mt-3 text-xs text-on-surface-variant">
            <span v-for="(bar, index) in revenueChart" :key="index">{{ bar.label }}</span>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Sales by Category</h3>
          <div class="space-y-4">
            <div v-for="cat in categorySales" :key="cat.name" class="flex items-center gap-3">
              <span class="text-sm text-on-surface w-24 truncate">{{ cat.name }}</span>
              <div class="flex-1 h-2 bg-surface-variant rounded-full overflow-hidden">
                <div class="h-full bg-primary rounded-full" :style="{ width: `${(cat.count / maxCategoryCount) * 100}%` }"></div>
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
            <input v-model="transactionSearch" type="text" placeholder="Search transactions..." class="bg-transparent border-none text-sm text-on-surface focus:ring-0 placeholder-on-surface-variant" />
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
              </tr>
            </thead>
            <tbody class="divide-y divide-outline-variant">
              <tr v-for="t in filteredTransactions" :key="t.id" class="hover:bg-surface-container-high">
                <td class="p-4">
                  <div class="font-medium text-on-surface">{{ t.item_name }}</div>
                  <div class="text-sm text-on-surface-variant">{{ t.category }}</div>
                </td>
                <td class="p-4">
                  <div class="flex items-center gap-2">
                    <div class="w-6 h-6 rounded-full bg-primary/20 flex items-center justify-center text-xs text-primary font-semibold">{{ t.buyer_initials }}</div>
                    <span>{{ t.buyer_name }}</span>
                  </div>
                </td>
                <td class="p-4 text-on-surface-variant">{{ formatDate(t.date) }}</td>
                <td class="p-4 font-semibold text-primary">${{ formatCurrency(t.price) }}</td>
                <td class="p-4">
                  <span class="px-2 py-1 rounded-full text-xs font-semibold uppercase" :class="t.status === 'completed' ? 'bg-secondary-container/20 text-secondary-container' : 'bg-primary/20 text-primary'">{{ t.status }}</span>
                </td>
              </tr>
              <tr v-if="filteredTransactions.length === 0">
                <td colspan="5" class="p-8 text-center text-on-surface-variant">No transactions found</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </UserNav>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { userDashboardService } from '@/services/userService'
import UserNav from '@/components/User_Nav.vue'

const loading = ref(true)
const error = ref(null)
const dateRange = ref('30d')
const transactionSearch = ref('')

const salesStats = ref({ total_revenue: 24500, revenue_growth: 15, total_sales: 48, sales_growth: 12, avg_price: 510, unique_buyers: 32 })

const revenueChart = ref([
  { label: 'Mon', value: 3200 }, { label: 'Tue', value: 2800 }, { label: 'Wed', value: 4100 },
  { label: 'Thu', value: 3600 }, { label: 'Fri', value: 5200 }, { label: 'Sat', value: 4800 }, { label: 'Sun', value: 3800 }
])

const categorySales = ref([
  { name: 'Road Bikes', count: 18, revenue: 12400 }, { name: 'Mountain Bikes', count: 12, revenue: 8200 },
  { name: 'Parts', count: 10, revenue: 2400 }, { name: 'Accessories', count: 8, revenue: 1500 }
])

const transactions = ref([
  { id: 1, item_name: 'Aero VeloX Pro 2024', category: 'Road Bike', buyer_name: 'Marcus Johnson', buyer_initials: 'MJ', date: '2026-06-23', price: 4200, status: 'completed' },
  { id: 2, item_name: 'TrailBlazer MTB XT', category: 'Mountain Bike', buyer_name: 'Sarah Chen', buyer_initials: 'SC', date: '2026-06-21', price: 2850, status: 'pending' }
])

const maxRevenue = computed(() => Math.max(...revenueChart.value.map(b => b.value), 1))
const maxCategoryCount = computed(() => Math.max(...categorySales.value.map(c => c.count), 1))

const filteredTransactions = computed(() => {
  if (!transactionSearch.value) return transactions.value
  const q = transactionSearch.value.toLowerCase()
  return transactions.value.filter(t => t.item_name.toLowerCase().includes(q) || t.buyer_name.toLowerCase().includes(q))
})

const loadSalesData = async () => {
  try {
    loading.value = true; error.value = null
    const userId = localStorage.getItem('userId') || '1'
    try {
      const response = await userDashboardService.getSales(userId)
      const data = response.data
      salesStats.value = { total_revenue: data.total_revenue || 24500, revenue_growth: data.revenue_growth || 15, total_sales: data.total_sales || 48, sales_growth: data.sales_growth || 12, avg_price: data.avg_price || 510, unique_buyers: data.unique_buyers || 32 }
    } catch (apiError) { console.log('Using demo data') }
  } catch (err) { error.value = 'Failed to load sales data.' }
  finally { loading.value = false }
}

const formatCurrency = (v) => (v || 0).toLocaleString('en-US')
const formatDate = (d) => d ? new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' }) : ''

onMounted(() => { loadSalesData() })
</script>