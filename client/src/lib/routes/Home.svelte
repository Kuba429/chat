<script lang="ts">
	import { navigateTo } from "svelte-router-spa";
	import { getUsername, setUsername } from "../username";
	let roomValue: string;
	let usernameValue = getUsername();
	const handleClick = (e) => {
		setUsername(usernameValue);
		navigateTo("/r/" + roomValue);
	};
</script>

<div>
	<label>Nickname<input bind:value={usernameValue} type="text" /></label>
	<label>Room<input bind:value={roomValue} type="text" /></label>
	<button on:click={handleClick}>Go</button>
</div>

<style>
	div {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
	}
</style>
