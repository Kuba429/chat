<script lang="ts">
	import { createEventDispatcher, getContext, onMount } from "svelte";
	import { v4 } from "uuid";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	import { getUsername, setUsername } from "../username";
	import Modal from "./Modal.svelte";

	const dispatch = createEventDispatcher();

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
		dispatch("close");
	};
	let inputHeight = "";
	let input: HTMLInputElement;
	onMount(() => {
		inputHeight = getComputedStyle(input).height;
	});
</script>

<Modal close={() => dispatch("close")}>
	<form on:submit={handleSubmit}>
		<input type="text" bind:value={username} bind:this={input} />
		<button style="--input-height: {inputHeight}" type="submit">Set</button>
	</form>
</Modal>

<style>
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
