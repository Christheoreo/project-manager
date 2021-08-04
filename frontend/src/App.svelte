<script lang="ts">
import { onDestroy, onMount } from 'svelte';

import { Router, Link, Route, navigate } from 'svelte-routing';
import type { Unsubscriber } from 'svelte/store';
import type { IRoute } from './interfaces/IRoute';
import Home from './routes/Home.svelte';
import Login from './routes/Login.svelte';
import Logout from './routes/Logout.svelte';
import NotFound from './routes/NotFound.svelte';
import Register from './routes/Register.svelte';
import Skeleton from './routes/Skeleton.svelte';
import { IsLoggedInStore } from './stores/IsLoggedInStore';

export let url = '';

let isLoggedInSubscriber: Unsubscriber;

const routes: IRoute[] = [
  { path: '/login', component: Login, protected: false, redirectIfLoggedIn: true },
  { path: '/register', component: Register, protected: false, redirectIfLoggedIn: true },
  {
    path: '/',
    component: Skeleton,
    protected: true,
    children: [
      { path: '', component: Home },
      { path: 'logout', component: Logout },
    ],
  },
  { path: '/*', component: NotFound }, // Make sure this is last!
];

const getAllRoutes = (routes: IRoute[]): IRoute[] => {
  let allRoutes: IRoute[] = [];

  routes.forEach((route) => {
    if (!route.children) {
      allRoutes.push(route);
      return;
    }

    allRoutes.push(...getAllRoutes(route.children));
  });

  return allRoutes;
};

/**
 * @todo - if we can create a stotre for window.location.pathname and do a similar check to the below, that would be cool!
 * if we are logged in, and we navigate to login programatically, it doesnt check auth! but if we refresh the page it does.
 */
onMount(() => {
  const hasToken = window.localStorage.getItem('token') != null;
  IsLoggedInStore.set(hasToken);
  isLoggedInSubscriber = IsLoggedInStore.subscribe((isLoggedIn) => {
    url = window.location.pathname;
    // const currentRoute = routes.find((r) => r.path == url);
    const currentRoute = getAllRoutes(routes).find((r) => r.path == url);
    if (!currentRoute) return;

    if (currentRoute.protected && !isLoggedIn) {
      navigate('login', { replace: true });
    }

    if (isLoggedIn && currentRoute.redirectIfLoggedIn) {
      navigate('/', { replace: true });
    }
  });
});

onDestroy(() => {
  isLoggedInSubscriber();
});
</script>

<!-- 
<Router url="{url}">
  {#each routes as route}
    <Route path="{route.path}" component="{route.component}" />
  {/each}
</Router> -->

<Router url="{url}">
  {#each routes as route}
    {#if route.children}
      {#each route.children as child}
        <Route path="{`${route.path}${child.path}`}">
          <svelte:component this="{route.component}">
            <svelte:component this="{child.component}" />
          </svelte:component>
        </Route>
      {/each}
    {:else}
      <Route path="{route.path}" component="{route.component}" />
    {/if}
  {/each}
</Router>
