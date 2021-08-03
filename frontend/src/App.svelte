<style global>
/**/
</style>

<script lang="ts">
import { onDestroy, onMount } from 'svelte';

import { Router, Link, Route, navigate } from 'svelte-routing';
import type { Unsubscriber } from 'svelte/store';
import Home from './routes/Home.svelte';
import Login from './routes/Login.svelte';
import Register from './routes/Register.svelte';
import { IsLoggedInStore } from './stores/IsLoggedInStore';

export let url = '';

let isLoggedInSubscriber: Unsubscriber;

const routes: { path: string; component: any; protected: boolean; redirectIfLoggedIn?: boolean }[] = [
  { path: '/login', component: Login, protected: false, redirectIfLoggedIn: true },
  { path: '/register', component: Register, protected: false, redirectIfLoggedIn: true },
  { path: '/', component: Home, protected: true },
];

/**
 * @todo - if we can create a stotre for window.location.pathname and do a similar check to the below, that would be cool!
 * if we are logged in, and we navigate to login programatically, it doesnt check auth! but if we refresh the page it does.
 */
onMount(() => {
  const hasToken = window.localStorage.getItem('token') != null;
  IsLoggedInStore.set(hasToken);
  isLoggedInSubscriber = IsLoggedInStore.subscribe((isLoggedIn) => {
    url = window.location.pathname;
    const currentRoute = routes.find((r) => r.path == url);
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

<Router url="{url}">
  {#each routes as route}
    <Route path="{route.path}" component="{route.component}" />
  {/each}
</Router>
