<script lang="ts">
	import { defaultUsername, getUsername } from "../username";
	import UsernameModal from "./UsernameModal.svelte";
	import ConnectionStatus from "./ConnectionStatus.svelte";

	let username = getUsername();
	let showModal = username.length < 1;
</script>

<header>
	<span on:click={() => (showModal = true)} class="username"
		>{username || defaultUsername}</span
	>
	<div>
		<ConnectionStatus />
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
