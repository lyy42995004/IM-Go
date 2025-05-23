<template>
  <div>
    <h2>查找用户或群组</h2>
    <input v-model="name" placeholder="名称" />
    <button @click="handleSearch">查找</button>
    <pre v-if="searchResult">{{ searchResult }}</pre>
    <p v-if="errorMessage">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { searchUserOrGroup } from '../api';

const name = ref('');
const searchResult = ref({});
const errorMessage = ref('');

const handleSearch = async () => {
  try {
    const response = await searchUserOrGroup(name.value);
    if (response.data.code === 200) {
      searchResult.value = response.data.data;
    } else {
      errorMessage.value = response.data.message;
    }
  } catch (error) {
    errorMessage.value = '查找失败，请稍后重试';
  }
};
</script>