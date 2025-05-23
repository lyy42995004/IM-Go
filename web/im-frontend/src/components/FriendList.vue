<template>
  <div>
    <h2>好友列表</h2>
    <input v-model="uuid" placeholder="用户 UUID" />
    <button @click="handleGetFriendList">获取好友列表</button>
    <ul v-if="friendList.length > 0">
      <li v-for="friend in friendList" :key="friend.uuid">{{ friend.username }}</li>
    </ul>
    <p v-if="errorMessage">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { getFriendList } from '../api';

const uuid = ref('');
const friendList = ref([]);
const errorMessage = ref('');

const handleGetFriendList = async () => {
  try {
    const response = await getFriendList(uuid.value);
    if (response.data.code === 200) {
      friendList.value = response.data.data;
    } else {
      errorMessage.value = response.data.message;
    }
  } catch (error) {
    errorMessage.value = '获取好友列表失败，请稍后重试';
  }
};
</script>