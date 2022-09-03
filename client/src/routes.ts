// .js because typescript complains about importing svelte components for some reason. Doesn't really matter here so i'm leaving it as is
import Home from "./lib/routes/home.svelte";
import Room from "./lib/routes/room.svelte";
const routes = [
	{
		name: "/",
		component: Home,
	},
	{
		name: "/r/:room",
		component: Room,
	},
];
export { routes };
