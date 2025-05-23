<template>
  <div class="auth-panel">
    <h2 class="text-xl font-bold mb-4">用户登录</h2>
    <div class="mb-3">
      <label class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
      <input 
        v-model="formData.username" 
        type="text" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请输入用户名"
      />
    </div>
    <div class="mb-4">
      <label class="block text-sm font-medium text-gray-700 mb-1">密码</label>
      <input 
        v-model="formData.password" 
        type="password" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请输入密码"
      />
    </div>
    <button 
      @click="handleLogin" 
      class="w-full bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-md transition-all shadow-md hover:shadow-lg"
    >
      登录
    </button>
    <p v-if="errorMessage" class="mt-3 text-sm text-red-500">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { userApi } from '@/api';

const formData = ref({
  username: '',
  password: ''
});
const errorMessage = ref('');
const loading = ref(false);

const handleLogin = async () => {
  if (!formData.value.username || !formData.value.password) {
    errorMessage.value = '请输入用户名和密码';
    return;
  }
  
  loading.value = true;
  errorMessage.value = '';
  
  try {
    const response = await userApi.login(formData.value);
    if (response.code === 200) {
      // 存储用户信息和 token
      localStorage.setItem('user', JSON.stringify(response.data));
      localStorage.setItem('token', response.data.token);
      
      // 跳转到首页或刷新页面
      window.location.reload();
    } else {
      errorMessage.value = response.message || '登录失败';
    }
  } catch (error) {
    errorMessage.value = '登录请求失败，请稍后重试';
    console.error('登录错误:', error);
  } finally {
    loading.value = false;
  }
};
</script>