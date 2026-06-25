<template>
  <div @click="$emit('click')" class="flex gap-3 items-start p-3 rounded-lg hover:bg-surface-container-high transition-colors cursor-pointer border border-transparent hover:border-outline-variant">
    <div class="w-10 h-10 rounded-full bg-surface-container overflow-hidden shrink-0">
      <img :alt="message.sender?.name || 'User'" class="w-full h-full object-cover" :src="message.sender?.avatar || 'https://via.placeholder.com/40'" />
    </div>
    <div class="flex-1 min-w-0">
      <div class="flex justify-between items-baseline mb-1">
        <span class="font-label-md text-label-md text-on-surface">{{ message.sender?.name || 'Unknown' }}</span>
        <span class="text-xs text-on-surface-variant shrink-0 ml-2">{{ message.time }}</span>
      </div>
      <p class="font-body-md text-sm text-on-surface-variant line-clamp-2">{{ message.preview }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MessageItem',
  props: { message: { type: Object, required: true } },
  emits: ['click']
}
</script>