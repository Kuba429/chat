<script lang="ts">
	export let close;
</script>

<div class="modal-background" on:click={close} />
<div class="modal">
	<slot />
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
</style>
