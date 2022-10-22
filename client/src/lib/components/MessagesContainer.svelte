<script lang="ts">
	import { messagesStore } from "../stores/messages";
	import Message from "./Message.svelte";
</script>

<div>
	{#each $messagesStore as message (message.Id)}
		<Message {message} />
	{/each}
	<!-- div below creates free space between header and first message when user scrolls to the top of the container. For some reason margin is not working as it's supposed to in this case. It's dumb but it works 🤷‍♂️ -->
	<div style="height: 10px; padding: 1px; opacity: 0;" />
</div>

<style>
	div {
		padding: 10px;
		height: 100%;
		overflow: auto;
		display: flex;
		flex-direction: column-reverse;
		gap: 1rem;
	}
</style>
