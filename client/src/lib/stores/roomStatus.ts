import { writable } from "svelte/store";
import type { roomStatus } from "../types";

export const roomStatusStore = writable<roomStatus>({
	room: "",
	length: 0,
	users: [],
});
