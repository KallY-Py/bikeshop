<template>
  <div class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4">
    <div class="bg-surface-container-high border border-outline-variant rounded-xl p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-6">
        <h2 class="font-headline-lg text-headline-lg text-on-surface">{{ listing?.id ? 'Edit Listing' : 'New Listing' }}</h2>
        <button @click="$emit('close')" class="text-on-surface-variant hover:text-on-surface transition-colors">
          <span class="material-symbols-outlined">close</span>
        </button>
      </div>
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Title *</label>
          <input v-model="form.title" type="text" required placeholder="e.g., Trek Madone SLR 9 2024" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Category *</label>
          <select v-model="form.category_id" required class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary">
            <option value="">Select category</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }} ({{ cat.type }})</option>
          </select>
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Price ($) *</label>
          <input v-model.number="form.price" type="number" required min="0" step="0.01" placeholder="0.00" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Condition *</label>
          <select v-model="form.condition_type" required class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary">
            <option value="used">Used</option>
            <option value="brand_new">Brand New</option>
            <option value="refurbished">Refurbished</option>
          </select>
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Image URL</label>
          <input 
            v-model="form.image_url"
            type="text"
            placeholder="https://example.com/bike-image.jpg"
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary"
          />
          <p class="text-xs text-on-surface-variant mt-1">Paste an image URL for your listing</p>
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Location</label>
          <input v-model="form.location" type="text" placeholder="City, State" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Description</label>
          <textarea v-model="form.description" rows="4" placeholder="Describe your bike, its condition, and any upgrades..." class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary resize-none"></textarea>
        </div>
        <div class="flex gap-3 justify-end pt-4">
          <button type="button" @click="$emit('close')" class="px-6 py-2.5 rounded-lg bg-surface-variant text-on-surface hover:bg-surface-bright transition-colors">Cancel</button>
          <button type="submit" :disabled="saving" class="px-6 py-2.5 rounded-lg bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md hover:brightness-110 transition-all">{{ listing?.id ? 'Update Listing' : 'Post Listing' }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ListingModal',
  props: {
    listing: { type: Object, default: null },
    categories: { type: Array, default: () => [] },
    saving: { type: Boolean, default: false }  // ADD THIS
  },
  emits: ['close', 'save'],
  data() {
    return {
      form: {
        title: '',
        category_id: '',
        price: 0,
        condition_type: 'used',
        location: '',
        description: '',
        image_url: '' 
      }
    }
  },
  mounted() {
    if (this.listing) {
      this.form = {
        title: this.listing.title || '',
        category_id: this.listing.category_id || '',
        price: this.listing.price || 0,
        condition_type: this.listing.condition_type || 'used',
        location: this.listing.location || '',
        description: this.listing.description || '',
        image_url: this.listing.image_url || '' 
      }
    }
  },
  methods: {
    handleSubmit() {
      this.$emit('save', { ...this.form })
    }
  }
}
</script>