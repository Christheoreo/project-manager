<script lang="ts">
import { navigate } from 'svelte-routing';
import Skeleton from './Skeleton.svelte';
const login = () => navigate('/login', { replace: true });
</script>

<svelte:head>
  <title>Project Manager - Home</title>
</svelte:head>

<!-- <Skeleton> -->
  <h1>Home</h1>

  <button type="button" on:click="{login}">Login</button>
<!-- </Skeleton> -->
