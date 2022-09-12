// fields are capitalized because go can only marshal/unmarshal structs with exported fields (capitalized = exported in go)
export type message = {
	Type: "join" | "message" | "leave" | "username_update";
	Data: string;
	Room: string;
	SenderName: string;
	SenderId: string;
	Date: number;
	Id: string;
	IsPending?: boolean; // only for messages sent by this user
};

export type roomStatus = {
	room: string;
	users: string[];
};
