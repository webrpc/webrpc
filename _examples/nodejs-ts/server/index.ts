import express from "express";
import { createExampleServiceApp } from "./server.gen";

const app = express();

app.use((req, res, next) => {
  res.setHeader("Access-Control-Allow-Origin", "*");
  res.setHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS");
  res.setHeader("Access-Control-Allow-Headers", "Content-Type");

  if (req.method === "OPTIONS") {
    res.status(200).end();

    return;
  }

  next();
});

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
      USERNAME: "webrpcfan",
      role: "admin"
    }
  })
});

app.use(exampleServiceApp);

app.listen(3000, () => {
  console.log("> Listening on port 3000");
});
