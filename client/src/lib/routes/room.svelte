<script lang="ts">
	import { onMount, onDestroy } from "svelte";
	import { Navigate } from "svelte-router-spa";
	import { Connection } from "../connection";
	export let currentRoute;
	let conn: Connection = null;
	const roomId = currentRoute.namedParams.room;
	onMount(() => {
		conn = new Connection(roomId);
	});
	onDestroy(() => {
		conn.close();
		conn = null;
	});
</script>

<main>
	room {roomId}
	<button on:click={() => conn.write()}>message</button>
	<Navigate to="/">home</Navigate>
</main>
