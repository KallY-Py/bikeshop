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
          <h1 class="font-headline-lg text-headline-lg-mobile md:text-headline-lg text-on-surface">Analytics</h1>
          <p class="font-body-md text-body-md text-on-surface-variant mt-1">Deep insights into your listing performance and market trends.</p>
        </div>
        <div class="flex gap-3">
          <select 
            v-model="timeRange"
            class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary transition-colors"
          >
            <option value="7d">Last 7 days</option>
            <option value="30d">Last 30 days</option>
            <option value="90d">Last 90 days</option>
            <option value="1y">This year</option>
          </select>
          <button 
            @click="exportData"
            class="bg-surface-container-low border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface hover:bg-surface-container-high transition-colors flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-lg">download</span>
            <span class="hidden sm:inline">Export</span>
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center h-64">
        <div class="text-center">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4"></div>
          <p class="text-on-surface-variant">Analyzing your data...</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-error/10 border border-error/20 rounded-xl p-8 text-center">
        <span class="material-symbols-outlined text-error text-4xl mb-4">error</span>
        <p class="text-error font-label-md mb-4">{{ error }}</p>
        <button 
          @click="loadAnalytics"
          class="bg-surface-variant hover:bg-surface-bright text-on-surface px-6 py-2 rounded-lg transition-colors"
        >
          Retry
        </button>
      </div>

      <!-- Analytics Content -->
      <div v-else>
        <!-- Performance Overview Cards -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-gutter">
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
            <div class="flex items-center gap-2 mb-2">
              <span class="material-symbols-outlined text-primary">visibility</span>
              <span class="font-label-sm text-label-sm text-on-surface-variant">Total Views</span>
            </div>
            <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_views }}</div>
            <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.views_growth }}%</div>
          </div>
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
            <div class="flex items-center gap-2 mb-2">
              <span class="material-symbols-outlined text-primary">favorite</span>
              <span class="font-label-sm text-label-sm text-on-surface-variant">Favorites</span>
            </div>
            <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_favorites }}</div>
            <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.favorites_growth }}%</div>
          </div>
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
            <div class="flex items-center gap-2 mb-2">
              <span class="material-symbols-outlined text-primary">chat</span>
              <span class="font-label-sm text-label-sm text-on-surface-variant">Inquiries</span>
            </div>
            <div class="font-display text-3xl text-on-surface">{{ analytics.performance.total_inquiries }}</div>
            <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.inquiries_growth }}%</div>
          </div>
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-4 shadow-sm">
            <div class="flex items-center gap-2 mb-2">
              <span class="material-symbols-outlined text-primary">trending_up</span>
              <span class="font-label-sm text-label-sm text-on-surface-variant">Conversion Rate</span>
            </div>
            <div class="font-display text-3xl text-on-surface">{{ analytics.performance.conversion_rate }}%</div>
            <div class="text-xs text-secondary-container mt-1">+{{ analytics.performance.conversion_growth }}%</div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-gutter mb-gutter">
          <!-- Views Over Time Chart -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Views & Inquiries Over Time</h3>
            <div class="h-64 flex items-end gap-2">
              <div 
                v-for="(point, index) in analytics.views_over_time" 
                :key="index"
                class="flex-1 bg-primary/20 hover:bg-primary/40 rounded-t transition-all cursor-pointer relative group"
                :style="{ height: `${(point.views / maxViews) * 100}%` }"
              >
                <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-surface-container-high text-on-surface text-xs px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                  {{ point.views }} views
                </div>
              </div>
            </div>
            <div class="flex justify-between mt-3 text-xs text-on-surface-variant">
              <span v-for="(point, index) in analytics.views_over_time" :key="index">{{ point.label }}</span>
            </div>
          </div>

          <!-- Top Performing Listings -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Top Performing Listings</h3>
            <div class="space-y-4">
              <div 
                v-for="(listing, index) in analytics.top_listings" 
                :key="listing.id"
                class="flex items-center gap-4 p-3 rounded-lg hover:bg-surface-container-high transition-colors"
              >
                <div class="w-8 h-8 rounded-full bg-primary/10 flex items-center justify-center text-primary font-bold text-sm">
                  #{{ index + 1 }}
                </div>
                <div class="w-12 h-12 rounded bg-surface overflow-hidden shrink-0">
                  <img :src="listing.image" :alt="listing.title" class="w-full h-full object-cover" />
                </div>
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

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-gutter mb-gutter">
          <!-- Traffic Sources -->
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

          <!-- Peak Activity Times -->
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

          <!-- Price Range Performance -->
          <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
            <h3 class="font-headline-md text-headline-md text-on-surface mb-6">Price Range Performance</h3>
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

        <!-- Insights & Recommendations -->
        <div class="bg-surface-container-low border border-outline-variant rounded-xl p-6 shadow-sm">
          <h3 class="font-headline-md text-headline-md text-on-surface mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-primary">tips_and_updates</span>
            Insights & Recommendations
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div 
              v-for="insight in analytics.insights" 
              :key="insight.id"
              class="p-4 rounded-lg border"
              :class="insight.type === 'success' ? 'border-secondary-container/30 bg-secondary-container/5' : 'border-primary/30 bg-primary/5'"
            >
              <div class="flex items-start gap-3">
                <span 
                  class="material-symbols-outlined text-lg mt-0.5"
                  :class="insight.type === 'success' ? 'text-secondary-container' : 'text-primary'"
                >
                  {{ insight.type === 'success' ? 'check_circle' : 'lightbulb' }}
                </span>
                <div>
                  <h4 class="font-label-md text-label-md text-on-surface mb-1">{{ insight.title }}</h4>
                  <p class="text-sm text-on-surface-variant">{{ insight.description }}</p>
                </div>
              </div>
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

<script>
import { ref, computed, watch } from 'vue'
import { userDashboardService } from '@/services/userService'

export default {
  name: 'AnalyticsView',
  
  setup() {
    // State
    const loading = ref(true)
    const error = ref(null)
    const timeRange = ref('30d')

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
      { name: 'Profile', icon: 'person', path: '/settings' }
    ]

    // Analytics data
    const analytics = ref({
      performance: {
        total_views: 0,
        views_growth: 0,
        total_favorites: 0,
        favorites_growth: 0,
        total_inquiries: 0,
        inquiries_growth: 0,
        conversion_rate: 0,
        conversion_growth: 0
      },
      views_over_time: [],
      top_listings: [],
      traffic_sources: [],
      peak_times: [],
      price_ranges: [],
      insights: []
    })

    // Computed
    const maxViews = computed(() => {
      if (!analytics.value.views_over_time.length) return 1
      return Math.max(...analytics.value.views_over_time.map(p => p.views), 1)
    })

    const maxPriceCount = computed(() => {
      if (!analytics.value.price_ranges.length) return 1
      return Math.max(...analytics.value.price_ranges.map(r => r.count), 1)
    })

    // Methods
    const loadAnalytics = async () => {
      try {
        loading.value = true
        error.value = null
        
        const userId = localStorage.getItem('userId') || '1'
        
        try {
          const response = await userDashboardService.getAnalytics(userId, timeRange.value)
          analytics.value = response.data
        } catch (apiError) {
          console.log('API not available, using demo data')
          // Demo data
          analytics.value = {
            performance: {
              total_views: 12580,
              views_growth: 23,
              total_favorites: 342,
              favorites_growth: 18,
              total_inquiries: 156,
              inquiries_growth: 12,
              conversion_rate: 8.5,
              conversion_growth: 5
            },
            views_over_time: [
              { label: 'Mon', views: 1200 },
              { label: 'Tue', views: 1800 },
              { label: 'Wed', views: 1600 },
              { label: 'Thu', views: 2100 },
              { label: 'Fri', views: 1900 },
              { label: 'Sat', views: 2400 },
              { label: 'Sun', views: 1580 }
            ],
            top_listings: [
              { id: 1, title: 'Aero VeloX Pro 2024', views: 3420, inquiries: 45, price: 4200, growth: 28, image: 'https://via.placeholder.com/48' },
              { id: 2, title: 'TrailBlazer MTB XT', views: 2890, inquiries: 38, price: 2850, growth: 22, image: 'https://via.placeholder.com/48' },
              { id: 3, title: 'Carbon Handlebar Set', views: 1560, inquiries: 25, price: 350, growth: 15, image: 'https://via.placeholder.com/48' },
              { id: 4, title: 'Shimano XTR Groupset', views: 1240, inquiries: 18, price: 1200, growth: 10, image: 'https://via.placeholder.com/48' }
            ],
            traffic_sources: [
              { name: 'Direct Search', percentage: 35, icon: 'search' },
              { name: 'Social Media', percentage: 28, icon: 'share' },
              { name: 'Marketplace Browse', percentage: 22, icon: 'storefront' },
              { name: 'External Links', percentage: 15, icon: 'link' }
            ],
            peak_times: [
              { day: 'Monday', percentage: 65, hours: '6-9 PM' },
              { day: 'Tuesday', percentage: 72, hours: '5-8 PM' },
              { day: 'Wednesday', percentage: 68, hours: '6-9 PM' },
              { day: 'Thursday', percentage: 85, hours: '4-8 PM' },
              { day: 'Friday', percentage: 78, hours: '3-7 PM' },
              { day: 'Saturday', percentage: 92, hours: '10 AM-2 PM' },
              { day: 'Sunday', percentage: 88, hours: '11 AM-3 PM' }
            ],
            price_ranges: [
              { label: '$0 - $500', count: 18 },
              { label: '$501 - $1,000', count: 24 },
              { label: '$1,001 - $2,000', count: 32 },
              { label: '$2,001 - $5,000', count: 28 },
              { label: '$5,001+', count: 8 }
            ],
            insights: [
              {
                id: 1,
                type: 'success',
                title: 'High-Performing Price Range',
                description: 'Your listings in the $1,001-$2,000 range receive 40% more views.'
              },
              {
                id: 2,
                type: 'tip',
                title: 'Best Time to Post',
                description: 'Listings posted on Thursday and Saturday receive the highest engagement.'
              },
              {
                id: 3,
                type: 'success',
                title: 'Image Quality Matters',
                description: 'Listings with 5+ high-quality images get 3x more inquiries.'
              },
              {
                id: 4,
                type: 'tip',
                title: 'Description Optimization',
                description: 'Detailed listings convert 25% better than those with minimal info.'
              }
            ]
          }
        }
        
      } catch (err) {
        console.error('Failed to load analytics:', err)
        error.value = 'Failed to load analytics data. Please try again.'
      } finally {
        loading.value = false
      }
    }

    const formatCurrency = (value) => {
      return (value || 0).toLocaleString('en-US')
    }

    const exportData = () => {
      console.log('Exporting analytics data...')
    }

    // Watch for time range changes
    watch(timeRange, () => {
      loadAnalytics()
    })

    // Load on mount
    loadAnalytics()

    return {
      loading,
      error,
      timeRange,
      navItems,
      mobileNavItems,
      analytics,
      maxViews,
      maxPriceCount,
      formatCurrency,
      exportData,
      loadAnalytics
    }
  }
}
</script>