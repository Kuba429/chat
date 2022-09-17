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
	:global(input[type="text"]) {
		padding: 10px;
		font-size: 1.3rem;
		background-color: var(--secondary);
		color: var(--contrasting);
		box-sizing: border-box;
		transition: ease border 0.2s;
		outline: none;
		border: 1px solid transparent;
		border-radius: 5px;
		max-width: 80vw;
	}
	:global(input[type="text"]:focus) {
		border: 1px solid var(--contrasting-dim);
	}
</style>
