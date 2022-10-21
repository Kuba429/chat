import { writable } from "svelte/store";

type head = {
	notificationCount: number;
};
export const headStore = writable<head>({ notificationCount: 0 });
