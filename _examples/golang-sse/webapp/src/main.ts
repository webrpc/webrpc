/// <reference lib="dom" />

import { effect, signal } from "@preact/signals-core";
import { Chat, Message } from "./rpc";
import "./style.css";

const api = new Chat("http://localhost:5173", fetch);

const messages = signal<Message[]>([]);

const onMessage = (message: Message) => {
  messages.value = [...messages.value, message];
};

const onError = (error: string) => {
  console.error(error);
};

const onOpen = () => {
  console.log("Connected");
};

api.subscribeMessages({ serverTimeoutSec: 10 }, { onMessage, onError, onOpen });

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

const form = document.querySelector("form") as HTMLFormElement;

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
const authorName = names[randomIndex];

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
