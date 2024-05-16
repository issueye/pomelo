import { createRouter, createWebHistory } from "vue-router";
import Main from "./views/main/index.vue";
import Root from "./views/Root.vue";
import Login from "./views/Login.vue";
import Todo from "./views/todo/index.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/todo", name: "todo", component: Todo },
    { path: "/login", name: "login", component: Login },
    { path: "/main", name: "main", component: Main },
    { path: "/users", component: () => import("./views/Users.vue") },
    { path: "/setup", component: () => import("./views/Setup.vue") },
  ],
});

export default router;
