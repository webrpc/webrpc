using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Numerics;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Interop.CSharp;
using Task = System.Threading.Tasks.Task;

namespace Interop.Tests
{
    // CSharpTests runs the generated C# client against the CSharp service
    // (server/main.go), and checks the serialization of the generated types.
    internal static class CSharpTests
    {
        // Fixture shared with server/main.go fixtureTypes().
        private static Types Fixture() => new Types
        {
            B = true,
            U8 = 255,
            U16 = 65535,
            U32 = 4294967295,
            U64 = 18446744073709551615,
            I8 = -128,
            I16 = -32768,
            I32 = -2147483648,
            I64 = -9223372036854775808,
            F32 = 1.5f,
            F64 = -0.000123456789,
            Str = "héllo \"wörld\" <&> ☃ \U0001F600",
            Bytes = new byte[] { 0x00, 0x01, 0xfe, 0xff },
            U8s = new byte[] { 0x01, 0x02 },
            Ts = new DateTimeOffset(2026, 9, 3, 12, 30, 15, TimeSpan.Zero).AddTicks(1234567),
            TsList = { new DateTimeOffset(1970, 1, 1, 0, 0, 0, TimeSpan.Zero), new DateTimeOffset(2026, 12, 31, 23, 59, 59, TimeSpan.Zero) },
            Big = BigInteger.Parse("-123456789012345678901234567890"),
            BigList = { BigInteger.Zero, BigInteger.Parse("340282366920938463463374607431768211456") },
            BigMap = { ["wei"] = BigInteger.Parse("1000000000000000000") },
            Address = "0xabc",
            Wallet = "0xdef",
            Labels = { "a", "b" },
            Kind = Kind.Admin,
            Kinds = { [Kind.User] = 1, [Kind.Admin] = 2 },
            Country = Country.CA,
            Reserved = Reserved.@class,
            ById = { [-1] = "minus one", [42] = "answer" },
            Nested = { [7] = new List<string> { "x", "y" } },
            Anything = new Dictionary<string, object?> { ["l"] = new object?[] { true, null }, ["n"] = 1, ["s"] = "two" },
            OptStr = null,
            OptBig = 42,
            OptTs = new DateTimeOffset(2000, 1, 2, 3, 4, 5, TimeSpan.Zero),
            OptKind = Kind.User,
            OptList = null,
            Renamed = "renamed",
        };

        public static void RunUnit(Check check)
        {
            check.Run("Unit.Serialize(Types)", () =>
            {
                var json = WebrpcJson.Serialize(Fixture());
                Assert.True(json.Contains("\"RENAMED_field\":\"renamed\""), "+ json renames the field: " + json);
                Assert.True(!json.Contains("hidden", StringComparison.OrdinalIgnoreCase), "+ json = - hides the field: " + json);
                Assert.True(!json.Contains("optStr") && !json.Contains("optList"), "null optional fields are omitted: " + json);
                Assert.True(json.Contains("\"ts\":\"2026-09-03T12:30:15.1234567Z\""), "timestamps are ISO 8601 in UTC: " + json);
                Assert.True(json.Contains("\"tsList\":[\"1970-01-01T00:00:00Z\",\"2026-12-31T23:59:59Z\"]"), "whole seconds have no fraction: " + json);
                Assert.True(json.Contains("\"big\":\"-123456789012345678901234567890\""), "bigints are strings: " + json);
                Assert.True(json.Contains("\"bigMap\":{\"wei\":\"1000000000000000000\"}"), "bigint map values are strings: " + json);
                Assert.True(json.Contains("\"optBig\":\"42\""), "optional bigints are strings: " + json);
                Assert.True(json.Contains("\"kind\":\"ADMIN\"") && json.Contains("\"kinds\":{\"USER\":1,\"ADMIN\":2}"), "integer enums use their + json value: " + json);
                Assert.True(json.Contains("\"country\":\"Canada\""), "string enums use their value: " + json);
                Assert.True(json.Contains("\"byId\":{\"-1\":\"minus one\",\"42\":\"answer\"}"), "integer map keys are strings: " + json);
                Assert.True(json.Contains("\"bytes\":\"AAH+/w==\""), "[]byte is base64: " + json);
                Assert.True(json.Contains("\"u8s\":\"AQI=\""), "[]uint8 is base64: " + json);
            });

            check.Run("Unit.Deserialize(Types)", () =>
            {
                var types = WebrpcJson.Deserialize<Types>("{\"ts\":\"2026-09-03T12:30:15.123456789Z\",\"tsList\":[\"2026-09-03T14:30:15+02:00\"],\"big\":12345678901234567890,\"kind\":\"USER\",\"country\":\"United States\",\"unknownField\":{\"a\":[1]},\"optTs\":null}");
                Assert.Equal(new DateTimeOffset(2026, 9, 3, 12, 30, 15, TimeSpan.Zero).AddTicks(1234567), types.Ts, "nanosecond timestamps are truncated to ticks");
                Assert.Equal(new DateTimeOffset(2026, 9, 3, 12, 30, 15, TimeSpan.Zero), types.TsList[0], "timestamp offsets");
                Assert.Equal(BigInteger.Parse("12345678901234567890"), types.Big, "bigint from a JSON number");
                Assert.Equal(Kind.User, types.Kind, "kind");
                Assert.Equal(Country.US, types.Country, "country");
                Assert.Equal(null, types.OptTs, "optTs");
                Assert.Equal("", types.Str, "missing string is empty");
                Assert.Equal(0, types.Labels.Count, "missing list is empty");
            });

            check.Run("Unit.Enums", () =>
            {
                Assert.Equal("USER", default(Kind).Value, "integer enum zero value is its 0 member");
                Assert.Equal(Kind.User, default(Kind), "default(Kind)");
                Assert.Equal("", default(Color).Value, "string enum zero value");
                Assert.True(Kind.Admin.IsKnown && !new Kind("SUPERUSER").IsKnown, "IsKnown");
                Assert.Equal(Color.red, (Color)"red", "implicit conversion from string");
                Assert.True(Color.red != Color.green && Color.red.Equals((object)Color.red), "equality");
                Assert.SequenceEqual(new[] { Kind.User, Kind.Admin }, Kind.KnownValues, "KnownValues");
                Assert.SequenceEqual(new[] { "Value", "IsKnown", "KnownValues", "Reserved", "class" }, new[] { Reserved.Value_, Reserved.IsKnown_, Reserved.KnownValues_, Reserved.Reserved_, Reserved.@class }.Select(r => r.Value), "reserved member names");
                var colors = WebrpcJson.Deserialize<List<Color>>("[\"red\",\"purple\"]");
                Assert.Equal("[\"red\",\"purple\"]", WebrpcJson.Serialize(colors), "unknown values round-trip");
            });

            check.Run("Unit.Keywords", () =>
            {
                var keywords = new Keywords { Class = "1", String = "2", Event = "3", Keywords_ = "4", ToString_ = "5", Equals_ = "6" };
                Assert.Equal("{\"class\":\"1\",\"string\":\"2\",\"event\":\"3\",\"keywords\":\"4\",\"toString\":\"5\",\"equals\":\"6\"}", WebrpcJson.Serialize(keywords), "json");
            });

            check.Run("Unit.Errors", () =>
            {
                var ex = new MismatchException("details");
                Assert.Equal(1001, ex.Code, "code");
                Assert.Equal(MismatchException.ErrorCode, ex.Code, "ErrorCode");
                Assert.Equal("Mismatch", ex.Name, "name");
                Assert.Equal("mismatch", ex.Message, "message");
                Assert.Equal(400, ex.HttpStatus, "status");
                Assert.True(ex.ToString().StartsWith("Interop.CSharp.MismatchException: Mismatch 1001: mismatch: details", StringComparison.Ordinal), "ToString: " + ex);
            });
        }

        public static async Task Run(Check check, HttpClient http, string url)
        {
            var api = new CSharpClient(url, http);

            await check.Run("CSharp.GetTypes", async () =>
            {
                var expected = Fixture();
                var got = await api.GetTypesAsync();
                Assert.Equal(expected.B, got.B, "b");
                Assert.Equal(expected.U8, got.U8, "u8");
                Assert.Equal(expected.U16, got.U16, "u16");
                Assert.Equal(expected.U32, got.U32, "u32");
                Assert.Equal(expected.U64, got.U64, "u64");
                Assert.Equal(expected.I8, got.I8, "i8");
                Assert.Equal(expected.I16, got.I16, "i16");
                Assert.Equal(expected.I32, got.I32, "i32");
                Assert.Equal(expected.I64, got.I64, "i64");
                Assert.Equal(expected.F32, got.F32, "f32");
                Assert.Equal(expected.F64, got.F64, "f64");
                Assert.Equal(expected.Str, got.Str, "str");
                Assert.SequenceEqual(expected.Bytes, got.Bytes, "bytes");
                Assert.SequenceEqual(expected.U8s, got.U8s, "u8s");
                Assert.Equal(expected.Ts, got.Ts, "ts");
                Assert.SequenceEqual(expected.TsList, got.TsList, "tsList");
                Assert.Equal(expected.Big, got.Big, "big");
                Assert.SequenceEqual(expected.BigList, got.BigList, "bigList");
                Assert.DictionaryEqual(expected.BigMap, got.BigMap, "bigMap");
                Assert.Equal(expected.Address, got.Address, "address");
                Assert.Equal(expected.Wallet, got.Wallet, "wallet");
                Assert.SequenceEqual(expected.Labels, got.Labels, "labels");
                Assert.Equal(expected.Kind, got.Kind, "kind");
                Assert.DictionaryEqual(expected.Kinds, got.Kinds, "kinds");
                Assert.Equal(expected.Country, got.Country, "country");
                Assert.Equal(expected.Reserved, got.Reserved, "reserved");
                Assert.DictionaryEqual(expected.ById, got.ById, "byId");
                Assert.SequenceEqual(expected.Nested[7], got.Nested[7], "nested");
                Assert.Equal("{\"l\":[true,null],\"n\":1,\"s\":\"two\"}", WebrpcJson.Serialize(got.Anything), "anything");
                Assert.Equal(expected.OptStr, got.OptStr, "optStr");
                Assert.Equal(expected.OptBig, got.OptBig, "optBig");
                Assert.Equal(expected.OptTs, got.OptTs, "optTs");
                Assert.Equal(expected.OptKind, got.OptKind, "optKind");
                Assert.Equal(null, got.OptList, "optList");
                Assert.Equal(expected.Renamed, got.Renamed, "renamed");
                await api.SendTypesAsync(got);
            });

            await check.Run("CSharp.SendTypes", () => api.SendTypesAsync(Fixture()));

            await check.Run("CSharp.SendTypes(mismatch)", async () =>
            {
                var types = Fixture();
                types.Big += 1;
                var ex = await Assert.Throws<MismatchException>(() => api.SendTypesAsync(types));
                Assert.True(ex.Cause != null && ex.Cause.Contains("Big"), "cause: " + ex.Cause);
            });

            await check.Run("CSharp.GetUnknownEnum", async () =>
            {
                var resp = await api.GetUnknownEnumAsync();
                Assert.Equal("purple", resp.Color.Value, "color");
                Assert.True(!resp.Color.IsKnown, "color.IsKnown");
                Assert.SequenceEqual(new[] { Color.red, new Color("blue") }, resp.Colors, "colors");
                Assert.DictionaryEqual(new Dictionary<Color, uint> { [Color.green] = 1, ["pink"] = 2 }, resp.ByColor, "byColor");
                await api.SendUnknownEnumAsync(resp.Color, resp.Colors);
            });

            await check.Run("CSharp.Keywords", async () =>
            {
                var resp = await api.KeywordsAsync("c", "t", new List<string> { "p1", "p2" });
                Assert.Equal("class=c cancellationToken=t params=p1,p2", resp.Event, "event");
                Assert.Equal("keywords", resp.Keywords.Keywords_, "keywords.keywords");
                Assert.Equal("toString", resp.Keywords.ToString_, "keywords.toString");
                var omitted = await api.KeywordsAsync("c", "t", null, CancellationToken.None);
                Assert.Equal("class=c cancellationToken=t params=", omitted.Event, "event (params omitted)");
            });

            await check.Run("CSharp.Shadow", async () =>
            {
                Assert.Equal("a,b,c,d", await api.ShadowAsync("a", "b", "c", "d"), "echo");
            });

            await check.Run("CSharp.Optional", async () =>
            {
                var ts = new DateTimeOffset(2026, 1, 2, 3, 4, 5, TimeSpan.Zero);
                var set = await api.OptionalAsync("s", 18446744073709551615, ts);
                Assert.Equal("s", set.Str, "str");
                Assert.Equal(18446744073709551615, set.Num, "num");
                Assert.Equal(ts, set.Ts, "ts");
                var unset = await api.OptionalAsync(null, null, null);
                Assert.True(unset.Str == null && unset.Num == null && unset.Ts == null, "unset outputs are null");
            });

            await check.Run("CSharp.Deprecated", async () =>
            {
#pragma warning disable CS0618
                Assert.True(await api.DeprecatedAsync(), "ok");
#pragma warning restore CS0618
                var obsolete = typeof(ICSharpClient).GetMethod(nameof(ICSharpClient.DeprecatedAsync))!.GetCustomAttributes(typeof(ObsoleteAttribute), false).Cast<ObsoleteAttribute>().Single();
                Assert.Equal("Use Optional() instead.", obsolete.Message, "[Obsolete] message");
            });

            await check.Run("CSharp.GetShadowed", async () =>
            {
                var resp = await api.GetShadowedAsync();
                Assert.Equal(1ul, resp.Task.Id, "task.id");
                Assert.Equal(2u, resp.Version.Major, "version.major");
                Assert.SequenceEqual(new[] { "three" }, resp.List.Items, "list.items");
            });

            await check.Run("CSharp.Countdown(stream)", async () =>
            {
                var messages = new List<CSharpClient.CountdownResponse>();
                await foreach (var message in api.CountdownAsync(3))
                {
                    messages.Add(message);
                }
                Assert.SequenceEqual(new uint[] { 3, 2, 1 }, messages.Select(m => m.N), "n");
                Assert.Equal(new DateTimeOffset(2026, 1, 1, 0, 0, 3, TimeSpan.Zero), messages[0].At, "at");
            });

            await check.Run("CSharp.Watch(succinct stream)", async () =>
            {
                var seqs = new List<uint>();
                await foreach (var message in api.WatchAsync(new WatchRequest { Key = "k", Count = 4 }))
                {
                    Assert.Equal("k", message.Key, "key");
                    seqs.Add(message.Seq);
                }
                Assert.SequenceEqual(new uint[] { 1, 2, 3, 4 }, seqs, "seq");
            });

            await check.Run("CSharp.CountdownFail(stream error)", async () =>
            {
                var got = new List<uint>();
                var ex = await Assert.Throws<CountdownFailedException>(async () =>
                {
                    await foreach (var n in api.CountdownFailAsync(5))
                    {
                        got.Add(n);
                    }
                });
                Assert.SequenceEqual(new uint[] { 5 }, got, "messages before the error");
                Assert.Equal("failed after 5", ex.Cause, "cause");
                Assert.Equal(500, ex.HttpStatus, "status");
            });

            await check.Run("CSharp.Countdown(cancel)", async () =>
            {
                using var cts = new CancellationTokenSource();
                await Assert.Throws<OperationCanceledException>(async () =>
                {
                    await foreach (var message in api.CountdownAsync(3, cts.Token))
                    {
                        cts.Cancel();
                    }
                });
            });

            await check.Run("CSharp.GetByName(REST)", async () =>
            {
                var resp = await api.GetByNameAsync("a b/c?d", "x&y");
                Assert.Equal("a b/c?d", resp.Name, "name");
                Assert.Equal("x&y", resp.Suffix, "suffix");
                Assert.Equal(null, (await api.GetByNameAsync("n", null)).Suffix, "suffix (omitted)");
            });

            await check.Run("CSharp.GetByName(invalid path param)", async () =>
            {
                foreach (var name in new[] { "", ".", ".." })
                {
                    var ex = await Assert.Throws<WebrpcBadRequestException>(() => api.GetByNameAsync(name, null), $"name \"{name}\"");
                    Assert.True(ex.Cause != null && ex.Cause.Contains("invalid path param 'name'"), "cause: " + ex.Cause);
                }
            });

            await check.Run("CSharp.Headers", async () =>
            {
                using var custom = new HttpClient(new HeaderHandler("X-Custom", "custom value") { InnerHandler = new HttpClientHandler() });
                var resp = await new CSharpClient(url, custom).HeadersAsync();
                Assert.Equal(Webrpc.HeaderValue, resp.Webrpc, "Webrpc header");
                Assert.Equal("custom value", resp.Custom, "X-Custom header (DelegatingHandler)");
            });

            await check.Run("CSharp.UploadDoc", async () =>
            {
                var content = Encoding.UTF8.GetBytes("hello");
                var resp = await api.UploadDocAsync(new Doc { Title = "T", File = WebrpcFile.FromBytes(content, "naïve \"résumé\".txt", "text/plain") });
                Assert.Equal("naïve \"résumé\".txt", resp.Name, "name");
                Assert.Equal("text/plain", resp.ContentType, "contentType");
                Assert.Equal((ulong)content.Length, resp.Size, "size");
                Assert.Equal("T", resp.Title, "title");
            });

            await check.Run("CSharp.Attach(succinct upload)", async () =>
            {
                var resp = await api.AttachAsync(new AttachRequest
                {
                    Id = 9,
                    Files = { WebrpcFile.FromBytes(new byte[] { 1, 2, 3 }, "a.bin"), new WebrpcFile(new NonSeekableStream(new byte[] { 4 }), "b.bin") },
                });
                Assert.Equal(9ul, resp.Id, "id");
                Assert.SequenceEqual(new[] { "a.bin", "b.bin" }, resp.Names, "names");
                Assert.SequenceEqual(new ulong[] { 3, 1 }, resp.Sizes, "sizes");
            });

            await check.Run("CSharp.DownloadDoc", async () =>
            {
                using var file = await api.DownloadDocAsync("naïve résumé.txt");
                Assert.Equal("naïve résumé.txt", file.Name, "name (RFC 5987)");
                Assert.Equal("text/plain; charset=utf-8", file.ContentType, "contentType");
                Assert.Equal((long?)Encoding.UTF8.GetByteCount("doc:naïve résumé.txt"), file.Size, "size");
                Assert.Equal("doc:naïve résumé.txt", Encoding.UTF8.GetString(await file.ReadAllBytesAsync()), "content");
            });

            await check.Run("CSharp.Cancellation", async () =>
            {
                using var cts = new CancellationTokenSource();
                cts.Cancel();
                await Assert.Throws<OperationCanceledException>(() => api.GetTypesAsync(cts.Token));
            });

            // The caller's token interrupts a stalled error response body.
            await check.Run("CSharp.Cancellation(stalled error body)", async () =>
            {
                var stalled = new CSharpClient(url + "/stall-error", http);
                var calls = new (string Name, Func<CancellationToken, Task> Call)[]
                {
                    ("download", async ct => (await stalled.DownloadDocAsync("x", ct)).Dispose()),
                    ("stream", async ct => { await foreach (var _ in stalled.CountdownAsync(1, ct)) { } }),
                    ("unary", ct => stalled.GetTypesAsync(ct)),
                };
                foreach (var (name, call) in calls)
                {
                    using var cts = new CancellationTokenSource(TimeSpan.FromMilliseconds(300));
                    var pending = call(cts.Token);
                    var finished = await Task.WhenAny(pending, Task.Delay(TimeSpan.FromSeconds(5)));
                    Assert.True(finished == pending, name + ": still pending 5s after cancellation");
                    await Assert.Throws<OperationCanceledException>(() => pending, name);
                }
            });

            await check.Run("CSharp.RequestFailed", async () =>
            {
                var ex = await Assert.Throws<WebrpcRequestFailedException>(() => new CSharpClient("http://127.0.0.1:1", http).GetTypesAsync());
                Assert.Equal(-1, ex.Code, "code");
                Assert.True(ex.InnerException is HttpRequestException, "inner: " + ex.InnerException);
            });

            await check.Run("CSharp.BadResponse(non-webrpc error)", async () =>
            {
                var ex = await Assert.Throws<WebrpcBadResponseException>(() => new CSharpClient(url + "/html-error", http).GetTypesAsync());
                Assert.Equal(502, ex.HttpStatus, "status");
                Assert.True(ex.Cause != null && ex.Cause.Contains("502 Bad Gateway"), "cause: " + ex.Cause);
            });

            await check.Run("CSharp.BadRoute", async () =>
            {
                var ex = await Assert.Throws<WebrpcException>(() => new CSharpClient(url + "/nope", http).GetTypesAsync());
                Assert.Equal(404, ex.HttpStatus, "status");
            });
        }

        private sealed class HeaderHandler : DelegatingHandler
        {
            private readonly string _name;
            private readonly string _value;

            public HeaderHandler(string name, string value)
            {
                _name = name;
                _value = value;
            }

            protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
            {
                request.Headers.Add(_name, _value);
                return base.SendAsync(request, cancellationToken);
            }
        }
    }

    // NonSeekableStream hides the length of its content, like a network stream.
    internal sealed class NonSeekableStream : MemoryStream
    {
        public NonSeekableStream(byte[] content) : base(content, false)
        {
        }

        public override bool CanSeek => false;

        public override long Length => throw new NotSupportedException();

        public override long Position
        {
            get => throw new NotSupportedException();
            set => throw new NotSupportedException();
        }
    }
}
