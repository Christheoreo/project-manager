<style global>
/**/
</style>

<script lang="ts">
import { Router, Link, Route } from 'svelte-routing';
import Home from './routes/Home.svelte';
import Login from './routes/Login.svelte';
import Register from './routes/Register.svelte';

export let url = '';
</script>

<Router url="{url}">
    <Route path="login" component="{Login}" />
    <Route path="register" component="{Register}" />
    <Route path="/"><Home /></Route>
</Router>
