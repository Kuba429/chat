<script lang="ts">
	import type { message } from "../types";

	export let message: message;
</script>

<div>
	<p>{message.SenderName}</p>
	<p>{message.Data}</p>
</div>
