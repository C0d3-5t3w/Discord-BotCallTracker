export interface Message {
    id: string;
    content: string;
    authorId: string;
    timestamp: Date;
}

export interface BotCall {
    callId: string;
    messageId: string;
    timestamp: Date;
}