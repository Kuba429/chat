import { writable } from "svelte/store";

type head = {
	title: string;
};
export const headStore = writable<head>({ title: "Chat" });
