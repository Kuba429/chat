<script lang="ts">
	import { Router } from "svelte-router-spa";
	import { routes } from "./routes";
</script>

<main>
	<Router {routes} />
</main>

<style>
	main {
		height: 100%;
	}
</style>
