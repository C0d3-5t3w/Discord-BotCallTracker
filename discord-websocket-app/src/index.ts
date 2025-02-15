import { DiscordClient } from './discordClient';
import { WebSocketServer } from './websocketServer';

const discordClient = new DiscordClient();
const websocketServer = new WebSocketServer();

discordClient.onMessage((message) => {
    websocketServer.broadcast(message);
});

discordClient.connect();
websocketServer.start();