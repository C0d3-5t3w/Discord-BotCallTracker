# Discord WebSocket App

This project is a simple application that connects to a Discord server using a WebSocket server to read messages and display the number of bot calls made.

## Features

- Connects to a Discord server and listens for messages.
- Tracks the number of bot calls made.
- Sets up a WebSocket server to communicate with clients.

## Prerequisites

- Node.js (version 14 or higher)
- TypeScript (version 4 or higher)
- A Discord bot token

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/discord-websocket-app.git
   ```

2. Navigate to the project directory:

   ```
   cd discord-websocket-app
   ```

3. Install the dependencies:

   ```
   npm install
   ```

## Configuration

Before running the application, you need to set up your Discord bot token. Create a `.env` file in the root directory and add the following line:

```
DISCORD_BOT_TOKEN=your_bot_token_here
```

## Usage

To start the application, run the following command:

```
npm start
```

This will initialize the Discord client and the WebSocket server. You can then connect to the WebSocket server using a WebSocket client to receive messages and the number of bot calls.

## WebSocket Server

The WebSocket server will broadcast messages received from the Discord server to all connected clients. It will also send updates on the number of bot calls made.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.