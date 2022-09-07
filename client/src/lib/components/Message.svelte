<script lang="ts">
	import { getContext } from "svelte";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	export let message: message;
	const conn: Connection = getContext("conn");
	let dateRaw: Date;
	let date: string;
	let isOwn: boolean;
	$: {
		// idk i don't really understand why i need to do this stuff inside this $ but it doesn't seem to be working without it. It's like the script tag only sees the first message (html sees all messages as it's supposed to though).
		isOwn = message.SenderId === conn.id;
		dateRaw = new Date(message.Date);
		const h = dateRaw.getHours().toString();
		const m = dateRaw.getMinutes().toString();
		// put 0s at the beginning of each single digit number
		date = `${h.length < 2 ? "0" + h : h}:${m.length < 2 ? "0" + m : m}`;
	}
</script>

<div class={isOwn ? "own" : ""}>
	<p class="username">
		<!-- Reversing order here instead of in css because changing display of p messes up span position (it's not perfectly aligned with text outside of span tag) -->
		{#if isOwn}
			<span class="date">{date}</span>
			{message.SenderName}
		{:else}
			{message.SenderName}
			<span class="date">{date}</span>
		{/if}
	</p>
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
	p.username {
		font-size: 1.2rem;
		font-weight: 500;
	}
	p.username span.date {
		font-size: 0.6em;
		opacity: 70%;
	}
</style>
