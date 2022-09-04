<script lang="ts">
	import type { message } from "../types";

	export let message: message;
</script>

<div>
	<p>{message.SenderName}</p>
	<p>{message.Data}</p>
</div>

<style>
	div {
		background-color: #353535;
		border-radius: 5px;
		padding: 10px;
	}
</style>
