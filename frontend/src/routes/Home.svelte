<script lang="ts">
import router from 'page';
import Skeleton from './Skeleton.svelte';
const login = () => router.replace('/login');
</script>

<svelte:head>
  <title>Project Manager - Home</title>
</svelte:head>

<!-- <Skeleton> -->
<h1>Home</h1>

<button type="button" on:click="{login}">Login</button>
<!-- </Skeleton> -->
