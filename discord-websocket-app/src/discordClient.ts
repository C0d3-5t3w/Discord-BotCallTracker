import { Client, GatewayIntentBits } from 'discord.js';
import { WebSocketServer } from './websocketServer';

export class DiscordClient {
    private client: Client;
    private websocketServer: WebSocketServer;
    private botCallCount: number;

    constructor(websocketServer: WebSocketServer) {
        this.websocketServer = websocketServer;
        this.botCallCount = 0;
        this.client = new Client({
            intents: [GatewayIntentBits.Guilds, GatewayIntentBits.GuildMessages, GatewayIntentBits.MessageContent],
        });
    }

    public connect(token: string): void {
        this.client.login(token);
        this.client.on('messageCreate', this.onMessage.bind(this));
    }

    private onMessage(message: any): void {
        if (message.author.bot) {
            this.botCallCount++;
            this.websocketServer.broadcast({ type: 'botCall', count: this.botCallCount });
        }
    }
}