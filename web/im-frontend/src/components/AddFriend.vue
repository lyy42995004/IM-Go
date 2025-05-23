<template>
  <div>
    <h2>添加好友</h2>
    <input v-model="friendRequest.uuid" placeholder="用户 UUID" />
    <input v-model="friendRequest.friendUsername" placeholder="好友用户名" />
    <button @click="handleAddFriend">添加好友</button>
    <p v-if="errorMessage">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { addFriend } from '../api';

const friendRequest = ref({
  uuid: '',
  friendUsername: ''
});
const errorMessage = ref('');

const handleAddFriend = async () => {
  try {
    const response = await addFriend(friendRequest.value);
    if (response.data.code === 200) {
      alert('添加好友成功');
    } else {
      errorMessage.value = response.data.message;
    }
  } catch (error) {
    errorMessage.value = '添加好友失败，请稍后重试';
  }
};
</script>