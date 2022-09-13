import type { message } from "./types";
import { getUsername } from "./username";
import { v4 } from "uuid";
import { messagesStore } from "./stores/messages";
import { roomStatusStore } from "./stores/roomStatus";
import { headStore } from "./stores/head";
import { connStatus } from "./stores/connStatus";

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
			connStatus.set(false);
		};
		// this.receive has no access to 'this' when it's assigned directly
		this.ws.onmessage = (m) => this.receive(m);
	}
	receive(messageRaw: MessageEvent<string>) {
		const message: message = JSON.parse(messageRaw.data);
		switch (message.Type) {
			case "message":
				if (message.SenderId === this.id) {
					messagesStore.update((state) => {
						state.find((m) => m.Id === message.Id).IsPending =
							false;
						return state;
					});
				} else {
					messagesStore.update((state) => [message, ...state]);
					notify(message);
				}
				break;
			case "join":
			case "leave":
			case "username_update":
				roomStatusStore.set({
					users: message.Data.split(","),
					room: this.room,
				});
		}
	}
	joinRoom() {
		if (this.ws.readyState === 1) {
			this.ws.send(
				JSON.stringify(<message>{
					Type: "join",
					Data: "",
					Room: this.room,
					SenderId: this.id,
					SenderName: getUsername(),
				})
			);
			connStatus.set(true);
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
	write(messageData: string, image: string = "") {
		if (messageData.length < 1 && image.length < 1) return;
		const message: message = {
			Type: "message",
			Data: messageData,
			Image: image,
			Room: this.room,
			SenderId: this.id,
			SenderName: getUsername(),
			Date: new Date().getTime(),
			Id: v4(),
		};
		messagesStore.update((state) => [
			{ ...message, IsPending: true },
			...state,
		]);
		this.ws.send(JSON.stringify(message));
	}
}

const notify = (() => {
	if ("Notification" in window) {
		Notification.requestPermission();
		document.addEventListener("visibilitychange", () => {
			if (!document.hidden)
				headStore.update((state) => ({ ...state, title: "Chat" }));
		});
		return (message: message) => {
			if (document.hidden) {
				headStore.update((state) => ({
					...state,
					title: `${message.SenderName}: ${message.Data}`,
				}));
				new Notification(message.SenderName, {
					body: message.Data,
					tag: message.Room,
				});
			}
		};
	}
	// Notifications not supported
	return (_: message) => {
		return;
	};
})();
