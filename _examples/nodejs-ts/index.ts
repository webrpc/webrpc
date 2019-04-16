import express from "express";
import { createExampleServiceApp } from "./server/server.gen";

const app = express();

const exampleServiceApp = createExampleServiceApp({
  Status: () => {
    return {
      status: false
    };
  },
  Ping: () => ({}),
  GetUser: () => ({
    code: 1,
    user: {
      id: 1,
      USERNAME: "",
      role: ""
    }
  })
});

app.use(exampleServiceApp);
