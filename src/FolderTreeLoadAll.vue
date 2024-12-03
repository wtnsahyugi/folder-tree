<template>
    <div class="folder-tree">
        <ul>
        <TreeNodeLoadAll
            v-for="item in treeData"
            :key="item.id"
            :node="item"
            :depth=0
        />
        </ul>
    </div>
</template>

<script>
    import { ref, onMounted } from 'vue';
    import TreeNodeLoadAll from './TreeNodeLoadAll.vue';

    export default {
        name: 'FolderTreeLoadAll',
        components: { TreeNodeLoadAll },
        setup() {
            const treeData = ref([]);

            // Fetch root-level folders (parent_id = 0)
            const fetchRootFolders = async () => {
                try {
                    const response = await fetch(`http://localhost:8081/folder/tree?id=0&depth=10`);
                    const data = await response.json();
                    treeData.value = [data.data]; // Adjust to match the API response structure
                } catch (error) {
                    console.error("Failed to fetch root folders:", error);
                }
            };

            onMounted(() => {
                fetchRootFolders();
            });

            return { treeData };
        },
    };
</script>

<style>
  .folder-tree ul {
    list-style: none;
    padding-left: 20px;
  }
</style>
