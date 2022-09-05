<script lang="ts">
	import { getContext } from "svelte";
	import type { Connection } from "../connection";

	const conn: Connection = getContext("conn");
	let inputValue = "";
	const send = () => {
		conn.write(inputValue);
		inputValue = "";
	};
	const handleKeypress = (e: KeyboardEvent) => {
		if (e.key === "Enter" && !e.shiftKey) {
			e.preventDefault();
			send();
		}
	};
</script>

<div>
	<div class="wrapper">
		<textarea
			rows="1"
			on:keypress={handleKeypress}
			bind:value={inputValue}
			placeholder="input here"
		/>
	</div>
	<button on:click={send}>send</button>
</div>

<style>
	div {
		margin: 10px;
		height: fit-content;
		display: flex;
	}
	.wrapper {
		width: 100%;
		height: fit-content;
		margin: 0;
		background-color: white;
		border-radius: 5px;
	}
	.wrapper textarea {
		font-size: 1.3rem;
		width: 100%;
		margin: 0;
		padding: 10px;
		background: transparent;
		outline: none;
		border: none;
		resize: none;
		border-radius: inherit;
	}
	div button {
		width: 4rem;
		height: 100%;
	}
</style>
