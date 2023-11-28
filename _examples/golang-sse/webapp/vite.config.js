export default {
  server: {
    strictPort: true,
    proxy: {
      "/rpc": {
        target: "http://localhost:4848",
        changeOrigin: true,
      },
    },
  },
};
