<script lang="ts">
	import { roomStatusStore } from "../stores/roomStatus";
	import { connStatus } from "../stores/connStatus";
</script>

{#if $connStatus}
	<span class="room-status"
		><i class="fa-sharp fa-solid fa-circle green" />
		{$roomStatusStore.users.length}
		<ul>
			<span>online: </span>
			{#each $roomStatusStore.users as user}
				<li>{user}</li>
			{/each}
		</ul>
	</span>
{:else}
	<i class="fa-sharp fa-solid fa-circle red" />
{/if}

<style>
	span.room-status {
		position: relative;
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 5px;
	}
	span.room-status ul {
		display: none;
		position: absolute;
		bottom: -100%;
		right: 0;
		box-sizing: content-box;
		transform: translateY(90%);
		/* padding: 10px; */
		padding: 0;
		border-radius: 10px;
		border-top-right-radius: 0px;
		border: 1px solid var(--contrasting-dim);
		background-color: var(--secondary);
		max-height: 50vh;
		overflow-y: auto;
	}
	span.room-status:hover ul {
		display: block;
	}
	li {
		list-style: none;
		/* margin-top: 10px; */
		padding: 5px 10px;
	}
	ul > span {
		margin: 0;
		padding: 0;
		padding-left: 10px;
		font-size: 80%;
		color: var(--contrasting-dim);
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
