<script lang="ts">
import { onDestroy, onMount } from 'svelte';
import type { Unsubscriber } from 'svelte/store';
import type { IRoute } from './interfaces/IRoute';
import Home from './routes/Home.svelte';
import Login from './routes/Login.svelte';
import Logout from './routes/Logout.svelte';
import NotFound from './routes/NotFound.svelte';
import Register from './routes/Register.svelte';
import Skeleton from './routes/Skeleton.svelte';
import { IsLoggedInStore } from './stores/IsLoggedInStore';

import router from 'page';

const routes: IRoute[] = [
  { path: '/login', component: Login, protected: false, redirectIfLoggedIn: true },
  { path: '/register', component: Register, protected: false, redirectIfLoggedIn: true },
  { path: '/', component: Home, protected: true },
  { path: 'logout', component: Logout, protected: true },
  { path: '/*', component: NotFound }, // Make sure this is last!
];
let page = Login;
let params: object;
let useSkeleton: boolean = false;

router('/login', () => {
  useSkeleton = false;
  page = Login;
});
router('/register', () => {
  if ($IsLoggedInStore) {
    page = Home;
  } else {
    useSkeleton = false;
    page = Register;
  }
});
router('/logout', () => {
  useSkeleton = false;
  page = Logout;
});
router('/', () => {
  useSkeleton = true;
  page = Home;
});

router.start();

onMount(() => {
  const hasToken = window.localStorage.getItem('token') != null;
  IsLoggedInStore.set(hasToken);
});

onDestroy(() => {});
</script>

{#if useSkeleton}
  <Skeleton>
    <svelte:component this="{page}" />
  </Skeleton>
{:else}
  <svelte:component this="{page}" />
{/if}
