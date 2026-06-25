<template>
  <UserNav>
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-gutter gap-4">
      <div>
        <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Analytics</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-1">Deep insights into your listing performance and market trends.</p>
      </div>
      <div class="flex gap-3">
        <select v-model="timeRange" class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors">
          <option value="7d">Last 7 days</option>
          <option value="30d">Last 30 days</option>
          <option value="90d">Last 90 days</option>
          <option value="1y">This year</option>
        </select>
        <button @click="exportData" class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface hover:bg-surface-container-high transition-colors flex items-center gap-2">
          <span class="material-symbols-outlined text-lg">download</span>
          <span class="hidden sm:inline">Export</span>
        </button>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
        <p class="text-on-surface-variant">Analyzing your data...</p>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
      <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
      <p class="text-error font-label-md mb-4">{{ error }}</p>
      <button @click="loadAnalytics" class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors">Retry</button>
    </div>

    <!-- Content -->
    <div v-else>
      <!-- Performance Cards -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-gutter">
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
          <div class="flex items-center gap-2 mb-2">
            <span class="material-symbols-outlined text-primary">visibility</span>
            <span class="font-label-sm text-on-surface-variant">Total Views</span>
          </div>
          <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_views }}</div>
          <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.views_growth }}%</div>
        </div>
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
          <div class="flex items-center gap-2 mb-2">
            <span class="material-symbols-outlined text-primary">favorite</span>
            <span class="font-label-sm text-on-surface-variant">Favorites</span>
          </div>
          <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_favorites }}</div>
          <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.favorites_growth }}%</div>
        </div>
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
          <div class="flex items-center gap-2 mb-2">
            <span class="material-symbols-outlined text-primary">chat</span>
            <span class="font-label-sm text-on-surface-variant">Inquiries</span>
          </div>
          <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_inquiries }}</div>
          <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.inquiries_growth }}%</div>
        </div>
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
          <div class="flex items-center gap-2 mb-2">
            <span class="material-symbols-outlined text-primary">trending_up</span>
            <span class="font-label-sm text-on-surface-variant">Conversion</span>
          </div>
          <div class="font-display text-3xl text-on-surface">{{ analytics.performance.conversion_rate }}%</div>
          <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.conversion_growth }}%</div>
        </div>
      </div>

      <!-- Charts Row -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-gutter mb-gutter">
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Views Over Time</h3>
          <div class="h-64 flex items-end gap-2">
            <div v-for="(point, i) in analytics.views_over_time" :key="i" class="flex-1 bg-primary/20 hover:bg-primary/40 rounded-t transition-all relative group" :style="{ height: `${(point.views / maxViews) * 100}%` }">
              <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-surface-container-high text-on-surface text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 whitespace-nowrap">{{ point.views }} views</div>
            </div>
          </div>
          <div class="flex justify-between mt-3 text-xs text-on-surface-variant">
            <span v-for="(point, i) in analytics.views_over_time" :key="i">{{ point.label }}</span>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Top Listings</h3>
          <div class="space-y-4">
            <div v-for="(listing, i) in analytics.top_listings" :key="listing.id" class="flex items-center gap-4 p-3 rounded-lg hover:bg-surface-container-high transition-colors">
              <div class="w-8 h-8 rounded-full bg-primary/10 flex items-center justify-center text-primary font-bold text-sm">#{{ i + 1 }}</div>
              <div class="w-10 h-10 rounded bg-surface overflow-hidden shrink-0 flex items-center justify-center text-on-surface-variant text-xs">IMG</div>
              <div class="flex-1 min-w-0">
                <div class="font-medium text-on-surface truncate">{{ listing.title }}</div>
                <div class="text-xs text-on-surface-variant">{{ listing.views }} views • {{ listing.inquiries }} inquiries</div>
              </div>
              <div class="text-right">
                <div class="text-primary font-semibold">${{ formatCurrency(listing.price) }}</div>
                <div class="text-xs text-secondary-container">+{{ listing.growth }}%</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Traffic & Peak Times -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-gutter mb-gutter">
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Traffic Sources</h3>
          <div class="space-y-4">
            <div v-for="source in analytics.traffic_sources" :key="source.name" class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <span class="material-symbols-outlined text-on-surface-variant">{{ source.icon }}</span>
                <div>
                  <div class="text-sm text-on-surface">{{ source.name }}</div>
                  <div class="text-xs text-on-surface-variant">{{ source.percentage }}%</div>
                </div>
              </div>
              <div class="w-24 h-2 bg-surface-variant rounded-full overflow-hidden">
                <div class="h-full bg-primary rounded-full" :style="{ width: `${source.percentage}%` }"></div>
              </div>
            </div>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Peak Activity</h3>
          <div class="space-y-3">
            <div v-for="time in analytics.peak_times" :key="time.day" class="flex items-center justify-between">
              <span class="text-sm text-on-surface w-20">{{ time.day }}</span>
              <div class="flex-1 mx-3 h-2 bg-surface-variant rounded-full overflow-hidden">
                <div class="h-full bg-tertiary rounded-full" :style="{ width: `${time.percentage}%` }"></div>
              </div>
              <span class="text-sm text-on-surface-variant w-16 text-right">{{ time.hours }}</span>
            </div>
          </div>
        </div>

        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Price Ranges</h3>
          <div class="space-y-3">
            <div v-for="range in analytics.price_ranges" :key="range.label" class="space-y-1">
              <div class="flex justify-between text-sm">
                <span class="text-on-surface">{{ range.label }}</span>
                <span class="text-on-surface-variant">{{ range.count }} sold</span>
              </div>
              <div class="h-2 bg-surface-variant rounded-full overflow-hidden">
                <div class="h-full bg-primary rounded-full" :style="{ width: `${(range.count / maxPriceCount) * 100}%` }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Insights -->
      <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
        <h3 class="font-headline-md text-headline-md text-on-surface mb-6 flex items-center gap-2">
          <span class="material-symbols-outlined text-primary">tips_and_updates</span>
          Insights & Recommendations
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div v-for="insight in analytics.insights" :key="insight.id" class="p-4 rounded-lg border" :class="insight.type === 'success' ? 'border-secondary-container/30 bg-secondary-container/5' : 'border-primary/30 bg-primary/5'">
            <div class="flex items-start gap-3">
              <span class="material-symbols-outlined text-lg mt-0.5" :class="insight.type === 'success' ? 'text-secondary-container' : 'text-primary'">{{ insight.type === 'success' ? 'check_circle' : 'lightbulb' }}</span>
              <div>
                <h4 class="font-label-md text-label-md text-on-surface mb-1">{{ insight.title }}</h4>
                <p class="text-sm text-on-surface-variant">{{ insight.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </UserNav>
</template>

<script>
import { ref, computed, watch } from 'vue'
import { userDashboardService } from '@/services/userService'
import UserNav from '@/components/User_Nav.vue'

export default {
  name: 'AnalyticsView',
  components: { UserNav },
  setup() {
    const loading = ref(true)
    const error = ref(null)
    const timeRange = ref('30d')

    const analytics = ref({
      performance: { total_views: 12580, views_growth: 23, total_favorites: 342, favorites_growth: 18, total_inquiries: 156, inquiries_growth: 12, conversion_rate: 8.5, conversion_growth: 5 },
      views_over_time: [
        { label: 'Mon', views: 1200 }, { label: 'Tue', views: 1800 }, { label: 'Wed', views: 1600 }, { label: 'Thu', views: 2100 }, { label: 'Fri', views: 1900 }, { label: 'Sat', views: 2400 }, { label: 'Sun', views: 1580 }
      ],
      top_listings: [
        { id: 1, title: 'Aero VeloX Pro 2024', views: 3420, inquiries: 45, price: 4200, growth: 28 },
        { id: 2, title: 'TrailBlazer MTB XT', views: 2890, inquiries: 38, price: 2850, growth: 22 }
      ],
      traffic_sources: [
        { name: 'Direct Search', percentage: 35, icon: 'search' }, { name: 'Social Media', percentage: 28, icon: 'share' }, { name: 'Browse', percentage: 22, icon: 'storefront' }, { name: 'External', percentage: 15, icon: 'link' }
      ],
      peak_times: [
        { day: 'Mon', percentage: 65, hours: '6-9 PM' }, { day: 'Tue', percentage: 72, hours: '5-8 PM' }, { day: 'Wed', percentage: 68, hours: '6-9 PM' }, { day: 'Thu', percentage: 85, hours: '4-8 PM' }, { day: 'Fri', percentage: 78, hours: '3-7 PM' }, { day: 'Sat', percentage: 92, hours: '10-2 PM' }, { day: 'Sun', percentage: 88, hours: '11-3 PM' }
      ],
      price_ranges: [
        { label: '$0-$500', count: 18 }, { label: '$501-$1K', count: 24 }, { label: '$1K-$2K', count: 32 }, { label: '$2K-$5K', count: 28 }, { label: '$5K+', count: 8 }
      ],
      insights: [
        { id: 1, type: 'success', title: 'High-Performing Price Range', description: 'Listings in $1,001-$2,000 get 40% more views.' },
        { id: 2, type: 'tip', title: 'Best Time to Post', description: 'Thursday & Saturday have highest engagement.' }
      ]
    })

    const maxViews = computed(() => Math.max(...analytics.value.views_over_time.map(p => p.views), 1))
    const maxPriceCount = computed(() => Math.max(...analytics.value.price_ranges.map(r => r.count), 1))

    const loadAnalytics = async () => {
      try { loading.value = true; error.value = null } catch (err) { error.value = 'Failed to load.' } finally { loading.value = false }
    }

    const formatCurrency = (v) => (v || 0).toLocaleString('en-US')
    const exportData = () => console.log('Export')

    watch(timeRange, () => loadAnalytics())
    loadAnalytics()

    return { loading, error, timeRange, analytics, maxViews, maxPriceCount, formatCurrency, exportData, loadAnalytics }
  }
}
</script>