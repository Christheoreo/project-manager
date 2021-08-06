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
  }, 1000);
});
</script>

<div class="jumbotron">
  <h1 class="display-4">Logging out</h1>
  <p class="lead">
    This is a simple hero unit, a simple jumbotron-style component for calling extra attention to featured content or
    information.
  </p>
  <hr class="my-4" />
  <p>It uses utility classes for typography and spacing to space content out within the larger container.</p>
  <a class="btn btn-primary btn-lg" href="#" role="button">Learn more</a>
</div>
