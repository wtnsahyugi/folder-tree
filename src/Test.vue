<script setup>
import { ref } from 'vue'
const count = ref(0)
const title = ref('title')
const items = ref([{ message: 'Foo' }, { message: 'Bar' }])

function addData() {
  count.value++
  items.value.push({message: "ff"+count.value})
}

</script>

<template>
  <button :id="title" @click="addData">Count is: {{ count }}</button>
  <ul>
    <li v-for="item in items" >
      <span :id="title">{{ item.message }}</span>
    </li>
  </ul>
</template>

<style scoped>
button {
  font-weight: bold;
}

#title {
  color: red;
}
</style>