import { effect, signal } from "@preact/signals-core";
import {
  Chat,
  Message,
  SubscribeMessagesReturn,
  WebrpcError,
  WebrpcErrorCodes,
  WebrpcStreamController,
} from "./rpc.gen";
import "./style.css";

// Create API client
const api = new Chat("http://localhost:4848", fetch);

// Create signal for messages and log
const messages = signal<Message[]>([]);
let lastMessageId: number | null = null;
const connectionStatus = signal<
  "connected" | "connecting" | "disconnected" | "finished"
>("connecting");

type Log = { type: "error" | "info" | "warn" | "debug"; log: string };
const log = signal<Log[]>([]);
const appendLog = (logValue: Log) => {
  log.value = [...log.value, logValue];
};

// Create message handlers
const onMessage = (message: SubscribeMessagesReturn) => {
  lastMessageId = message.message.id;

  if (messages.value.length >= 1000) {
    messages.value = [...messages.value.slice(-999), message.message];
  } else {
    messages.value = [...messages.value, message.message];
  }
};

let retryCount = 0;
const maxDelay = 30; // seconds

const onError = (error: WebrpcError, reconnect: () => void) => {
  console.error("onError()", error);

  switch (error.code) {
    case WebrpcErrorCodes.WebrpcStreamFinished:
      appendLog({ type: "info", log: `Chat closed` });
      connectionStatus.value = "finished";
      return;
    case WebrpcErrorCodes.WebrpcClientDisconnected:
    case WebrpcErrorCodes.WebrpcRequestFailed /* THIS IS WRONG */:
      appendLog({ type: "debug", log: `Disconnected` });
      connectionStatus.value = "disconnected";
      return;
  }

  appendLog({ type: "error", log: String(error) });
  connectionStatus.value = "connecting";

  const backoffTime = Math.min(maxDelay, Math.pow(2, retryCount)) * 1000;
  appendLog({
    type: "debug",
    log: `Reconnecting in ${backoffTime / 1000}s`,
  });
  retryCount++;

  setTimeout(reconnect, backoffTime);
};

const onOpen = () => {
  console.log("onOpen()");
  connectionStatus.value = "connected";
  retryCount = 0;
  appendLog({ type: "info", log: "Connected" });
};

const onClose = () => {
  console.log("onClose()");
  connectionStatus.value = "disconnected";
  appendLog({ type: "info", log: "Disconnected" });
};

const username = randomUsername();
const usernameEl = document.querySelector("#username");
if (usernameEl) {
  usernameEl.innerHTML = username;
}

// const controller = new AbortController();
// const abortSignal = controller.signal;

const toggleConnectHandler = () => {
  if (connectionStatus.value == "connected") {
    controller.abort("user disconnected");
  } else if (connectionStatus.value == "connecting") {
    appendLog({ type: "debug", log: "Reconnecting..." });
  } else {
    connect();
    // appendLog({ type: "debug", log: "?" });
  }
};

let controller: WebrpcStreamController;

function connect() {
  appendLog({ type: "debug", log: "Connecting..." });

  // Subscribe to messages
  controller = api.subscribeMessages(
    { username, lastMessageId },
    { onMessage, onError, onOpen, onClose }
  );
}

connect();

// Update chatbox
const chatbox = document.querySelector("#chat");

effect(() => {
  if (chatbox) {
    chatbox.innerHTML = `
      ${messages.value
        .map(
          (message) => `
          <li>
          <div class="name">
            <span>${message.username}</span>
					</div>
          <div class="message">
            ${message.text}
          </div>
          <div class="msg-time">${formatTime(message.createdAt)}</div>
          </li>`
        )
        .join("")}
  `;
    chatbox.scrollTo(0, chatbox.scrollHeight);
  }
});

// Send new message on submit
const form = document.querySelector("form") as HTMLFormElement;
const textField = form.elements.namedItem("text") as HTMLInputElement;

function sendMessage(event: Event) {
  event.preventDefault();
  const msg = textField.value;

  if (connectionStatus.value != "connected" || msg.length == 0) {
    return;
  }

  try {
    api.sendMessage({
      username,
      text: msg,
    });
  } catch (e) {
    console.error(e);
  }
  textField.value = "";
}

form.addEventListener("submit", sendMessage);

// Update log
const logEl = document.querySelector("#log");

effect(() => {
  if (logEl) {
    logEl.innerHTML = `
      ${log.value
        .map((log) => `<div class=${log.type}>${log.log}</div>`)
        .join("")}
  `;
    logEl.scrollTo(0, logEl.scrollHeight);
  }
});

// Abort when disconnect button is clicked
const toggleConnectButton = document.getElementById(
  "toggle-connect"
) as HTMLButtonElement;
toggleConnectButton.addEventListener("click", toggleConnectHandler);

effect(() => {
  switch (connectionStatus.value) {
    case "connected":
      toggleConnectButton.innerText = "Disconnect";
      toggleConnectButton.disabled = false;
      break;
    case "disconnected":
      toggleConnectButton.innerText = "Re-connect";
      toggleConnectButton.disabled = false;
      break;
    case "finished":
      toggleConnectButton.innerText = "Chat closed";
      toggleConnectButton.disabled = true;
      break;

    default:
      toggleConnectButton.innerText = "Reconnecting..";
      toggleConnectButton.disabled = true;
      break;
  }
});

function randomUsername() {
  const names = [
    "Chuck Norris",
    "Mr. Bean",
    "Bugs Bunny",
    "Homer Simpson",
    "SpongeBob",
    "Patrick Star",
    "Mario",
    "Luigi",
    "Yoda",
    "Batman",
    "Iron Man",
    "Spider-Man",
    "Captain America",
    "Wonder Woman",
    "Superman",
    "Mickey Mouse",
    "Donald Duck",
    "Shrek",
    "Darth Vader",
    "Harry Potter",
  ];

  const randomIndex = Math.floor(Math.random() * names.length);
  return names[randomIndex];
}

function formatTime(dateString: string) {
  const date = new Date(dateString);
  const hours =
    date.getHours() > 9 ? `${date.getHours()}` : `0${date.getHours()}`;
  const minutes =
    date.getMinutes() > 9 ? `${date.getMinutes()}` : `0${date.getMinutes()}`;
  const seconds =
    date.getSeconds() > 9 ? `${date.getSeconds()}` : `0${date.getSeconds()}`;

  return `${hours}:${minutes}:${seconds}`;
}
