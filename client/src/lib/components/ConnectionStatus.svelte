<script lang="ts">
	import { roomStatusStore } from "../stores/roomStatus";
	import { connStatus } from "../stores/connStatus";
	// todo: display usernames of active users on hover; maybe modal on click for mobile
</script>

{#if $connStatus}
	<span
		><i class="fa-sharp fa-solid fa-circle green" />
		{$roomStatusStore.users.length}</span
	>
{:else}
	<i class="fa-sharp fa-solid fa-circle red" />
{/if}

<style>
	span {
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
