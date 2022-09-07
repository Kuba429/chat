import type { message } from "./types";
import { getUsername } from "./username";
import { v4 } from "uuid";
// crypto.randomUUID threw on my phone's browser so i used uuid package
export class Connection {
	ws: WebSocket;
	id: string;
	room: string;
	constructor(room: string) {
		this.ws = new WebSocket(import.meta.env.VITE_WS_SERVER);
		this.id = v4();
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
	write(message: string) {
		if (message.length < 1) return;
		this.ws.send(
			JSON.stringify(<message>{
				Type: "message",
				Data: message,
				Room: this.room,
				SenderId: this.id,
				SenderName: getUsername() || "guest",
				Date: new Date().getTime(),
			})
		);
	}
}
