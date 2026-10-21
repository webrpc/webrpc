using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace Interop.Tests
{
    // Runs the C# client interoperability tests; see ../test.sh.
    //
    // Usage: dotnet run -- <TestApi URL> <TestApiRest URL> <CSharp URL> <Example URL>
    internal static class Program
    {
        public static async Task<int> Main(string[] args)
        {
            if (args.Length != 4)
            {
                Console.Error.WriteLine("usage: <TestApi URL> <TestApiRest URL> <CSharp URL> <Example URL>");
                return 2;
            }
#if NEWTONSOFT
            Console.WriteLine("C# client interop tests (-json=newtonsoft)");
#else
            Console.WriteLine("C# client interop tests (-json=stj)");
#endif
            var check = new Check();
            using var http = new HttpClient();

            CSharpTests.RunUnit(check);
            await TestApiTests.Run(check, http, args[0]);
            await TestApiRestTests.Run(check, http, args[1]);
            await CSharpTests.Run(check, http, args[2]);
            await ExampleTests.Run(check, http, args[3]);

            if (check.Failures.Count > 0)
            {
                Console.Error.WriteLine($"\n{check.Failures.Count} failed:");
                foreach (var failure in check.Failures)
                {
                    Console.Error.WriteLine(failure + "\n");
                }
                return 1;
            }
            Console.WriteLine("PASS");
            return 0;
        }
    }
}
