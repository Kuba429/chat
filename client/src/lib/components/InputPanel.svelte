<script lang="ts">
	import { getContext, onMount } from "svelte";
	import { writable } from "svelte/store";
	import type { Connection } from "../connection";
	import FileInput from "./FileInput.svelte";

	const conn: Connection = getContext("conn");
	let textarea: HTMLTextAreaElement;
	const imageStore = writable("");
	let textareaOgHeight = 46; // this value is updated on mount. The first value is just a lucky guess

	// handles height of textarea (when message contains more than 1 line)
	const handleInput = () => {
		// first set height to default value. Otherwise it decreases by only a few pixels per invoke
		textarea.style.height = textareaOgHeight + "px";
		textarea.style.height = textarea.scrollHeight + "px";
		if (textarea.value.length) {
			// allow cycling again after input has been emptied (manually or otherwise); notice that the "send" function invokes this handler
			preventCycling();
		} else {
			allowCycling();
		}
	};

	const sentMessages: string[] = [];
	let shouldCycle = true; // should arrows cycle through sent messages
	let cycleIndex = -1;
	const preventCycling = () => {
		shouldCycle = false;
		cycleIndex = -1;
	};
	const allowCycling = () => {
		shouldCycle = true;
		cycleIndex = -1;
	};

	const handleKeyDown = (e: KeyboardEvent) => {
		// handles whether "enter" should send a message or insert new line
		if (e.key === "Enter" && !e.shiftKey) {
			e.preventDefault();
			send();
		} else if (
			shouldCycle &&
			(e.key === "ArrowUp" || e.key === "ArrowDown")
		) {
			// cycle through messages
			e.preventDefault();
			if (e.key === "ArrowUp" && cycleIndex < sentMessages.length - 1) {
				textarea.value = sentMessages[++cycleIndex];
			} else if (cycleIndex > 0) {
				textarea.value = sentMessages[--cycleIndex];
			}
		} else if (textarea.value.length > 0 && !shouldCycle) {
			// prevent cycling; use arrows to move around text (regular behavior)
			preventCycling();
		}
	};
	const send = () => {
		conn.write(textarea.value, $imageStore);
		sentMessages.unshift(textarea.value);
		textarea.value = "";
		imageStore.set("");
		handleInput();
	};
	// set button height to textarea's original height (when textarea has only 1 row)
	onMount(() => {
		const height = parseInt(getComputedStyle(textarea).height);
		// only change the height when height can be determined
		if (isNaN(height)) return;
		textareaOgHeight = height;
	});
</script>

<div id="input-panel">
	<div class="wrapper">
		<textarea
			rows="1"
			on:keydown={handleKeyDown}
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
