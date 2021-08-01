<script lang="ts">
  import { navigate } from "svelte-routing";
  const login = () => navigate("/login", { replace: true });
</script>

<svelte:head>
  <title>Project Manager - Home</title>
</svelte:head>

<h1>Home</h1>

<button type="button" on:click={login}>Login</button>
