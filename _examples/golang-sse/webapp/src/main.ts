import { effect, signal } from "@preact/signals-core";
import { Chat, Message, SubscribeMessagesReturn, WebrpcError } from "./rpc.gen";
import "./style.css";

// Create client
const api = new Chat("http://localhost:4848", fetch);

// Create signal for messages and log
const messages = signal<Message[]>([]);
const connectionStatus = signal<
  "connected" | "error" | "aborted" | "disconnected"
>("disconnected");

type Log = { type: "error" | "info" | "warn"; log: string };
const log = signal<Log[]>([]);
const appendLog = (logValue: Log) => {
  log.value = [...log.value, logValue];
};

// Create message handlers
const onMessage = (message: SubscribeMessagesReturn) => {
  console.log("onMessage()", message);
  messages.value = [...messages.value, message.message];
};

let retryCount = 0;
const maxDelay = 30; // seconds

const onError = (error: WebrpcError, reconnect?: () => void) => {
  connectionStatus.value = "error";
  console.error("onError()", error);
  if (error.message == "AbortError") {
    connectionStatus.value = "aborted";
    appendLog({ type: "warn", log: "Connection closed by abort signal" });
    // TODO: Reconnect
  } else {
    appendLog({ type: "error", log: "Lost connection" });
    appendLog({ type: "error", log: String(error) });
    if (reconnect) {
      appendLog({
        type: "warn",
        log: `Attempting reconnect ${retryCount + 1}`,
      });
      retryCount++;
      const backoffTime = Math.min(maxDelay, Math.pow(2, retryCount)) * 1000;
      setTimeout(reconnect, backoffTime);
    }
  }
};

const onOpen = () => {
  console.log("onOpen()");
  connectionStatus.value = "connected";
  appendLog({ type: "info", log: "Connected" });
};

const onClose = () => {
  console.log("onClose()");
  connectionStatus.value = "disconnected";
  appendLog({ type: "info", log: "Disconnected" });
};

const username = randomUserName();

const controller = new AbortController();
const abortSignal = controller.signal;

const toggleConnectHandler = () => {
  if (connectionStatus.value == "connected") {
    controller.abort();
    connectionStatus.value = "aborted";
  } else if (connectionStatus.value == "aborted") {
    // reconnect();
  }
};

// Subscribe to messages
api.subscribeMessages(
  { username },
  { onMessage, onError, onOpen, onClose, signal: abortSignal }
);

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

function sendMessage(event: Event) {
  const textField = form.elements.namedItem("text") as HTMLInputElement;
  try {
    api.sendMessage({
      username,
      text: textField.value,
    });
  } catch (e) {
    console.error(e);
  }
  textField.value = "";
  event.preventDefault();
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
    case "aborted":
      toggleConnectButton.innerText = "Connect";
      toggleConnectButton.disabled = false;
      break;
    case "error":
      toggleConnectButton.innerText = "Connection error";
      toggleConnectButton.disabled = true;
      break;
  }
});

function randomUserName() {
  const names = [
    "Chuck Norris",
    "Mr. Bean",
    "Bugs Bunny",
    "Homer Simpson",
    "SpongeBob",
    "Patrick Star",
    "Pikachu",
    "Mario",
    "Luigi",
    "Yoda",
  ];
  const randomIndex = Math.floor(Math.random() * names.length);
  return names[randomIndex];
}

function formatTime(dateString: string) {
  const date = new Date(dateString);
  const hours = date.getHours();
  const minutes = date.getMinutes();
  const seconds = date.getSeconds();

  return `${hours}:${minutes}:${seconds}`;
}
