<template>
    <div class="folder-contents">
      <h2>Folder: {{ selectedFolder?.value.name || 'No Folder Selected' }}</h2>
      <ul>
        <li v-for="file in selectedFolder?.value.child" :key="file.id">
          {{ file.name }}
        </li>
      </ul>

    </div>
  </template>

  <script setup>

  const props = defineProps({
    selectedFolder: {
      type: Object,
      required: true
    },
  });

</script>

<style>
    .folder-contents {
        padding: 20px;
    }
</style>