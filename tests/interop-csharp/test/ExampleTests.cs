using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Security.Cryptography;
using System.Threading.Tasks;
using Interop.Example;

namespace Interop.Tests
{
    // ExampleTests runs the generated C# client against the golang-basics
    // example server (_examples/golang-basics): succinct methods, enums with
    // + json values, file metadata, deprecated methods.
    internal static class ExampleTests
    {
        public static async Task Run(Check check, HttpClient http, string url)
        {
            var api = new ExampleClient(url, http);

            await check.Run("Example.Ping", () => api.PingAsync());

            await check.Run("Example.GetUserV2", async () =>
            {
                var resp = await api.GetUserV2Async(new GetUserRequest { UserId = 7, Prefs = { ["lang"] = "en" } });
                Assert.Equal(200u, resp.Code, "code");
                Assert.Equal(7ul, resp.User.Id, "user.id");
                Assert.Equal("hihi", resp.User.Username, "user.USERNAME");
                Assert.Equal(Kind.Admin, resp.User.Kind, "user.kind");
                Assert.Equal("ADMIN", resp.User.Kind.Value, "user.kind wire value");
                Assert.Equal(Intent.openSession, resp.User.Intent, "user.intent");
                Assert.Equal("https://www.google.com/images/john-doe.jpg", resp.Profile, "profile");
            });

            await check.Run("Example.GetUserV2(not found)", async () =>
            {
                var ex = await Assert.Throws<UserNotFoundException>(() => api.GetUserV2Async(new GetUserRequest { UserId = 31337 }));
                Assert.Equal(1000, ex.Code, "code");
                Assert.Equal(404, ex.HttpStatus, "status");
                Assert.Equal("unknown user id 31337", ex.Cause, "cause");
            });

            // The server rejects calls of methods annotated + @deprecated.
            await check.Run("Example.GetUser(deprecated)", async () =>
            {
#pragma warning disable CS0618
                await Assert.Throws<WebrpcException>(() => api.GetUserAsync(new GetUserRequest { UserId = 7 }));
#pragma warning restore CS0618
            });

            await check.Run("Example.FindUser", async () =>
            {
                var resp = await api.FindUserAsync(new SearchFilter { Q = "jane" });
                Assert.Equal("jane", resp.Name, "name");
                Assert.Equal("jane", resp.User.Username, "user.USERNAME");
                Assert.Equal(Kind.User, resp.User.Kind, "user.kind (zero value)");
            });

            await check.Run("Example.GetIntents", async () =>
            {
                Assert.SequenceEqual(new[] { Intent.openSession, Intent.closeSession, Intent.validateSession }, await api.GetIntentsAsync(), "intents");
            });

            await check.Run("Example.CountIntents", async () =>
            {
                var count = await api.CountIntentsAsync(1);
                Assert.DictionaryEqual(new Dictionary<Intent, uint> { [Intent.openSession] = 1, [Intent.closeSession] = 2, [Intent.validateSession] = 3 }, count, "count");
            });

            await check.Run("Example.UploadAvatar+DownloadAvatar", async () =>
            {
                var avatar = RandomNumberGenerator.GetBytes(1024);
                var uploaded = await api.UploadAvatarAsync(new UploadAvatarRequest { UserId = 42, Avatar = WebrpcFile.FromBytes(avatar, "me.png", "image/png") });
                Assert.Equal((ulong)avatar.Length, uploaded.Size, "size");

                using var download = await api.DownloadAvatarAsync(new DownloadAvatarRequest { UserId = 42 });
                Assert.Equal("me.png", download.Name, "name");
                Assert.Equal("image/png", download.ContentType, "contentType");
                Assert.Equal((long?)avatar.Length, download.Size, "size");
                Assert.SequenceEqual(avatar, await download.ReadAllBytesAsync(), "content");
            });

            await check.Run("Example.DownloadAvatar(not found)", async () =>
            {
                var ex = await Assert.Throws<UserNotFoundException>(() => api.DownloadAvatarAsync(new DownloadAvatarRequest { UserId = 999999 }));
                Assert.Equal("no avatar for user 999999", ex.Cause, "cause");
            });
        }
    }
}
