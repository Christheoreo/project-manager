import Vue from "vue";
import VueRouter, { NavigationGuardNext, Route, RouteConfig } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import $store from "@/store";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home,
    beforeEnter: (to: Route, from: Route, next: NavigationGuardNext) => {
      if (!$store.state.isLoggedIn) {
        next("/login");
        return;
      }
      next();
    },
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    beforeEnter: (to: Route, from: Route, next: NavigationGuardNext) => {
      if ($store.state.isLoggedIn) {
        next("/");
        return;
      }
      next();
    },
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
