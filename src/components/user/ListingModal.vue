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
        <!-- Title -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Title *</label>
          <input v-model="form.title" type="text" required placeholder="e.g., Trek Madone SLR 9 2024" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        
        <!-- Category -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Category *</label>
          <select v-model="form.category_id" required class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary">
            <option value="">Select category</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }} ({{ cat.type }})</option>
          </select>
        </div>
        
        <!-- Price -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Price ($) *</label>
          <input v-model.number="form.price" type="number" required min="0" step="0.01" placeholder="0.00" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        
        <!-- Condition -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Condition *</label>
          <select v-model="form.condition_type" required class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary">
            <option value="used">Used</option>
            <option value="brand_new">Brand New</option>
            <option value="refurbished">Refurbished</option>
          </select>
        </div>
        
        <!-- Location -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Location</label>
          <input v-model="form.location" type="text" placeholder="City, State" class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary" />
        </div>
        
        <!-- Image Upload (File Only) -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">
            Listing Image <span class="text-on-surface-variant text-xs">(JPEG or PNG only, max 5MB)</span>
          </label>
          
          <!-- Drop Zone -->
          <div 
            @click="$refs.fileInput.click()"
            @dragover.prevent="dragOver = true"
            @dragleave.prevent="dragOver = false"
            @drop.prevent="handleDrop"
            class="border-2 border-dashed rounded-lg p-6 text-center cursor-pointer transition-colors"
            :class="dragOver ? 'border-primary bg-primary/5' : 'border-outline-variant hover:border-primary/50'"
          >
            <input 
              ref="fileInput"
              type="file" 
              accept="image/jpeg,image/png"
              @change="handleFileSelect"
              class="hidden"
            />
            
            <!-- No file selected -->
            <div v-if="!imagePreview">
              <span class="material-symbols-outlined text-4xl text-on-surface-variant mb-2">cloud_upload</span>
              <p class="text-on-surface-variant font-label-md mb-1">Click to upload or drag and drop</p>
              <p class="text-on-surface-variant text-xs">JPEG or PNG only • Max 5MB</p>
            </div>
            
            <!-- Preview -->
            <div v-else class="relative">
              <img :src="imagePreview" class="max-h-48 mx-auto rounded-lg object-cover" />
              <button 
                type="button"
                @click.stop="removeImage"
                class="absolute top-2 right-2 bg-black/60 hover:bg-black/80 text-white rounded-full p-1 transition-colors"
              >
                <span class="material-symbols-outlined text-lg">close</span>
              </button>
              <p class="text-on-surface-variant text-xs mt-2">{{ selectedFile?.name }}</p>
            </div>
          </div>
          
          <!-- Upload progress -->
          <div v-if="uploading" class="mt-2 flex items-center gap-2 text-sm text-primary">
            <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-primary"></div>
            Uploading image...
          </div>
        </div>
        
        <!-- Description -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Description</label>
          <textarea v-model="form.description" rows="4" placeholder="Describe your bike, its condition, and any upgrades..." class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary resize-none"></textarea>
        </div>
        
        <!-- Buttons -->
        <div class="flex gap-3 justify-end pt-4">
          <button type="button" @click="$emit('close')" class="px-6 py-2.5 rounded-lg bg-surface-variant text-on-surface hover:bg-surface-bright transition-colors">Cancel</button>
          <button 
            type="submit" 
            :disabled="saving || uploading" 
            class="px-6 py-2.5 rounded-lg bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md hover:brightness-110 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
          >
            <span v-if="uploading" class="animate-spin rounded-full h-4 w-4 border-b-2 border-black"></span>
            {{ uploading ? 'Uploading...' : saving ? 'Saving...' : (listing?.id ? 'Update Listing' : 'Post Listing') }}
          </button>
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
    saving: { type: Boolean, default: false }
  },
  emits: ['close', 'save'],
  data() {
    return {
      selectedFile: null,
      imagePreview: null,
      uploading: false,
      dragOver: false,
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
      // Show existing image
      if (this.listing.image_url) {
        this.imagePreview = this.listing.image_url
      }
    }
  },
  methods: {
    handleFileSelect(event) {
      const file = event.target.files[0]
      this.processFile(file)
    },
    
    handleDrop(event) {
      this.dragOver = false
      const file = event.dataTransfer.files[0]
      this.processFile(file)
    },
    
    processFile(file) {
      if (!file) return
      
      // Validate file type - JPEG or PNG only
      const validTypes = ['image/jpeg', 'image/png']
      if (!validTypes.includes(file.type)) {
        alert('Please select a JPEG or PNG image only.')
        return
      }
      
      // Validate file size (max 5MB)
      if (file.size > 5 * 1024 * 1024) {
        alert('Image must be less than 5MB.')
        return
      }
      
      this.selectedFile = file
      this.form.image_url = ''
      
      // Show preview
      const reader = new FileReader()
      reader.onload = (e) => {
        this.imagePreview = e.target.result
      }
      reader.readAsDataURL(file)
    },
    
    removeImage() {
      this.selectedFile = null
      this.imagePreview = null
      this.form.image_url = ''
      // Reset file input
      if (this.$refs.fileInput) {
        this.$refs.fileInput.value = ''
      }
    },
    
    async uploadImage() {
      if (!this.selectedFile) return this.form.image_url || null
      
      this.uploading = true
      try {
        const formData = new FormData()
        formData.append('image', this.selectedFile)
        
        const response = await fetch('http://localhost:3000/api/upload', {
          method: 'POST',
          body: formData
        })
        
        if (!response.ok) {
          const errData = await response.json()
          throw new Error(errData.error || 'Upload failed')
        }
        
        const data = await response.json()
        console.log('Image uploaded:', data.image_url)
        return data.image_url
      } catch (err) {
        console.error('Image upload failed:', err)
        alert('Failed to upload image: ' + err.message)
        return null
      } finally {
        this.uploading = false
      }
    },
    
    async handleSubmit() {
      // Upload image first if file selected
      if (this.selectedFile) {
        const imageUrl = await this.uploadImage()
        if (imageUrl) {
          this.form.image_url = imageUrl
        } else {
          return // Don't submit if upload failed
        }
      }
      
      // Require image for new listings
      if (!this.listing?.id && !this.form.image_url) {
        alert('Please upload an image for your listing.')
        return
      }
      
      this.$emit('save', { ...this.form })
    }
  }
}
</script>