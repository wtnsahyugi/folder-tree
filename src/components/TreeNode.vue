<template>
  <li :style="{ paddingLeft: `${depth * 20}px` }">
    <div class="node" @click="toggle">
      <span v-if="is_folder" class="folder-icon">
        <i :class="isOpen ? 'icon-folder-open' : 'icon-folder-closed'"></i>
      </span>
      <span v-else class="file-icon">
        <i class="icon-file"></i>
      </span>
      {{ node.name }}
    </div>
    <ul v-if="isOpen">
      <TreeNode
        v-for="child in node.child || []"
        :key="child.id"
        :node="child"
        :depth="depth + 1"
        @load-children="handleChildLoad"
      />
    </ul>
  </li>
</template>

<script setup>
  import { ref, inject, watch } from 'vue';

  const props = defineProps({
    node: {
      type: Object,
      required: true,
    },
    depth: {
      type: Number,
      default: 0,
    },
  });

  const folderParent = inject('openFolder')

  const is_folder = props.node.is_folder;
  const emit = defineEmits();

  const isOpen = ref(false);
  const selectedFolder = ref(null);

  const toggle = () => {
      console.log(props.node.id, props.node)
      if (!isOpen.value && props.node.is_folder) {
          // Load children only when opening a folder for the first time
          emit("load-children", props.node.id, props.node);
      }

      selectedFolder.value = props.node;
      emit('update:selected-folder', selectedFolder)

      isOpen.value = !isOpen.value;
  };


  const handleChildLoad = (childId, childNode) => {
      emit("load-children", childId, childNode);
  };

  watch(folderParent, (newFolder) => {
      if (newFolder && newFolder.value.id === props.node.id) {
          isOpen.value = true;
      }
  },
  { deep: true });
</script>

<style>
  .node {
    cursor: pointer;
    padding: 5px;
  }
  .toggle-icon {
    margin-right: 5px;
  }

  .folder-icon .icon-folder-open,
  .folder-icon .icon-folder-closed, .icon-file {
    width: 16px; /* Consistent width */
    height: 16px; /* Consistent height */
    display: inline-block;
    background-size: contain; /* Scale the image to fit */
    background-repeat: no-repeat; /* Prevent tiling */
  }

  /* Folder icons */
  .folder-icon .icon-folder-open {
    content: url('../assets/folder_open.png'); /* Replace with actual folder open icon */
  }

  .folder-icon .icon-folder-closed {
    content: url('../assets/folder_close.png'); /* Replace with actual folder closed icon */
  }

  /* File icon */
  .file-icon .icon-file {
    content: url('../assets/document.png'); /* Replace with actual file icon */
  }

  /* Child padding dynamically handled via inline style */
  li {
    list-style: none;
  }
</style>
