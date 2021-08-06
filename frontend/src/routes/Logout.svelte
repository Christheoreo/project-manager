<style type="text/css">
.jumbotron {
  padding: 4rem 2rem;
  margin-bottom: 2rem;
  background-color: var(--bs-light);
  border-radius: 0.3rem;
}
</style>

<script lang="ts">
import { onMount } from 'svelte';
import router from 'page';
import { IsLoggedInStore } from '../stores/IsLoggedInStore';

onMount(() => {
  window.localStorage.clear();
  IsLoggedInStore.set(false);
  setTimeout(() => {
    router.replace('/login');
  }, 3000);
});
</script>

<div class="jumbotron">
  <h1 class="display-4">Logging out</h1>
  <p class="lead">
    Currently logging you out.
  </p>
  <p class="lead">
    Have a great day!
  </p>
</div>
