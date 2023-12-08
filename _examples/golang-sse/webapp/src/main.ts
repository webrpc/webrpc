import { effect, signal } from "@preact/signals-core";
import { Chat, Message, SubscribeMessagesReturn, WebrpcError } from "./rpc";
import "./style.css";

//Create client
const api = new Chat("http://localhost:4848", fetch);

//Create signal for messages and log
const messages = signal<Message[]>([]);

type Log = { type: "error" | "info" | "warn"; log: string };
const log = signal<Log[]>([]);

//Create message handlers
const onMessage = (message: SubscribeMessagesReturn) => {
  console.log("onMessage()", message);
  messages.value = [...messages.value, message.message];
};

const onError = (error: WebrpcError) => {
  console.error("onError()", error);
  if (error.message == "AbortError") {
    log.value = [
      ...log.value,
      { type: "warn", log: "Connection closed by abort signal" },
    ];
  } else {
    log.value = [...log.value, { type: "error", log: String(error) }];
  }
};

const onOpen = () => {
  console.log("onOpen()");
  log.value = [...log.value, { type: "info", log: "Connected" }];
};

const onClose = () => {
  console.log("onClose()");
  log.value = [...log.value, { type: "warn", log: "Disconnected" }];
};

const username = randomUserName();

const controller = new AbortController();
const abortSignal = controller.signal;

const abort = () => {
  controller.abort();
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
const disconnectButton = document.getElementById(
  "disconnect"
) as HTMLButtonElement;
disconnectButton.addEventListener("click", abort);

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
