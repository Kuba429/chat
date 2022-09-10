<script lang="ts">
	import { roomStatusStore } from "../stores/roomStatus";
	import { getUsername } from "../username";
	import UsernameModal from "./UsernameModal.svelte";

	let username = getUsername();
	let showModal = false;
</script>

<header>
	<span on:click={() => (showModal = true)} class="username">{username}</span>
	<div>
		<span>room: {$roomStatusStore.users.length}</span>
	</div>
</header>
{#if showModal}
	<UsernameModal
		on:close={() => {
			username = getUsername();
			showModal = false;
		}}
	/>
{/if}

<style>
	header {
		padding: 10px;
		background-color: var(--secondary);
		display: flex;
		justify-content: space-between;
		align-items: flex-end;
	}
	span.username {
		font-size: 1.5rem;
	}
</style>
