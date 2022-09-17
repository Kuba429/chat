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
		input.select();
	});
</script>

<Modal close={() => dispatch("close")}>
	<form on:submit={handleSubmit}>
		<input
			type="text"
			placeholder="Username"
			bind:value={username}
			bind:this={input}
		/>
		<button style="--input-height: {inputHeight}" type="submit">Set</button>
	</form>
</Modal>

<style>
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
