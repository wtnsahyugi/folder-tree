<template>
    <div class="folder-contents">
      <h2>Folder: {{ selectedFolder?.value.name || 'No Folder Selected' }}</h2>
      <ul>
        <li v-for="file in selectedFolder?.value.child" :key="file.id">
          <div class="node" @click="clickHandler(file)">
            <span v-if="file.is_folder" class="folder-icon">
              <i class="icon-folder-closed"></i>
            </span>
            <span v-else class="file-icon">
              <i class="icon-file"></i>
            </span>
            {{ file.name }}
          </div>
        </li>
      </ul>

    </div>
  </template>

  <script setup>

  import { ref } from 'vue';
  const props = defineProps({
    selectedFolder: {
      type: Object,
      required: true
    },
  });

  const folderContent = ref(null);
  const emit = defineEmits();
  const loadChildren = async (parentId, parentNode) => {
      try {
          const response = await fetch(`http://localhost:8081/folder-trees?id=${parentId}&depth=2`);
          const respData = await response.json();
          parentNode.child = respData.data.child; // Populate the child nodes
          folderContent.value = parentNode;
          folderContent.contentid = 'content-'+parentId
          emit('update:selected-folder-content', folderContent, 'content-'+parentId)
      } catch (error) {
          console.error(`Failed to load children for parent ID ${parentId}:`, error);
      }
  };
  const clickHandler = (file) => {
    loadChildren(file.id, file);
  };


</script>

<style>
    .folder-contents {
        padding: 20px;
    }
</style>