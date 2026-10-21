using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using Interop.Test;

namespace Interop.Tests
{
    // TestApiRestTests ports the generated-client cases of the webrpc REST
    // routes suite (tests/client/rest.go RunRESTTests) to the C# client, run
    // against `webrpc-test -server -rest`. The suite's raw HTTP cases test
    // the server, not the client, and are left out.
    internal static class TestApiRestTests
    {
        public static async Task Run(Check check, HttpClient http, string url)
        {
            var api = new TestApiRestClient(url, http);

            // POST /rpc/items: request fields in the JSON body.
            await check.Run("TestApiRest.CreateItem", async () =>
            {
                var created = await api.CreateItemAsync(new CreateItemRequest { Item = new Item { Id = 7, Name = "alice" } });
                Assert.Equal("alice", created.Item.Name, "item.name");
            });

            // GET /rpc/items/{id}: path param plus optional query param.
            await check.Run("TestApiRest.GetItem", async () =>
            {
                var got = await api.GetItemAsync(new GetItemRequest { Id = 7, Details = true });
                Assert.Equal("alice (details)", got.Item.Name, "item.name");
            });

            // GET /rpc/items: repeated query keys, required scalar, absent optional.
            await check.Run("TestApiRest.ListItems", async () =>
            {
                var list = await api.ListItemsAsync(new ListItemsRequest { Tags = { "a", "b" }, Limit = 5 });
                Assert.True(list.Items.Any(item => item.Name == "q=false tags=[a b] limit=5"), $"echo item not found in [{string.Join(", ", list.Items.Select(i => i.Name))}]");
            });

            // PUT /rpc/items/{id}: path param plus JSON body (normal form).
            await check.Run("TestApiRest.UpdateItem", async () =>
            {
                var updated = await api.UpdateItemAsync(7, "bob");
                Assert.Equal("bob", updated.Name, "name");
            });

            // PATCH /rpc/items/{id}: optional body field present, then absent.
            await check.Run("TestApiRest.PatchItem", async () =>
            {
                var patched = await api.PatchItemAsync(7, "carol");
                Assert.Equal("carol", patched.Name, "name");
            });
            await check.Run("TestApiRest.PatchItem(null)", async () =>
            {
                var patched = await api.PatchItemAsync(7, null);
                Assert.Equal("carol", patched.Name, "name");
            });

            // QUERY /rpc/items/search: the experimental QUERY verb.
            await check.Run("TestApiRest.QueryItems", async () =>
            {
                var found = await api.QueryItemsAsync("carol");
                Assert.Equal(1, found.Count, "count");
                Assert.Equal("carol", found[0].Name, "name");
            });

            // GET /rpc/items/echo/params: every scalar query param type round-trips.
            await check.Run("TestApiRest.EchoParams", async () =>
            {
                var ts = new DateTimeOffset(2026, 9, 3, 12, 30, 15, TimeSpan.Zero);
                var echo = await api.EchoParamsAsync("hello world/&?=", -42, true, 1.5, ts, Status.NOT_AVAILABLE, new List<string> { "x y", "z" }, null);
                Assert.Equal("str=hello world/&?= num=-42 flag=true fl=1.5 ts=2026-09-03T12:30:15Z status=NOT_AVAILABLE tags=[x y z] opt=<nil>", echo, "echo");
            });

            // GET /rpc/items/{id}/children/{childId}: multiple path params.
            await check.Run("TestApiRest.GetChild", async () =>
            {
                Assert.Equal("/items/1/children/2", await api.GetChildAsync(1, 2), "path");
            });

            // GET /rpc/items/{id}/status/latest: literal segments after the path param.
            await check.Run("TestApiRest.GetItemStatus", async () =>
            {
                Assert.Equal("latest status of item 7: carol", await api.GetItemStatusAsync(7), "status");
            });

            // A method without a route keeps the default webrpc dispatch.
            await check.Run("TestApiRest.RpcMethod", async () =>
            {
                Assert.Equal("carol", (await api.RpcMethodAsync(7)).Name, "name");
            });

            // The client calls the canonical (first) route of a method.
            await check.Run("TestApiRest.LookupItem", async () =>
            {
                Assert.Equal("carol", (await api.LookupItemAsync(7)).Name, "name");
            });
            await check.Run("TestApiRest.FindItems", async () =>
            {
                var found = await api.FindItemsAsync("carol");
                Assert.Equal(1, found.Count, "count");
                Assert.Equal("carol", found[0].Name, "name");
            });

            // PUT /rpc/items: declared `+ success 200|201`.
            await check.Run("TestApiRest.UpsertItem", async () =>
            {
                var upserted = await api.UpsertItemAsync(new Item { Id = 10, Name = "zoe" });
                Assert.Equal("zoe", upserted.Name, "name");
            });

            // DELETE /rpc/items/{id}: declared 204, the empty body is the zero result.
            await check.Run("TestApiRest.DeleteItem", async () =>
            {
                await api.DeleteItemAsync(7);
                await Assert.Throws<WebrpcException>(() => api.GetItemAsync(new GetItemRequest { Id = 7 }), "GetItem() after DeleteItem()");
            });

            // A body route without inputs.
            await check.Run("TestApiRest.ResetItems", () => api.ResetItemsAsync());

            // POST /rpc/items/purge: declared `+ success 204`.
            await check.Run("TestApiRest.PurgeItems", () => api.PurgeItemsAsync());

            // An empty required list sends no query keys and decodes as empty.
            await check.Run("TestApiRest.ListItems(empty tags)", async () =>
            {
                var list = await api.ListItemsAsync(new ListItemsRequest { Tags = new List<string>(), Limit = 1 });
                Assert.True(list.Items.Count > 0 && list.Items[0].Name == "q=false tags=[] limit=1", $"unexpected items [{string.Join(", ", list.Items.Select(i => i.Name))}]");
            });
        }
    }
}
