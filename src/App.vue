<script setup>
// import HelloWorld from './components/HelloWorld.vue'
// import FolderTreeLoadAll from './FolderTreeLoadAll.vue';
import {ref} from 'vue'
import FolderTree from './FolderTree.vue';
import FolderContent from './FolderContent.vue';

const selectedFolder = ref(null); // Hold reference to the currently selected folder

const updateSelectedFolder = (folder) => {
  console.log('Selected folder:', folder); // Log the selected folder to the console
    selectedFolder.value = folder; // Update the selected folder
};
</script>

<template>
  <div style="display: flex; height: 100vh;">
    <div style="flex: 1;">
      <FolderTree @update:selected-folder="updateSelectedFolder" />
    </div>
    <div style="flex: 2;">
      <FolderContent :selectedFolder="selectedFolder" />
    </div>
  </div>
</template>
