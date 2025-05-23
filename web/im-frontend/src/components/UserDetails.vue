<template>
  <div class="auth-panel">
    <h2 class="text-xl font-bold mb-4">用户注册</h2>
    <div class="mb-3">
      <label class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
      <input 
        v-model="formData.username" 
        type="text" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请输入用户名"
      />
    </div>
    <div class="mb-3">
      <label class="block text-sm font-medium text-gray-700 mb-1">密码</label>
      <input 
        v-model="formData.password" 
        type="password" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请输入密码"
      />
    </div>
    <div class="mb-4">
      <label class="block text-sm font-medium text-gray-700 mb-1">确认密码</label>
      <input 
        v-model="formData.confirmPassword" 
        type="password" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请再次输入密码"
      />
    </div>
    <button 
      @click="handleRegister" 
      class="w-full bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-md transition-all shadow-md hover:shadow-lg"
    >
      注册
    </button>
    <p v-if="errorMessage" class="mt-3 text-sm text-red-500">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { userApi } from '@/api';

const formData = ref({
  username: '',
  password: '',
  confirmPassword: ''
});
const errorMessage = ref('');
const loading = ref(false);

const handleRegister = async () => {
  if (!formData.value.username || !formData.value.password) {
    errorMessage.value = '请输入用户名和密码';
    return;
  }
  
  if (formData.value.password !== formData.value.confirmPassword) {
    errorMessage.value = '两次输入的密码不一致';
    return;
  }
  
  loading.value = true;
  errorMessage.value = '';
  
  try {
    const response = await userApi.register(formData.value);
    if (response.code === 200) {
      errorMessage.value = '注册成功，请登录';
      // 清空表单
      formData.value = {
        username: '',
        password: '',
        confirmPassword: ''
      };
    } else {
      errorMessage.value = response.message || '注册失败';
    }
  } catch (error) {
    errorMessage.value = '注册请求失败，请稍后重试';
    console.error('注册错误:', error);
  } finally {
    loading.value = false;
  }
};
</script>