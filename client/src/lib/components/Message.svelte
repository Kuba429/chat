<script lang="ts">
	import { getContext } from "svelte";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	export let message: message;
	const conn: Connection = getContext("conn");
	const dateRaw = new Date(message.Date);
	const date = `${dateRaw.getHours()}:${dateRaw.getMinutes()}`;
</script>

<div class={message.SenderId === conn.id ? "own" : ""}>
	<p class="username">{message.SenderName} <span>{date}</span></p>
	<p>{message.Data}</p>
</div>

<style>
	div {
		background-color: #353535;
		border-radius: 5px;
		padding: 10px;
		word-break: break-word;
		width: fit-content;
		height: fit-content;
		max-width: 70vw;
	}
	/* message sent by this user */
	div.own {
		align-self: flex-end;
		text-align: right;
		background-color: var(--contrasting-dimm);
		color: var(--primary);
	}

	p {
		padding: 0;
		margin: 0;
	}
	.username {
		font-size: 1.2rem;
		font-weight: 500;
	}
	.username span {
		font-size: 0.6em;
		opacity: 70%;
	}
</style>
