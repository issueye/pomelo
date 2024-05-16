<template>
    <div class="flex flex-col h-screen">
        <div class="flex flex-row bg-[#F2F2F2] border-b border-b-[#D9D9D9]">
            <div class="flex flex-grow items-center" style="--wails-draggable: drag">
                <div class="text-black font-semibold border-red-1-dashed pl-4"> 这是程序标题 </div>
            </div>
            <div class="flex h-10 justify-end flex-0">
                <button
                    class="w-10 h-10 text-xl flex flex-col items-center justify-center hover:bg-[#E9E9E9] dark:hover:bg-[#2D2D2D] hover:text-gray-500">
                    <Icon icon="material-symbols-light:shelf-auto-hide-outline-sharp" @click="Hide()" />
                </button>
                <button
                    class="w-10 h-10 text-xl flex flex-col items-center justify-center hover:bg-[#E9E9E9] dark:hover:bg-[#2D2D2D] hover:text-gray-500">
                    <Icon icon="fluent:minimize-24-regular" @click="WindowMinimise()" />
                </button>
                <button
                    class="w-10 h-10 text-xl flex flex-col items-center justify-center hover:bg-[#E9E9E9] dark:hover:bg-[#2D2D2D]">
                    <Icon icon="fluent:maximize-16-regular" @click="WindowToggleMaximise()" />
                </button>
                <button
                    class="w-10 h-10 text-xl flex flex-col items-center justify-center hover:bg-[#C13124] dark:hover:bg-[#C13124] hover:text-white">
                    <Icon icon="material-symbols:close" @click="Quit()" />
                </button>
            </div>
        </div>
        <div class="flex flex-grow" style="--wails-draggable: none">
            <div class="w-12 bg-[#F9F9F9] border-r border-b-[#D9D9D9]">
                <nav
                    class="flex-none flex flex-col justify-between items-center text-center select-none z-20 bg-gray-50/10 dark:bg-slate-900/80">
                    <div class="mt-5 my-2 flex flex-col gap-6 text-2xl text-gray-500 dark:text-gray-200">
                        <router-link v-for="item in menu" :key="item.text" :to="item.href" v-slot="{ isActive }">
                            <Icon :icon="item.icon" :class="isActive && activeClass" />
                        </router-link>
                    </div>
                    <div class="my-2 flex flex-col gap-4 text-2xl text-gray-500 dark:text-gray-200">
                        <DarkMode />
                        <router-link to="/setup" v-slot="{ isActive }">
                            <Icon icon="material-symbols:menu-rounded" :class="isActive && activeClass" />
                        </router-link>
                    </div>
                </nav>
            </div>

            <div class="w-full">
                <router-view />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { WindowToggleMaximise, WindowMinimise, Quit, Hide } from '../../../wailsjs/runtime/runtime';

// 侧边栏导航
const menu = [
    { text: "首页", href: "/main", icon: "material-symbols:home-app-logo" },
    { text: "用户", href: "/users", icon: "material-symbols:supervisor-account-outline-rounded" },
    { text: "TODO", href: "/todo", icon: "ic:twotone-done-outline" },
]

// 侧边栏导航激活样式
const activeClass = 'text-blue-600 dark:text-blue-400'

</script>