<script lang="ts">
	const ws = new WebSocket("ws://localhost:3000/ws");
	const id = crypto.randomUUID();
	ws.onopen = () => console.log("conn open");
	ws.onmessage = (m) => console.log(m);
	const sendJoin = () => {
		ws.send(
			JSON.stringify({
				type: "join",
				data: "",
				room: "test_room",
				senderId: id,
				senderName: "Kuba",
			})
		);
	};
	const sendLeave = () => {
		ws.send(
			JSON.stringify({
				type: "leave",
				data: "",
				room: "test_room",
				senderId: id,
				senderName: "Kuba",
			})
		);
	};
	const sendMessage = () => {
		ws.send(
			JSON.stringify({
				type: "message",
				data: "some message lol",
				room: "test_room",
				senderId: id,
				senderName: "Kuba",
			})
		);
	};
</script>

<main>
	<button on:click={sendJoin}>join</button>
	<button on:click={sendLeave}>leave</button>
	<button on:click={sendMessage}>message</button>
</main>

<style>
</style>
