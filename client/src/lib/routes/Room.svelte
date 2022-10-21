<script lang="ts">
	import { onDestroy, setContext } from "svelte";
	import { Connection } from "../connection";
	import MessagesContainer from "../components/MessagesContainer.svelte";
	import InputPanel from "../components/InputPanel.svelte";
	import Header from "../components/Header.svelte";
	import { headStore } from "../stores/head";
	export let currentRoute: any; // TODO maybe add a better type
	const roomId = currentRoute.namedParams.room;
	let conn: Connection = new Connection(roomId);
	onDestroy(() => {
		conn.close();
		conn = null;
	});

	setContext("conn", conn);
	let titleNumber = "";
	$: titleNumber = $headStore.notificationCount
		? `(+${$headStore.notificationCount})`
		: "";
</script>

<svelte:head>
	<title>{titleNumber} Chat</title>
</svelte:head>
<div>
	<Header />
	<MessagesContainer />
	<InputPanel />
</div>

<style>
	div {
		height: 100%;
		background-color: var(--primary);
		display: flex;
		flex-direction: column;
		/* padding: 10px; */
	}
</style>
