<template>
    <div class="folder-tree">
        <ul>
        <TreeNode
            v-for="item in treeData"
            :key="item.id"
            :node="item"
            :depth=0
            @load-children="loadChildren"
        />
        </ul>
    </div>
</template>

<script>
    import { ref, onMounted } from 'vue';
    import TreeNode from './TreeNode.vue';

    export default {
        name: 'FolderTree',
        components: { TreeNode },
        setup(props, { emit }) {
            const treeData = ref([]);
            const selectedFolder = ref(null);

            // Fetch root-level folders (parent_id = 0)
            const fetchRootFolders = async () => {
                try {
                    const response = await fetch(`http://localhost:8081/folder/tree?id=0&depth=2`);
                    const data = await response.json();
                    treeData.value = [data.data]; // Adjust to match the API response structure
                } catch (error) {
                    console.error("Failed to fetch root folders:", error);
                }
            };

            // Load children dynamically for a specific parent_id
            const loadChildren = async (parentId, parentNode) => {
                try {
                    const response = await fetch(`http://localhost:8081/folder/tree?id=${parentId}&depth=2`);
                    const respData = await response.json();
                    parentNode.child = respData.data.child; // Populate the child nodes
                    selectedFolder.value = parentNode;
                    emit('update:selected-folder', selectedFolder)
                } catch (error) {
                    console.error(`Failed to load children for parent ID ${parentId}:`, error);
                }
            };

            onMounted(() => {
                fetchRootFolders();
            });

            return { treeData, loadChildren };
        },
    };
</script>

<style>
  .folder-tree ul {
    list-style: none;
    padding-left: 20px;
  }
</style>
