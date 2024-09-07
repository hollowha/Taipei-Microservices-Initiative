const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
});

// const fs = require("fs");

// module.exports = {
//   devServer: {
//     https: {
//       key: fs.readFileSync("server.key"),
//       cert: fs.readFileSync("server.cert"),
//     },
//   },
// };

// module.exports = {
//   devServer: {
//     allowedHosts: "all", // Allow connections from external hosts
//     host: "0.0.0.0", // Listen on all network interfaces
//     port: 8080, // The port you want to use
//     https: true, // Enable HTTPS on the dev server
//     client: {
//       webSocketURL: {
//         protocol: "wss", // Use secure WebSocket
//         hostname: "your-tunnel-url.loca.lt", // Replace with your tunnel URL
//         port: 443, // Port for HTTPS (443)
//         pathname: "/ws", // WebSocket path
//       },
//     },
//   },
// };

// module.exports = {
//   devServer: {
//     https: false, // Disable HTTPS
//     host: "0.0.0.0",
//     port: 8080,
//     allowedHosts: "all",
//   },
// };

module.exports = {
  devServer: {
    host: "0.0.0.0",
    port: 8080,
    client: {
      webSocketURL: {
        protocol: "wss",
        hostname: "your-ngrok-url.ngrok.io",
        port: 443,
      },
    },
    allowedHosts: "all", // Allow any host
  },
};
