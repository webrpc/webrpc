import { ExampleService } from "./client.gen";

const exampleService = new ExampleService(
  "http://localhost:3000",
  (init, options) => fetch(init, options)
);

document.addEventListener("DOMContentLoaded", () => {
  const userContainer = document.getElementsByClassName("js-user")[0];
  const loadUserButton = document.getElementsByClassName("js-load-user-btn")[0];
  const userNameText = document.getElementsByClassName("js-username")[0];

  loadUserButton.addEventListener("click", () => {
    exampleService
      .getUser({
        userID: 1,
        header: {}
      })
      .then(({ user }) => {
        userContainer.classList.add("loaded");

        userNameText.textContent = user.USERNAME;
      });
  });
});
