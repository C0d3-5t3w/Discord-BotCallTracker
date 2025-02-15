import WebSocket from 'ws';
import { BotCall } from './types';

export class WebSocketServer {
    private wss: WebSocket.Server;
    private botCallCount: number;

    constructor(port: number) {
        this.wss = new WebSocket.Server({ port });
        this.botCallCount = 0;

        this.wss.on('connection', (ws) => {
            console.log('Client connected');
            ws.on('message', (message: string) => {
                this.handleMessage(ws, message);
            });
        });
    }

    private handleMessage(ws: WebSocket, message: string) {
        // Handle incoming messages from clients if needed
        console.log(`Received message: ${message}`);
    }

    public incrementBotCall() {
        this.botCallCount++;
        this.broadcast(`Bot calls made: ${this.botCallCount}`);
    }

    public broadcast(data: string) {
        this.wss.clients.forEach((client) => {
            if (client.readyState === WebSocket.OPEN) {
                client.send(data);
            }
        });
    }

    public start() {
        console.log(`WebSocket server started on port ${this.wss.options.port}`);
    }
}