// fields are capitalized because go can only marshal/unmarshal structs with exported fields (capitalized = exported in go)
export type message = {
	Type: "join" | "message" | "leave";
	Data: string;
	Room: string;
	SenderName: string;
	SenderId: string;
};
