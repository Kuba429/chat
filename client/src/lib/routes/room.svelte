<script lang="ts">
	import { onMount, onDestroy } from "svelte";
	import { Connection } from "../connection";
	import MessagesContainer from "../components/MessagesContainer.svelte";
	import InputPanel from "../components/InputPanel.svelte";
	import type { message } from "../types";
	export let currentRoute;
	let conn: Connection = null;
	const roomId = currentRoute.namedParams.room;
	let messages: message[] = [];
	onMount(() => {
		conn = new Connection(roomId);
		conn.ws.onmessage = (m) => {
			messages = [...messages, JSON.parse(m.data)];
			console.log(messages);
		};
	});
	onDestroy(() => {
		conn.close();
		conn = null;
	});

	const handleClick = () => {
		conn.write();
	};
</script>

<main>
	<button on:click={() => handleClick()}>message</button>
	<MessagesContainer {messages} />
	<InputPanel />
</main>
