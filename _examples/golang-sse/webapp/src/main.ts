import { effect, signal } from "@preact/signals-core";
import { Chat, Message, SubscribeMessagesReturn } from "./rpc";
import "./style.css";

//Create client
const api = new Chat("http://localhost:5173", fetch);

//Create signal for messages
const messages = signal<Message[]>([]);

//Create message handlers
const onMessage = (message: SubscribeMessagesReturn) => {
  messages.value = [...messages.value, message.message];
};

const onError = (error: unknown) => {
  console.error(error);
};

const onOpen = () => {
  console.log("Connected");
};

const onClose = () => {
  console.log("Conection closed");
};

//Subscribe to messages
api.subscribeMessages(
  { serverTimeoutSec: 10 },
  { onMessage, onError, onOpen, onClose }
);

// Update chatbox
const chatbox = document.querySelector("#chatbox");

effect(() => {
  if (chatbox) {
    chatbox.innerHTML = `
      ${messages.value
        .map(
          (message) => `
          <p class="speechbubble">${message.text}<span class="username">${message.authorName}</span</p>`
        )
        .join("")}
  `;
    chatbox.scrollTo(0, chatbox.scrollHeight);
  }
});

// Send new message on submit
const form = document.querySelector("form") as HTMLFormElement;
const authorName = randomUserName();

function sendMessage(event: Event) {
  const textField = form.elements.namedItem("text") as HTMLInputElement;
  try {
    api.sendMessage({
      authorName,
      text: textField.value,
    });
  } catch (e) {
    console.error(e);
  }
  textField.value = "";
  event.preventDefault();
}

form.addEventListener("submit", sendMessage);

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
