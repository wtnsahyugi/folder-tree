<script setup>
// import HelloWorld from './components/HelloWorld.vue'
// import FolderTreeLoadAll from './FolderTreeLoadAll.vue';
import {ref} from 'vue'
import FolderTree from './components/FolderTree.vue';
import FolderContent from './components/FolderContent.vue';

const selectedFolder = ref(null); // Hold reference to the currently selected folder

const updateSelectedFolder = (folder) => {
    selectedFolder.value = folder; // Update the selected folder
};
</script>

<template>
  <div style="display: flex; height: 100vh;">
    <div style="flex: 1;">
      <FolderTree @update:selected-folder="updateSelectedFolder" />
    </div>
    <div style="flex: 2;">
      <FolderContent :selectedFolder="selectedFolder" @update:selected-folder="updateSelectedFolder" />
    </div>
  </div>
</template>
