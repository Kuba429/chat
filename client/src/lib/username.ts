const usernameKey = "chat-sv-go-username";
export const setUsername = (username: string) => {
	localStorage.setItem(usernameKey, username !== "guest" ? username : "");
};
export const getUsername = () => {
	return localStorage.getItem(usernameKey) || "guest";
};
