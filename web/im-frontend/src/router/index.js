import { createRouter, createWebHistory } from 'vue-router';
import AuthLogin from '@/components/AuthLogin.vue';
import AuthRegister from '@/components/AuthRegister.vue';
import UserProfile from '@/components/UserProfile.vue';
import FriendList from '@/components/FriendList.vue';
import AddFriend from '@/components/AddFriend.vue';

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'Login',
    component: AuthLogin
  },
  {
    path: '/register',
    name: 'Register',
    component: AuthRegister
  },
  {
    path: '/profile',
    name: 'Profile',
    component: UserProfile,
    meta: { requiresAuth: true }
  },
  {
    path: '/friends',
    name: 'Friends',
    component: FriendList,
    meta: { requiresAuth: true }
  },
  {
    path: '/add-friend',
    name: 'AddFriend',
    component: AddFriend,
    meta: { requiresAuth: true }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 路由守卫
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !localStorage.getItem('user')) {
    next('/login');
  } else {
    next();
  }
});

export default router;