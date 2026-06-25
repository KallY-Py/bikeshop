<template>
  <div class="bg-surface-container-low border border-outline-variant rounded-xl overflow-hidden shadow-sm hover:shadow-md transition-all group">
    <div class="relative h-48 bg-surface overflow-hidden">
      <img :src="listing.image || 'https://via.placeholder.com/400x300'" :alt="listing.title" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300" @error="onImageError"/>
      <div class="absolute top-3 right-3 flex gap-2">
        <span class="px-2 py-1 rounded-full text-xs font-semibold uppercase tracking-wider" :class="statusClass">{{ listing.status }}</span>
      </div>
      <div class="absolute bottom-3 left-3 flex gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
        <button @click="$emit('edit', listing)" class="bg-surface/90 backdrop-blur-sm p-2 rounded-lg hover:bg-surface transition-colors">
          <span class="material-symbols-outlined text-on-surface text-lg">edit</span>
        </button>
        <button @click="$emit('delete', listing)" class="bg-surface/90 backdrop-blur-sm p-2 rounded-lg hover:bg-error/20 transition-colors">
          <span class="material-symbols-outlined text-error text-lg">delete</span>
        </button>
      </div>
    </div>
    <div class="p-4">
      <h3 class="font-headline-md text-headline-md text-on-surface mb-2 line-clamp-1">{{ listing.title }}</h3>
      <p class="text-sm text-on-surface-variant mb-3 line-clamp-2">{{ listing.description || 'No description' }}</p>
      <div class="flex justify-between items-center">
        <span class="font-display text-2xl text-primary">${{ formatPrice(listing.price) }}</span>
        <div class="flex items-center gap-1 text-on-surface-variant text-sm">
          <span class="material-symbols-outlined text-lg">visibility</span>
          <span>{{ listing.views || 0 }}</span>
        </div>
      </div>
      <div class="mt-3 pt-3 border-t border-outline-variant flex justify-between items-center text-sm text-on-surface-variant">
        <span>{{ listing.category || 'Uncategorized' }}</span>
        <span>{{ formatDate(listing.created_at) }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ListingCard',
  props: {
    listing: { type: Object, required: true }
  },
  emits: ['edit', 'delete', 'view'],
  computed: {
    statusClass() {
      const s = this.listing.status
      if (s === 'active' || s === 'approved') return 'bg-secondary-container/20 text-secondary-container'
      if (s === 'sold') return 'bg-tertiary/20 text-tertiary'
      if (s === 'rejected') return 'bg-error/20 text-error'
      return 'bg-primary/20 text-primary'
    }
  },
  methods: {
    formatPrice(price) {
      return (price || 0).toLocaleString()
    },
    formatDate(d) {
      if (!d) return ''
      return new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
    },
    onImageError(e) {
      e.target.src = 'https://via.placeholder.com/400x300?text=No+Image'
    }
  }
}
</script>