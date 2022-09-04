import type { message } from "./types";

export class Connection {
	ws: WebSocket;
	id: string;
	room: string;
	constructor(room: string) {
		this.ws = new WebSocket("ws://localhost:3000/ws"); // store url in .env
		this.id = crypto.randomUUID();
		this.room = room;

		// event handlers
		this.ws.onopen = () => {
			console.log("open");
			this.joinRoom();
		};
		this.ws.onclose = () => {
			console.log("close");
		};
	}
	joinRoom() {
		if (this.ws.readyState === 1) {
			this.ws.send(
				JSON.stringify(<message>{
					Type: "join",
					Data: "",
					Room: this.room,
					SenderId: this.id,
					SenderName: "Kuba",
				})
			);
		} else {
			setTimeout(() => {
				// if the socket isn't connected yet try to join again
				this.joinRoom();
			}, 1000);
		}
	}
	close() {
		this.ws.close();
	}
	write() {
		this.ws.send(
			JSON.stringify(<message>{
				Type: "message",
				Data: "some data lol",
				Room: this.room,
				SenderId: this.id,
				SenderName: "Kuba",
			})
		);
	}
}
