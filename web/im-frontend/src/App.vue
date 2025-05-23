<template>
  <div id="app">
    <UserLogin />
    <UserRegister />
    <UserDetails />
    <SearchUserGroup />
    <FriendList />
    <AddFriend />
  </div>
</template>

<script setup>
import UserLogin from './components/UserLogin.vue';
import UserRegister from './components/UserRegister.vue';
import UserDetails from './components/UserDetails.vue';
import SearchUserGroup from './components/SearchUserGroup.vue';
import FriendList from './components/FriendList.vue';
import AddFriend from './components/AddFriend.vue';
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>