<script lang="ts">
	import type { Writable } from "svelte/store";

	import { toBase64 } from "../toBase64";

	export let imageStore: Writable<string>;
	const handleInput = async (e: Event) => {
		imageStore.set(await toBase64((e.target as HTMLInputElement).files[0]));
	};
	const handleClick = () => {
		imageStore.set("");
	};
</script>

<label>
	<i class="fa-solid fa-image" />
	<input on:input={handleInput} type="file" name="image-input" />
</label>

{#if $imageStore.length > 0}
	<div on:click={handleClick} class="preview">
		<img src={$imageStore} alt="preview" />
	</div>
{/if}

<style>
	.preview {
		position: absolute;
		height: fit-content;
		width: fit-content;
		max-width: 100%;
		max-height: 80vh;
		top: 0;
		right: 0;
		z-index: 1;
		transform: translateY(-100%);
		box-sizing: content-box;
		margin-top: -10px;
		background-color: var(--secondary);
		border: none;
		border-radius: 10px;
	}
	.preview img {
		object-fit: contain;
		max-width: 100%;
		width: 400px;
		max-height: 100%;
	}
	label {
		height: var(--height);
		width: var(--height);
		display: flex;
		justify-content: center;
		align-items: center;
	}
	input[type="file"] {
		/* input doesn't seem to care about 'cursor: pointer' which causes the label to have a spot with wrong cursor */
		position: absolute;
		top: 0;
		left: 0;
		width: 0;
		height: 0;
		opacity: 0%;
	}
</style>
