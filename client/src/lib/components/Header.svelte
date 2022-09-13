<script lang="ts">
	import { roomStatusStore } from "../stores/roomStatus";
	import { getUsername } from "../username";
	import UsernameModal from "./UsernameModal.svelte";
	import { connStatus } from "../stores/connStatus";

	let username = getUsername();
	let showModal = false;
</script>

<header>
	<span on:click={() => (showModal = true)} class="username">{username}</span>
	<div>
		{#if $connStatus}
			<span
				><i class="fa-sharp fa-solid fa-circle green" />
				{$roomStatusStore.users.length}</span
			>
		{:else}
			<i class="fa-sharp fa-solid fa-circle red" />
		{/if}
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
	div > span {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 5px;
	}
	i {
		font-size: 80%;
	}
	i.green {
		color: var(--green-online);
	}
	i.red {
		color: var(--red-offline);
	}
</style>
