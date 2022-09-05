<script lang="ts">
	import { onDestroy, setContext } from "svelte";
	import { Connection } from "../connection";
	import MessagesContainer from "../components/MessagesContainer.svelte";
	import InputPanel from "../components/InputPanel.svelte";
	import type { message } from "../types";
	import { writable } from "svelte/store";
	import Header from "../components/Header.svelte";
	export let currentRoute;
	const roomId = currentRoute.namedParams.room;
	let conn: Connection = new Connection(roomId);
	onDestroy(() => {
		conn.close();
		conn = null;
	});

	setContext("messagesStore", writable([] as message[]));
	setContext("conn", conn);
</script>

<div>
	<Header />
	<MessagesContainer />
	<InputPanel />
</div>

<style>
	div {
		height: 100%;
		display: flex;
		flex-direction: column;
		/* padding: 10px; */
	}
</style>
