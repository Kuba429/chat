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
