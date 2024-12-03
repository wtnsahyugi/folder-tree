<script setup>
import { ref } from 'vue'

const caret = ref('caret')
const nested = ref('nested')

const items = ref([
    {
        name: 'Beverages',
        children: [
            { name: 'Water', children: [] },
            { name: 'Coffee', children: [] },
            {
                name: 'Tea',
                children: [
                    { name: 'Black Tea', children: [] },
                    { name: 'White Tea', children: [] },
                    {
                        name: 'Green Tea',
                        children: [
                            { name: 'Sencha', children: [] },
                            { name: 'Gyokuro', children: [] },
                            { name: 'Matcha', children: [] },
                            { name: 'Pi Lo Chun', children: [] }
                        ]
                    }
                ]
            }
        ]
    }
]);

function expandData(event) {
    const target = event.target
    const liElement = target.closest('li');

    // Toggle visibility of the nested items by switching classes
    const nestedList = liElement.querySelector('ul');
    if (nestedList) {
        nestedList.classList.toggle('active'); // toggle visibility
        target.classList.toggle('caret-down'); // toggle caret direction
    }
}
</script>

<template>
    <ul id="myUL">
        <li><span @click="expandData" :class="caret">Beverages</span>
            <ul :class="nested">
            <li>Water</li>
            <li>Coffee</li>
            <li><span  @click="expandData" :class="caret">Tea</span>
                <ul :class="nested">
                <li>Black Tea</li>
                <li>White Tea</li>
                <li><span  @click="expandData" :class="caret">Green Tea</span>
                    <ul :class="nested">
                    <li>Sencha</li>
                    <li>Gyokuro</li>
                    <li>Matcha</li>
                    <li>Pi Lo Chun</li>
                    </ul>
                </li>
                </ul>
            </li>
            </ul>
        </li>
    </ul>
</template>

<style scoped>
ul, #myUL {
  list-style-type: none;
}

#myUL {
  margin: 0;
  padding: 0;
}

.caret {
  cursor: pointer;
  user-select: none; /* Prevent text selection */
}

.caret::before {
  content: "\25B6";
  color: black;
  display: inline-block;
  margin-right: 6px;
}

/* Rotate the caret/arrow icon when clicked on (using JavaScript) */
.caret-down::before {
  transform: rotate(90deg);
}

/* Hide the nested list */
.nested {
  display: none;
}

/* Show the nested list when the user clicks on the caret/arrow (with JavaScript) */
.active {
  display: block;
}
</style>