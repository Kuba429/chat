<script lang="ts">
	import type { message } from "../types";

	export let messages: message[];
</script>

<div>
	{#each messages as message}
		<p>{message.Data}</p>
	{/each}
</div>
