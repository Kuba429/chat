<script lang="ts">
	import { onMount } from "svelte";
	import { navigateTo } from "svelte-router-spa";
	import { defaultUsername, getUsername, setUsername } from "../username";
	let roomValue: string;
	let usernameValue = getUsername() || defaultUsername;
	const handleSubmit = (e) => {
		setUsername(usernameValue);
		navigateTo("/r/" + roomValue);
	};
	let onlineUsers = 0;
	onMount(async () => {
		onlineUsers = 4;
		const res = await fetch(
			import.meta.env.VITE_HTTP_SERVER + "/getUserCount"
		);
		onlineUsers = +(await (await res.blob()).text());
	});
</script>

<form on:submit={handleSubmit}>
	<input
		max="16"
		placeholder="Username"
		bind:value={usernameValue}
		type="text"
	/>
	<input placeholder="Room" bind:value={roomValue} type="text" />
	<button type="submit"><i class="fa-solid fa-arrow-right" /></button>
</form>
{#if onlineUsers > 0}
	<span>{onlineUsers} <i class="fa-sharp fa-solid fa-circle green" /></span>
{/if}

<style>
	:global(#app) {
		background: var(--primary);
	}
	span {
		position: absolute;
		bottom: 0;
		right: 0;
		margin: 10px;
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 5px;
	}
	span > i {
		color: var(--green-online);
		font-size: 80%;
	}
	form {
		display: flex;
		gap: 10px;
		flex-wrap: wrap;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
	}
	form * {
		color: var(--contrasting);
		font-size: 1.3rem;
	}
	input {
		padding: 10px;
		background-color: var(--secondary);
		color: var(--contrasting);
		box-sizing: border-box;
		transition: ease border 0.2s;
		outline: none;
		border: 1px solid transparent;
		border-radius: 5px;
	}

	button {
		transition: ease border 0.1s;
		aspect-ratio: 1/1;
		border: dashed 2px transparent;
		margin: 10px;
		padding: 0;
		height: 46px;
		cursor: pointer;
		background-color: var(--secondary);
		color: inherit;
		border-radius: 50%;
	}
	button:hover,
	button:focus {
		border-color: var(--contrasting-dim);
	}
</style>
