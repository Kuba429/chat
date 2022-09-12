<script lang="ts">
	import { getContext, onMount } from "svelte";
	import { writable } from "svelte/store";
	import type { Connection } from "../connection";
	import FileInput from "./FileInput.svelte";

	const conn: Connection = getContext("conn");
	let textarea: HTMLTextAreaElement;
	const imageStore = writable("");
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
		conn.write(textarea.value, $imageStore);
		textarea.value = "";
		imageStore.set("");
		handleInput();
	};
	// set button height to textarea's original height (when textarea has only 1 row)
	onMount(() => {
		textareaOgHeight = parseInt(getComputedStyle(textarea).height);
	});
</script>

<div id="input-panel">
	<div class="wrapper">
		<textarea
			rows="1"
			on:keypress={handleKeypress}
			on:input={handleInput}
			bind:this={textarea}
			placeholder="Type here..."
		/>
	</div>
	<div style="--height: {textareaOgHeight}px" class="buttons">
		<FileInput {imageStore} />
		<button on:click={send}>
			<i class="fa-solid fa-paper-plane" />
		</button>
	</div>
</div>

<style>
	div#input-panel {
		position: relative;
	}
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
		overflow: hidden;
	}
	textarea {
		transition: ease border 0.2s;
		background-color: var(--secondary);
		border: 1px solid var(--secondary);
	}
	textarea:focus {
		border: 1px solid var(--contrasting-dim);
	}
	.buttons {
		display: flex;
		height: fit-content;
		width: fit-content;
		padding: 0;
		margin: 0;
	}
	:global(.buttons > *) {
		height: var(--height);
		width: var(--height);
		font-size: 1rem;
		background: transparent;
		box-sizing: border-box;
		border: 2px dashed transparent;
		transition: ease border 0.1s;
		border-radius: 50%;
		cursor: pointer;
		color: var(--contrasting);
	}
	:global(.buttons > *:hover) {
		border-color: var(--contrasting-dim);
	}
	:global(.buttons *) {
		cursor: pointer;
	}
</style>
