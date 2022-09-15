export const usernameLSKey = "chat-sv-go-username";
export const defaultUsername = "guest";
export const setUsername = (username: string) => {
	localStorage.setItem(
		usernameLSKey,
		username !== defaultUsername ? username : ""
	);
};
export const getUsername = () => {
	return localStorage.getItem(usernameLSKey);
};
