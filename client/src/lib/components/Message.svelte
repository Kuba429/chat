<script lang="ts">
	import { getContext, onMount } from "svelte";
	import type { Connection } from "../connection";
	import type { message } from "../types";
	export let message: message;
	const conn: Connection = getContext("conn");
	let dateRaw: Date;
	let date: string;
	let isOwn: boolean;
	let preElement;

	isOwn = message.SenderId === conn.id;
	dateRaw = new Date(message.Date);
	const h = dateRaw.getHours().toString();
	const m = dateRaw.getMinutes().toString();
	// put 0s at the beginning of each single digit number
	date = `${h.length < 2 ? "0" + h : h}:${m.length < 2 ? "0" + m : m}`;

	onMount(() => {
		// replace urls with anchor tags
		preElement.innerHTML = preElement.innerHTML.replace(
			/((https|http)?:\/\/[^\s]+)/g,
			(url: string) => `<a href="${url}" target="_blank">${url}</a>`
		);
	});
</script>

<div class={`${isOwn ? "own" : ""} ${message?.IsPending ? "dim" : ""}`}>
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
	<pre bind:this={preElement}>{message.Data}</pre>
	{#if message.Image}
		<img src={message.Image} alt="image sent by {message.SenderName}" />
	{/if}
</div>

<style>
	div {
		position: relative;
		background-color: #353535;
		border-radius: 15px;
		border-bottom-left-radius: 0px;
		padding: 10px;

		width: fit-content;
		height: fit-content;
		max-width: 70vw;
	}
	/* message sent by this user */
	div.own {
		border-radius: 15px;
		border-bottom-right-radius: 0px;
		align-self: flex-end;
		text-align: right;
		background-color: var(--contrasting-dim);
		color: var(--primary);
	}
	div.dim {
		opacity: 0.8;
	}
	div::after {
		content: "";
		position: absolute;
		bottom: 0;
		right: unset;
		left: 0;
		transform: translateY(100%);
		background-color: transparent;
		width: 0;
		height: 0;
		border-style: solid;
		border-width: 7px 7px 0 0;
		border-color: var(--secondary) transparent transparent transparent;
	}
	div.own::after {
		right: 0;
		left: unset;
		border-width: 0 7px 7px 0;
		border-color: transparent var(--contrasting-dim) transparent transparent;
	}
	p,
	pre {
		padding: 0;
		margin: 0;
		word-break: break-word;
		white-space: pre-wrap;
	}
	:global(div pre > a) {
		color: inherit;
	}
	p.username {
		font-size: 1.2rem;
		font-weight: 500;
	}
	p.username span.date {
		font-size: 0.6em;
		opacity: 70%;
	}
	img {
		object-fit: contain;
		width: 400px;
		max-width: 100%;
		max-height: 80vh;
	}
</style>
