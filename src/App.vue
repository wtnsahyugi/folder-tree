<script setup>
// import HelloWorld from './components/HelloWorld.vue'
// import FolderTreeLoadAll from './FolderTreeLoadAll.vue';
import {ref, provide} from 'vue'
import FolderTree from './components/FolderTree.vue';
import FolderContent from './components/FolderContent.vue';

const selectedFolder = ref(null); // Hold reference to the currently selected folder
const contentTrigger = ref(null)

provide('openFolder', contentTrigger)

const updateSelectedFolder = (folder, triggerBy) => {
    selectedFolder.value = folder; // Update the selected folder
    console.log(selectedFolder.value, folder)
    if (triggerBy === 'content-'+folder.value.id && folder.value.is_folder) {
      contentTrigger.value = folder
    }
};
</script>

<template>
  <div style="display: flex; height: 100vh;">
    <div style="flex: 1;">
      <FolderTree @update:selected-folder-app="updateSelectedFolder" :contentTrigger="contentTrigger"/>
    </div>
    <div style="flex: 2;">
      <FolderContent :selectedFolder="selectedFolder" @update:selected-folder-content="updateSelectedFolder" />
    </div>
  </div>
</template>
