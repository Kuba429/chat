<script lang="ts">
	import { getContext, onMount } from "svelte";
	import type { Writable } from "svelte/store";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	import Message from "./Message.svelte";

	const messagesStore: Writable<message[]> = getContext("messagesStore");
	const conn: Connection = getContext("conn");
	onMount(() => {
		conn.ws.onmessage = (m) => {
			messagesStore.set([JSON.parse(m.data), ...$messagesStore]);
		};
	});
</script>

<div>
	{#each $messagesStore as message}
		<Message {message} />
	{/each}
</div>

<style>
	div {
		margin: 10px;
		height: 100%;
		overflow: auto;
		display: flex;
		flex-direction: column-reverse;
		gap: 1rem;
	}
</style>
