<script setup lang="ts">
import Home from './views/home/index.vue';
// import { Icon } from '@iconify/vue';
// import DarkMode from './components/DarkMode.vue';
// import { WindowMinimise, Quit, WindowToggleMaximise } from "../wailsjs/runtime/runtime";
// import { ref } from 'vue';
// import { Auth } from './store/store';

// 侧边栏导航
// const menu = [
//   { text: "首页", href: "/home", icon: "material-symbols:home-app-logo" },
//   { text: "用户", href: "/users", icon: "material-symbols:supervisor-account-outline-rounded" },
// ]
// 侧边栏导航激活样式
// const activeClass = 'text-blue-600 dark:text-blue-400'
// 是否非 Mac 平台
// const isNotMac = navigator.userAgent.toUpperCase().indexOf('MAC') < 0;
// 是否最大化
// const isMaximised = ref(false);
</script>

<template>
  <Home />
</template>
