using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Interop.Test;

namespace Interop.Tests
{
    // TestApiTests ports the webrpc interoperability client suite
    // (tests/client/client.go RunTests) to the generated C# client, run
    // against `webrpc-test -server`.
    internal static class TestApiTests
    {
        // File fixtures of the test server, with non-UTF8 bytes to catch encoding bugs.
        private const string FixtureFileName = "test.bin";
        private static readonly byte[] FixtureFileContent = Bytes("webrpc interop test file\x00\x01\x02\xff\xfe binary content");
        private static readonly byte[][] FixtureFilesContent = { Bytes("first file"), Bytes("second file\x00\xff") };

        private static byte[] Bytes(string latin1) => Encoding.Latin1.GetBytes(latin1);

        public static async Task Run(Check check, HttpClient http, string url)
        {
            var api = new TestApiClient(url, http);

            await check.Run("TestApi.GetEmpty", () => api.GetEmptyAsync());

            await check.Run("TestApi.GetError", async () =>
            {
                var ex = await Assert.Throws<WebrpcException>(() => api.GetErrorAsync());
                Assert.Equal(0, ex.Code, "code");
            });

            await check.Run("TestApi.GetOne+SendOne", async () =>
            {
                var one = await api.GetOneAsync();
                Assert.Equal(1, one.Id, "id");
                Assert.Equal("one", one.Name, "name");
                await api.SendOneAsync(one);
            });

            await check.Run("TestApi.GetMulti+SendMulti", async () =>
            {
                var multi = await api.GetMultiAsync();
                Assert.Equal("three", multi.Three.Name, "three");
                await api.SendMultiAsync(multi.One, multi.Two, multi.Three);
            });

            await check.Run("TestApi.GetComplex+SendComplex", async () =>
            {
                var complex = await api.GetComplexAsync();
                Assert.SequenceEqual(new long[] { 1, 2, 3, 4534643543 }, complex.NumsList, "numsList");
                Assert.Equal("John-Doe", complex.MapOfUsers["admin"].Username, "mapOfUsers.admin.USERNAME");
                Assert.Equal(Status.AVAILABLE, complex.Status, "status");
                await api.SendComplexAsync(complex);
            });

            await check.Run("TestApi.SendComplex(constructed)", () => api.SendComplexAsync(new Complex
            {
                Meta = { ["1"] = "23", ["2"] = 24 },
                MetaNestedExample = { ["1"] = new Dictionary<string, uint> { ["2"] = 1 } },
                NamesList = { "John", "Alice", "Jakob" },
                NumsList = { 1, 2, 3, 4534643543 },
                DoubleArray = { new List<string> { "testing" }, new List<string> { "api" } },
                ListOfMaps = { new Dictionary<string, uint> { ["john"] = 1, ["alice"] = 2, ["Jakob"] = 251 } },
                ListOfUsers = { new User { Id = 1, Username = "John-Doe", Role = "admin" } },
                MapOfUsers = { ["admin"] = new User { Id = 1, Username = "John-Doe", Role = "admin" } },
                User = new User { Id = 1, Username = "John-Doe", Role = "admin" },
                Status = Status.AVAILABLE,
            }));

            await check.Run("TestApi.GetEnumList", async () =>
            {
                var list = await api.GetEnumListAsync();
                Assert.SequenceEqual(new[] { Status.AVAILABLE, Status.NOT_AVAILABLE }, list, "list");
            });

            await check.Run("TestApi.GetEnumMap", async () =>
            {
                var map = await api.GetEnumMapAsync();
                Assert.DictionaryEqual(new Dictionary<Access, ulong> { [Access.READ] = 1, [Access.WRITE] = 2 }, map, "map");
            });

            await check.Run("TestApi.UploadFile", async () =>
            {
                var size = await api.UploadFileAsync(FixtureFileName, WebrpcFile.FromBytes(FixtureFileContent, FixtureFileName));
                Assert.Equal((ulong)FixtureFileContent.Length, size, "size");
            });

            await check.Run("TestApi.UploadFile(non-seekable stream)", async () =>
            {
                using var body = new NonSeekableStream(FixtureFileContent);
                var size = await api.UploadFileAsync(FixtureFileName, new WebrpcFile(body, FixtureFileName));
                Assert.Equal((ulong)FixtureFileContent.Length, size, "size");
            });

            await check.Run("TestApi.UploadFiles", async () =>
            {
                var files = FixtureFilesContent.Select((content, i) => WebrpcFile.FromBytes(content, $"file-{i}.bin")).ToList();
                var count = await api.UploadFilesAsync(files);
                Assert.Equal((uint)FixtureFilesContent.Length, count, "count");
            });

            await check.Run("TestApi.DownloadFile", async () =>
            {
                using var file = await api.DownloadFileAsync(FixtureFileName);
                Assert.SequenceEqual(FixtureFileContent, await file.ReadAllBytesAsync(), "content");
                Assert.Equal(FixtureFileName, file.Name, "name");
                Assert.Equal("application/octet-stream", file.ContentType, "contentType");
                Assert.Equal((long?)FixtureFileContent.Length, file.Size, "size");
            });

            await check.Run("TestApi.EchoFile", async () =>
            {
                using var file = await api.EchoFileAsync(WebrpcFile.FromBytes(FixtureFileContent, FixtureFileName));
                Assert.SequenceEqual(FixtureFileContent, await file.ReadAllBytesAsync(), "content");
            });

            await check.Run("TestApi.DownloadFile(error)", async () =>
            {
                var ex = await Assert.Throws<UnexpectedValueException>(() => api.DownloadFileAsync("nope.bin"));
                Assert.Equal(400, ex.HttpStatus, "status");
            });

            var schemaErrors = new (int Code, Type Type, string Name, string Message, int Status, string? Cause)[]
            {
                (0, typeof(WebrpcEndpointException), "WebrpcEndpoint", "endpoint error", 400, "failed to read file: unexpected EOF"),
                (1, typeof(UnauthorizedException), "Unauthorized", "unauthorized", 401, "failed to verify JWT token"),
                (2, typeof(ExpiredTokenException), "ExpiredToken", "expired token", 401, null),
                (3, typeof(InvalidTokenException), "InvalidToken", "invalid token", 401, null),
                (4, typeof(DeactivatedException), "Deactivated", "account deactivated", 403, null),
                (5, typeof(ConfirmAccountException), "ConfirmAccount", "confirm your email", 403, null),
                (6, typeof(AccessDeniedException), "AccessDenied", "access denied", 403, null),
                (7, typeof(MissingArgumentException), "MissingArgument", "missing argument", 400, null),
                (8, typeof(UnexpectedValueException), "UnexpectedValue", "unexpected value", 400, null),
                (100, typeof(RateLimitedException), "RateLimited", "too many requests", 429, "1000 req/min exceeded"),
                (101, typeof(DatabaseDownException), "DatabaseDown", "service outage", 503, null),
                (102, typeof(ElasticDownException), "ElasticDown", "search is degraded", 503, null),
                (103, typeof(Interop.Test.NotImplementedException), "NotImplemented", "not implemented", 501, null),
                (200, typeof(UserNotFoundException), "UserNotFound", "user not found", 400, null),
                (201, typeof(UserBusyException), "UserBusy", "user busy", 400, null),
                (202, typeof(InvalidUsernameException), "InvalidUsername", "invalid username", 400, null),
                (300, typeof(FileTooBigException), "FileTooBig", "file is too big (max 1GB)", 400, null),
                (301, typeof(FileInfectedException), "FileInfected", "file is infected", 400, null),
                (302, typeof(FileTypeException), "FileType", "unsupported file type", 400, ".wav is not supported"),
            };
            foreach (var tc in schemaErrors)
            {
                await check.Run($"TestApi.GetSchemaError({tc.Code})", async () =>
                {
                    var ex = await Assert.Throws<WebrpcException>(() => api.GetSchemaErrorAsync(tc.Code));
                    Assert.Equal(tc.Type, ex.GetType(), "type");
                    Assert.Equal(tc.Code, ex.Code, "code");
                    Assert.Equal(tc.Name, ex.Name, "name");
                    Assert.Equal(tc.Message, ex.Message, "message");
                    Assert.Equal(tc.Status, ex.HttpStatus, "status");
                    Assert.Equal(tc.Cause, ex.Cause, "cause");
                });
            }

            await check.Run("TestApi.GetSchemaError(none)", () => api.GetSchemaErrorAsync(-1));
        }
    }
}
