<script lang="ts">
	import { createEventDispatcher, getContext, onMount } from "svelte";
	import { v4 } from "uuid";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	import { getUsername, setUsername } from "../username";
	const dispatch = createEventDispatcher();
	const close = () => dispatch("close");

	const conn: Connection = getContext("conn");
	let username = getUsername();
	const handleSubmit = (e: Event) => {
		e.preventDefault();
		setUsername(username);
		conn.ws.send(
			JSON.stringify(<message>{
				Type: "username_update",
				Data: "",
				Room: conn.room,
				SenderName: username,
				Date: new Date().getTime(),
				SenderId: conn.id,
				Id: v4(),
			})
		);
		close();
	};
	let inputHeight = "";
	let input: HTMLInputElement;
	onMount(() => {
		inputHeight = getComputedStyle(input).height;
	});
</script>

<div class="modal-background" on:click={close} />
<div class="modal">
	<form on:submit={handleSubmit}>
		<input type="text" bind:value={username} bind:this={input} />
		<button style="--input-height: {inputHeight}" type="submit">Set</button>
	</form>
</div>

<style>
	.modal-background {
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		position: absolute;
		z-index: 10;
		background-color: rgba(var(--primary-rgb), 0.7);
		backdrop-filter: blur(10px);
		-webkit-backdrop-filter: blur(10px);
	}
	div.modal {
		position: absolute;
		z-index: 20;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
	}
	input {
		padding: 10px;
		font-size: 1.3rem;
		background-color: var(--secondary);
		color: var(--contrasting);
		box-sizing: border-box;
		transition: ease border 0.2s;
		outline: none;
		border: 1px solid transparent;
		border-radius: 5px;
	}
	input:focus {
		border: 1px solid var(--contrasting-dim);
	}
	form {
		display: flex;
		gap: 1rem;
		flex-wrap: wrap;
		justify-content: center;
		align-items: center;
	}
	button {
		transition: ease border 0.1s;
		width: var(--input-height);
		height: var(--input-height);
		border: dashed 2px transparent;
		margin: 0;
		padding: 0;

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
