<template>
  <div>
    <h2>用户注册</h2>
    <input v-model="user.username" placeholder="用户名" />
    <input v-model="user.password" placeholder="密码" type="password" />
    <button @click="handleRegister">注册</button>
    <p v-if="errorMessage">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { register } from '../api';

const user = ref({
  username: '',
  password: ''
});
const errorMessage = ref('');

const handleRegister = async () => {
  try {
    const response = await register(user.value);
    if (response.data.code === 200) {
      alert('注册成功');
    } else {
      errorMessage.value = response.data.message;
    }
  } catch (error) {
    errorMessage.value = '注册失败，请稍后重试';
  }
};
</script>