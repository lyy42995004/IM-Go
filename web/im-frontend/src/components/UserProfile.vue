<template>
  <div class="user-profile">
    <h2 class="text-xl font-bold mb-4">用户详情</h2>
    <div class="mb-4">
      <label class="block text-sm font-medium text-gray-700 mb-1">用户 UUID</label>
      <input 
        v-model="uuid" 
        type="text" 
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary/50"
        placeholder="请输入用户 UUID"
      />
    </div>
    <button 
      @click="fetchUserInfo" 
      class="bg-primary hover:bg-primary/90 text-white font-medium py-2 px-4 rounded-md transition-all shadow-md hover:shadow-lg mb-4"
    >
      获取详情
    </button>
    
    <div v-if="userInfo.uuid" class="bg-white p-4 rounded-md shadow-md">
      <div class="flex items-start mb-4">
        <img 
          :src="userInfo.avatar || 'https://picsum.photos/200/200?random=1'" 
          alt="用户头像" 
          class="w-16 h-16 rounded-full object-cover mr-4"
        />
        <div>
          <h3 class="font-bold text-lg">{{ userInfo.nickname || userInfo.username }}</h3>
          <p class="text-sm text-gray-500 mt-1">ID: {{ userInfo.uuid }}</p>
          <p class="text-sm text-gray-500 mt-1">状态: {{ userInfo.status === 1 ? '在线' : '离线' }}</p>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <p class="text-sm text-gray-500">用户名</p>
          <p class="font-medium">{{ userInfo.username }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-500">邮箱</p>
          <p class="font-medium">{{ userInfo.email || '未设置' }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-500">手机号</p>
          <p class="font-medium">{{ userInfo.phone || '未设置' }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-500">注册时间</p>
          <p class="font-medium">{{ formatDate(userInfo.createTime) }}</p>
        </div>
      </div>
    </div>
    
    <p v-if="errorMessage" class="mt-3 text-sm text-red-500">{{ errorMessage }}</p>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { userApi } from '@/api';

const uuid = ref('');
const userInfo = ref({});
const errorMessage = ref('');
const loading = ref(false);

// 格式化日期函数
const formatDate = (timestamp) => {
  if (!timestamp) return '';
  const date = new Date(timestamp);
  return date.toLocaleString();
};

const fetchUserInfo = async () => {
  if (!uuid.value) {
    errorMessage.value = '请输入用户 UUID';
    return;
  }
  
  loading.value = true;
  errorMessage.value = '';
  
  try {
    const response = await userApi.getInfo(uuid.value);
    if (response.code === 200) {
      userInfo.value = response.data;
    } else {
      errorMessage.value = response.message || '获取详情失败';
    }
  } catch (error) {
    errorMessage.value = '请求失败，请稍后重试';
    console.error('获取用户详情错误:', error);
  } finally {
    loading.value = false;
  }
};
</script>