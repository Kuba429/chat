import { writable } from "svelte/store";
import type { message } from "../types";

export const messagesStore = writable<message[]>([]);
