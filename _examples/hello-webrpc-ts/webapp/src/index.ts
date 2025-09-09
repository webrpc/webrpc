//
// hello-webrpc-ts example program.
// --
// see the `hello-api.ridl` schema where the API is defined
// and the client.gen.ts file is code-generated which we use below..
//
import * as client from "./client.gen";

const api = new client.ExampleService("http://127.0.01:4242", fetch);

async function main() {
  //
  // Ping the API
  //
  console.log("[A] webrpc -- calling Ping() rpc method (expecting true):");

  try {
    const resp = await api.ping();
    console.log("[A]", resp.status);
  } catch (err) {
    console.log("[A]", { err });
  }

  //
  // Get a user from the API
  //
  console.log(
    "[B] webrpc -- calling GetUser() rpc method (expecting User object):"
  );

  try {
    const { user }: client.GetUserReturn = await api.getUser({ userID: 966 });
    console.log("[B]", user);
    console.log("[B] welcome user ID", user.id, "with username", user.USERNAME);
  } catch (err) {
    console.log("[B]", { err });
  }

  //
  // Get an unknown user from the API -- we expect a 404 from the server
  //
  console.log(
    "[C] webrpc -- calling GetUser() rpc method of an unknown user (expecting HTTP 400):"
  );

  try {
    const resp = await api.getUser({ userID: 911 });
    console.log("[C]", resp.user);
  } catch (err) {
    console.log("[C]", { err });
  }

  //
  // Get a user from the API -- different code syntax example
  //
  console.log(
    "[D] webrpc -- calling GetUser() rpc method (expecting User object):"
  );

  await api
    .getUser({ userID: 1337 })
    .then((resp) => {
      const user = resp.user;
      console.log("[D]", user);
      const meta = user.meta;
      console.log("[D]", { meta });
      console.log(
        "[D] welcome user ID",
        user.id,
        "with username",
        user.USERNAME
      );
    })
    .catch((err) => {
      console.log("[D]", { err });
    });

  //
  // Find users from the API
  //
  try {
    const resp = await api.findUsers({ q: "a-z" });
    console.log("[E]", { resp });
  } catch (err) {
    console.log("[E]", { err });
  }
}

main();
