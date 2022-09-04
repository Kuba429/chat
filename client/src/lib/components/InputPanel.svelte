<script lang="ts">
	import { getContext } from "svelte";
	import type { Connection } from "../connection";

	const conn: Connection = getContext("conn");
	let inputValue = "";
	const handleClick = () => {
		conn.write(inputValue);
		inputValue = "";
	};
</script>

<div>
	<input bind:value={inputValue} type="text" placeholder="input here" />
	<button on:click={() => handleClick()}>send</button>
</div>

<style>
	div {
		height: 2rem;
		width: 100%;
		display: flex;
	}
	div input {
		width: 100%;
		height: 100%;
	}
	div button {
		width: 4rem;
		height: 100%;
	}
</style>
