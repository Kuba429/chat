<script lang="ts">
	import { Router } from "svelte-router-spa";
	import Header from "./lib/components/Header.svelte";
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
