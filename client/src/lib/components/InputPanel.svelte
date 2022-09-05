<script lang="ts">
	import { getContext, onMount } from "svelte";
	import type { Connection } from "../connection";

	const conn: Connection = getContext("conn");
	let textarea: HTMLTextAreaElement;
	let textareaOgHeight = 46; // this value is updated on mount. The first value is just a lucky guess

	// handles height of textarea (when message includes more than 1 line)
	const handleInput = () => {
		// first set height to default value. Otherwise it decreases by only a few pixels per invoke
		textarea.style.height = textareaOgHeight + "px";
		textarea.style.height = textarea.scrollHeight + "px";
	};
	// handles whether "enter" should send a message or insert new line
	const handleKeypress = (e: KeyboardEvent) => {
		if (e.key === "Enter" && !e.shiftKey) {
			e.preventDefault();
			send();
		}
	};
	const send = () => {
		conn.write(textarea.value);
		textarea.value = "";
		handleInput();
	};
	// set button height to textarea's original height (when textarea has only 1 row)
	onMount(() => {
		textareaOgHeight = parseInt(getComputedStyle(textarea).height);
	});
</script>

<div>
	<div class="wrapper">
		<textarea
			rows="1"
			on:keypress={handleKeypress}
			on:input={handleInput}
			bind:this={textarea}
			placeholder="Type here..."
		/>
	</div>
	<button style="--height: {textareaOgHeight}px" on:click={send}>▶</button>
</div>

<style>
	div {
		margin: 10px;
		height: fit-content;
		display: flex;
		align-items: flex-end;
		gap: 10px;
	}
	.wrapper {
		width: 100%;
		height: fit-content;
		margin: 0;
		border-radius: 5px;
	}
	textarea {
		box-sizing: border-box;
		font-size: 1.3rem;
		color: white;
		width: 100%;
		margin: 0;
		padding: 10px;
		outline: none;
		border: none;
		resize: none;
		border-radius: inherit;
	}
	textarea {
		transition: ease background 0.1s;
		background-color: transparent;
		border: 1px solid var(--secondary);
	}
	textarea:focus {
		background-color: var(--secondary);
	}
	button {
		height: var(--height);
		width: var(--height);
		background: var(--secondary);
		border: 0;
		border-radius: 50%;
		cursor: pointer;
		color: var(--contrasting);
	}
</style>
