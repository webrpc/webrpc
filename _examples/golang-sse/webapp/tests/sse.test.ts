import { beforeEach, test, expect, vi } from "vitest";

import {
  WebrpcError,
  Chat,
  SubscribeMessagesReturn,
  Message,
  WebrpcServerPanicError,
  WebrpcRequestFailedError,
  WebrpcStreamLostError,
} from "../src/rpc";

const data = [
  {
    id: 1,
    text: "Hello",
    authorName: "John",
    createdAt: new Date().toDateString(),
  },
  {
    id: 2,
    text: "Hi",
    authorName: "Joe",
    createdAt: new Date().toDateString(),
  },
  {
    id: 3,
    text: "How are you?",
    authorName: "John",
    createdAt: new Date().toDateString(),
  },
] satisfies Message[];

function createMockFetch(
  {
    status,
    body,
    errorAfter,
    closeStream,
  }: {
    status?: number;
    body?: any;
    errorAfter?: number;
    closeStream?: boolean;
  } = {
    status: 200,
    body: undefined,
    errorAfter: undefined,
    closeStream: true,
  }
) {
  return function mockFetch(input: RequestInfo | URL, init?: RequestInit) {
    const stream = new ReadableStream({
      start(controller) {
        data.forEach((item) => {
          const buffer = new TextEncoder().encode(
            JSON.stringify({ message: item }) + "\n"
          );
          controller.enqueue(buffer);
        });

        if (errorAfter) {
          setTimeout(() => {
            controller.error();
          }, errorAfter);
        } else if (closeStream) {
          controller.close();
        }
      },
    });

    body = body || stream;

    return Promise.resolve(
      new Response(body, {
        status,
        headers: {
          "Content-Type": "application/x-ndjson",
        },
      })
    );
  };
}

let onMessage = (msg: any) => {};
let onError = (err: WebrpcError) => {};

beforeEach(() => {
  onMessage = (msg: any) => {};
  onError = (err: WebrpcError) => {};
});

test("call onOpen right before opening stream", async () => {
  const mockFetch = createMockFetch();
  const api = new Chat("", mockFetch);

  const onOpen = vi.fn();
  await api.subscribeMessages(
    { serverTimeoutSec: 10 },
    { onOpen, onMessage, onError }
  );

  expect(onOpen).toHaveBeenCalled();
});

test("call onMessage with correct data", async () => {
  const mockFetch = createMockFetch();
  const api = new Chat("", mockFetch);

  let messages: Message[] = [];
  const onMessage = (msg: SubscribeMessagesReturn) => {
    messages.push(msg.message);
  };
  await api.subscribeMessages({ serverTimeoutSec: 10 }, { onMessage, onError });

  expect(messages).toEqual(data);
});

test("call onClose when stream is done", async () => {
  const mockFetch = createMockFetch();
  const api = new Chat("", mockFetch);

  const onClose = vi.fn();

  await api.subscribeMessages(
    { serverTimeoutSec: 10 },
    { onMessage, onError, onClose }
  );

  expect(onClose).toHaveBeenCalled();
});

test("call onError with WebrpcServerPanicError on server panic", async () => {
  const mockFetch = createMockFetch({
    status: 500,
    body: JSON.stringify({ code: -6 }),
  });

  const api = new Chat("", mockFetch);
  let error: WebrpcError | undefined;

  const onError = (err: WebrpcError) => {
    error = err;
  };

  await api.subscribeMessages({ serverTimeoutSec: 10 }, { onMessage, onError });

  expect(error).toEqual(new WebrpcServerPanicError());
});

test("call onError with WebrpcStreamLostError on stream error", async () => {
  const mockFetch = createMockFetch({ errorAfter: 100 });
  const api = new Chat("", mockFetch);
  let error: WebrpcError | undefined;

  const onError = (err: WebrpcError) => {
    error = err;
  };

  await api.subscribeMessages({ serverTimeoutSec: 10 }, { onMessage, onError });

  expect(error).toEqual(new WebrpcStreamLostError());
});

test("call onError with WebrpcStreamLostError on stream timeout", async () => {
  const mockFetch = createMockFetch({ closeStream: false });
  const api = new Chat("", mockFetch);
  let error: WebrpcError | undefined;

  const onError = (err: WebrpcError) => {
    error = err;
  };

  await api.subscribeMessages({ serverTimeoutSec: 10 }, { onMessage, onError });

  expect(error).toEqual(new WebrpcStreamLostError());
});
